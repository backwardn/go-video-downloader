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
 * instagram/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/instagram.py
 */

package instagram

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError            λ.Object
	InfoExtractor             λ.Object
	InstagramIE               λ.Object
	InstagramPlaylistIE       λ.Object
	InstagramTagIE            λ.Object
	InstagramUserIE           λ.Object
	ϒcompat_HTTPError         λ.Object
	ϒcompat_str               λ.Object
	ϒget_element_by_attribute λ.Object
	ϒint_or_none              λ.Object
	ϒlowercase_escape         λ.Object
	ϒstd_headers              λ.Object
	ϒtry_get                  λ.Object
	ϒurl_or_none              λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒcompat_HTTPError = Ωcompat.ϒcompat_HTTPError
		ExtractorError = Ωutils.ExtractorError
		ϒget_element_by_attribute = Ωutils.ϒget_element_by_attribute
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒlowercase_escape = Ωutils.ϒlowercase_escape
		ϒstd_headers = Ωutils.ϒstd_headers
		ϒtry_get = Ωutils.ϒtry_get
		ϒurl_or_none = Ωutils.ϒurl_or_none
		InstagramIE = λ.Cal(λ.TypeType, λ.StrLiteral("InstagramIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				InstagramIE__VALID_URL λ.Object
			)
			InstagramIE__VALID_URL = λ.StrLiteral("(?P<url>https?://(?:www\\.)?instagram\\.com/(?:p|tv)/(?P<id>[^/?#&]+))")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": InstagramIE__VALID_URL,
			})
		}())
		InstagramPlaylistIE = λ.Cal(λ.TypeType, λ.StrLiteral("InstagramPlaylistIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		InstagramUserIE = λ.Cal(λ.TypeType, λ.StrLiteral("InstagramUserIE"), λ.NewTuple(InstagramPlaylistIE), func() λ.Dict {
			var (
				InstagramUserIE__VALID_URL λ.Object
			)
			InstagramUserIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?instagram\\.com/(?P<id>[^/]{2,})/?(?:$|[?#])")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": InstagramUserIE__VALID_URL,
			})
		}())
		InstagramTagIE = λ.Cal(λ.TypeType, λ.StrLiteral("InstagramTagIE"), λ.NewTuple(InstagramPlaylistIE), func() λ.Dict {
			var (
				InstagramTagIE__VALID_URL λ.Object
			)
			InstagramTagIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?instagram\\.com/explore/tags/(?P<id>[^/]+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": InstagramTagIE__VALID_URL,
			})
		}())
	})
}
