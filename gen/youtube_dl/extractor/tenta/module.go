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
 * tenta/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/tenta.py
 */

package tenta

import (
	Ωparse "github.com/tenta-browser/go-video-downloader/gen/urllib/parse"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	TentaIE       λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		TentaIE = λ.Cal(λ.TypeType, λ.NewStr("TentaIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TentaIE__TESTS        λ.Object
				TentaIE__VALID_URL    λ.Object
				TentaIE__real_extract λ.Object
			)
			TentaIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?tenta\\.com/how-to-download-videos")
			TentaIE__TESTS = λ.NewList(λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("https://tenta.com/how-to-download-videos"),
				λ.NewStr("md5"): λ.NewStr("815a64f01f7ebfb1607220396f3ffa7d"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):    λ.NewStr("howto"),
					λ.NewStr("ext"):   λ.NewStr("mp4"),
					λ.NewStr("title"): λ.NewStr("How To Download Videos & Save Securely | Tenta VPN Browser"),
				}),
			}))
			TentaIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself      = λargs[0]
						ϒtitle     λ.Object
						ϒurl       = λargs[1]
						ϒvideo_id  λ.Object
						ϒvideo_url λ.Object
						ϒwebpage   λ.Object
					)
					ϒvideo_id = λ.NewStr("howto")
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					ϒtitle = λ.Call(λ.GetAttr(ϒself, "_og_search_title", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					ϒvideo_url = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("<source[^>]+src=[\"\\'](.+?)[\"\\'][^>]+type=[\"\\']video/mp4[\"\\'][^>]*/>"), ϒwebpage, λ.NewStr("video_url"))
					ϒvideo_url = λ.Cal(Ωparse.ϒurljoin, ϒurl, ϒvideo_url)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):    ϒvideo_id,
						λ.NewStr("title"): ϒtitle,
						λ.NewStr("url"):   ϒvideo_url,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        TentaIE__TESTS,
				λ.NewStr("_VALID_URL"):    TentaIE__VALID_URL,
				λ.NewStr("_real_extract"): TentaIE__real_extract,
			})
		}())
	})
}