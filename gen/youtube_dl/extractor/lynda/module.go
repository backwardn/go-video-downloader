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
 * lynda/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/lynda.py
 */

package lynda

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError      λ.Object
	InfoExtractor       λ.Object
	LyndaBaseIE         λ.Object
	LyndaCourseIE       λ.Object
	LyndaIE             λ.Object
	ϒcompat_str         λ.Object
	ϒint_or_none        λ.Object
	ϒurlencode_postdata λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒurlencode_postdata = Ωutils.ϒurlencode_postdata
		LyndaBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("LyndaBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				LyndaBaseIE__NETRC_MACHINE   λ.Object
				LyndaBaseIE__login           λ.Object
				LyndaBaseIE__real_initialize λ.Object
			)
			LyndaBaseIE__NETRC_MACHINE = λ.StrLiteral("lynda")
			LyndaBaseIE__real_initialize = λ.NewFunction("_real_initialize",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
					)
					λ.Calm(ϒself, "_login")
					return λ.None
				})
			LyndaBaseIE__login = λ.NewFunction("_login",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒpassword      λ.Object
						ϒpassword_form λ.Object
						ϒself          = λargs[0]
						ϒsignin_form   λ.Object
						ϒsignin_page   λ.Object
						ϒsignin_url    λ.Object
						ϒusername      λ.Object
						τmp0           λ.Object
					)
					τmp0 = λ.Calm(ϒself, "_get_login_info")
					ϒusername = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒpassword = λ.GetItem(τmp0, λ.IntLiteral(1))
					if ϒusername == λ.None {
						return λ.None
					}
					ϒsignin_page = λ.Calm(ϒself, "_download_webpage", λ.GetAttr(ϒself, "_SIGNIN_URL", nil), λ.None, λ.StrLiteral("Downloading signin page"))
					if λ.IsTrue(λ.Cal(λ.BuiltinAny, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒp   λ.Object
									τmp0 λ.Object
									τmp1 λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
									λ.StrLiteral("isLoggedIn\\s*:\\s*true"),
									λ.StrLiteral("logout\\.aspx"),
									λ.StrLiteral(">Log out<"),
								))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒp = τmp1
									λgy.Yield(λ.Cal(Ωre.ϒsearch, ϒp, ϒsignin_page))
								}
								return λ.None
							})
						})))) {
						return λ.None
					}
					ϒsignin_form = λ.Calm(ϒself, "_search_regex", λ.StrLiteral("(?s)(<form[^>]+data-form-name=[\"\\']signin[\"\\'][^>]*>.+?</form>)"), ϒsignin_page, λ.StrLiteral("signin form"))
					τmp0 = λ.Calm(ϒself, "_login_step", ϒsignin_form, λ.GetAttr(ϒself, "_PASSWORD_URL", nil), λ.DictLiteral(map[string]λ.Object{
						"email": ϒusername,
					}), λ.StrLiteral("Submitting email"), λ.GetAttr(ϒself, "_SIGNIN_URL", nil))
					ϒsignin_page = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒsignin_url = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒpassword_form = λ.GetItem(ϒsignin_page, λ.StrLiteral("body"))
					λ.Calm(ϒself, "_login_step", ϒpassword_form, λ.GetAttr(ϒself, "_USER_URL", nil), λ.DictLiteral(map[string]λ.Object{
						"email":    ϒusername,
						"password": ϒpassword,
					}), λ.StrLiteral("Submitting password"), ϒsignin_url)
					return λ.None
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_NETRC_MACHINE":   LyndaBaseIE__NETRC_MACHINE,
				"_login":           LyndaBaseIE__login,
				"_real_initialize": LyndaBaseIE__real_initialize,
			})
		}())
		LyndaIE = λ.Cal(λ.TypeType, λ.StrLiteral("LyndaIE"), λ.NewTuple(LyndaBaseIE), func() λ.Dict {
			var (
				LyndaIE_IE_NAME       λ.Object
				LyndaIE__VALID_URL    λ.Object
				LyndaIE__real_extract λ.Object
			)
			LyndaIE_IE_NAME = λ.StrLiteral("lynda")
			LyndaIE__VALID_URL = λ.StrLiteral("(?x)\n                    https?://\n                        (?:www\\.)?(?:lynda\\.com|educourse\\.ga)/\n                        (?:\n                            (?:[^/]+/){2,3}(?P<course_id>\\d+)|\n                            player/embed\n                        )/\n                        (?P<id>\\d+)\n                    ")
			LyndaIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcdn                   λ.Object
						ϒconviva               λ.Object
						ϒcourse_id             λ.Object
						ϒduration              λ.Object
						ϒfmts                  λ.Object
						ϒformat_id             λ.Object
						ϒformat_url            λ.Object
						ϒformats               λ.Object
						ϒformats_dict          λ.Object
						ϒmobj                  λ.Object
						ϒplay                  λ.Object
						ϒprioritized_stream    λ.Object
						ϒprioritized_stream_id λ.Object
						ϒprioritized_streams   λ.Object
						ϒquery                 λ.Object
						ϒself                  = λargs[0]
						ϒsubtitles             λ.Object
						ϒtitle                 λ.Object
						ϒurl                   = λargs[1]
						ϒurls                  λ.Object
						ϒvideo                 λ.Object
						ϒvideo_id              λ.Object
						τmp0                   λ.Object
						τmp1                   λ.Object
						τmp2                   λ.Object
						τmp3                   λ.Object
						τmp4                   λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("id"))
					ϒcourse_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("course_id"))
					ϒquery = λ.DictLiteral(map[string]λ.Object{
						"videoId": ϒvideo_id,
						"type":    λ.StrLiteral("video"),
					})
					ϒvideo = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.StrLiteral("https://www.lynda.com/ajax/player"),
						ϒvideo_id,
						λ.StrLiteral("Downloading video JSON"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
						{Name: "query", Value: ϒquery},
					})
					if !λ.IsTrue(ϒvideo) {
						λ.SetItem(ϒquery, λ.StrLiteral("courseId"), ϒcourse_id)
						ϒplay = λ.Calm(ϒself, "_download_json", λ.Mod(λ.StrLiteral("https://www.lynda.com/ajax/course/%s/%s/play"), λ.NewTuple(
							ϒcourse_id,
							ϒvideo_id,
						)), ϒvideo_id, λ.StrLiteral("Downloading play JSON"))
						if !λ.IsTrue(ϒplay) {
							λ.Calm(ϒself, "_raise_unavailable", ϒvideo_id)
						}
						ϒformats = λ.NewList()
						τmp0 = λ.Cal(λ.BuiltinIter, ϒplay)
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒformats_dict = τmp1
							ϒurls = λ.Calm(ϒformats_dict, "get", λ.StrLiteral("urls"))
							if !λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒurls, λ.DictType)) {
								continue
							}
							ϒcdn = λ.Calm(ϒformats_dict, "get", λ.StrLiteral("name"))
							τmp2 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒurls, "items"))
							for {
								if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
									break
								}
								τmp4 = τmp3
								ϒformat_id = λ.GetItem(τmp4, λ.IntLiteral(0))
								ϒformat_url = λ.GetItem(τmp4, λ.IntLiteral(1))
								if !λ.IsTrue(ϒformat_url) {
									continue
								}
								λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
									"url": ϒformat_url,
									"format_id": func() λ.Object {
										if λ.IsTrue(ϒcdn) {
											return λ.Mod(λ.StrLiteral("%s-%s"), λ.NewTuple(
												ϒcdn,
												ϒformat_id,
											))
										} else {
											return ϒformat_id
										}
									}(),
									"height": λ.Cal(ϒint_or_none, ϒformat_id),
								}))
							}
						}
						λ.Calm(ϒself, "_sort_formats", ϒformats)
						ϒconviva = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
							λ.StrLiteral("https://www.lynda.com/ajax/player/conviva"),
							ϒvideo_id,
							λ.StrLiteral("Downloading conviva JSON"),
						), λ.KWArgs{
							{Name: "query", Value: ϒquery},
						})
						return λ.DictLiteral(map[string]λ.Object{
							"id":           ϒvideo_id,
							"title":        λ.GetItem(ϒconviva, λ.StrLiteral("VideoTitle")),
							"description":  λ.Calm(ϒconviva, "get", λ.StrLiteral("VideoDescription")),
							"release_year": λ.Cal(ϒint_or_none, λ.Calm(ϒconviva, "get", λ.StrLiteral("ReleaseYear"))),
							"duration":     λ.Cal(ϒint_or_none, λ.Calm(ϒconviva, "get", λ.StrLiteral("Duration"))),
							"creator":      λ.Calm(ϒconviva, "get", λ.StrLiteral("Author")),
							"formats":      ϒformats,
						})
					}
					if λ.Contains(ϒvideo, λ.StrLiteral("Status")) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.StrLiteral("lynda returned error: %s"), λ.GetItem(ϒvideo, λ.StrLiteral("Message")))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					if λ.Calm(ϒvideo, "get", λ.StrLiteral("HasAccess")) == λ.False {
						λ.Calm(ϒself, "_raise_unavailable", ϒvideo_id)
					}
					ϒvideo_id = λ.Cal(ϒcompat_str, func() λ.Object {
						if λv := λ.Calm(ϒvideo, "get", λ.StrLiteral("ID")); λ.IsTrue(λv) {
							return λv
						} else {
							return ϒvideo_id
						}
					}())
					ϒduration = λ.Cal(ϒint_or_none, λ.Calm(ϒvideo, "get", λ.StrLiteral("DurationInSeconds")))
					ϒtitle = λ.GetItem(ϒvideo, λ.StrLiteral("Title"))
					ϒformats = λ.NewList()
					ϒfmts = λ.Calm(ϒvideo, "get", λ.StrLiteral("Formats"))
					if λ.IsTrue(ϒfmts) {
						λ.Calm(ϒformats, "extend", λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
							nil,
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
									var (
										ϒf   λ.Object
										τmp0 λ.Object
										τmp1 λ.Object
									)
									τmp0 = λ.Cal(λ.BuiltinIter, ϒfmts)
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										ϒf = τmp1
										if λ.IsTrue(λ.Calm(ϒf, "get", λ.StrLiteral("Url"))) {
											λgy.Yield(λ.DictLiteral(map[string]λ.Object{
												"url":      λ.GetItem(ϒf, λ.StrLiteral("Url")),
												"ext":      λ.Calm(ϒf, "get", λ.StrLiteral("Extension")),
												"width":    λ.Cal(ϒint_or_none, λ.Calm(ϒf, "get", λ.StrLiteral("Width"))),
												"height":   λ.Cal(ϒint_or_none, λ.Calm(ϒf, "get", λ.StrLiteral("Height"))),
												"filesize": λ.Cal(ϒint_or_none, λ.Calm(ϒf, "get", λ.StrLiteral("FileSize"))),
												"format_id": func() λ.Object {
													if λ.IsTrue(λ.Calm(ϒf, "get", λ.StrLiteral("Resolution"))) {
														return λ.Cal(ϒcompat_str, λ.Calm(ϒf, "get", λ.StrLiteral("Resolution")))
													} else {
														return λ.None
													}
												}(),
											}))
										}
									}
									return λ.None
								})
							}))))
					}
					ϒprioritized_streams = λ.Calm(ϒvideo, "get", λ.StrLiteral("PrioritizedStreams"))
					if λ.IsTrue(ϒprioritized_streams) {
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒprioritized_streams, "items"))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							τmp2 = τmp1
							ϒprioritized_stream_id = λ.GetItem(τmp2, λ.IntLiteral(0))
							ϒprioritized_stream = λ.GetItem(τmp2, λ.IntLiteral(1))
							λ.Calm(ϒformats, "extend", λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
								nil,
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
										var (
											ϒformat_id λ.Object
											ϒvideo_url λ.Object
											τmp0       λ.Object
											τmp1       λ.Object
											τmp2       λ.Object
										)
										τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒprioritized_stream, "items"))
										for {
											if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
												break
											}
											τmp2 = τmp1
											ϒformat_id = λ.GetItem(τmp2, λ.IntLiteral(0))
											ϒvideo_url = λ.GetItem(τmp2, λ.IntLiteral(1))
											λgy.Yield(λ.DictLiteral(map[string]λ.Object{
												"url":    ϒvideo_url,
												"height": λ.Cal(ϒint_or_none, ϒformat_id),
												"format_id": λ.Mod(λ.StrLiteral("%s-%s"), λ.NewTuple(
													ϒprioritized_stream_id,
													ϒformat_id,
												)),
											}))
										}
										return λ.None
									})
								}))))
						}
					}
					λ.Calm(ϒself, "_check_formats", ϒformats, ϒvideo_id)
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒsubtitles = λ.Calm(ϒself, "extract_subtitles", ϒvideo_id)
					return λ.DictLiteral(map[string]λ.Object{
						"id":        ϒvideo_id,
						"title":     ϒtitle,
						"duration":  ϒduration,
						"subtitles": ϒsubtitles,
						"formats":   ϒformats,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":       LyndaIE_IE_NAME,
				"_VALID_URL":    LyndaIE__VALID_URL,
				"_real_extract": LyndaIE__real_extract,
			})
		}())
		LyndaCourseIE = λ.Cal(λ.TypeType, λ.StrLiteral("LyndaCourseIE"), λ.NewTuple(LyndaBaseIE), func() λ.Dict {
			var (
				LyndaCourseIE__VALID_URL λ.Object
			)
			LyndaCourseIE__VALID_URL = λ.StrLiteral("https?://(?:www|m)\\.(?:lynda\\.com|educourse\\.ga)/(?P<coursepath>(?:[^/]+/){2,3}(?P<courseid>\\d+))-2\\.html")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": LyndaCourseIE__VALID_URL,
			})
		}())
	})
}
