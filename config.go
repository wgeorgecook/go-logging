package gologging

import "go.uber.org/zap"

type lvl int

const (
	_ lvl = iota
	DEBUG
	INFO
	ERROR
	FATAL
)

type config struct {
	level  lvl
	logger *zap.Logger
}

var c *config

type configOption func(*config)

// WithLevel defines which level logs
// are written. Renders all log functions
// no-op if they are lower level than provided.
// Must call Stop() before exiting application.
func WithLevel(l lvl) configOption {
	return func(c *config) {
		c.level = l
	}
}

// Init starts a Zap production logger
// either with the provided level or INFO level.
func Init(opts ...configOption) error {
	if c != nil {
		return nil
	}

	c = &config{
		level: INFO,
	}
	for _, opt := range opts {
		opt(c)
	}

	return c.start()
}

// Stop flushes the underlying
// Zap logger.
func Stop() error {
	return c.flush()
}
