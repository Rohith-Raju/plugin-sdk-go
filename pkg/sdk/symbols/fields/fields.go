/*
Copyright (C) 2021 The Falco Authors.

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

package fields

import "C"
import (
	"encoding/json"

	"github.com/falcosecurity/plugin-sdk-go/pkg/ptr"
	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk"
)

var (
	fields []sdk.FieldEntry
	buf    ptr.StringBuffer
)

func SetFields(f []sdk.FieldEntry) {
	fields = f
}

func Fields() []sdk.FieldEntry {
	return fields
}

//export plugin_get_fields
func plugin_get_fields() *C.char {
	b, err := json.Marshal(&fields)
	if err != nil {
		return nil
	}
	buf.Write(string(b))
	return (*C.char)(buf.CharPtr())
}