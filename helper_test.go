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

var (
	stenoSchemaLoader = gojsonschema.NewReferenceLoader("file://./testdata/steno.schema.json")
)

func HelperTestGetLogger(n string, l logrus.Level, f *Formatter) (*Logger, *bytes.Buffer) {
	var buffer *bytes.Buffer = new(bytes.Buffer)
	var logrusLogger *logrus.Logger = &logrus.Logger{
		Out: buffer,
		Formatter: f,
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
	HelperTestVerifyIgnoreContext(t, actualBuffer, actualFile, []string{})
}

func HelperTestVerifyIgnoreContext(t *testing.T, actualBuffer *bytes.Buffer, actualFile string, ignoreContextKeysForSchema []string) {
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
	var actualForValidationRootNode map[string]interface{}
	if err = json.Unmarshal([]byte(actualAsString), &actualForValidationRootNode); err != nil {
		t.Errorf("Unmarshal of actual failed because %v in buffer %s", err, actualAsString)
		return
	}
	hideIgnoredKeys(actualForValidationRootNode, ignoreContextKeysForSchema)
	var actualForValidationAsByteArray []byte
	if actualForValidationAsByteArray, err = json.Marshal(actualForValidationRootNode); err != nil {
		t.Errorf("Remarshal of actual failed because %v for %s", err, actualForValidationRootNode)
		return
	}
	var actualForValidationJsonLoader = gojsonschema.NewStringLoader(string(actualForValidationAsByteArray))
	if result, err = gojsonschema.Validate(stenoSchemaLoader, actualForValidationJsonLoader); err != nil {
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

func hideIgnoredKeys(r map[string]interface{}, ignoreContextKeys []string) {
	if node, ok := r["context"].(map[string]interface{}); ok {
		for i := range ignoreContextKeys {
			delete(node, ignoreContextKeys[i])
		}
	}
}

func normalize(r map[string]interface{}) {
	if _, ok := r["id"]; ok {
		r["id"] = "<ID>"
	}
	if _, ok := r["time"]; ok {
		r["time"] = "<TIME>"
	}
	if node, ok := r["context"].(map[string]interface{}); ok {
		if _, ok := node["host"]; ok {
			node["host"] = "<HOST>"
		}
		if _, ok := node["processId"]; ok {
			node["processId"] = "<PROCESS_ID>"
		}
	}
}
