package artifacts

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func SetupRMQConnection(retryFunc func() error, exchangeName, exchangeType string) (*amqp.Connection, *amqp.Channel, error) {
	// Establish connection with RabbitMQ
	rmqConn, err := dialRMQ()
	if err != nil {
		return nil, nil, err
	}

	// Spawn a goroutine to handle any RabbitMQ related errors
	go handleRabbitMQErrors(rmqConn, retryFunc)

	// Create Channel
	rmqChannel, err := createChannel(rmqConn)
	if err != nil {
		return nil, nil, err
	}

	return rmqConn, rmqChannel, declareRealtimeExchange(rmqChannel, exchangeName, exchangeType)
}

// Open channel with janus virtualhost connection
func createChannel(rmqConn *amqp.Connection) (*amqp.Channel, error) {
	rmqChan, err := rmqConn.Channel()
	return rmqChan, err
}

// Declare exchange `janus-realtime-notifications` to listen for realtime notification pushes from asgard
func declareRealtimeExchange(rmqChan *amqp.Channel, exchangeName, exchangeType string) (err error) {
	err = rmqChan.ExchangeDeclare(
		// viper.GetString("rmq.realtimeExchange"),     // name of the exchange
		// viper.GetString("rmq.realtimeExchangeType"), // type
		exchangeName, // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return fmt.Errorf("Error declaring exchange: janus-realtime-notifications on vhost: janus. Error: %s", err)
	}
	return nil
}

// Looks out for any RabbitMQ errors/closure and re-establishes connection
// and initializes the realtime consumer
func handleRabbitMQErrors(rmqConn *amqp.Connection, consumerInitFn func() error) {
	rmqError := <-rmqConn.NotifyClose(make(chan *amqp.Error))
	if rmqError == nil {
		return
	}
	// RabbitMQ connection closure triggered. Log and send a sentry alert
	rmqCloseErr := fmt.Errorf("RMQ Connection closed. Code: %d: Reason: %s", rmqError.Code, rmqError.Reason)
	log.Println(rmqCloseErr)

	// Check if the error is a "ConnectionError" or a "ChannelError".
	// Other error codes don't trigger retrial
	if isBrokerError(rmqError.Code) {
		for {
			if retryError := retryBrokerConnections(consumerInitFn); retryError == nil {
				break
			}
		}
	}
}

func retryBrokerConnections(consumerSetupFn func() error) error {
	// Setup consumers
	if reloadErr := consumerSetupFn(); reloadErr != nil {
		log.Println("Failed to restore RMQ consumers: ", reloadErr)
		// raven.CaptureErrorAndWait(fmt.Errorf("Failed to restore RMQ consumers: %v", reloadErr), nil)
		return reloadErr
	}
	return nil
}

func isBrokerError(code int) bool {
	switch code {
	case
		// Channel Errors
		311, // amqp.ContentTooLarge
		313, // amqp.NoConsumers
		403, // amqp.AccessRefused
		404, // amqp.NotFound
		405, // amqp.ResourceLocked
		406: // amqp.PreconditionFailed
		return true

	case
		// Connection Errors
		320, // amqp.ConnectionForced
		402, // amqp.InvalidPath
		501, // amqp.FrameError
		502, // amqp.SyntaxError
		503, // amqp.CommandInvalid
		504, // amqp.ChannelError
		505, // amqp.UnexpectedFrame
		506, // amqp.ResourceError
		530, // amqp.NotAllowed
		540, // amqp.NotImplemented
		541: // amqp.InternalError
		return true

	default:
		return false
	}
}
