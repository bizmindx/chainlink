package logger

import (
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting/types"
)

// ocrLogger is an implemenation of the Logger interface for OCR
// See: https://github.com/smartcontractkit/offchain-reporting-design/blob/master/prototype/offchainreporting/types/logger.go#L3

var _ ocrtypes.Logger = &ocrLogger{}

type ocrLogger struct {
	internal *Logger
}

func NewOCRLogger(internal *Logger) ocrtypes.Logger {
	return &ocrLogger{
		internal: internal,
	}
}

func (ol *ocrLogger) Trace(msg string, fields ocrtypes.LogFields) {
	ol.internal.Tracew(msg, toKeysAndValues(fields))
}

func (ol *ocrLogger) Debug(msg string, fields ocrtypes.LogFields) {
	ol.internal.Debugw(msg, toKeysAndValues(fields))
}

func (ol *ocrLogger) Info(msg string, fields ocrtypes.LogFields) {
	ol.internal.Infow(msg, toKeysAndValues(fields))
}

func (ol *ocrLogger) Warn(msg string, fields ocrtypes.LogFields) {
	ol.internal.Warnw(msg, toKeysAndValues(fields))
}

func (ol *ocrLogger) Error(msg string, fields ocrtypes.LogFields) {
	ol.internal.Errorw(msg, toKeysAndValues(fields))
}

func (ol *ocrLogger) Trace(msg string, fields ocrtypes.LogFields) {
	// @@TODO ???
	ol.internal.Debugw(msg, toKeysAndValues(fields))
}

// Helpers

func toKeysAndValues(fields ocrtypes.LogFields) []interface{} {
	out := []interface{}{}
	for key, val := range fields {
		out = append(out, key, val)
	}
	return out
}
