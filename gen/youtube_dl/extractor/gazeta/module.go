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
 * gazeta/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/gazeta.py
 */

package gazeta

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	GazetaIE      λ.Object
	InfoExtractor λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		GazetaIE = λ.Cal(λ.TypeType, λ.StrLiteral("GazetaIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				GazetaIE__VALID_URL    λ.Object
				GazetaIE__real_extract λ.Object
			)
			GazetaIE__VALID_URL = λ.StrLiteral("(?P<url>https?://(?:www\\.)?gazeta\\.ru/(?:[^/]+/)?video/(?:main/)*(?:\\d{4}/\\d{2}/\\d{2}/)?(?P<id>[A-Za-z0-9-_.]+)\\.s?html)")
			GazetaIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id λ.Object
						ϒembed_page λ.Object
						ϒembed_url  λ.Object
						ϒmobj       λ.Object
						ϒself       = λargs[0]
						ϒurl        = λargs[1]
						ϒvideo_id   λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒdisplay_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("id"))
					ϒembed_url = λ.Mod(λ.StrLiteral("%s?p=embed"), λ.Calm(ϒmobj, "group", λ.StrLiteral("url")))
					ϒembed_page = λ.Calm(ϒself, "_download_webpage", ϒembed_url, ϒdisplay_id, λ.StrLiteral("Downloading embed page"))
					ϒvideo_id = λ.Calm(ϒself, "_search_regex", λ.StrLiteral("<div[^>]*?class=\"eagleplayer\"[^>]*?data-id=\"([^\"]+)\""), ϒembed_page, λ.StrLiteral("video id"))
					return λ.Calm(ϒself, "url_result", λ.Mod(λ.StrLiteral("eagleplatform:gazeta.media.eagleplatform.com:%s"), ϒvideo_id), λ.StrLiteral("EaglePlatform"))
				})
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL":    GazetaIE__VALID_URL,
				"_real_extract": GazetaIE__real_extract,
			})
		}())
	})
}
