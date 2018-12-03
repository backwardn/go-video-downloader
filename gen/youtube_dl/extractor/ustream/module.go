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
 * ustream/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/ustream.py
 */

package ustream

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError   λ.Object
	InfoExtractor    λ.Object
	UstreamChannelIE λ.Object
	UstreamIE        λ.Object
	ϒcompat_str      λ.Object
	ϒfloat_or_none   λ.Object
	ϒint_or_none     λ.Object
	ϒmimetype2ext    λ.Object
	ϒstr_or_none     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒmimetype2ext = Ωutils.ϒmimetype2ext
		ϒstr_or_none = Ωutils.ϒstr_or_none
		UstreamIE = λ.Cal(λ.TypeType, λ.NewStr("UstreamIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				UstreamIE__VALID_URL λ.Object
			)
			UstreamIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?ustream\\.tv/(?P<type>recorded|embed|embed/recorded)/(?P<id>\\d+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): UstreamIE__VALID_URL,
			})
		}())
		UstreamChannelIE = λ.Cal(λ.TypeType, λ.NewStr("UstreamChannelIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				UstreamChannelIE__VALID_URL λ.Object
			)
			UstreamChannelIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?ustream\\.tv/channel/(?P<slug>.+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): UstreamChannelIE__VALID_URL,
			})
		}())
	})
}
