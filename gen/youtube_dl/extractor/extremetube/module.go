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
 * extremetube/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/extremetube.py
 */

package extremetube

import (
	Ωkeezmovies "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/keezmovies"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtremeTubeIE λ.Object
	KeezMoviesIE  λ.Object
	ϒstr_to_int   λ.Object
)

func init() {
	λ.InitModule(func() {
		ϒstr_to_int = Ωutils.ϒstr_to_int
		KeezMoviesIE = Ωkeezmovies.KeezMoviesIE
		ExtremeTubeIE = λ.Cal(λ.TypeType, λ.StrLiteral("ExtremeTubeIE"), λ.NewTuple(KeezMoviesIE), func() λ.Dict {
			var (
				ExtremeTubeIE__VALID_URL    λ.Object
				ExtremeTubeIE__real_extract λ.Object
			)
			ExtremeTubeIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?extremetube\\.com/(?:[^/]+/)?video/(?P<id>[^/#?&]+)")
			ExtremeTubeIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒinfo       λ.Object
						ϒself       = λargs[0]
						ϒuploader   λ.Object
						ϒurl        = λargs[1]
						ϒview_count λ.Object
						ϒwebpage    λ.Object
						τmp0        λ.Object
					)
					τmp0 = λ.Calm(ϒself, "_extract_info", ϒurl)
					ϒwebpage = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒinfo = λ.GetItem(τmp0, λ.IntLiteral(1))
					if !λ.IsTrue(λ.GetItem(ϒinfo, λ.StrLiteral("title"))) {
						λ.SetItem(ϒinfo, λ.StrLiteral("title"), λ.Calm(ϒself, "_search_regex", λ.StrLiteral("<h1[^>]+title=\"([^\"]+)\"[^>]*>"), ϒwebpage, λ.StrLiteral("title")))
					}
					ϒuploader = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("Uploaded by:\\s*</[^>]+>\\s*<a[^>]+>(.+?)</a>"),
						ϒwebpage,
						λ.StrLiteral("uploader"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒview_count = λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("Views:\\s*</[^>]+>\\s*<[^>]+>([\\d,\\.]+)</"),
						ϒwebpage,
						λ.StrLiteral("view count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					λ.Calm(ϒinfo, "update", λ.DictLiteral(map[string]λ.Object{
						"uploader":   ϒuploader,
						"view_count": ϒview_count,
					}))
					return ϒinfo
				})
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL":    ExtremeTubeIE__VALID_URL,
				"_real_extract": ExtremeTubeIE__real_extract,
			})
		}())
	})
}
