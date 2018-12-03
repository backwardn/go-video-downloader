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
 * dbtv/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/dbtv.py
 */

package dbtv

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	DBTVIE        λ.Object
	InfoExtractor λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		DBTVIE = λ.Cal(λ.TypeType, λ.NewStr("DBTVIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				DBTVIE__VALID_URL λ.Object
			)
			DBTVIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?dbtv\\.no/(?:[^/]+/)?(?P<id>[0-9]+)(?:#(?P<display_id>.+))?")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): DBTVIE__VALID_URL,
			})
		}())
	})
}
