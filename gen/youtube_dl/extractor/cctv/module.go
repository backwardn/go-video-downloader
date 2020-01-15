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
 * cctv/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/cctv.py
 */

package cctv

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	CCTVIE             λ.Object
	InfoExtractor      λ.Object
	ϒcompat_str        λ.Object
	ϒfloat_or_none     λ.Object
	ϒtry_get           λ.Object
	ϒunified_timestamp λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒtry_get = Ωutils.ϒtry_get
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		CCTVIE = λ.Cal(λ.TypeType, λ.StrLiteral("CCTVIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				CCTVIE__VALID_URL    λ.Object
				CCTVIE__real_extract λ.Object
			)
			CCTVIE__VALID_URL = λ.StrLiteral("https?://(?:(?:[^/]+)\\.(?:cntv|cctv)\\.(?:com|cn)|(?:www\\.)?ncpa-classic\\.com)/(?:[^/]+/)*?(?P<id>[^/?#&]+?)(?:/index)?(?:\\.s?html|[?#&]|$)")
			CCTVIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒchapters_key λ.Object
						ϒdata         λ.Object
						ϒdescription  λ.Object
						ϒduration     λ.Object
						ϒformats      λ.Object
						ϒhls_url      λ.Object
						ϒquality      λ.Object
						ϒself         = λargs[0]
						ϒtimestamp    λ.Object
						ϒtitle        λ.Object
						ϒuploader     λ.Object
						ϒurl          = λargs[1]
						ϒvideo        λ.Object
						ϒvideo_id     λ.Object
						ϒvideo_url    λ.Object
						ϒwebpage      λ.Object
						τmp0          λ.Object
						τmp1          λ.Object
						τmp2          λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒvideo_id)
					ϒvideo_id = λ.Calm(ϒself, "_search_regex", λ.NewList(
						λ.StrLiteral("var\\s+guid\\s*=\\s*[\"\\']([\\da-fA-F]+)"),
						λ.StrLiteral("videoCenterId[\"\\']\\s*,\\s*[\"\\']([\\da-fA-F]+)"),
						λ.StrLiteral("changePlayer\\s*\\(\\s*[\"\\']([\\da-fA-F]+)"),
						λ.StrLiteral("load[Vv]ideo\\s*\\(\\s*[\"\\']([\\da-fA-F]+)"),
						λ.StrLiteral("var\\s+initMyAray\\s*=\\s*[\"\\']([\\da-fA-F]+)"),
						λ.StrLiteral("var\\s+ids\\s*=\\s*\\[[\"\\']([\\da-fA-F]+)"),
					), ϒwebpage, λ.StrLiteral("video id"))
					ϒdata = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.StrLiteral("http://vdn.apps.cntv.cn/api/getHttpVideoInfo.do"),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "query", Value: λ.DictLiteral(map[string]λ.Object{
							"pid":      ϒvideo_id,
							"url":      ϒurl,
							"idl":      λ.IntLiteral(32),
							"idlr":     λ.IntLiteral(32),
							"modifyed": λ.StrLiteral("false"),
						})},
					})
					ϒtitle = λ.GetItem(ϒdata, λ.StrLiteral("title"))
					ϒformats = λ.NewList()
					ϒvideo = λ.Calm(ϒdata, "get", λ.StrLiteral("video"))
					if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒvideo, λ.DictType)) {
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.EnumerateIteratorType, λ.NewTuple(
							λ.StrLiteral("lowChapters"),
							λ.StrLiteral("chapters"),
						)))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							τmp2 = τmp1
							ϒquality = λ.GetItem(τmp2, λ.IntLiteral(0))
							ϒchapters_key = λ.GetItem(τmp2, λ.IntLiteral(1))
							ϒvideo_url = λ.Cal(ϒtry_get, ϒvideo, λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(λ.GetItem(λ.GetItem(ϒx, ϒchapters_key), λ.IntLiteral(0)), λ.StrLiteral("url"))
								}), ϒcompat_str)
							if λ.IsTrue(ϒvideo_url) {
								λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
									"url":        ϒvideo_url,
									"format_id":  λ.StrLiteral("http"),
									"quality":    ϒquality,
									"preference": λ.Neg(λ.IntLiteral(1)),
								}))
							}
						}
					}
					ϒhls_url = λ.Cal(ϒtry_get, ϒdata, λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "x"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒx = λargs[0]
							)
							return λ.GetItem(ϒx, λ.StrLiteral("hls_url"))
						}), ϒcompat_str)
					if λ.IsTrue(ϒhls_url) {
						ϒhls_url = λ.Cal(Ωre.ϒsub, λ.StrLiteral("maxbr=\\d+&?"), λ.StrLiteral(""), ϒhls_url)
						λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
							ϒhls_url,
							ϒvideo_id,
							λ.StrLiteral("mp4"),
						), λ.KWArgs{
							{Name: "entry_protocol", Value: λ.StrLiteral("m3u8_native")},
							{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
							{Name: "fatal", Value: λ.False},
						}))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒuploader = λ.Calm(ϒdata, "get", λ.StrLiteral("editer_name"))
					ϒdescription = λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
						λ.StrLiteral("description"),
						ϒwebpage,
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					ϒtimestamp = λ.Cal(ϒunified_timestamp, λ.Calm(ϒdata, "get", λ.StrLiteral("f_pgmtime")))
					ϒduration = λ.Cal(ϒfloat_or_none, λ.Cal(ϒtry_get, ϒvideo, λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "x"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒx = λargs[0]
							)
							return λ.GetItem(ϒx, λ.StrLiteral("totalLength"))
						})))
					return λ.DictLiteral(map[string]λ.Object{
						"id":          ϒvideo_id,
						"title":       ϒtitle,
						"description": ϒdescription,
						"uploader":    ϒuploader,
						"timestamp":   ϒtimestamp,
						"duration":    ϒduration,
						"formats":     ϒformats,
					})
				})
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL":    CCTVIE__VALID_URL,
				"_real_extract": CCTVIE__real_extract,
			})
		}())
	})
}
