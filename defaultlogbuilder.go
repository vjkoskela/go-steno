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

var (
	_ LogBuilder = (*DefaultLogBuilder)(nil)
)

// DefaultLogBuilder is the default LogBuilder implementation that satisfies the LogBuilder contract.
type DefaultLogBuilder struct {
	logger *logrus.Logger
	level logrus.Level
	event string
	loggerName string
	message string
	err error
	data map[string]interface{}
	context map[string]interface{}
}

func NewDefaultLogBuilder(l *logrus.Logger, v logrus.Level, n string) *DefaultLogBuilder {
	return &DefaultLogBuilder{
			logger: l,
			level: v,
			loggerName: n,
			data: make(map[string]interface{}),
			context: make(map[string]interface{})}
}

func (dlb *DefaultLogBuilder) SetEvent(event string) LogBuilder {
	dlb.event = event
	return dlb
}

func (dlb *DefaultLogBuilder) SetMessage(message string) LogBuilder {
	dlb.message = message
	return dlb
}

func (dlb *DefaultLogBuilder) SetError(err error) LogBuilder {
	dlb.err = err
	return dlb
}

func (dlb *DefaultLogBuilder) AddData(key string, value interface{}) LogBuilder {
	dlb.data[key] = value
	return dlb
}

func (dlb *DefaultLogBuilder) AddContext(key string, value interface{}) LogBuilder {
	dlb.context[key] = value
	return dlb
}

func (dlb *DefaultLogBuilder) Log() {
	var entry *logrus.Entry = MarkerMaps.Encode(
		dlb.logger,
		dlb.event,
		dlb.loggerName,
		dlb.data,
		dlb.context,
		dlb.err,
	)
	output(entry, dlb.message, dlb.logger, dlb.level)
}

func output(e *logrus.Entry, m string, l *logrus.Logger, v logrus.Level) {
	switch v {
	default:
		e.Info(m)
	case logrus.DebugLevel:
		e.Debug(m)
	case logrus.InfoLevel:
		e.Info(m)
	case logrus.WarnLevel:
		e.Warn(m)
	case logrus.ErrorLevel:
		e.Error(m)
	case logrus.FatalLevel:
		e.Fatal(m)
	case logrus.PanicLevel:
		e.Panic(m)
	}
}