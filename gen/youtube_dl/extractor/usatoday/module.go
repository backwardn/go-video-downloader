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
 * usatoday/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/usatoday.py
 */

package usatoday

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError            λ.Object
	InfoExtractor             λ.Object
	USATodayIE                λ.Object
	ϒcompat_str               λ.Object
	ϒget_element_by_attribute λ.Object
	ϒparse_duration           λ.Object
	ϒupdate_url_query         λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒget_element_by_attribute = Ωutils.ϒget_element_by_attribute
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		ExtractorError = Ωutils.ExtractorError
		ϒcompat_str = Ωcompat.ϒcompat_str
		USATodayIE = λ.Cal(λ.TypeType, λ.NewStr("USATodayIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				USATodayIE_BRIGHTCOVE_URL_TEMPLATE λ.Object
				USATodayIE__TEST                   λ.Object
				USATodayIE__VALID_URL              λ.Object
				USATodayIE__real_extract           λ.Object
			)
			USATodayIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?usatoday\\.com/(?:[^/]+/)*(?P<id>[^?/#]+)")
			USATodayIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("http://www.usatoday.com/media/cinematic/video/81729424/us-france-warn-syrian-regime-ahead-of-new-peace-talks/"),
				λ.NewStr("md5"): λ.NewStr("4d40974481fa3475f8bccfd20c5361f8"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):          λ.NewStr("81729424"),
					λ.NewStr("ext"):         λ.NewStr("mp4"),
					λ.NewStr("title"):       λ.NewStr("US, France warn Syrian regime ahead of new peace talks"),
					λ.NewStr("timestamp"):   λ.NewInt(1457891045),
					λ.NewStr("description"): λ.NewStr("md5:7e50464fdf2126b0f533748d3c78d58f"),
					λ.NewStr("uploader_id"): λ.NewStr("29906170001"),
					λ.NewStr("upload_date"): λ.NewStr("20160313"),
				}),
			})
			USATodayIE_BRIGHTCOVE_URL_TEMPLATE = λ.NewStr("http://players.brightcove.net/29906170001/38a9eecc-bdd8-42a3-ba14-95397e48b3f8_default/index.html?videoId=%s")
			USATodayIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id    λ.Object
						ϒself          = λargs[0]
						ϒui_video_data λ.Object
						ϒurl           = λargs[1]
						ϒvideo_data    λ.Object
						ϒwebpage       λ.Object
					)
					ϒdisplay_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), λ.Cal(ϒupdate_url_query, ϒurl, λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("ajax"): λ.NewStr("true"),
					})), ϒdisplay_id)
					ϒui_video_data = λ.Cal(ϒget_element_by_attribute, λ.NewStr("class"), λ.NewStr("ui-video-data"), ϒwebpage)
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒui_video_data))) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.NewStr("no video on the webpage")), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒvideo_data = λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), ϒui_video_data, ϒdisplay_id)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("_type"):       λ.NewStr("url_transparent"),
						λ.NewStr("url"):         λ.Mod(λ.GetAttr(ϒself, "BRIGHTCOVE_URL_TEMPLATE", nil), λ.GetItem(ϒvideo_data, λ.NewStr("brightcove_id"))),
						λ.NewStr("id"):          λ.Cal(ϒcompat_str, λ.GetItem(ϒvideo_data, λ.NewStr("id"))),
						λ.NewStr("title"):       λ.GetItem(ϒvideo_data, λ.NewStr("title")),
						λ.NewStr("thumbnail"):   λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("thumbnail")),
						λ.NewStr("description"): λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("description")),
						λ.NewStr("duration"):    λ.Cal(ϒparse_duration, λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("length"))),
						λ.NewStr("ie_key"):      λ.NewStr("BrightcoveNew"),
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("BRIGHTCOVE_URL_TEMPLATE"): USATodayIE_BRIGHTCOVE_URL_TEMPLATE,
				λ.NewStr("_TEST"):                   USATodayIE__TEST,
				λ.NewStr("_VALID_URL"):              USATodayIE__VALID_URL,
				λ.NewStr("_real_extract"):           USATodayIE__real_extract,
			})
		}())
	})
}