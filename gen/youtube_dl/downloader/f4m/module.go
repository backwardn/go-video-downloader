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
 * f4m/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/downloader/f4m.py
 */

package f4m

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωfragment "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/downloader/fragment"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	FragmentFD                    λ.Object
	ϒcompat_b64decode             λ.Object
	ϒcompat_etree_fromstring      λ.Object
	ϒcompat_urllib_parse_urlparse λ.Object
	ϒfix_xml_ampersands           λ.Object
	ϒxpath_text                   λ.Object
)

func init() {
	λ.InitModule(func() {
		FragmentFD = Ωfragment.FragmentFD
		ϒcompat_b64decode = Ωcompat.ϒcompat_b64decode
		ϒcompat_etree_fromstring = Ωcompat.ϒcompat_etree_fromstring
		ϒcompat_urllib_parse_urlparse = Ωcompat.ϒcompat_urllib_parse_urlparse
		ϒfix_xml_ampersands = Ωutils.ϒfix_xml_ampersands
		ϒxpath_text = Ωutils.ϒxpath_text
	})
}
