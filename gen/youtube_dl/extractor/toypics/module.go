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
 * toypics/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/toypics.py
 */

package toypics

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	ToypicsIE     λ.Object
	ToypicsUserIE λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ToypicsIE = λ.Cal(λ.TypeType, λ.StrLiteral("ToypicsIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ToypicsIE__VALID_URL    λ.Object
				ToypicsIE__real_extract λ.Object
			)
			ToypicsIE__VALID_URL = λ.StrLiteral("https?://videos\\.toypics\\.net/view/(?P<id>[0-9]+)")
			ToypicsIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒformats  λ.Object
						ϒself     = λargs[0]
						ϒtitle    λ.Object
						ϒuploader λ.Object
						ϒurl      = λargs[1]
						ϒvideo_id λ.Object
						ϒwebpage  λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒvideo_id)
					ϒformats = λ.GetItem(λ.GetItem(λ.Calm(ϒself, "_parse_html5_media_entries", ϒurl, ϒwebpage, ϒvideo_id), λ.IntLiteral(0)), λ.StrLiteral("formats"))
					ϒtitle = λ.Calm(ϒself, "_html_search_regex", λ.NewList(
						λ.StrLiteral("<h1[^>]+class=[\"\\']view-video-title[^>]+>([^<]+)</h"),
						λ.StrLiteral("<title>([^<]+) - Toypics</title>"),
					), ϒwebpage, λ.StrLiteral("title"))
					ϒuploader = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("More videos from <strong>([^<]+)</strong>"),
						ϒwebpage,
						λ.StrLiteral("uploader"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					return λ.DictLiteral(map[string]λ.Object{
						"id":        ϒvideo_id,
						"formats":   ϒformats,
						"title":     ϒtitle,
						"uploader":  ϒuploader,
						"age_limit": λ.IntLiteral(18),
					})
				})
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL":    ToypicsIE__VALID_URL,
				"_real_extract": ToypicsIE__real_extract,
			})
		}())
		ToypicsUserIE = λ.Cal(λ.TypeType, λ.StrLiteral("ToypicsUserIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ToypicsUserIE__VALID_URL λ.Object
			)
			ToypicsUserIE__VALID_URL = λ.StrLiteral("https?://videos\\.toypics\\.net/(?!view)(?P<id>[^/?#&]+)")
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL": ToypicsUserIE__VALID_URL,
			})
		}())
	})
}
