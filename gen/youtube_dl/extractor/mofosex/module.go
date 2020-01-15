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
 * mofosex/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/mofosex.py
 */

package mofosex

import (
	Ωkeezmovies "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/keezmovies"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	KeezMoviesIE     λ.Object
	MofosexIE        λ.Object
	ϒint_or_none     λ.Object
	ϒstr_to_int      λ.Object
	ϒunified_strdate λ.Object
)

func init() {
	λ.InitModule(func() {
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒstr_to_int = Ωutils.ϒstr_to_int
		ϒunified_strdate = Ωutils.ϒunified_strdate
		KeezMoviesIE = Ωkeezmovies.KeezMoviesIE
		MofosexIE = λ.Cal(λ.TypeType, λ.StrLiteral("MofosexIE"), λ.NewTuple(KeezMoviesIE), func() λ.Dict {
			var (
				MofosexIE__VALID_URL    λ.Object
				MofosexIE__real_extract λ.Object
			)
			MofosexIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?mofosex\\.com/videos/(?P<id>\\d+)/(?P<display_id>[^/?#&.]+)\\.html")
			MofosexIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdislike_count λ.Object
						ϒinfo          λ.Object
						ϒlike_count    λ.Object
						ϒself          = λargs[0]
						ϒupload_date   λ.Object
						ϒurl           = λargs[1]
						ϒview_count    λ.Object
						ϒwebpage       λ.Object
						τmp0           λ.Object
					)
					τmp0 = λ.Calm(ϒself, "_extract_info", ϒurl)
					ϒwebpage = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒinfo = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒview_count = λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("VIEWS:</span>\\s*([\\d,.]+)"),
						ϒwebpage,
						λ.StrLiteral("view count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒlike_count = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("id=[\"\\']amountLikes[\"\\'][^>]*>(\\d+)"),
						ϒwebpage,
						λ.StrLiteral("like count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒdislike_count = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("id=[\"\\']amountDislikes[\"\\'][^>]*>(\\d+)"),
						ϒwebpage,
						λ.StrLiteral("like count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒupload_date = λ.Cal(ϒunified_strdate, λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("Added:</span>([^<]+)"),
						ϒwebpage,
						λ.StrLiteral("upload date"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					λ.Calm(ϒinfo, "update", λ.DictLiteral(map[string]λ.Object{
						"view_count":    ϒview_count,
						"like_count":    ϒlike_count,
						"dislike_count": ϒdislike_count,
						"upload_date":   ϒupload_date,
						"thumbnail":     λ.Calm(ϒself, "_og_search_thumbnail", ϒwebpage),
					}))
					return ϒinfo
				})
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL":    MofosexIE__VALID_URL,
				"_real_extract": MofosexIE__real_extract,
			})
		}())
	})
}
