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
 * washingtonpost/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/washingtonpost.py
 */

package washingtonpost

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor           λ.Object
	WashingtonPostArticleIE λ.Object
	WashingtonPostIE        λ.Object
	ϒint_or_none            λ.Object
	ϒstrip_jsonp            λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒstrip_jsonp = Ωutils.ϒstrip_jsonp
		WashingtonPostIE = λ.Cal(λ.TypeType, λ.StrLiteral("WashingtonPostIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				WashingtonPostIE_IE_NAME       λ.Object
				WashingtonPostIE__VALID_URL    λ.Object
				WashingtonPostIE__real_extract λ.Object
			)
			WashingtonPostIE_IE_NAME = λ.StrLiteral("washingtonpost")
			WashingtonPostIE__VALID_URL = λ.StrLiteral("(?:washingtonpost:|https?://(?:www\\.)?washingtonpost\\.com/video/(?:[^/]+/)*)(?P<id>[\\da-f]{8}-[\\da-f]{4}-[\\da-f]{4}-[\\da-f]{4}-[\\da-f]{12})")
			WashingtonPostIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒformats          λ.Object
						ϒhas_width        λ.Object
						ϒm3u8_format      λ.Object
						ϒm3u8_formats     λ.Object
						ϒs                λ.Object
						ϒs_url            λ.Object
						ϒself             = λargs[0]
						ϒsource_media_url λ.Object
						ϒtitle            λ.Object
						ϒurl              = λargs[1]
						ϒurls             λ.Object
						ϒvbr              λ.Object
						ϒvideo_data       λ.Object
						ϒvideo_id         λ.Object
						ϒvideo_type       λ.Object
						ϒwidth            λ.Object
						τmp0              λ.Object
						τmp1              λ.Object
						τmp2              λ.Object
						τmp3              λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒvideo_data = λ.GetItem(λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.Mod(λ.StrLiteral("http://www.washingtonpost.com/posttv/c/videojson/%s?resType=jsonp"), ϒvideo_id),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "transform_source", Value: ϒstrip_jsonp},
					}), λ.IntLiteral(0)), λ.StrLiteral("contentConfig"))
					ϒtitle = λ.GetItem(ϒvideo_data, λ.StrLiteral("title"))
					ϒurls = λ.NewList()
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒvideo_data, "get", λ.StrLiteral("streams"), λ.NewList()))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒs = τmp1
						ϒs_url = λ.Calm(ϒs, "get", λ.StrLiteral("url"))
						if λ.IsTrue(func() λ.Object {
							if λv := λ.NewBool(!λ.IsTrue(ϒs_url)); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewBool(λ.Contains(ϒurls, ϒs_url))
							}
						}()) {
							continue
						}
						λ.Calm(ϒurls, "append", ϒs_url)
						ϒvideo_type = λ.Calm(ϒs, "get", λ.StrLiteral("type"))
						if λ.IsTrue(λ.Eq(ϒvideo_type, λ.StrLiteral("smil"))) {
							continue
						} else {
							if λ.IsTrue(func() λ.Object {
								if λv := λ.NewBool(λ.Contains(λ.NewTuple(
									λ.StrLiteral("ts"),
									λ.StrLiteral("hls"),
								), ϒvideo_type)); !λ.IsTrue(λv) {
									return λv
								} else {
									return func() λ.Object {
										if λv := λ.NewBool(λ.Contains(ϒs_url, λ.StrLiteral("_master.m3u8"))); λ.IsTrue(λv) {
											return λv
										} else {
											return λ.NewBool(λ.Contains(ϒs_url, λ.StrLiteral("_mobile.m3u8")))
										}
									}()
								}
							}()) {
								ϒm3u8_formats = λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
									ϒs_url,
									ϒvideo_id,
									λ.StrLiteral("mp4"),
									λ.StrLiteral("m3u8_native"),
								), λ.KWArgs{
									{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
									{Name: "fatal", Value: λ.False},
								})
								τmp2 = λ.Cal(λ.BuiltinIter, ϒm3u8_formats)
								for {
									if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
										break
									}
									ϒm3u8_format = τmp3
									ϒwidth = λ.Calm(ϒm3u8_format, "get", λ.StrLiteral("width"))
									if !λ.IsTrue(ϒwidth) {
										continue
									}
									ϒvbr = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
										λ.Mod(λ.StrLiteral("%d_%d_(\\d+)"), λ.NewTuple(
											ϒwidth,
											λ.GetItem(ϒm3u8_format, λ.StrLiteral("height")),
										)),
										λ.GetItem(ϒm3u8_format, λ.StrLiteral("url")),
										λ.StrLiteral("vbr"),
									), λ.KWArgs{
										{Name: "default", Value: λ.None},
									})
									if λ.IsTrue(ϒvbr) {
										λ.Calm(ϒm3u8_format, "update", λ.DictLiteral(map[string]λ.Object{
											"vbr": λ.Cal(ϒint_or_none, ϒvbr),
										}))
									}
								}
								λ.Calm(ϒformats, "extend", ϒm3u8_formats)
							} else {
								ϒwidth = λ.Cal(ϒint_or_none, λ.Calm(ϒs, "get", λ.StrLiteral("width")))
								ϒvbr = λ.Cal(ϒint_or_none, λ.Calm(ϒs, "get", λ.StrLiteral("bitrate")))
								ϒhas_width = λ.Ne(ϒwidth, λ.IntLiteral(0))
								λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
									"format_id": func() λ.Object {
										if λ.IsTrue(ϒwidth) {
											return λ.Mod(λ.StrLiteral("%s-%d-%d"), λ.NewTuple(
												ϒvideo_type,
												ϒwidth,
												ϒvbr,
											))
										} else {
											return ϒvideo_type
										}
									}(),
									"vbr": func() λ.Object {
										if λ.IsTrue(ϒhas_width) {
											return ϒvbr
										} else {
											return λ.None
										}
									}(),
									"width":  ϒwidth,
									"height": λ.Cal(ϒint_or_none, λ.Calm(ϒs, "get", λ.StrLiteral("height"))),
									"acodec": λ.Calm(ϒs, "get", λ.StrLiteral("audioCodec")),
									"vcodec": func() λ.Object {
										if λ.IsTrue(ϒhas_width) {
											return λ.Calm(ϒs, "get", λ.StrLiteral("videoCodec"))
										} else {
											return λ.StrLiteral("none")
										}
									}(),
									"filesize": λ.Cal(ϒint_or_none, λ.Calm(ϒs, "get", λ.StrLiteral("fileSize"))),
									"url":      ϒs_url,
									"ext":      λ.StrLiteral("mp4"),
									"protocol": func() λ.Object {
										if λ.Contains(λ.NewTuple(
											λ.StrLiteral("ts"),
											λ.StrLiteral("hls"),
										), ϒvideo_type) {
											return λ.StrLiteral("m3u8_native")
										} else {
											return λ.None
										}
									}(),
								}))
							}
						}
					}
					ϒsource_media_url = λ.Calm(ϒvideo_data, "get", λ.StrLiteral("sourceMediaURL"))
					if λ.IsTrue(ϒsource_media_url) {
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"format_id": λ.StrLiteral("source_media"),
							"url":       ϒsource_media_url,
						}))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats, λ.NewTuple(
						λ.StrLiteral("width"),
						λ.StrLiteral("height"),
						λ.StrLiteral("vbr"),
						λ.StrLiteral("filesize"),
						λ.StrLiteral("tbr"),
						λ.StrLiteral("format_id"),
					))
					return λ.DictLiteral(map[string]λ.Object{
						"id":          ϒvideo_id,
						"title":       ϒtitle,
						"description": λ.Calm(ϒvideo_data, "get", λ.StrLiteral("blurb")),
						"uploader":    λ.Calm(λ.Calm(ϒvideo_data, "get", λ.StrLiteral("credits"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("source")),
						"formats":     ϒformats,
						"duration":    λ.Cal(ϒint_or_none, λ.Calm(ϒvideo_data, "get", λ.StrLiteral("videoDuration")), λ.IntLiteral(100)),
						"timestamp":   λ.Cal(ϒint_or_none, λ.Calm(λ.Calm(ϒvideo_data, "get", λ.StrLiteral("dateConfig"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("dateFirstPublished")), λ.IntLiteral(1000)),
					})
				})
			return λ.DictLiteral(map[string]λ.Object{
				"IE_NAME":       WashingtonPostIE_IE_NAME,
				"_VALID_URL":    WashingtonPostIE__VALID_URL,
				"_real_extract": WashingtonPostIE__real_extract,
			})
		}())
		WashingtonPostArticleIE = λ.Cal(λ.TypeType, λ.StrLiteral("WashingtonPostArticleIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				WashingtonPostArticleIE__VALID_URL λ.Object
				WashingtonPostArticleIE_suitable   λ.Object
			)
			WashingtonPostArticleIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?washingtonpost\\.com/(?:[^/]+/)*(?P<id>[^/?#]+)")
			WashingtonPostArticleIE_suitable = λ.NewFunction("suitable",
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
						if λ.IsTrue(λ.Calm(WashingtonPostIE, "suitable", ϒurl)) {
							return λ.False
						} else {
							return λ.Calm(λ.Cal(λ.SuperType, WashingtonPostArticleIE, ϒcls), "suitable", ϒurl)
						}
					}()
				})
			WashingtonPostArticleIE_suitable = λ.Cal(λ.ClassMethodType, WashingtonPostArticleIE_suitable)
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL": WashingtonPostArticleIE__VALID_URL,
				"suitable":   WashingtonPostArticleIE_suitable,
			})
		}())
	})
}
