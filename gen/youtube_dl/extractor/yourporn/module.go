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
 * yourporn/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/yourporn.py
 */

package yourporn

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor   λ.Object
	YourPornIE      λ.Object
	ϒcompat_str     λ.Object
	ϒparse_duration λ.Object
	ϒurljoin        λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒurljoin = Ωutils.ϒurljoin
		YourPornIE = λ.Cal(λ.TypeType, λ.StrLiteral("YourPornIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				YourPornIE__VALID_URL    λ.Object
				YourPornIE__real_extract λ.Object
			)
			YourPornIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?sxyprn\\.com/post/(?P<id>[^/?#&.]+)")
			YourPornIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒc         λ.Object
						ϒduration  λ.Object
						ϒnum       λ.Object
						ϒparts     λ.Object
						ϒself      = λargs[0]
						ϒthumbnail λ.Object
						ϒtitle     λ.Object
						ϒurl       = λargs[1]
						ϒvideo_id  λ.Object
						ϒvideo_url λ.Object
						ϒwebpage   λ.Object
						τmp0       λ.Object
						τmp1       λ.Object
						τmp2       λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒvideo_id)
					ϒparts = λ.Calm(λ.GetItem(λ.Calm(ϒself, "_parse_json", λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("data-vnfo=([\"\\'])(?P<data>{.+?})\\1"),
						ϒwebpage,
						λ.StrLiteral("data info"),
					), λ.KWArgs{
						{Name: "group", Value: λ.StrLiteral("data")},
					}), ϒvideo_id), ϒvideo_id), "split", λ.StrLiteral("/"))
					ϒnum = λ.IntLiteral(0)
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Add(λ.GetItem(ϒparts, λ.IntLiteral(6)), λ.GetItem(ϒparts, λ.IntLiteral(7))))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒc = τmp1
						if λ.IsTrue(λ.Calm(ϒc, "isnumeric")) {
							τmp2 = λ.IAdd(ϒnum, λ.Cal(λ.IntType, ϒc))
							ϒnum = τmp2
						}
					}
					λ.SetItem(ϒparts, λ.IntLiteral(5), λ.Cal(ϒcompat_str, λ.Sub(λ.Cal(λ.IntType, λ.GetItem(ϒparts, λ.IntLiteral(5))), ϒnum)))
					τmp0 = λ.IAdd(λ.GetItem(ϒparts, λ.IntLiteral(1)), λ.StrLiteral("8"))
					λ.SetItem(ϒparts, λ.IntLiteral(1), τmp0)
					ϒvideo_url = λ.Cal(ϒurljoin, ϒurl, λ.Calm(λ.StrLiteral("/"), "join", ϒparts))
					ϒtitle = λ.Calm(func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.StrLiteral("<[^>]+\\bclass=[\"\\']PostEditTA[^>]+>([^<]+)"),
							ϒwebpage,
							λ.StrLiteral("title"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒself, "_og_search_description", ϒwebpage)
						}
					}(), "strip")
					ϒthumbnail = λ.Calm(ϒself, "_og_search_thumbnail", ϒwebpage)
					ϒduration = λ.Cal(ϒparse_duration, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("duration\\s*:\\s*<[^>]+>([\\d:]+)"),
						ϒwebpage,
						λ.StrLiteral("duration"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					}))
					return λ.DictLiteral(map[string]λ.Object{
						"id":        ϒvideo_id,
						"url":       ϒvideo_url,
						"title":     ϒtitle,
						"thumbnail": ϒthumbnail,
						"duration":  ϒduration,
						"age_limit": λ.IntLiteral(18),
						"ext":       λ.StrLiteral("mp4"),
					})
				})
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL":    YourPornIE__VALID_URL,
				"_real_extract": YourPornIE__real_extract,
			})
		}())
	})
}
