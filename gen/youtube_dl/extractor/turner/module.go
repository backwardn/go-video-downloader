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
 * turner/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/turner.py
 */

package turner

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωadobepass "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/adobepass"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	AdobePassIE       λ.Object
	ExtractorError    λ.Object
	TurnerBaseIE      λ.Object
	ϒcompat_str       λ.Object
	ϒdetermine_ext    λ.Object
	ϒfloat_or_none    λ.Object
	ϒint_or_none      λ.Object
	ϒparse_duration   λ.Object
	ϒstrip_or_none    λ.Object
	ϒupdate_url_query λ.Object
	ϒurl_or_none      λ.Object
	ϒxpath_attr       λ.Object
	ϒxpath_text       λ.Object
)

func init() {
	λ.InitModule(func() {
		AdobePassIE = Ωadobepass.AdobePassIE
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒxpath_text = Ωutils.ϒxpath_text
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒxpath_attr = Ωutils.ϒxpath_attr
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		ExtractorError = Ωutils.ExtractorError
		ϒstrip_or_none = Ωutils.ϒstrip_or_none
		ϒurl_or_none = Ωutils.ϒurl_or_none
		TurnerBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("TurnerBaseIE"), λ.NewTuple(AdobePassIE), func() λ.Dict {
			var (
				TurnerBaseIE__extract_cvp_info  λ.Object
				TurnerBaseIE__extract_timestamp λ.Object
			)
			TurnerBaseIE__extract_timestamp = λ.NewFunction("_extract_timestamp",
				[]λ.Param{
					{Name: "self"},
					{Name: "video_data"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself       = λargs[0]
						ϒvideo_data = λargs[1]
					)
					_ = ϒself
					return λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_attr, ϒvideo_data, λ.StrLiteral("dateCreated"), λ.StrLiteral("uts")))
				})
			TurnerBaseIE__extract_cvp_info = λ.NewFunction("_extract_cvp_info",
				[]λ.Param{
					{Name: "self"},
					{Name: "data_src"},
					{Name: "video_id"},
					{Name: "path_data", Def: λ.DictLiteral(map[λ.Object]λ.Object{})},
					{Name: "ap_data", Def: λ.DictLiteral(map[λ.Object]λ.Object{})},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒap_data          = λargs[4]
						ϒbase_path_data   λ.Object
						ϒcontent_id       λ.Object
						ϒdata_src         = λargs[1]
						ϒext              λ.Object
						ϒf                λ.Object
						ϒformat_id        λ.Object
						ϒformats          λ.Object
						ϒis_live          λ.Object
						ϒlang             λ.Object
						ϒm3u8_formats     λ.Object
						ϒmedia_src        λ.Object
						ϒmobj             λ.Object
						ϒpath_data        = λargs[3]
						ϒrex              λ.Object
						ϒsecure_path_data λ.Object
						ϒself             = λargs[0]
						ϒsource           λ.Object
						ϒsubtitles        λ.Object
						ϒthumbnails       λ.Object
						ϒtitle            λ.Object
						ϒtrack            λ.Object
						ϒtrack_url        λ.Object
						ϒurls             λ.Object
						ϒvideo_data       λ.Object
						ϒvideo_file       λ.Object
						ϒvideo_id         = λargs[2]
						ϒvideo_url        λ.Object
						τmp0              λ.Object
						τmp1              λ.Object
						τmp2              λ.Object
						τmp3              λ.Object
					)
					ϒvideo_data = λ.Calm(ϒself, "_download_xml", ϒdata_src, ϒvideo_id)
					ϒvideo_id = λ.GetItem(λ.GetAttr(ϒvideo_data, "attrib", nil), λ.StrLiteral("id"))
					ϒtitle = λ.Call(ϒxpath_text, λ.NewArgs(
						ϒvideo_data,
						λ.StrLiteral("headline"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.True},
					})
					ϒcontent_id = func() λ.Object {
						if λv := λ.Cal(ϒxpath_text, ϒvideo_data, λ.StrLiteral("contentId")); λ.IsTrue(λv) {
							return λv
						} else {
							return ϒvideo_id
						}
					}()
					ϒurls = λ.NewList()
					ϒformats = λ.NewList()
					ϒrex = λ.Cal(Ωre.ϒcompile, λ.StrLiteral("(?P<width>[0-9]+)x(?P<height>[0-9]+)(?:_(?P<bitrate>[0-9]+))?"))
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒvideo_data, "findall", λ.StrLiteral(".//file")))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒvideo_file = τmp1
						ϒvideo_url = λ.Calm(λ.GetAttr(ϒvideo_file, "text", nil), "strip")
						if !λ.IsTrue(ϒvideo_url) {
							continue
						}
						ϒext = λ.Cal(ϒdetermine_ext, ϒvideo_url)
						if λ.IsTrue(λ.Calm(ϒvideo_url, "startswith", λ.StrLiteral("/mp4:protected/"))) {
							continue
						} else {
							if λ.IsTrue(λ.Calm(ϒvideo_url, "startswith", λ.StrLiteral("/secure/"))) {
								ϒsecure_path_data = λ.Calm(ϒpath_data, "get", λ.StrLiteral("secure"))
								if !λ.IsTrue(ϒsecure_path_data) {
									continue
								}
								ϒvideo_url = λ.Calm(ϒself, "_add_akamai_spe_token", λ.GetItem(ϒsecure_path_data, λ.StrLiteral("tokenizer_src")), λ.Add(λ.GetItem(ϒsecure_path_data, λ.StrLiteral("media_src")), ϒvideo_url), ϒcontent_id, ϒap_data)
							} else {
								if !λ.IsTrue(λ.Cal(Ωre.ϒmatch, λ.StrLiteral("https?://"), ϒvideo_url)) {
									ϒbase_path_data = λ.Calm(ϒpath_data, "get", ϒext, λ.Calm(ϒpath_data, "get", λ.StrLiteral("default"), λ.DictLiteral(map[λ.Object]λ.Object{})))
									ϒmedia_src = λ.Calm(ϒbase_path_data, "get", λ.StrLiteral("media_src"))
									if !λ.IsTrue(ϒmedia_src) {
										continue
									}
									ϒvideo_url = λ.Add(ϒmedia_src, ϒvideo_url)
								}
							}
						}
						if λ.Contains(ϒurls, ϒvideo_url) {
							continue
						}
						λ.Calm(ϒurls, "append", ϒvideo_url)
						ϒformat_id = λ.Calm(ϒvideo_file, "get", λ.StrLiteral("bitrate"))
						if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("smil"))) {
							λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_smil_formats", nil), λ.NewArgs(
								ϒvideo_url,
								ϒvideo_id,
							), λ.KWArgs{
								{Name: "fatal", Value: λ.False},
							}))
						} else {
							if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("m3u8"))) {
								ϒm3u8_formats = λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
									ϒvideo_url,
									ϒvideo_id,
									λ.StrLiteral("mp4"),
								), λ.KWArgs{
									{Name: "m3u8_id", Value: func() λ.Object {
										if λv := ϒformat_id; λ.IsTrue(λv) {
											return λv
										} else {
											return λ.StrLiteral("hls")
										}
									}()},
									{Name: "fatal", Value: λ.False},
								})
								if λ.IsTrue(func() λ.Object {
									if λv := λ.NewBool(λ.Contains(ϒvideo_url, λ.StrLiteral("/secure/"))); !λ.IsTrue(λv) {
										return λv
									} else {
										return λ.NewBool(λ.Contains(ϒvideo_url, λ.StrLiteral("?hdnea=")))
									}
								}()) {
									τmp2 = λ.Cal(λ.BuiltinIter, ϒm3u8_formats)
									for {
										if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
											break
										}
										ϒf = τmp3
										λ.SetItem(ϒf, λ.StrLiteral("_seekable"), λ.False)
									}
								}
								λ.Calm(ϒformats, "extend", ϒm3u8_formats)
							} else {
								if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("f4m"))) {
									λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
										λ.Cal(ϒupdate_url_query, ϒvideo_url, λ.DictLiteral(map[string]string{
											"hdcore": "3.7.0",
										})),
										ϒvideo_id,
									), λ.KWArgs{
										{Name: "f4m_id", Value: func() λ.Object {
											if λv := ϒformat_id; λ.IsTrue(λv) {
												return λv
											} else {
												return λ.StrLiteral("hds")
											}
										}()},
										{Name: "fatal", Value: λ.False},
									}))
								} else {
									ϒf = λ.DictLiteral(map[string]λ.Object{
										"format_id": ϒformat_id,
										"url":       ϒvideo_url,
										"ext":       ϒext,
									})
									ϒmobj = λ.Calm(ϒrex, "search", λ.Add(ϒformat_id, ϒvideo_url))
									if λ.IsTrue(ϒmobj) {
										λ.Calm(ϒf, "update", λ.DictLiteral(map[string]λ.Object{
											"width":  λ.Cal(λ.IntType, λ.Calm(ϒmobj, "group", λ.StrLiteral("width"))),
											"height": λ.Cal(λ.IntType, λ.Calm(ϒmobj, "group", λ.StrLiteral("height"))),
											"tbr":    λ.Cal(ϒint_or_none, λ.Calm(ϒmobj, "group", λ.StrLiteral("bitrate"))),
										}))
									} else {
										if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒformat_id, ϒcompat_str)) {
											if λ.IsTrue(λ.Calm(ϒformat_id, "isdigit")) {
												λ.SetItem(ϒf, λ.StrLiteral("tbr"), λ.Cal(λ.IntType, ϒformat_id))
											} else {
												ϒmobj = λ.Cal(Ωre.ϒmatch, λ.StrLiteral("ios_(audio|[0-9]+)$"), ϒformat_id)
												if λ.IsTrue(ϒmobj) {
													if λ.IsTrue(λ.Eq(λ.Calm(ϒmobj, "group", λ.IntLiteral(1)), λ.StrLiteral("audio"))) {
														λ.Calm(ϒf, "update", λ.DictLiteral(map[string]string{
															"vcodec": "none",
															"ext":    "m4a",
														}))
													} else {
														λ.SetItem(ϒf, λ.StrLiteral("tbr"), λ.Cal(λ.IntType, λ.Calm(ϒmobj, "group", λ.IntLiteral(1))))
													}
												}
											}
										}
									}
									λ.Calm(ϒformats, "append", ϒf)
								}
							}
						}
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒsubtitles = λ.DictLiteral(map[λ.Object]λ.Object{})
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒvideo_data, "findall", λ.StrLiteral("closedCaptions/source")))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒsource = τmp1
						τmp2 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒsource, "findall", λ.StrLiteral("track")))
						for {
							if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
								break
							}
							ϒtrack = τmp3
							ϒtrack_url = λ.Cal(ϒurl_or_none, λ.Calm(ϒtrack, "get", λ.StrLiteral("url")))
							if λ.IsTrue(func() λ.Object {
								if λv := λ.NewBool(!λ.IsTrue(ϒtrack_url)); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Calm(ϒtrack_url, "endswith", λ.StrLiteral("/big"))
								}
							}()) {
								continue
							}
							ϒlang = func() λ.Object {
								if λv := λ.Calm(ϒtrack, "get", λ.StrLiteral("lang")); λ.IsTrue(λv) {
									return λv
								} else if λv := λ.Calm(ϒtrack, "get", λ.StrLiteral("label")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.StrLiteral("en")
								}
							}()
							λ.Calm(λ.Calm(ϒsubtitles, "setdefault", ϒlang, λ.NewList()), "append", λ.DictLiteral(map[string]λ.Object{
								"url": ϒtrack_url,
								"ext": λ.Calm(λ.DictLiteral(map[string]string{
									"scc":     "scc",
									"webvtt":  "vtt",
									"smptett": "tt",
								}), "get", λ.Calm(ϒsource, "get", λ.StrLiteral("format"))),
							}))
						}
					}
					ϒthumbnails = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒimage λ.Object
									τmp0   λ.Object
									τmp1   λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒvideo_data, "findall", λ.StrLiteral("images/image")))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒimage = τmp1
									λgy.Yield(λ.DictLiteral(map[string]λ.Object{
										"id":     λ.Calm(ϒimage, "get", λ.StrLiteral("cut")),
										"url":    λ.GetAttr(ϒimage, "text", nil),
										"width":  λ.Cal(ϒint_or_none, λ.Calm(ϒimage, "get", λ.StrLiteral("width"))),
										"height": λ.Cal(ϒint_or_none, λ.Calm(ϒimage, "get", λ.StrLiteral("height"))),
									}))
								}
								return λ.None
							})
						})))
					ϒis_live = λ.Eq(λ.Cal(ϒxpath_text, ϒvideo_data, λ.StrLiteral("isLive")), λ.StrLiteral("true"))
					return λ.DictLiteral(map[string]λ.Object{
						"id": ϒvideo_id,
						"title": func() λ.Object {
							if λ.IsTrue(ϒis_live) {
								return λ.Calm(ϒself, "_live_title", ϒtitle)
							} else {
								return ϒtitle
							}
						}(),
						"formats":     ϒformats,
						"subtitles":   ϒsubtitles,
						"thumbnails":  ϒthumbnails,
						"thumbnail":   λ.Cal(ϒxpath_text, ϒvideo_data, λ.StrLiteral("poster")),
						"description": λ.Cal(ϒstrip_or_none, λ.Cal(ϒxpath_text, ϒvideo_data, λ.StrLiteral("description"))),
						"duration": λ.Cal(ϒparse_duration, func() λ.Object {
							if λv := λ.Cal(ϒxpath_text, ϒvideo_data, λ.StrLiteral("length")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(ϒxpath_text, ϒvideo_data, λ.StrLiteral("trt"))
							}
						}()),
						"timestamp":      λ.Calm(ϒself, "_extract_timestamp", ϒvideo_data),
						"upload_date":    λ.Cal(ϒxpath_attr, ϒvideo_data, λ.StrLiteral("metas"), λ.StrLiteral("version")),
						"series":         λ.Cal(ϒxpath_text, ϒvideo_data, λ.StrLiteral("showTitle")),
						"season_number":  λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒvideo_data, λ.StrLiteral("seasonNumber"))),
						"episode_number": λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒvideo_data, λ.StrLiteral("episodeNumber"))),
						"is_live":        ϒis_live,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_extract_cvp_info":  TurnerBaseIE__extract_cvp_info,
				"_extract_timestamp": TurnerBaseIE__extract_timestamp,
			})
		}())
	})
}
