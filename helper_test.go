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
	"io/ioutil"
	"encoding/json"
	"reflect"
	"testing"
	"github.com/Sirupsen/logrus"
	"github.com/xeipuuv/gojsonschema"
)

func HelperTestGetLogger(n string, l logrus.Level) (*Logger, *bytes.Buffer) {
	var buffer *bytes.Buffer = new(bytes.Buffer)
	var logrusLogger *logrus.Logger = &logrus.Logger{
		Out: buffer,
		Formatter: formatter,
		Level: l,
	}
	return GetLoggerForLogger(n, logrusLogger), buffer
}

func HelperTestVerifyEmpty(t *testing.T, actualBuffer *bytes.Buffer) {
	if actualBuffer.Len() > 0 {
		t.Errorf("Expected actual buffer to be empty but contains %s", actualBuffer.String())
		return
	}
}

func HelperTestVerify(t *testing.T, actualBuffer *bytes.Buffer, actualFile string) {
	var err error
	var expectedBuffer []byte
	var result *gojsonschema.Result;
	if expectedBuffer, err = ioutil.ReadFile(actualFile); err != nil {
		t.Errorf("Failed to read actual file %s", actualFile)
		return
	}
	var expectedRootNode map[string]interface{}
	if err = json.Unmarshal(expectedBuffer, &expectedRootNode); err != nil {
		t.Errorf("Unmarshal of expected failed because %v in buffer %s", err, string(expectedBuffer[:]))
		return
	}
	var actualAsString string = actualBuffer.String()
	var actualRootNode map[string]interface{}
	if err = json.Unmarshal([]byte(actualAsString), &actualRootNode); err != nil {
		t.Errorf("Unmarshal of actual failed because %v in buffer %s", err, actualAsString)
		return
	}
	var actualJsonLoader = gojsonschema.NewStringLoader(actualAsString)
	if result, err = gojsonschema.Validate(stenoSchemaLoader, actualJsonLoader); err != nil {
		t.Errorf("Validation against json schema failed because %v", err)
		return
	} else if !result.Valid() {
		t.Errorf("Actual log message is not valid steno because %v actual is %s", result.Errors(), actualAsString)
		return
	}
	normalize(actualRootNode)
	if !reflect.DeepEqual(expectedRootNode, actualRootNode) {
		t.Errorf("Actual log message does not match expected; expected is %v but actual was %v ", expectedRootNode, actualRootNode)
		return
	}
}

func normalize(r map[string]interface{}) {
	r["id"] = "<ID>"
	r["time"] = "<TIME>"
	if node, ok := r["context"].(map[string]interface{}); ok {
		node["host"] = "<HOST>"
		node["processId"] = "<PROCESS_ID>"
	}
}
