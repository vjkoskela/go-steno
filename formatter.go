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
	"os"
	"strconv"
	"time"
	"github.com/pborman/uuid"
	"github.com/Sirupsen/logrus"
)

const (
	globalDefaultEventName = "log"
)

var (
	newLine = []byte("\n")
	hostname = "<UNKNOWN>"
	processId = "<UNKNOWN>"
)

func init() {
	var value string
	var err error

	// Hostname
	value, err = os.Hostname()
	if (err == nil) {
		hostname = value
	}
	// Process id
	processId = strconv.Itoa(os.Getpid())
}

type Formatter struct {
	logEventName string
	injectContextHost bool
	injectContextProcess bool
	injectContextLogger bool
}

func NewFormatter() *Formatter {
	return &Formatter{
		logEventName: globalDefaultEventName,
		injectContextHost: true,
		injectContextProcess: true,
		injectContextLogger: false,
	}
}

func (sf *Formatter) Format(e *logrus.Entry) (result []byte, err error) {
	var initialStenoBuffer []byte = make([]byte, 0, 512)
	var initialUserBuffer []byte = make([]byte, 0, 256)
	var stenoBuffer *bytes.Buffer = bytes.NewBuffer(initialStenoBuffer)
	var userBuffer *bytes.Buffer = bytes.NewBuffer(initialUserBuffer)
	var value []byte

	// Begin steno wrapper
	if _, err = stenoBuffer.WriteString("{"); err != nil {
		return
	}
	if err = writeKeyStringValue(stenoBuffer, "time", sf.getTime(e)); err != nil {
		return
	}
	if err = writeKeyStringValue(stenoBuffer, "name", sf.getEventName(e, sf.logEventName)); err != nil {
		return
	}
	if err = writeKeyStringValue(stenoBuffer, "level", sf.getLevel(e)); err != nil {
		return
	}

	// Encode user data; handle errors more gracefully by encoding them into the buffer if possible
	if value, err = sf.getData(e); err != nil {
		if err = writeError(stenoBuffer, err); err != nil {
			return
		}
	} else if err = writeKeyJsonValue(userBuffer, "data", value); err != nil {
		if err = writeError(stenoBuffer, err); err != nil {
			return
		}
	} else {
		if value, err = sf.getContext(e); err != nil {
			if err = writeError(stenoBuffer, err); err != nil {
				return
			}
		} else if err = writeKeyJsonValue(userBuffer, "context", value); err != nil {
			if err = writeError(stenoBuffer, err); err != nil {
				return
			}
		} else {
			if value, err = sf.getError(e); err != nil {
				if err = writeError(userBuffer, err); err != nil {
					return
				}
			} else if value != nil {
				if err = writeKeyJsonValue(userBuffer, "exception", value); err != nil {
					if err = writeError(stenoBuffer, err); err != nil {
						return
					}
				} else if _, err = stenoBuffer.Write(userBuffer.Bytes()); err != nil {
					return
				}
			} else if _, err = stenoBuffer.Write(userBuffer.Bytes()); err != nil {
				return
			}
		}
	}

	// Complete steno wrapper
	if err = writeKeyStringValue(stenoBuffer, "id", uuid.New()); err != nil {
		return
	}
	if err = writeKeyStringValue(stenoBuffer, "version", "0"); err != nil {
		return
	}
	stenoBuffer.Write(newLine);
	result = stenoBuffer.Bytes()
	result[len(result) - len(newLine) - 1] = '}'
	return
}

func (sf *Formatter) LogEventName() string {
	return sf.logEventName
}

func (sf *Formatter) SetLogEventName(v string) {
	sf.logEventName = v
}

func (sf *Formatter) InjectContextHost() bool {
	return sf.injectContextHost
}

func (sf *Formatter) SetInjectContextHost(v bool) {
	sf.injectContextHost = v
}

func (sf *Formatter) InjectContextProcess() bool {
	return sf.injectContextProcess
}

func (sf *Formatter) SetInjectContextProcess(v bool) {
	sf.injectContextProcess = v
}

func (sf *Formatter) InjectContextLogger() bool {
	return sf.injectContextLogger
}

func (sf *Formatter) SetInjectContextLogger(v bool) {
	sf.injectContextLogger = v
}

func (sf *Formatter) getTime(e *logrus.Entry) string {
	return e.Time.UTC().Format(time.RFC3339Nano)
}

func (sf *Formatter) getEventName(e *logrus.Entry, defaultName string) string {
	var marker interface{} = e.Data[MarkerKey]
	switch marker := marker.(type) {
	default:
		return defaultName
	case *MapsMarker:
		var name string
		if name = marker.ParseEvent(e); name != "" {
			return name
		}
		return defaultName
	}
}

func (sf *Formatter) getLevel(e *logrus.Entry) string {
	switch e.Level {
	case logrus.DebugLevel:
		return "debug"
	case logrus.InfoLevel:
		return "info"
	case logrus.WarnLevel:
		return "warn"
	case logrus.ErrorLevel:
		return "crit"
	case logrus.FatalLevel:
		return "fatal"
	case logrus.PanicLevel:
		return "fatal"
	}
	return "unknown"
}

