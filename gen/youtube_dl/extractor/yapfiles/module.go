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
 * yapfiles/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/yapfiles.py
 */

package yapfiles

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError λ.Object
	InfoExtractor  λ.Object
	YapFilesIE     λ.Object
	ϒint_or_none   λ.Object
	ϒqualities     λ.Object
	ϒunescapeHTML  λ.Object
	ϒurl_or_none   λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒqualities = Ωutils.ϒqualities
		ϒunescapeHTML = Ωutils.ϒunescapeHTML
		ϒurl_or_none = Ωutils.ϒurl_or_none
		YapFilesIE = λ.Cal(λ.TypeType, λ.NewStr("YapFilesIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				YapFilesIE__VALID_URL    λ.Object
				YapFilesIE__YAPFILES_URL λ.Object
			)
			YapFilesIE__YAPFILES_URL = λ.NewStr("//(?:(?:www|api)\\.)?yapfiles\\.ru/get_player/*\\?.*?\\bv=(?P<id>\\w+)")
			YapFilesIE__VALID_URL = λ.Mod(λ.NewStr("https?:%s"), YapFilesIE__YAPFILES_URL)
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"):    YapFilesIE__VALID_URL,
				λ.NewStr("_YAPFILES_URL"): YapFilesIE__YAPFILES_URL,
			})
		}())
	})
}
