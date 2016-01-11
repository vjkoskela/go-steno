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
	"testing"
)

func TestNoOpLogBuilder(t *testing.T) {
	t.Parallel()
	var nolb LogBuilder = new(NoOpLogBuilder)
	var r LogBuilder
	if r = nolb.SetEvent("event"); r != nolb {
		t.Error("SetEvent did not return nolb")
	}
	if r = nolb.SetError(errors.New("an error")); r != nolb {
		t.Error("SetError did not return nolb")
	}
	if r = nolb.SetMessage("message"); r != nolb {
		t.Error("SetMessage did not return nolb")
	}
	if r = nolb.AddData("k", "v"); r != nolb {
		t.Error("AddData did not return nolb")
	}
	if r = nolb.AddContext("k", "v"); r != nolb {
		t.Error("AddData did not return nolb")
	}
}