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
 * limelight/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/limelight.py
 */

package limelight

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError         λ.Object
	InfoExtractor          λ.Object
	LimelightBaseIE        λ.Object
	LimelightChannelIE     λ.Object
	LimelightChannelListIE λ.Object
	LimelightMediaIE       λ.Object
	ϒcompat_HTTPError      λ.Object
	ϒdetermine_ext         λ.Object
	ϒfloat_or_none         λ.Object
	ϒint_or_none           λ.Object
	ϒsmuggle_url           λ.Object
	ϒtry_get               λ.Object
	ϒunsmuggle_url         λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_HTTPError = Ωcompat.ϒcompat_HTTPError
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒsmuggle_url = Ωutils.ϒsmuggle_url
		ϒtry_get = Ωutils.ϒtry_get
		ϒunsmuggle_url = Ωutils.ϒunsmuggle_url
		ExtractorError = Ωutils.ExtractorError
		LimelightBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("LimelightBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				LimelightBaseIE__PLAYLIST_SERVICE_URL  λ.Object
				LimelightBaseIE__call_playlist_service λ.Object
				LimelightBaseIE__extract               λ.Object
				LimelightBaseIE__extract_info          λ.Object
			)
			LimelightBaseIE__PLAYLIST_SERVICE_URL = λ.StrLiteral("http://production-ps.lvp.llnw.net/r/PlaylistService/%s/%s/%s")
			LimelightBaseIE__call_playlist_service = λ.NewFunction("_call_playlist_service",
				[]λ.Param{
					{Name: "self"},
					{Name: "item_id"},
					{Name: "method"},
					{Name: "fatal", Def: λ.True},
					{Name: "referer", Def: λ.None},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒerror   λ.Object
						ϒfatal   = λargs[3]
						ϒheaders λ.Object
						ϒitem_id = λargs[1]
						ϒmethod  = λargs[2]
						ϒreferer = λargs[4]
						ϒself    = λargs[0]
						τmp0     λ.Object
						τmp1     λ.Object
					)
					ϒheaders = λ.DictLiteral(map[λ.Object]λ.Object{})
					if λ.IsTrue(ϒreferer) {
						λ.SetItem(ϒheaders, λ.StrLiteral("Referer"), ϒreferer)
					}
					τmp0, τmp1 = func() (λexit λ.Object, λret λ.Object) {
						defer λ.CatchMulti(
							nil,
							&λ.Catcher{ExtractorError, func(λex λ.BaseException) {
								var ϒe λ.Object = λex
								if λ.IsTrue(func() λ.Object {
									if λv := λ.Cal(λ.BuiltinIsInstance, λ.GetAttr(ϒe, "cause", nil), ϒcompat_HTTPError); !λ.IsTrue(λv) {
										return λv
									} else {
										return λ.Eq(λ.GetAttr(λ.GetAttr(ϒe, "cause", nil), "code", nil), λ.IntLiteral(403))
									}
								}()) {
									ϒerror = λ.GetItem(λ.GetItem(λ.Calm(ϒself, "_parse_json", λ.Calm(λ.Calm(λ.GetAttr(ϒe, "cause", nil), "read"), "decode"), ϒitem_id), λ.StrLiteral("detail")), λ.StrLiteral("contentAccessPermission"))
									if λ.IsTrue(λ.Eq(ϒerror, λ.StrLiteral("CountryDisabled"))) {
										λ.Calm(ϒself, "raise_geo_restricted")
									}
									panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(ϒerror), λ.KWArgs{
										{Name: "expected", Value: λ.True},
									})))
								}
								panic(λ.Raise(λex))
							}},
						)
						λexit, λret = λ.BlockExitReturn, λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
							λ.Mod(λ.GetAttr(ϒself, "_PLAYLIST_SERVICE_URL", nil), λ.NewTuple(
								λ.GetAttr(ϒself, "_PLAYLIST_SERVICE_PATH", nil),
								ϒitem_id,
								ϒmethod,
							)),
							ϒitem_id,
							λ.Mod(λ.StrLiteral("Downloading PlaylistService %s JSON"), ϒmethod),
						), λ.KWArgs{
							{Name: "fatal", Value: ϒfatal},
							{Name: "headers", Value: ϒheaders},
						})
						return
						return λ.BlockExitNormally, nil
					}()
					if τmp0 == λ.BlockExitReturn {
						return τmp1
					}
					return λ.None
				})
			LimelightBaseIE__extract = λ.NewFunction("_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "item_id"},
					{Name: "pc_method"},
					{Name: "mobile_method"},
					{Name: "referer", Def: λ.None},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒitem_id       = λargs[1]
						ϒmobile        λ.Object
						ϒmobile_method = λargs[3]
						ϒpc            λ.Object
						ϒpc_method     = λargs[2]
						ϒreferer       = λargs[4]
						ϒself          = λargs[0]
					)
					ϒpc = λ.Call(λ.GetAttr(ϒself, "_call_playlist_service", nil), λ.NewArgs(
						ϒitem_id,
						ϒpc_method,
					), λ.KWArgs{
						{Name: "referer", Value: ϒreferer},
					})
					ϒmobile = λ.Call(λ.GetAttr(ϒself, "_call_playlist_service", nil), λ.NewArgs(
						ϒitem_id,
						ϒmobile_method,
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
						{Name: "referer", Value: ϒreferer},
					})
					return λ.NewTuple(
						ϒpc,
						ϒmobile,
					)
				})
			LimelightBaseIE__extract_info = λ.NewFunction("_extract_info",
				[]λ.Param{
					{Name: "self"},
					{Name: "pc"},
					{Name: "mobile"},
					{Name: "i"},
					{Name: "referer"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						CDN_HOSTS        λ.Object
						ϒcc              λ.Object
						ϒcc_url          λ.Object
						ϒcdn_host        λ.Object
						ϒclosed_captions λ.Object
						ϒext             λ.Object
						ϒflag            λ.Object
						ϒfmt             λ.Object
						ϒformat_id       λ.Object
						ϒformats         λ.Object
						ϒget_item        λ.Object
						ϒget_meta        λ.Object
						ϒheight          λ.Object
						ϒhttp_fmt        λ.Object
						ϒhttp_format_id  λ.Object
						ϒhttp_host       λ.Object
						ϒhttp_url        λ.Object
						ϒi               = λargs[3]
						ϒlang            λ.Object
						ϒmedia_url       λ.Object
						ϒmobile          = λargs[2]
						ϒmobile_item     λ.Object
						ϒmobile_url      λ.Object
						ϒpc              = λargs[1]
						ϒpc_item         λ.Object
						ϒreferer         = λargs[4]
						ϒrtmp            λ.Object
						ϒself            = λargs[0]
						ϒstream          λ.Object
						ϒstream_url      λ.Object
						ϒsubtitles       λ.Object
						ϒtitle           λ.Object
						ϒurls            λ.Object
						ϒvbr             λ.Object
						ϒvideo_id        λ.Object
						ϒwidth           λ.Object
						τmp0             λ.Object
						τmp1             λ.Object
						τmp2             λ.Object
						τmp3             λ.Object
						τmp4             λ.Object
					)
					ϒget_item = λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "x"},
							{Name: "y"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒx = λargs[0]
								ϒy = λargs[1]
							)
							return func() λ.Object {
								if λv := λ.Cal(ϒtry_get, ϒx, λ.NewFunction("<lambda>",
									[]λ.Param{
										{Name: "x"},
									},
									0, false, false,
									func(λargs []λ.Object) λ.Object {
										var (
											ϒx = λargs[0]
										)
										return λ.GetItem(λ.GetItem(ϒx, ϒy), ϒi)
									}), λ.DictType); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.DictLiteral(map[λ.Object]λ.Object{})
								}
							}()
						})
					ϒpc_item = λ.Cal(ϒget_item, ϒpc, λ.StrLiteral("playlistItems"))
					ϒmobile_item = λ.Cal(ϒget_item, ϒmobile, λ.StrLiteral("mediaList"))
					ϒvideo_id = func() λ.Object {
						if λv := λ.Calm(ϒpc_item, "get", λ.StrLiteral("mediaId")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.GetItem(ϒmobile_item, λ.StrLiteral("mediaId"))
						}
					}()
					ϒtitle = func() λ.Object {
						if λv := λ.Calm(ϒpc_item, "get", λ.StrLiteral("title")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.GetItem(ϒmobile_item, λ.StrLiteral("title"))
						}
					}()
					ϒformats = λ.NewList()
					ϒurls = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒpc_item, "get", λ.StrLiteral("streams"), λ.NewList()))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒstream = τmp1
						ϒstream_url = λ.Calm(ϒstream, "get", λ.StrLiteral("url"))
						if λ.IsTrue(func() λ.Object {
							if λv := λ.NewBool(!λ.IsTrue(ϒstream_url)); λ.IsTrue(λv) {
								return λv
							} else if λv := λ.Calm(ϒstream, "get", λ.StrLiteral("drmProtected")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewBool(λ.Contains(ϒurls, ϒstream_url))
							}
						}()) {
							continue
						}
						λ.Calm(ϒurls, "append", ϒstream_url)
						ϒext = λ.Cal(ϒdetermine_ext, ϒstream_url)
						if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("f4m"))) {
							λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
								ϒstream_url,
								ϒvideo_id,
							), λ.KWArgs{
								{Name: "f4m_id", Value: λ.StrLiteral("hds")},
								{Name: "fatal", Value: λ.False},
							}))
						} else {
							ϒfmt = λ.DictLiteral(map[string]λ.Object{
								"url": ϒstream_url,
								"abr": λ.Cal(ϒfloat_or_none, λ.Calm(ϒstream, "get", λ.StrLiteral("audioBitRate"))),
								"fps": λ.Cal(ϒfloat_or_none, λ.Calm(ϒstream, "get", λ.StrLiteral("videoFrameRate"))),
								"ext": ϒext,
							})
							ϒwidth = λ.Cal(ϒint_or_none, λ.Calm(ϒstream, "get", λ.StrLiteral("videoWidthInPixels")))
							ϒheight = λ.Cal(ϒint_or_none, λ.Calm(ϒstream, "get", λ.StrLiteral("videoHeightInPixels")))
							ϒvbr = λ.Cal(ϒfloat_or_none, λ.Calm(ϒstream, "get", λ.StrLiteral("videoBitRate")))
							if λ.IsTrue(func() λ.Object {
								if λv := ϒwidth; λ.IsTrue(λv) {
									return λv
								} else if λv := ϒheight; λ.IsTrue(λv) {
									return λv
								} else {
									return ϒvbr
								}
							}()) {
								λ.Calm(ϒfmt, "update", λ.DictLiteral(map[string]λ.Object{
									"width":  ϒwidth,
									"height": ϒheight,
									"vbr":    ϒvbr,
								}))
							} else {
								λ.SetItem(ϒfmt, λ.StrLiteral("vcodec"), λ.StrLiteral("none"))
							}
							ϒrtmp = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("^(?P<url>rtmpe?://(?P<host>[^/]+)/(?P<app>.+))/(?P<playpath>mp[34]:.+)$"), ϒstream_url)
							if λ.IsTrue(ϒrtmp) {
								ϒformat_id = λ.StrLiteral("rtmp")
								if λ.IsTrue(λ.Calm(ϒstream, "get", λ.StrLiteral("videoBitRate"))) {
									τmp2 = λ.IAdd(ϒformat_id, λ.Mod(λ.StrLiteral("-%d"), λ.Cal(ϒint_or_none, λ.GetItem(ϒstream, λ.StrLiteral("videoBitRate")))))
									ϒformat_id = τmp2
								}
								ϒhttp_format_id = λ.Calm(ϒformat_id, "replace", λ.StrLiteral("rtmp"), λ.StrLiteral("http"))
								CDN_HOSTS = λ.NewTuple(
									λ.NewTuple(
										λ.StrLiteral("delvenetworks.com"),
										λ.StrLiteral("cpl.delvenetworks.com"),
									),
									λ.NewTuple(
										λ.StrLiteral("video.llnw.net"),
										λ.StrLiteral("s2.content.video.llnw.net"),
									),
								)
								τmp2 = λ.Cal(λ.BuiltinIter, CDN_HOSTS)
								for {
									if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
										break
									}
									τmp4 = τmp3
									ϒcdn_host = λ.GetItem(τmp4, λ.IntLiteral(0))
									ϒhttp_host = λ.GetItem(τmp4, λ.IntLiteral(1))
									if !λ.Contains(λ.Calm(λ.Calm(ϒrtmp, "group", λ.StrLiteral("host")), "lower"), ϒcdn_host) {
										continue
									}
									ϒhttp_url = λ.Mod(λ.StrLiteral("http://%s/%s"), λ.NewTuple(
										ϒhttp_host,
										λ.GetItem(λ.Calm(ϒrtmp, "group", λ.StrLiteral("playpath")), λ.NewSlice(λ.IntLiteral(4), λ.None, λ.None)),
									))
									λ.Calm(ϒurls, "append", ϒhttp_url)
									if λ.IsTrue(λ.Calm(ϒself, "_is_valid_url", ϒhttp_url, ϒvideo_id, ϒhttp_format_id)) {
										ϒhttp_fmt = λ.Calm(ϒfmt, "copy")
										λ.Calm(ϒhttp_fmt, "update", λ.DictLiteral(map[string]λ.Object{
											"url":       ϒhttp_url,
											"format_id": ϒhttp_format_id,
										}))
										λ.Calm(ϒformats, "append", ϒhttp_fmt)
										break
									}
								}
								λ.Calm(ϒfmt, "update", λ.DictLiteral(map[string]λ.Object{
									"url":       λ.Calm(ϒrtmp, "group", λ.StrLiteral("url")),
									"play_path": λ.Calm(ϒrtmp, "group", λ.StrLiteral("playpath")),
									"app":       λ.Calm(ϒrtmp, "group", λ.StrLiteral("app")),
									"ext":       λ.StrLiteral("flv"),
									"format_id": ϒformat_id,
								}))
							}
							λ.Calm(ϒformats, "append", ϒfmt)
						}
					}
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒmobile_item, "get", λ.StrLiteral("mobileUrls"), λ.NewList()))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒmobile_url = τmp1
						ϒmedia_url = λ.Calm(ϒmobile_url, "get", λ.StrLiteral("mobileUrl"))
						ϒformat_id = λ.Calm(ϒmobile_url, "get", λ.StrLiteral("targetMediaPlatform"))
						if λ.IsTrue(func() λ.Object {
							if λv := λ.NewBool(!λ.IsTrue(ϒmedia_url)); λ.IsTrue(λv) {
								return λv
							} else if λv := λ.NewBool(λ.Contains(λ.NewTuple(
								λ.StrLiteral("Widevine"),
								λ.StrLiteral("SmoothStreaming"),
							), ϒformat_id)); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewBool(λ.Contains(ϒurls, ϒmedia_url))
							}
						}()) {
							continue
						}
						λ.Calm(ϒurls, "append", ϒmedia_url)
						ϒext = λ.Cal(ϒdetermine_ext, ϒmedia_url)
						if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("m3u8"))) {
							λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
								ϒmedia_url,
								ϒvideo_id,
								λ.StrLiteral("mp4"),
								λ.StrLiteral("m3u8_native"),
							), λ.KWArgs{
								{Name: "m3u8_id", Value: ϒformat_id},
								{Name: "fatal", Value: λ.False},
							}))
						} else {
							if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("f4m"))) {
								λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
									ϒstream_url,
									ϒvideo_id,
								), λ.KWArgs{
									{Name: "f4m_id", Value: ϒformat_id},
									{Name: "fatal", Value: λ.False},
								}))
							} else {
								λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
									"url":        ϒmedia_url,
									"format_id":  ϒformat_id,
									"preference": λ.Neg(λ.IntLiteral(1)),
									"ext":        ϒext,
								}))
							}
						}
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒsubtitles = λ.DictLiteral(map[λ.Object]λ.Object{})
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒmobile_item, "get", λ.StrLiteral("flags")))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒflag = τmp1
						if λ.IsTrue(λ.Eq(ϒflag, λ.StrLiteral("ClosedCaptions"))) {
							ϒclosed_captions = func() λ.Object {
								if λv := λ.Calm(ϒself, "_call_playlist_service", ϒvideo_id, λ.StrLiteral("getClosedCaptionsDetailsByMediaId"), λ.False, ϒreferer); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.NewList()
								}
							}()
							τmp2 = λ.Cal(λ.BuiltinIter, ϒclosed_captions)
							for {
								if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
									break
								}
								ϒcc = τmp3
								ϒcc_url = λ.Calm(ϒcc, "get", λ.StrLiteral("webvttFileUrl"))
								if !λ.IsTrue(ϒcc_url) {
									continue
								}
								ϒlang = func() λ.Object {
									if λv := λ.Calm(ϒcc, "get", λ.StrLiteral("languageCode")); λ.IsTrue(λv) {
										return λv
									} else {
										return λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
											λ.StrLiteral("/[a-z]{2}\\.vtt"),
											ϒcc_url,
											λ.StrLiteral("lang"),
										), λ.KWArgs{
											{Name: "default", Value: λ.StrLiteral("en")},
										})
									}
								}()
								λ.Calm(λ.Calm(ϒsubtitles, "setdefault", ϒlang, λ.NewList()), "append", λ.DictLiteral(map[string]λ.Object{
									"url": ϒcc_url,
								}))
							}
							break
						}
					}
					ϒget_meta = λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "x"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒx = λargs[0]
							)
							return func() λ.Object {
								if λv := λ.Calm(ϒpc_item, "get", ϒx); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Calm(ϒmobile_item, "get", ϒx)
								}
							}()
						})
					return λ.DictLiteral(map[string]λ.Object{
						"id":          ϒvideo_id,
						"title":       ϒtitle,
						"description": λ.Cal(ϒget_meta, λ.StrLiteral("description")),
						"formats":     ϒformats,
						"duration":    λ.Cal(ϒfloat_or_none, λ.Cal(ϒget_meta, λ.StrLiteral("durationInMilliseconds")), λ.IntLiteral(1000)),
						"thumbnail": func() λ.Object {
							if λv := λ.Cal(ϒget_meta, λ.StrLiteral("previewImageUrl")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(ϒget_meta, λ.StrLiteral("thumbnailImageUrl"))
							}
						}(),
						"subtitles": ϒsubtitles,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_PLAYLIST_SERVICE_URL":  LimelightBaseIE__PLAYLIST_SERVICE_URL,
				"_call_playlist_service": LimelightBaseIE__call_playlist_service,
				"_extract":               LimelightBaseIE__extract,
				"_extract_info":          LimelightBaseIE__extract_info,
			})
		}())
		LimelightMediaIE = λ.Cal(λ.TypeType, λ.StrLiteral("LimelightMediaIE"), λ.NewTuple(LimelightBaseIE), func() λ.Dict {
			var (
				LimelightMediaIE_IE_NAME                λ.Object
				LimelightMediaIE__PLAYLIST_SERVICE_PATH λ.Object
				LimelightMediaIE__VALID_URL             λ.Object
				LimelightMediaIE__real_extract          λ.Object
			)
			LimelightMediaIE_IE_NAME = λ.StrLiteral("limelight")
			LimelightMediaIE__VALID_URL = λ.StrLiteral("(?x)\n                        (?:\n                            limelight:media:|\n                            https?://\n                                (?:\n                                    link\\.videoplatform\\.limelight\\.com/media/|\n                                    assets\\.delvenetworks\\.com/player/loader\\.swf\n                                )\n                                \\?.*?\\bmediaId=\n                        )\n                        (?P<id>[a-z0-9]{32})\n                    ")
			LimelightMediaIE__PLAYLIST_SERVICE_PATH = λ.StrLiteral("media")
			LimelightMediaIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒmobile        λ.Object
						ϒpc            λ.Object
						ϒself          = λargs[0]
						ϒsmuggled_data λ.Object
						ϒsource_url    λ.Object
						ϒurl           = λargs[1]
						ϒvideo_id      λ.Object
						τmp0           λ.Object
					)
					τmp0 = λ.Cal(ϒunsmuggle_url, ϒurl, λ.DictLiteral(map[λ.Object]λ.Object{}))
					ϒurl = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒsmuggled_data = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒsource_url = λ.Calm(ϒsmuggled_data, "get", λ.StrLiteral("source_url"))
					λ.Calm(ϒself, "_initialize_geo_bypass", λ.DictLiteral(map[string]λ.Object{
						"countries": λ.Calm(ϒsmuggled_data, "get", λ.StrLiteral("geo_countries")),
					}))
					τmp0 = λ.Calm(ϒself, "_extract", ϒvideo_id, λ.StrLiteral("getPlaylistByMediaId"), λ.StrLiteral("getMobilePlaylistByMediaId"), ϒsource_url)
					ϒpc = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒmobile = λ.GetItem(τmp0, λ.IntLiteral(1))
					return λ.Calm(ϒself, "_extract_info", ϒpc, ϒmobile, λ.IntLiteral(0), ϒsource_url)
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":                LimelightMediaIE_IE_NAME,
				"_PLAYLIST_SERVICE_PATH": LimelightMediaIE__PLAYLIST_SERVICE_PATH,
				"_VALID_URL":             LimelightMediaIE__VALID_URL,
				"_real_extract":          LimelightMediaIE__real_extract,
			})
		}())
		LimelightChannelIE = λ.Cal(λ.TypeType, λ.StrLiteral("LimelightChannelIE"), λ.NewTuple(LimelightBaseIE), func() λ.Dict {
			var (
				LimelightChannelIE__VALID_URL λ.Object
			)
			LimelightChannelIE__VALID_URL = λ.StrLiteral("(?x)\n                        (?:\n                            limelight:channel:|\n                            https?://\n                                (?:\n                                    link\\.videoplatform\\.limelight\\.com/media/|\n                                    assets\\.delvenetworks\\.com/player/loader\\.swf\n                                )\n                                \\?.*?\\bchannelId=\n                        )\n                        (?P<id>[a-z0-9]{32})\n                    ")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": LimelightChannelIE__VALID_URL,
			})
		}())
		LimelightChannelListIE = λ.Cal(λ.TypeType, λ.StrLiteral("LimelightChannelListIE"), λ.NewTuple(LimelightBaseIE), func() λ.Dict {
			var (
				LimelightChannelListIE__VALID_URL λ.Object
			)
			LimelightChannelListIE__VALID_URL = λ.StrLiteral("(?x)\n                        (?:\n                            limelight:channel_list:|\n                            https?://\n                                (?:\n                                    link\\.videoplatform\\.limelight\\.com/media/|\n                                    assets\\.delvenetworks\\.com/player/loader\\.swf\n                                )\n                                \\?.*?\\bchannelListId=\n                        )\n                        (?P<id>[a-z0-9]{32})\n                    ")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": LimelightChannelListIE__VALID_URL,
			})
		}())
	})
}
