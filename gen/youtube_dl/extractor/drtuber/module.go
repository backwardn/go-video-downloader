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
 * drtuber/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/drtuber.py
 */

package drtuber

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	DrTuberIE       λ.Object
	InfoExtractor   λ.Object
	NO_DEFAULT      λ.Object
	ϒint_or_none    λ.Object
	ϒparse_duration λ.Object
	ϒstr_to_int     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒint_or_none = Ωutils.ϒint_or_none
		NO_DEFAULT = Ωutils.NO_DEFAULT
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒstr_to_int = Ωutils.ϒstr_to_int
		DrTuberIE = λ.Cal(λ.TypeType, λ.StrLiteral("DrTuberIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				DrTuberIE__VALID_URL    λ.Object
				DrTuberIE__real_extract λ.Object
			)
			DrTuberIE__VALID_URL = λ.StrLiteral("https?://(?:(?:www|m)\\.)?drtuber\\.com/(?:video|embed)/(?P<id>\\d+)(?:/(?P<display_id>[\\w-]+))?")
			DrTuberIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcategories    λ.Object
						ϒcats_str      λ.Object
						ϒcomment_count λ.Object
						ϒdislike_count λ.Object
						ϒdisplay_id    λ.Object
						ϒduration      λ.Object
						ϒextract_count λ.Object
						ϒformat_id     λ.Object
						ϒformats       λ.Object
						ϒlike_count    λ.Object
						ϒmobj          λ.Object
						ϒself          = λargs[0]
						ϒthumbnail     λ.Object
						ϒtitle         λ.Object
						ϒurl           = λargs[1]
						ϒvideo_data    λ.Object
						ϒvideo_id      λ.Object
						ϒvideo_url     λ.Object
						ϒwebpage       λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
						τmp2           λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("id"))
					ϒdisplay_id = func() λ.Object {
						if λv := λ.Calm(ϒmobj, "group", λ.StrLiteral("display_id")); λ.IsTrue(λv) {
							return λv
						} else {
							return ϒvideo_id
						}
					}()
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", λ.Mod(λ.StrLiteral("http://www.drtuber.com/video/%s"), ϒvideo_id), ϒdisplay_id)
					ϒvideo_data = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.StrLiteral("http://www.drtuber.com/player_config_json/"),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "query", Value: λ.DictLiteral(map[string]λ.Object{
							"vid":       ϒvideo_id,
							"embed":     λ.IntLiteral(0),
							"aid":       λ.IntLiteral(0),
							"domain_id": λ.IntLiteral(0),
						})},
					})
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(λ.GetItem(ϒvideo_data, λ.StrLiteral("files")), "items"))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒformat_id = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒvideo_url = λ.GetItem(τmp2, λ.IntLiteral(1))
						if λ.IsTrue(ϒvideo_url) {
							λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
								"format_id": ϒformat_id,
								"quality": func() λ.Object {
									if λ.IsTrue(λ.Eq(ϒformat_id, λ.StrLiteral("hq"))) {
										return λ.IntLiteral(2)
									} else {
										return λ.IntLiteral(1)
									}
								}(),
								"url": ϒvideo_url,
							}))
						}
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒduration = func() λ.Object {
						if λv := λ.Cal(ϒint_or_none, λ.Calm(ϒvideo_data, "get", λ.StrLiteral("duration"))); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(ϒparse_duration, λ.Calm(ϒvideo_data, "get", λ.StrLiteral("duration_format")))
						}
					}()
					ϒtitle = λ.Calm(ϒself, "_html_search_regex", λ.NewTuple(
						λ.StrLiteral("<h1[^>]+class=[\"\\']title[^>]+>([^<]+)"),
						λ.StrLiteral("<title>([^<]+)\\s*@\\s+DrTuber"),
						λ.StrLiteral("class=\"title_watch\"[^>]*><(?:p|h\\d+)[^>]*>([^<]+)<"),
						λ.StrLiteral("<p[^>]+class=\"title_substrate\">([^<]+)</p>"),
						λ.StrLiteral("<title>([^<]+) - \\d+"),
					), ϒwebpage, λ.StrLiteral("title"))
					ϒthumbnail = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("poster=\"([^\"]+)\""),
						ϒwebpage,
						λ.StrLiteral("thumbnail"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒextract_count = λ.NewFunction("extract_count",
						[]λ.Param{
							{Name: "id_"},
							{Name: "name"},
							{Name: "default", Def: NO_DEFAULT},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒdefault = λargs[2]
								ϒid_     = λargs[0]
								ϒname    = λargs[1]
							)
							return λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
								λ.Mod(λ.StrLiteral("<span[^>]+(?:class|id)=\"%s\"[^>]*>([\\d,\\.]+)</span>"), ϒid_),
								ϒwebpage,
								λ.Mod(λ.StrLiteral("%s count"), ϒname),
							), λ.KWArgs{
								{Name: "default", Value: ϒdefault},
								{Name: "fatal", Value: λ.False},
							}))
						})
					ϒlike_count = λ.Cal(ϒextract_count, λ.StrLiteral("rate_likes"), λ.StrLiteral("like"))
					ϒdislike_count = λ.Call(ϒextract_count, λ.NewArgs(
						λ.StrLiteral("rate_dislikes"),
						λ.StrLiteral("dislike"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					ϒcomment_count = λ.Cal(ϒextract_count, λ.StrLiteral("comments_count"), λ.StrLiteral("comment"))
					ϒcats_str = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("<div[^>]+class=\"categories_list\">(.+?)</div>"),
						ϒwebpage,
						λ.StrLiteral("categories"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒcategories = func() λ.Object {
						if !λ.IsTrue(ϒcats_str) {
							return λ.NewList()
						} else {
							return λ.Cal(Ωre.ϒfindall, λ.StrLiteral("<a title=\"([^\"]+)\""), ϒcats_str)
						}
					}()
					return λ.DictLiteral(map[string]λ.Object{
						"id":            ϒvideo_id,
						"display_id":    ϒdisplay_id,
						"formats":       ϒformats,
						"title":         ϒtitle,
						"thumbnail":     ϒthumbnail,
						"like_count":    ϒlike_count,
						"dislike_count": ϒdislike_count,
						"comment_count": ϒcomment_count,
						"categories":    ϒcategories,
						"age_limit":     λ.Calm(ϒself, "_rta_search", ϒwebpage),
						"duration":      ϒduration,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    DrTuberIE__VALID_URL,
				"_real_extract": DrTuberIE__real_extract,
			})
		}())
	})
}
