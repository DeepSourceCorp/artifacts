package artifacts

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
)

// Utility function to retry any routine
func Retry(attempts int, sleep time.Duration, f func() error) (err error) {
	for i := 0; ; i++ {
		err = f()
		if err == nil {
			return
		}
		if i >= (attempts - 1) {
			break
		}
		time.Sleep(sleep)
		log.Println("Retrying job after error:", err)
	}
	sentry.CaptureException(err)
	return fmt.Errorf("Failed to trigger the job aftert %d attempts, last error: %s", attempts, err)
}

// Get a HTTP client for interacting with k8s REST API
func GetNewHTTPClient(certFilePath string) (*http.Client, error) {
	var httpClient *http.Client

	// if auth is true, send request with the certificate
	caCert, err := ioutil.ReadFile(certFilePath)
	if err != nil {
		return httpClient, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	httpClient = &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
				RootCAs:            caCertPool,
				MinVersion:         tls.VersionTLS12,
				MaxVersion:         0,
			},
		},
	}

	return httpClient, nil
}
