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
 * flickr/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/flickr.py
 */

package flickr

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError                 λ.Object
	FlickrIE                       λ.Object
	InfoExtractor                  λ.Object
	ϒcompat_str                    λ.Object
	ϒcompat_urllib_parse_urlencode λ.Object
	ϒint_or_none                   λ.Object
	ϒqualities                     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒcompat_urllib_parse_urlencode = Ωcompat.ϒcompat_urllib_parse_urlencode
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒqualities = Ωutils.ϒqualities
		FlickrIE = λ.Cal(λ.TypeType, λ.StrLiteral("FlickrIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				FlickrIE__API_BASE_URL λ.Object
				FlickrIE__LICENSES     λ.Object
				FlickrIE__VALID_URL    λ.Object
				FlickrIE__call_api     λ.Object
				FlickrIE__real_extract λ.Object
			)
			FlickrIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.|secure\\.)?flickr\\.com/photos/[\\w\\-_@]+/(?P<id>\\d+)")
			FlickrIE__API_BASE_URL = λ.StrLiteral("https://api.flickr.com/services/rest?")
			FlickrIE__LICENSES = λ.DictLiteral(map[string]string{
				"0":  "All Rights Reserved",
				"1":  "Attribution-NonCommercial-ShareAlike",
				"2":  "Attribution-NonCommercial",
				"3":  "Attribution-NonCommercial-NoDerivs",
				"4":  "Attribution",
				"5":  "Attribution-ShareAlike",
				"6":  "Attribution-NoDerivs",
				"7":  "No known copyright restrictions",
				"8":  "United States government work",
				"9":  "Public Domain Dedication (CC0)",
				"10": "Public Domain Work",
			})
			FlickrIE__call_api = λ.NewFunction("_call_api",
				[]λ.Param{
					{Name: "self"},
					{Name: "method"},
					{Name: "video_id"},
					{Name: "api_key"},
					{Name: "note"},
					{Name: "secret", Def: λ.None},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒapi_key  = λargs[3]
						ϒdata     λ.Object
						ϒmethod   = λargs[1]
						ϒnote     = λargs[4]
						ϒquery    λ.Object
						ϒsecret   = λargs[5]
						ϒself     = λargs[0]
						ϒvideo_id = λargs[2]
					)
					ϒquery = λ.DictLiteral(map[string]λ.Object{
						"photo_id":       ϒvideo_id,
						"method":         λ.Mod(λ.StrLiteral("flickr.%s"), ϒmethod),
						"api_key":        ϒapi_key,
						"format":         λ.StrLiteral("json"),
						"nojsoncallback": λ.IntLiteral(1),
					})
					if λ.IsTrue(ϒsecret) {
						λ.SetItem(ϒquery, λ.StrLiteral("secret"), ϒsecret)
					}
					ϒdata = λ.Calm(ϒself, "_download_json", λ.Add(λ.GetAttr(ϒself, "_API_BASE_URL", nil), λ.Cal(ϒcompat_urllib_parse_urlencode, ϒquery)), ϒvideo_id, ϒnote)
					if λ.IsTrue(λ.Ne(λ.GetItem(ϒdata, λ.StrLiteral("stat")), λ.StrLiteral("ok"))) {
						panic(λ.Raise(λ.Cal(ExtractorError, λ.GetItem(ϒdata, λ.StrLiteral("message")))))
					}
					return ϒdata
				})
			FlickrIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒapi_key       λ.Object
						ϒformats       λ.Object
						ϒowner         λ.Object
						ϒpreference    λ.Object
						ϒself          = λargs[0]
						ϒstream        λ.Object
						ϒstream_type   λ.Object
						ϒstreams       λ.Object
						ϒuploader_id   λ.Object
						ϒuploader_path λ.Object
						ϒuploader_url  λ.Object
						ϒurl           = λargs[1]
						ϒvideo_id      λ.Object
						ϒvideo_info    λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒapi_key = λ.GetItem(λ.Calm(ϒself, "_download_json", λ.StrLiteral("https://www.flickr.com/hermes_error_beacon.gne"), ϒvideo_id, λ.StrLiteral("Downloading api key")), λ.StrLiteral("site_key"))
					ϒvideo_info = λ.GetItem(λ.Calm(ϒself, "_call_api", λ.StrLiteral("photos.getInfo"), ϒvideo_id, ϒapi_key, λ.StrLiteral("Downloading video info")), λ.StrLiteral("photo"))
					if λ.IsTrue(λ.Eq(λ.GetItem(ϒvideo_info, λ.StrLiteral("media")), λ.StrLiteral("video"))) {
						ϒstreams = λ.GetItem(λ.Calm(ϒself, "_call_api", λ.StrLiteral("video.getStreamInfo"), ϒvideo_id, ϒapi_key, λ.StrLiteral("Downloading streams info"), λ.GetItem(ϒvideo_info, λ.StrLiteral("secret"))), λ.StrLiteral("streams"))
						ϒpreference = λ.Cal(ϒqualities, λ.NewList(
							λ.StrLiteral("288p"),
							λ.StrLiteral("iphone_wifi"),
							λ.StrLiteral("100"),
							λ.StrLiteral("300"),
							λ.StrLiteral("700"),
							λ.StrLiteral("360p"),
							λ.StrLiteral("appletv"),
							λ.StrLiteral("720p"),
							λ.StrLiteral("1080p"),
							λ.StrLiteral("orig"),
						))
						ϒformats = λ.NewList()
						τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(ϒstreams, λ.StrLiteral("stream")))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒstream = τmp1
							ϒstream_type = λ.Cal(ϒcompat_str, λ.Calm(ϒstream, "get", λ.StrLiteral("type")))
							λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
								"format_id":  ϒstream_type,
								"url":        λ.GetItem(ϒstream, λ.StrLiteral("_content")),
								"preference": λ.Cal(ϒpreference, ϒstream_type),
							}))
						}
						λ.Calm(ϒself, "_sort_formats", ϒformats)
						ϒowner = λ.Calm(ϒvideo_info, "get", λ.StrLiteral("owner"), λ.DictLiteral(map[λ.Object]λ.Object{}))
						ϒuploader_id = λ.Calm(ϒowner, "get", λ.StrLiteral("nsid"))
						ϒuploader_path = func() λ.Object {
							if λv := λ.Calm(ϒowner, "get", λ.StrLiteral("path_alias")); λ.IsTrue(λv) {
								return λv
							} else {
								return ϒuploader_id
							}
						}()
						ϒuploader_url = func() λ.Object {
							if λ.IsTrue(ϒuploader_path) {
								return λ.Mod(λ.StrLiteral("https://www.flickr.com/photos/%s/"), ϒuploader_path)
							} else {
								return λ.None
							}
						}()
						return λ.DictLiteral(map[string]λ.Object{
							"id":            ϒvideo_id,
							"title":         λ.GetItem(λ.GetItem(ϒvideo_info, λ.StrLiteral("title")), λ.StrLiteral("_content")),
							"description":   λ.Calm(λ.Calm(ϒvideo_info, "get", λ.StrLiteral("description"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("_content")),
							"formats":       ϒformats,
							"timestamp":     λ.Cal(ϒint_or_none, λ.Calm(ϒvideo_info, "get", λ.StrLiteral("dateuploaded"))),
							"duration":      λ.Cal(ϒint_or_none, λ.Calm(λ.Calm(ϒvideo_info, "get", λ.StrLiteral("video"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("duration"))),
							"uploader_id":   ϒuploader_id,
							"uploader":      λ.Calm(ϒowner, "get", λ.StrLiteral("realname")),
							"uploader_url":  ϒuploader_url,
							"comment_count": λ.Cal(ϒint_or_none, λ.Calm(λ.Calm(ϒvideo_info, "get", λ.StrLiteral("comments"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("_content"))),
							"view_count":    λ.Cal(ϒint_or_none, λ.Calm(ϒvideo_info, "get", λ.StrLiteral("views"))),
							"tags": λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
								nil,
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
										var (
											ϒtag λ.Object
											τmp0 λ.Object
											τmp1 λ.Object
										)
										τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(λ.Calm(ϒvideo_info, "get", λ.StrLiteral("tags"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("tag"), λ.NewList()))
										for {
											if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
												break
											}
											ϒtag = τmp1
											λgy.Yield(λ.Calm(ϒtag, "get", λ.StrLiteral("_content")))
										}
										return λ.None
									})
								}))),
							"license": λ.Calm(λ.GetAttr(ϒself, "_LICENSES", nil), "get", λ.Calm(ϒvideo_info, "get", λ.StrLiteral("license"))),
						})
					} else {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.StrLiteral("not a video")), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					return λ.None
				})
			return λ.DictLiteral(map[string]λ.Object{
				"_API_BASE_URL": FlickrIE__API_BASE_URL,
				"_LICENSES":     FlickrIE__LICENSES,
				"_VALID_URL":    FlickrIE__VALID_URL,
				"_call_api":     FlickrIE__call_api,
				"_real_extract": FlickrIE__real_extract,
			})
		}())
	})
}
