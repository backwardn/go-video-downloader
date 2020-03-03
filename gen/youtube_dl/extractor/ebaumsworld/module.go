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
 * ebaumsworld/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/ebaumsworld.py
 */

package ebaumsworld

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	EbaumsWorldIE λ.Object
	InfoExtractor λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		EbaumsWorldIE = λ.Cal(λ.TypeType, λ.StrLiteral("EbaumsWorldIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				EbaumsWorldIE__VALID_URL    λ.Object
				EbaumsWorldIE__real_extract λ.Object
			)
			EbaumsWorldIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?ebaumsworld\\.com/videos/[^/]+/(?P<id>\\d+)")
			EbaumsWorldIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒconfig    λ.Object
						ϒself      = λargs[0]
						ϒurl       = λargs[1]
						ϒvideo_id  λ.Object
						ϒvideo_url λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒconfig = λ.Calm(ϒself, "_download_xml", λ.Mod(λ.StrLiteral("http://www.ebaumsworld.com/video/player/%s"), ϒvideo_id), ϒvideo_id)
					ϒvideo_url = λ.GetAttr(λ.Calm(ϒconfig, "find", λ.StrLiteral("file")), "text", nil)
					return λ.DictLiteral(map[string]λ.Object{
						"id":          ϒvideo_id,
						"title":       λ.GetAttr(λ.Calm(ϒconfig, "find", λ.StrLiteral("title")), "text", nil),
						"url":         ϒvideo_url,
						"description": λ.GetAttr(λ.Calm(ϒconfig, "find", λ.StrLiteral("description")), "text", nil),
						"thumbnail":   λ.GetAttr(λ.Calm(ϒconfig, "find", λ.StrLiteral("image")), "text", nil),
						"uploader":    λ.GetAttr(λ.Calm(ϒconfig, "find", λ.StrLiteral("username")), "text", nil),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    EbaumsWorldIE__VALID_URL,
				"_real_extract": EbaumsWorldIE__real_extract,
			})
		}())
	})
}
