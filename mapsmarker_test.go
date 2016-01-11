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
	"errors"
	"reflect"
	"testing"
	"github.com/Sirupsen/logrus"
)

var (
	logger *logrus.Logger = logrus.New()
	mm *MapsMarker = new(MapsMarker)
	emptyEntry *logrus.Entry = logrus.NewEntry(logger)
)

func TestMapsMarkerEncode(t *testing.T) {
	t.Parallel()
	var expectedEvent string = "my_event"
	var expectedLoggerName string = "my_logger"
	var expectedData map[string]interface{} = map[string]interface{}{"foo":"bar","one":1,"pi":3.14,}
	var expectedContext map[string]interface{} = map[string]interface{}{"bar":"foo","two":2,"2pi":6.28,}
	var expectedError error = errors.New("this is an error")
	var e *logrus.Entry = mm.Encode(
		logger,
		expectedEvent,
		expectedLoggerName,
		expectedData,
		expectedContext,
		expectedError)
	if e.Data["event"] != expectedEvent {
		t.Errorf("Encode failed to encode event")
	}
	if e.Data["logger"] != expectedLoggerName {
		t.Errorf("Encode failed to encode logger")
	}
	if !reflect.DeepEqual(e.Data["data"], expectedData) {
		t.Errorf("Encode failed to encode data")
	}
	if !reflect.DeepEqual(e.Data["context"], expectedContext) {
		t.Errorf("Encode failed to encode context")
	}
	if e.Data["error"] != expectedError {
		t.Errorf("Encode failed to encode error")
	}
}

func TestMapsMarkerParseName(t *testing.T) {
	t.Parallel()
	var expectedName string = "my_event"
	var e *logrus.Entry = logrus.WithField("event", expectedName)
	var actualName string
	if actualName = mm.ParseEvent(e); actualName != expectedName {
		t.Errorf("ParseEvent failed; expected '%s' instead actual '%s'", expectedName, actualName)
	}
	if actualName = mm.ParseEvent(emptyEntry); actualName != "" {
		t.Errorf("ParseEvent failed; expected nil instead actual '%s'", actualName)
	}
}

func TestMapsMarkerParseLoggerName(t *testing.T) {
	t.Parallel()
	var expectedLoggerName string = "my_logger"
	var e *logrus.Entry = logrus.WithField("logger", expectedLoggerName)
	var actualLoggerName string
	if actualLoggerName = mm.ParseLoggerName(e); actualLoggerName != expectedLoggerName {
		t.Errorf("ParseLoggerName failed; expected '%s' instead actual '%s'", expectedLoggerName, actualLoggerName)
	}
	if actualLoggerName = mm.ParseLoggerName(emptyEntry); actualLoggerName != "" {
		t.Errorf("ParseLoggerName failed; expected nil instead actual '%s'", actualLoggerName)
	}
}

func TestMapsMarkerParseData(t *testing.T) {
	t.Parallel()
	var expectedData map[string]interface{} = map[string]interface{}{"foo":"bar","one":1,"pi":3.14,}
	var e *logrus.Entry = logrus.WithField("data", expectedData)
	var actualData map[string]interface{}
	if actualData = mm.ParseData(e); !reflect.DeepEqual(actualData, expectedData) {
		t.Errorf("ParseData failed; expected '%v' instead actual '%v'", expectedData, actualData)
	}
	if actualData = mm.ParseData(emptyEntry); actualData != nil {
		t.Errorf("ParseData failed; expected nil instead actual '%v'", actualData)
	}
}

func TestMapsMarkerParseContext(t *testing.T) {
	t.Parallel()
	var expectedContext map[string]interface{} = map[string]interface{}{"foo":"bar","one":1,"pi":3.14,}
	var e *logrus.Entry = logrus.WithField("context", expectedContext)
	var actualContext map[string]interface{}
	if actualContext = mm.ParseContext(e); !reflect.DeepEqual(actualContext, expectedContext) {
		t.Errorf("ParseContext failed; expected '%v' instead actual '%v'", expectedContext, actualContext)
	}
	if actualContext = mm.ParseContext(emptyEntry); actualContext != nil {
		t.Errorf("ParseContext failed; expected nil instead actual '%v'", actualContext)
	}
}

func TestMapsMarkerParseError(t *testing.T) {
	t.Parallel()
	var expectedError error = errors.New("this is an error")
	var e *logrus.Entry = logrus.WithField("error", expectedError)
	var actualError error
	if actualError = mm.ParseError(e); actualError != expectedError {
		t.Errorf("ParseLoggerError failed; expected '%s' instead actual '%s'", expectedError, actualError)
	}
	if actualError = mm.ParseError(emptyEntry); actualError != nil {
		t.Errorf("ParseLoggerError failed; expected nil instead actual '%s'", expectedError, actualError)
	}
}
