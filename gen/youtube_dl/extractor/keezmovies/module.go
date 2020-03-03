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
 * keezmovies/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/keezmovies.py
 */

package keezmovies

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError               λ.Object
	InfoExtractor                λ.Object
	KeezMoviesIE                 λ.Object
	ϒcompat_urllib_parse_unquote λ.Object
	ϒdetermine_ext               λ.Object
	ϒint_or_none                 λ.Object
	ϒstr_to_int                  λ.Object
	ϒstrip_or_none               λ.Object
	ϒurl_or_none                 λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_urllib_parse_unquote = Ωcompat.ϒcompat_urllib_parse_unquote
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒstr_to_int = Ωutils.ϒstr_to_int
		ϒstrip_or_none = Ωutils.ϒstrip_or_none
		ϒurl_or_none = Ωutils.ϒurl_or_none
		KeezMoviesIE = λ.Cal(λ.TypeType, λ.StrLiteral("KeezMoviesIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				KeezMoviesIE__VALID_URL    λ.Object
				KeezMoviesIE__extract_info λ.Object
				KeezMoviesIE__real_extract λ.Object
			)
			KeezMoviesIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?keezmovies\\.com/video/(?:(?P<display_id>[^/]+)-)?(?P<id>\\d+)")
			KeezMoviesIE__extract_info = λ.NewFunction("_extract_info",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
					{Name: "fatal", Def: λ.True},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id     λ.Object
						ϒduration       λ.Object
						ϒencrypted      λ.Object
						ϒextract_format λ.Object
						ϒfatal          = λargs[2]
						ϒflashvars      λ.Object
						ϒformat_urls    λ.Object
						ϒformats        λ.Object
						ϒkey            λ.Object
						ϒmobj           λ.Object
						ϒself           = λargs[0]
						ϒthumbnail      λ.Object
						ϒtitle          λ.Object
						ϒurl            = λargs[1]
						ϒvalue          λ.Object
						ϒvideo_id       λ.Object
						ϒvideo_url      λ.Object
						ϒwebpage        λ.Object
						τmp0            λ.Object
						τmp1            λ.Object
						τmp2            λ.Object
					)
					_ = τmp0
					_ = τmp1
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("id"))
					ϒdisplay_id = func() λ.Object {
						if λv := func() λ.Object {
							if λ.Contains(λ.Calm(ϒmobj, "groupdict"), λ.StrLiteral("display_id")) {
								return λ.Calm(ϒmobj, "group", λ.StrLiteral("display_id"))
							} else {
								return λ.None
							}
						}(); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒmobj, "group", λ.StrLiteral("id"))
						}
					}()
					ϒwebpage = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
						ϒurl,
						ϒdisplay_id,
					), λ.KWArgs{
						{Name: "headers", Value: λ.DictLiteral(map[string]string{
							"Cookie": "age_verified=1",
						})},
					})
					ϒformats = λ.NewList()
					ϒformat_urls = λ.Cal(λ.SetType)
					ϒtitle = λ.None
					ϒthumbnail = λ.None
					ϒduration = λ.None
					ϒencrypted = λ.False
					ϒextract_format = λ.NewFunction("extract_format",
						[]λ.Param{
							{Name: "format_url"},
							{Name: "height", Def: λ.None},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒformat_url = λargs[0]
								ϒheight     = λargs[1]
								ϒtbr        λ.Object
							)
							ϒformat_url = λ.Cal(ϒurl_or_none, ϒformat_url)
							if λ.IsTrue(func() λ.Object {
								if λv := λ.NewBool(!λ.IsTrue(ϒformat_url)); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.NewBool(!λ.IsTrue(λ.Calm(ϒformat_url, "startswith", λ.NewTuple(
										λ.StrLiteral("http"),
										λ.StrLiteral("//"),
									))))
								}
							}()) {
								return λ.None
							}
							if λ.Contains(ϒformat_urls, ϒformat_url) {
								return λ.None
							}
							λ.Calm(ϒformat_urls, "add", ϒformat_url)
							ϒtbr = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("[/_](\\d+)[kK][/_]"),
								ϒformat_url,
								λ.StrLiteral("tbr"),
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							}))
							if !λ.IsTrue(ϒheight) {
								ϒheight = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
									λ.StrLiteral("[/_](\\d+)[pP][/_]"),
									ϒformat_url,
									λ.StrLiteral("height"),
								), λ.KWArgs{
									{Name: "default", Value: λ.None},
								}))
							}
							if λ.IsTrue(ϒencrypted) {
								ϒformat_url = λ.Calm(λ.Cal(λ.None, λ.None, ϒtitle, λ.IntLiteral(32)), "decode", λ.StrLiteral("utf-8"))
							}
							λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
								"url": ϒformat_url,
								"format_id": func() λ.Object {
									if λ.IsTrue(ϒheight) {
										return λ.Mod(λ.StrLiteral("%dp"), ϒheight)
									} else {
										return λ.None
									}
								}(),
								"height": ϒheight,
								"tbr":    ϒtbr,
							}))
							return λ.None
						})
					ϒflashvars = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
						λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.StrLiteral("flashvars\\s*=\\s*({.+?});"),
							ϒwebpage,
							λ.StrLiteral("flashvars"),
						), λ.KWArgs{
							{Name: "default", Value: λ.StrLiteral("{}")},
						}),
						ϒdisplay_id,
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					if λ.IsTrue(ϒflashvars) {
						ϒtitle = λ.Calm(ϒflashvars, "get", λ.StrLiteral("video_title"))
						ϒthumbnail = λ.Calm(ϒflashvars, "get", λ.StrLiteral("image_url"))
						ϒduration = λ.Cal(ϒint_or_none, λ.Calm(ϒflashvars, "get", λ.StrLiteral("video_duration")))
						ϒencrypted = λ.NewBool(λ.Calm(ϒflashvars, "get", λ.StrLiteral("encrypted")) == λ.True)
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒflashvars, "items"))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							τmp2 = τmp1
							ϒkey = λ.GetItem(τmp2, λ.IntLiteral(0))
							ϒvalue = λ.GetItem(τmp2, λ.IntLiteral(1))
							ϒmobj = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("quality_(\\d+)[pP]"), ϒkey)
							if λ.IsTrue(ϒmobj) {
								λ.Cal(ϒextract_format, ϒvalue, λ.Cal(λ.IntType, λ.Calm(ϒmobj, "group", λ.IntLiteral(1))))
							}
						}
						ϒvideo_url = λ.Calm(ϒflashvars, "get", λ.StrLiteral("video_url"))
						if λ.IsTrue(func() λ.Object {
							if λv := ϒvideo_url; !λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(ϒdetermine_ext, ϒvideo_url, λ.None)
							}
						}()) {
							λ.Cal(ϒextract_format, ϒvideo_url)
						}
					}
					ϒvideo_url = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("flashvars\\.video_url\\s*=\\s*([\"\\'])(?P<url>http.+?)\\1"),
						ϒwebpage,
						λ.StrLiteral("video url"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
						{Name: "group", Value: λ.StrLiteral("url")},
					})
					if λ.IsTrue(ϒvideo_url) {
						λ.Cal(ϒextract_format, λ.Cal(ϒcompat_urllib_parse_unquote, ϒvideo_url))
					}
					if !λ.IsTrue(ϒformats) {
						if λ.Contains(ϒwebpage, λ.StrLiteral("title=\"This video is no longer available\"")) {
							panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.StrLiteral("Video %s is no longer available"), ϒvideo_id)), λ.KWArgs{
								{Name: "expected", Value: λ.True},
							})))
						}
					}
					τmp0, τmp1 = func() (λexit λ.Object, λret λ.Object) {
						defer λ.CatchMulti(
							nil,
							&λ.Catcher{ExtractorError, func(λex λ.BaseException) {
								if λ.IsTrue(ϒfatal) {
									panic(λ.Raise(λex))
								}
							}},
						)
						λ.Calm(ϒself, "_sort_formats", ϒformats)
						return λ.BlockExitNormally, nil
					}()
					if !λ.IsTrue(ϒtitle) {
						ϒtitle = λ.Calm(ϒself, "_html_search_regex", λ.StrLiteral("<h1[^>]*>([^<]+)"), ϒwebpage, λ.StrLiteral("title"))
					}
					return λ.NewTuple(
						ϒwebpage,
						λ.DictLiteral(map[string]λ.Object{
							"id":         ϒvideo_id,
							"display_id": ϒdisplay_id,
							"title":      λ.Cal(ϒstrip_or_none, ϒtitle),
							"thumbnail":  ϒthumbnail,
							"duration":   ϒduration,
							"age_limit":  λ.IntLiteral(18),
							"formats":    ϒformats,
						}),
					)
				})
			KeezMoviesIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒinfo    λ.Object
						ϒself    = λargs[0]
						ϒurl     = λargs[1]
						ϒwebpage λ.Object
						τmp0     λ.Object
					)
					τmp0 = λ.Call(λ.GetAttr(ϒself, "_extract_info", nil), λ.NewArgs(ϒurl), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒwebpage = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒinfo = λ.GetItem(τmp0, λ.IntLiteral(1))
					if !λ.IsTrue(λ.GetItem(ϒinfo, λ.StrLiteral("formats"))) {
						return λ.Calm(ϒself, "url_result", ϒurl, λ.StrLiteral("Generic"))
					}
					λ.SetItem(ϒinfo, λ.StrLiteral("view_count"), λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("<b>([\\d,.]+)</b> Views?"),
						ϒwebpage,
						λ.StrLiteral("view count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})))
					return ϒinfo
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    KeezMoviesIE__VALID_URL,
				"_extract_info": KeezMoviesIE__extract_info,
				"_real_extract": KeezMoviesIE__real_extract,
			})
		}())
	})
}
