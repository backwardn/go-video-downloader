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
 * sbs/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/sbs.py
 */

package sbs

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError λ.Object
	InfoExtractor  λ.Object
	SBSIE          λ.Object
	ϒsmuggle_url   λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒsmuggle_url = Ωutils.ϒsmuggle_url
		ExtractorError = Ωutils.ExtractorError
		SBSIE = λ.Cal(λ.TypeType, λ.StrLiteral("SBSIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				SBSIE__VALID_URL    λ.Object
				SBSIE__real_extract λ.Object
			)
			SBSIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?sbs\\.com\\.au/(?:ondemand|news)/video/(?:single/)?(?P<id>[0-9]+)")
			SBSIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒerror           λ.Object
						ϒerror_code      λ.Object
						ϒerror_message   λ.Object
						ϒplayer_params   λ.Object
						ϒself            = λargs[0]
						ϒtheplatform_url λ.Object
						ϒurl             = λargs[1]
						ϒurls            λ.Object
						ϒvideo_data      λ.Object
						ϒvideo_id        λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒplayer_params = λ.Calm(ϒself, "_download_json", λ.Mod(λ.StrLiteral("http://www.sbs.com.au/api/video_pdkvars/id/%s?form=json"), ϒvideo_id), ϒvideo_id)
					ϒerror = λ.Calm(ϒplayer_params, "get", λ.StrLiteral("error"))
					if λ.IsTrue(ϒerror) {
						ϒerror_message = λ.StrLiteral("Sorry, The video you are looking for does not exist.")
						ϒvideo_data = func() λ.Object {
							if λv := λ.Calm(ϒerror, "get", λ.StrLiteral("results")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.DictLiteral(map[λ.Object]λ.Object{})
							}
						}()
						ϒerror_code = λ.Calm(ϒerror, "get", λ.StrLiteral("errorCode"))
						if λ.IsTrue(λ.Eq(ϒerror_code, λ.StrLiteral("ComingSoon"))) {
							ϒerror_message = λ.Mod(λ.StrLiteral("%s is not yet available."), λ.Calm(ϒvideo_data, "get", λ.StrLiteral("title"), λ.StrLiteral("")))
						} else {
							if λ.Contains(λ.NewTuple(
								λ.StrLiteral("Forbidden"),
								λ.StrLiteral("intranetAccessOnly"),
							), ϒerror_code) {
								ϒerror_message = λ.StrLiteral("Sorry, This video cannot be accessed via this website")
							} else {
								if λ.IsTrue(λ.Eq(ϒerror_code, λ.StrLiteral("Expired"))) {
									ϒerror_message = λ.Mod(λ.StrLiteral("Sorry, %s is no longer available."), λ.Calm(ϒvideo_data, "get", λ.StrLiteral("title"), λ.StrLiteral("")))
								}
							}
						}
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.StrLiteral("%s said: %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							ϒerror_message,
						))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒurls = λ.GetItem(ϒplayer_params, λ.StrLiteral("releaseUrls"))
					ϒtheplatform_url = func() λ.Object {
						if λv := λ.Calm(ϒurls, "get", λ.StrLiteral("progressive")); λ.IsTrue(λv) {
							return λv
						} else if λv := λ.Calm(ϒurls, "get", λ.StrLiteral("html")); λ.IsTrue(λv) {
							return λv
						} else if λv := λ.Calm(ϒurls, "get", λ.StrLiteral("standard")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.GetItem(ϒplayer_params, λ.StrLiteral("relatedItemsURL"))
						}
					}()
					return λ.DictLiteral(map[string]λ.Object{
						"_type":  λ.StrLiteral("url_transparent"),
						"ie_key": λ.StrLiteral("ThePlatform"),
						"id":     ϒvideo_id,
						"url": λ.Cal(ϒsmuggle_url, λ.Calm(ϒself, "_proto_relative_url", ϒtheplatform_url), λ.DictLiteral(map[string]λ.Object{
							"force_smil_url": λ.True,
						})),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    SBSIE__VALID_URL,
				"_real_extract": SBSIE__real_extract,
			})
		}())
	})
}
