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
 * once/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/once.py
 */

package once

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	OnceIE        λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		OnceIE = λ.Cal(λ.TypeType, λ.StrLiteral("OnceIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				OnceIE__VALID_URL λ.Object
			)
			OnceIE__VALID_URL = λ.StrLiteral("https?://.+?\\.unicornmedia\\.com/now/(?:ads/vmap/)?[^/]+/[^/]+/(?P<domain_id>[^/]+)/(?P<application_id>[^/]+)/(?:[^/]+/)?(?P<media_item_id>[^/]+)/content\\.(?:once|m3u8|mp4)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": OnceIE__VALID_URL,
			})
		}())
	})
}
