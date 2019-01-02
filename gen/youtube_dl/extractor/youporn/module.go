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
 * youporn/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/youporn.py
 */

package youporn

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor      λ.Object
	YouPornIE          λ.Object
	ϒint_or_none       λ.Object
	ϒsanitized_Request λ.Object
	ϒstr_to_int        λ.Object
	ϒunescapeHTML      λ.Object
	ϒunified_strdate   λ.Object
	ϒurl_or_none       λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒsanitized_Request = Ωutils.ϒsanitized_Request
		ϒstr_to_int = Ωutils.ϒstr_to_int
		ϒunescapeHTML = Ωutils.ϒunescapeHTML
		ϒunified_strdate = Ωutils.ϒunified_strdate
		ϒurl_or_none = Ωutils.ϒurl_or_none
		YouPornIE = λ.Cal(λ.TypeType, λ.NewStr("YouPornIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				YouPornIE__TESTS        λ.Object
				YouPornIE__VALID_URL    λ.Object
				YouPornIE__real_extract λ.Object
			)
			YouPornIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?youporn\\.com/watch/(?P<id>\\d+)/(?P<display_id>[^/?#&]+)")
			YouPornIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.youporn.com/watch/505835/sex-ed-is-it-safe-to-masturbate-daily/"),
					λ.NewStr("md5"): λ.NewStr("3744d24c50438cf5b6f6d59feb5055c2"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):             λ.NewStr("505835"),
						λ.NewStr("display_id"):     λ.NewStr("sex-ed-is-it-safe-to-masturbate-daily"),
						λ.NewStr("ext"):            λ.NewStr("mp4"),
						λ.NewStr("title"):          λ.NewStr("Sex Ed: Is It Safe To Masturbate Daily?"),
						λ.NewStr("description"):    λ.NewStr("Love & Sex Answers: http://bit.ly/DanAndJenn -- Is It Unhealthy To Masturbate Daily?"),
						λ.NewStr("thumbnail"):      λ.NewStr("re:^https?://.*\\.jpg$"),
						λ.NewStr("uploader"):       λ.NewStr("Ask Dan And Jennifer"),
						λ.NewStr("upload_date"):    λ.NewStr("20101217"),
						λ.NewStr("average_rating"): λ.IntType,
						λ.NewStr("view_count"):     λ.IntType,
						λ.NewStr("comment_count"):  λ.IntType,
						λ.NewStr("categories"):     λ.ListType,
						λ.NewStr("tags"):           λ.ListType,
						λ.NewStr("age_limit"):      λ.NewInt(18),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.youporn.com/watch/561726/big-tits-awesome-brunette-on-amazing-webcam-show/?from=related3&al=2&from_id=561726&pos=4"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):             λ.NewStr("561726"),
						λ.NewStr("display_id"):     λ.NewStr("big-tits-awesome-brunette-on-amazing-webcam-show"),
						λ.NewStr("ext"):            λ.NewStr("mp4"),
						λ.NewStr("title"):          λ.NewStr("Big Tits Awesome Brunette On amazing webcam show"),
						λ.NewStr("description"):    λ.NewStr("http://sweetlivegirls.com Big Tits Awesome Brunette On amazing webcam show.mp4"),
						λ.NewStr("thumbnail"):      λ.NewStr("re:^https?://.*\\.jpg$"),
						λ.NewStr("uploader"):       λ.NewStr("Unknown"),
						λ.NewStr("upload_date"):    λ.NewStr("20110418"),
						λ.NewStr("average_rating"): λ.IntType,
						λ.NewStr("view_count"):     λ.IntType,
						λ.NewStr("comment_count"):  λ.IntType,
						λ.NewStr("categories"):     λ.ListType,
						λ.NewStr("tags"):           λ.ListType,
						λ.NewStr("age_limit"):      λ.NewInt(18),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
			)
			YouPornIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒage_limit       λ.Object
						ϒaverage_rating  λ.Object
						ϒbitrate         λ.Object
						ϒcategories      λ.Object
						ϒcomment_count   λ.Object
						ϒdefinition      λ.Object
						ϒdefinitions     λ.Object
						ϒdescription     λ.Object
						ϒdisplay_id      λ.Object
						ϒencrypted_link  λ.Object
						ϒextract_tag_box λ.Object
						ϒf               λ.Object
						ϒformats         λ.Object
						ϒheight          λ.Object
						ϒlink            λ.Object
						ϒlinks           λ.Object
						ϒmobj            λ.Object
						ϒrequest         λ.Object
						ϒself            = λargs[0]
						ϒsources         λ.Object
						ϒtags            λ.Object
						ϒthumbnail       λ.Object
						ϒtitle           λ.Object
						ϒupload_date     λ.Object
						ϒuploader        λ.Object
						ϒurl             = λargs[1]
						ϒvideo_id        λ.Object
						ϒvideo_url       λ.Object
						ϒview_count      λ.Object
						ϒwebpage         λ.Object
						τmp0             λ.Object
						τmp1             λ.Object
						τmp2             λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					ϒdisplay_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("display_id"))
					ϒrequest = λ.Cal(ϒsanitized_Request, ϒurl)
					λ.Cal(λ.GetAttr(ϒrequest, "add_header", nil), λ.NewStr("Cookie"), λ.NewStr("age_verified=1"))
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒrequest, ϒdisplay_id)
					ϒtitle = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewList(
								λ.NewStr("(?:video_titles|videoTitle)\\s*[:=]\\s*([\"\\'])(?P<title>(?:(?!\\1).)+)\\1"),
								λ.NewStr("<h1[^>]+class=[\"\\']heading\\d?[\"\\'][^>]*>(?P<title>[^<]+)<"),
							),
							ϒwebpage,
							λ.NewStr("title"),
						), λ.KWArgs{
							{Name: "group", Value: λ.NewStr("title")},
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else if λv := λ.Call(λ.GetAttr(ϒself, "_og_search_title", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
								λ.NewStr("title"),
								ϒwebpage,
							), λ.KWArgs{
								{Name: "fatal", Value: λ.True},
							})
						}
					}()
					ϒlinks = λ.NewList()
					ϒdefinitions = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
						λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("mediaDefinition\\s*=\\s*(\\[.+?\\]);"),
							ϒwebpage,
							λ.NewStr("media definitions"),
						), λ.KWArgs{
							{Name: "default", Value: λ.NewStr("[]")},
						}),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					if λ.IsTrue(ϒdefinitions) {
						τmp0 = λ.Cal(λ.BuiltinIter, ϒdefinitions)
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒdefinition = τmp1
							if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒdefinition, λ.DictType)))) {
								continue
							}
							ϒvideo_url = λ.Cal(ϒurl_or_none, λ.Cal(λ.GetAttr(ϒdefinition, "get", nil), λ.NewStr("videoUrl")))
							if λ.IsTrue(ϒvideo_url) {
								λ.Cal(λ.GetAttr(ϒlinks, "append", nil), ϒvideo_url)
							}
						}
					}
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfindall, λ.NewStr("<a[^>]+href=([\"\\'])(http.+?)\\1[^>]+title=[\"\\']Download [Vv]ideo"), ϒwebpage))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						_ = λ.GetItem(τmp2, λ.NewInt(0))
						ϒlink = λ.GetItem(τmp2, λ.NewInt(1))
						λ.Cal(λ.GetAttr(ϒlinks, "append", nil), ϒlink)
					}
					ϒsources = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("(?s)sources\\s*:\\s*({.+?})"),
						ϒwebpage,
						λ.NewStr("sources"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒsources) {
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfindall, λ.NewStr("[^:]+\\s*:\\s*([\"\\'])(http.+?)\\1"), ϒsources))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							τmp2 = τmp1
							_ = λ.GetItem(τmp2, λ.NewInt(0))
							ϒlink = λ.GetItem(τmp2, λ.NewInt(1))
							λ.Cal(λ.GetAttr(ϒlinks, "append", nil), ϒlink)
						}
					}
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfindall, λ.NewStr("(?:videoSrc|videoIpadUrl|html5PlayerSrc)\\s*[:=]\\s*([\"\\'])(http.+?)\\1"), ϒwebpage))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						_ = λ.GetItem(τmp2, λ.NewInt(0))
						ϒlink = λ.GetItem(τmp2, λ.NewInt(1))
						λ.Cal(λ.GetAttr(ϒlinks, "append", nil), ϒlink)
					}
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfindall, λ.NewStr("encryptedQuality\\d{3,4}URL\\s*=\\s*([\"\\'])([\\da-zA-Z+/=]+)\\1"), ϒwebpage))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						_ = λ.GetItem(τmp2, λ.NewInt(0))
						ϒencrypted_link = λ.GetItem(τmp2, λ.NewInt(1))
						λ.Cal(λ.GetAttr(ϒlinks, "append", nil), λ.Cal(λ.GetAttr(λ.Cal(λ.None, ϒencrypted_link, ϒtitle, λ.NewInt(32)), "decode", nil), λ.NewStr("utf-8")))
					}
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.SetType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒlink λ.Object
									τmp0  λ.Object
									τmp1  λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, ϒlinks)
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒlink = τmp1
									λgen.Yield(λ.Cal(ϒunescapeHTML, ϒlink))
								}
								return λ.None
							})
						}))))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒvideo_url = τmp1
						ϒf = λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"): ϒvideo_url,
						})
						ϒmobj = λ.Cal(Ωre.ϒsearch, λ.NewStr("(?P<height>\\d{3,4})[pP]_(?P<bitrate>\\d+)[kK]_\\d+/"), ϒvideo_url)
						if λ.IsTrue(ϒmobj) {
							ϒheight = λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("height")))
							ϒbitrate = λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("bitrate")))
							λ.Cal(λ.GetAttr(ϒf, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("format_id"): λ.Mod(λ.NewStr("%dp-%dk"), λ.NewTuple(
									ϒheight,
									ϒbitrate,
								)),
								λ.NewStr("height"): ϒheight,
								λ.NewStr("tbr"):    ϒbitrate,
							}))
						}
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒf)
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒdescription = λ.Call(λ.GetAttr(ϒself, "_og_search_description", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					ϒthumbnail = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("(?:imageurl\\s*=|poster\\s*:)\\s*([\"\\'])(?P<thumbnail>.+?)\\1"),
						ϒwebpage,
						λ.NewStr("thumbnail"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
						{Name: "group", Value: λ.NewStr("thumbnail")},
					})
					ϒuploader = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("(?s)<div[^>]+class=[\"\\']submitByLink[\"\\'][^>]*>(.+?)</div>"),
						ϒwebpage,
						λ.NewStr("uploader"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒupload_date = λ.Cal(ϒunified_strdate, λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewList(
							λ.NewStr("Date\\s+[Aa]dded:\\s*<span>([^<]+)"),
							λ.NewStr("(?s)<div[^>]+class=[\"\\']videoInfo(?:Date|Time)[\"\\'][^>]*>(.+?)</div>"),
						),
						ϒwebpage,
						λ.NewStr("upload date"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒage_limit = λ.Cal(λ.GetAttr(ϒself, "_rta_search", nil), ϒwebpage)
					ϒaverage_rating = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("<div[^>]+class=[\"\\']videoRatingPercentage[\"\\'][^>]*>(\\d+)%</div>"),
						ϒwebpage,
						λ.NewStr("average rating"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒview_count = λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("(?s)<div[^>]+class=([\"\\']).*?\\bvideoInfoViews\\b.*?\\1[^>]*>.*?(?P<count>[\\d,.]+)<"),
						ϒwebpage,
						λ.NewStr("view count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
						{Name: "group", Value: λ.NewStr("count")},
					}))
					ϒcomment_count = λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr(">All [Cc]omments? \\(([\\d,.]+)\\)"),
						ϒwebpage,
						λ.NewStr("comment count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒextract_tag_box = λ.NewFunction("extract_tag_box",
						[]λ.Param{
							{Name: "regex"},
							{Name: "title"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒregex   = λargs[0]
								ϒtag_box λ.Object
								ϒtitle   = λargs[1]
							)
							ϒtag_box = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								ϒregex,
								ϒwebpage,
								ϒtitle,
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							})
							if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒtag_box))) {
								return λ.NewList()
							}
							return λ.Cal(Ωre.ϒfindall, λ.NewStr("<a[^>]+href=[^>]+>([^<]+)"), ϒtag_box)
						})
					ϒcategories = λ.Cal(ϒextract_tag_box, λ.NewStr("(?s)Categories:.*?</[^>]+>(.+?)</div>"), λ.NewStr("categories"))
					ϒtags = λ.Cal(ϒextract_tag_box, λ.NewStr("(?s)Tags:.*?</div>\\s*<div[^>]+class=[\"\\']tagBoxContent[\"\\'][^>]*>(.+?)</div>"), λ.NewStr("tags"))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):             ϒvideo_id,
						λ.NewStr("display_id"):     ϒdisplay_id,
						λ.NewStr("title"):          ϒtitle,
						λ.NewStr("description"):    ϒdescription,
						λ.NewStr("thumbnail"):      ϒthumbnail,
						λ.NewStr("uploader"):       ϒuploader,
						λ.NewStr("upload_date"):    ϒupload_date,
						λ.NewStr("average_rating"): ϒaverage_rating,
						λ.NewStr("view_count"):     ϒview_count,
						λ.NewStr("comment_count"):  ϒcomment_count,
						λ.NewStr("categories"):     ϒcategories,
						λ.NewStr("tags"):           ϒtags,
						λ.NewStr("age_limit"):      ϒage_limit,
						λ.NewStr("formats"):        ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        YouPornIE__TESTS,
				λ.NewStr("_VALID_URL"):    YouPornIE__VALID_URL,
				λ.NewStr("_real_extract"): YouPornIE__real_extract,
			})
		}())
	})
}