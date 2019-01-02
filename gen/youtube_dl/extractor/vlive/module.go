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
 * vlive/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/vlive.py
 */

package vlive

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError                 λ.Object
	InfoExtractor                  λ.Object
	VLiveChannelIE                 λ.Object
	VLiveIE                        λ.Object
	VLivePlaylistIE                λ.Object
	ϒcompat_str                    λ.Object
	ϒcompat_urllib_parse_urlencode λ.Object
	ϒdict_get                      λ.Object
	ϒfloat_or_none                 λ.Object
	ϒint_or_none                   λ.Object
	ϒremove_start                  λ.Object
	ϒtry_get                       λ.Object
	ϒurlencode_postdata            λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_urllib_parse_urlencode = Ωcompat.ϒcompat_urllib_parse_urlencode
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒdict_get = Ωutils.ϒdict_get
		ExtractorError = Ωutils.ExtractorError
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒremove_start = Ωutils.ϒremove_start
		ϒtry_get = Ωutils.ϒtry_get
		ϒurlencode_postdata = Ωutils.ϒurlencode_postdata
		VLiveIE = λ.Cal(λ.TypeType, λ.NewStr("VLiveIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				VLiveIE_IE_NAME            λ.Object
				VLiveIE__TESTS             λ.Object
				VLiveIE__VALID_URL         λ.Object
				VLiveIE__get_common_fields λ.Object
				VLiveIE__real_extract      λ.Object
				VLiveIE__replay            λ.Object
				VLiveIE_suitable           λ.Object
			)
			VLiveIE_IE_NAME = λ.NewStr("vlive")
			VLiveIE__VALID_URL = λ.NewStr("https?://(?:(?:www|m)\\.)?vlive\\.tv/video/(?P<id>[0-9]+)")
			VLiveIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.vlive.tv/video/1326"),
					λ.NewStr("md5"): λ.NewStr("cc7314812855ce56de70a06a27314983"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):         λ.NewStr("1326"),
						λ.NewStr("ext"):        λ.NewStr("mp4"),
						λ.NewStr("title"):      λ.NewStr("[V LIVE] Girl's Day's Broadcast"),
						λ.NewStr("creator"):    λ.NewStr("Girl's Day"),
						λ.NewStr("view_count"): λ.IntType,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.vlive.tv/video/16937"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):         λ.NewStr("16937"),
						λ.NewStr("ext"):        λ.NewStr("mp4"),
						λ.NewStr("title"):      λ.NewStr("[V LIVE] 첸백시 걍방"),
						λ.NewStr("creator"):    λ.NewStr("EXO"),
						λ.NewStr("view_count"): λ.IntType,
						λ.NewStr("subtitles"):  λ.NewStr("mincount:12"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
			)
			VLiveIE_suitable = λ.NewFunction("suitable",
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
						if λ.IsTrue(λ.Cal(λ.GetAttr(VLivePlaylistIE, "suitable", nil), ϒurl)) {
							return λ.False
						} else {
							return λ.Cal(λ.GetAttr(λ.Cal(λ.SuperType, VLiveIE, ϒcls), "suitable", nil), ϒurl)
						}
					}()
				})
			VLiveIE_suitable = λ.Cal(λ.ClassMethodType, VLiveIE_suitable)
			VLiveIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						VIDEO_PARAMS_FIELD λ.Object
						VIDEO_PARAMS_RE    λ.Object
						ϒkey               λ.Object
						ϒlong_video_id     λ.Object
						ϒparams            λ.Object
						ϒself              = λargs[0]
						ϒstatus            λ.Object
						ϒurl               = λargs[1]
						ϒvideo_id          λ.Object
						ϒwebpage           λ.Object
						τmp0               λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), λ.Mod(λ.NewStr("https://www.vlive.tv/video/%s"), ϒvideo_id), ϒvideo_id)
					VIDEO_PARAMS_RE = λ.NewStr("\\bvlive\\.video\\.init\\(([^)]+)")
					VIDEO_PARAMS_FIELD = λ.NewStr("video params")
					ϒparams = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
						λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							VIDEO_PARAMS_RE,
							ϒwebpage,
							VIDEO_PARAMS_FIELD,
						), λ.KWArgs{
							{Name: "default", Value: λ.NewStr("")},
						}),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "transform_source", Value: λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "s"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒs = λargs[0]
								)
								return λ.Add(λ.Add(λ.NewStr("["), ϒs), λ.NewStr("]"))
							})},
						{Name: "fatal", Value: λ.False},
					})
					if λ.IsTrue(func() λ.Object {
						if λv := λ.NewBool(!λ.IsTrue(ϒparams)); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Lt(λ.Cal(λ.BuiltinLen, ϒparams), λ.NewInt(7))
						}
					}()) {
						ϒparams = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), VIDEO_PARAMS_RE, ϒwebpage, VIDEO_PARAMS_FIELD)
						ϒparams = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
							nil,
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
									var (
										ϒp   λ.Object
										τmp0 λ.Object
										τmp1 λ.Object
									)
									τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.None, λ.NewStr("\\s*,\\s*"), ϒparams))
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										ϒp = τmp1
										λgen.Yield(λ.Cal(λ.GetAttr(ϒp, "strip", nil), λ.NewStr("\"")))
									}
									return λ.None
								})
							})))
					}
					τmp0 = λ.NewTuple(
						λ.GetItem(ϒparams, λ.NewInt(2)),
						λ.GetItem(ϒparams, λ.NewInt(5)),
						λ.GetItem(ϒparams, λ.NewInt(6)),
					)
					ϒstatus = λ.GetItem(τmp0, λ.NewInt(0))
					ϒlong_video_id = λ.GetItem(τmp0, λ.NewInt(1))
					ϒkey = λ.GetItem(τmp0, λ.NewInt(2))
					ϒstatus = λ.Cal(ϒremove_start, ϒstatus, λ.NewStr("PRODUCT_"))
					if λ.IsTrue(λ.NewBool(λ.Contains(λ.NewTuple(
						λ.NewStr("LIVE_ON_AIR"),
						λ.NewStr("BIG_EVENT_ON_AIR"),
					), ϒstatus))) {
						return λ.Cal(λ.GetAttr(ϒself, "_live", nil), ϒvideo_id, ϒwebpage)
					} else {
						if λ.IsTrue(λ.NewBool(λ.Contains(λ.NewTuple(
							λ.NewStr("VOD_ON_AIR"),
							λ.NewStr("BIG_EVENT_INTRO"),
						), ϒstatus))) {
							if λ.IsTrue(func() λ.Object {
								if λv := ϒlong_video_id; !λ.IsTrue(λv) {
									return λv
								} else {
									return ϒkey
								}
							}()) {
								return λ.Cal(λ.GetAttr(ϒself, "_replay", nil), ϒvideo_id, ϒwebpage, ϒlong_video_id, ϒkey)
							} else {
								ϒstatus = λ.NewStr("COMING_SOON")
							}
						}
					}
					if λ.IsTrue(λ.Eq(ϒstatus, λ.NewStr("LIVE_END"))) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.NewStr("Uploading for replay. Please wait...")), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					} else {
						if λ.IsTrue(λ.Eq(ϒstatus, λ.NewStr("COMING_SOON"))) {
							panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.NewStr("Coming soon!")), λ.KWArgs{
								{Name: "expected", Value: λ.True},
							})))
						} else {
							if λ.IsTrue(λ.Eq(ϒstatus, λ.NewStr("CANCELED"))) {
								panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.NewStr("We are sorry, but the live broadcast has been canceled.")), λ.KWArgs{
									{Name: "expected", Value: λ.True},
								})))
							} else {
								panic(λ.Raise(λ.Cal(ExtractorError, λ.Mod(λ.NewStr("Unknown status %s"), ϒstatus))))
							}
						}
					}
					return λ.None
				})
			VLiveIE__get_common_fields = λ.NewFunction("_get_common_fields",
				[]λ.Param{
					{Name: "self"},
					{Name: "webpage"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcreator   λ.Object
						ϒself      = λargs[0]
						ϒthumbnail λ.Object
						ϒtitle     λ.Object
						ϒwebpage   = λargs[1]
					)
					ϒtitle = λ.Cal(λ.GetAttr(ϒself, "_og_search_title", nil), ϒwebpage)
					ϒcreator = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("<div[^>]+class=\"info_area\"[^>]*>\\s*<a\\s+[^>]*>([^<]+)"),
						ϒwebpage,
						λ.NewStr("creator"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒthumbnail = λ.Cal(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), ϒwebpage)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("title"):     ϒtitle,
						λ.NewStr("creator"):   ϒcreator,
						λ.NewStr("thumbnail"): ϒthumbnail,
					})
				})
			VLiveIE__replay = λ.NewFunction("_replay",
				[]λ.Param{
					{Name: "self"},
					{Name: "video_id"},
					{Name: "webpage"},
					{Name: "long_video_id"},
					{Name: "key"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcaption       λ.Object
						ϒformats       λ.Object
						ϒinfo          λ.Object
						ϒkey           = λargs[4]
						ϒlang          λ.Object
						ϒlong_video_id = λargs[3]
						ϒplayinfo      λ.Object
						ϒself          = λargs[0]
						ϒsubtitles     λ.Object
						ϒvideo_id      = λargs[1]
						ϒview_count    λ.Object
						ϒwebpage       = λargs[2]
						τmp0           λ.Object
						τmp1           λ.Object
					)
					ϒplayinfo = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.NewStr("http://global.apis.naver.com/rmcnmv/rmcnmv/vod_play_videoInfo.json?%s"), λ.Cal(ϒcompat_urllib_parse_urlencode, λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("videoId"): ϒlong_video_id,
						λ.NewStr("key"):     ϒkey,
						λ.NewStr("ptc"):     λ.NewStr("http"),
						λ.NewStr("doct"):    λ.NewStr("json"),
						λ.NewStr("cpt"):     λ.NewStr("vtt"),
					}))), ϒvideo_id)
					ϒformats = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒvid λ.Object
									τmp0 λ.Object
									τmp1 λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒplayinfo, "get", nil), λ.NewStr("videos"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("list"), λ.NewList()))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒvid = τmp1
									if λ.IsTrue(λ.Cal(λ.GetAttr(ϒvid, "get", nil), λ.NewStr("source"))) {
										λgen.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
											λ.NewStr("url"):       λ.GetItem(ϒvid, λ.NewStr("source")),
											λ.NewStr("format_id"): λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒvid, "get", nil), λ.NewStr("encodingOption"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("name")),
											λ.NewStr("abr"):       λ.Cal(ϒfloat_or_none, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒvid, "get", nil), λ.NewStr("bitrate"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("audio"))),
											λ.NewStr("vbr"):       λ.Cal(ϒfloat_or_none, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒvid, "get", nil), λ.NewStr("bitrate"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("video"))),
											λ.NewStr("width"):     λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒvid, "get", nil), λ.NewStr("encodingOption"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("width"))),
											λ.NewStr("height"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒvid, "get", nil), λ.NewStr("encodingOption"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("height"))),
											λ.NewStr("filesize"):  λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvid, "get", nil), λ.NewStr("size"))),
										}))
									}
								}
								return λ.None
							})
						})))
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒview_count = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒplayinfo, "get", nil), λ.NewStr("meta"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("count")))
					ϒsubtitles = λ.NewDictWithTable(map[λ.Object]λ.Object{})
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒplayinfo, "get", nil), λ.NewStr("captions"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("list"), λ.NewList()))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒcaption = τmp1
						ϒlang = λ.Cal(ϒdict_get, ϒcaption, λ.NewTuple(
							λ.NewStr("locale"),
							λ.NewStr("language"),
							λ.NewStr("country"),
							λ.NewStr("label"),
						))
						if λ.IsTrue(func() λ.Object {
							if λv := ϒlang; !λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(λ.GetAttr(ϒcaption, "get", nil), λ.NewStr("source"))
							}
						}()) {
							λ.SetItem(ϒsubtitles, ϒlang, λ.NewList(λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("ext"): λ.NewStr("vtt"),
								λ.NewStr("url"): λ.GetItem(ϒcaption, λ.NewStr("source")),
							})))
						}
					}
					ϒinfo = λ.Cal(λ.GetAttr(ϒself, "_get_common_fields", nil), ϒwebpage)
					λ.Cal(λ.GetAttr(ϒinfo, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):         ϒvideo_id,
						λ.NewStr("formats"):    ϒformats,
						λ.NewStr("view_count"): ϒview_count,
						λ.NewStr("subtitles"):  ϒsubtitles,
					}))
					return ϒinfo
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):            VLiveIE_IE_NAME,
				λ.NewStr("_TESTS"):             VLiveIE__TESTS,
				λ.NewStr("_VALID_URL"):         VLiveIE__VALID_URL,
				λ.NewStr("_get_common_fields"): VLiveIE__get_common_fields,
				λ.NewStr("_real_extract"):      VLiveIE__real_extract,
				λ.NewStr("_replay"):            VLiveIE__replay,
				λ.NewStr("suitable"):           VLiveIE_suitable,
			})
		}())
		VLiveChannelIE = λ.Cal(λ.TypeType, λ.NewStr("VLiveChannelIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				VLiveChannelIE__VALID_URL λ.Object
			)
			VLiveChannelIE__VALID_URL = λ.NewStr("https?://channels\\.vlive\\.tv/(?P<id>[0-9A-Z]+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): VLiveChannelIE__VALID_URL,
			})
		}())
		VLivePlaylistIE = λ.Cal(λ.TypeType, λ.NewStr("VLivePlaylistIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				VLivePlaylistIE__VALID_URL λ.Object
			)
			VLivePlaylistIE__VALID_URL = λ.NewStr("https?://(?:(?:www|m)\\.)?vlive\\.tv/video/(?P<video_id>[0-9]+)/playlist/(?P<id>[0-9]+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): VLivePlaylistIE__VALID_URL,
			})
		}())
	})
}