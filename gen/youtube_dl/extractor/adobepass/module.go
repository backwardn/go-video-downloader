// Code generated by transpiler. DO NOT EDIT.

/**
 * Go Video Downloader
 *
 *    Copyright 2019 Tenta, LLC
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
 * adobepass/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/adobepass.py
 */

package adobepass

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	AdobePassIE         λ.Object
	ExtractorError      λ.Object
	InfoExtractor       λ.Object
	NO_DEFAULT          λ.Object
	ϒcompat_kwargs      λ.Object
	ϒunescapeHTML       λ.Object
	ϒunified_timestamp  λ.Object
	ϒurlencode_postdata λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_kwargs = Ωcompat.ϒcompat_kwargs
		ϒunescapeHTML = Ωutils.ϒunescapeHTML
		ϒurlencode_postdata = Ωutils.ϒurlencode_postdata
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ExtractorError = Ωutils.ExtractorError
		NO_DEFAULT = Ωutils.NO_DEFAULT
		AdobePassIE = λ.Cal(λ.TypeType, λ.StrLiteral("AdobePassIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				AdobePassIE__download_webpage_handle λ.Object
			)
			AdobePassIE__download_webpage_handle = λ.NewFunction("_download_webpage_handle",
				[]λ.Param{
					{Name: "self"},
				},
				0, true, true,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒargs    = λargs[1]
						ϒheaders λ.Object
						ϒkwargs  = λargs[2]
						ϒself    = λargs[0]
					)
					ϒheaders = λ.Calm(ϒself, "geo_verification_headers")
					λ.Calm(ϒheaders, "update", λ.Calm(ϒkwargs, "get", λ.StrLiteral("headers"), λ.DictLiteral(map[λ.Object]λ.Object{})))
					λ.SetItem(ϒkwargs, λ.StrLiteral("headers"), ϒheaders)
					return λ.Call(λ.GetAttr(λ.Cal(λ.SuperType, AdobePassIE, ϒself), "_download_webpage_handle", nil), λ.NewArgs(λ.Unpack(λ.AsStarred(ϒargs))...), λ.KWArgs{
						{Name: "", Value: λ.Cal(ϒcompat_kwargs, ϒkwargs)},
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_download_webpage_handle": AdobePassIE__download_webpage_handle,
			})
		}())
	})
}
