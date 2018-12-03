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
 * jsinterp/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/jsinterp.py
 */

package jsinterp

import (
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError     λ.Object
	ϒ_ASSIGN_OPERATORS λ.Object
	ϒ_OPERATORS        λ.Object
)

func init() {
	λ.InitModule(func() {
		ExtractorError = Ωutils.ExtractorError
		ϒ_OPERATORS = λ.NewList(
			λ.NewTuple(
				λ.NewStr("|"),
				λ.None,
			),
			λ.NewTuple(
				λ.NewStr("^"),
				λ.None,
			),
			λ.NewTuple(
				λ.NewStr("&"),
				λ.None,
			),
			λ.NewTuple(
				λ.NewStr(">>"),
				λ.None,
			),
			λ.NewTuple(
				λ.NewStr("<<"),
				λ.None,
			),
			λ.NewTuple(
				λ.NewStr("-"),
				λ.None,
			),
			λ.NewTuple(
				λ.NewStr("+"),
				λ.None,
			),
			λ.NewTuple(
				λ.NewStr("%"),
				λ.None,
			),
			λ.NewTuple(
				λ.NewStr("/"),
				λ.None,
			),
			λ.NewTuple(
				λ.NewStr("*"),
				λ.None,
			),
		)
		ϒ_ASSIGN_OPERATORS = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
			nil,
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
					var (
						ϒop     λ.Object
						ϒopfunc λ.Object
						τmp0    λ.Object
						τmp1    λ.Object
						τmp2    λ.Object
					)
					τmp0 = λ.Cal(λ.BuiltinIter, ϒ_OPERATORS)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒop = λ.GetItem(τmp2, λ.NewInt(0))
						ϒopfunc = λ.GetItem(τmp2, λ.NewInt(1))
						λgen.Yield(λ.NewTuple(
							λ.Add(ϒop, λ.NewStr("=")),
							ϒopfunc,
						))
					}
					return λ.None
				})
			})))
		λ.Cal(λ.GetAttr(ϒ_ASSIGN_OPERATORS, "append", nil), λ.NewTuple(
			λ.NewStr("="),
			λ.NewFunction("<lambda>",
				[]λ.Param{
					{Name: "cur"},
					{Name: "right"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcur   = λargs[0]
						ϒright = λargs[1]
					)
					_ = ϒcur
					return ϒright
				}),
		))
	})
}
