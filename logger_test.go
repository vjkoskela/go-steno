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
	"github.com/xeipuuv/gojsonschema"
)

// TODO:
// 1) Test Fatal and Panic methods using a mock logrus logger.
// 2) Test no-op stub methods from standard logger.
// 3) Test output method to info w/o stack trace method from standard logger.

const (
	testDataPath = "./testdata/logger_test/"
)

var (
	formatter *Formatter = NewFormatter()
	stenoSchemaLoader = gojsonschema.NewReferenceLoader("file://./testdata/steno.schema.json")
)

func TestLoggerDebugBuilder(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerDebugBuilder", logrus.DebugLevel)
	logger.DebugBuilder().SetMessage("TestLoggerDebugBuilder").Log()
	HelperTestVerify(t, buffer, testDataPath + "TestLoggerDebugBuilder.expected.json")
}

func TestLoggerDebugBuilderLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerDebugBuilderLevelSuppressed", logrus.InfoLevel)
	logger.DebugBuilder().SetMessage("TestLoggerDebugBuilderLevelSuppressed").Log()
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerInfoBuilder(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerInfoBuilder", logrus.InfoLevel)
	logger.InfoBuilder().SetMessage("TestLoggerInfoBuilder").Log()
	HelperTestVerify(t, buffer, testDataPath + "TestLoggerInfoBuilder.expected.json")
}

func TestLoggerInfoBuilderLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerInfoBuilderLevelSuppressed", logrus.WarnLevel)
	logger.InfoBuilder().SetMessage("TestLoggerInfoBuilderLevelSuppressed").Log()
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerWarnBuilder(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarnBuilder", logrus.WarnLevel)
	logger.WarnBuilder().SetMessage("TestLoggerWarnBuilder").Log()
	HelperTestVerify(t, buffer, testDataPath + "TestLoggerWarnBuilder.expected.json")
}

func TestLoggerWarnBuilderLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarnBuilderLevelSuppressed", logrus.ErrorLevel)
	logger.WarnBuilder().SetMessage("TestLoggerWarnBuilderLevelSuppressed").Log()
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerWarningBuilder(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarningBuilder", logrus.WarnLevel)
	logger.WarningBuilder().SetMessage("TestLoggerWarningBuilder").Log()
	HelperTestVerify(t, buffer, testDataPath + "TestLoggerWarningBuilder.expected.json")
}

func TestLoggerWarningBuilderLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarningBuilderLevelSuppressed", logrus.ErrorLevel)
	logger.WarningBuilder().SetMessage("TestLoggerWarningBuilderLevelSuppressed").Log()
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerErrorBuilder(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerErrorBuilder", logrus.ErrorLevel)
	logger.ErrorBuilder().SetMessage("TestLoggerErrorBuilder").Log()
	HelperTestVerify(t, buffer, testDataPath + "TestLoggerErrorBuilder.expected.json")
}

func TestLoggerErrorBuilderLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerErrorBuilderLevelSuppressed", logrus.FatalLevel)
	logger.ErrorBuilder().SetMessage("TestLoggerErrorBuilderLevelSuppressed").Log()
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerPrint(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerPrint", logrus.InfoLevel)
	logger.Print("TestLoggerPrint")
	HelperTestVerify(t, buffer, testDataPath + "TestLoggerPrint.expected.json")
}

func TestLoggerPrintLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerPrintLevelSuppressed", logrus.WarnLevel)
	logger.Print("TestLoggerPrintLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerPrintf(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerPrintf", logrus.InfoLevel)
	logger.Printf("%sLogger%s", "Test", "Printf")
	HelperTestVerify(t, buffer, testDataPath + "TestLoggerPrintf.expected.json")
}

func TestLoggerPrintfLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerPrintfLevelSuppressed", logrus.WarnLevel)
	logger.Printf("TestLoggerPrintfLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerPrintln(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerPrintln", logrus.InfoLevel)
	logger.Println("TestLoggerPrintln")
	HelperTestVerify(t, buffer, testDataPath + "TestLoggerPrintln.expected.json")
}

func TestLoggerPrintlnLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerPrintlnLevelSuppressed", logrus.WarnLevel)
	logger.Print("TestLoggerPrintlnLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerDebug(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerDebug", logrus.DebugLevel)
	logger.Debug("TestLoggerDebug")
	HelperTestVerify(t, buffer, testDataPath + "TestLoggerDebug.expected.json")
}

func TestLoggerDebugLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerDebugLevelSuppressed", logrus.InfoLevel)
	logger.Debug("TestLoggerDebugLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerDebugf(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerDebugf", logrus.DebugLevel)
	logger.Debugf("%sLogger%s", "Test", "Debugf")
	HelperTestVerify(t, buffer, testDataPath + "TestLoggerDebugf.expected.json")
}

func TestLoggerDebugfLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerDebugfLevelSuppressed", logrus.InfoLevel)
	logger.Debugf("TestLoggerDebugfLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerDebugln(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerDebugln", logrus.DebugLevel)
	logger.Debugln("TestLoggerDebugln")
	HelperTestVerify(t, buffer, testDataPath + "TestLoggerDebugln.expected.json")
}

func TestLoggerDebuglnLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerDebuglnLevelSuppressed", logrus.InfoLevel)
	logger.Debug("TestLoggerDebuglnLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerInfo(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerInfo", logrus.InfoLevel)
	logger.Info("TestLoggerInfo")
	HelperTestVerify(t, buffer, testDataPath + "TestLoggerInfo.expected.json")
}

func TestLoggerInfoLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerInfoLevelSuppressed", logrus.WarnLevel)
	logger.Info("TestLoggerInfoLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerInfof(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerInfof", logrus.InfoLevel)
	logger.Infof("%sLogger%s", "Test", "Infof")
	HelperTestVerify(t, buffer, testDataPath + "TestLoggerInfof.expected.json")
}

func TestLoggerInfofLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerInfofLevelSuppressed", logrus.WarnLevel)
	logger.Infof("TestLoggerInfofLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerInfoln(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerInfoln", logrus.InfoLevel)
	logger.Infoln("TestLoggerInfoln")
	HelperTestVerify(t, buffer, testDataPath + "TestLoggerInfoln.expected.json")
}

func TestLoggerInfolnLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerInfolnLevelSuppressed", logrus.WarnLevel)
	logger.Info("TestLoggerInfolnLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerWarn(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarn", logrus.WarnLevel)
	logger.Warn("TestLoggerWarn")
	HelperTestVerify(t, buffer, testDataPath + "TestLoggerWarn.expected.json")
}

func TestLoggerWarnLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarnLevelSuppressed", logrus.ErrorLevel)
	logger.Warn("TestLoggerWarnLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerWarnf(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarnf", logrus.WarnLevel)
	logger.Warnf("%sLogger%s", "Test", "Warnf")
	HelperTestVerify(t, buffer, testDataPath + "TestLoggerWarnf.expected.json")
}

func TestLoggerWarnfLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarnfLevelSuppressed", logrus.ErrorLevel)
	logger.Warnf("TestLoggerWarnfLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerWarnln(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarnln", logrus.WarnLevel)
	logger.Warnln("TestLoggerWarnln")
	HelperTestVerify(t, buffer, testDataPath + "TestLoggerWarnln.expected.json")
}

func TestLoggerWarnlnLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarnlnLevelSuppressed", logrus.ErrorLevel)
	logger.Warn("TestLoggerWarnlnLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerWarning(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarning", logrus.WarnLevel)
	logger.Warning("TestLoggerWarning")
	HelperTestVerify(t, buffer, testDataPath + "TestLoggerWarning.expected.json")
}

func TestLoggerWarningLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarningLevelSuppressed", logrus.ErrorLevel)
	logger.Warning("TestLoggerWarningLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerWarningf(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarningf", logrus.WarnLevel)
	logger.Warningf("%sLogger%s", "Test", "Warningf")
	HelperTestVerify(t, buffer, testDataPath + "TestLoggerWarningf.expected.json")
}

func TestLoggerWarningfLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarningfLevelSuppressed", logrus.ErrorLevel)
	logger.Warningf("TestLoggerWarningfLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerWarningln(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarningln", logrus.WarnLevel)
	logger.Warningln("TestLoggerWarningln")
	HelperTestVerify(t, buffer, testDataPath + "TestLoggerWarningln.expected.json")
}

func TestLoggerWarninglnLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarninglnLevelSuppressed", logrus.ErrorLevel)
	logger.Warning("TestLoggerWarninglnLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerError(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerError", logrus.ErrorLevel)
	logger.Error("TestLoggerError")
	HelperTestVerify(t, buffer, testDataPath + "TestLoggerError.expected.json")
}

func TestLoggerErrorLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerErrorLevelSuppressed", logrus.FatalLevel)
	logger.Error("TestLoggerErrorLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerErrorf(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerErrorf", logrus.ErrorLevel)
	logger.Errorf("%sLogger%s", "Test", "Errorf")
	HelperTestVerify(t, buffer, testDataPath + "TestLoggerErrorf.expected.json")
}

func TestLoggerErrorfLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerErrorfLevelSuppressed", logrus.FatalLevel)
	logger.Errorf("TestLoggerErrorfLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerErrorln(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerErrorln", logrus.ErrorLevel)
	logger.Errorln("TestLoggerErrorln")
	HelperTestVerify(t, buffer, testDataPath + "TestLoggerErrorln.expected.json")
}

func TestLoggerErrorlnLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerErrorlnLevelSuppressed", logrus.FatalLevel)
	logger.Error("TestLoggerErrorlnLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerWithField(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWithField", logrus.InfoLevel)
	logger.WithField("foo", "bar").Info("TestLoggerWithField")
	HelperTestVerify(t, buffer, testDataPath + "TestLoggerWithField.expected.json")
}

func TestLoggerWithFields(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWithFields", logrus.InfoLevel)
	logger.WithFields(logrus.Fields{"foo": "bar", "one": 1, "pi": 3.14,}).Info("TestLoggerWithFields")
	HelperTestVerify(t, buffer, testDataPath + "TestLoggerWithFields.expected.json")
}
