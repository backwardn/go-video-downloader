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
 * stretchinternet/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/stretchinternet.py
 */

package stretchinternet

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor     λ.Object
	StretchInternetIE λ.Object
	ϒint_or_none      λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒint_or_none = Ωutils.ϒint_or_none
		StretchInternetIE = λ.Cal(λ.TypeType, λ.StrLiteral("StretchInternetIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				StretchInternetIE__VALID_URL    λ.Object
				StretchInternetIE__real_extract λ.Object
			)
			StretchInternetIE__VALID_URL = λ.StrLiteral("https?://portal\\.stretchinternet\\.com/[^/]+/(?:portal|full)\\.htm\\?.*?\\beventId=(?P<id>\\d+)")
			StretchInternetIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒevent    λ.Object
						ϒself     = λargs[0]
						ϒurl      = λargs[1]
						ϒvideo_id λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒevent = λ.GetItem(λ.Calm(ϒself, "_download_json", λ.Add(λ.StrLiteral("https://api.stretchinternet.com/trinity/event/tcg/"), ϒvideo_id), ϒvideo_id), λ.IntLiteral(0))
					return λ.DictLiteral(map[string]λ.Object{
						"id":        ϒvideo_id,
						"title":     λ.GetItem(ϒevent, λ.StrLiteral("title")),
						"timestamp": λ.Cal(ϒint_or_none, λ.Calm(ϒevent, "get", λ.StrLiteral("dateCreated")), λ.IntLiteral(1000)),
						"url":       λ.Add(λ.StrLiteral("https://"), λ.GetItem(λ.GetItem(λ.GetItem(ϒevent, λ.StrLiteral("media")), λ.IntLiteral(0)), λ.StrLiteral("url"))),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    StretchInternetIE__VALID_URL,
				"_real_extract": StretchInternetIE__real_extract,
			})
		}())
	})
}
