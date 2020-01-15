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
 * pearvideo/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/pearvideo.py
 */

package pearvideo

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor      λ.Object
	PearVideoIE        λ.Object
	ϒqualities         λ.Object
	ϒunified_timestamp λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒqualities = Ωutils.ϒqualities
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		PearVideoIE = λ.Cal(λ.TypeType, λ.StrLiteral("PearVideoIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				PearVideoIE__VALID_URL    λ.Object
				PearVideoIE__real_extract λ.Object
			)
			PearVideoIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?pearvideo\\.com/video_(?P<id>\\d+)")
			PearVideoIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdescription λ.Object
						ϒformats     λ.Object
						ϒquality     λ.Object
						ϒself        = λargs[0]
						ϒtimestamp   λ.Object
						ϒtitle       λ.Object
						ϒurl         = λargs[1]
						ϒvideo_id    λ.Object
						ϒwebpage     λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒvideo_id)
					ϒquality = λ.Cal(ϒqualities, λ.NewTuple(
						λ.StrLiteral("ldflv"),
						λ.StrLiteral("ld"),
						λ.StrLiteral("sdflv"),
						λ.StrLiteral("sd"),
						λ.StrLiteral("hdflv"),
						λ.StrLiteral("hd"),
						λ.StrLiteral("src"),
					))
					ϒformats = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒmobj λ.Object
									τmp0  λ.Object
									τmp1  λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfinditer, λ.StrLiteral("(?P<id>[a-zA-Z]+)Url\\s*=\\s*([\"\\'])(?P<url>(?:https?:)?//.+?)\\2"), ϒwebpage))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒmobj = τmp1
									λgy.Yield(λ.DictLiteral(map[string]λ.Object{
										"url":       λ.Calm(ϒmobj, "group", λ.StrLiteral("url")),
										"format_id": λ.Calm(ϒmobj, "group", λ.StrLiteral("id")),
										"quality":   λ.Cal(ϒquality, λ.Calm(ϒmobj, "group", λ.StrLiteral("id"))),
									}))
								}
								return λ.None
							})
						})))
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒtitle = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewTuple(
							λ.StrLiteral("<h1[^>]+\\bclass=([\"\\'])video-tt\\1[^>]*>(?P<value>[^<]+)"),
							λ.StrLiteral("<[^>]+\\bdata-title=([\"\\'])(?P<value>(?:(?!\\1).)+)\\1"),
						),
						ϒwebpage,
						λ.StrLiteral("title"),
					), λ.KWArgs{
						{Name: "group", Value: λ.StrLiteral("value")},
					})
					ϒdescription = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewTuple(
								λ.StrLiteral("<div[^>]+\\bclass=([\"\\'])summary\\1[^>]*>(?P<value>[^<]+)"),
								λ.StrLiteral("<[^>]+\\bdata-summary=([\"\\'])(?P<value>(?:(?!\\1).)+)\\1"),
							),
							ϒwebpage,
							λ.StrLiteral("description"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
							{Name: "group", Value: λ.StrLiteral("value")},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒself, "_html_search_meta", λ.StrLiteral("Description"), ϒwebpage)
						}
					}()
					ϒtimestamp = λ.Cal(ϒunified_timestamp, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("<div[^>]+\\bclass=[\"\\']date[\"\\'][^>]*>([^<]+)"),
						ϒwebpage,
						λ.StrLiteral("timestamp"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					return λ.DictLiteral(map[string]λ.Object{
						"id":          ϒvideo_id,
						"title":       ϒtitle,
						"description": ϒdescription,
						"timestamp":   ϒtimestamp,
						"formats":     ϒformats,
					})
				})
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL":    PearVideoIE__VALID_URL,
				"_real_extract": PearVideoIE__real_extract,
			})
		}())
	})
}
