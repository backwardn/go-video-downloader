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
 * tvn24/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/tvn24.py
 */

package tvn24

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	NO_DEFAULT    λ.Object
	TVN24IE       λ.Object
	ϒint_or_none  λ.Object
	ϒunescapeHTML λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒint_or_none = Ωutils.ϒint_or_none
		NO_DEFAULT = Ωutils.NO_DEFAULT
		ϒunescapeHTML = Ωutils.ϒunescapeHTML
		TVN24IE = λ.Cal(λ.TypeType, λ.StrLiteral("TVN24IE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TVN24IE__VALID_URL    λ.Object
				TVN24IE__real_extract λ.Object
			)
			TVN24IE__VALID_URL = λ.StrLiteral("https?://(?:(?:[^/]+)\\.)?tvn24(?:bis)?\\.pl/(?:[^/]+/)*(?P<id>[^/]+)")
			TVN24IE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdescription  λ.Object
						ϒdisplay_id   λ.Object
						ϒextract_json λ.Object
						ϒformat_id    λ.Object
						ϒformats      λ.Object
						ϒquality_data λ.Object
						ϒself         = λargs[0]
						ϒshare_params λ.Object
						ϒthumbnail    λ.Object
						ϒtitle        λ.Object
						ϒurl          = λargs[1]
						ϒvideo_id     λ.Object
						ϒwebpage      λ.Object
						τmp0          λ.Object
						τmp1          λ.Object
						τmp2          λ.Object
					)
					ϒdisplay_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒdisplay_id)
					ϒtitle = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_og_search_title", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒself, "_search_regex", λ.StrLiteral("<h\\d+[^>]+class=[\"\\']magazineItemHeader[^>]+>(.+?)</h"), ϒwebpage, λ.StrLiteral("title"))
						}
					}()
					ϒextract_json = λ.NewFunction("extract_json",
						[]λ.Param{
							{Name: "attr"},
							{Name: "name"},
							{Name: "default", Def: NO_DEFAULT},
							{Name: "fatal", Def: λ.True},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒattr    = λargs[0]
								ϒdefault = λargs[2]
								ϒfatal   = λargs[3]
								ϒname    = λargs[1]
							)
							return λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
								func() λ.Object {
									if λv := λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
										λ.Mod(λ.StrLiteral("\\b%s=([\"\\'])(?P<json>(?!\\1).+?)\\1"), ϒattr),
										ϒwebpage,
										ϒname,
									), λ.KWArgs{
										{Name: "group", Value: λ.StrLiteral("json")},
										{Name: "default", Value: ϒdefault},
										{Name: "fatal", Value: ϒfatal},
									}); λ.IsTrue(λv) {
										return λv
									} else {
										return λ.StrLiteral("{}")
									}
								}(),
								ϒdisplay_id,
							), λ.KWArgs{
								{Name: "transform_source", Value: ϒunescapeHTML},
								{Name: "fatal", Value: ϒfatal},
							})
						})
					ϒquality_data = λ.Cal(ϒextract_json, λ.StrLiteral("data-quality"), λ.StrLiteral("formats"))
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒquality_data, "items"))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒformat_id = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒurl = λ.GetItem(τmp2, λ.IntLiteral(1))
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"url":       ϒurl,
							"format_id": ϒformat_id,
							"height":    λ.Cal(ϒint_or_none, λ.Calm(ϒformat_id, "rstrip", λ.StrLiteral("p"))),
						}))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒdescription = λ.Call(λ.GetAttr(ϒself, "_og_search_description", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					ϒthumbnail = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("\\bdata-poster=([\"\\'])(?P<url>(?!\\1).+?)\\1"),
								ϒwebpage,
								λ.StrLiteral("thumbnail"),
							), λ.KWArgs{
								{Name: "group", Value: λ.StrLiteral("url")},
							})
						}
					}()
					ϒvideo_id = λ.None
					ϒshare_params = λ.Call(ϒextract_json, λ.NewArgs(
						λ.StrLiteral("data-share-params"),
						λ.StrLiteral("share params"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒshare_params, λ.DictType)) {
						ϒvideo_id = λ.Calm(ϒshare_params, "get", λ.StrLiteral("id"))
					}
					if !λ.IsTrue(ϒvideo_id) {
						ϒvideo_id = func() λ.Object {
							if λv := λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("data-vid-id=[\"\\'](\\d+)"),
								ϒwebpage,
								λ.StrLiteral("video id"),
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							}); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
									λ.StrLiteral(",(\\d+)\\.html"),
									ϒurl,
									λ.StrLiteral("video id"),
								), λ.KWArgs{
									{Name: "default", Value: ϒdisplay_id},
								})
							}
						}()
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id":          ϒvideo_id,
						"title":       ϒtitle,
						"description": ϒdescription,
						"thumbnail":   ϒthumbnail,
						"formats":     ϒformats,
					})
				})
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL":    TVN24IE__VALID_URL,
				"_real_extract": TVN24IE__real_extract,
			})
		}())
	})
}
