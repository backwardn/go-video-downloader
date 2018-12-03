// Code generated by transpiler. DO NOT EDIT.

/**
 * Go Video Downloader
 *
 *    Copyright 2018 Tenta, LLC
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
 * bitchute/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/bitchute.py
 */

package bitchute

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	BitChuteChannelIE   λ.Object
	BitChuteIE          λ.Object
	InfoExtractor       λ.Object
	ϒurlencode_postdata λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒurlencode_postdata = Ωutils.ϒurlencode_postdata
		BitChuteIE = λ.Cal(λ.TypeType, λ.NewStr("BitChuteIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				BitChuteIE__TESTS        λ.Object
				BitChuteIE__VALID_URL    λ.Object
				BitChuteIE__real_extract λ.Object
			)
			BitChuteIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?bitchute\\.com/(?:video|embed|torrent/[^/]+)/(?P<id>[^/?#&]+)")
			BitChuteIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.bitchute.com/video/szoMrox2JEI/"),
					λ.NewStr("md5"): λ.NewStr("66c4a70e6bfc40dcb6be3eb1d74939eb"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("szoMrox2JEI"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Fuck bitches get money"),
						λ.NewStr("description"): λ.NewStr("md5:3f21f6fb5b1d17c3dee9cf6b5fe60b3a"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.jpg$"),
						λ.NewStr("uploader"):    λ.NewStr("Victoria X Rave"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.bitchute.com/embed/lbb5G1hjPhw/"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.bitchute.com/torrent/Zee5BE49045h/szoMrox2JEI.webtorrent"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			BitChuteIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdescription λ.Object
						ϒformats     λ.Object
						ϒself        = λargs[0]
						ϒthumbnail   λ.Object
						ϒtitle       λ.Object
						ϒuploader    λ.Object
						ϒurl         = λargs[1]
						ϒvideo_id    λ.Object
						ϒwebpage     λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
						λ.Mod(λ.NewStr("https://www.bitchute.com/video/%s"), ϒvideo_id),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "headers", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("User-Agent"): λ.NewStr("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.57 Safari/537.36"),
						})},
					})
					ϒtitle = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
							λ.NewTuple(
								λ.NewStr("<[^>]+\\bid=[\"\\']video-title[^>]+>([^<]+)"),
								λ.NewStr("<title>([^<]+)"),
							),
							ϒwebpage,
							λ.NewStr("title"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else if λv := λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
							λ.NewStr("description"),
							ϒwebpage,
							λ.NewStr("title"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒself, "_og_search_description", nil), ϒwebpage)
						}
					}()
					ϒformats = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒmobj λ.Object
									τmp0  λ.Object
									τmp1  λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfinditer, λ.NewStr("addWebSeed\\s*\\(\\s*([\"\\'])(?P<url>(?:(?!\\1).)+)\\1"), ϒwebpage))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒmobj = τmp1
									λgen.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("url"): λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("url")),
									}))
								}
								return λ.None
							})
						})))
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒdescription = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("(?s)<div\\b[^>]+\\bclass=[\"\\']full hidden[^>]+>(.+?)</div>"),
						ϒwebpage,
						λ.NewStr("description"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒthumbnail = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewStr("twitter:image:src"), ϒwebpage, λ.NewStr("thumbnail"))
						}
					}()
					ϒuploader = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("(?s)<p\\b[^>]+\\bclass=[\"\\']video-author[^>]+>(.+?)</p>"),
						ϒwebpage,
						λ.NewStr("uploader"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("thumbnail"):   ϒthumbnail,
						λ.NewStr("uploader"):    ϒuploader,
						λ.NewStr("formats"):     ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        BitChuteIE__TESTS,
				λ.NewStr("_VALID_URL"):    BitChuteIE__VALID_URL,
				λ.NewStr("_real_extract"): BitChuteIE__real_extract,
			})
		}())
		BitChuteChannelIE = λ.Cal(λ.TypeType, λ.NewStr("BitChuteChannelIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				BitChuteChannelIE__VALID_URL λ.Object
			)
			BitChuteChannelIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?bitchute\\.com/channel/(?P<id>[^/?#&]+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): BitChuteChannelIE__VALID_URL,
			})
		}())
	})
}
