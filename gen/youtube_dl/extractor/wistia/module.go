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
 * wistia/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/wistia.py
 */

package wistia

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError λ.Object
	InfoExtractor  λ.Object
	WistiaIE       λ.Object
	ϒfloat_or_none λ.Object
	ϒint_or_none   λ.Object
	ϒunescapeHTML  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒunescapeHTML = Ωutils.ϒunescapeHTML
		WistiaIE = λ.Cal(λ.TypeType, λ.StrLiteral("WistiaIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				WistiaIE__VALID_URL λ.Object
			)
			WistiaIE__VALID_URL = λ.StrLiteral("(?:wistia:|https?://(?:fast\\.)?wistia\\.(?:net|com)/embed/(?:iframe|medias)/)(?P<id>[a-z0-9]{10})")
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL": WistiaIE__VALID_URL,
			})
		}())
	})
}
