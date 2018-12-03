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
 * presstv/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/presstv.py
 */

package presstv

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	PressTVIE     λ.Object
	ϒremove_start λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒremove_start = Ωutils.ϒremove_start
		PressTVIE = λ.Cal(λ.TypeType, λ.NewStr("PressTVIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				PressTVIE__TEST         λ.Object
				PressTVIE__VALID_URL    λ.Object
				PressTVIE__real_extract λ.Object
			)
			PressTVIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?presstv\\.ir/[^/]+/(?P<y>\\d+)/(?P<m>\\d+)/(?P<d>\\d+)/(?P<id>\\d+)/(?P<display_id>[^/]+)?")
			PressTVIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("http://www.presstv.ir/Detail/2016/04/09/459911/Australian-sewerage-treatment-facility-/"),
				λ.NewStr("md5"): λ.NewStr("5d7e3195a447cb13e9267e931d8dd5a5"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):          λ.NewStr("459911"),
					λ.NewStr("display_id"):  λ.NewStr("Australian-sewerage-treatment-facility-"),
					λ.NewStr("ext"):         λ.NewStr("mp4"),
					λ.NewStr("title"):       λ.NewStr("Organic mattresses used to clean waste water"),
					λ.NewStr("upload_date"): λ.NewStr("20160409"),
					λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.jpg"),
					λ.NewStr("description"): λ.NewStr("md5:20002e654bbafb6908395a5c0cfcd125"),
				}),
			})
			PressTVIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒ_formats    λ.Object
						ϒbase_url    λ.Object
						ϒdescription λ.Object
						ϒdisplay_id  λ.Object
						ϒformats     λ.Object
						ϒmobj        λ.Object
						ϒself        = λargs[0]
						ϒthumbnail   λ.Object
						ϒtitle       λ.Object
						ϒupload_date λ.Object
						ϒurl         = λargs[1]
						ϒvideo_id    λ.Object
						ϒvideo_url   λ.Object
						ϒwebpage     λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					ϒdisplay_id = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("display_id")); λ.IsTrue(λv) {
							return λv
						} else {
							return ϒvideo_id
						}
					}()
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒdisplay_id)
					ϒvideo_url = λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_hidden_inputs", nil), ϒwebpage), λ.NewStr("inpPlayback"))
					ϒbase_url = λ.NewStr("http://192.99.219.222:82/presstv")
					ϒ_formats = λ.NewList(
						λ.NewTuple(
							λ.NewInt(180),
							λ.NewStr("_low200.mp4"),
						),
						λ.NewTuple(
							λ.NewInt(360),
							λ.NewStr("_low400.mp4"),
						),
						λ.NewTuple(
							λ.NewInt(720),
							λ.NewStr("_low800.mp4"),
						),
						λ.NewTuple(
							λ.NewInt(1080),
							λ.NewStr(".mp4"),
						),
					)
					ϒformats = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒextension λ.Object
									ϒheight    λ.Object
									τmp0       λ.Object
									τmp1       λ.Object
									τmp2       λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, ϒ_formats)
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									τmp2 = τmp1
									ϒheight = λ.GetItem(τmp2, λ.NewInt(0))
									ϒextension = λ.GetItem(τmp2, λ.NewInt(1))
									λgen.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("url"):       λ.Add(λ.Add(ϒbase_url, λ.GetItem(ϒvideo_url, λ.NewSlice(λ.None, λ.Neg(λ.NewInt(4)), λ.None))), ϒextension),
										λ.NewStr("format_id"): λ.Mod(λ.NewStr("%dp"), ϒheight),
										λ.NewStr("height"):    ϒheight,
									}))
								}
								return λ.None
							})
						})))
					ϒtitle = λ.Cal(ϒremove_start, λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
						λ.NewStr("title"),
						ϒwebpage,
					), λ.KWArgs{
						{Name: "fatal", Value: λ.True},
					}), λ.NewStr("PressTV-"))
					ϒthumbnail = λ.Cal(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), ϒwebpage)
					ϒdescription = λ.Cal(λ.GetAttr(ϒself, "_og_search_description", nil), ϒwebpage)
					ϒupload_date = λ.Mod(λ.NewStr("%04d%02d%02d"), λ.NewTuple(
						λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("y"))),
						λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("m"))),
						λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("d"))),
					))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("display_id"):  ϒdisplay_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("formats"):     ϒformats,
						λ.NewStr("thumbnail"):   ϒthumbnail,
						λ.NewStr("upload_date"): ϒupload_date,
						λ.NewStr("description"): ϒdescription,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TEST"):         PressTVIE__TEST,
				λ.NewStr("_VALID_URL"):    PressTVIE__VALID_URL,
				λ.NewStr("_real_extract"): PressTVIE__real_extract,
			})
		}())
	})
}
