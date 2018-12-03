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
 * webcaster/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/webcaster.py
 */

package webcaster

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor   λ.Object
	WebcasterFeedIE λ.Object
	WebcasterIE     λ.Object
	ϒdetermine_ext  λ.Object
	ϒxpath_text     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ϒxpath_text = Ωutils.ϒxpath_text
		WebcasterIE = λ.Cal(λ.TypeType, λ.NewStr("WebcasterIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				WebcasterIE__VALID_URL λ.Object
			)
			WebcasterIE__VALID_URL = λ.NewStr("https?://bl\\.webcaster\\.pro/(?:quote|media)/start/free_(?P<id>[^/]+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): WebcasterIE__VALID_URL,
			})
		}())
		WebcasterFeedIE = λ.Cal(λ.TypeType, λ.NewStr("WebcasterFeedIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				WebcasterFeedIE__VALID_URL λ.Object
			)
			WebcasterFeedIE__VALID_URL = λ.NewStr("https?://bl\\.webcaster\\.pro/feed/start/free_(?P<id>[^/]+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): WebcasterFeedIE__VALID_URL,
			})
		}())
	})
}
