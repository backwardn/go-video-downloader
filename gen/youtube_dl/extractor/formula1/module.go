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
 * formula1/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/formula1.py
 */

package formula1

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	Formula1IE    λ.Object
	InfoExtractor λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		Formula1IE = λ.Cal(λ.TypeType, λ.StrLiteral("Formula1IE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				Formula1IE__VALID_URL    λ.Object
				Formula1IE__real_extract λ.Object
			)
			Formula1IE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?formula1\\.com/(?:content/fom-website/)?en/video/\\d{4}/\\d{1,2}/(?P<id>.+?)\\.html")
			Formula1IE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id        λ.Object
						ϒooyala_embed_code λ.Object
						ϒself              = λargs[0]
						ϒurl               = λargs[1]
						ϒwebpage           λ.Object
					)
					ϒdisplay_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒdisplay_id)
					ϒooyala_embed_code = λ.Calm(ϒself, "_search_regex", λ.StrLiteral("data-videoid=\"([^\"]+)\""), ϒwebpage, λ.StrLiteral("ooyala embed code"))
					return λ.Calm(ϒself, "url_result", λ.Mod(λ.StrLiteral("ooyala:%s"), ϒooyala_embed_code), λ.StrLiteral("Ooyala"), ϒooyala_embed_code)
				})
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL":    Formula1IE__VALID_URL,
				"_real_extract": Formula1IE__real_extract,
			})
		}())
	})
}
