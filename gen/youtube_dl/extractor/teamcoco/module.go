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
 * teamcoco/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/teamcoco.py
 */

package teamcoco

import (
	Ωjson "github.com/tenta-browser/go-video-downloader/gen/json"
	Ωturner "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/turner"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError  λ.Object
	TeamcocoIE      λ.Object
	TurnerBaseIE    λ.Object
	ϒdetermine_ext  λ.Object
	ϒint_or_none    λ.Object
	ϒmimetype2ext   λ.Object
	ϒparse_duration λ.Object
	ϒparse_iso8601  λ.Object
	ϒqualities      λ.Object
)

func init() {
	λ.InitModule(func() {
		TurnerBaseIE = Ωturner.TurnerBaseIE
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒmimetype2ext = Ωutils.ϒmimetype2ext
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒqualities = Ωutils.ϒqualities
		TeamcocoIE = λ.Cal(λ.TypeType, λ.StrLiteral("TeamcocoIE"), λ.NewTuple(TurnerBaseIE), func() λ.Dict {
			var (
				TeamcocoIE__RECORD_TEMPL λ.Object
				TeamcocoIE__VALID_URL    λ.Object
				TeamcocoIE__graphql_call λ.Object
				TeamcocoIE__real_extract λ.Object
			)
			TeamcocoIE__VALID_URL = λ.StrLiteral("https?://(?:\\w+\\.)?teamcoco\\.com/(?P<id>([^/]+/)*[^/?#]+)")
			TeamcocoIE__RECORD_TEMPL = λ.StrLiteral("id\n        title\n        teaser\n        publishOn\n        thumb {\n          preview\n        }\n        tags {\n          name\n        }\n        duration\n        turnerMediaId\n        turnerMediaAuthToken")
			TeamcocoIE__graphql_call = λ.NewFunction("_graphql_call",
				[]λ.Param{
					{Name: "self"},
					{Name: "query_template"},
					{Name: "object_type"},
					{Name: "object_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒfind_object    λ.Object
						ϒobject_id      = λargs[3]
						ϒobject_type    = λargs[2]
						ϒquery_template = λargs[1]
						ϒself           = λargs[0]
					)
					ϒfind_object = λ.Add(λ.StrLiteral("find"), ϒobject_type)
					return λ.GetItem(λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.StrLiteral("https://teamcoco.com/graphql"),
						ϒobject_id,
					), λ.KWArgs{
						{Name: "data", Value: λ.Calm(λ.Cal(Ωjson.ϒdumps, λ.DictLiteral(map[string]λ.Object{
							"query": λ.Mod(ϒquery_template, λ.NewTuple(
								ϒfind_object,
								ϒobject_id,
							)),
						})), "encode")},
						{Name: "headers", Value: λ.DictLiteral(map[string]string{
							"Content-Type": "application/json",
						})},
					}), λ.StrLiteral("data")), ϒfind_object)
				})
			TeamcocoIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒchild         λ.Object
						ϒdisplay_id    λ.Object
						ϒext           λ.Object
						ϒformat_id     λ.Object
						ϒformats       λ.Object
						ϒget_quality   λ.Object
						ϒinfo          λ.Object
						ϒmedia_id      λ.Object
						ϒrecord        λ.Object
						ϒresponse      λ.Object
						ϒself          = λargs[0]
						ϒsrc           λ.Object
						ϒsrc_url       λ.Object
						ϒtbr           λ.Object
						ϒurl           = λargs[1]
						ϒvideo_id      λ.Object
						ϒvideo_sources λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
					)
					ϒdisplay_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒresponse = λ.Calm(ϒself, "_graphql_call", λ.Mod(λ.StrLiteral("{\n  %%s(slug: \"%%s\") {\n    ... on RecordSlug {\n      record {\n        %s\n      }\n    }\n    ... on PageSlug {\n      child {\n        id\n      }\n    }\n    ... on NotFoundSlug {\n      status\n    }\n  }\n}"), λ.GetAttr(ϒself, "_RECORD_TEMPL", nil)), λ.StrLiteral("Slug"), ϒdisplay_id)
					if λ.IsTrue(λ.Calm(ϒresponse, "get", λ.StrLiteral("status"))) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.StrLiteral("This video is no longer available.")), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒchild = λ.Calm(ϒresponse, "get", λ.StrLiteral("child"))
					if λ.IsTrue(ϒchild) {
						ϒrecord = λ.Calm(ϒself, "_graphql_call", λ.Mod(λ.StrLiteral("{\n  %%s(id: \"%%s\") {\n    ... on Video {\n      %s\n    }\n  }\n}"), λ.GetAttr(ϒself, "_RECORD_TEMPL", nil)), λ.StrLiteral("Record"), λ.GetItem(ϒchild, λ.StrLiteral("id")))
					} else {
						ϒrecord = λ.GetItem(ϒresponse, λ.StrLiteral("record"))
					}
					ϒvideo_id = λ.GetItem(ϒrecord, λ.StrLiteral("id"))
					ϒinfo = λ.DictLiteral(map[string]λ.Object{
						"id":          ϒvideo_id,
						"display_id":  ϒdisplay_id,
						"title":       λ.GetItem(ϒrecord, λ.StrLiteral("title")),
						"thumbnail":   λ.Calm(λ.Calm(ϒrecord, "get", λ.StrLiteral("thumb"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("preview")),
						"description": λ.Calm(ϒrecord, "get", λ.StrLiteral("teaser")),
						"duration":    λ.Cal(ϒparse_duration, λ.Calm(ϒrecord, "get", λ.StrLiteral("duration"))),
						"timestamp":   λ.Cal(ϒparse_iso8601, λ.Calm(ϒrecord, "get", λ.StrLiteral("publishOn"))),
					})
					ϒmedia_id = λ.Calm(ϒrecord, "get", λ.StrLiteral("turnerMediaId"))
					if λ.IsTrue(ϒmedia_id) {
						λ.Calm(ϒself, "_initialize_geo_bypass", λ.DictLiteral(map[string]λ.Object{
							"countries": λ.NewList(λ.StrLiteral("US")),
						}))
						λ.Calm(ϒinfo, "update", λ.Calm(ϒself, "_extract_ngtv_info", ϒmedia_id, λ.DictLiteral(map[string]λ.Object{
							"accessToken":     λ.GetItem(ϒrecord, λ.StrLiteral("turnerMediaAuthToken")),
							"accessTokenType": λ.StrLiteral("jws"),
						})))
					} else {
						ϒvideo_sources = λ.GetItem(λ.GetItem(λ.Calm(ϒself, "_download_json", λ.Add(λ.StrLiteral("https://teamcoco.com/_truman/d/"), ϒvideo_id), ϒvideo_id), λ.StrLiteral("meta")), λ.StrLiteral("src"))
						if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒvideo_sources, λ.DictType)) {
							ϒvideo_sources = λ.Calm(ϒvideo_sources, "values")
						}
						ϒformats = λ.NewList()
						ϒget_quality = λ.Cal(ϒqualities, λ.NewList(
							λ.StrLiteral("low"),
							λ.StrLiteral("sd"),
							λ.StrLiteral("hd"),
							λ.StrLiteral("uhd"),
						))
						τmp0 = λ.Cal(λ.BuiltinIter, ϒvideo_sources)
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒsrc = τmp1
							if !λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒsrc, λ.DictType)) {
								continue
							}
							ϒsrc_url = λ.Calm(ϒsrc, "get", λ.StrLiteral("src"))
							if !λ.IsTrue(ϒsrc_url) {
								continue
							}
							ϒformat_id = λ.Calm(ϒsrc, "get", λ.StrLiteral("label"))
							ϒext = λ.Cal(ϒdetermine_ext, ϒsrc_url, λ.Cal(ϒmimetype2ext, λ.Calm(ϒsrc, "get", λ.StrLiteral("type"))))
							if λ.IsTrue(func() λ.Object {
								if λv := λ.Eq(ϒformat_id, λ.StrLiteral("hls")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Eq(ϒext, λ.StrLiteral("m3u8"))
								}
							}()) {
								if λ.IsTrue(λ.Calm(ϒsrc_url, "startswith", λ.StrLiteral("/"))) {
									ϒsrc_url = λ.Add(λ.StrLiteral("http://ht.cdn.turner.com/tbs/big/teamcoco"), ϒsrc_url)
								}
								λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
									ϒsrc_url,
									ϒvideo_id,
									λ.StrLiteral("mp4"),
								), λ.KWArgs{
									{Name: "m3u8_id", Value: ϒformat_id},
									{Name: "fatal", Value: λ.False},
								}))
							} else {
								if λ.IsTrue(λ.Calm(ϒsrc_url, "startswith", λ.StrLiteral("/mp4:protected/"))) {
									continue
								}
								ϒtbr = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
									λ.StrLiteral("(\\d+)k\\.mp4"),
									ϒsrc_url,
									λ.StrLiteral("tbr"),
								), λ.KWArgs{
									{Name: "default", Value: λ.None},
								}))
								λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
									"url":       ϒsrc_url,
									"ext":       ϒext,
									"tbr":       ϒtbr,
									"format_id": ϒformat_id,
									"quality":   λ.Cal(ϒget_quality, ϒformat_id),
								}))
							}
						}
						λ.Calm(ϒself, "_sort_formats", ϒformats)
						λ.SetItem(ϒinfo, λ.StrLiteral("formats"), ϒformats)
					}
					return ϒinfo
				})
			return λ.DictLiteral(map[string]λ.Object{
				"_RECORD_TEMPL": TeamcocoIE__RECORD_TEMPL,
				"_VALID_URL":    TeamcocoIE__VALID_URL,
				"_graphql_call": TeamcocoIE__graphql_call,
				"_real_extract": TeamcocoIE__real_extract,
			})
		}())
	})
}
