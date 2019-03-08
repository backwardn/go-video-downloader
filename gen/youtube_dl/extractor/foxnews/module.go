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
 * foxnews/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/foxnews.py
 */

package foxnews

import (
	Ωamp "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/amp"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	AMPIE            λ.Object
	FoxNewsArticleIE λ.Object
	FoxNewsIE        λ.Object
	InfoExtractor    λ.Object
)

func init() {
	λ.InitModule(func() {
		AMPIE = Ωamp.AMPIE
		InfoExtractor = Ωcommon.InfoExtractor
		FoxNewsIE = λ.Cal(λ.TypeType, λ.NewStr("FoxNewsIE"), λ.NewTuple(AMPIE), func() λ.Dict {
			var (
				FoxNewsIE__VALID_URL λ.Object
			)
			FoxNewsIE__VALID_URL = λ.NewStr("https?://(?P<host>video\\.(?:insider\\.)?fox(?:news|business)\\.com)/v/(?:video-embed\\.html\\?video_id=)?(?P<id>\\d+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): FoxNewsIE__VALID_URL,
			})
		}())
		FoxNewsArticleIE = λ.Cal(λ.TypeType, λ.NewStr("FoxNewsArticleIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				FoxNewsArticleIE__VALID_URL λ.Object
			)
			FoxNewsArticleIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?(?:insider\\.)?foxnews\\.com/(?!v)([^/]+/)+(?P<id>[a-z-]+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): FoxNewsArticleIE__VALID_URL,
			})
		}())
	})
}
