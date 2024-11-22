package gologging

import "go.uber.org/zap"

func (c *config) start() error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	c.logger = logger
	return nil
}

func (c *config) flush() error {
	return c.logger.Sync()
}
