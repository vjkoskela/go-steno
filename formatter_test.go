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
	"testing"
	"github.com/Sirupsen/logrus"
	"errors"
)

type subWidget struct {
	Name string
}

type widget struct {
	Name string
	Parts *[]subWidget
}

// TODO:
// 1) Test catastropic failure resulting in no log message.
// 2) Test data/context serialization failure resulting in skeleton message.

const (
	formatterTestDataPath = "./testdata/formatter_test/"
)

func TestFormatterDefaults(t *testing.T) {
	var formatter *Formatter = NewFormatter()
	if v := formatter.InjectContextHost(); v != true {
		t.Errorf("Incorrect default value for injectContextHost %v", v)
	}
	if v := formatter.InjectContextLogger(); v != false {
		t.Errorf("Incorrect default value for injectContextLogger %v", v)
	}
	if v := formatter.InjectContextProcess(); v != true {
		t.Errorf("Incorrect default value for injectContextProcess %v", v)
	}
	if v := formatter.LogEventName(); v != "log" {
		t.Errorf("Incorrect default value for log event name %v", v)
	}
}

func TestFormatterLevelMapping(t *testing.T) {
	var formatter *Formatter = NewFormatter()
	if v := formatter.getLevel(&logrus.Entry{Level: logrus.DebugLevel}); v != "debug" {
		t.Errorf("Incorrect level for DebugLevel %v", v)
	}
	if v := formatter.getLevel(&logrus.Entry{Level: logrus.InfoLevel}); v != "info" {
		t.Errorf("Incorrect level for InfoLevel %v", v)
	}
	if v := formatter.getLevel(&logrus.Entry{Level: logrus.WarnLevel}); v != "warn" {
		t.Errorf("Incorrect level for WarnLevel %v", v)
	}
	if v := formatter.getLevel(&logrus.Entry{Level: logrus.ErrorLevel}); v != "crit" {
		t.Errorf("Incorrect level for ErrorLevel %v", v)
	}
	if v := formatter.getLevel(&logrus.Entry{Level: logrus.FatalLevel}); v != "fatal" {
		t.Errorf("Incorrect level for FatalLevel %v", v)
	}
	if v := formatter.getLevel(&logrus.Entry{Level: logrus.PanicLevel}); v != "fatal" {
		t.Errorf("Incorrect level for PanicLevel %v", v)
	}
}

func TestFormatterGlobalDefaultEventName(t *testing.T) {
	t.Parallel()
	var formatter *Formatter = NewFormatter()
	logger, buffer := HelperTestGetLogger("TestFormatterGlobalDefaultEventName", logrus.DebugLevel, formatter)
	logger.DebugBuilder().SetMessage("TestFormatterGlobalDefaultEventName").Log()
	HelperTestVerify(t, buffer, formatterTestDataPath + "TestFormatterGlobalDefaultEventName.expected.json")
}

func TestFormatterConfiguredDefaultEventName(t *testing.T) {
	t.Parallel()
	var formatter *Formatter = NewFormatter()
	formatter.SetLogEventName("default_event")
	logger, buffer := HelperTestGetLogger("TestFormatterConfiguredDefaultEventName", logrus.DebugLevel, formatter)
	logger.DebugBuilder().SetMessage("TestFormatterConfiguredDefaultEventName").Log()
	HelperTestVerify(t, buffer, formatterTestDataPath + "TestFormatterConfiguredDefaultEventName.expected.json")
}

func TestFormatterSpecifiedEventName(t *testing.T) {
	t.Parallel()
	var formatter *Formatter = NewFormatter()
	formatter.SetLogEventName("default_event")
	logger, buffer := HelperTestGetLogger("TestFormatterSpecifiedEventName", logrus.DebugLevel, formatter)
	logger.DebugBuilder().SetEvent("custom_event").SetMessage("TestFormatterSpecifiedEventName").Log()
	HelperTestVerify(t, buffer, formatterTestDataPath + "TestFormatterSpecifiedEventName.expected.json")
}

func TestFormatterEnableLoggerName(t *testing.T) {
	t.Parallel()
	var formatter *Formatter = NewFormatter()
	formatter.SetInjectContextLogger(true)
	logger, buffer := HelperTestGetLogger("TestFormatterEnableLoggerName", logrus.DebugLevel, formatter)
	logger.DebugBuilder().SetMessage("TestFormatterEnableLoggerName").Log()
	HelperTestVerifyIgnoreContext(t, buffer, formatterTestDataPath + "TestFormatterEnableLoggerName.expected.json", []string{"logger"})
}

