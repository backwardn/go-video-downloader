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
 * acast/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/acast.py
 */

package acast

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ACastChannelIE     λ.Object
	ACastIE            λ.Object
	InfoExtractor      λ.Object
	ϒclean_html        λ.Object
	ϒcompat_str        λ.Object
	ϒfloat_or_none     λ.Object
	ϒint_or_none       λ.Object
	ϒtry_get           λ.Object
	ϒunified_timestamp λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒclean_html = Ωutils.ϒclean_html
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒtry_get = Ωutils.ϒtry_get
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ACastIE = λ.Cal(λ.TypeType, λ.StrLiteral("ACastIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ACastIE_IE_NAME       λ.Object
				ACastIE__VALID_URL    λ.Object
				ACastIE__real_extract λ.Object
			)
			ACastIE_IE_NAME = λ.StrLiteral("acast")
			ACastIE__VALID_URL = λ.StrLiteral("(?x)\n                    https?://\n                        (?:\n                            (?:(?:embed|www)\\.)?acast\\.com/|\n                            play\\.acast\\.com/s/\n                        )\n                        (?P<channel>[^/]+)/(?P<id>[^/#?]+)\n                    ")
			ACastIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcast_data   λ.Object
						ϒchannel     λ.Object
						ϒdisplay_id  λ.Object
						ϒe           λ.Object
						ϒepisode_url λ.Object
						ϒmedia_url   λ.Object
						ϒs           λ.Object
						ϒself        = λargs[0]
						ϒtitle       λ.Object
						ϒurl         = λargs[1]
						τmp0         λ.Object
					)
					τmp0 = λ.Calm(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups")
					ϒchannel = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒdisplay_id = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒs = λ.Calm(ϒself, "_download_json", λ.Mod(λ.StrLiteral("https://feeder.acast.com/api/v1/shows/%s/episodes/%s"), λ.NewTuple(
						ϒchannel,
						ϒdisplay_id,
					)), ϒdisplay_id)
					ϒmedia_url = λ.GetItem(ϒs, λ.StrLiteral("url"))
					if λ.IsTrue(λ.Cal(Ωre.ϒsearch, λ.StrLiteral("[0-9a-f]{8}-(?:[0-9a-f]{4}-){3}[0-9a-f]{12}"), ϒdisplay_id)) {
						ϒepisode_url = λ.Calm(ϒs, "get", λ.StrLiteral("episodeUrl"))
						if λ.IsTrue(ϒepisode_url) {
							ϒdisplay_id = ϒepisode_url
						} else {
							τmp0 = λ.Calm(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), λ.GetItem(ϒs, λ.StrLiteral("link"))), "groups")
							ϒchannel = λ.GetItem(τmp0, λ.IntLiteral(0))
							ϒdisplay_id = λ.GetItem(τmp0, λ.IntLiteral(1))
						}
					}
					ϒcast_data = λ.GetItem(λ.Calm(ϒself, "_download_json", λ.Mod(λ.StrLiteral("https://play-api.acast.com/splash/%s/%s"), λ.NewTuple(
						ϒchannel,
						ϒdisplay_id,
					)), ϒdisplay_id), λ.StrLiteral("result"))
					ϒe = λ.GetItem(ϒcast_data, λ.StrLiteral("episode"))
					ϒtitle = func() λ.Object {
						if λv := λ.Calm(ϒe, "get", λ.StrLiteral("name")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.GetItem(ϒs, λ.StrLiteral("title"))
						}
					}()
					return λ.DictLiteral(map[string]λ.Object{
						"id":         λ.Cal(ϒcompat_str, λ.GetItem(ϒe, λ.StrLiteral("id"))),
						"display_id": ϒdisplay_id,
						"url":        ϒmedia_url,
						"title":      ϒtitle,
						"description": func() λ.Object {
							if λv := λ.Calm(ϒe, "get", λ.StrLiteral("summary")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(ϒclean_html, func() λ.Object {
									if λv := λ.Calm(ϒe, "get", λ.StrLiteral("description")); λ.IsTrue(λv) {
										return λv
									} else {
										return λ.Calm(ϒs, "get", λ.StrLiteral("description"))
									}
								}())
							}
						}(),
						"thumbnail": λ.Calm(ϒe, "get", λ.StrLiteral("image")),
						"timestamp": λ.Cal(ϒunified_timestamp, func() λ.Object {
							if λv := λ.Calm(ϒe, "get", λ.StrLiteral("publishingDate")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(ϒs, "get", λ.StrLiteral("publishDate"))
							}
						}()),
						"duration": λ.Cal(ϒfloat_or_none, func() λ.Object {
							if λv := λ.Calm(ϒe, "get", λ.StrLiteral("duration")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(ϒs, "get", λ.StrLiteral("duration"))
							}
						}()),
						"filesize": λ.Cal(ϒint_or_none, λ.Calm(ϒe, "get", λ.StrLiteral("contentLength"))),
						"creator": λ.Cal(ϒtry_get, ϒcast_data, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("show")), λ.StrLiteral("author"))
							}), ϒcompat_str),
						"series": λ.Cal(ϒtry_get, ϒcast_data, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("show")), λ.StrLiteral("name"))
							}), ϒcompat_str),
						"season_number":  λ.Cal(ϒint_or_none, λ.Calm(ϒe, "get", λ.StrLiteral("seasonNumber"))),
						"episode":        ϒtitle,
						"episode_number": λ.Cal(ϒint_or_none, λ.Calm(ϒe, "get", λ.StrLiteral("episodeNumber"))),
					})
				})
			return λ.DictLiteral(map[string]λ.Object{
				"IE_NAME":       ACastIE_IE_NAME,
				"_VALID_URL":    ACastIE__VALID_URL,
				"_real_extract": ACastIE__real_extract,
			})
		}())
		ACastChannelIE = λ.Cal(λ.TypeType, λ.StrLiteral("ACastChannelIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ACastChannelIE__VALID_URL λ.Object
				ACastChannelIE_suitable   λ.Object
			)
			ACastChannelIE__VALID_URL = λ.StrLiteral("(?x)\n                    https?://\n                        (?:\n                            (?:www\\.)?acast\\.com/|\n                            play\\.acast\\.com/s/\n                        )\n                        (?P<id>[^/#?]+)\n                    ")
			ACastChannelIE_suitable = λ.NewFunction("suitable",
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
						if λ.IsTrue(λ.Calm(ACastIE, "suitable", ϒurl)) {
							return λ.False
						} else {
							return λ.Calm(λ.Cal(λ.SuperType, ACastChannelIE, ϒcls), "suitable", ϒurl)
						}
					}()
				})
			ACastChannelIE_suitable = λ.Cal(λ.ClassMethodType, ACastChannelIE_suitable)
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL": ACastChannelIE__VALID_URL,
				"suitable":   ACastChannelIE_suitable,
			})
		}())
	})
}
