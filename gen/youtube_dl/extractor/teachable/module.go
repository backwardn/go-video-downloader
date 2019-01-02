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
 * teachable/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/teachable.py
 */

package teachable

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωwistia "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/wistia"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError        λ.Object
	InfoExtractor         λ.Object
	TeachableBaseIE       λ.Object
	WistiaIE              λ.Object
	ϒclean_html           λ.Object
	ϒcompat_str           λ.Object
	ϒget_element_by_class λ.Object
	ϒurlencode_postdata   λ.Object
	ϒurljoin              λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		WistiaIE = Ωwistia.WistiaIE
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒclean_html = Ωutils.ϒclean_html
		ExtractorError = Ωutils.ExtractorError
		ϒget_element_by_class = Ωutils.ϒget_element_by_class
		ϒurlencode_postdata = Ωutils.ϒurlencode_postdata
		ϒurljoin = Ωutils.ϒurljoin
		TeachableBaseIE = λ.Cal(λ.TypeType, λ.NewStr("TeachableBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TeachableBaseIE__SITES               λ.Object
				TeachableBaseIE__URL_PREFIX          λ.Object
				TeachableBaseIE__VALID_URL_SUB_TUPLE λ.Object
			)
			TeachableBaseIE__URL_PREFIX = λ.NewStr("teachable:")
			TeachableBaseIE__SITES = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("upskillcourses.com"):     λ.NewStr("upskill"),
				λ.NewStr("academy.gns3.com"):       λ.NewStr("gns3"),
				λ.NewStr("academyhacker.com"):      λ.NewStr("academyhacker"),
				λ.NewStr("stackskills.com"):        λ.NewStr("stackskills"),
				λ.NewStr("market.saleshacker.com"): λ.NewStr("saleshacker"),
				λ.NewStr("learnability.org"):       λ.NewStr("learnability"),
				λ.NewStr("edurila.com"):            λ.NewStr("edurila"),
			})
			TeachableBaseIE__VALID_URL_SUB_TUPLE = λ.NewTuple(
				TeachableBaseIE__URL_PREFIX,
				λ.Cal(λ.GetAttr(λ.NewStr("|"), "join", nil), λ.Cal(λ.NewFunction("<generator>",
					nil,
					0, false, false,
					func(λargs []λ.Object) λ.Object {
						return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
							var (
								ϒsite λ.Object
								τmp0  λ.Object
								τmp1  λ.Object
							)
							τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(TeachableBaseIE__SITES, "keys", nil)))
							for {
								if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
									break
								}
								ϒsite = τmp1
								λgen.Yield(λ.Cal(Ωre.ϒescape, ϒsite))
							}
							return λ.None
						})
					}))),
			)
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_SITES"):               TeachableBaseIE__SITES,
				λ.NewStr("_URL_PREFIX"):          TeachableBaseIE__URL_PREFIX,
				λ.NewStr("_VALID_URL_SUB_TUPLE"): TeachableBaseIE__VALID_URL_SUB_TUPLE,
			})
		}())
	})
}