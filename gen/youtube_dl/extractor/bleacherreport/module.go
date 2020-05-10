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
 * bleacherreport/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/bleacherreport.py
 */

package bleacherreport

import (
	Ωamp "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/amp"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	AMPIE               λ.Object
	BleacherReportCMSIE λ.Object
	BleacherReportIE    λ.Object
	ExtractorError      λ.Object
	InfoExtractor       λ.Object
	ϒint_or_none        λ.Object
	ϒparse_iso8601      λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		AMPIE = Ωamp.AMPIE
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		BleacherReportIE = λ.Cal(λ.TypeType, λ.StrLiteral("BleacherReportIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				BleacherReportIE__VALID_URL    λ.Object
				BleacherReportIE__real_extract λ.Object
			)
			BleacherReportIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?bleacherreport\\.com/articles/(?P<id>\\d+)")
			BleacherReportIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒarticle_data  λ.Object
						ϒarticle_id    λ.Object
						ϒinfo          λ.Object
						ϒprimary_photo λ.Object
						ϒself          = λargs[0]
						ϒthumbnails    λ.Object
						ϒurl           = λargs[1]
						ϒvideo         λ.Object
						ϒvideo_type    λ.Object
					)
					ϒarticle_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒarticle_data = λ.GetItem(λ.Calm(ϒself, "_download_json", λ.Mod(λ.StrLiteral("http://api.bleacherreport.com/api/v1/articles/%s"), ϒarticle_id), ϒarticle_id), λ.StrLiteral("article"))
					ϒthumbnails = λ.NewList()
					ϒprimary_photo = λ.Calm(ϒarticle_data, "get", λ.StrLiteral("primaryPhoto"))
					if λ.IsTrue(ϒprimary_photo) {
						ϒthumbnails = λ.NewList(λ.DictLiteral(map[string]λ.Object{
							"url":    λ.GetItem(ϒprimary_photo, λ.StrLiteral("url")),
							"width":  λ.Calm(ϒprimary_photo, "get", λ.StrLiteral("width")),
							"height": λ.Calm(ϒprimary_photo, "get", λ.StrLiteral("height")),
						}))
					}
					ϒinfo = λ.DictLiteral(map[string]λ.Object{
						"_type":         λ.StrLiteral("url_transparent"),
						"id":            ϒarticle_id,
						"title":         λ.GetItem(ϒarticle_data, λ.StrLiteral("title")),
						"uploader":      λ.Calm(λ.Calm(ϒarticle_data, "get", λ.StrLiteral("author"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("name")),
						"uploader_id":   λ.Calm(ϒarticle_data, "get", λ.StrLiteral("authorId")),
						"timestamp":     λ.Cal(ϒparse_iso8601, λ.Calm(ϒarticle_data, "get", λ.StrLiteral("createdAt"))),
						"thumbnails":    ϒthumbnails,
						"comment_count": λ.Cal(ϒint_or_none, λ.Calm(ϒarticle_data, "get", λ.StrLiteral("commentsCount"))),
						"view_count":    λ.Cal(ϒint_or_none, λ.Calm(ϒarticle_data, "get", λ.StrLiteral("hitCount"))),
					})
					ϒvideo = λ.Calm(ϒarticle_data, "get", λ.StrLiteral("video"))
					if λ.IsTrue(ϒvideo) {
						ϒvideo_type = λ.GetItem(ϒvideo, λ.StrLiteral("type"))
						if λ.Contains(λ.NewTuple(
							λ.StrLiteral("cms.bleacherreport.com"),
							λ.StrLiteral("vid.bleacherreport.com"),
						), ϒvideo_type) {
							λ.SetItem(ϒinfo, λ.StrLiteral("url"), λ.Mod(λ.StrLiteral("http://bleacherreport.com/video_embed?id=%s"), λ.GetItem(ϒvideo, λ.StrLiteral("id"))))
						} else {
							if λ.IsTrue(λ.Eq(ϒvideo_type, λ.StrLiteral("ooyala.com"))) {
								λ.SetItem(ϒinfo, λ.StrLiteral("url"), λ.Mod(λ.StrLiteral("ooyala:%s"), λ.GetItem(ϒvideo, λ.StrLiteral("id"))))
							} else {
								if λ.IsTrue(λ.Eq(ϒvideo_type, λ.StrLiteral("youtube.com"))) {
									λ.SetItem(ϒinfo, λ.StrLiteral("url"), λ.GetItem(ϒvideo, λ.StrLiteral("id")))
								} else {
									if λ.IsTrue(λ.Eq(ϒvideo_type, λ.StrLiteral("vine.co"))) {
										λ.SetItem(ϒinfo, λ.StrLiteral("url"), λ.Mod(λ.StrLiteral("https://vine.co/v/%s"), λ.GetItem(ϒvideo, λ.StrLiteral("id"))))
									} else {
										λ.SetItem(ϒinfo, λ.StrLiteral("url"), λ.Add(ϒvideo_type, λ.GetItem(ϒvideo, λ.StrLiteral("id"))))
									}
								}
							}
						}
						return ϒinfo
					} else {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.StrLiteral("no video in the article")), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					return λ.None
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    BleacherReportIE__VALID_URL,
				"_real_extract": BleacherReportIE__real_extract,
			})
		}())
		BleacherReportCMSIE = λ.Cal(λ.TypeType, λ.StrLiteral("BleacherReportCMSIE"), λ.NewTuple(AMPIE), func() λ.Dict {
			var (
				BleacherReportCMSIE__VALID_URL    λ.Object
				BleacherReportCMSIE__real_extract λ.Object
			)
			BleacherReportCMSIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?bleacherreport\\.com/video_embed\\?id=(?P<id>[0-9a-f-]{36}|\\d{5})")
			BleacherReportCMSIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒinfo     λ.Object
						ϒself     = λargs[0]
						ϒurl      = λargs[1]
						ϒvideo_id λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒinfo = λ.Calm(ϒself, "_extract_feed_info", λ.Mod(λ.StrLiteral("http://vid.bleacherreport.com/videos/%s.akamai"), ϒvideo_id))
					λ.SetItem(ϒinfo, λ.StrLiteral("id"), ϒvideo_id)
					return ϒinfo
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    BleacherReportCMSIE__VALID_URL,
				"_real_extract": BleacherReportCMSIE__real_extract,
			})
		}())
	})
}
