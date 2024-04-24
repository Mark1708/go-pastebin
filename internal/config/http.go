package config

import (
	"net"
	"os"
	"time"

	"github.com/pkg/errors"
)

const (
	httpHostEnvName              = "HTTP_HOST"
	httpPortEnvName              = "HTTP_PORT"
	httpReadTimeoutEnvName       = "HTTP_TIMEOUT_READ"
	httpReadHeaderTimeoutEnvName = "HTTP_TIMEOUT_READ_HEADER"
	httpWriteTimeoutEnvName      = "HTTP_TIMEOUT_WRITE" //nolint:gosec // Так надо
	httpIdleTimeoutEnvName       = "HTTP_TIMEOUT_IDLE"
)

type HTTPConfig interface {
	Address() string
	ReadTimeout() time.Duration
	ReadHeaderTimeout() time.Duration
	WriteTimeout() time.Duration
	IdleTimeout() time.Duration
}

type httpConfig struct {
	host              string
	port              string
	readTimeout       time.Duration
	readHeaderTimeout time.Duration
	writeTimeout      time.Duration
	idleTimeout       time.Duration
}

func NewHTTPConfig() (HTTPConfig, error) {
	host := os.Getenv(httpHostEnvName)

	port := os.Getenv(httpPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("http port not found")
	}

	readTimeout, err := parseDurationEnv("readTimeout", httpReadTimeoutEnvName)
	if err != nil {
		return nil, err
	}

	readHeaderTimeout, err := parseDurationEnv("readHeaderTimeout", httpReadHeaderTimeoutEnvName)
	if err != nil {
		return nil, err
	}

	writeTimeout, err := parseDurationEnv("writeTimeout", httpWriteTimeoutEnvName)
	if err != nil {
		return nil, err
	}

	idleTimeout, err := parseDurationEnv("idleTimeout", httpIdleTimeoutEnvName)
	if err != nil {
		return nil, err
	}

	return &httpConfig{
		host:              host,
		port:              port,
		readTimeout:       readTimeout,
		readHeaderTimeout: readHeaderTimeout,
		writeTimeout:      writeTimeout,
		idleTimeout:       idleTimeout,
	}, nil
}

func (cfg *httpConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}

func (cfg *httpConfig) ReadTimeout() time.Duration {
	return cfg.readTimeout
}

func (cfg *httpConfig) ReadHeaderTimeout() time.Duration {
	return cfg.readHeaderTimeout
}

func (cfg *httpConfig) WriteTimeout() time.Duration {
	return cfg.writeTimeout
}

func (cfg *httpConfig) IdleTimeout() time.Duration {
	return cfg.idleTimeout
}

func parseDurationEnv(variableName string, envName string) (time.Duration, error) {
	durationStr := os.Getenv(envName)
	if len(durationStr) == 0 {
		return time.Millisecond, errors.New("idleTimeout not found")
	}
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return time.Millisecond, errors.New(variableName + " has incorrect format")
	}
	return duration, nil
}
