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
 * dispeak/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/dispeak.py
 */

package dispeak

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	DigitallySpeakingIE λ.Object
	InfoExtractor       λ.Object
	ϒint_or_none        λ.Object
	ϒparse_duration     λ.Object
	ϒremove_end         λ.Object
	ϒxpath_element      λ.Object
	ϒxpath_text         λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒremove_end = Ωutils.ϒremove_end
		ϒxpath_element = Ωutils.ϒxpath_element
		ϒxpath_text = Ωutils.ϒxpath_text
		DigitallySpeakingIE = λ.Cal(λ.TypeType, λ.StrLiteral("DigitallySpeakingIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				DigitallySpeakingIE__VALID_URL    λ.Object
				DigitallySpeakingIE__parse_mp4    λ.Object
				DigitallySpeakingIE__real_extract λ.Object
			)
			DigitallySpeakingIE__VALID_URL = λ.StrLiteral("https?://(?:s?evt\\.dispeak|events\\.digitallyspeaking)\\.com/(?:[^/]+/)+xml/(?P<id>[^.]+)\\.xml")
			DigitallySpeakingIE__parse_mp4 = λ.NewFunction("_parse_mp4",
				[]λ.Param{
					{Name: "self"},
					{Name: "metadata"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒa_format      λ.Object
						ϒabr           λ.Object
						ϒbitrate       λ.Object
						ϒformats       λ.Object
						ϒhttp_host     λ.Object
						ϒmetadata      = λargs[1]
						ϒmobj          λ.Object
						ϒmp4_video     λ.Object
						ϒself          = λargs[0]
						ϒstream_name   λ.Object
						ϒtbr           λ.Object
						ϒurl           λ.Object
						ϒvbr           λ.Object
						ϒvideo_formats λ.Object
						ϒvideo_path    λ.Object
						ϒvideo_root    λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
					)
					ϒvideo_formats = λ.NewList()
					ϒvideo_root = λ.None
					ϒmp4_video = λ.Call(ϒxpath_text, λ.NewArgs(
						ϒmetadata,
						λ.StrLiteral("./mp4video"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if ϒmp4_video != λ.None {
						ϒmobj = λ.Cal(Ωre.ϒmatch, λ.StrLiteral("(?P<root>https?://.*?/).*"), ϒmp4_video)
						ϒvideo_root = λ.Calm(ϒmobj, "group", λ.StrLiteral("root"))
					}
					if ϒvideo_root == λ.None {
						ϒhttp_host = λ.Call(ϒxpath_text, λ.NewArgs(
							ϒmetadata,
							λ.StrLiteral("httpHost"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
						if λ.IsTrue(ϒhttp_host) {
							ϒvideo_root = λ.Mod(λ.StrLiteral("http://%s/"), ϒhttp_host)
						}
					}
					if ϒvideo_root == λ.None {
						ϒvideo_root = λ.StrLiteral("http://s3-2u.digitallyspeaking.com/")
					}
					ϒformats = λ.Calm(ϒmetadata, "findall", λ.StrLiteral("./MBRVideos/MBRVideo"))
					if !λ.IsTrue(ϒformats) {
						return λ.None
					}
					τmp0 = λ.Cal(λ.BuiltinIter, ϒformats)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒa_format = τmp1
						ϒstream_name = λ.Call(ϒxpath_text, λ.NewArgs(
							ϒa_format,
							λ.StrLiteral("streamName"),
						), λ.KWArgs{
							{Name: "fatal", Value: λ.True},
						})
						ϒvideo_path = λ.Calm(λ.Cal(Ωre.ϒmatch, λ.StrLiteral("mp4\\:(?P<path>.*)"), ϒstream_name), "group", λ.StrLiteral("path"))
						ϒurl = λ.Add(ϒvideo_root, ϒvideo_path)
						ϒbitrate = λ.Cal(ϒxpath_text, ϒa_format, λ.StrLiteral("bitrate"))
						ϒtbr = λ.Cal(ϒint_or_none, ϒbitrate)
						ϒvbr = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.StrLiteral("-(\\d+)\\.mp4"),
							ϒvideo_path,
							λ.StrLiteral("vbr"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}))
						ϒabr = func() λ.Object {
							if λ.IsTrue(func() λ.Object {
								if λv := ϒtbr; !λ.IsTrue(λv) {
									return λv
								} else {
									return ϒvbr
								}
							}()) {
								return λ.Sub(ϒtbr, ϒvbr)
							} else {
								return λ.None
							}
						}()
						λ.Calm(ϒvideo_formats, "append", λ.DictLiteral(map[string]λ.Object{
							"format_id": ϒbitrate,
							"url":       ϒurl,
							"tbr":       ϒtbr,
							"vbr":       ϒvbr,
							"abr":       ϒabr,
						}))
					}
					return ϒvideo_formats
				})
			DigitallySpeakingIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒmetadata        λ.Object
						ϒself            = λargs[0]
						ϒurl             = λargs[1]
						ϒvideo_formats   λ.Object
						ϒvideo_id        λ.Object
						ϒxml_description λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒxml_description = λ.Calm(ϒself, "_download_xml", ϒurl, ϒvideo_id)
					ϒmetadata = λ.Cal(ϒxpath_element, ϒxml_description, λ.StrLiteral("metadata"))
					ϒvideo_formats = λ.Calm(ϒself, "_parse_mp4", ϒmetadata)
					if ϒvideo_formats == λ.None {
						ϒvideo_formats = λ.Calm(ϒself, "_parse_flv", ϒmetadata)
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id":      ϒvideo_id,
						"formats": ϒvideo_formats,
						"title": λ.Call(ϒxpath_text, λ.NewArgs(
							ϒmetadata,
							λ.StrLiteral("title"),
						), λ.KWArgs{
							{Name: "fatal", Value: λ.True},
						}),
						"duration": λ.Cal(ϒparse_duration, λ.Cal(ϒxpath_text, ϒmetadata, λ.StrLiteral("endTime"))),
						"creator":  λ.Cal(ϒxpath_text, ϒmetadata, λ.StrLiteral("speaker")),
					})
				})
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL":    DigitallySpeakingIE__VALID_URL,
				"_parse_mp4":    DigitallySpeakingIE__parse_mp4,
				"_real_extract": DigitallySpeakingIE__real_extract,
			})
		}())
	})
}
