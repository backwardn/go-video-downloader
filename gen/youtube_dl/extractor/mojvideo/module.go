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
 * mojvideo/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/mojvideo.py
 */

package mojvideo

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError  λ.Object
	InfoExtractor   λ.Object
	MojvideoIE      λ.Object
	ϒparse_duration λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ExtractorError = Ωutils.ExtractorError
		ϒparse_duration = Ωutils.ϒparse_duration
		MojvideoIE = λ.Cal(λ.TypeType, λ.StrLiteral("MojvideoIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				MojvideoIE__VALID_URL    λ.Object
				MojvideoIE__real_extract λ.Object
			)
			MojvideoIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?mojvideo\\.com/video-(?P<display_id>[^/]+)/(?P<id>[a-f0-9]+)")
			MojvideoIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id λ.Object
						ϒduration   λ.Object
						ϒerror_desc λ.Object
						ϒmobj       λ.Object
						ϒplayerapi  λ.Object
						ϒself       = λargs[0]
						ϒthumbnail  λ.Object
						ϒtitle      λ.Object
						ϒurl        = λargs[1]
						ϒvideo_id   λ.Object
						ϒvideo_url  λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("id"))
					ϒdisplay_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("display_id"))
					ϒplayerapi = λ.Calm(ϒself, "_download_webpage", λ.Mod(λ.StrLiteral("http://www.mojvideo.com/playerapi.php?v=%s&t=1"), ϒvideo_id), ϒdisplay_id)
					if λ.Contains(ϒplayerapi, λ.StrLiteral("<error>true</error>")) {
						ϒerror_desc = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
							λ.StrLiteral("<errordesc>([^<]*)</errordesc>"),
							ϒplayerapi,
							λ.StrLiteral("error description"),
						), λ.KWArgs{
							{Name: "fatal", Value: λ.False},
						})
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.StrLiteral("%s said: %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							ϒerror_desc,
						))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒtitle = λ.Calm(ϒself, "_html_search_regex", λ.StrLiteral("<title>([^<]+)</title>"), ϒplayerapi, λ.StrLiteral("title"))
					ϒvideo_url = λ.Calm(ϒself, "_html_search_regex", λ.StrLiteral("<file>([^<]+)</file>"), ϒplayerapi, λ.StrLiteral("video URL"))
					ϒthumbnail = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("<preview>([^<]+)</preview>"),
						ϒplayerapi,
						λ.StrLiteral("thumbnail"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒduration = λ.Cal(ϒparse_duration, λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("<duration>([^<]+)</duration>"),
						ϒplayerapi,
						λ.StrLiteral("duration"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					return λ.DictLiteral(map[string]λ.Object{
						"id":         ϒvideo_id,
						"display_id": ϒdisplay_id,
						"url":        ϒvideo_url,
						"title":      ϒtitle,
						"thumbnail":  ϒthumbnail,
						"duration":   ϒduration,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    MojvideoIE__VALID_URL,
				"_real_extract": MojvideoIE__real_extract,
			})
		}())
	})
}
