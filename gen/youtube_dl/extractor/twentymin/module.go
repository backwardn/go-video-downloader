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
 * twentymin/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/twentymin.py
 */

package twentymin

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor   λ.Object
	TwentyMinutenIE λ.Object
	ϒint_or_none    λ.Object
	ϒtry_get        λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒtry_get = Ωutils.ϒtry_get
		TwentyMinutenIE = λ.Cal(λ.TypeType, λ.NewStr("TwentyMinutenIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TwentyMinutenIE_IE_NAME       λ.Object
				TwentyMinutenIE__TESTS        λ.Object
				TwentyMinutenIE__VALID_URL    λ.Object
				TwentyMinutenIE__real_extract λ.Object
			)
			TwentyMinutenIE_IE_NAME = λ.NewStr("20min")
			TwentyMinutenIE__VALID_URL = λ.NewStr("(?x)\n                    https?://\n                        (?:www\\.)?20min\\.ch/\n                        (?:\n                            videotv/*\\?.*?\\bvid=|\n                            videoplayer/videoplayer\\.html\\?.*?\\bvideoId@\n                        )\n                        (?P<id>\\d+)\n                    ")
			TwentyMinutenIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.20min.ch/videotv/?vid=469148&cid=2"),
					λ.NewStr("md5"): λ.NewStr("e7264320db31eed8c38364150c12496e"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):        λ.NewStr("469148"),
						λ.NewStr("ext"):       λ.NewStr("mp4"),
						λ.NewStr("title"):     λ.NewStr("85 000 Franken für 15 perfekte Minuten"),
						λ.NewStr("thumbnail"): λ.NewStr("re:https?://.*\\.jpg$"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.20min.ch/videoplayer/videoplayer.html?params=client@twentyDE|videoId@523629"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("523629"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("So kommen Sie bei Eis und Schnee sicher an"),
						λ.NewStr("description"): λ.NewStr("md5:117c212f64b25e3d95747e5276863f7d"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:https?://.*\\.jpg$"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.20min.ch/videotv/?cid=44&vid=468738"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			TwentyMinutenIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdescription   λ.Object
						ϒdislike_count λ.Object
						ϒextract_count λ.Object
						ϒformats       λ.Object
						ϒlike_count    λ.Object
						ϒself          = λargs[0]
						ϒthumbnail     λ.Object
						ϒtitle         λ.Object
						ϒurl           = λargs[1]
						ϒvideo         λ.Object
						ϒvideo_id      λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒvideo = λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.NewStr("http://api.20min.ch/video/%s/show"), ϒvideo_id), ϒvideo_id), λ.NewStr("content"))
					ϒtitle = λ.GetItem(ϒvideo, λ.NewStr("title"))
					ϒformats = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒformat_id λ.Object
									ϒp         λ.Object
									ϒquality   λ.Object
									τmp0       λ.Object
									τmp1       λ.Object
									τmp2       λ.Object
									τmp3       λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.EnumerateIteratorType, λ.NewList(
									λ.NewTuple(
										λ.NewStr("sd"),
										λ.NewStr(""),
									),
									λ.NewTuple(
										λ.NewStr("hd"),
										λ.NewStr("h"),
									),
								)))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									τmp2 = τmp1
									ϒquality = λ.GetItem(τmp2, λ.NewInt(0))
									τmp3 = λ.GetItem(τmp2, λ.NewInt(1))
									ϒformat_id = λ.GetItem(τmp3, λ.NewInt(0))
									ϒp = λ.GetItem(τmp3, λ.NewInt(1))
									λgen.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("format_id"): ϒformat_id,
										λ.NewStr("url"): λ.Mod(λ.NewStr("http://podcast.20min-tv.ch/podcast/20min/%s%s.mp4"), λ.NewTuple(
											ϒvideo_id,
											ϒp,
										)),
										λ.NewStr("quality"): ϒquality,
									}))
								}
								return λ.None
							})
						})))
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒdescription = λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("lead"))
					ϒthumbnail = λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("thumbnail"))
					ϒextract_count = λ.NewFunction("extract_count",
						[]λ.Param{
							{Name: "kind"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒkind = λargs[0]
							)
							return λ.Cal(ϒtry_get, ϒvideo, λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.Cal(ϒint_or_none, λ.GetItem(λ.GetItem(ϒx, λ.NewStr("communityobject")), λ.Mod(λ.NewStr("thumbs_%s"), ϒkind)))
								}))
						})
					ϒlike_count = λ.Cal(ϒextract_count, λ.NewStr("up"))
					ϒdislike_count = λ.Cal(ϒextract_count, λ.NewStr("down"))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):            ϒvideo_id,
						λ.NewStr("title"):         ϒtitle,
						λ.NewStr("description"):   ϒdescription,
						λ.NewStr("thumbnail"):     ϒthumbnail,
						λ.NewStr("like_count"):    ϒlike_count,
						λ.NewStr("dislike_count"): ϒdislike_count,
						λ.NewStr("formats"):       ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       TwentyMinutenIE_IE_NAME,
				λ.NewStr("_TESTS"):        TwentyMinutenIE__TESTS,
				λ.NewStr("_VALID_URL"):    TwentyMinutenIE__VALID_URL,
				λ.NewStr("_real_extract"): TwentyMinutenIE__real_extract,
			})
		}())
	})
}
