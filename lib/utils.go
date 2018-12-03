/**
 * Go Video Downloader
 *
 *    Copyright 2018 Tenta, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * For any questions, please contact developer@tenta.io
 *
 * utils.go: Miscellaneous utilities
 */

package lib

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

// StringlikeVal extracts the string value from a str or bytes typed object.
// Returns true in the second value if the string was a str.
func StringlikeVal(o rnt.Object) (string, bool) {
	if o.IsInstance(rnt.BytesType) {
		return string(o.(rnt.Bytes).Value()), false
	}
	return o.(rnt.Str).Value(), true
}

// NewStringlike builds a str or bytes object having val as value,
// depending on st (True->str, False->bytes).
func NewStringlike(val string, st bool) rnt.Object {
	if !st {
		return rnt.NewBytes([]byte(val)...)
	}
	return rnt.NewStr(val)
}
