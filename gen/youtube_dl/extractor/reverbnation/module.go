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
 * reverbnation/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/reverbnation.py
 */

package reverbnation

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor  λ.Object
	ReverbNationIE λ.Object
	ϒqualities     λ.Object
	ϒstr_or_none   λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒqualities = Ωutils.ϒqualities
		ϒstr_or_none = Ωutils.ϒstr_or_none
		ReverbNationIE = λ.Cal(λ.TypeType, λ.StrLiteral("ReverbNationIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ReverbNationIE__VALID_URL    λ.Object
				ReverbNationIE__real_extract λ.Object
			)
			ReverbNationIE__VALID_URL = λ.StrLiteral("^https?://(?:www\\.)?reverbnation\\.com/.*?/song/(?P<id>\\d+).*?$")
			ReverbNationIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						THUMBNAILS  λ.Object
						ϒapi_res    λ.Object
						ϒquality    λ.Object
						ϒself       = λargs[0]
						ϒsong_id    λ.Object
						ϒthumb_key  λ.Object
						ϒthumbnails λ.Object
						ϒurl        = λargs[1]
						τmp0        λ.Object
						τmp1        λ.Object
					)
					ϒsong_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒapi_res = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.Mod(λ.StrLiteral("https://api.reverbnation.com/song/%s"), ϒsong_id),
						ϒsong_id,
					), λ.KWArgs{
						{Name: "note", Value: λ.Mod(λ.StrLiteral("Downloading information of song %s"), ϒsong_id)},
					})
					THUMBNAILS = λ.NewTuple(
						λ.StrLiteral("thumbnail"),
						λ.StrLiteral("image"),
					)
					ϒquality = λ.Cal(ϒqualities, THUMBNAILS)
					ϒthumbnails = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, THUMBNAILS)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒthumb_key = τmp1
						if λ.IsTrue(λ.Calm(ϒapi_res, "get", ϒthumb_key)) {
							λ.Calm(ϒthumbnails, "append", λ.DictLiteral(map[string]λ.Object{
								"url":        λ.GetItem(ϒapi_res, ϒthumb_key),
								"preference": λ.Cal(ϒquality, ϒthumb_key),
							}))
						}
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id":          ϒsong_id,
						"title":       λ.GetItem(ϒapi_res, λ.StrLiteral("name")),
						"url":         λ.GetItem(ϒapi_res, λ.StrLiteral("url")),
						"uploader":    λ.Calm(λ.Calm(ϒapi_res, "get", λ.StrLiteral("artist"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("name")),
						"uploader_id": λ.Cal(ϒstr_or_none, λ.Calm(λ.Calm(ϒapi_res, "get", λ.StrLiteral("artist"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("id"))),
						"thumbnails":  ϒthumbnails,
						"ext":         λ.StrLiteral("mp3"),
						"vcodec":      λ.StrLiteral("none"),
					})
				})
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL":    ReverbNationIE__VALID_URL,
				"_real_extract": ReverbNationIE__real_extract,
			})
		}())
	})
}
