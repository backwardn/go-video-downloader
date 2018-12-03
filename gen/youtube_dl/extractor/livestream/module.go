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
 * livestream/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/livestream.py
 */

package livestream

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωparse "github.com/tenta-browser/go-video-downloader/gen/urllib/parse"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor         λ.Object
	LivestreamIE          λ.Object
	LivestreamOriginalIE  λ.Object
	LivestreamShortenerIE λ.Object
	ϒcompat_str           λ.Object
	ϒdetermine_ext        λ.Object
	ϒfind_xpath_attr      λ.Object
	ϒfloat_or_none        λ.Object
	ϒint_or_none          λ.Object
	ϒorderedSet           λ.Object
	ϒparse_iso8601        λ.Object
	ϒupdate_url_query     λ.Object
	ϒxpath_attr           λ.Object
	ϒxpath_text           λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒfind_xpath_attr = Ωutils.ϒfind_xpath_attr
		ϒxpath_attr = Ωutils.ϒxpath_attr
		ϒxpath_text = Ωutils.ϒxpath_text
		ϒorderedSet = Ωutils.ϒorderedSet
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		LivestreamIE = λ.Cal(λ.TypeType, λ.NewStr("LivestreamIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				LivestreamIE_IE_NAME             λ.Object
				LivestreamIE__API_URL_TEMPLATE   λ.Object
				LivestreamIE__TESTS              λ.Object
				LivestreamIE__VALID_URL          λ.Object
				LivestreamIE__extract_video_info λ.Object
				LivestreamIE__parse_smil_formats λ.Object
				LivestreamIE__real_extract       λ.Object
			)
			LivestreamIE_IE_NAME = λ.NewStr("livestream")
			LivestreamIE__VALID_URL = λ.NewStr("https?://(?:new\\.)?livestream\\.com/(?:accounts/(?P<account_id>\\d+)|(?P<account_name>[^/]+))/(?:events/(?P<event_id>\\d+)|(?P<event_name>[^/]+))(?:/videos/(?P<id>\\d+))?")
			LivestreamIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://new.livestream.com/CoheedandCambria/WebsterHall/videos/4719370"),
					λ.NewStr("md5"): λ.NewStr("53274c76ba7754fb0e8d072716f2292b"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("4719370"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Live from Webster Hall NYC"),
						λ.NewStr("timestamp"):   λ.NewInt(1350008072),
						λ.NewStr("upload_date"): λ.NewStr("20121012"),
						λ.NewStr("duration"):    λ.NewFloat(5968.0),
						λ.NewStr("like_count"):  λ.IntType,
						λ.NewStr("view_count"):  λ.IntType,
						λ.NewStr("thumbnail"):   λ.NewStr("re:^http://.*\\.jpg$"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://new.livestream.com/tedx/cityenglish"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("title"): λ.NewStr("TEDCity2.0 (English)"),
						λ.NewStr("id"):    λ.NewStr("2245590"),
					}),
					λ.NewStr("playlist_mincount"): λ.NewInt(4),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://new.livestream.com/chess24/tatasteelchess"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("title"): λ.NewStr("Tata Steel Chess"),
						λ.NewStr("id"):    λ.NewStr("3705884"),
					}),
					λ.NewStr("playlist_mincount"): λ.NewInt(60),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://new.livestream.com/accounts/362/events/3557232/videos/67864563/player?autoPlay=false&height=360&mute=false&width=640"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://livestream.com/bsww/concacafbeachsoccercampeonato2015"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			LivestreamIE__API_URL_TEMPLATE = λ.NewStr("http://livestream.com/api/accounts/%s/events/%s")
			LivestreamIE__parse_smil_formats = λ.NewFunction("_parse_smil_formats",
				[]λ.Param{
					{Name: "self"},
					{Name: "smil"},
					{Name: "smil_url"},
					{Name: "video_id"},
					{Name: "namespace", Def: λ.None},
					{Name: "f4m_params", Def: λ.None},
					{Name: "transform_rtmp_url", Def: λ.None},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒbase               λ.Object
						ϒbase_ele           λ.Object
						ϒf4m_params         = λargs[5]
						ϒformats            λ.Object
						ϒfurl               λ.Object
						ϒnamespace          = λargs[4]
						ϒself               = λargs[0]
						ϒsmil               = λargs[1]
						ϒsmil_url           = λargs[2]
						ϒtbr                λ.Object
						ϒtransform_rtmp_url = λargs[6]
						ϒvideo_id           = λargs[3]
						ϒvideo_nodes        λ.Object
						ϒvn                 λ.Object
						τmp0                λ.Object
						τmp1                λ.Object
						τmp2                λ.Object
					)
					_ = ϒf4m_params
					_ = ϒsmil_url
					_ = ϒtransform_rtmp_url
					_ = ϒvideo_id
					ϒbase_ele = λ.Cal(ϒfind_xpath_attr, ϒsmil, λ.Cal(λ.GetAttr(ϒself, "_xpath_ns", nil), λ.NewStr(".//meta"), ϒnamespace), λ.NewStr("name"), λ.NewStr("httpBase"))
					ϒbase = func() λ.Object {
						if λ.IsTrue(λ.NewBool(ϒbase_ele != λ.None)) {
							return λ.Cal(λ.GetAttr(ϒbase_ele, "get", nil), λ.NewStr("content"))
						} else {
							return λ.NewStr("http://livestreamvod-f.akamaihd.net/")
						}
					}()
					ϒformats = λ.NewList()
					ϒvideo_nodes = λ.Cal(λ.GetAttr(ϒsmil, "findall", nil), λ.Cal(λ.GetAttr(ϒself, "_xpath_ns", nil), λ.NewStr(".//video"), ϒnamespace))
					τmp0 = λ.Cal(λ.BuiltinIter, ϒvideo_nodes)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒvn = τmp1
						ϒtbr = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(λ.GetAttr(ϒvn, "attrib", nil), "get", nil), λ.NewStr("system-bitrate")), λ.NewInt(1000))
						ϒfurl = λ.Cal(ϒupdate_url_query, λ.Cal(Ωparse.ϒurljoin, ϒbase, λ.GetItem(λ.GetAttr(ϒvn, "attrib", nil), λ.NewStr("src"))), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("v"):  λ.NewStr("3.0.3"),
							λ.NewStr("fp"): λ.NewStr("WIN% 14,0,0,145"),
						}))
						if λ.IsTrue(λ.NewBool(λ.Contains(λ.GetAttr(ϒvn, "attrib", nil), λ.NewStr("clipBegin")))) {
							τmp2 = λ.IAdd(ϒfurl, λ.Add(λ.NewStr("&ssek="), λ.GetItem(λ.GetAttr(ϒvn, "attrib", nil), λ.NewStr("clipBegin"))))
							ϒfurl = τmp2
						}
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"):        ϒfurl,
							λ.NewStr("format_id"):  λ.Mod(λ.NewStr("smil_%d"), ϒtbr),
							λ.NewStr("ext"):        λ.NewStr("flv"),
							λ.NewStr("tbr"):        ϒtbr,
							λ.NewStr("preference"): λ.Neg(λ.NewInt(1000)),
						}))
					}
					return ϒformats
				})
			LivestreamIE__extract_video_info = λ.NewFunction("_extract_video_info",
				[]λ.Param{
					{Name: "self"},
					{Name: "video_data"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						FORMAT_KEYS λ.Object
						ϒbitrate    λ.Object
						ϒcomments   λ.Object
						ϒext        λ.Object
						ϒf4m_url    λ.Object
						ϒformat_id  λ.Object
						ϒformats    λ.Object
						ϒkey        λ.Object
						ϒm3u8_url   λ.Object
						ϒself       = λargs[0]
						ϒsmil_url   λ.Object
						ϒvideo_data = λargs[1]
						ϒvideo_id   λ.Object
						ϒvideo_url  λ.Object
						τmp0        λ.Object
						τmp1        λ.Object
						τmp2        λ.Object
					)
					ϒvideo_id = λ.Cal(ϒcompat_str, λ.GetItem(ϒvideo_data, λ.NewStr("id")))
					FORMAT_KEYS = λ.NewTuple(
						λ.NewTuple(
							λ.NewStr("sd"),
							λ.NewStr("progressive_url"),
						),
						λ.NewTuple(
							λ.NewStr("hd"),
							λ.NewStr("progressive_url_hd"),
						),
					)
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, FORMAT_KEYS)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒformat_id = λ.GetItem(τmp2, λ.NewInt(0))
						ϒkey = λ.GetItem(τmp2, λ.NewInt(1))
						ϒvideo_url = λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), ϒkey)
						if λ.IsTrue(ϒvideo_url) {
							ϒext = λ.Cal(ϒdetermine_ext, ϒvideo_url)
							if λ.IsTrue(λ.Eq(ϒext, λ.NewStr("m3u8"))) {
								continue
							}
							ϒbitrate = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.Mod(λ.NewStr("(\\d+)\\.%s"), ϒext),
								ϒvideo_url,
								λ.NewStr("bitrate"),
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							}))
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("url"):       ϒvideo_url,
								λ.NewStr("format_id"): ϒformat_id,
								λ.NewStr("tbr"):       ϒbitrate,
								λ.NewStr("ext"):       ϒext,
							}))
						}
					}
					ϒsmil_url = λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("smil_url"))
					if λ.IsTrue(ϒsmil_url) {
						λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_smil_formats", nil), λ.NewArgs(
							ϒsmil_url,
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "fatal", Value: λ.False},
						}))
					}
					ϒm3u8_url = λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("m3u8_url"))
					if λ.IsTrue(ϒm3u8_url) {
						λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
							ϒm3u8_url,
							ϒvideo_id,
							λ.NewStr("mp4"),
							λ.NewStr("m3u8_native"),
						), λ.KWArgs{
							{Name: "m3u8_id", Value: λ.NewStr("hls")},
							{Name: "fatal", Value: λ.False},
						}))
					}
					ϒf4m_url = λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("f4m_url"))
					if λ.IsTrue(ϒf4m_url) {
						λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
							ϒf4m_url,
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "f4m_id", Value: λ.NewStr("hds")},
							{Name: "fatal", Value: λ.False},
						}))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒcomments = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒcomment λ.Object
									τmp0     λ.Object
									τmp1     λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("comments"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("data"), λ.NewList()))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒcomment = τmp1
									λgen.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("author_id"): λ.Cal(λ.GetAttr(ϒcomment, "get", nil), λ.NewStr("author_id")),
										λ.NewStr("author"):    λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒcomment, "get", nil), λ.NewStr("author"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("full_name")),
										λ.NewStr("id"):        λ.Cal(λ.GetAttr(ϒcomment, "get", nil), λ.NewStr("id")),
										λ.NewStr("text"):      λ.GetItem(ϒcomment, λ.NewStr("text")),
										λ.NewStr("timestamp"): λ.Cal(ϒparse_iso8601, λ.Cal(λ.GetAttr(ϒcomment, "get", nil), λ.NewStr("created_at"))),
									}))
								}
								return λ.None
							})
						})))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):            ϒvideo_id,
						λ.NewStr("formats"):       ϒformats,
						λ.NewStr("title"):         λ.GetItem(ϒvideo_data, λ.NewStr("caption")),
						λ.NewStr("description"):   λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("description")),
						λ.NewStr("thumbnail"):     λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("thumbnail_url")),
						λ.NewStr("duration"):      λ.Cal(ϒfloat_or_none, λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("duration")), λ.NewInt(1000)),
						λ.NewStr("timestamp"):     λ.Cal(ϒparse_iso8601, λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("publish_at"))),
						λ.NewStr("like_count"):    λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("likes"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("total")),
						λ.NewStr("comment_count"): λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("comments"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("total")),
						λ.NewStr("view_count"):    λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("views")),
						λ.NewStr("comments"):      ϒcomments,
					})
				})
			LivestreamIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒaccount    λ.Object
						ϒapi_url    λ.Object
						ϒevent      λ.Object
						ϒevent_data λ.Object
						ϒmobj       λ.Object
						ϒself       = λargs[0]
						ϒurl        = λargs[1]
						ϒvideo_data λ.Object
						ϒvideo_id   λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					ϒevent = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("event_id")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("event_name"))
						}
					}()
					ϒaccount = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("account_id")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("account_name"))
						}
					}()
					ϒapi_url = λ.Mod(λ.GetAttr(ϒself, "_API_URL_TEMPLATE", nil), λ.NewTuple(
						ϒaccount,
						ϒevent,
					))
					if λ.IsTrue(ϒvideo_id) {
						ϒvideo_data = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Add(ϒapi_url, λ.Mod(λ.NewStr("/videos/%s"), ϒvideo_id)), ϒvideo_id)
						return λ.Cal(λ.GetAttr(ϒself, "_extract_video_info", nil), ϒvideo_data)
					} else {
						ϒevent_data = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), ϒapi_url, ϒvideo_id)
						return λ.Cal(λ.GetAttr(ϒself, "_extract_event", nil), ϒevent_data)
					}
					return λ.None
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):             LivestreamIE_IE_NAME,
				λ.NewStr("_API_URL_TEMPLATE"):   LivestreamIE__API_URL_TEMPLATE,
				λ.NewStr("_TESTS"):              LivestreamIE__TESTS,
				λ.NewStr("_VALID_URL"):          LivestreamIE__VALID_URL,
				λ.NewStr("_extract_video_info"): LivestreamIE__extract_video_info,
				λ.NewStr("_parse_smil_formats"): LivestreamIE__parse_smil_formats,
				λ.NewStr("_real_extract"):       LivestreamIE__real_extract,
			})
		}())
		LivestreamOriginalIE = λ.Cal(λ.TypeType, λ.NewStr("LivestreamOriginalIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				LivestreamOriginalIE__VALID_URL λ.Object
			)
			LivestreamOriginalIE__VALID_URL = λ.NewStr("(?x)https?://original\\.livestream\\.com/\n        (?P<user>[^/\\?#]+)(?:/(?P<type>video|folder)\n        (?:(?:\\?.*?Id=|/)(?P<id>.*?)(&|$))?)?\n        ")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): LivestreamOriginalIE__VALID_URL,
			})
		}())
		LivestreamShortenerIE = λ.Cal(λ.TypeType, λ.NewStr("LivestreamShortenerIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				LivestreamShortenerIE__VALID_URL λ.Object
			)
			LivestreamShortenerIE__VALID_URL = λ.NewStr("https?://livestre\\.am/(?P<id>.+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): LivestreamShortenerIE__VALID_URL,
			})
		}())
	})
}
