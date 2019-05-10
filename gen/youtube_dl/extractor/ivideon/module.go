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
 * ivideon/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/ivideon.py
 */

package ivideon

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωparse "github.com/tenta-browser/go-video-downloader/gen/urllib/parse"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor                  λ.Object
	IvideonIE                      λ.Object
	ϒcompat_urllib_parse_urlencode λ.Object
	ϒqualities                     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_urllib_parse_urlencode = Ωcompat.ϒcompat_urllib_parse_urlencode
		ϒqualities = Ωutils.ϒqualities
		IvideonIE = λ.Cal(λ.TypeType, λ.NewStr("IvideonIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				IvideonIE_IE_NAME       λ.Object
				IvideonIE__QUALITIES    λ.Object
				IvideonIE__TESTS        λ.Object
				IvideonIE__VALID_URL    λ.Object
				IvideonIE__real_extract λ.Object
			)
			IvideonIE_IE_NAME = λ.NewStr("ivideon")
			IvideonIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?ivideon\\.com/tv/(?:[^/]+/)*camera/(?P<id>\\d+-[\\da-f]+)/(?P<camera_id>\\d+)")
			IvideonIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.ivideon.com/tv/camera/100-916ca13b5c4ad9f564266424a026386d/0/"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("100-916ca13b5c4ad9f564266424a026386d"),
						λ.NewStr("ext"):         λ.NewStr("flv"),
						λ.NewStr("title"):       λ.NewStr("re:^Касса [0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}$"),
						λ.NewStr("description"): λ.NewStr("Основное предназначение - запись действий кассиров. Плюс общий вид."),
						λ.NewStr("is_live"):     λ.True,
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.ivideon.com/tv/camera/100-c4ee4cb9ede885cf62dfbe93d7b53783/589824/?lang=ru"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.ivideon.com/tv/map/22.917923/-31.816406/16/camera/100-e7bc16c7d4b5bbd633fd5350b66dfa9a/0"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			IvideonIE__QUALITIES = λ.NewTuple(
				λ.NewStr("low"),
				λ.NewStr("mid"),
				λ.NewStr("hi"),
			)
			IvideonIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcamera_id     λ.Object
						ϒcamera_info   λ.Object
						ϒcamera_name   λ.Object
						ϒcamera_url    λ.Object
						ϒconfig        λ.Object
						ϒconfig_string λ.Object
						ϒdescription   λ.Object
						ϒformats       λ.Object
						ϒmobj          λ.Object
						ϒquality       λ.Object
						ϒself          = λargs[0]
						ϒserver_id     λ.Object
						ϒurl           = λargs[1]
						ϒwebpage       λ.Object
						τmp0           λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					τmp0 = λ.NewTuple(
						λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id")),
						λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("camera_id")),
					)
					ϒserver_id = λ.GetItem(τmp0, λ.NewInt(0))
					ϒcamera_id = λ.GetItem(τmp0, λ.NewInt(1))
					τmp0 = λ.NewTuple(
						λ.None,
						λ.None,
					)
					ϒcamera_name = λ.GetItem(τmp0, λ.NewInt(0))
					ϒdescription = λ.GetItem(τmp0, λ.NewInt(1))
					ϒcamera_url = λ.Cal(Ωparse.ϒurljoin, ϒurl, λ.Mod(λ.NewStr("/tv/camera/%s/%s/"), λ.NewTuple(
						ϒserver_id,
						ϒcamera_id,
					)))
					ϒwebpage = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
						ϒcamera_url,
						ϒserver_id,
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					if λ.IsTrue(ϒwebpage) {
						ϒconfig_string = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("var\\s+config\\s*=\\s*({.+?});"),
							ϒwebpage,
							λ.NewStr("config"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
						if λ.IsTrue(ϒconfig_string) {
							ϒconfig = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
								ϒconfig_string,
								ϒserver_id,
							), λ.KWArgs{
								{Name: "fatal", Value: λ.False},
							})
							ϒcamera_info = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒconfig, "get", nil), λ.NewStr("ivTvAppOptions"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("currentCameraInfo"))
							if λ.IsTrue(ϒcamera_info) {
								ϒcamera_name = λ.Cal(λ.GetAttr(ϒcamera_info, "get", nil), λ.NewStr("camera_name"))
								ϒdescription = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒcamera_info, "get", nil), λ.NewStr("misc"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("description"))
							}
						}
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒcamera_name))) {
							ϒcamera_name = func() λ.Object {
								if λv := λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
									λ.NewStr("name"),
									ϒwebpage,
									λ.NewStr("camera name"),
								), λ.KWArgs{
									{Name: "default", Value: λ.None},
								}); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
										λ.NewStr("<h1[^>]+class=\"b-video-title\"[^>]*>([^<]+)"),
										ϒwebpage,
										λ.NewStr("camera name"),
									), λ.KWArgs{
										{Name: "default", Value: λ.None},
									})
								}
							}()
						}
					}
					ϒquality = λ.Cal(ϒqualities, λ.GetAttr(ϒself, "_QUALITIES", nil))
					ϒformats = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒformat_id λ.Object
									τmp0       λ.Object
									τmp1       λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.GetAttr(ϒself, "_QUALITIES", nil))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒformat_id = τmp1
									λgy.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("url"): λ.Mod(λ.NewStr("https://streaming.ivideon.com/flv/live?%s"), λ.Cal(ϒcompat_urllib_parse_urlencode, λ.NewDictWithTable(map[λ.Object]λ.Object{
											λ.NewStr("server"):    ϒserver_id,
											λ.NewStr("camera"):    ϒcamera_id,
											λ.NewStr("sessionId"): λ.NewStr("demo"),
											λ.NewStr("q"):         λ.Cal(ϒquality, ϒformat_id),
										}))),
										λ.NewStr("format_id"): ϒformat_id,
										λ.NewStr("ext"):       λ.NewStr("flv"),
										λ.NewStr("quality"):   λ.Cal(ϒquality, ϒformat_id),
									}))
								}
								return λ.None
							})
						})))
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"): ϒserver_id,
						λ.NewStr("title"): λ.Cal(λ.GetAttr(ϒself, "_live_title", nil), func() λ.Object {
							if λv := ϒcamera_name; λ.IsTrue(λv) {
								return λv
							} else {
								return ϒserver_id
							}
						}()),
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("is_live"):     λ.True,
						λ.NewStr("formats"):     ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       IvideonIE_IE_NAME,
				λ.NewStr("_QUALITIES"):    IvideonIE__QUALITIES,
				λ.NewStr("_TESTS"):        IvideonIE__TESTS,
				λ.NewStr("_VALID_URL"):    IvideonIE__VALID_URL,
				λ.NewStr("_real_extract"): IvideonIE__real_extract,
			})
		}())
	})
}
