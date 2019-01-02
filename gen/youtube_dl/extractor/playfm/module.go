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
 * playfm/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/playfm.py
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
		PlayFMIE = λ.Cal(λ.TypeType, λ.NewStr("PlayFMIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				PlayFMIE_IE_NAME       λ.Object
				PlayFMIE__TEST         λ.Object
				PlayFMIE__VALID_URL    λ.Object
				PlayFMIE__real_extract λ.Object
			)
			PlayFMIE_IE_NAME = λ.NewStr("play.fm")
			PlayFMIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?play\\.fm/(?P<slug>(?:[^/]+/)+(?P<id>[^/]+))/?(?:$|[?#])")
			PlayFMIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("https://www.play.fm/dan-drastic/sven-tasnadi-leipzig-electronic-music-batofar-paris-fr-2014-07-12"),
				λ.NewStr("md5"): λ.NewStr("c505f8307825a245d0c7ad1850001f22"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):            λ.NewStr("71276"),
					λ.NewStr("ext"):           λ.NewStr("mp3"),
					λ.NewStr("title"):         λ.NewStr("Sven Tasnadi - LEIPZIG ELECTRONIC MUSIC @ Batofar (Paris,FR) - 2014-07-12"),
					λ.NewStr("description"):   λ.NewStr(""),
					λ.NewStr("duration"):      λ.NewInt(5627),
					λ.NewStr("timestamp"):     λ.NewInt(1406033781),
					λ.NewStr("upload_date"):   λ.NewStr("20140722"),
					λ.NewStr("uploader"):      λ.NewStr("Dan Drastic"),
					λ.NewStr("uploader_id"):   λ.NewStr("71170"),
					λ.NewStr("view_count"):    λ.IntType,
					λ.NewStr("comment_count"): λ.IntType,
				}),
			})
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
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					ϒslug = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("slug"))
					ϒrecordings = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.NewStr("http://v2api.play.fm/recordings/slug/%s"), ϒslug), ϒvideo_id)
					ϒerror = λ.Cal(λ.GetAttr(ϒrecordings, "get", nil), λ.NewStr("error"))
					if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒerror, λ.DictType)) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("%s returned error: %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							λ.Cal(λ.GetAttr(ϒerror, "get", nil), λ.NewStr("message")),
						))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒaudio_url = λ.GetItem(ϒrecordings, λ.NewStr("audio"))
					ϒvideo_id = λ.Cal(ϒcompat_str, func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒrecordings, "get", nil), λ.NewStr("id")); λ.IsTrue(λv) {
							return λv
						} else {
							return ϒvideo_id
						}
					}())
					ϒtitle = λ.GetItem(ϒrecordings, λ.NewStr("title"))
					ϒdescription = λ.Cal(λ.GetAttr(ϒrecordings, "get", nil), λ.NewStr("description"))
					ϒduration = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒrecordings, "get", nil), λ.NewStr("recordingDuration")))
					ϒtimestamp = λ.Cal(ϒparse_iso8601, λ.Cal(λ.GetAttr(ϒrecordings, "get", nil), λ.NewStr("created_at")))
					ϒuploader = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒrecordings, "get", nil), λ.NewStr("page"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("title"))
					ϒuploader_id = λ.Cal(ϒcompat_str, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒrecordings, "get", nil), λ.NewStr("page"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("id")))
					ϒview_count = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒrecordings, "get", nil), λ.NewStr("playCount")))
					ϒcomment_count = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒrecordings, "get", nil), λ.NewStr("commentCount")))
					ϒcategories = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒtag λ.Object
									τmp0 λ.Object
									τmp1 λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒrecordings, "get", nil), λ.NewStr("tags"), λ.NewList()))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒtag = τmp1
									if λ.IsTrue(λ.Cal(λ.GetAttr(ϒtag, "get", nil), λ.NewStr("name"))) {
										λgen.Yield(λ.GetItem(ϒtag, λ.NewStr("name")))
									}
								}
								return λ.None
							})
						})))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):            ϒvideo_id,
						λ.NewStr("url"):           ϒaudio_url,
						λ.NewStr("title"):         ϒtitle,
						λ.NewStr("description"):   ϒdescription,
						λ.NewStr("duration"):      ϒduration,
						λ.NewStr("timestamp"):     ϒtimestamp,
						λ.NewStr("uploader"):      ϒuploader,
						λ.NewStr("uploader_id"):   ϒuploader_id,
						λ.NewStr("view_count"):    ϒview_count,
						λ.NewStr("comment_count"): ϒcomment_count,
						λ.NewStr("categories"):    ϒcategories,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       PlayFMIE_IE_NAME,
				λ.NewStr("_TEST"):         PlayFMIE__TEST,
				λ.NewStr("_VALID_URL"):    PlayFMIE__VALID_URL,
				λ.NewStr("_real_extract"): PlayFMIE__real_extract,
			})
		}())
	})
}