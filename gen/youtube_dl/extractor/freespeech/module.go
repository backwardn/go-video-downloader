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
 * freespeech/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/freespeech.py
 */

package freespeech

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωyoutube "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/youtube"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	FreespeechIE  λ.Object
	InfoExtractor λ.Object
	YoutubeIE     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		YoutubeIE = Ωyoutube.YoutubeIE
		FreespeechIE = λ.Cal(λ.TypeType, λ.StrLiteral("FreespeechIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				FreespeechIE_IE_NAME       λ.Object
				FreespeechIE__VALID_URL    λ.Object
				FreespeechIE__real_extract λ.Object
			)
			FreespeechIE_IE_NAME = λ.StrLiteral("freespeech.org")
			FreespeechIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?freespeech\\.org/stories/(?P<id>.+)")
			FreespeechIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id  λ.Object
						ϒself        = λargs[0]
						ϒurl         = λargs[1]
						ϒwebpage     λ.Object
						ϒyoutube_url λ.Object
					)
					ϒdisplay_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒdisplay_id)
					ϒyoutube_url = λ.Calm(ϒself, "_search_regex", λ.StrLiteral("data-video-url=\"([^\"]+)\""), ϒwebpage, λ.StrLiteral("youtube url"))
					return λ.Calm(ϒself, "url_result", ϒyoutube_url, λ.Calm(YoutubeIE, "ie_key"))
				})
			return λ.DictLiteral(map[string]λ.Object{
				"IE_NAME":       FreespeechIE_IE_NAME,
				"_VALID_URL":    FreespeechIE__VALID_URL,
				"_real_extract": FreespeechIE__real_extract,
			})
		}())
	})
}
