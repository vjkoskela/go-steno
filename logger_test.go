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

// TODO:
// 1) Test Fatal and Panic methods using a mock logrus logger.
// 2) Test no-op stub methods from standard logger.
// 3) Test output method to info w/o stack trace method from standard logger.

const (
	loggerTestDataPath = "./testdata/logger_test/"
)

var (
	loggerTestFormatter *Formatter = NewFormatter()
)

func TestLoggerDebugBuilder(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerDebugBuilder", logrus.DebugLevel, loggerTestFormatter)
	logger.DebugBuilder().SetMessage("TestLoggerDebugBuilder").Log()
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerDebugBuilder.expected.json")
}

func TestLoggerDebugBuilderLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerDebugBuilderLevelSuppressed", logrus.InfoLevel, loggerTestFormatter)
	logger.DebugBuilder().SetMessage("TestLoggerDebugBuilderLevelSuppressed").Log()
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerInfoBuilder(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerInfoBuilder", logrus.InfoLevel, loggerTestFormatter)
	logger.InfoBuilder().SetMessage("TestLoggerInfoBuilder").Log()
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerInfoBuilder.expected.json")
}

func TestLoggerInfoBuilderLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerInfoBuilderLevelSuppressed", logrus.WarnLevel, loggerTestFormatter)
	logger.InfoBuilder().SetMessage("TestLoggerInfoBuilderLevelSuppressed").Log()
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerWarnBuilder(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarnBuilder", logrus.WarnLevel, loggerTestFormatter)
	logger.WarnBuilder().SetMessage("TestLoggerWarnBuilder").Log()
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerWarnBuilder.expected.json")
}

func TestLoggerWarnBuilderLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarnBuilderLevelSuppressed", logrus.ErrorLevel, loggerTestFormatter)
	logger.WarnBuilder().SetMessage("TestLoggerWarnBuilderLevelSuppressed").Log()
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerWarningBuilder(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarningBuilder", logrus.WarnLevel, loggerTestFormatter)
	logger.WarningBuilder().SetMessage("TestLoggerWarningBuilder").Log()
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerWarningBuilder.expected.json")
}

func TestLoggerWarningBuilderLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarningBuilderLevelSuppressed", logrus.ErrorLevel, loggerTestFormatter)
	logger.WarningBuilder().SetMessage("TestLoggerWarningBuilderLevelSuppressed").Log()
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerErrorBuilder(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerErrorBuilder", logrus.ErrorLevel, loggerTestFormatter)
	logger.ErrorBuilder().SetMessage("TestLoggerErrorBuilder").Log()
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerErrorBuilder.expected.json")
}

func TestLoggerErrorBuilderLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerErrorBuilderLevelSuppressed", logrus.FatalLevel, loggerTestFormatter)
	logger.ErrorBuilder().SetMessage("TestLoggerErrorBuilderLevelSuppressed").Log()
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerPrint(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerPrint", logrus.InfoLevel, loggerTestFormatter)
	logger.Print("TestLoggerPrint")
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerPrint.expected.json")
}

func TestLoggerPrintLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerPrintLevelSuppressed", logrus.WarnLevel, loggerTestFormatter)
	logger.Print("TestLoggerPrintLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerPrintf(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerPrintf", logrus.InfoLevel, loggerTestFormatter)
	logger.Printf("%sLogger%s", "Test", "Printf")
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerPrintf.expected.json")
}

func TestLoggerPrintfLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerPrintfLevelSuppressed", logrus.WarnLevel, loggerTestFormatter)
	logger.Printf("TestLoggerPrintfLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerPrintln(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerPrintln", logrus.InfoLevel, loggerTestFormatter)
	logger.Println("TestLoggerPrintln")
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerPrintln.expected.json")
}

func TestLoggerPrintlnLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerPrintlnLevelSuppressed", logrus.WarnLevel, loggerTestFormatter)
	logger.Print("TestLoggerPrintlnLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerDebug(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerDebug", logrus.DebugLevel, loggerTestFormatter)
	logger.Debug("TestLoggerDebug")
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerDebug.expected.json")
}

func TestLoggerDebugLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerDebugLevelSuppressed", logrus.InfoLevel, loggerTestFormatter)
	logger.Debug("TestLoggerDebugLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerDebugf(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerDebugf", logrus.DebugLevel, loggerTestFormatter)
	logger.Debugf("%sLogger%s", "Test", "Debugf")
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerDebugf.expected.json")
}

func TestLoggerDebugfLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerDebugfLevelSuppressed", logrus.InfoLevel, loggerTestFormatter)
	logger.Debugf("TestLoggerDebugfLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerDebugln(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerDebugln", logrus.DebugLevel, loggerTestFormatter)
	logger.Debugln("TestLoggerDebugln")
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerDebugln.expected.json")
}

func TestLoggerDebuglnLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerDebuglnLevelSuppressed", logrus.InfoLevel, loggerTestFormatter)
	logger.Debug("TestLoggerDebuglnLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerInfo(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerInfo", logrus.InfoLevel, loggerTestFormatter)
	logger.Info("TestLoggerInfo")
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerInfo.expected.json")
}

func TestLoggerInfoLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerInfoLevelSuppressed", logrus.WarnLevel, loggerTestFormatter)
	logger.Info("TestLoggerInfoLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerInfof(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerInfof", logrus.InfoLevel, loggerTestFormatter)
	logger.Infof("%sLogger%s", "Test", "Infof")
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerInfof.expected.json")
}

func TestLoggerInfofLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerInfofLevelSuppressed", logrus.WarnLevel, loggerTestFormatter)
	logger.Infof("TestLoggerInfofLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerInfoln(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerInfoln", logrus.InfoLevel, loggerTestFormatter)
	logger.Infoln("TestLoggerInfoln")
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerInfoln.expected.json")
}

func TestLoggerInfolnLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerInfolnLevelSuppressed", logrus.WarnLevel, loggerTestFormatter)
	logger.Info("TestLoggerInfolnLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerWarn(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarn", logrus.WarnLevel, loggerTestFormatter)
	logger.Warn("TestLoggerWarn")
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerWarn.expected.json")
}

func TestLoggerWarnLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarnLevelSuppressed", logrus.ErrorLevel, loggerTestFormatter)
	logger.Warn("TestLoggerWarnLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerWarnf(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarnf", logrus.WarnLevel, loggerTestFormatter)
	logger.Warnf("%sLogger%s", "Test", "Warnf")
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerWarnf.expected.json")
}

func TestLoggerWarnfLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarnfLevelSuppressed", logrus.ErrorLevel, loggerTestFormatter)
	logger.Warnf("TestLoggerWarnfLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerWarnln(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarnln", logrus.WarnLevel, loggerTestFormatter)
	logger.Warnln("TestLoggerWarnln")
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerWarnln.expected.json")
}

func TestLoggerWarnlnLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarnlnLevelSuppressed", logrus.ErrorLevel, loggerTestFormatter)
	logger.Warn("TestLoggerWarnlnLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerWarning(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarning", logrus.WarnLevel, loggerTestFormatter)
	logger.Warning("TestLoggerWarning")
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerWarning.expected.json")
}

func TestLoggerWarningLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarningLevelSuppressed", logrus.ErrorLevel, loggerTestFormatter)
	logger.Warning("TestLoggerWarningLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerWarningf(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarningf", logrus.WarnLevel, loggerTestFormatter)
	logger.Warningf("%sLogger%s", "Test", "Warningf")
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerWarningf.expected.json")
}

func TestLoggerWarningfLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarningfLevelSuppressed", logrus.ErrorLevel, loggerTestFormatter)
	logger.Warningf("TestLoggerWarningfLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerWarningln(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarningln", logrus.WarnLevel, loggerTestFormatter)
	logger.Warningln("TestLoggerWarningln")
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerWarningln.expected.json")
}

func TestLoggerWarninglnLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWarninglnLevelSuppressed", logrus.ErrorLevel, loggerTestFormatter)
	logger.Warning("TestLoggerWarninglnLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerError(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerError", logrus.ErrorLevel, loggerTestFormatter)
	logger.Error("TestLoggerError")
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerError.expected.json")
}

func TestLoggerErrorLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerErrorLevelSuppressed", logrus.FatalLevel, loggerTestFormatter)
	logger.Error("TestLoggerErrorLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerErrorf(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerErrorf", logrus.ErrorLevel, loggerTestFormatter)
	logger.Errorf("%sLogger%s", "Test", "Errorf")
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerErrorf.expected.json")
}

func TestLoggerErrorfLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerErrorfLevelSuppressed", logrus.FatalLevel, loggerTestFormatter)
	logger.Errorf("TestLoggerErrorfLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerErrorln(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerErrorln", logrus.ErrorLevel, loggerTestFormatter)
	logger.Errorln("TestLoggerErrorln")
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerErrorln.expected.json")
}

func TestLoggerErrorlnLevelSuppressed(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerErrorlnLevelSuppressed", logrus.FatalLevel, loggerTestFormatter)
	logger.Error("TestLoggerErrorlnLevelSuppressed")
	HelperTestVerifyEmpty(t, buffer)
}

func TestLoggerWithField(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWithField", logrus.InfoLevel, loggerTestFormatter)
	logger.WithField("foo", "bar").Info("TestLoggerWithField")
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerWithField.expected.json")
}

func TestLoggerWithFields(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWithFields", logrus.InfoLevel, loggerTestFormatter)
	logger.WithFields(logrus.Fields{"foo": "bar", "one": 1, "pi": 3.14,}).Info("TestLoggerWithFields")
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerWithFields.expected.json")
}

func TestLoggerWithError(t *testing.T) {
	t.Parallel()
	logger, buffer := HelperTestGetLogger("TestLoggerWithError", logrus.InfoLevel, loggerTestFormatter)
	logger.WithError(errors.New("This is an error")).Info("TestLoggerWithError")
	HelperTestVerify(t, buffer, loggerTestDataPath + "TestLoggerWithError.expected.json")
}
