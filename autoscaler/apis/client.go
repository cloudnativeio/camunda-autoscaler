package apis

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"

	"github.com/motemen/go-loghttp"
	"github.com/sirupsen/logrus"
	"github.com/trx35479/camunda-autoscaler/autoscaler/log"
)

var logger = log.NewLogger()

type Client struct {
	HTTPClient  *http.Client
	Certificate []byte
}

func (c *Client) client(msg string) (*http.Client, error) {
	var caCertPool *x509.CertPool
	if c.Certificate != nil {
		caCertPool = x509.NewCertPool()
		if ok := caCertPool.AppendCertsFromPEM(c.Certificate); !ok {
			return nil, fmt.Errorf("could not decode cert")
		}
	} else {
		caCertPool = nil
	}

	c.HTTPClient = &http.Client{
		Transport: &loghttp.Transport{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs: caCertPool,
				},
			},
			LogRequest: func(req *http.Request) {
				logger.WithFields(logrus.Fields{
					"protocol":    req.Proto,
					"method":      req.Method,
					"user_agent":  req.UserAgent(),
					"remote_host": req.Host,
					"path":        req.URL.Path,
				}).Info(fmt.Sprintf("%s-Req", msg))
			},
			LogResponse: func(resp *http.Response) {
				logger.WithFields(logrus.Fields{
					"protocol":    resp.Proto,
					"method":      resp.Request.Method,
					"user_agent":  resp.Request.UserAgent(),
					"status_code": resp.StatusCode,
					"remote_host": resp.Request.Host,
					"path":        resp.Request.URL.Path,
				}).Info(fmt.Sprintf("%s-Resp", msg))
			},
		},
	}

	return c.HTTPClient, nil
}
