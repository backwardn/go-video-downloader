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
 * formula1/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/formula1.py
 */

package formula1

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	Formula1IE    λ.Object
	InfoExtractor λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		Formula1IE = λ.Cal(λ.TypeType, λ.NewStr("Formula1IE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				Formula1IE__TESTS        λ.Object
				Formula1IE__VALID_URL    λ.Object
				Formula1IE__real_extract λ.Object
			)
			Formula1IE__VALID_URL = λ.NewStr("https?://(?:www\\.)?formula1\\.com/(?:content/fom-website/)?en/video/\\d{4}/\\d{1,2}/(?P<id>.+?)\\.html")
			Formula1IE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.formula1.com/content/fom-website/en/video/2016/5/Race_highlights_-_Spain_2016.html"),
					λ.NewStr("md5"): λ.NewStr("8c79e54be72078b26b89e0e111c0502b"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):    λ.NewStr("JvYXJpMzE6pArfHWm5ARp5AiUmD-gibV"),
						λ.NewStr("ext"):   λ.NewStr("mp4"),
						λ.NewStr("title"): λ.NewStr("Race highlights - Spain 2016"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
					λ.NewStr("add_ie"): λ.NewList(λ.NewStr("Ooyala")),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.formula1.com/en/video/2016/5/Race_highlights_-_Spain_2016.html"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			Formula1IE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id        λ.Object
						ϒooyala_embed_code λ.Object
						ϒself              = λargs[0]
						ϒurl               = λargs[1]
						ϒwebpage           λ.Object
					)
					ϒdisplay_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒdisplay_id)
					ϒooyala_embed_code = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("data-videoid=\"([^\"]+)\""), ϒwebpage, λ.NewStr("ooyala embed code"))
					return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), λ.Mod(λ.NewStr("ooyala:%s"), ϒooyala_embed_code), λ.NewStr("Ooyala"), ϒooyala_embed_code)
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        Formula1IE__TESTS,
				λ.NewStr("_VALID_URL"):    Formula1IE__VALID_URL,
				λ.NewStr("_real_extract"): Formula1IE__real_extract,
			})
		}())
	})
}