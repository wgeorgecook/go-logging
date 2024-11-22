package gologging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logOption func(*fields)

// WithRequestID sets the provided request_id
// on the Zap structured logger
func WithRequestID(requestID string) logOption {
	return func(f *fields) {
		f.requestId = requestID
	}
}

// WithError sets the provided error on the
// Zap structured logger
func WithError(err error) logOption {
	return func(f *fields) {
		f.err = err
	}
}

type fields struct {
	err       error
	requestId string
}

func setOpts(opts ...logOption) *fields {
	var f = &fields{}
	for _, opt := range opts {
		opt(f)
	}

	return f
}

func setZapFields(f *fields) []zapcore.Field {
	var zapFields = make([]zapcore.Field, 0, 2)
	if f.requestId != "" {
		zapFields = append(zapFields, zap.String("request_id", f.requestId))
	}
	if f.err != nil {
		zapFields = append(zapFields, zap.Error(f.err))
	}

	return zapFields
}

// Info logs the provided LogOptions using the Zap structured
// logger, provided the level provided to Init() is <= INFO.
func Info(msg string, opts ...logOption) {
	if c.level > INFO {
		return
	}

	f := setOpts(opts...)
	zapFields := setZapFields(f)
	c.logger.Info(msg, zapFields...)

}

// Debug logs the provided LogOptions using the Zap structured
// logger, provided the level provided to Init() is <= DEBUG.
func Debug(msg string, opts ...logOption) {
	if c.level > DEBUG {
		return
	}

	f := setOpts(opts...)
	zapFields := setZapFields(f)
	c.logger.Debug(msg, zapFields...)
}

// Error logs the provided error and LogOptions using the Zap structured
// logger, provided the level provided to Init() is <= ERROR.
func Error(msg string, err error, opts ...logOption) {
	if c.level > ERROR {
		return
	}
	f := setOpts(opts...)
	if f.err == nil {
		f.err = err
	}
	zapFields := setZapFields(f)
	c.logger.Error(msg, zapFields...)
}

// Fatal logs the provided LogOptions using the Zap structured
// logger, provided the level provided to Init() is <= FATAL.
func Fatal(msg string, opts ...logOption) {
	if c.level > FATAL {
		return
	}
	f := setOpts(opts...)
	zapFields := setZapFields(f)
	c.logger.Fatal(msg, zapFields...)
}
