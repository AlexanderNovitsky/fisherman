package logger_test

import (
	"bytes"
	"fisherman/infrastructure/logger"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var formattingParams = []interface{}{1, "s", 44.3}
var emptyParamas = []interface{}{}

func TestErrorWithLogLevel(t *testing.T) {
	t.Parallel()

	testData := []struct {
		message string
		output  string
		level   logger.LogLevel
	}{
		{message: "error", output: "error\n", level: logger.Debug},
		{message: "error", output: "error\n", level: logger.Info},
		{message: "error", output: "error\n", level: logger.Error},
		{message: "error", output: "", level: logger.None},
		{message: "", output: "\n", level: logger.Error},
		{message: "\n\n", output: "\n\n\n", level: logger.Error},
	}

	for _, tt := range testData {
		t.Run(fmt.Sprintf("Error message '%s' for level: %d", tt.message, tt.level), func(t *testing.T) {
			output := mockLogModule()
			log := logger.NewConsoleLogger(logger.OutputConfig{LogLevel: tt.level})
			log.Error(tt.message)
			assert.Equal(t, tt.output, output.String())
		})
	}
}

func TestErrorfWithLogLevel(t *testing.T) {
	output := mockLogModule()

	testData := []struct {
		message string
		output  string
		params  []interface{}
		level   logger.LogLevel
	}{
		{message: "error %d %s %f", params: formattingParams, output: "error 1 s 44.300000\n", level: logger.Error},
		{message: "error", params: emptyParamas, output: "error\n", level: logger.Error},
		{message: "error %s", params: emptyParamas, output: "error %!s(MISSING)\n", level: logger.Error},
		{message: "error %d %s %f", params: formattingParams, output: "error 1 s 44.300000\n", level: logger.Debug},
		{message: "error", params: emptyParamas, output: "error\n", level: logger.Debug},
		{message: "error %s", params: emptyParamas, output: "error %!s(MISSING)\n", level: logger.Debug},
		{message: "error %d %s %f", params: formattingParams, output: "error 1 s 44.300000\n", level: logger.Info},
		{message: "error", params: emptyParamas, output: "error\n", level: logger.Info},
		{message: "error %s", params: emptyParamas, output: "error %!s(MISSING)\n", level: logger.Info},
		{message: "error %d %s %f", params: formattingParams, output: "", level: logger.None},
		{message: "error", params: emptyParamas, output: "", level: logger.None},
		{message: "error %s", params: emptyParamas, output: "", level: logger.None},
	}

	for _, tt := range testData {
		t.Run(fmt.Sprintf("Errorf message '%s' for level: %d", tt.message, tt.level), func(t *testing.T) {
			output.Reset()
			log := logger.NewConsoleLogger(logger.OutputConfig{LogLevel: tt.level})
			log.Errorf(tt.message, tt.params...)
			assert.Equal(t, tt.output, output.String())
		})
	}
}

func TestDebugWithLogLevel(t *testing.T) {
	output := mockLogModule()

	testData := []struct {
		message string
		output  string
		level   logger.LogLevel
	}{
		{message: "debug", output: "debug\n", level: logger.Debug},
		{message: "debug", output: "", level: logger.Info},
		{message: "debug", output: "", level: logger.Error},
		{message: "debug", output: "", level: logger.None},
		{message: "", output: "\n", level: logger.Debug},
		{message: "\n\n", output: "\n\n\n", level: logger.Debug},
	}

	for _, tt := range testData {
		t.Run(fmt.Sprintf("Debug message '%s' for level: %d", tt.message, tt.level), func(t *testing.T) {
			output.Reset()
			log := logger.NewConsoleLogger(logger.OutputConfig{LogLevel: tt.level})
			log.Debug(tt.message)
			assert.Equal(t, tt.output, output.String())
		})
	}
}

func TestDebugfWithLogLevel(t *testing.T) {
	output := mockLogModule()

	testData := []struct {
		message string
		output  string
		params  []interface{}
		level   logger.LogLevel
	}{
		{message: "debug %d %s %f", params: formattingParams, output: "debug 1 s 44.300000\n", level: logger.Debug},
		{message: "debug", params: emptyParamas, output: "debug\n", level: logger.Debug},
		{message: "debug %s", params: emptyParamas, output: "debug %!s(MISSING)\n", level: logger.Debug},
		{message: "debug %d %s %f", params: formattingParams, output: "", level: logger.Error},
		{message: "debug", params: emptyParamas, output: "", level: logger.Error},
		{message: "debug %s", params: emptyParamas, output: "", level: logger.Error},
		{message: "debug %d %s %f", params: formattingParams, output: "", level: logger.Info},
		{message: "debug", params: emptyParamas, output: "", level: logger.Info},
		{message: "debug %s", params: emptyParamas, output: "", level: logger.Info},
		{message: "debug %d %s %f", params: formattingParams, output: "", level: logger.None},
		{message: "debug", params: emptyParamas, output: "", level: logger.None},
		{message: "debug %s", params: emptyParamas, output: "", level: logger.None},
	}

	for _, tt := range testData {
		t.Run(fmt.Sprintf("Debugf message '%s' for level: %d", tt.message, tt.level), func(t *testing.T) {
			output.Reset()
			log := logger.NewConsoleLogger(logger.OutputConfig{LogLevel: tt.level})
			log.Debugf(tt.message, tt.params...)
			assert.Equal(t, tt.output, output.String())
		})
	}
}

func TestInfoWithLogLevel(t *testing.T) {
	output := mockLogModule()

	testData := []struct {
		message string
		output  string
		level   logger.LogLevel
	}{
		{message: "info", output: "info\n", level: logger.Debug},
		{message: "info", output: "info\n", level: logger.Info},
		{message: "info", output: "", level: logger.Error},
		{message: "info", output: "", level: logger.None},
		{message: "", output: "\n", level: logger.Info},
		{message: "\n\n", output: "\n\n\n", level: logger.Info},
	}

	for _, tt := range testData {
		t.Run(fmt.Sprintf("Info message '%s' for level: %d", tt.message, tt.level), func(t *testing.T) {
			output.Reset()
			log := logger.NewConsoleLogger(logger.OutputConfig{LogLevel: tt.level})
			log.Info(tt.message)
			assert.Equal(t, tt.output, output.String())
		})
	}
}

func TestInfofWithLogLevel(t *testing.T) {
	output := mockLogModule()

	testData := []struct {
		message string
		output  string
		params  []interface{}
		level   logger.LogLevel
	}{
		{message: "info %d %s %f", params: formattingParams, output: "", level: logger.Error},
		{message: "info", params: emptyParamas, output: "", level: logger.Error},
		{message: "info %s", params: emptyParamas, output: "", level: logger.Error},
		{message: "info %d %s %f", params: formattingParams, output: "info 1 s 44.300000\n", level: logger.Debug},
		{message: "info", params: emptyParamas, output: "info\n", level: logger.Debug},
		{message: "info %s", params: emptyParamas, output: "info %!s(MISSING)\n", level: logger.Debug},
		{message: "info %d %s %f", params: formattingParams, output: "info 1 s 44.300000\n", level: logger.Info},
		{message: "info", params: emptyParamas, output: "info\n", level: logger.Info},
		{message: "info %s", params: emptyParamas, output: "info %!s(MISSING)\n", level: logger.Info},
		{message: "info %d %s %f", params: formattingParams, output: "", level: logger.None},
		{message: "info", params: emptyParamas, output: "", level: logger.None},
		{message: "info %s", params: emptyParamas, output: "", level: logger.None},
	}

	for _, tt := range testData {
		t.Run(fmt.Sprintf("Infof message '%s' for level: %d", tt.message, tt.level), func(t *testing.T) {
			output.Reset()
			log := logger.NewConsoleLogger(logger.OutputConfig{LogLevel: tt.level})
			log.Infof(tt.message, tt.params...)
			assert.Equal(t, tt.output, output.String())
		})
	}
}

func TestWrite(t *testing.T) {
	output := mockLogModule()

	testData := []string{
		"demo",
		"multiline demo\nmultiline demo\nmultiline demo",
		"multiline demo win\r\nmultiline demo win\r\nmultiline demo win",
		"",
		"\t\t\t",
	}

	for _, message := range testData {
		t.Run(fmt.Sprintf("Write message '%s' correctly", message), func(t *testing.T) {
			output.Reset()
			log := logger.NewConsoleLogger(logger.OutputConfig{LogLevel: logger.Debug})
			bytesCount, err := log.Write([]byte(message))
			assert.Equal(t, message, output.String())
			assert.NoError(t, err)
			assert.Equal(t, len([]byte(message)), bytesCount)
		})
	}
}

func mockLogModule() *bytes.Buffer {
	output := bytes.NewBufferString("")
	log.SetOutput(output)
	log.SetFlags(0)
	return output
}
