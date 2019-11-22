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
 * viewlift/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/viewlift.py
 */

package viewlift

import (
	Ωbase64 "github.com/tenta-browser/go-video-downloader/gen/base64"
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError               λ.Object
	InfoExtractor                λ.Object
	ViewLiftBaseIE               λ.Object
	ViewLiftEmbedIE              λ.Object
	ViewLiftIE                   λ.Object
	ϒclean_html                  λ.Object
	ϒcompat_urllib_parse_unquote λ.Object
	ϒdetermine_ext               λ.Object
	ϒint_or_none                 λ.Object
	ϒjs_to_json                  λ.Object
	ϒparse_age_limit             λ.Object
	ϒparse_duration              λ.Object
	ϒtry_get                     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_urllib_parse_unquote = Ωcompat.ϒcompat_urllib_parse_unquote
		ExtractorError = Ωutils.ExtractorError
		ϒclean_html = Ωutils.ϒclean_html
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒjs_to_json = Ωutils.ϒjs_to_json
		ϒparse_age_limit = Ωutils.ϒparse_age_limit
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒtry_get = Ωutils.ϒtry_get
		ViewLiftBaseIE = λ.Cal(λ.TypeType, λ.NewStr("ViewLiftBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ViewLiftBaseIE__DOMAINS_REGEX λ.Object
			)
			ViewLiftBaseIE__DOMAINS_REGEX = λ.NewStr("(?:(?:main\\.)?snagfilms|snagxtreme|funnyforfree|kiddovid|winnersview|(?:monumental|lax)sportsnetwork|vayafilm)\\.com|hoichoi\\.tv")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_DOMAINS_REGEX"): ViewLiftBaseIE__DOMAINS_REGEX,
			})
		}())
		ViewLiftEmbedIE = λ.Cal(λ.TypeType, λ.NewStr("ViewLiftEmbedIE"), λ.NewTuple(ViewLiftBaseIE), func() λ.Dict {
			var (
				ViewLiftEmbedIE__VALID_URL    λ.Object
				ViewLiftEmbedIE__real_extract λ.Object
			)
			ViewLiftEmbedIE__VALID_URL = λ.Mod(λ.NewStr("https?://(?:(?:www|embed)\\.)?(?:%s)/embed/player\\?.*\\bfilmId=(?P<id>[\\da-f]{8}-(?:[\\da-f]{4}-){3}[\\da-f]{12})"), λ.GetAttr(ViewLiftBaseIE, "_DOMAINS_REGEX", nil))
			ViewLiftEmbedIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒbitrate          λ.Object
						ϒext              λ.Object
						ϒfield_preference λ.Object
						ϒfile_            λ.Object
						ϒformat_id        λ.Object
						ϒformats          λ.Object
						ϒhas_bitrate      λ.Object
						ϒheight           λ.Object
						ϒhls_url          λ.Object
						ϒself             = λargs[0]
						ϒsource           λ.Object
						ϒsources          λ.Object
						ϒtitle            λ.Object
						ϒtype_            λ.Object
						ϒurl              = λargs[1]
						ϒvideo_id         λ.Object
						ϒwebpage          λ.Object
						τmp0              λ.Object
						τmp1              λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					if λ.IsTrue(λ.NewBool(λ.Contains(ϒwebpage, λ.NewStr(">This film is not playable in your area.<")))) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("Film %s is not playable in your area."), ϒvideo_id)), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒformats = λ.NewList()
					ϒhas_bitrate = λ.False
					ϒsources = λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("(?s)sources:\\s*(\\[.+?\\]),"),
						ϒwebpage,
						λ.NewStr("sources"),
					), λ.KWArgs{
						{Name: "default", Value: λ.NewStr("[]")},
					}), ϒvideo_id, ϒjs_to_json)
					τmp0 = λ.Cal(λ.BuiltinIter, ϒsources)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒsource = τmp1
						ϒfile_ = λ.Cal(λ.GetAttr(ϒsource, "get", nil), λ.NewStr("file"))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒfile_))) {
							continue
						}
						ϒtype_ = λ.Cal(λ.GetAttr(ϒsource, "get", nil), λ.NewStr("type"))
						ϒext = λ.Cal(ϒdetermine_ext, ϒfile_)
						ϒformat_id = func() λ.Object {
							if λv := λ.Cal(λ.GetAttr(ϒsource, "get", nil), λ.NewStr("label")); λ.IsTrue(λv) {
								return λv
							} else {
								return ϒext
							}
						}()
						if λ.IsTrue(λ.Cal(λ.BuiltinAll, λ.Cal(λ.NewFunction("<generator>",
							nil,
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
									var (
										ϒv   λ.Object
										τmp0 λ.Object
										τmp1 λ.Object
									)
									τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
										ϒtype_,
										ϒext,
									))
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										ϒv = τmp1
										λgy.Yield(λ.NewBool(λ.Contains(λ.NewTuple(
											λ.NewStr("m3u8"),
											λ.NewStr("hls"),
										), ϒv)))
									}
									return λ.None
								})
							})))) {
							λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
								ϒfile_,
								ϒvideo_id,
								λ.NewStr("mp4"),
								λ.NewStr("m3u8_native"),
							), λ.KWArgs{
								{Name: "m3u8_id", Value: λ.NewStr("hls")},
								{Name: "fatal", Value: λ.False},
							}))
						} else {
							ϒbitrate = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.NewList(
									λ.NewStr("(\\d+)kbps"),
									λ.Mod(λ.NewStr("_\\d{1,2}x\\d{1,2}_(\\d{3,})\\.%s"), ϒext),
								),
								ϒfile_,
								λ.NewStr("bitrate"),
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							}))
							if λ.IsTrue(func() λ.Object {
								if λv := λ.NewBool(!λ.IsTrue(ϒhas_bitrate)); !λ.IsTrue(λv) {
									return λv
								} else {
									return ϒbitrate
								}
							}()) {
								ϒhas_bitrate = λ.True
							}
							ϒheight = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.NewStr("^(\\d+)[pP]$"),
								ϒformat_id,
								λ.NewStr("height"),
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							}))
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("url"): ϒfile_,
								λ.NewStr("format_id"): λ.Mod(λ.NewStr("http-%s%s"), λ.NewTuple(
									ϒformat_id,
									func() λ.Object {
										if λ.IsTrue(ϒbitrate) {
											return λ.Mod(λ.NewStr("-%dk"), ϒbitrate)
										} else {
											return λ.NewStr("")
										}
									}(),
								)),
								λ.NewStr("tbr"):    ϒbitrate,
								λ.NewStr("height"): ϒheight,
							}))
						}
					}
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒformats))) {
						ϒhls_url = λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("filmInfo\\.src\\s*=\\s*({.+?});"), ϒwebpage, λ.NewStr("src")), ϒvideo_id, ϒjs_to_json), λ.NewStr("src"))
						ϒformats = λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
							ϒhls_url,
							ϒvideo_id,
							λ.NewStr("mp4"),
							λ.NewStr("m3u8_native"),
						), λ.KWArgs{
							{Name: "m3u8_id", Value: λ.NewStr("hls")},
							{Name: "fatal", Value: λ.False},
						})
					}
					ϒfield_preference = func() λ.Object {
						if λ.IsTrue(ϒhas_bitrate) {
							return λ.None
						} else {
							return λ.NewTuple(
								λ.NewStr("height"),
								λ.NewStr("tbr"),
								λ.NewStr("format_id"),
							)
						}
					}()
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats, ϒfield_preference)
					ϒtitle = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewList(
						λ.NewStr("title\\s*:\\s*'([^']+)'"),
						λ.NewStr("<title>([^<]+)</title>"),
					), ϒwebpage, λ.NewStr("title"))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):      ϒvideo_id,
						λ.NewStr("title"):   ϒtitle,
						λ.NewStr("formats"): ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"):    ViewLiftEmbedIE__VALID_URL,
				λ.NewStr("_real_extract"): ViewLiftEmbedIE__real_extract,
			})
		}())
		ViewLiftIE = λ.Cal(λ.TypeType, λ.NewStr("ViewLiftIE"), λ.NewTuple(ViewLiftBaseIE), func() λ.Dict {
			var (
				ViewLiftIE__VALID_URL    λ.Object
				ViewLiftIE__real_extract λ.Object
				ViewLiftIE_suitable      λ.Object
			)
			ViewLiftIE__VALID_URL = λ.Mod(λ.NewStr("https?://(?:www\\.)?(?P<domain>%s)(?:/(?:films/title|show|(?:news/)?videos?))?/(?P<id>[^?#]+)"), λ.GetAttr(ViewLiftBaseIE, "_DOMAINS_REGEX", nil))
			ViewLiftIE_suitable = λ.NewFunction("suitable",
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
						if λ.IsTrue(λ.Cal(λ.GetAttr(ViewLiftEmbedIE, "suitable", nil), ϒurl)) {
							return λ.False
						} else {
							return λ.Cal(λ.GetAttr(λ.Cal(λ.SuperType, ViewLiftIE, ϒcls), "suitable", nil), ϒurl)
						}
					}()
				})
			ViewLiftIE_suitable = λ.Cal(λ.ClassMethodType, ViewLiftIE_suitable)
			ViewLiftIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒbitrate             λ.Object
						ϒcategories          λ.Object
						ϒcontent_data        λ.Object
						ϒdata                λ.Object
						ϒdescription         λ.Object
						ϒdisplay_id          λ.Object
						ϒdomain              λ.Object
						ϒduration            λ.Object
						ϒfilm_id             λ.Object
						ϒformats             λ.Object
						ϒgist                λ.Object
						ϒheight              λ.Object
						ϒhls_url             λ.Object
						ϒinfo                λ.Object
						ϒinitial_store_state λ.Object
						ϒitem                λ.Object
						ϒk                   λ.Object
						ϒmodules             λ.Object
						ϒmpeg_video_assets   λ.Object
						ϒself                = λargs[0]
						ϒsnag                λ.Object
						ϒthumbnail           λ.Object
						ϒtitle               λ.Object
						ϒtoken               λ.Object
						ϒurl                 = λargs[1]
						ϒvideo_asset         λ.Object
						ϒvideo_asset_url     λ.Object
						ϒvideo_assets        λ.Object
						ϒwebpage             λ.Object
						τmp0                 λ.Object
						τmp1                 λ.Object
					)
					τmp0 = λ.Cal(λ.GetAttr(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups", nil))
					ϒdomain = λ.GetItem(τmp0, λ.NewInt(0))
					ϒdisplay_id = λ.GetItem(τmp0, λ.NewInt(1))
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒdisplay_id)
					if λ.IsTrue(λ.NewBool(λ.Contains(ϒwebpage, λ.NewStr(">Sorry, the Film you're looking for is not available.<")))) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("Film %s is not available."), ϒdisplay_id)), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒinitial_store_state = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("window\\.initialStoreState\\s*=.*?JSON\\.parse\\(unescape\\(atob\\('([^']+)'\\)\\)\\)"),
						ϒwebpage,
						λ.NewStr("Initial Store State"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒinitial_store_state) {
						ϒmodules = λ.GetItem(λ.GetItem(λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), λ.Cal(ϒcompat_urllib_parse_unquote, λ.Cal(λ.GetAttr(λ.Cal(Ωbase64.ϒb64decode, ϒinitial_store_state), "decode", nil))), ϒdisplay_id), λ.NewStr("page")), λ.NewStr("data")), λ.NewStr("modules"))
						ϒcontent_data = λ.Cal(λ.BuiltinNext, λ.Cal(λ.NewFunction("<generator>",
							nil,
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
									var (
										ϒm   λ.Object
										τmp0 λ.Object
										τmp1 λ.Object
									)
									τmp0 = λ.Cal(λ.BuiltinIter, ϒmodules)
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										ϒm = τmp1
										if λ.IsTrue(λ.Eq(λ.Cal(λ.GetAttr(ϒm, "get", nil), λ.NewStr("moduleType")), λ.NewStr("VideoDetailModule"))) {
											λgy.Yield(λ.GetItem(λ.GetItem(ϒm, λ.NewStr("contentData")), λ.NewInt(0)))
										}
									}
									return λ.None
								})
							})))
						ϒgist = λ.GetItem(ϒcontent_data, λ.NewStr("gist"))
						ϒfilm_id = λ.GetItem(ϒgist, λ.NewStr("id"))
						ϒtitle = λ.GetItem(ϒgist, λ.NewStr("title"))
						ϒvideo_assets = λ.Cal(ϒtry_get, ϒcontent_data, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("streamingInfo")), λ.NewStr("videoAssets"))
							}), λ.DictType)
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒvideo_assets))) {
							ϒtoken = λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
								λ.NewStr("https://prod-api.viewlift.com/identity/anonymous-token"),
								ϒfilm_id,
								λ.NewStr("Downloading authorization token"),
							), λ.KWArgs{
								{Name: "query", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("site"): λ.NewStr("snagfilms"),
								})},
							}), λ.NewStr("authorizationToken"))
							ϒvideo_assets = λ.GetItem(λ.GetItem(λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
								λ.NewStr("https://prod-api.viewlift.com/entitlement/video/status"),
								ϒfilm_id,
							), λ.KWArgs{
								{Name: "headers", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("Authorization"): ϒtoken,
									λ.NewStr("Referer"):       ϒurl,
								})},
								{Name: "query", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("id"): ϒfilm_id,
								})},
							}), λ.NewStr("video")), λ.NewStr("streamingInfo")), λ.NewStr("videoAssets"))
						}
						ϒformats = λ.NewList()
						ϒmpeg_video_assets = func() λ.Object {
							if λv := λ.Cal(λ.GetAttr(ϒvideo_assets, "get", nil), λ.NewStr("mpeg")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewList()
							}
						}()
						τmp0 = λ.Cal(λ.BuiltinIter, ϒmpeg_video_assets)
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒvideo_asset = τmp1
							ϒvideo_asset_url = λ.Cal(λ.GetAttr(ϒvideo_asset, "get", nil), λ.NewStr("url"))
							if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒvideo_asset))) {
								continue
							}
							ϒbitrate = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo_asset, "get", nil), λ.NewStr("bitrate")))
							ϒheight = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.NewStr("^_?(\\d+)[pP]$"),
								λ.Cal(λ.GetAttr(ϒvideo_asset, "get", nil), λ.NewStr("renditionValue")),
								λ.NewStr("height"),
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							}))
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("url"): ϒvideo_asset_url,
								λ.NewStr("format_id"): λ.Mod(λ.NewStr("http%s"), func() λ.Object {
									if λ.IsTrue(ϒbitrate) {
										return λ.Mod(λ.NewStr("-%d"), ϒbitrate)
									} else {
										return λ.NewStr("")
									}
								}()),
								λ.NewStr("tbr"):    ϒbitrate,
								λ.NewStr("height"): ϒheight,
								λ.NewStr("vcodec"): λ.Cal(λ.GetAttr(ϒvideo_asset, "get", nil), λ.NewStr("codec")),
							}))
						}
						ϒhls_url = λ.Cal(λ.GetAttr(ϒvideo_assets, "get", nil), λ.NewStr("hls"))
						if λ.IsTrue(ϒhls_url) {
							λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
								ϒhls_url,
								ϒfilm_id,
								λ.NewStr("mp4"),
								λ.NewStr("m3u8_native"),
							), λ.KWArgs{
								{Name: "m3u8_id", Value: λ.NewStr("hls")},
								{Name: "fatal", Value: λ.False},
							}))
						}
						λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats, λ.NewTuple(
							λ.NewStr("height"),
							λ.NewStr("tbr"),
							λ.NewStr("format_id"),
						))
						ϒinfo = λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("id"):          ϒfilm_id,
							λ.NewStr("display_id"):  ϒdisplay_id,
							λ.NewStr("title"):       ϒtitle,
							λ.NewStr("description"): λ.Cal(λ.GetAttr(ϒgist, "get", nil), λ.NewStr("description")),
							λ.NewStr("thumbnail"):   λ.Cal(λ.GetAttr(ϒgist, "get", nil), λ.NewStr("videoImageUrl")),
							λ.NewStr("duration"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒgist, "get", nil), λ.NewStr("runtime"))),
							λ.NewStr("age_limit"):   λ.Cal(ϒparse_age_limit, λ.Cal(λ.GetAttr(ϒcontent_data, "get", nil), λ.NewStr("parentalRating"))),
							λ.NewStr("timestamp"):   λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒgist, "get", nil), λ.NewStr("publishDate")), λ.NewInt(1000)),
							λ.NewStr("formats"):     ϒformats,
						})
						τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
							λ.NewStr("categories"),
							λ.NewStr("tags"),
						))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒk = τmp1
							λ.SetItem(ϒinfo, ϒk, λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
								nil,
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
										var (
											ϒv   λ.Object
											τmp0 λ.Object
											τmp1 λ.Object
										)
										τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒcontent_data, "get", nil), ϒk, λ.NewList()))
										for {
											if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
												break
											}
											ϒv = τmp1
											if λ.IsTrue(λ.Cal(λ.GetAttr(ϒv, "get", nil), λ.NewStr("title"))) {
												λgy.Yield(λ.GetItem(ϒv, λ.NewStr("title")))
											}
										}
										return λ.None
									})
								}))))
						}
						return ϒinfo
					} else {
						ϒfilm_id = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("filmId=([\\da-f-]{36})\""), ϒwebpage, λ.NewStr("film id"))
						ϒsnag = λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("Snag\\.page\\.data\\s*=\\s*(\\[.+?\\]);"),
							ϒwebpage,
							λ.NewStr("snag"),
						), λ.KWArgs{
							{Name: "default", Value: λ.NewStr("[]")},
						}), ϒdisplay_id)
						τmp0 = λ.Cal(λ.BuiltinIter, ϒsnag)
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒitem = τmp1
							if λ.IsTrue(λ.Eq(λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒitem, "get", nil), λ.NewStr("data"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("film"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("id")), ϒfilm_id)) {
								ϒdata = λ.GetItem(λ.GetItem(ϒitem, λ.NewStr("data")), λ.NewStr("film"))
								ϒtitle = λ.GetItem(ϒdata, λ.NewStr("title"))
								ϒdescription = λ.Cal(ϒclean_html, λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("synopsis")))
								ϒthumbnail = λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("image"))
								ϒduration = λ.Cal(ϒint_or_none, func() λ.Object {
									if λv := λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("duration")); λ.IsTrue(λv) {
										return λv
									} else {
										return λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("runtime"))
									}
								}())
								ϒcategories = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
									nil,
									0, false, false,
									func(λargs []λ.Object) λ.Object {
										return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
											var (
												ϒcategory λ.Object
												τmp0      λ.Object
												τmp1      λ.Object
											)
											τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("categories"), λ.NewList()))
											for {
												if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
													break
												}
												ϒcategory = τmp1
												if λ.IsTrue(λ.Cal(λ.GetAttr(ϒcategory, "get", nil), λ.NewStr("title"))) {
													λgy.Yield(λ.GetItem(ϒcategory, λ.NewStr("title")))
												}
											}
											return λ.None
										})
									})))
								break
							}
						}
						if τmp1 == λ.AfterLast {
							ϒtitle = λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewTuple(
								λ.NewStr("itemprop=\"title\">([^<]+)<"),
								λ.NewStr("(?s)itemprop=\"title\">(.+?)<div"),
							), ϒwebpage, λ.NewStr("title"))
							ϒdescription = func() λ.Object {
								if λv := λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
									λ.NewStr("(?s)<div itemprop=\"description\" class=\"film-synopsis-inner \">(.+?)</div>"),
									ϒwebpage,
									λ.NewStr("description"),
								), λ.KWArgs{
									{Name: "default", Value: λ.None},
								}); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Cal(λ.GetAttr(ϒself, "_og_search_description", nil), ϒwebpage)
								}
							}()
							ϒthumbnail = λ.Cal(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), ϒwebpage)
							ϒduration = λ.Cal(ϒparse_duration, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.NewStr("<span itemprop=\"duration\" class=\"film-duration strong\">([^<]+)<"),
								ϒwebpage,
								λ.NewStr("duration"),
							), λ.KWArgs{
								{Name: "fatal", Value: λ.False},
							}))
							ϒcategories = λ.Cal(Ωre.ϒfindall, λ.NewStr("<a href=\"/movies/[^\"]+\">([^<]+)</a>"), ϒwebpage)
						}
						return λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("_type"): λ.NewStr("url_transparent"),
							λ.NewStr("url"): λ.Mod(λ.NewStr("http://%s/embed/player?filmId=%s"), λ.NewTuple(
								ϒdomain,
								ϒfilm_id,
							)),
							λ.NewStr("id"):          ϒfilm_id,
							λ.NewStr("display_id"):  ϒdisplay_id,
							λ.NewStr("title"):       ϒtitle,
							λ.NewStr("description"): ϒdescription,
							λ.NewStr("thumbnail"):   ϒthumbnail,
							λ.NewStr("duration"):    ϒduration,
							λ.NewStr("categories"):  ϒcategories,
							λ.NewStr("ie_key"):      λ.NewStr("ViewLiftEmbed"),
						})
					}
					return λ.None
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"):    ViewLiftIE__VALID_URL,
				λ.NewStr("_real_extract"): ViewLiftIE__real_extract,
				λ.NewStr("suitable"):      ViewLiftIE_suitable,
			})
		}())
	})
}
