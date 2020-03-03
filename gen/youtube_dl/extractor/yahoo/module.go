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
 * yahoo/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/yahoo.py
 */

package yahoo

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωbrightcove "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/brightcove"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	BrightcoveNewIE     λ.Object
	InfoExtractor       λ.Object
	SearchInfoExtractor λ.Object
	YahooGyaOIE         λ.Object
	YahooGyaOPlayerIE   λ.Object
	YahooIE             λ.Object
	YahooJapanNewsIE    λ.Object
	ϒclean_html         λ.Object
	ϒcompat_str         λ.Object
	ϒint_or_none        λ.Object
	ϒmimetype2ext       λ.Object
	ϒparse_iso8601      λ.Object
	ϒsmuggle_url        λ.Object
	ϒtry_get            λ.Object
	ϒurl_or_none        λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		SearchInfoExtractor = Ωcommon.SearchInfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒclean_html = Ωutils.ϒclean_html
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒmimetype2ext = Ωutils.ϒmimetype2ext
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒsmuggle_url = Ωutils.ϒsmuggle_url
		ϒtry_get = Ωutils.ϒtry_get
		ϒurl_or_none = Ωutils.ϒurl_or_none
		BrightcoveNewIE = Ωbrightcove.BrightcoveNewIE
		YahooIE = λ.Cal(λ.TypeType, λ.StrLiteral("YahooIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				YahooIE__VALID_URL    λ.Object
				YahooIE__real_extract λ.Object
			)
			YahooIE__VALID_URL = λ.StrLiteral("(?P<url>https?://(?:(?P<country>[a-zA-Z]{2}(?:-[a-zA-Z]{2})?|malaysia)\\.)?(?:[\\da-zA-Z_-]+\\.)?yahoo\\.com/(?:[^/]+/)*(?P<id>[^?&#]*-[0-9]+(?:-[a-z]+)?)\\.html)")
			YahooIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒapi_base      λ.Object
						ϒcc            λ.Object
						ϒcc_url        λ.Object
						ϒcontent       λ.Object
						ϒcountry       λ.Object
						ϒcover         λ.Object
						ϒcover_url     λ.Object
						ϒdisplay_id    λ.Object
						ϒe             λ.Object
						ϒentries       λ.Object
						ϒfmt           λ.Object
						ϒfmts          λ.Object
						ϒformats       λ.Object
						ϒhost          λ.Object
						ϒi             λ.Object
						ϒiframe_url    λ.Object
						ϒis_live       λ.Object
						ϒitem          λ.Object
						ϒmedia_obj     λ.Object
						ϒmsg           λ.Object
						ϒpath          λ.Object
						ϒs             λ.Object
						ϒs_url         λ.Object
						ϒself          = λargs[0]
						ϒseries_info   λ.Object
						ϒstreaming_url λ.Object
						ϒsubtitles     λ.Object
						ϒtbr           λ.Object
						ϒthumb         λ.Object
						ϒthumb_url     λ.Object
						ϒthumbnails    λ.Object
						ϒtitle         λ.Object
						ϒurl           = λargs[1]
						ϒurls          λ.Object
						ϒuuid          λ.Object
						ϒvideo         λ.Object
						ϒvideo_id      λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
						τmp2           λ.Object
						τmp3           λ.Object
					)
					τmp0 = λ.Calm(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups")
					ϒurl = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒcountry = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒdisplay_id = λ.GetItem(τmp0, λ.IntLiteral(2))
					if !λ.IsTrue(ϒcountry) {
						ϒcountry = λ.StrLiteral("us")
					} else {
						ϒcountry = λ.GetItem(λ.Calm(ϒcountry, "split", λ.StrLiteral("-")), λ.IntLiteral(0))
					}
					ϒapi_base = λ.Mod(λ.StrLiteral("https://%s.yahoo.com/_td/api/resource/"), ϒcountry)
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.EnumerateIteratorType, λ.NewList(
						λ.Add(λ.StrLiteral("url="), ϒurl),
						λ.Add(λ.StrLiteral("ymedia-alias="), ϒdisplay_id),
					)))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒi = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒuuid = λ.GetItem(τmp2, λ.IntLiteral(1))
						ϒcontent = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
							λ.Add(ϒapi_base, λ.Mod(λ.StrLiteral("content;getDetailView=true;uuids=[\"%s\"]"), ϒuuid)),
							ϒdisplay_id,
							λ.StrLiteral("Downloading content JSON metadata"),
						), λ.KWArgs{
							{Name: "fatal", Value: λ.Eq(ϒi, λ.IntLiteral(1))},
						})
						if λ.IsTrue(ϒcontent) {
							ϒitem = λ.GetItem(λ.GetItem(ϒcontent, λ.StrLiteral("items")), λ.IntLiteral(0))
							break
						}
					}
					if λ.IsTrue(λ.Ne(λ.Calm(ϒitem, "get", λ.StrLiteral("type")), λ.StrLiteral("video"))) {
						ϒentries = λ.NewList()
						ϒcover = func() λ.Object {
							if λv := λ.Calm(ϒitem, "get", λ.StrLiteral("cover")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.DictLiteral(map[λ.Object]λ.Object{})
							}
						}()
						if λ.IsTrue(λ.Eq(λ.Calm(ϒcover, "get", λ.StrLiteral("type")), λ.StrLiteral("yvideo"))) {
							ϒcover_url = λ.Calm(ϒcover, "get", λ.StrLiteral("url"))
							if λ.IsTrue(ϒcover_url) {
								λ.Calm(ϒentries, "append", λ.Calm(ϒself, "url_result", ϒcover_url, λ.StrLiteral("Yahoo"), λ.Calm(ϒcover, "get", λ.StrLiteral("uuid"))))
							}
						}
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒitem, "get", λ.StrLiteral("body"), λ.NewList()))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒe = τmp1
							if λ.IsTrue(λ.Eq(λ.Calm(ϒe, "get", λ.StrLiteral("type")), λ.StrLiteral("videoIframe"))) {
								ϒiframe_url = λ.Calm(ϒe, "get", λ.StrLiteral("url"))
								if !λ.IsTrue(ϒiframe_url) {
									continue
								}
								λ.Calm(ϒentries, "append", λ.Calm(ϒself, "url_result", ϒiframe_url))
							}
						}
						return λ.Calm(ϒself, "playlist_result", ϒentries, λ.Calm(ϒitem, "get", λ.StrLiteral("uuid")), λ.Calm(ϒitem, "get", λ.StrLiteral("title")), λ.Calm(ϒitem, "get", λ.StrLiteral("summary")))
					}
					ϒvideo_id = λ.GetItem(ϒitem, λ.StrLiteral("uuid"))
					ϒvideo = λ.GetItem(λ.Calm(ϒself, "_download_json", λ.Add(ϒapi_base, λ.Mod(λ.StrLiteral("VideoService.videos;view=full;video_ids=[\"%s\"]"), ϒvideo_id)), ϒvideo_id, λ.StrLiteral("Downloading video JSON metadata")), λ.IntLiteral(0))
					ϒtitle = λ.GetItem(ϒvideo, λ.StrLiteral("title"))
					if λ.IsTrue(λ.Eq(ϒcountry, λ.StrLiteral("malaysia"))) {
						ϒcountry = λ.StrLiteral("my")
					}
					ϒis_live = λ.Eq(λ.Calm(ϒvideo, "get", λ.StrLiteral("live_state")), λ.StrLiteral("live"))
					ϒfmts = func() λ.Object {
						if λ.IsTrue(ϒis_live) {
							return λ.NewTuple(λ.StrLiteral("m3u8"))
						} else {
							return λ.NewTuple(
								λ.StrLiteral("webm"),
								λ.StrLiteral("mp4"),
							)
						}
					}()
					ϒurls = λ.NewList()
					ϒformats = λ.NewList()
					ϒsubtitles = λ.DictLiteral(map[λ.Object]λ.Object{})
					τmp0 = λ.Cal(λ.BuiltinIter, ϒfmts)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒfmt = τmp1
						ϒmedia_obj = λ.GetItem(λ.GetItem(λ.GetItem(λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
							λ.Add(λ.StrLiteral("https://video-api.yql.yahoo.com/v1/video/sapi/streams/"), ϒvideo_id),
							ϒvideo_id,
							λ.Mod(λ.StrLiteral("Downloading %s JSON metadata"), ϒfmt),
						), λ.KWArgs{
							{Name: "headers", Value: λ.Calm(ϒself, "geo_verification_headers")},
							{Name: "query", Value: λ.DictLiteral(map[string]λ.Object{
								"format": ϒfmt,
								"region": λ.Calm(ϒcountry, "upper"),
							})},
						}), λ.StrLiteral("query")), λ.StrLiteral("results")), λ.StrLiteral("mediaObj")), λ.IntLiteral(0))
						ϒmsg = λ.Calm(λ.Calm(ϒmedia_obj, "get", λ.StrLiteral("status"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("msg"))
						τmp2 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒmedia_obj, "get", λ.StrLiteral("streams"), λ.NewList()))
						for {
							if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
								break
							}
							ϒs = τmp3
							ϒhost = λ.Calm(ϒs, "get", λ.StrLiteral("host"))
							ϒpath = λ.Calm(ϒs, "get", λ.StrLiteral("path"))
							if λ.IsTrue(func() λ.Object {
								if λv := λ.NewBool(!λ.IsTrue(ϒhost)); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.NewBool(!λ.IsTrue(ϒpath))
								}
							}()) {
								continue
							}
							ϒs_url = λ.Add(ϒhost, ϒpath)
							if λ.IsTrue(λ.Eq(λ.Calm(ϒs, "get", λ.StrLiteral("format")), λ.StrLiteral("m3u8"))) {
								λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
									ϒs_url,
									ϒvideo_id,
									λ.StrLiteral("mp4"),
								), λ.KWArgs{
									{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
									{Name: "fatal", Value: λ.False},
								}))
								continue
							}
							ϒtbr = λ.Cal(ϒint_or_none, λ.Calm(ϒs, "get", λ.StrLiteral("bitrate")))
							λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
								"url": ϒs_url,
								"format_id": λ.Add(ϒfmt, func() λ.Object {
									if λ.IsTrue(ϒtbr) {
										return λ.Mod(λ.StrLiteral("-%d"), ϒtbr)
									} else {
										return λ.StrLiteral("")
									}
								}()),
								"width":  λ.Cal(ϒint_or_none, λ.Calm(ϒs, "get", λ.StrLiteral("width"))),
								"height": λ.Cal(ϒint_or_none, λ.Calm(ϒs, "get", λ.StrLiteral("height"))),
								"tbr":    ϒtbr,
								"fps":    λ.Cal(ϒint_or_none, λ.Calm(ϒs, "get", λ.StrLiteral("framerate"))),
							}))
						}
						τmp2 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒmedia_obj, "get", λ.StrLiteral("closedcaptions"), λ.NewList()))
						for {
							if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
								break
							}
							ϒcc = τmp3
							ϒcc_url = λ.Calm(ϒcc, "get", λ.StrLiteral("url"))
							if λ.IsTrue(func() λ.Object {
								if λv := λ.NewBool(!λ.IsTrue(ϒcc_url)); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.NewBool(λ.Contains(ϒurls, ϒcc_url))
								}
							}()) {
								continue
							}
							λ.Calm(ϒurls, "append", ϒcc_url)
							λ.Calm(λ.Calm(ϒsubtitles, "setdefault", func() λ.Object {
								if λv := λ.Calm(ϒcc, "get", λ.StrLiteral("lang")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.StrLiteral("en-US")
								}
							}(), λ.NewList()), "append", λ.DictLiteral(map[string]λ.Object{
								"url": ϒcc_url,
								"ext": λ.Cal(ϒmimetype2ext, λ.Calm(ϒcc, "get", λ.StrLiteral("content_type"))),
							}))
						}
					}
					ϒstreaming_url = λ.Calm(ϒvideo, "get", λ.StrLiteral("streaming_url"))
					if λ.IsTrue(func() λ.Object {
						if λv := ϒstreaming_url; !λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewBool(!λ.IsTrue(ϒis_live))
						}
					}()) {
						λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
							ϒstreaming_url,
							ϒvideo_id,
							λ.StrLiteral("mp4"),
							λ.StrLiteral("m3u8_native"),
						), λ.KWArgs{
							{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
							{Name: "fatal", Value: λ.False},
						}))
					}
					if λ.IsTrue(func() λ.Object {
						if λv := λ.NewBool(!λ.IsTrue(ϒformats)); !λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Eq(ϒmsg, λ.StrLiteral("geo restricted"))
						}
					}()) {
						λ.Calm(ϒself, "raise_geo_restricted")
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒthumbnails = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒvideo, "get", λ.StrLiteral("thumbnails"), λ.NewList()))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒthumb = τmp1
						ϒthumb_url = λ.Calm(ϒthumb, "get", λ.StrLiteral("url"))
						if !λ.IsTrue(ϒthumb_url) {
							continue
						}
						λ.Calm(ϒthumbnails, "append", λ.DictLiteral(map[string]λ.Object{
							"id":     λ.Calm(ϒthumb, "get", λ.StrLiteral("tag")),
							"url":    λ.Calm(ϒthumb, "get", λ.StrLiteral("url")),
							"width":  λ.Cal(ϒint_or_none, λ.Calm(ϒthumb, "get", λ.StrLiteral("width"))),
							"height": λ.Cal(ϒint_or_none, λ.Calm(ϒthumb, "get", λ.StrLiteral("height"))),
						}))
					}
					ϒseries_info = func() λ.Object {
						if λv := λ.Calm(ϒvideo, "get", λ.StrLiteral("series_info")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.DictLiteral(map[λ.Object]λ.Object{})
						}
					}()
					return λ.DictLiteral(map[string]λ.Object{
						"id": ϒvideo_id,
						"title": func() λ.Object {
							if λ.IsTrue(ϒis_live) {
								return λ.Calm(ϒself, "_live_title", ϒtitle)
							} else {
								return ϒtitle
							}
						}(),
						"formats":        ϒformats,
						"display_id":     ϒdisplay_id,
						"thumbnails":     ϒthumbnails,
						"description":    λ.Cal(ϒclean_html, λ.Calm(ϒvideo, "get", λ.StrLiteral("description"))),
						"timestamp":      λ.Cal(ϒparse_iso8601, λ.Calm(ϒvideo, "get", λ.StrLiteral("publish_time"))),
						"subtitles":      ϒsubtitles,
						"duration":       λ.Cal(ϒint_or_none, λ.Calm(ϒvideo, "get", λ.StrLiteral("duration"))),
						"view_count":     λ.Cal(ϒint_or_none, λ.Calm(ϒvideo, "get", λ.StrLiteral("view_count"))),
						"is_live":        ϒis_live,
						"series":         λ.Calm(ϒvideo, "get", λ.StrLiteral("show_name")),
						"season_number":  λ.Cal(ϒint_or_none, λ.Calm(ϒseries_info, "get", λ.StrLiteral("season_number"))),
						"episode_number": λ.Cal(ϒint_or_none, λ.Calm(ϒseries_info, "get", λ.StrLiteral("episode_number"))),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    YahooIE__VALID_URL,
				"_real_extract": YahooIE__real_extract,
			})
		}())
		YahooGyaOPlayerIE = λ.Cal(λ.TypeType, λ.StrLiteral("YahooGyaOPlayerIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				YahooGyaOPlayerIE__VALID_URL λ.Object
			)
			YahooGyaOPlayerIE__VALID_URL = λ.StrLiteral("https?://(?:gyao\\.yahoo\\.co\\.jp/(?:player|episode/[^/]+)|streaming\\.yahoo\\.co\\.jp/c/y)/(?P<id>\\d+/v\\d+/v\\d+|[\\da-f]{8}-[\\da-f]{4}-[\\da-f]{4}-[\\da-f]{4}-[\\da-f]{12})")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": YahooGyaOPlayerIE__VALID_URL,
			})
		}())
		YahooGyaOIE = λ.Cal(λ.TypeType, λ.StrLiteral("YahooGyaOIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				YahooGyaOIE__VALID_URL λ.Object
			)
			YahooGyaOIE__VALID_URL = λ.StrLiteral("https?://(?:gyao\\.yahoo\\.co\\.jp/(?:p|title/[^/]+)|streaming\\.yahoo\\.co\\.jp/p/y)/(?P<id>\\d+/v\\d+|[\\da-f]{8}-[\\da-f]{4}-[\\da-f]{4}-[\\da-f]{4}-[\\da-f]{12})")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": YahooGyaOIE__VALID_URL,
			})
		}())
		YahooJapanNewsIE = λ.Cal(λ.TypeType, λ.StrLiteral("YahooJapanNewsIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				YahooJapanNewsIE__VALID_URL λ.Object
			)
			YahooJapanNewsIE__VALID_URL = λ.StrLiteral("https?://(?P<host>(?:news|headlines)\\.yahoo\\.co\\.jp)[^\\d]*(?P<id>\\d[\\d-]*\\d)?")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": YahooJapanNewsIE__VALID_URL,
			})
		}())
	})
}
