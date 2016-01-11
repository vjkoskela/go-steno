/*
Copyright 2016 Ville Koskela

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package gosteno

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"
	"github.com/pborman/uuid"
	"github.com/Sirupsen/logrus"
)

func TestLoggerFactoryGetLogger(t *testing.T) {
	// WARNING: This test mucks with the default global logger. So it must be run serially.
	var originalWriter io.Writer = logrus.StandardLogger().Out
	var originalFormatter logrus.Formatter = logrus.StandardLogger().Formatter

	// Create a formatter that includes the logger name
	var formatter *Formatter = NewFormatter()
	formatter.SetInjectContextLogger(true)

	// Create a logrus logger to a buffer
	var buffer *bytes.Buffer = new(bytes.Buffer)
	logrus.SetOutput(buffer)
	logrus.SetFormatter(formatter)

	// Obtain a named logger for the default global logrus logger
	var expectedLoggerName string = "my_logger_name-" + uuid.New()
	var expectedMessage string = "Hello World-" + uuid.New()
	var logger *Logger = GetLogger(expectedLoggerName)
	logger.InfoBuilder().SetMessage(expectedMessage).Log()

	// Unmarshal the buffer and interrogate the results for the logger name and message
	unmarshallAndIterrogate(t, buffer, expectedLoggerName, expectedMessage)

	// Restore the original writer on the default global logger
	logrus.SetOutput(originalWriter)
	logrus.SetFormatter(originalFormatter)
}

func TestLoggerFactoryGetLoggerForLogger(t *testing.T) {
	t.Parallel()
	// Create a formatter that includes the logger name
	var formatter *Formatter = NewFormatter()
	formatter.SetInjectContextLogger(true)

	// Create a logrus logger to a buffer
	var buffer *bytes.Buffer = new(bytes.Buffer)
	var logrusLogger *logrus.Logger = &logrus.Logger{
		Out: buffer,
		Formatter: formatter,
		Level: logrus.InfoLevel,
	}

	// Obtain a named logger for the custom logrus logger
	var expectedLoggerName string = "my_logger_name-" + uuid.New()
	var expectedMessage string = "Hello World-" + uuid.New()
	var logger *Logger = GetLoggerForLogger(expectedLoggerName, logrusLogger)
	logger.InfoBuilder().SetMessage(expectedMessage).Log()

	// Unmarshal the buffer and interrogate the results for the logger name and message
	unmarshallAndIterrogate(t, buffer, expectedLoggerName, expectedMessage)
}

func unmarshallAndIterrogate(t *testing.T, b *bytes.Buffer, l string, m string) {
	var err error
	var rootNode map[string]interface{}
	if err = json.Unmarshal(b.Bytes(), &rootNode); err != nil {
		t.Errorf("Unmarshal failed because %v in buffer %v", err, b)
		return
	}

	if node, ok := rootNode["data"].(map[string]interface{}); ok {
		if value, ok := node["message"].(string); ok {
			if value != m {
				t.Errorf("Expected message not found; message=%v", node)
			}
		} else {
			t.Errorf("Message node not found; buffer=%v", b)
		}
	} else {
		t.Errorf("Data node not found; buffer=%v", b)
	}

	if node, ok := rootNode["context"].(map[string]interface{}); ok {
		if value, ok := node["logger"].(string); ok {
			if value != l {
				t.Errorf("Expected logger name not found; message=%v", node)
			}
		} else {
			t.Errorf("Logger node not found; buffer=%v", b)
		}
	} else {
		t.Errorf("Context node not found; buffer=%v", b)
	}
}