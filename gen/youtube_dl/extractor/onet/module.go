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
 * onet/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/onet.py
 */

package onet

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError        λ.Object
	InfoExtractor         λ.Object
	NO_DEFAULT            λ.Object
	OnetBaseIE            λ.Object
	OnetChannelIE         λ.Object
	OnetIE                λ.Object
	OnetMVPIE             λ.Object
	OnetPlIE              λ.Object
	ϒdetermine_ext        λ.Object
	ϒfloat_or_none        λ.Object
	ϒget_element_by_class λ.Object
	ϒint_or_none          λ.Object
	ϒjs_to_json           λ.Object
	ϒparse_iso8601        λ.Object
	ϒremove_start         λ.Object
	ϒstrip_or_none        λ.Object
	ϒurl_basename         λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒget_element_by_class = Ωutils.ϒget_element_by_class
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒjs_to_json = Ωutils.ϒjs_to_json
		NO_DEFAULT = Ωutils.NO_DEFAULT
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒremove_start = Ωutils.ϒremove_start
		ϒstrip_or_none = Ωutils.ϒstrip_or_none
		ϒurl_basename = Ωutils.ϒurl_basename
		OnetBaseIE = λ.Cal(λ.TypeType, λ.NewStr("OnetBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				OnetBaseIE__extract_from_id λ.Object
				OnetBaseIE__search_mvp_id   λ.Object
			)
			OnetBaseIE__search_mvp_id = λ.NewFunction("_search_mvp_id",
				[]λ.Param{
					{Name: "self"},
					{Name: "webpage"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself    = λargs[0]
						ϒwebpage = λargs[1]
					)
					return λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("id=([\"\\'])mvp:(?P<id>.+?)\\1"),
						ϒwebpage,
						λ.NewStr("mvp id"),
					), λ.KWArgs{
						{Name: "group", Value: λ.NewStr("id")},
					})
				})
			OnetBaseIE__extract_from_id = λ.NewFunction("_extract_from_id",
				[]λ.Param{
					{Name: "self"},
					{Name: "video_id"},
					{Name: "webpage", Def: λ.None},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdescription  λ.Object
						ϒduration     λ.Object
						ϒerror        λ.Object
						ϒext          λ.Object
						ϒf            λ.Object
						ϒformat_id    λ.Object
						ϒformat_list  λ.Object
						ϒformats      λ.Object
						ϒformats_dict λ.Object
						ϒmeta         λ.Object
						ϒresponse     λ.Object
						ϒself         = λargs[0]
						ϒtimestamp    λ.Object
						ϒtitle        λ.Object
						ϒvideo        λ.Object
						ϒvideo_id     = λargs[1]
						ϒvideo_url    λ.Object
						ϒwebpage      = λargs[2]
						τmp0          λ.Object
						τmp1          λ.Object
						τmp2          λ.Object
						τmp3          λ.Object
						τmp4          λ.Object
						τmp5          λ.Object
					)
					ϒresponse = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.NewStr("http://qi.ckm.onetapi.pl/"),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "query", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("body[id]"):                    ϒvideo_id,
							λ.NewStr("body[jsonrpc]"):               λ.NewStr("2.0"),
							λ.NewStr("body[method]"):                λ.NewStr("get_asset_detail"),
							λ.NewStr("body[params][ID_Publikacji]"): ϒvideo_id,
							λ.NewStr("body[params][Service]"):       λ.NewStr("www.onet.pl"),
							λ.NewStr("content-type"):                λ.NewStr("application/jsonp"),
							λ.NewStr("x-onet-app"):                  λ.NewStr("player.front.onetapi.pl"),
						})},
					})
					ϒerror = λ.Cal(λ.GetAttr(ϒresponse, "get", nil), λ.NewStr("error"))
					if λ.IsTrue(ϒerror) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("%s said: %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							λ.GetItem(ϒerror, λ.NewStr("message")),
						))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒvideo = λ.Cal(λ.GetAttr(λ.GetItem(ϒresponse, λ.NewStr("result")), "get", nil), λ.NewStr("0"))
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(λ.GetItem(ϒvideo, λ.NewStr("formats")), "items", nil)))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						_ = λ.GetItem(τmp2, λ.NewInt(0))
						ϒformats_dict = λ.GetItem(τmp2, λ.NewInt(1))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒformats_dict, λ.DictType)))) {
							continue
						}
						τmp2 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒformats_dict, "items", nil)))
						for {
							if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
								break
							}
							τmp4 = τmp3
							ϒformat_id = λ.GetItem(τmp4, λ.NewInt(0))
							ϒformat_list = λ.GetItem(τmp4, λ.NewInt(1))
							if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒformat_list, λ.ListType)))) {
								continue
							}
							τmp4 = λ.Cal(λ.BuiltinIter, ϒformat_list)
							for {
								if τmp5 = λ.NextDefault(τmp4, λ.AfterLast); τmp5 == λ.AfterLast {
									break
								}
								ϒf = τmp5
								ϒvideo_url = λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("url"))
								if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒvideo_url))) {
									continue
								}
								ϒext = λ.Cal(ϒdetermine_ext, ϒvideo_url)
								if λ.IsTrue(λ.Eq(ϒformat_id, λ.NewStr("ism"))) {
									λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_ism_formats", nil), λ.NewArgs(
										ϒvideo_url,
										ϒvideo_id,
										λ.NewStr("mss"),
									), λ.KWArgs{
										{Name: "fatal", Value: λ.False},
									}))
								} else {
									if λ.IsTrue(λ.Eq(ϒext, λ.NewStr("mpd"))) {
										λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_mpd_formats", nil), λ.NewArgs(
											ϒvideo_url,
											ϒvideo_id,
										), λ.KWArgs{
											{Name: "mpd_id", Value: λ.NewStr("dash")},
											{Name: "fatal", Value: λ.False},
										}))
									} else {
										λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
											λ.NewStr("url"):       ϒvideo_url,
											λ.NewStr("format_id"): ϒformat_id,
											λ.NewStr("height"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("vertical_resolution"))),
											λ.NewStr("width"):     λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("horizontal_resolution"))),
											λ.NewStr("abr"):       λ.Cal(ϒfloat_or_none, λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("audio_bitrate"))),
											λ.NewStr("vbr"):       λ.Cal(ϒfloat_or_none, λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("video_bitrate"))),
										}))
									}
								}
							}
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒmeta = λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("meta"), λ.NewDictWithTable(map[λ.Object]λ.Object{}))
					ϒtitle = func() λ.Object {
						if λv := func() λ.Object {
							if λ.IsTrue(ϒwebpage) {
								return λ.Call(λ.GetAttr(ϒself, "_og_search_title", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
									{Name: "default", Value: λ.None},
								})
							} else {
								return λ.None
							}
						}(); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.GetItem(ϒmeta, λ.NewStr("title"))
						}
					}()
					ϒdescription = func() λ.Object {
						if λv := func() λ.Object {
							if λ.IsTrue(ϒwebpage) {
								return λ.Call(λ.GetAttr(ϒself, "_og_search_description", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
									{Name: "default", Value: λ.None},
								})
							} else {
								return λ.None
							}
						}(); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒmeta, "get", nil), λ.NewStr("description"))
						}
					}()
					ϒduration = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒmeta, "get", nil), λ.NewStr("length")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒmeta, "get", nil), λ.NewStr("lenght"))
						}
					}()
					ϒtimestamp = λ.Cal(ϒparse_iso8601, λ.Cal(λ.GetAttr(ϒmeta, "get", nil), λ.NewStr("addDate")), λ.NewStr(" "))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("duration"):    ϒduration,
						λ.NewStr("timestamp"):   ϒtimestamp,
						λ.NewStr("formats"):     ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_extract_from_id"): OnetBaseIE__extract_from_id,
				λ.NewStr("_search_mvp_id"):   OnetBaseIE__search_mvp_id,
			})
		}())
		OnetMVPIE = λ.Cal(λ.TypeType, λ.NewStr("OnetMVPIE"), λ.NewTuple(OnetBaseIE), func() λ.Dict {
			var (
				OnetMVPIE__VALID_URL λ.Object
			)
			OnetMVPIE__VALID_URL = λ.NewStr("onetmvp:(?P<id>\\d+\\.\\d+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): OnetMVPIE__VALID_URL,
			})
		}())
		OnetIE = λ.Cal(λ.TypeType, λ.NewStr("OnetIE"), λ.NewTuple(OnetBaseIE), func() λ.Dict {
			var (
				OnetIE_IE_NAME       λ.Object
				OnetIE__TEST         λ.Object
				OnetIE__VALID_URL    λ.Object
				OnetIE__real_extract λ.Object
			)
			OnetIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?onet\\.tv/[a-z]/[a-z]+/(?P<display_id>[0-9a-z-]+)/(?P<id>[0-9a-z]+)")
			OnetIE_IE_NAME = λ.NewStr("onet.tv")
			OnetIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("http://onet.tv/k/openerfestival/open-er-festival-2016-najdziwniejsze-wymagania-gwiazd/qbpyqc"),
				λ.NewStr("md5"): λ.NewStr("e3ffbf47590032ac3f27249204173d50"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):          λ.NewStr("qbpyqc"),
					λ.NewStr("display_id"):  λ.NewStr("open-er-festival-2016-najdziwniejsze-wymagania-gwiazd"),
					λ.NewStr("ext"):         λ.NewStr("mp4"),
					λ.NewStr("title"):       λ.NewStr("Open'er Festival 2016: najdziwniejsze wymagania gwiazd"),
					λ.NewStr("description"): λ.NewStr("Trzy samochody, których nigdy nie użyto, prywatne spa, hotel dekorowany czarnym suknem czy nielegalne używki. Organizatorzy koncertów i festiwali muszą stawać przed nie lada wyzwaniem zapraszając gwia..."),
					λ.NewStr("upload_date"): λ.NewStr("20160705"),
					λ.NewStr("timestamp"):   λ.NewInt(1467721580),
				}),
			})
			OnetIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id λ.Object
						ϒinfo_dict  λ.Object
						ϒmobj       λ.Object
						ϒmvp_id     λ.Object
						ϒself       = λargs[0]
						ϒurl        = λargs[1]
						ϒvideo_id   λ.Object
						ϒwebpage    λ.Object
						τmp0        λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					τmp0 = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("display_id"), λ.NewStr("id"))
					ϒdisplay_id = λ.GetItem(τmp0, λ.NewInt(0))
					ϒvideo_id = λ.GetItem(τmp0, λ.NewInt(1))
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒdisplay_id)
					ϒmvp_id = λ.Cal(λ.GetAttr(ϒself, "_search_mvp_id", nil), ϒwebpage)
					ϒinfo_dict = λ.Cal(λ.GetAttr(ϒself, "_extract_from_id", nil), ϒmvp_id, ϒwebpage)
					λ.Cal(λ.GetAttr(ϒinfo_dict, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):         ϒvideo_id,
						λ.NewStr("display_id"): ϒdisplay_id,
					}))
					return ϒinfo_dict
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       OnetIE_IE_NAME,
				λ.NewStr("_TEST"):         OnetIE__TEST,
				λ.NewStr("_VALID_URL"):    OnetIE__VALID_URL,
				λ.NewStr("_real_extract"): OnetIE__real_extract,
			})
		}())
		OnetChannelIE = λ.Cal(λ.TypeType, λ.NewStr("OnetChannelIE"), λ.NewTuple(OnetBaseIE), func() λ.Dict {
			var (
				OnetChannelIE_IE_NAME       λ.Object
				OnetChannelIE__TEST         λ.Object
				OnetChannelIE__VALID_URL    λ.Object
				OnetChannelIE__real_extract λ.Object
			)
			OnetChannelIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?onet\\.tv/[a-z]/(?P<id>[a-z]+)(?:[?#]|$)")
			OnetChannelIE_IE_NAME = λ.NewStr("onet.tv:channel")
			OnetChannelIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("http://onet.tv/k/openerfestival"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):          λ.NewStr("openerfestival"),
					λ.NewStr("title"):       λ.NewStr("Open'er Festival Live"),
					λ.NewStr("description"): λ.NewStr("Dziękujemy, że oglądaliście transmisje. Zobaczcie nasze relacje i wywiady z artystami."),
				}),
				λ.NewStr("playlist_mincount"): λ.NewInt(46),
			})
			OnetChannelIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒchannel_description λ.Object
						ϒchannel_id          λ.Object
						ϒchannel_title       λ.Object
						ϒcurrent_clip_info   λ.Object
						ϒentries             λ.Object
						ϒmatches             λ.Object
						ϒself                = λargs[0]
						ϒurl                 = λargs[1]
						ϒvideo_id            λ.Object
						ϒvideo_name          λ.Object
						ϒwebpage             λ.Object
					)
					ϒchannel_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒchannel_id)
					ϒcurrent_clip_info = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
						λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("var\\s+currentClip\\s*=\\s*({[^}]+})"), ϒwebpage, λ.NewStr("video info")),
						ϒchannel_id,
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
								return λ.Cal(ϒjs_to_json, λ.Cal(Ωre.ϒsub, λ.NewStr("\\'\\s*\\+\\s*\\'"), λ.NewStr(""), ϒs))
							})},
					})
					ϒvideo_id = λ.Cal(ϒremove_start, λ.GetItem(ϒcurrent_clip_info, λ.NewStr("ckmId")), λ.NewStr("mvp:"))
					ϒvideo_name = λ.Cal(ϒurl_basename, λ.GetItem(ϒcurrent_clip_info, λ.NewStr("url")))
					if λ.IsTrue(λ.Cal(λ.GetAttr(λ.GetAttr(λ.GetAttr(ϒself, "_downloader", nil), "params", nil), "get", nil), λ.NewStr("noplaylist"))) {
						λ.Cal(λ.GetAttr(ϒself, "to_screen", nil), λ.Mod(λ.NewStr("Downloading just video %s because of --no-playlist"), ϒvideo_name))
						return λ.Cal(λ.GetAttr(ϒself, "_extract_from_id", nil), ϒvideo_id, ϒwebpage)
					}
					λ.Cal(λ.GetAttr(ϒself, "to_screen", nil), λ.Mod(λ.NewStr("Downloading channel %s - add --no-playlist to just download video %s"), λ.NewTuple(
						ϒchannel_id,
						ϒvideo_name,
					)))
					ϒmatches = λ.Cal(Ωre.ϒfindall, λ.NewStr("<a[^>]+href=[\\'\"](https?://(?:www\\.)?onet\\.tv/[a-z]/[a-z]+/[0-9a-z-]+/[0-9a-z]+)"), ϒwebpage)
					ϒentries = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒvideo_link λ.Object
									τmp0        λ.Object
									τmp1        λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, ϒmatches)
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒvideo_link = τmp1
									λgen.Yield(λ.Cal(λ.GetAttr(ϒself, "url_result", nil), ϒvideo_link, λ.Cal(λ.GetAttr(OnetIE, "ie_key", nil))))
								}
								return λ.None
							})
						})))
					ϒchannel_title = λ.Cal(ϒstrip_or_none, λ.Cal(ϒget_element_by_class, λ.NewStr("o_channelName"), ϒwebpage))
					ϒchannel_description = λ.Cal(ϒstrip_or_none, λ.Cal(ϒget_element_by_class, λ.NewStr("o_channelDesc"), ϒwebpage))
					return λ.Cal(λ.GetAttr(ϒself, "playlist_result", nil), ϒentries, ϒchannel_id, ϒchannel_title, ϒchannel_description)
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       OnetChannelIE_IE_NAME,
				λ.NewStr("_TEST"):         OnetChannelIE__TEST,
				λ.NewStr("_VALID_URL"):    OnetChannelIE__VALID_URL,
				λ.NewStr("_real_extract"): OnetChannelIE__real_extract,
			})
		}())
		OnetPlIE = λ.Cal(λ.TypeType, λ.NewStr("OnetPlIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				OnetPlIE__VALID_URL λ.Object
			)
			OnetPlIE__VALID_URL = λ.NewStr("https?://(?:[^/]+\\.)?(?:onet|businessinsider\\.com|plejada)\\.pl/(?:[^/]+/)+(?P<id>[0-9a-z]+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): OnetPlIE__VALID_URL,
			})
		}())
	})
}