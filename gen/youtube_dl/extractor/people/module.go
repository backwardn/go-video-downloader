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
 * people/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/people.py
 */

package people

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	PeopleIE      λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		PeopleIE = λ.Cal(λ.TypeType, λ.StrLiteral("PeopleIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				PeopleIE__VALID_URL    λ.Object
				PeopleIE__real_extract λ.Object
			)
			PeopleIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?people\\.com/people/videos/0,,(?P<id>\\d+),00\\.html")
			PeopleIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
						ϒurl  = λargs[1]
					)
					return λ.Calm(ϒself, "url_result", λ.Mod(λ.StrLiteral("http://players.brightcove.net/416418724/default_default/index.html?videoId=ref:%s"), λ.Calm(ϒself, "_match_id", ϒurl)), λ.StrLiteral("BrightcoveNew"))
				})
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL":    PeopleIE__VALID_URL,
				"_real_extract": PeopleIE__real_extract,
			})
		}())
	})
}
