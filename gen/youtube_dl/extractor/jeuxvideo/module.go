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
 * jeuxvideo/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/jeuxvideo.py
 */

package jeuxvideo

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	JeuxVideoIE   λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		JeuxVideoIE = λ.Cal(λ.TypeType, λ.NewStr("JeuxVideoIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				JeuxVideoIE__TESTS        λ.Object
				JeuxVideoIE__VALID_URL    λ.Object
				JeuxVideoIE__real_extract λ.Object
			)
			JeuxVideoIE__VALID_URL = λ.NewStr("https?://.*?\\.jeuxvideo\\.com/.*/(.*?)\\.htm")
			JeuxVideoIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.jeuxvideo.com/reportages-videos-jeux/0004/00046170/tearaway-playstation-vita-gc-2013-tearaway-nous-presente-ses-papiers-d-identite-00115182.htm"),
					λ.NewStr("md5"): λ.NewStr("046e491afb32a8aaac1f44dd4ddd54ee"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("114765"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Tearaway : GC 2013 : Tearaway nous présente ses papiers d'identité"),
						λ.NewStr("description"): λ.NewStr("Lorsque les développeurs de LittleBigPlanet proposent un nouveau titre, on ne peut que s'attendre à un résultat original et fort attrayant."),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.jeuxvideo.com/videos/chroniques/434220/l-histoire-du-jeu-video-la-saturn.htm"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			JeuxVideoIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒconfig     λ.Object
						ϒconfig_url λ.Object
						ϒformats    λ.Object
						ϒmobj       λ.Object
						ϒself       = λargs[0]
						ϒtitle      λ.Object
						ϒurl        = λargs[1]
						ϒvideo_id   λ.Object
						ϒwebpage    λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒtitle = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewInt(1))
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒtitle)
					ϒtitle = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewStr("name"), ϒwebpage); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒself, "_og_search_title", nil), ϒwebpage)
						}
					}()
					ϒconfig_url = λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewStr("data-src(?:set-video)?=\"(/contenu/medias/video\\.php.*?)\""), ϒwebpage, λ.NewStr("config URL"))
					ϒconfig_url = λ.Add(λ.NewStr("http://www.jeuxvideo.com"), ϒconfig_url)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("id=(\\d+)"), ϒconfig_url, λ.NewStr("video ID"))
					ϒconfig = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), ϒconfig_url, ϒtitle, λ.NewStr("Downloading JSON config"))
					ϒformats = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒsource λ.Object
									τmp0    λ.Object
									τmp1    λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.ReversedIteratorType, λ.GetItem(ϒconfig, λ.NewStr("sources"))))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒsource = τmp1
									λgy.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("url"):        λ.GetItem(ϒsource, λ.NewStr("file")),
										λ.NewStr("format_id"):  λ.GetItem(ϒsource, λ.NewStr("label")),
										λ.NewStr("resolution"): λ.GetItem(ϒsource, λ.NewStr("label")),
									}))
								}
								return λ.None
							})
						})))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("formats"):     ϒformats,
						λ.NewStr("description"): λ.Cal(λ.GetAttr(ϒself, "_og_search_description", nil), ϒwebpage),
						λ.NewStr("thumbnail"):   λ.Cal(λ.GetAttr(ϒconfig, "get", nil), λ.NewStr("image")),
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        JeuxVideoIE__TESTS,
				λ.NewStr("_VALID_URL"):    JeuxVideoIE__VALID_URL,
				λ.NewStr("_real_extract"): JeuxVideoIE__real_extract,
			})
		}())
	})
}