func TestFormatterDisableProcess(t *testing.T) {
	t.Parallel()
	var formatter *Formatter = NewFormatter()
	formatter.SetInjectContextProcess(false)
	logger, buffer := HelperTestGetLogger("TestFormatterDisableProcess", logrus.DebugLevel, formatter)
	logger.DebugBuilder().SetMessage("TestFormatterDisableProcess").Log()
	HelperTestVerify(t, buffer, formatterTestDataPath + "TestFormatterDisableProcess.expected.json")
}

func TestFormatterDisableHost(t *testing.T) {
	t.Parallel()
	var formatter *Formatter = NewFormatter()
	formatter.SetInjectContextHost(false)
	logger, buffer := HelperTestGetLogger("TestFormatterDisableHost", logrus.DebugLevel, formatter)
	logger.DebugBuilder().SetMessage("TestFormatterDisableHost").Log()
	HelperTestVerify(t, buffer, formatterTestDataPath + "TestFormatterDisableHost.expected.json")
}

func TestFormatterEmptyContext(t *testing.T) {
	t.Parallel()
	var formatter *Formatter = NewFormatter()
	formatter.SetInjectContextHost(false)
	formatter.SetInjectContextProcess(false)
	logger, buffer := HelperTestGetLogger("TestFormatterEmptyContext", logrus.DebugLevel, formatter)
	logger.DebugBuilder().SetMessage("TestFormatterEmptyContext").Log()
	HelperTestVerify(t, buffer, formatterTestDataPath + "TestFormatterEmptyContext.expected.json")
}

func TestFormatterComplexContext(t *testing.T) {
	t.Parallel()
	var formatter *Formatter = NewFormatter()
	formatter.SetInjectContextHost(false)
	formatter.SetInjectContextProcess(false)
	logger, buffer := HelperTestGetLogger("TestFormatterComplexContext", logrus.DebugLevel, formatter)
	logger.DebugBuilder().
		AddContext("foo", "bar").
		AddContext("one", 1).
		AddContext("pi", 3.14).
		AddContext("map", map[string]interface{}{"a":"A","b":"B",}).
		AddContext("list", []int{1,2}).
		AddContext("obj", createWidget("TestFormatterComplexContext")).
		SetMessage("TestFormatterComplexContext").
		Log()
	HelperTestVerifyIgnoreContext(
		t,
		buffer,
		formatterTestDataPath + "TestFormatterComplexContext.expected.json",
		[]string{"foo", "one", "pi", "map", "list", "obj"})
}

func TestFormatterEmptyData(t *testing.T) {
	t.Parallel()
	var formatter *Formatter = NewFormatter()
	logger, buffer := HelperTestGetLogger("TestFormatterEmptyData", logrus.DebugLevel, formatter)
	logger.DebugBuilder().Log()
	HelperTestVerify(t, buffer, formatterTestDataPath + "TestFormatterEmptyData.expected.json")
}

func TestFormatterComplexData(t *testing.T) {
	t.Parallel()
	var formatter *Formatter = NewFormatter()
	formatter.SetInjectContextHost(false)
	formatter.SetInjectContextProcess(false)
	logger, buffer := HelperTestGetLogger("TestFormatterComplexData", logrus.DebugLevel, formatter)
	logger.DebugBuilder().
		AddData("foo", "bar").
		AddData("one", 1).
		AddData("pi", 3.14).
		AddData("map", map[string]interface{}{"a":"A","b":"B",}).
		AddData("list", []int{1,2}).
		AddData("obj", createWidget("TestFormatterComplexData")).
	SetMessage("TestFormatterComplexData").
	Log()
	HelperTestVerify(t, buffer, formatterTestDataPath + "TestFormatterComplexData.expected.json")
}

func TestFormatterWithError(t *testing.T) {
	t.Parallel()
	var formatter *Formatter = NewFormatter()
	logger, buffer := HelperTestGetLogger("TestFormatterWithError", logrus.DebugLevel, formatter)
	logger.DebugBuilder().SetError(errors.New("This is an error")).SetMessage("TestFormatterWithError").Log()
	HelperTestVerify(t, buffer, formatterTestDataPath + "TestFormatterWithError.expected.json")
}

func TestFormatterWithLogrusError(t *testing.T) {
	t.Parallel()
	var formatter *Formatter = NewFormatter()
	logger, buffer := HelperTestGetLogger("TestFormatterWithLogrusError", logrus.DebugLevel, formatter)
	logger.WithError(errors.New("This is an error")).Debug("TestFormatterWithLogrusError")
	HelperTestVerify(t, buffer, formatterTestDataPath + "TestFormatterWithLogrusError.expected.json")
}

func createWidget(name string) (widget) {
	var parts []subWidget = make([]subWidget, 2, 2)
	parts[0] = *new(subWidget)
	parts[0].Name = name + "-1"
	parts[1] = *new(subWidget)
	parts[1].Name = name + "-2"
	var w *widget = new(widget)
	w.Name = name
	w.Parts = &parts
	return *w
}