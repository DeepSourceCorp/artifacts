// Publisher is a message publisher that publishes resutls to Asgard.
package publisher

import "context"

// Extra information passed to the Publish() method.  This is currently only
// used to pass the authorization token in HTTP mode.
type Extra struct {
	Token string `json:"token"`
}

type Payload interface {
	Bytes() ([]byte, error)
}

type Publisher interface {
	Publish(ctx context.Context, payload Payload, extra *Extra) error
}
