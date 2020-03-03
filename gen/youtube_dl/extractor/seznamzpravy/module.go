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
 * seznamzpravy/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/seznamzpravy.py
 */

package seznamzpravy

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor                 λ.Object
	SeznamZpravyArticleIE         λ.Object
	SeznamZpravyIE                λ.Object
	ϒ_raw_id                      λ.Object
	ϒcompat_parse_qs              λ.Object
	ϒcompat_str                   λ.Object
	ϒcompat_urllib_parse_urlparse λ.Object
	ϒint_or_none                  λ.Object
	ϒparse_codecs                 λ.Object
	ϒtry_get                      λ.Object
	ϒurljoin                      λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_parse_qs = Ωcompat.ϒcompat_parse_qs
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒcompat_urllib_parse_urlparse = Ωcompat.ϒcompat_urllib_parse_urlparse
		ϒurljoin = Ωutils.ϒurljoin
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_codecs = Ωutils.ϒparse_codecs
		ϒtry_get = Ωutils.ϒtry_get
		ϒ_raw_id = λ.NewFunction("_raw_id",
			[]λ.Param{
				{Name: "src_url"},
			},
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				var (
					ϒsrc_url = λargs[0]
				)
				return λ.GetItem(λ.Calm(λ.GetAttr(λ.Cal(ϒcompat_urllib_parse_urlparse, ϒsrc_url), "path", nil), "split", λ.StrLiteral("/")), λ.Neg(λ.IntLiteral(1)))
			})
		SeznamZpravyIE = λ.Cal(λ.TypeType, λ.StrLiteral("SeznamZpravyIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				SeznamZpravyIE__VALID_URL           λ.Object
				SeznamZpravyIE__extract_sdn_formats λ.Object
				SeznamZpravyIE__real_extract        λ.Object
			)
			SeznamZpravyIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?seznamzpravy\\.cz/iframe/player\\?.*\\bsrc=")
			SeznamZpravyIE__extract_sdn_formats = λ.NewFunction("_extract_sdn_formats",
				[]λ.Param{
					{Name: "self"},
					{Name: "sdn_url"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdash_rel_url λ.Object
						ϒf            λ.Object
						ϒformat_data  λ.Object
						ϒformat_id    λ.Object
						ϒformats      λ.Object
						ϒget_url      λ.Object
						ϒheight       λ.Object
						ϒhls_rel_url  λ.Object
						ϒmp4_formats  λ.Object
						ϒpls          λ.Object
						ϒrelative_url λ.Object
						ϒsdn_data     λ.Object
						ϒsdn_url      = λargs[1]
						ϒself         = λargs[0]
						ϒvideo_id     = λargs[2]
						ϒwidth        λ.Object
						τmp0          λ.Object
						τmp1          λ.Object
						τmp2          λ.Object
						τmp3          λ.Object
						τmp4          λ.Object
					)
					_ = τmp2
					_ = τmp3
					ϒsdn_data = λ.Calm(ϒself, "_download_json", ϒsdn_url, ϒvideo_id)
					if λ.IsTrue(λ.Calm(ϒsdn_data, "get", λ.StrLiteral("Location"))) {
						ϒsdn_url = λ.GetItem(ϒsdn_data, λ.StrLiteral("Location"))
						ϒsdn_data = λ.Calm(ϒself, "_download_json", ϒsdn_url, ϒvideo_id)
					}
					ϒformats = λ.NewList()
					ϒmp4_formats = func() λ.Object {
						if λv := λ.Cal(ϒtry_get, ϒsdn_data, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("data")), λ.StrLiteral("mp4"))
							}), λ.DictType); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.DictLiteral(map[λ.Object]λ.Object{})
						}
					}()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒmp4_formats, "items"))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒformat_id = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒformat_data = λ.GetItem(τmp2, λ.IntLiteral(1))
						ϒrelative_url = λ.Calm(ϒformat_data, "get", λ.StrLiteral("url"))
						if !λ.IsTrue(ϒrelative_url) {
							continue
						}
						τmp2, τmp3 = func() (λexit λ.Object, λret λ.Object) {
							defer λ.CatchMulti(
								nil,
								&λ.Catcher{λ.NewTuple(
									λ.TypeErrorType,
									λ.ValueErrorType,
								), func(λex λ.BaseException) {
									τmp4 = λ.NewTuple(
										λ.None,
										λ.None,
									)
									ϒwidth = λ.GetItem(τmp4, λ.IntLiteral(0))
									ϒheight = λ.GetItem(τmp4, λ.IntLiteral(1))
								}},
							)
							τmp4 = λ.Calm(ϒformat_data, "get", λ.StrLiteral("resolution"))
							ϒwidth = λ.GetItem(τmp4, λ.IntLiteral(0))
							ϒheight = λ.GetItem(τmp4, λ.IntLiteral(1))
							return λ.BlockExitNormally, nil
						}()
						ϒf = λ.DictLiteral(map[string]λ.Object{
							"url":       λ.Cal(ϒurljoin, ϒsdn_url, ϒrelative_url),
							"format_id": λ.Mod(λ.StrLiteral("http-%s"), ϒformat_id),
							"tbr": λ.Call(ϒint_or_none, λ.NewArgs(λ.Calm(ϒformat_data, "get", λ.StrLiteral("bandwidth"))), λ.KWArgs{
								{Name: "scale", Value: λ.IntLiteral(1000)},
							}),
							"width":  λ.Cal(ϒint_or_none, ϒwidth),
							"height": λ.Cal(ϒint_or_none, ϒheight),
						})
						λ.Calm(ϒf, "update", λ.Cal(ϒparse_codecs, λ.Calm(ϒformat_data, "get", λ.StrLiteral("codec"))))
						λ.Calm(ϒformats, "append", ϒf)
					}
					ϒpls = λ.Calm(ϒsdn_data, "get", λ.StrLiteral("pls"), λ.DictLiteral(map[λ.Object]λ.Object{}))
					ϒget_url = λ.NewFunction("get_url",
						[]λ.Param{
							{Name: "format_id"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒformat_id = λargs[0]
							)
							return λ.Cal(ϒtry_get, ϒpls, λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(λ.GetItem(ϒx, ϒformat_id), λ.StrLiteral("url"))
								}), ϒcompat_str)
						})
					ϒdash_rel_url = λ.Cal(ϒget_url, λ.StrLiteral("dash"))
					if λ.IsTrue(ϒdash_rel_url) {
						λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_mpd_formats", nil), λ.NewArgs(
							λ.Cal(ϒurljoin, ϒsdn_url, ϒdash_rel_url),
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "mpd_id", Value: λ.StrLiteral("dash")},
							{Name: "fatal", Value: λ.False},
						}))
					}
					ϒhls_rel_url = λ.Cal(ϒget_url, λ.StrLiteral("hls"))
					if λ.IsTrue(ϒhls_rel_url) {
						λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
							λ.Cal(ϒurljoin, ϒsdn_url, ϒhls_rel_url),
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "ext", Value: λ.StrLiteral("mp4")},
							{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
							{Name: "fatal", Value: λ.False},
						}))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					return ϒformats
				})
			SeznamZpravyIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒduration  λ.Object
						ϒformats   λ.Object
						ϒparams    λ.Object
						ϒself      = λargs[0]
						ϒseries    λ.Object
						ϒsrc       λ.Object
						ϒthumbnail λ.Object
						ϒtitle     λ.Object
						ϒurl       = λargs[1]
						ϒvideo_id  λ.Object
					)
					ϒparams = λ.Cal(ϒcompat_parse_qs, λ.GetAttr(λ.Cal(ϒcompat_urllib_parse_urlparse, ϒurl), "query", nil))
					ϒsrc = λ.GetItem(λ.GetItem(ϒparams, λ.StrLiteral("src")), λ.IntLiteral(0))
					ϒtitle = λ.GetItem(λ.GetItem(ϒparams, λ.StrLiteral("title")), λ.IntLiteral(0))
					ϒvideo_id = λ.GetItem(λ.Calm(ϒparams, "get", λ.StrLiteral("contentId"), λ.NewList(λ.Cal(ϒ_raw_id, ϒsrc))), λ.IntLiteral(0))
					ϒformats = λ.Calm(ϒself, "_extract_sdn_formats", λ.Add(ϒsrc, λ.StrLiteral("spl2,2,VOD")), ϒvideo_id)
					ϒduration = λ.Cal(ϒint_or_none, λ.GetItem(λ.Calm(ϒparams, "get", λ.StrLiteral("duration"), λ.NewList(λ.None)), λ.IntLiteral(0)))
					ϒseries = λ.GetItem(λ.Calm(ϒparams, "get", λ.StrLiteral("series"), λ.NewList(λ.None)), λ.IntLiteral(0))
					ϒthumbnail = λ.GetItem(λ.Calm(ϒparams, "get", λ.StrLiteral("poster"), λ.NewList(λ.None)), λ.IntLiteral(0))
					return λ.DictLiteral(map[string]λ.Object{
						"id":        ϒvideo_id,
						"title":     ϒtitle,
						"thumbnail": ϒthumbnail,
						"duration":  ϒduration,
						"series":    ϒseries,
						"formats":   ϒformats,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":           SeznamZpravyIE__VALID_URL,
				"_extract_sdn_formats": SeznamZpravyIE__extract_sdn_formats,
				"_real_extract":        SeznamZpravyIE__real_extract,
			})
		}())
		SeznamZpravyArticleIE = λ.Cal(λ.TypeType, λ.StrLiteral("SeznamZpravyArticleIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				SeznamZpravyArticleIE__VALID_URL λ.Object
			)
			SeznamZpravyArticleIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?(?:seznam\\.cz/zpravy|seznamzpravy\\.cz)/clanek/(?:[^/?#&]+)-(?P<id>\\d+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": SeznamZpravyArticleIE__VALID_URL,
			})
		}())
	})
}
