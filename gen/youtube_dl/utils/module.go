// Code generated by transpiler. DO NOT EDIT.

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
 * utils/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/utils.py
 */

package utils

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ϒcompat_HTMLParseError λ.Object
	τmp0                   λ.Object
	τmp1                   λ.Object
)

func init() {
	λ.InitModule(func() {
		ϒcompat_HTMLParseError = Ωcompat.ϒcompat_HTMLParseError
		λ.Cal(λ.GetAttr(λ.None, "extend", nil), λ.NewList(
			λ.NewStr("%d-%m-%Y"),
			λ.NewStr("%d.%m.%Y"),
			λ.NewStr("%d.%m.%y"),
			λ.NewStr("%d/%m/%Y"),
			λ.NewStr("%d/%m/%y"),
			λ.NewStr("%d/%m/%Y %H:%M:%S"),
		))
		λ.Cal(λ.GetAttr(λ.None, "extend", nil), λ.NewList(
			λ.NewStr("%m-%d-%Y"),
			λ.NewStr("%m.%d.%Y"),
			λ.NewStr("%m/%d/%Y"),
			λ.NewStr("%m/%d/%y"),
			λ.NewStr("%m/%d/%Y %H:%M:%S"),
		))
		if λ.IsTrue(λ.Ge(λ.None, λ.NewTuple(
			λ.NewInt(2),
			λ.NewInt(7),
		))) {

		} else {

		}
		if λ.IsTrue(λ.Eq(λ.None, λ.NewStr("win32"))) {
			λ.SetAttr(λ.None, "argtypes", λ.NewList(
				λ.GetAttr(λ.GetAttr(λ.None, "wintypes", nil), "HANDLE", nil),
				λ.GetAttr(λ.GetAttr(λ.None, "wintypes", nil), "DWORD", nil),
				λ.GetAttr(λ.GetAttr(λ.None, "wintypes", nil), "DWORD", nil),
				λ.GetAttr(λ.GetAttr(λ.None, "wintypes", nil), "DWORD", nil),
				λ.GetAttr(λ.GetAttr(λ.None, "wintypes", nil), "DWORD", nil),
				λ.Cal(λ.GetAttr(λ.None, "POINTER", nil), λ.None),
			))
			λ.SetAttr(λ.None, "restype", λ.GetAttr(λ.GetAttr(λ.None, "wintypes", nil), "BOOL", nil))
			λ.SetAttr(λ.None, "argtypes", λ.NewList(
				λ.GetAttr(λ.GetAttr(λ.None, "wintypes", nil), "HANDLE", nil),
				λ.GetAttr(λ.GetAttr(λ.None, "wintypes", nil), "DWORD", nil),
				λ.GetAttr(λ.GetAttr(λ.None, "wintypes", nil), "DWORD", nil),
				λ.GetAttr(λ.GetAttr(λ.None, "wintypes", nil), "DWORD", nil),
				λ.Cal(λ.GetAttr(λ.None, "POINTER", nil), λ.None),
			))
			λ.SetAttr(λ.None, "restype", λ.GetAttr(λ.GetAttr(λ.None, "wintypes", nil), "BOOL", nil))
		} else {
			τmp1, τmp0 = func() (λexit λ.Object, λret λ.Object) {
				defer λ.CatchMulti(
					nil,
					&λ.Catcher{λ.ImportErrorType, func(λex λ.BaseException) {

					}},
				)

				return λ.BlockExitNormally, nil
			}()
		}
	})
}
