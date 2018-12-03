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
 * nexx/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/nexx.py
 */

package nexx

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError      λ.Object
	InfoExtractor       λ.Object
	NexxEmbedIE         λ.Object
	NexxIE              λ.Object
	ϒcompat_str         λ.Object
	ϒint_or_none        λ.Object
	ϒparse_duration     λ.Object
	ϒtry_get            λ.Object
	ϒurlencode_postdata λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒtry_get = Ωutils.ϒtry_get
		ϒurlencode_postdata = Ωutils.ϒurlencode_postdata
		NexxIE = λ.Cal(λ.TypeType, λ.NewStr("NexxIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				NexxIE__VALID_URL λ.Object
			)
			NexxIE__VALID_URL = λ.NewStr("(?x)\n                        (?:\n                            https?://api\\.nexx(?:\\.cloud|cdn\\.com)/v3/(?P<domain_id>\\d+)/videos/byid/|\n                            nexx:(?:(?P<domain_id_s>\\d+):)?|\n                            https?://arc\\.nexx\\.cloud/api/video/\n                        )\n                        (?P<id>\\d+)\n                    ")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): NexxIE__VALID_URL,
			})
		}())
		NexxEmbedIE = λ.Cal(λ.TypeType, λ.NewStr("NexxEmbedIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				NexxEmbedIE__VALID_URL λ.Object
			)
			NexxEmbedIE__VALID_URL = λ.NewStr("https?://embed\\.nexx(?:\\.cloud|cdn\\.com)/\\d+/(?P<id>[^/?#&]+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): NexxEmbedIE__VALID_URL,
			})
		}())
	})
}
