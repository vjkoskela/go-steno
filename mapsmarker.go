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
	"github.com/Sirupsen/logrus"
)

const (
	EVENT_DATA_EVENT_KEY string = "event"
	EVENT_DATA_LOGGER_KEY string = "logger"
	EVENT_DATA_DATA_KEY string = "data"
	EVENT_DATA_CONTEXT_KEY string = "context"
	EVENT_DATA_ERROR_KEY string = "error"
)

var (
	_ Marker = (*MapsMarker)(nil)
)

// Maps marker implementation.
type MapsMarker struct { }

// Encode.
func (smm *MapsMarker) Encode(
		logger *logrus.Logger,
		event string,
		loggerName string,
		data map[string]interface{},
		context map[string]interface{},
		err error) *logrus.Entry {

	var entry *logrus.Entry = logrus.NewEntry(logger).WithFields(logrus.Fields{
		MarkerKey: smm,
		EVENT_DATA_EVENT_KEY: event,
		EVENT_DATA_LOGGER_KEY: loggerName,
		EVENT_DATA_DATA_KEY: data,
		EVENT_DATA_CONTEXT_KEY: context,
		EVENT_DATA_ERROR_KEY: err,})
	return entry
}

// Parse event name from event.
func (smm *MapsMarker) ParseEvent(e *logrus.Entry) string {
	var v interface{}
	v = e.Data[EVENT_DATA_EVENT_KEY]
	switch v := v.(type) {
	default:
		return ""
	case string:
		return v
	}
}

// Parse logger name from event.
func (smm *MapsMarker) ParseLoggerName(e *logrus.Entry) string {
	var v interface{}
	v = e.Data[EVENT_DATA_LOGGER_KEY]
	switch v := v.(type) {
	default:
		return ""
	case string:
		return v
	}
}

// Parse data from event.
func (smm *MapsMarker) ParseData(e *logrus.Entry) map[string]interface{} {
	var v interface{}
	v = e.Data[EVENT_DATA_DATA_KEY]
	switch v := v.(type) {
	default:
		return nil
	case map[string]interface{}:
		return v
	}
}

// Parse context from event.
func (smm *MapsMarker) ParseContext(e *logrus.Entry) map[string]interface{} {
	var v interface{}
	v = e.Data[EVENT_DATA_CONTEXT_KEY]
	switch v := v.(type) {
	default:
		return nil
	case map[string]interface{}:
		return v
	}
}

// Parse data from event.
func (smm *MapsMarker) ParseError(e *logrus.Entry) error {
	var v interface{}
	v = e.Data[EVENT_DATA_ERROR_KEY]
	switch v := v.(type) {
	default:
		return nil
	case error:
		return v
	}
}
