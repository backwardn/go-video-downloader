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
 * dw/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/dw.py
 */

package dw

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	DWArticleIE      λ.Object
	DWIE             λ.Object
	InfoExtractor    λ.Object
	ϒint_or_none     λ.Object
	ϒunified_strdate λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒunified_strdate = Ωutils.ϒunified_strdate
		DWIE = λ.Cal(λ.TypeType, λ.StrLiteral("DWIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				DWIE_IE_NAME       λ.Object
				DWIE__VALID_URL    λ.Object
				DWIE__real_extract λ.Object
			)
			DWIE_IE_NAME = λ.StrLiteral("dw")
			DWIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?dw\\.com/(?:[^/]+/)+(?:av|e)-(?P<id>\\d+)")
			DWIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒformats       λ.Object
						ϒhidden_inputs λ.Object
						ϒmedia_id      λ.Object
						ϒself          = λargs[0]
						ϒtitle         λ.Object
						ϒupload_date   λ.Object
						ϒurl           = λargs[1]
						ϒwebpage       λ.Object
					)
					ϒmedia_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒmedia_id)
					ϒhidden_inputs = λ.Calm(ϒself, "_hidden_inputs", ϒwebpage)
					ϒtitle = λ.GetItem(ϒhidden_inputs, λ.StrLiteral("media_title"))
					ϒmedia_id = func() λ.Object {
						if λv := λ.Calm(ϒhidden_inputs, "get", λ.StrLiteral("media_id")); λ.IsTrue(λv) {
							return λv
						} else {
							return ϒmedia_id
						}
					}()
					if λ.IsTrue(func() λ.Object {
						if λv := λ.Eq(λ.Calm(ϒhidden_inputs, "get", λ.StrLiteral("player_type")), λ.StrLiteral("video")); !λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Eq(λ.Calm(ϒhidden_inputs, "get", λ.StrLiteral("stream_file")), λ.StrLiteral("1"))
						}
					}()) {
						ϒformats = λ.Call(λ.GetAttr(ϒself, "_extract_smil_formats", nil), λ.NewArgs(
							λ.Mod(λ.StrLiteral("http://www.dw.com/smil/v-%s"), ϒmedia_id),
							ϒmedia_id,
						), λ.KWArgs{
							{Name: "transform_source", Value: λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "s"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒs = λargs[0]
									)
									return λ.Calm(ϒs, "replace", λ.StrLiteral("rtmp://tv-od.dw.de/flash/"), λ.StrLiteral("http://tv-download.dw.de/dwtv_video/flv/"))
								})},
						})
						λ.Calm(ϒself, "_sort_formats", ϒformats)
					} else {
						ϒformats = λ.NewList(λ.DictLiteral(map[string]λ.Object{
							"url": λ.GetItem(ϒhidden_inputs, λ.StrLiteral("file_name")),
						}))
					}
					ϒupload_date = λ.Calm(ϒhidden_inputs, "get", λ.StrLiteral("display_date"))
					if !λ.IsTrue(ϒupload_date) {
						ϒupload_date = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
							λ.StrLiteral("<span[^>]+class=\"date\">([0-9.]+)\\s*\\|"),
							ϒwebpage,
							λ.StrLiteral("upload date"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
						ϒupload_date = λ.Cal(ϒunified_strdate, ϒupload_date)
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id":          ϒmedia_id,
						"title":       ϒtitle,
						"description": λ.Calm(ϒself, "_og_search_description", ϒwebpage),
						"thumbnail":   λ.Calm(ϒhidden_inputs, "get", λ.StrLiteral("preview_image")),
						"duration":    λ.Cal(ϒint_or_none, λ.Calm(ϒhidden_inputs, "get", λ.StrLiteral("file_duration"))),
						"upload_date": ϒupload_date,
						"formats":     ϒformats,
					})
				})
			return λ.DictLiteral(map[string]λ.Object{
				"IE_NAME":       DWIE_IE_NAME,
				"_VALID_URL":    DWIE__VALID_URL,
				"_real_extract": DWIE__real_extract,
			})
		}())
		DWArticleIE = λ.Cal(λ.TypeType, λ.StrLiteral("DWArticleIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				DWArticleIE__VALID_URL λ.Object
			)
			DWArticleIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?dw\\.com/(?:[^/]+/)+a-(?P<id>\\d+)")
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL": DWArticleIE__VALID_URL,
			})
		}())
	})
}
