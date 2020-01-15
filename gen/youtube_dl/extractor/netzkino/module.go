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
 * netzkino/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/netzkino.py
 */

package netzkino

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor  λ.Object
	NetzkinoIE     λ.Object
	ϒclean_html    λ.Object
	ϒint_or_none   λ.Object
	ϒjs_to_json    λ.Object
	ϒparse_iso8601 λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒclean_html = Ωutils.ϒclean_html
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒjs_to_json = Ωutils.ϒjs_to_json
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		NetzkinoIE = λ.Cal(λ.TypeType, λ.StrLiteral("NetzkinoIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				NetzkinoIE__VALID_URL    λ.Object
				NetzkinoIE__real_extract λ.Object
			)
			NetzkinoIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?netzkino\\.de/\\#!/(?P<category>[^/]+)/(?P<id>[^/]+)")
			NetzkinoIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒapi_info      λ.Object
						ϒapi_url       λ.Object
						ϒavo_js        λ.Object
						ϒcategory_id   λ.Object
						ϒcomments      λ.Object
						ϒcustom_fields λ.Object
						ϒfilm_fn       λ.Object
						ϒformats       λ.Object
						ϒinfo          λ.Object
						ϒmobj          λ.Object
						ϒproduction_js λ.Object
						ϒself          = λargs[0]
						ϒsuffix        λ.Object
						ϒtemplates     λ.Object
						ϒurl           = λargs[1]
						ϒvideo_id      λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒcategory_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("category"))
					ϒvideo_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("id"))
					ϒapi_url = λ.Mod(λ.StrLiteral("http://api.netzkino.de.simplecache.net/capi-2.0a/categories/%s.json?d=www"), ϒcategory_id)
					ϒapi_info = λ.Calm(ϒself, "_download_json", ϒapi_url, ϒvideo_id)
					ϒinfo = λ.Cal(λ.BuiltinNext, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒp   λ.Object
									τmp0 λ.Object
									τmp1 λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(ϒapi_info, λ.StrLiteral("posts")))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒp = τmp1
									if λ.IsTrue(λ.Eq(λ.GetItem(ϒp, λ.StrLiteral("slug")), ϒvideo_id)) {
										λgy.Yield(ϒp)
									}
								}
								return λ.None
							})
						})))
					ϒcustom_fields = λ.GetItem(ϒinfo, λ.StrLiteral("custom_fields"))
					ϒproduction_js = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
						λ.StrLiteral("http://www.netzkino.de/beta/dist/production.min.js"),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "note", Value: λ.StrLiteral("Downloading player code")},
					})
					ϒavo_js = λ.Calm(ϒself, "_search_regex", λ.StrLiteral("var urlTemplate=(\\{.*?\"\\})"), ϒproduction_js, λ.StrLiteral("URL templates"))
					ϒtemplates = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
						ϒavo_js,
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "transform_source", Value: ϒjs_to_json},
					})
					ϒsuffix = λ.DictLiteral(map[string]string{
						"hds": ".mp4/manifest.f4m",
						"hls": ".mp4/master.m3u8",
						"pmd": ".mp4",
					})
					ϒfilm_fn = λ.GetItem(λ.GetItem(ϒcustom_fields, λ.StrLiteral("Streaming")), λ.IntLiteral(0))
					ϒformats = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒkey λ.Object
									ϒtpl λ.Object
									τmp0 λ.Object
									τmp1 λ.Object
									τmp2 λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒtemplates, "items"))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									τmp2 = τmp1
									ϒkey = λ.GetItem(τmp2, λ.IntLiteral(0))
									ϒtpl = λ.GetItem(τmp2, λ.IntLiteral(1))
									λgy.Yield(λ.DictLiteral(map[string]λ.Object{
										"format_id": ϒkey,
										"ext":       λ.StrLiteral("mp4"),
										"url":       λ.Add(λ.Calm(ϒtpl, "replace", λ.StrLiteral("{}"), ϒfilm_fn), λ.GetItem(ϒsuffix, ϒkey)),
									}))
								}
								return λ.None
							})
						})))
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒcomments = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒc   λ.Object
									τmp0 λ.Object
									τmp1 λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒinfo, "get", λ.StrLiteral("comments"), λ.NewList()))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒc = τmp1
									λgy.Yield(λ.DictLiteral(map[string]λ.Object{
										"timestamp": λ.Call(ϒparse_iso8601, λ.NewArgs(λ.Calm(ϒc, "get", λ.StrLiteral("date"))), λ.KWArgs{
											{Name: "delimiter", Value: λ.StrLiteral(" ")},
										}),
										"id":     λ.GetItem(ϒc, λ.StrLiteral("id")),
										"author": λ.GetItem(ϒc, λ.StrLiteral("name")),
										"html":   λ.GetItem(ϒc, λ.StrLiteral("content")),
										"parent": func() λ.Object {
											if λ.IsTrue(λ.Eq(λ.Calm(ϒc, "get", λ.StrLiteral("parent"), λ.IntLiteral(0)), λ.IntLiteral(0))) {
												return λ.StrLiteral("root")
											} else {
												return λ.GetItem(ϒc, λ.StrLiteral("parent"))
											}
										}(),
									}))
								}
								return λ.None
							})
						})))
					return λ.DictLiteral(map[string]λ.Object{
						"id":        ϒvideo_id,
						"formats":   ϒformats,
						"comments":  ϒcomments,
						"title":     λ.GetItem(ϒinfo, λ.StrLiteral("title")),
						"age_limit": λ.Cal(ϒint_or_none, λ.GetItem(λ.Calm(ϒcustom_fields, "get", λ.StrLiteral("FSK")), λ.IntLiteral(0))),
						"timestamp": λ.Call(ϒparse_iso8601, λ.NewArgs(λ.Calm(ϒinfo, "get", λ.StrLiteral("date"))), λ.KWArgs{
							{Name: "delimiter", Value: λ.StrLiteral(" ")},
						}),
						"description":    λ.Cal(ϒclean_html, λ.Calm(ϒinfo, "get", λ.StrLiteral("content"))),
						"thumbnail":      λ.Calm(ϒinfo, "get", λ.StrLiteral("thumbnail")),
						"playlist_title": λ.Calm(ϒapi_info, "get", λ.StrLiteral("title")),
						"playlist_id":    ϒcategory_id,
					})
				})
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL":    NetzkinoIE__VALID_URL,
				"_real_extract": NetzkinoIE__real_extract,
			})
		}())
	})
}
