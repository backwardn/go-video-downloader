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
 * tvanouvelles/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/tvanouvelles.py
 */

package tvanouvelles

import (
	Ωbrightcove "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/brightcove"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	BrightcoveNewIE       λ.Object
	InfoExtractor         λ.Object
	TVANouvellesArticleIE λ.Object
	TVANouvellesIE        λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		BrightcoveNewIE = Ωbrightcove.BrightcoveNewIE
		TVANouvellesIE = λ.Cal(λ.TypeType, λ.StrLiteral("TVANouvellesIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TVANouvellesIE_BRIGHTCOVE_URL_TEMPLATE λ.Object
				TVANouvellesIE__VALID_URL              λ.Object
				TVANouvellesIE__real_extract           λ.Object
			)
			TVANouvellesIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?tvanouvelles\\.ca/videos/(?P<id>\\d+)")
			TVANouvellesIE_BRIGHTCOVE_URL_TEMPLATE = λ.StrLiteral("http://players.brightcove.net/1741764581/default_default/index.html?videoId=%s")
			TVANouvellesIE__real_extract = λ.NewFunction("_real_extract",
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
					return λ.Calm(ϒself, "url_result", λ.Mod(λ.GetAttr(ϒself, "BRIGHTCOVE_URL_TEMPLATE", nil), ϒbrightcove_id), λ.Calm(BrightcoveNewIE, "ie_key"), ϒbrightcove_id)
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"BRIGHTCOVE_URL_TEMPLATE": TVANouvellesIE_BRIGHTCOVE_URL_TEMPLATE,
				"_VALID_URL":              TVANouvellesIE__VALID_URL,
				"_real_extract":           TVANouvellesIE__real_extract,
			})
		}())
		TVANouvellesArticleIE = λ.Cal(λ.TypeType, λ.StrLiteral("TVANouvellesArticleIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TVANouvellesArticleIE__VALID_URL λ.Object
				TVANouvellesArticleIE_suitable   λ.Object
			)
			TVANouvellesArticleIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?tvanouvelles\\.ca/(?:[^/]+/)+(?P<id>[^/?#&]+)")
			TVANouvellesArticleIE_suitable = λ.NewFunction("suitable",
				[]λ.Param{
					{Name: "cls"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcls = λargs[0]
						ϒurl = λargs[1]
					)
					return func() λ.Object {
						if λ.IsTrue(λ.Calm(TVANouvellesIE, "suitable", ϒurl)) {
							return λ.False
						} else {
							return λ.Calm(λ.Cal(λ.SuperType, TVANouvellesArticleIE, ϒcls), "suitable", ϒurl)
						}
					}()
				})
			TVANouvellesArticleIE_suitable = λ.Cal(λ.ClassMethodType, TVANouvellesArticleIE_suitable)
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": TVANouvellesArticleIE__VALID_URL,
				"suitable":   TVANouvellesArticleIE_suitable,
			})
		}())
	})
}
