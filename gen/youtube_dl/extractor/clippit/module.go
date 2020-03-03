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
 * clippit/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/clippit.py
 */

package clippit

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ClippitIE      λ.Object
	InfoExtractor  λ.Object
	ϒparse_iso8601 λ.Object
	ϒqualities     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒqualities = Ωutils.ϒqualities
		ClippitIE = λ.Cal(λ.TypeType, λ.StrLiteral("ClippitIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ClippitIE__VALID_URL    λ.Object
				ClippitIE__real_extract λ.Object
			)
			ClippitIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?clippituser\\.tv/c/(?P<id>[a-z]+)")
			ClippitIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						FORMATS       λ.Object
						ϒformat_id    λ.Object
						ϒformats      λ.Object
						ϒmatch        λ.Object
						ϒquality      λ.Object
						ϒself         = λargs[0]
						ϒthumbnail    λ.Object
						ϒtimestamp    λ.Object
						ϒtitle        λ.Object
						ϒuploader     λ.Object
						ϒuploader_url λ.Object
						ϒurl          = λargs[1]
						ϒvideo_id     λ.Object
						ϒwebpage      λ.Object
						τmp0          λ.Object
						τmp1          λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒvideo_id)
					ϒtitle = λ.Calm(ϒself, "_html_search_regex", λ.StrLiteral("<title.*>(.+?)</title>"), ϒwebpage, λ.StrLiteral("title"))
					FORMATS = λ.NewTuple(
						λ.StrLiteral("sd"),
						λ.StrLiteral("hd"),
					)
					ϒquality = λ.Cal(ϒqualities, FORMATS)
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, FORMATS)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒformat_id = τmp1
						ϒurl = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
							λ.Mod(λ.StrLiteral("data-%s-file=\"(.+?)\""), ϒformat_id),
							ϒwebpage,
							λ.StrLiteral("url"),
						), λ.KWArgs{
							{Name: "fatal", Value: λ.False},
						})
						if !λ.IsTrue(ϒurl) {
							continue
						}
						ϒmatch = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("/(?P<height>\\d+)\\.mp4"), ϒurl)
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"url":       ϒurl,
							"format_id": ϒformat_id,
							"quality":   λ.Cal(ϒquality, ϒformat_id),
							"height": func() λ.Object {
								if λ.IsTrue(ϒmatch) {
									return λ.Cal(λ.IntType, λ.Calm(ϒmatch, "group", λ.StrLiteral("height")))
								} else {
									return λ.None
								}
							}(),
						}))
					}
					ϒuploader = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("class=\"username\".*>\\s+(.+?)\\n"),
						ϒwebpage,
						λ.StrLiteral("uploader"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒuploader_url = func() λ.Object {
						if λ.IsTrue(ϒuploader) {
							return λ.Add(λ.StrLiteral("https://www.clippituser.tv/p/"), ϒuploader)
						} else {
							return λ.None
						}
					}()
					ϒtimestamp = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("datetime=\"(.+?)\""),
						ϒwebpage,
						λ.StrLiteral("date"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒthumbnail = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("data-image=\"(.+?)\""),
						ϒwebpage,
						λ.StrLiteral("thumbnail"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					return λ.DictLiteral(map[string]λ.Object{
						"id":           ϒvideo_id,
						"title":        ϒtitle,
						"formats":      ϒformats,
						"uploader":     ϒuploader,
						"uploader_url": ϒuploader_url,
						"timestamp":    λ.Cal(ϒparse_iso8601, ϒtimestamp),
						"description":  λ.Calm(ϒself, "_og_search_description", ϒwebpage),
						"thumbnail":    ϒthumbnail,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    ClippitIE__VALID_URL,
				"_real_extract": ClippitIE__real_extract,
			})
		}())
	})
}
