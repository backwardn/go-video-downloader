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
 * camdemy/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/camdemy.py
 */

package camdemy

import (
	Ωparse "github.com/tenta-browser/go-video-downloader/gen/urllib/parse"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	CamdemyFolderIE                λ.Object
	CamdemyIE                      λ.Object
	InfoExtractor                  λ.Object
	ϒclean_html                    λ.Object
	ϒcompat_urllib_parse_urlencode λ.Object
	ϒparse_duration                λ.Object
	ϒstr_to_int                    λ.Object
	ϒunified_strdate               λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_urllib_parse_urlencode = Ωcompat.ϒcompat_urllib_parse_urlencode
		ϒclean_html = Ωutils.ϒclean_html
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒstr_to_int = Ωutils.ϒstr_to_int
		ϒunified_strdate = Ωutils.ϒunified_strdate
		CamdemyIE = λ.Cal(λ.TypeType, λ.StrLiteral("CamdemyIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				CamdemyIE__VALID_URL    λ.Object
				CamdemyIE__real_extract λ.Object
			)
			CamdemyIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?camdemy\\.com/media/(?P<id>\\d+)")
			CamdemyIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdescription   λ.Object
						ϒfile_list_doc λ.Object
						ϒfile_name     λ.Object
						ϒoembed_obj    λ.Object
						ϒself          = λargs[0]
						ϒsrc_from      λ.Object
						ϒthumb_url     λ.Object
						ϒtitle         λ.Object
						ϒupload_date   λ.Object
						ϒurl           = λargs[1]
						ϒvideo_folder  λ.Object
						ϒvideo_id      λ.Object
						ϒvideo_url     λ.Object
						ϒview_count    λ.Object
						ϒwebpage       λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒvideo_id)
					ϒsrc_from = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("class=['\\\"]srcFrom['\\\"][^>]*>Sources?(?:\\s+from)?\\s*:\\s*<a[^>]+(?:href|title)=(['\\\"])(?P<url>(?:(?!\\1).)+)\\1"),
						ϒwebpage,
						λ.StrLiteral("external source"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
						{Name: "group", Value: λ.StrLiteral("url")},
					})
					if λ.IsTrue(ϒsrc_from) {
						return λ.Calm(ϒself, "url_result", ϒsrc_from)
					}
					ϒoembed_obj = λ.Calm(ϒself, "_download_json", λ.Add(λ.StrLiteral("http://www.camdemy.com/oembed/?format=json&url="), ϒurl), ϒvideo_id)
					ϒtitle = λ.GetItem(ϒoembed_obj, λ.StrLiteral("title"))
					ϒthumb_url = λ.GetItem(ϒoembed_obj, λ.StrLiteral("thumbnail_url"))
					ϒvideo_folder = λ.Cal(Ωparse.ϒurljoin, ϒthumb_url, λ.StrLiteral("video/"))
					ϒfile_list_doc = λ.Calm(ϒself, "_download_xml", λ.Cal(Ωparse.ϒurljoin, ϒvideo_folder, λ.StrLiteral("fileList.xml")), ϒvideo_id, λ.StrLiteral("Downloading filelist XML"))
					ϒfile_name = λ.GetAttr(λ.Calm(ϒfile_list_doc, "find", λ.StrLiteral("./video/item/fileName")), "text", nil)
					ϒvideo_url = λ.Cal(Ωparse.ϒurljoin, ϒvideo_folder, ϒfile_name)
					ϒupload_date = λ.Cal(ϒunified_strdate, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral(">published on ([^<]+)<"),
						ϒwebpage,
						λ.StrLiteral("upload date"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					}))
					ϒview_count = λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("role=[\"\\']viewCnt[\"\\'][^>]*>([\\d,.]+) views"),
						ϒwebpage,
						λ.StrLiteral("view count"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					}))
					ϒdescription = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
							λ.StrLiteral("description"),
							ϒwebpage,
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(ϒclean_html, λ.Calm(ϒoembed_obj, "get", λ.StrLiteral("description")))
						}
					}()
					return λ.DictLiteral(map[string]λ.Object{
						"id":          ϒvideo_id,
						"url":         ϒvideo_url,
						"title":       ϒtitle,
						"thumbnail":   ϒthumb_url,
						"description": ϒdescription,
						"creator":     λ.Calm(ϒoembed_obj, "get", λ.StrLiteral("author_name")),
						"duration":    λ.Cal(ϒparse_duration, λ.Calm(ϒoembed_obj, "get", λ.StrLiteral("duration"))),
						"upload_date": ϒupload_date,
						"view_count":  ϒview_count,
					})
				})
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL":    CamdemyIE__VALID_URL,
				"_real_extract": CamdemyIE__real_extract,
			})
		}())
		CamdemyFolderIE = λ.Cal(λ.TypeType, λ.StrLiteral("CamdemyFolderIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				CamdemyFolderIE__VALID_URL λ.Object
			)
			CamdemyFolderIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?camdemy\\.com/folder/(?P<id>\\d+)")
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL": CamdemyFolderIE__VALID_URL,
			})
		}())
	})
}
