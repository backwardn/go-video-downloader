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
 * tube8/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/tube8.py
 */

package tube8

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωkeezmovies "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/keezmovies"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	KeezMoviesIE λ.Object
	Tube8IE      λ.Object
	ϒint_or_none λ.Object
	ϒstr_to_int  λ.Object
)

func init() {
	λ.InitModule(func() {
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒstr_to_int = Ωutils.ϒstr_to_int
		KeezMoviesIE = Ωkeezmovies.KeezMoviesIE
		Tube8IE = λ.Cal(λ.TypeType, λ.StrLiteral("Tube8IE"), λ.NewTuple(KeezMoviesIE), func() λ.Dict {
			var (
				Tube8IE__VALID_URL    λ.Object
				Tube8IE__real_extract λ.Object
			)
			Tube8IE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?tube8\\.com/(?:[^/]+/)+(?P<display_id>[^/]+)/(?P<id>\\d+)")
			Tube8IE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcategories    λ.Object
						ϒcategory      λ.Object
						ϒcomment_count λ.Object
						ϒdescription   λ.Object
						ϒdislike_count λ.Object
						ϒinfo          λ.Object
						ϒlike_count    λ.Object
						ϒself          = λargs[0]
						ϒtags          λ.Object
						ϒtags_str      λ.Object
						ϒuploader      λ.Object
						ϒurl           = λargs[1]
						ϒview_count    λ.Object
						ϒwebpage       λ.Object
						τmp0           λ.Object
					)
					τmp0 = λ.Calm(ϒself, "_extract_info", ϒurl)
					ϒwebpage = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒinfo = λ.GetItem(τmp0, λ.IntLiteral(1))
					if !λ.IsTrue(λ.GetItem(ϒinfo, λ.StrLiteral("title"))) {
						λ.SetItem(ϒinfo, λ.StrLiteral("title"), λ.Calm(ϒself, "_html_search_regex", λ.StrLiteral("videoTitle\\s*=\\s*\"([^\"]+)"), ϒwebpage, λ.StrLiteral("title")))
					}
					ϒdescription = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("(?s)Description:</dt>\\s*<dd>(.+?)</dd>"),
						ϒwebpage,
						λ.StrLiteral("description"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒuploader = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("<span class=\"username\">\\s*(.+?)\\s*<"),
						ϒwebpage,
						λ.StrLiteral("uploader"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒlike_count = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("rupVar\\s*=\\s*\"(\\d+)\""),
						ϒwebpage,
						λ.StrLiteral("like count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒdislike_count = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("rdownVar\\s*=\\s*\"(\\d+)\""),
						ϒwebpage,
						λ.StrLiteral("dislike count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒview_count = λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("Views:\\s*</dt>\\s*<dd>([\\d,\\.]+)"),
						ϒwebpage,
						λ.StrLiteral("view count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒcomment_count = λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("<span id=\"allCommentsCount\">(\\d+)</span>"),
						ϒwebpage,
						λ.StrLiteral("comment count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒcategory = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("Category:\\s*</dt>\\s*<dd>\\s*<a[^>]+href=[^>]+>([^<]+)"),
						ϒwebpage,
						λ.StrLiteral("category"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒcategories = func() λ.Object {
						if λ.IsTrue(ϒcategory) {
							return λ.NewList(ϒcategory)
						} else {
							return λ.None
						}
					}()
					ϒtags_str = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("(?s)Tags:\\s*</dt>\\s*<dd>(.+?)</(?!a)"),
						ϒwebpage,
						λ.StrLiteral("tags"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒtags = func() λ.Object {
						if λ.IsTrue(ϒtags_str) {
							return λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
								nil,
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
										var (
											ϒt   λ.Object
											τmp0 λ.Object
											τmp1 λ.Object
										)
										τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfindall, λ.StrLiteral("<a[^>]+href=[^>]+>([^<]+)"), ϒtags_str))
										for {
											if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
												break
											}
											ϒt = τmp1
											λgy.Yield(ϒt)
										}
										return λ.None
									})
								})))
						} else {
							return λ.None
						}
					}()
					λ.Calm(ϒinfo, "update", λ.DictLiteral(map[string]λ.Object{
						"description":   ϒdescription,
						"uploader":      ϒuploader,
						"view_count":    ϒview_count,
						"like_count":    ϒlike_count,
						"dislike_count": ϒdislike_count,
						"comment_count": ϒcomment_count,
						"categories":    ϒcategories,
						"tags":          ϒtags,
					}))
					return ϒinfo
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    Tube8IE__VALID_URL,
				"_real_extract": Tube8IE__real_extract,
			})
		}())
	})
}
