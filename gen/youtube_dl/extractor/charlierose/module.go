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
 * charlierose/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/charlierose.py
 */

package charlierose

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	CharlieRoseIE λ.Object
	InfoExtractor λ.Object
	ϒremove_end   λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒremove_end = Ωutils.ϒremove_end
		CharlieRoseIE = λ.Cal(λ.TypeType, λ.StrLiteral("CharlieRoseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				CharlieRoseIE__PLAYER_BASE  λ.Object
				CharlieRoseIE__VALID_URL    λ.Object
				CharlieRoseIE__real_extract λ.Object
			)
			CharlieRoseIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?charlierose\\.com/(?:video|episode)(?:s|/player)/(?P<id>\\d+)")
			CharlieRoseIE__PLAYER_BASE = λ.StrLiteral("https://charlierose.com/video/player/%s")
			CharlieRoseIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒinfo_dict λ.Object
						ϒself      = λargs[0]
						ϒtitle     λ.Object
						ϒurl       = λargs[1]
						ϒvideo_id  λ.Object
						ϒwebpage   λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", λ.Mod(λ.GetAttr(ϒself, "_PLAYER_BASE", nil), ϒvideo_id), ϒvideo_id)
					ϒtitle = λ.Cal(ϒremove_end, λ.Calm(ϒself, "_og_search_title", ϒwebpage), λ.StrLiteral(" - Charlie Rose"))
					ϒinfo_dict = λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_parse_html5_media_entries", nil), λ.NewArgs(
						λ.Mod(λ.GetAttr(ϒself, "_PLAYER_BASE", nil), ϒvideo_id),
						ϒwebpage,
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "m3u8_entry_protocol", Value: λ.StrLiteral("m3u8_native")},
					}), λ.IntLiteral(0))
					λ.Calm(ϒself, "_sort_formats", λ.GetItem(ϒinfo_dict, λ.StrLiteral("formats")))
					λ.Calm(ϒself, "_remove_duplicate_formats", λ.GetItem(ϒinfo_dict, λ.StrLiteral("formats")))
					λ.Calm(ϒinfo_dict, "update", λ.DictLiteral(map[string]λ.Object{
						"id":          ϒvideo_id,
						"title":       ϒtitle,
						"thumbnail":   λ.Calm(ϒself, "_og_search_thumbnail", ϒwebpage),
						"description": λ.Calm(ϒself, "_og_search_description", ϒwebpage),
					}))
					return ϒinfo_dict
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_PLAYER_BASE":  CharlieRoseIE__PLAYER_BASE,
				"_VALID_URL":    CharlieRoseIE__VALID_URL,
				"_real_extract": CharlieRoseIE__real_extract,
			})
		}())
	})
}
