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
 * playfm/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/playfm.py
 */

package playfm

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError λ.Object
	InfoExtractor  λ.Object
	PlayFMIE       λ.Object
	ϒcompat_str    λ.Object
	ϒint_or_none   λ.Object
	ϒparse_iso8601 λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		PlayFMIE = λ.Cal(λ.TypeType, λ.StrLiteral("PlayFMIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				PlayFMIE_IE_NAME       λ.Object
				PlayFMIE__VALID_URL    λ.Object
				PlayFMIE__real_extract λ.Object
			)
			PlayFMIE_IE_NAME = λ.StrLiteral("play.fm")
			PlayFMIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?play\\.fm/(?P<slug>(?:[^/]+/)+(?P<id>[^/]+))/?(?:$|[?#])")
			PlayFMIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒaudio_url     λ.Object
						ϒcategories    λ.Object
						ϒcomment_count λ.Object
						ϒdescription   λ.Object
						ϒduration      λ.Object
						ϒerror         λ.Object
						ϒmobj          λ.Object
						ϒrecordings    λ.Object
						ϒself          = λargs[0]
						ϒslug          λ.Object
						ϒtimestamp     λ.Object
						ϒtitle         λ.Object
						ϒuploader      λ.Object
						ϒuploader_id   λ.Object
						ϒurl           = λargs[1]
						ϒvideo_id      λ.Object
						ϒview_count    λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("id"))
					ϒslug = λ.Calm(ϒmobj, "group", λ.StrLiteral("slug"))
					ϒrecordings = λ.Calm(ϒself, "_download_json", λ.Mod(λ.StrLiteral("http://v2api.play.fm/recordings/slug/%s"), ϒslug), ϒvideo_id)
					ϒerror = λ.Calm(ϒrecordings, "get", λ.StrLiteral("error"))
					if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒerror, λ.DictType)) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.StrLiteral("%s returned error: %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							λ.Calm(ϒerror, "get", λ.StrLiteral("message")),
						))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒaudio_url = λ.GetItem(ϒrecordings, λ.StrLiteral("audio"))
					ϒvideo_id = λ.Cal(ϒcompat_str, func() λ.Object {
						if λv := λ.Calm(ϒrecordings, "get", λ.StrLiteral("id")); λ.IsTrue(λv) {
							return λv
						} else {
							return ϒvideo_id
						}
					}())
					ϒtitle = λ.GetItem(ϒrecordings, λ.StrLiteral("title"))
					ϒdescription = λ.Calm(ϒrecordings, "get", λ.StrLiteral("description"))
					ϒduration = λ.Cal(ϒint_or_none, λ.Calm(ϒrecordings, "get", λ.StrLiteral("recordingDuration")))
					ϒtimestamp = λ.Cal(ϒparse_iso8601, λ.Calm(ϒrecordings, "get", λ.StrLiteral("created_at")))
					ϒuploader = λ.Calm(λ.Calm(ϒrecordings, "get", λ.StrLiteral("page"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("title"))
					ϒuploader_id = λ.Cal(ϒcompat_str, λ.Calm(λ.Calm(ϒrecordings, "get", λ.StrLiteral("page"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("id")))
					ϒview_count = λ.Cal(ϒint_or_none, λ.Calm(ϒrecordings, "get", λ.StrLiteral("playCount")))
					ϒcomment_count = λ.Cal(ϒint_or_none, λ.Calm(ϒrecordings, "get", λ.StrLiteral("commentCount")))
					ϒcategories = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒtag λ.Object
									τmp0 λ.Object
									τmp1 λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒrecordings, "get", λ.StrLiteral("tags"), λ.NewList()))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒtag = τmp1
									if λ.IsTrue(λ.Calm(ϒtag, "get", λ.StrLiteral("name"))) {
										λgy.Yield(λ.GetItem(ϒtag, λ.StrLiteral("name")))
									}
								}
								return λ.None
							})
						})))
					return λ.DictLiteral(map[string]λ.Object{
						"id":            ϒvideo_id,
						"url":           ϒaudio_url,
						"title":         ϒtitle,
						"description":   ϒdescription,
						"duration":      ϒduration,
						"timestamp":     ϒtimestamp,
						"uploader":      ϒuploader,
						"uploader_id":   ϒuploader_id,
						"view_count":    ϒview_count,
						"comment_count": ϒcomment_count,
						"categories":    ϒcategories,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":       PlayFMIE_IE_NAME,
				"_VALID_URL":    PlayFMIE__VALID_URL,
				"_real_extract": PlayFMIE__real_extract,
			})
		}())
	})
}
