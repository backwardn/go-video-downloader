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
 * photobucket/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/photobucket.py
 */

package photobucket

import (
	Ωjson "github.com/tenta-browser/go-video-downloader/gen/json"
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor                λ.Object
	PhotobucketIE                λ.Object
	ϒcompat_urllib_parse_unquote λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_urllib_parse_unquote = Ωcompat.ϒcompat_urllib_parse_unquote
		PhotobucketIE = λ.Cal(λ.TypeType, λ.StrLiteral("PhotobucketIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				PhotobucketIE__VALID_URL    λ.Object
				PhotobucketIE__real_extract λ.Object
			)
			PhotobucketIE__VALID_URL = λ.StrLiteral("https?://(?:[a-z0-9]+\\.)?photobucket\\.com/.*(([\\?\\&]current=)|_)(?P<id>.*)\\.(?P<ext>(flv)|(mp4))")
			PhotobucketIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒinfo            λ.Object
						ϒinfo_json       λ.Object
						ϒmobj            λ.Object
						ϒself            = λargs[0]
						ϒurl             = λargs[1]
						ϒvideo_extension λ.Object
						ϒvideo_id        λ.Object
						ϒwebpage         λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("id"))
					ϒvideo_extension = λ.Calm(ϒmobj, "group", λ.StrLiteral("ext"))
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒvideo_id)
					λ.Calm(ϒself, "report_extraction", ϒvideo_id)
					ϒinfo_json = λ.Calm(ϒself, "_search_regex", λ.StrLiteral("Pb\\.Data\\.Shared\\.put\\(Pb\\.Data\\.Shared\\.MEDIA, (.*?)\\);"), ϒwebpage, λ.StrLiteral("info json"))
					ϒinfo = λ.Cal(Ωjson.ϒloads, ϒinfo_json)
					ϒurl = λ.Cal(ϒcompat_urllib_parse_unquote, λ.Calm(ϒself, "_html_search_regex", λ.StrLiteral("file=(.+\\.mp4)"), λ.GetItem(λ.GetItem(ϒinfo, λ.StrLiteral("linkcodes")), λ.StrLiteral("html")), λ.StrLiteral("url")))
					return λ.DictLiteral(map[string]λ.Object{
						"id":        ϒvideo_id,
						"url":       ϒurl,
						"uploader":  λ.GetItem(ϒinfo, λ.StrLiteral("username")),
						"timestamp": λ.GetItem(ϒinfo, λ.StrLiteral("creationDate")),
						"title":     λ.GetItem(ϒinfo, λ.StrLiteral("title")),
						"ext":       ϒvideo_extension,
						"thumbnail": λ.GetItem(ϒinfo, λ.StrLiteral("thumbUrl")),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    PhotobucketIE__VALID_URL,
				"_real_extract": PhotobucketIE__real_extract,
			})
		}())
	})
}
