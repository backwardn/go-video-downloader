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
 * piksel/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/piksel.py
 */

package piksel

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError λ.Object
	InfoExtractor  λ.Object
	PikselIE       λ.Object
	ϒcompat_str    λ.Object
	ϒdict_get      λ.Object
	ϒint_or_none   λ.Object
	ϒparse_iso8601 λ.Object
	ϒunescapeHTML  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ExtractorError = Ωutils.ExtractorError
		ϒdict_get = Ωutils.ϒdict_get
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒunescapeHTML = Ωutils.ϒunescapeHTML
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		PikselIE = λ.Cal(λ.TypeType, λ.StrLiteral("PikselIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				PikselIE__VALID_URL    λ.Object
				PikselIE__real_extract λ.Object
			)
			PikselIE__VALID_URL = λ.StrLiteral("https?://player\\.piksel\\.com/v/(?:refid/[^/]+/prefid/)?(?P<id>[a-z0-9_]+)")
			PikselIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒabr         λ.Object
						ϒapp_token   λ.Object
						ϒasset_file  λ.Object
						ϒasset_type  λ.Object
						ϒcaption     λ.Object
						ϒcaption_url λ.Object
						ϒdisplay_id  λ.Object
						ϒfailure     λ.Object
						ϒformat_id   λ.Object
						ϒformats     λ.Object
						ϒhttp_url    λ.Object
						ϒm3u8_url    λ.Object
						ϒresponse    λ.Object
						ϒself        = λargs[0]
						ϒsubtitles   λ.Object
						ϒtbr         λ.Object
						ϒtitle       λ.Object
						ϒurl         = λargs[1]
						ϒvbr         λ.Object
						ϒvideo_data  λ.Object
						ϒvideo_id    λ.Object
						ϒwebpage     λ.Object
						τmp0         λ.Object
						τmp1         λ.Object
					)
					ϒdisplay_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒdisplay_id)
					ϒvideo_id = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("data-de-program-uuid=[\\'\"]([a-z0-9]+)"),
						ϒwebpage,
						λ.StrLiteral("program uuid"),
					), λ.KWArgs{
						{Name: "default", Value: ϒdisplay_id},
					})
					ϒapp_token = λ.Calm(ϒself, "_search_regex", λ.NewList(
						λ.StrLiteral("clientAPI\\s*:\\s*\"([^\"]+)\""),
						λ.StrLiteral("data-de-api-key\\s*=\\s*\"([^\"]+)\""),
					), ϒwebpage, λ.StrLiteral("app token"))
					ϒresponse = λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.Mod(λ.StrLiteral("http://player.piksel.com/ws/ws_program/api/%s/mode/json/apiv/5"), ϒapp_token),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "query", Value: λ.DictLiteral(map[string]λ.Object{
							"v": ϒvideo_id,
						})},
					}), λ.StrLiteral("response"))
					ϒfailure = λ.Calm(ϒresponse, "get", λ.StrLiteral("failure"))
					if λ.IsTrue(ϒfailure) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.GetItem(λ.GetItem(ϒresponse, λ.StrLiteral("failure")), λ.StrLiteral("reason"))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒvideo_data = λ.GetItem(λ.GetItem(λ.GetItem(ϒresponse, λ.StrLiteral("WsProgramResponse")), λ.StrLiteral("program")), λ.StrLiteral("asset"))
					ϒtitle = λ.GetItem(ϒvideo_data, λ.StrLiteral("title"))
					ϒformats = λ.NewList()
					ϒm3u8_url = λ.Cal(ϒdict_get, ϒvideo_data, λ.NewList(
						λ.StrLiteral("m3u8iPadURL"),
						λ.StrLiteral("ipadM3u8Url"),
						λ.StrLiteral("m3u8AndroidURL"),
						λ.StrLiteral("m3u8iPhoneURL"),
						λ.StrLiteral("iphoneM3u8Url"),
					))
					if λ.IsTrue(ϒm3u8_url) {
						λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
							ϒm3u8_url,
							ϒvideo_id,
							λ.StrLiteral("mp4"),
							λ.StrLiteral("m3u8_native"),
						), λ.KWArgs{
							{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
							{Name: "fatal", Value: λ.False},
						}))
					}
					ϒasset_type = λ.Cal(ϒdict_get, ϒvideo_data, λ.NewList(
						λ.StrLiteral("assetType"),
						λ.StrLiteral("asset_type"),
					))
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒvideo_data, "get", λ.StrLiteral("assetFiles"), λ.NewList()))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒasset_file = τmp1
						ϒhttp_url = λ.Calm(ϒasset_file, "get", λ.StrLiteral("http_url"))
						if !λ.IsTrue(ϒhttp_url) {
							continue
						}
						ϒtbr = λ.None
						ϒvbr = λ.Cal(ϒint_or_none, λ.Calm(ϒasset_file, "get", λ.StrLiteral("videoBitrate")), λ.IntLiteral(1024))
						ϒabr = λ.Cal(ϒint_or_none, λ.Calm(ϒasset_file, "get", λ.StrLiteral("audioBitrate")), λ.IntLiteral(1024))
						if λ.IsTrue(λ.Eq(ϒasset_type, λ.StrLiteral("video"))) {
							ϒtbr = λ.Add(ϒvbr, ϒabr)
						} else {
							if λ.IsTrue(λ.Eq(ϒasset_type, λ.StrLiteral("audio"))) {
								ϒtbr = ϒabr
							}
						}
						ϒformat_id = λ.NewList(λ.StrLiteral("http"))
						if λ.IsTrue(ϒtbr) {
							λ.Calm(ϒformat_id, "append", λ.Cal(ϒcompat_str, ϒtbr))
						}
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"format_id": λ.Calm(λ.StrLiteral("-"), "join", ϒformat_id),
							"url":       λ.Cal(ϒunescapeHTML, ϒhttp_url),
							"vbr":       ϒvbr,
							"abr":       ϒabr,
							"width":     λ.Cal(ϒint_or_none, λ.Calm(ϒasset_file, "get", λ.StrLiteral("videoWidth"))),
							"height":    λ.Cal(ϒint_or_none, λ.Calm(ϒasset_file, "get", λ.StrLiteral("videoHeight"))),
							"filesize":  λ.Cal(ϒint_or_none, λ.Calm(ϒasset_file, "get", λ.StrLiteral("filesize"))),
							"tbr":       ϒtbr,
						}))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒsubtitles = λ.DictLiteral(map[λ.Object]λ.Object{})
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒvideo_data, "get", λ.StrLiteral("captions"), λ.NewList()))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒcaption = τmp1
						ϒcaption_url = λ.Calm(ϒcaption, "get", λ.StrLiteral("url"))
						if λ.IsTrue(ϒcaption_url) {
							λ.Calm(λ.Calm(ϒsubtitles, "setdefault", λ.Calm(ϒcaption, "get", λ.StrLiteral("locale"), λ.StrLiteral("en")), λ.NewList()), "append", λ.DictLiteral(map[string]λ.Object{
								"url": ϒcaption_url,
							}))
						}
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id":          ϒvideo_id,
						"title":       ϒtitle,
						"description": λ.Calm(ϒvideo_data, "get", λ.StrLiteral("description")),
						"thumbnail":   λ.Calm(ϒvideo_data, "get", λ.StrLiteral("thumbnailUrl")),
						"timestamp":   λ.Cal(ϒparse_iso8601, λ.Calm(ϒvideo_data, "get", λ.StrLiteral("dateadd"))),
						"formats":     ϒformats,
						"subtitles":   ϒsubtitles,
					})
				})
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL":    PikselIE__VALID_URL,
				"_real_extract": PikselIE__real_extract,
			})
		}())
	})
}
