package main

import (
	"errors"
	"os"
	"github.com/pborman/uuid"
	"github.com/Sirupsen/logrus"
	"github.com/vjkoskela/gosteno"
)

func main() {
	// Creation Option 1: Default Underlying Logrus Logger (recommended)
	// NOTE: The logrus logger and the logger factory are "decoupled" via the global default logrus logger
	var logrusLogger *logrus.Logger = logrus.StandardLogger()
	var formatter *gosteno.Formatter = gosteno.NewFormatter()
	formatter.SetInjectContextLogger(true)
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(formatter)
	logrus.SetLevel(logrus.DebugLevel)
	var logger *gosteno.Logger = gosteno.GetLogger("examples.main")

	// Creation Option 2: New Underlying Logrus Logger
	/*
	var formatter *gosteno.Formatter = gosteno.NewFormatter()
	formatter.SetInjectContextLogger(true)
	var logrusLogger *logrus.Logger = &logrus.Logger{
		Out: os.Stdout,
		Formatter: formatter,
		Level: logrus.DebugLevel,
	}
	var logger *gosteno.Logger = gosteno.GetLoggerForLogger("test.main", logrusLogger)
	*/

	// Sample code: Existing standard Go logging
	logger.Debug("This is a vanilla debug message")
	logger.Info("This is a vanilla info message")
	logger.Warn("This is a vanilla warn message")
	logger.Error("This is a vanilla error message")

	// Sample code: Existing github.com/Sirupsen/logrus logging
	logger.WithField("foo", "bar").Info("This is a logrus info message with a single field")
	logger.WithFields(logrus.Fields{
		"foo": "bar",
		"one": 1,
		"pi": 3.14,
		"array": []string{"a","b","c"},
		"map": map[string]interface{}{"one":1,"two":2,"three":3,},
	}).Info("This is a logrus info message with multiple fields")
	logger.WithError(errors.New("This is an error")).Warn("This is a logrus warn message with an error")
	logger.WithField("foo", "bar").WithError(errors.New("This is an error")).Error("This is a logrus error message with a single field and error")

	// Sample code: Log builder
	logger.DebugBuilder().SetMessage("This is a log builder debug message").Log()
	logger.InfoBuilder().SetEvent("my_event").SetMessage("This is a log builder info message with event").Log()
	logger.WarnBuilder().
			SetEvent("my_event").
			SetMessage("This is a warn builder info message with event and error").
			SetError(errors.New("This is also another error")).
			Log()
	logger.ErrorBuilder().
			SetEvent("my_event").
			SetMessage("This is a log builder info message with event, error, data and context").
			SetError(errors.New("This is also another error")).
			AddContext("requestId", uuid.New()).
			AddData("userId", uuid.New()).
			Log()

	// Sample code: Manual event encoding
	// NOTE: Provided for completeness, this is _not_ recommended
	logger.WithFields(logrus.Fields{
		gosteno.MarkerKey: gosteno.MarkerMaps,
		"data": map[string]interface{}{
			"foo": "bar",
			"one": 1,
			"pi": 3.14,
			"array": []string{"a","b","c"},
			"map": map[string]interface{}{"one":1,"two":2,"three":3,},
		},
		"context": map[string]interface{}{
			"requestId": uuid.New(),
		},
		"error": errors.New("This is an error"),
	}).Info("This was encoded by hand")

	// Sample code: Assisted event encoding
	// NOTE: Provided for completeness, this is _not_ recommended
	gosteno.MarkerMaps.Encode(
		logrusLogger,
		"custom_event",	// event name
		"my_logger",	// logger name
		map[string]interface{}{
			"bar": "foo",
			"two": 2,
			"2pi": 6.28,
			"array": []string{"a","b","c"},
			"map": map[string]interface{}{"one":1,"two":2,"three":3,},
		},
		map[string]interface{}{
			"requestId": uuid.New(),
		},
		errors.New("This is also an error"),
	).Info("This used was encoded by the marker")
}