func (sf *Formatter) getData(e *logrus.Entry) (jsonBytes []byte, err error) {
	var data map[string]interface{}
	var marker interface{} = e.Data[MarkerKey]
	var hasValidMarker bool = true
	switch marker := marker.(type) {
	default:
		hasValidMarker = false
		data = e.Data
	case *MapsMarker:
		data = marker.ParseData(e)
	}
	var buffer bytes.Buffer
	if _, err = buffer.WriteString("{"); err != nil {
		return
	}
	if e.Message != "" {
		if err = writeKeyStringValue(&buffer, "message", e.Message); err != nil {
			return
		}
	}
	for key, value := range data {
		// Favor explicit message in event (if not empty) over any data with the same key
		if key == "message" && e.Message != "" {
			continue
		}
		// Suppress error in event if processing raw event data (e.g. without valid marker)
		if key == logrus.ErrorKey && !hasValidMarker {
			continue
		}
		var valueJsonBytes []byte
		if valueJsonBytes, err = json.Marshal(value); err != nil {
			return
		}
		if err = writeKeyJsonValue(&buffer, key, valueJsonBytes); err != nil {
			return
		}
	}
	if buffer.Len() == 1 {
		if _, err = buffer.WriteString("}"); err != nil {
			return
		}
		jsonBytes = buffer.Bytes()
	} else {
		jsonBytes = buffer.Bytes()
		jsonBytes[len(jsonBytes) - 1] = '}'
	}
	return
}

func (sf *Formatter) getContext(e *logrus.Entry) (jsonBytes []byte, err error) {
	var context map[string]interface{}
	var loggerName string
	var marker interface{} = e.Data[MarkerKey]
	switch marker := marker.(type) {
	default:
		context = nil
		loggerName = ""
	case *MapsMarker:
		context = marker.ParseContext(e)
		loggerName = marker.ParseLoggerName(e)
	}
	var buffer bytes.Buffer
	if _, err = buffer.WriteString("{"); err != nil {
		return
	}
	for key, value := range context {
		var valueJsonBytes []byte
		if valueJsonBytes, err = json.Marshal(value); err != nil {
			return
		}
		if err = writeKeyJsonValue(&buffer, key, valueJsonBytes); err != nil {
			return
		}
	}
	if sf.injectContextHost {
		if err = writeKeyStringValue(&buffer, "host", hostname); err != nil {
			return
		}
	}
	if sf.injectContextProcess {
		if err = writeKeyStringValue(&buffer, "processId", processId); err != nil {
			return
		}
	}
	if sf.injectContextLogger && loggerName != "" {
		if err = writeKeyStringValue(&buffer, "logger", loggerName); err != nil {
			return
		}
	}
	if buffer.Len() == 1 {
		if _, err = buffer.WriteString("}"); err != nil {
			return
		}
		jsonBytes = buffer.Bytes()
	} else {
		jsonBytes = buffer.Bytes()
		jsonBytes[len(jsonBytes) - 1] = '}'
	}
	return
}

func (sf *Formatter) getError(e *logrus.Entry) (jsonBytes []byte, err error) {
	var entryError error
	var marker interface{} = e.Data[MarkerKey]
	switch marker := marker.(type) {
	default:
		entryError = nil
		if logrusError, ok := e.Data[logrus.ErrorKey].(error); ok {
			entryError = logrusError
		}
	case *MapsMarker:
		entryError = marker.ParseError(e)
	}
	if entryError != nil {
		var buffer bytes.Buffer
		if _, err = buffer.WriteString("{"); err != nil {
			return
		}
		if err = writeKeyStringValue(&buffer, "type", "error"); err != nil {
			return
		}
		if err = writeKeyStringValue(&buffer, "message", entryError.Error()); err != nil {
			return
		}
		if err = writeKeyJsonValue(&buffer, "backtrace", []byte("[]")); err != nil {
			return
		}

		jsonBytes = buffer.Bytes()
		jsonBytes[len(jsonBytes) - 1] = '}'
	}
	return
}

func writeKeyStringValue(buffer *bytes.Buffer, key string, value string) (err error) {
	var bytes []byte
	if bytes, err = json.Marshal(key); err != nil {
		return
	} else if _, err = buffer.Write(bytes); err != nil {
		return
	}
	if _, err = buffer.WriteString(":"); err != nil {
		return
	}
	if bytes, err = json.Marshal(value); err != nil {
		return
	} else if _, err = buffer.Write(bytes); err != nil {
		return
	}
	if _, err = buffer.WriteString(","); err != nil {
		return
	}
	return nil
}

func writeKeyJsonValue(buffer *bytes.Buffer, key string, jsonBytes []byte) (err error) {
	var bytes []byte
	if bytes, err = json.Marshal(key); err != nil {
		return
	} else if _, err = buffer.Write(bytes); err != nil {
		return
	}
	if _, err = buffer.WriteString(":"); err != nil {
		return
	}
	if _, err = buffer.Write(jsonBytes); err != nil {
		return
	}
	if _, err = buffer.WriteString(","); err != nil {
		return
	}
	return nil
}

func writeError(buffer *bytes.Buffer, internalError error) (err error) {
	var internalErrorAsMap = map[string]interface{} {
		"message": internalError.Error(),
	}
	var bytes []byte
	if bytes, err = json.Marshal(internalErrorAsMap); err != nil {
		return
	}
	if _, err = buffer.Write(bytes); err != nil {
		return
	}
	if _, err = buffer.WriteString(","); err != nil {
		return
	}
	return nil
}
