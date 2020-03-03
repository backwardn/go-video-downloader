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
 * br/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/br.py
 */

package br

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	BRIE            λ.Object
	BRMediathekIE   λ.Object
	ExtractorError  λ.Object
	InfoExtractor   λ.Object
	ϒdetermine_ext  λ.Object
	ϒint_or_none    λ.Object
	ϒparse_duration λ.Object
	ϒparse_iso8601  λ.Object
	ϒxpath_element  λ.Object
	ϒxpath_text     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒxpath_element = Ωutils.ϒxpath_element
		ϒxpath_text = Ωutils.ϒxpath_text
		BRIE = λ.Cal(λ.TypeType, λ.StrLiteral("BRIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				BRIE__VALID_URL          λ.Object
				BRIE__extract_formats    λ.Object
				BRIE__extract_thumbnails λ.Object
				BRIE__real_extract       λ.Object
			)
			BRIE__VALID_URL = λ.StrLiteral("(?P<base_url>https?://(?:www\\.)?br(?:-klassik)?\\.de)/(?:[a-z0-9\\-_]+/)+(?P<id>[a-z0-9\\-_]+)\\.html")
			BRIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒbase_url       λ.Object
						ϒbroadcast_date λ.Object
						ϒdisplay_id     λ.Object
						ϒmedia          λ.Object
						ϒmedia_id       λ.Object
						ϒmedias         λ.Object
						ϒpage           λ.Object
						ϒself           = λargs[0]
						ϒurl            = λargs[1]
						ϒxml            λ.Object
						ϒxml_media      λ.Object
						ϒxml_url        λ.Object
						τmp0            λ.Object
						τmp1            λ.Object
					)
					τmp0 = λ.Calm(λ.Cal(Ωre.ϒsearch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups")
					ϒbase_url = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒdisplay_id = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒdisplay_id)
					ϒxml_url = λ.Calm(ϒself, "_search_regex", λ.StrLiteral("return BRavFramework\\.register\\(BRavFramework\\('avPlayer_(?:[a-f0-9-]{36})'\\)\\.setup\\({dataURL:'(/(?:[a-z0-9\\-]+/)+[a-z0-9/~_.-]+)'}\\)\\);"), ϒpage, λ.StrLiteral("XMLURL"))
					ϒxml = λ.Calm(ϒself, "_download_xml", λ.Add(ϒbase_url, ϒxml_url), ϒdisplay_id)
					ϒmedias = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Add(λ.Calm(ϒxml, "findall", λ.StrLiteral("video")), λ.Calm(ϒxml, "findall", λ.StrLiteral("audio"))))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒxml_media = τmp1
						ϒmedia_id = λ.Calm(ϒxml_media, "get", λ.StrLiteral("externalId"))
						ϒmedia = λ.DictLiteral(map[string]λ.Object{
							"id":          ϒmedia_id,
							"title":       λ.Cal(ϒxpath_text, ϒxml_media, λ.StrLiteral("title"), λ.StrLiteral("title"), λ.True),
							"duration":    λ.Cal(ϒparse_duration, λ.Cal(ϒxpath_text, ϒxml_media, λ.StrLiteral("duration"))),
							"formats":     λ.Calm(ϒself, "_extract_formats", λ.Cal(ϒxpath_element, ϒxml_media, λ.StrLiteral("assets")), ϒmedia_id),
							"thumbnails":  λ.Calm(ϒself, "_extract_thumbnails", λ.Cal(ϒxpath_element, ϒxml_media, λ.StrLiteral("teaserImage/variants")), ϒbase_url),
							"description": λ.Cal(ϒxpath_text, ϒxml_media, λ.StrLiteral("desc")),
							"webpage_url": λ.Cal(ϒxpath_text, ϒxml_media, λ.StrLiteral("permalink")),
							"uploader":    λ.Cal(ϒxpath_text, ϒxml_media, λ.StrLiteral("author")),
						})
						ϒbroadcast_date = λ.Cal(ϒxpath_text, ϒxml_media, λ.StrLiteral("broadcastDate"))
						if λ.IsTrue(ϒbroadcast_date) {
							λ.SetItem(ϒmedia, λ.StrLiteral("upload_date"), λ.Calm(λ.StrLiteral(""), "join", λ.Cal(λ.ReversedIteratorType, λ.Calm(ϒbroadcast_date, "split", λ.StrLiteral(".")))))
						}
						λ.Calm(ϒmedias, "append", ϒmedia)
					}
					if λ.IsTrue(λ.Gt(λ.Cal(λ.BuiltinLen, ϒmedias), λ.IntLiteral(1))) {
						λ.Calm(λ.GetAttr(ϒself, "_downloader", nil), "report_warning", λ.StrLiteral("found multiple medias; please report this with the video URL to http://yt-dl.org/bug"))
					}
					if !λ.IsTrue(ϒmedias) {
						panic(λ.Raise(λ.Cal(ExtractorError, λ.StrLiteral("No media entries found"))))
					}
					return λ.GetItem(ϒmedias, λ.IntLiteral(0))
				})
			BRIE__extract_formats = λ.NewFunction("_extract_formats",
				[]λ.Param{
					{Name: "self"},
					{Name: "assets"},
					{Name: "media_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒasset            λ.Object
						ϒasset_type       λ.Object
						ϒassets           = λargs[1]
						ϒformat_info      λ.Object
						ϒformat_url       λ.Object
						ϒformats          λ.Object
						ϒhttp_format_info λ.Object
						ϒmedia_id         = λargs[2]
						ϒrtmp_format_info λ.Object
						ϒself             = λargs[0]
						ϒserver_prefix    λ.Object
						τmp0              λ.Object
						τmp1              λ.Object
					)
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒassets, "findall", λ.StrLiteral("asset")))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒasset = τmp1
						ϒformat_url = λ.Cal(ϒxpath_text, ϒasset, λ.NewList(
							λ.StrLiteral("downloadUrl"),
							λ.StrLiteral("url"),
						))
						ϒasset_type = λ.Calm(ϒasset, "get", λ.StrLiteral("type"))
						if λ.IsTrue(λ.Calm(ϒasset_type, "startswith", λ.StrLiteral("HDS"))) {
							λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
								λ.Add(ϒformat_url, λ.StrLiteral("?hdcore=3.2.0")),
								ϒmedia_id,
							), λ.KWArgs{
								{Name: "f4m_id", Value: λ.StrLiteral("hds")},
								{Name: "fatal", Value: λ.False},
							}))
						} else {
							if λ.IsTrue(λ.Calm(ϒasset_type, "startswith", λ.StrLiteral("HLS"))) {
								λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
									ϒformat_url,
									ϒmedia_id,
									λ.StrLiteral("mp4"),
									λ.StrLiteral("m3u8_native"),
								), λ.KWArgs{
									{Name: "m3u8_id", Value: λ.StrLiteral("hds")},
									{Name: "fatal", Value: λ.False},
								}))
							} else {
								ϒformat_info = λ.DictLiteral(map[string]λ.Object{
									"ext":       λ.Cal(ϒxpath_text, ϒasset, λ.StrLiteral("mediaType")),
									"width":     λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒasset, λ.StrLiteral("frameWidth"))),
									"height":    λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒasset, λ.StrLiteral("frameHeight"))),
									"tbr":       λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒasset, λ.StrLiteral("bitrateVideo"))),
									"abr":       λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒasset, λ.StrLiteral("bitrateAudio"))),
									"vcodec":    λ.Cal(ϒxpath_text, ϒasset, λ.StrLiteral("codecVideo")),
									"acodec":    λ.Cal(ϒxpath_text, ϒasset, λ.StrLiteral("codecAudio")),
									"container": λ.Cal(ϒxpath_text, ϒasset, λ.StrLiteral("mediaType")),
									"filesize":  λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒasset, λ.StrLiteral("size"))),
								})
								ϒformat_url = λ.Calm(ϒself, "_proto_relative_url", ϒformat_url)
								if λ.IsTrue(ϒformat_url) {
									ϒhttp_format_info = λ.Calm(ϒformat_info, "copy")
									λ.Calm(ϒhttp_format_info, "update", λ.DictLiteral(map[string]λ.Object{
										"url":       ϒformat_url,
										"format_id": λ.Mod(λ.StrLiteral("http-%s"), ϒasset_type),
									}))
									λ.Calm(ϒformats, "append", ϒhttp_format_info)
								}
								ϒserver_prefix = λ.Cal(ϒxpath_text, ϒasset, λ.StrLiteral("serverPrefix"))
								if λ.IsTrue(ϒserver_prefix) {
									ϒrtmp_format_info = λ.Calm(ϒformat_info, "copy")
									λ.Calm(ϒrtmp_format_info, "update", λ.DictLiteral(map[string]λ.Object{
										"url":       ϒserver_prefix,
										"play_path": λ.Cal(ϒxpath_text, ϒasset, λ.StrLiteral("fileName")),
										"format_id": λ.Mod(λ.StrLiteral("rtmp-%s"), ϒasset_type),
									}))
									λ.Calm(ϒformats, "append", ϒrtmp_format_info)
								}
							}
						}
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					return ϒformats
				})
			BRIE__extract_thumbnails = λ.NewFunction("_extract_thumbnails",
				[]λ.Param{
					{Name: "self"},
					{Name: "variants"},
					{Name: "base_url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒbase_url   = λargs[2]
						ϒself       = λargs[0]
						ϒthumbnails λ.Object
						ϒvariants   = λargs[1]
					)
					_ = ϒself
					ϒthumbnails = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒvariant λ.Object
									τmp0     λ.Object
									τmp1     λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒvariants, "findall", λ.StrLiteral("variant")))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒvariant = τmp1
									if λ.IsTrue(λ.Cal(ϒxpath_text, ϒvariant, λ.StrLiteral("url"))) {
										λgy.Yield(λ.DictLiteral(map[string]λ.Object{
											"url":    λ.Add(ϒbase_url, λ.Cal(ϒxpath_text, ϒvariant, λ.StrLiteral("url"))),
											"width":  λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒvariant, λ.StrLiteral("width"))),
											"height": λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒvariant, λ.StrLiteral("height"))),
										}))
									}
								}
								return λ.None
							})
						})))
					λ.Call(λ.GetAttr(ϒthumbnails, "sort", nil), nil, λ.KWArgs{
						{Name: "key", Value: λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.Mul(λ.GetItem(ϒx, λ.StrLiteral("width")), λ.GetItem(ϒx, λ.StrLiteral("height")))
							})},
						{Name: "reverse", Value: λ.True},
					})
					return ϒthumbnails
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":          BRIE__VALID_URL,
				"_extract_formats":    BRIE__extract_formats,
				"_extract_thumbnails": BRIE__extract_thumbnails,
				"_real_extract":       BRIE__real_extract,
			})
		}())
		BRMediathekIE = λ.Cal(λ.TypeType, λ.StrLiteral("BRMediathekIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				BRMediathekIE__VALID_URL λ.Object
			)
			BRMediathekIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?br\\.de/mediathek/video/[^/?&#]*?-(?P<id>av:[0-9a-f]{24})")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": BRMediathekIE__VALID_URL,
			})
		}())
	})
}
