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
 * vuclip/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/vuclip.py
 */

package vuclip

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError                λ.Object
	InfoExtractor                 λ.Object
	VuClipIE                      λ.Object
	ϒcompat_urllib_parse_urlparse λ.Object
	ϒparse_duration               λ.Object
	ϒremove_end                   λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_urllib_parse_urlparse = Ωcompat.ϒcompat_urllib_parse_urlparse
		ExtractorError = Ωutils.ExtractorError
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒremove_end = Ωutils.ϒremove_end
		VuClipIE = λ.Cal(λ.TypeType, λ.StrLiteral("VuClipIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				VuClipIE__VALID_URL    λ.Object
				VuClipIE__real_extract λ.Object
			)
			VuClipIE__VALID_URL = λ.StrLiteral("https?://(?:m\\.)?vuclip\\.com/w\\?.*?cid=(?P<id>[0-9]+)")
			VuClipIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒad_m       λ.Object
						ϒadfree_url λ.Object
						ϒduration   λ.Object
						ϒerror_msg  λ.Object
						ϒformats    λ.Object
						ϒself       = λargs[0]
						ϒtitle      λ.Object
						ϒurl        = λargs[1]
						ϒurlr       λ.Object
						ϒvideo_id   λ.Object
						ϒvideo_url  λ.Object
						ϒwebpage    λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒvideo_id)
					ϒad_m = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("value=\"No.*?\" onClick=\"location.href='([^\"']+)'\""), ϒwebpage)
					if λ.IsTrue(ϒad_m) {
						ϒurlr = λ.Cal(ϒcompat_urllib_parse_urlparse, ϒurl)
						ϒadfree_url = λ.Add(λ.Add(λ.Add(λ.GetAttr(ϒurlr, "scheme", nil), λ.StrLiteral("://")), λ.GetAttr(ϒurlr, "netloc", nil)), λ.Calm(ϒad_m, "group", λ.IntLiteral(1)))
						ϒwebpage = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
							ϒadfree_url,
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "note", Value: λ.StrLiteral("Download post-ad page")},
						})
					}
					ϒerror_msg = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("<p class=\"message\">(.*?)</p>"),
						ϒwebpage,
						λ.StrLiteral("error message"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒerror_msg) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.StrLiteral("%s said: %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							ϒerror_msg,
						))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒvideo_url = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("<a[^>]+href=\"([^\"]+)\"[^>]*><img[^>]+src=\"[^\"]*/play\\.gif"),
						ϒwebpage,
						λ.StrLiteral("video URL"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒvideo_url) {
						ϒformats = λ.NewList(λ.DictLiteral(map[string]λ.Object{
							"url": ϒvideo_url,
						}))
					} else {
						ϒformats = λ.GetItem(λ.GetItem(λ.Calm(ϒself, "_parse_html5_media_entries", ϒurl, ϒwebpage, ϒvideo_id), λ.IntLiteral(0)), λ.StrLiteral("formats"))
					}
					ϒtitle = λ.Cal(ϒremove_end, λ.Calm(λ.Calm(ϒself, "_html_search_regex", λ.StrLiteral("<title>(.*?)-\\s*Vuclip</title>"), ϒwebpage, λ.StrLiteral("title")), "strip"), λ.StrLiteral(" - Video"))
					ϒduration = λ.Cal(ϒparse_duration, λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("[(>]([0-9]+:[0-9]+)(?:<span|\\))"),
						ϒwebpage,
						λ.StrLiteral("duration"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					return λ.DictLiteral(map[string]λ.Object{
						"id":       ϒvideo_id,
						"formats":  ϒformats,
						"title":    ϒtitle,
						"duration": ϒduration,
					})
				})
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL":    VuClipIE__VALID_URL,
				"_real_extract": VuClipIE__real_extract,
			})
		}())
	})
}
