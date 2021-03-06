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
 * hitrecord/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/hitrecord.py
 */

package hitrecord

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	HitRecordIE    λ.Object
	InfoExtractor  λ.Object
	ϒclean_html    λ.Object
	ϒcompat_str    λ.Object
	ϒfloat_or_none λ.Object
	ϒint_or_none   λ.Object
	ϒtry_get       λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒclean_html = Ωutils.ϒclean_html
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒtry_get = Ωutils.ϒtry_get
		HitRecordIE = λ.Cal(λ.TypeType, λ.StrLiteral("HitRecordIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				HitRecordIE__VALID_URL    λ.Object
				HitRecordIE__real_extract λ.Object
			)
			HitRecordIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?hitrecord\\.org/records/(?P<id>\\d+)")
			HitRecordIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself      = λargs[0]
						ϒtags      λ.Object
						ϒtags_list λ.Object
						ϒtitle     λ.Object
						ϒurl       = λargs[1]
						ϒvideo     λ.Object
						ϒvideo_id  λ.Object
						ϒvideo_url λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒvideo = λ.Calm(ϒself, "_download_json", λ.Mod(λ.StrLiteral("https://hitrecord.org/api/web/records/%s"), ϒvideo_id), ϒvideo_id)
					ϒtitle = λ.GetItem(ϒvideo, λ.StrLiteral("title"))
					ϒvideo_url = λ.GetItem(λ.GetItem(ϒvideo, λ.StrLiteral("source_url")), λ.StrLiteral("mp4_url"))
					ϒtags = λ.None
					ϒtags_list = λ.Cal(ϒtry_get, ϒvideo, λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "x"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒx = λargs[0]
							)
							return λ.GetItem(ϒx, λ.StrLiteral("tags"))
						}), λ.ListType)
					if λ.IsTrue(ϒtags_list) {
						ϒtags = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
							nil,
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
									var (
										ϒt   λ.Object
										τmp0 λ.Object
										τmp1 λ.Object
									)
									τmp0 = λ.Cal(λ.BuiltinIter, ϒtags_list)
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										ϒt = τmp1
										if λ.IsTrue(func() λ.Object {
											if λv := λ.Cal(λ.BuiltinIsInstance, ϒt, λ.DictType); !λ.IsTrue(λv) {
												return λv
											} else if λv := λ.Calm(ϒt, "get", λ.StrLiteral("text")); !λ.IsTrue(λv) {
												return λv
											} else {
												return λ.Cal(λ.BuiltinIsInstance, λ.GetItem(ϒt, λ.StrLiteral("text")), ϒcompat_str)
											}
										}()) {
											λgy.Yield(λ.GetItem(ϒt, λ.StrLiteral("text")))
										}
									}
									return λ.None
								})
							})))
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id":          ϒvideo_id,
						"url":         ϒvideo_url,
						"title":       ϒtitle,
						"description": λ.Cal(ϒclean_html, λ.Calm(ϒvideo, "get", λ.StrLiteral("body"))),
						"duration":    λ.Cal(ϒfloat_or_none, λ.Calm(ϒvideo, "get", λ.StrLiteral("duration")), λ.IntLiteral(1000)),
						"timestamp":   λ.Cal(ϒint_or_none, λ.Calm(ϒvideo, "get", λ.StrLiteral("created_at_i"))),
						"uploader": λ.Cal(ϒtry_get, ϒvideo, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("user")), λ.StrLiteral("username"))
							}), ϒcompat_str),
						"uploader_id": λ.Cal(ϒtry_get, ϒvideo, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.Cal(ϒcompat_str, λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("user")), λ.StrLiteral("id")))
							})),
						"view_count":    λ.Cal(ϒint_or_none, λ.Calm(ϒvideo, "get", λ.StrLiteral("total_views_count"))),
						"like_count":    λ.Cal(ϒint_or_none, λ.Calm(ϒvideo, "get", λ.StrLiteral("hearts_count"))),
						"comment_count": λ.Cal(ϒint_or_none, λ.Calm(ϒvideo, "get", λ.StrLiteral("comments_count"))),
						"tags":          ϒtags,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    HitRecordIE__VALID_URL,
				"_real_extract": HitRecordIE__real_extract,
			})
		}())
	})
}
