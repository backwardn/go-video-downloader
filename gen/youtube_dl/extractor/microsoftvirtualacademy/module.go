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
 * microsoftvirtualacademy/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/microsoftvirtualacademy.py
 */

package microsoftvirtualacademy

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor                   λ.Object
	MicrosoftVirtualAcademyBaseIE   λ.Object
	MicrosoftVirtualAcademyCourseIE λ.Object
	MicrosoftVirtualAcademyIE       λ.Object
	ϒcompat_xpath                   λ.Object
	ϒint_or_none                    λ.Object
	ϒparse_duration                 λ.Object
	ϒsmuggle_url                    λ.Object
	ϒunsmuggle_url                  λ.Object
	ϒxpath_text                     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_xpath = Ωcompat.ϒcompat_xpath
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒsmuggle_url = Ωutils.ϒsmuggle_url
		ϒunsmuggle_url = Ωutils.ϒunsmuggle_url
		ϒxpath_text = Ωutils.ϒxpath_text
		MicrosoftVirtualAcademyBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("MicrosoftVirtualAcademyBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				MicrosoftVirtualAcademyBaseIE__extract_base_url          λ.Object
				MicrosoftVirtualAcademyBaseIE__extract_chapter_and_title λ.Object
			)
			MicrosoftVirtualAcademyBaseIE__extract_base_url = λ.NewFunction("_extract_base_url",
				[]λ.Param{
					{Name: "self"},
					{Name: "course_id"},
					{Name: "display_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcourse_id  = λargs[1]
						ϒdisplay_id = λargs[2]
						ϒself       = λargs[0]
					)
					return λ.Calm(ϒself, "_download_json", λ.Mod(λ.StrLiteral("https://api-mlxprod.microsoft.com/services/products/anonymous/%s"), ϒcourse_id), ϒdisplay_id, λ.StrLiteral("Downloading course base URL"))
				})
			MicrosoftVirtualAcademyBaseIE__extract_chapter_and_title = λ.NewFunction("_extract_chapter_and_title",
				[]λ.Param{
					{Name: "self"},
					{Name: "title"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒm     λ.Object
						ϒself  = λargs[0]
						ϒtitle = λargs[1]
					)
					_ = ϒself
					if !λ.IsTrue(ϒtitle) {
						return λ.NewTuple(
							λ.None,
							λ.None,
						)
					}
					ϒm = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("(?P<chapter>\\d+)\\s*\\|\\s*(?P<title>.+)"), ϒtitle)
					return func() λ.Object {
						if λ.IsTrue(ϒm) {
							return λ.NewTuple(
								λ.Cal(λ.IntType, λ.Calm(ϒm, "group", λ.StrLiteral("chapter"))),
								λ.Calm(ϒm, "group", λ.StrLiteral("title")),
							)
						} else {
							return λ.NewTuple(
								λ.None,
								ϒtitle,
							)
						}
					}()
				})
			return λ.DictLiteral(map[string]λ.Object{
				"_extract_base_url":          MicrosoftVirtualAcademyBaseIE__extract_base_url,
				"_extract_chapter_and_title": MicrosoftVirtualAcademyBaseIE__extract_chapter_and_title,
			})
		}())
		MicrosoftVirtualAcademyIE = λ.Cal(λ.TypeType, λ.StrLiteral("MicrosoftVirtualAcademyIE"), λ.NewTuple(MicrosoftVirtualAcademyBaseIE), func() λ.Dict {
			var (
				MicrosoftVirtualAcademyIE_IE_NAME       λ.Object
				MicrosoftVirtualAcademyIE__VALID_URL    λ.Object
				MicrosoftVirtualAcademyIE__real_extract λ.Object
			)
			MicrosoftVirtualAcademyIE_IE_NAME = λ.StrLiteral("mva")
			MicrosoftVirtualAcademyIE__VALID_URL = λ.Mod(λ.StrLiteral("(?:%s:|https?://(?:mva\\.microsoft|(?:www\\.)?microsoftvirtualacademy)\\.com/[^/]+/training-courses/[^/?#&]+-)(?P<course_id>\\d+)(?::|\\?l=)(?P<id>[\\da-zA-Z]+_\\d+)"), MicrosoftVirtualAcademyIE_IE_NAME)
			MicrosoftVirtualAcademyIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒacodec        λ.Object
						ϒbase_url      λ.Object
						ϒcodec         λ.Object
						ϒcodecs        λ.Object
						ϒcourse_id     λ.Object
						ϒformats       λ.Object
						ϒheight        λ.Object
						ϒmobj          λ.Object
						ϒself          = λargs[0]
						ϒsettings      λ.Object
						ϒsmuggled_data λ.Object
						ϒsource        λ.Object
						ϒsources       λ.Object
						ϒsources_type  λ.Object
						ϒsubtitle_url  λ.Object
						ϒsubtitles     λ.Object
						ϒtitle         λ.Object
						ϒurl           = λargs[1]
						ϒvcodec        λ.Object
						ϒvideo_id      λ.Object
						ϒvideo_mode    λ.Object
						ϒvideo_url     λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
						τmp2           λ.Object
						τmp3           λ.Object
						τmp4           λ.Object
					)
					τmp0 = λ.Cal(ϒunsmuggle_url, ϒurl, λ.DictLiteral(map[λ.Object]λ.Object{}))
					ϒurl = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒsmuggled_data = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒcourse_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("course_id"))
					ϒvideo_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("id"))
					ϒbase_url = func() λ.Object {
						if λv := λ.Calm(ϒsmuggled_data, "get", λ.StrLiteral("base_url")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒself, "_extract_base_url", ϒcourse_id, ϒvideo_id)
						}
					}()
					ϒsettings = λ.Calm(ϒself, "_download_xml", λ.Mod(λ.StrLiteral("%s/content/content_%s/videosettings.xml?v=1"), λ.NewTuple(
						ϒbase_url,
						ϒvideo_id,
					)), ϒvideo_id, λ.StrLiteral("Downloading video settings XML"))
					τmp0 = λ.Calm(ϒself, "_extract_chapter_and_title", λ.Call(ϒxpath_text, λ.NewArgs(
						ϒsettings,
						λ.StrLiteral(".//Title"),
						λ.StrLiteral("title"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.True},
					}))
					_ = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒtitle = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒsettings, "findall", λ.Cal(ϒcompat_xpath, λ.StrLiteral(".//MediaSources"))))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒsources = τmp1
						ϒsources_type = λ.Calm(ϒsources, "get", λ.StrLiteral("videoType"))
						τmp2 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒsources, "findall", λ.Cal(ϒcompat_xpath, λ.StrLiteral("./MediaSource"))))
						for {
							if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
								break
							}
							ϒsource = τmp3
							ϒvideo_url = λ.GetAttr(ϒsource, "text", nil)
							if λ.IsTrue(func() λ.Object {
								if λv := λ.NewBool(!λ.IsTrue(ϒvideo_url)); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.NewBool(!λ.IsTrue(λ.Calm(ϒvideo_url, "startswith", λ.StrLiteral("http"))))
								}
							}()) {
								continue
							}
							if λ.IsTrue(λ.Eq(ϒsources_type, λ.StrLiteral("smoothstreaming"))) {
								λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_ism_formats", nil), λ.NewArgs(
									ϒvideo_url,
									ϒvideo_id,
									λ.StrLiteral("mss"),
								), λ.KWArgs{
									{Name: "fatal", Value: λ.False},
								}))
								continue
							}
							ϒvideo_mode = λ.Calm(ϒsource, "get", λ.StrLiteral("videoMode"))
							ϒheight = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("^(\\d+)[pP]$"),
								func() λ.Object {
									if λv := ϒvideo_mode; λ.IsTrue(λv) {
										return λv
									} else {
										return λ.StrLiteral("")
									}
								}(),
								λ.StrLiteral("height"),
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							}))
							ϒcodec = λ.Calm(ϒsource, "get", λ.StrLiteral("codec"))
							τmp4 = λ.Mul(λ.NewList(λ.None), λ.IntLiteral(2))
							ϒacodec = λ.GetItem(τmp4, λ.IntLiteral(0))
							ϒvcodec = λ.GetItem(τmp4, λ.IntLiteral(1))
							if λ.IsTrue(ϒcodec) {
								ϒcodecs = λ.Calm(ϒcodec, "split", λ.StrLiteral(","))
								if λ.IsTrue(λ.Eq(λ.Cal(λ.BuiltinLen, ϒcodecs), λ.IntLiteral(2))) {
									τmp4 = ϒcodecs
									ϒacodec = λ.GetItem(τmp4, λ.IntLiteral(0))
									ϒvcodec = λ.GetItem(τmp4, λ.IntLiteral(1))
								} else {
									if λ.IsTrue(λ.Eq(λ.Cal(λ.BuiltinLen, ϒcodecs), λ.IntLiteral(1))) {
										ϒvcodec = λ.GetItem(ϒcodecs, λ.IntLiteral(0))
									}
								}
							}
							λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
								"url":       ϒvideo_url,
								"format_id": ϒvideo_mode,
								"height":    ϒheight,
								"acodec":    ϒacodec,
								"vcodec":    ϒvcodec,
							}))
						}
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒsubtitles = λ.DictLiteral(map[λ.Object]λ.Object{})
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒsettings, "findall", λ.Cal(ϒcompat_xpath, λ.StrLiteral(".//MarkerResourceSource"))))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒsource = τmp1
						ϒsubtitle_url = λ.GetAttr(ϒsource, "text", nil)
						if !λ.IsTrue(ϒsubtitle_url) {
							continue
						}
						λ.Calm(λ.Calm(ϒsubtitles, "setdefault", λ.StrLiteral("en"), λ.NewList()), "append", λ.DictLiteral(map[string]λ.Object{
							"url": λ.Mod(λ.StrLiteral("%s/%s"), λ.NewTuple(
								ϒbase_url,
								ϒsubtitle_url,
							)),
							"ext": λ.Calm(ϒsource, "get", λ.StrLiteral("type")),
						}))
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id":        ϒvideo_id,
						"title":     ϒtitle,
						"subtitles": ϒsubtitles,
						"formats":   ϒformats,
					})
				})
			return λ.DictLiteral(map[string]λ.Object{
				"IE_NAME":       MicrosoftVirtualAcademyIE_IE_NAME,
				"_VALID_URL":    MicrosoftVirtualAcademyIE__VALID_URL,
				"_real_extract": MicrosoftVirtualAcademyIE__real_extract,
			})
		}())
		MicrosoftVirtualAcademyCourseIE = λ.Cal(λ.TypeType, λ.StrLiteral("MicrosoftVirtualAcademyCourseIE"), λ.NewTuple(MicrosoftVirtualAcademyBaseIE), func() λ.Dict {
			var (
				MicrosoftVirtualAcademyCourseIE_IE_NAME    λ.Object
				MicrosoftVirtualAcademyCourseIE__VALID_URL λ.Object
				MicrosoftVirtualAcademyCourseIE_suitable   λ.Object
			)
			MicrosoftVirtualAcademyCourseIE_IE_NAME = λ.StrLiteral("mva:course")
			MicrosoftVirtualAcademyCourseIE__VALID_URL = λ.Mod(λ.StrLiteral("(?:%s:|https?://(?:mva\\.microsoft|(?:www\\.)?microsoftvirtualacademy)\\.com/[^/]+/training-courses/(?P<display_id>[^/?#&]+)-)(?P<id>\\d+)"), MicrosoftVirtualAcademyCourseIE_IE_NAME)
			MicrosoftVirtualAcademyCourseIE_suitable = λ.NewFunction("suitable",
				[]λ.Param{
					{Name: "cls"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcls = λargs[0]
						ϒurl = λargs[1]
					)
					return func() λ.Object {
						if λ.IsTrue(λ.Calm(MicrosoftVirtualAcademyIE, "suitable", ϒurl)) {
							return λ.False
						} else {
							return λ.Calm(λ.Cal(λ.SuperType, MicrosoftVirtualAcademyCourseIE, ϒcls), "suitable", ϒurl)
						}
					}()
				})
			MicrosoftVirtualAcademyCourseIE_suitable = λ.Cal(λ.ClassMethodType, MicrosoftVirtualAcademyCourseIE_suitable)
			return λ.DictLiteral(map[string]λ.Object{
				"IE_NAME":    MicrosoftVirtualAcademyCourseIE_IE_NAME,
				"_VALID_URL": MicrosoftVirtualAcademyCourseIE__VALID_URL,
				"suitable":   MicrosoftVirtualAcademyCourseIE_suitable,
			})
		}())
	})
}
