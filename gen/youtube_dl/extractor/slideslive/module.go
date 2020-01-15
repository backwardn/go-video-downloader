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
 * slideslive/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/slideslive.py
 */

package slideslive

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	SlidesLiveIE  λ.Object
	ϒsmuggle_url  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒsmuggle_url = Ωutils.ϒsmuggle_url
		SlidesLiveIE = λ.Cal(λ.TypeType, λ.StrLiteral("SlidesLiveIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				SlidesLiveIE__VALID_URL    λ.Object
				SlidesLiveIE__real_extract λ.Object
			)
			SlidesLiveIE__VALID_URL = λ.StrLiteral("https?://slideslive\\.com/(?P<id>[0-9]+)")
			SlidesLiveIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒinfo         λ.Object
						ϒself         = λargs[0]
						ϒservice_id   λ.Object
						ϒservice_name λ.Object
						ϒurl          = λargs[1]
						ϒvideo_data   λ.Object
						ϒvideo_id     λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒvideo_data = λ.Calm(ϒself, "_download_json", λ.Add(λ.StrLiteral("https://ben.slideslive.com/player/"), ϒvideo_id), ϒvideo_id)
					ϒservice_name = λ.Calm(λ.GetItem(ϒvideo_data, λ.StrLiteral("video_service_name")), "lower")
					if !λ.Contains(λ.NewTuple(
						λ.StrLiteral("url"),
						λ.StrLiteral("vimeo"),
						λ.StrLiteral("youtube"),
					), ϒservice_name) {
						panic(λ.Raise(λ.Cal(λ.AssertionErrorType)))
					}
					ϒservice_id = λ.GetItem(ϒvideo_data, λ.StrLiteral("video_service_id"))
					ϒinfo = λ.DictLiteral(map[string]λ.Object{
						"id":        ϒvideo_id,
						"thumbnail": λ.Calm(ϒvideo_data, "get", λ.StrLiteral("thumbnail")),
						"url":       ϒservice_id,
					})
					if λ.IsTrue(λ.Eq(ϒservice_name, λ.StrLiteral("url"))) {
						λ.SetItem(ϒinfo, λ.StrLiteral("title"), λ.GetItem(ϒvideo_data, λ.StrLiteral("title")))
					} else {
						λ.Calm(ϒinfo, "update", λ.DictLiteral(map[string]λ.Object{
							"_type":  λ.StrLiteral("url_transparent"),
							"ie_key": λ.Calm(ϒservice_name, "capitalize"),
							"title":  λ.Calm(ϒvideo_data, "get", λ.StrLiteral("title")),
						}))
						if λ.IsTrue(λ.Eq(ϒservice_name, λ.StrLiteral("vimeo"))) {
							λ.SetItem(ϒinfo, λ.StrLiteral("url"), λ.Cal(ϒsmuggle_url, λ.Add(λ.StrLiteral("https://player.vimeo.com/video/"), ϒservice_id), λ.DictLiteral(map[string]λ.Object{
								"http_headers": λ.DictLiteral(map[string]λ.Object{
									"Referer": ϒurl,
								}),
							})))
						}
					}
					return ϒinfo
				})
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL":    SlidesLiveIE__VALID_URL,
				"_real_extract": SlidesLiveIE__real_extract,
			})
		}())
	})
}
