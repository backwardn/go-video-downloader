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
 * outsidetv/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/outsidetv.py
 */

package outsidetv

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	OutsideTVIE   λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		OutsideTVIE = λ.Cal(λ.TypeType, λ.NewStr("OutsideTVIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				OutsideTVIE__TESTS        λ.Object
				OutsideTVIE__VALID_URL    λ.Object
				OutsideTVIE__real_extract λ.Object
			)
			OutsideTVIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?outsidetv\\.com/(?:[^/]+/)*?play/[a-zA-Z0-9]{8}/\\d+/\\d+/(?P<id>[a-zA-Z0-9]{8})")
			OutsideTVIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.outsidetv.com/category/snow/play/ZjQYboH6/1/10/Hdg0jukV/4"),
					λ.NewStr("md5"): λ.NewStr("192d968fedc10b2f70ec31865ffba0da"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("Hdg0jukV"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Home - Jackson Ep 1 | Arbor Snowboards"),
						λ.NewStr("description"): λ.NewStr("md5:41a12e94f3db3ca253b04bb1e8d8f4cd"),
						λ.NewStr("upload_date"): λ.NewStr("20181225"),
						λ.NewStr("timestamp"):   λ.NewInt(1545742800),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.outsidetv.com/home/play/ZjQYboH6/1/10/Hdg0jukV/4"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			OutsideTVIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒjw_media_id λ.Object
						ϒself        = λargs[0]
						ϒurl         = λargs[1]
					)
					ϒjw_media_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), λ.Add(λ.NewStr("jwplatform:"), ϒjw_media_id), λ.NewStr("JWPlatform"), ϒjw_media_id)
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        OutsideTVIE__TESTS,
				λ.NewStr("_VALID_URL"):    OutsideTVIE__VALID_URL,
				λ.NewStr("_real_extract"): OutsideTVIE__real_extract,
			})
		}())
	})
}