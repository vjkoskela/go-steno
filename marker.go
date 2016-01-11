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
import "github.com/Sirupsen/logrus"

// Marker interface.
type Marker interface {

	// Parse event name from event.
	ParseEvent(e *logrus.Entry) string

	// Parse logger from event.
	ParseLoggerName(e *logrus.Entry) string

	// Parse data from event
	ParseData(e *logrus.Entry) map[string]interface{}

	// Parse context from event
	ParseContext(e *logrus.Entry) map[string]interface{}

	// Parse error from event
	ParseError(e *logrus.Entry) error
}
