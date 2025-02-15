// Copyright © 2022 The sealos Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"github.com/labring/endpoints-operator/library/convert"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func runtimeConvertUnstructured(from runtime.Object) *unstructured.Unstructured {
	to, ok := from.(*unstructured.Unstructured)
	if ok {
		return to
	}
	if to, err := convert.ResourceToUnstructured(from); err == nil {
		return to
	}
	return nil
}
