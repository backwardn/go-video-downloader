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
 * sonyliv/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/sonyliv.py
 */

package sonyliv

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	SonyLIVIE     λ.Object
	ϒsmuggle_url  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒsmuggle_url = Ωutils.ϒsmuggle_url
		SonyLIVIE = λ.Cal(λ.TypeType, λ.StrLiteral("SonyLIVIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				SonyLIVIE_BRIGHTCOVE_URL_TEMPLATE λ.Object
				SonyLIVIE__VALID_URL              λ.Object
				SonyLIVIE__real_extract           λ.Object
			)
			SonyLIVIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?sonyliv\\.com/details/[^/]+/(?P<id>\\d+)")
			SonyLIVIE_BRIGHTCOVE_URL_TEMPLATE = λ.StrLiteral("http://players.brightcove.net/5182475815001/default_default/index.html?videoId=ref:%s")
			SonyLIVIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒbrightcove_id λ.Object
						ϒself          = λargs[0]
						ϒurl           = λargs[1]
					)
					ϒbrightcove_id = λ.Calm(ϒself, "_match_id", ϒurl)
					return λ.Calm(ϒself, "url_result", λ.Cal(ϒsmuggle_url, λ.Mod(λ.GetAttr(ϒself, "BRIGHTCOVE_URL_TEMPLATE", nil), ϒbrightcove_id), λ.DictLiteral(map[string]λ.Object{
						"geo_countries": λ.NewList(λ.StrLiteral("IN")),
						"referrer":      ϒurl,
					})), λ.StrLiteral("BrightcoveNew"), ϒbrightcove_id)
				})
			return λ.DictLiteral(map[string]λ.Object{
				"BRIGHTCOVE_URL_TEMPLATE": SonyLIVIE_BRIGHTCOVE_URL_TEMPLATE,
				"_VALID_URL":              SonyLIVIE__VALID_URL,
				"_real_extract":           SonyLIVIE__real_extract,
			})
		}())
	})
}
