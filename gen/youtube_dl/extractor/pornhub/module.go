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
 * pornhub/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/pornhub.py
 */

package pornhub

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError        λ.Object
	InfoExtractor         λ.Object
	PornHubIE             λ.Object
	PornHubPlaylistBaseIE λ.Object
	PornHubPlaylistIE     λ.Object
	PornHubUserVideosIE   λ.Object
	ϒcompat_str           λ.Object
	ϒint_or_none          λ.Object
	ϒjs_to_json           λ.Object
	ϒorderedSet           λ.Object
	ϒstr_to_int           λ.Object
	ϒurl_or_none          λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒjs_to_json = Ωutils.ϒjs_to_json
		ϒorderedSet = Ωutils.ϒorderedSet
		ϒstr_to_int = Ωutils.ϒstr_to_int
		ϒurl_or_none = Ωutils.ϒurl_or_none
		PornHubIE = λ.Cal(λ.TypeType, λ.NewStr("PornHubIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				PornHubIE__TESTS         λ.Object
				PornHubIE__VALID_URL     λ.Object
				PornHubIE__extract_count λ.Object
				PornHubIE__real_extract  λ.Object
			)
			PornHubIE__VALID_URL = λ.NewStr("(?x)\n                    https?://\n                        (?:\n                            (?:[^/]+\\.)?(?P<host>pornhub\\.(?:com|net))/(?:(?:view_video\\.php|video/show)\\?viewkey=|embed/)|\n                            (?:www\\.)?thumbzilla\\.com/video/\n                        )\n                        (?P<id>[\\da-z]+)\n                    ")
			PornHubIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.pornhub.com/view_video.php?viewkey=648719015"),
					λ.NewStr("md5"): λ.NewStr("1e19b41231a02eba417839222ac9d58e"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):            λ.NewStr("648719015"),
						λ.NewStr("ext"):           λ.NewStr("mp4"),
						λ.NewStr("title"):         λ.NewStr("Seductive Indian beauty strips down and fingers her pink pussy"),
						λ.NewStr("uploader"):      λ.NewStr("Babes"),
						λ.NewStr("upload_date"):   λ.NewStr("20130628"),
						λ.NewStr("duration"):      λ.NewInt(361),
						λ.NewStr("view_count"):    λ.IntType,
						λ.NewStr("like_count"):    λ.IntType,
						λ.NewStr("dislike_count"): λ.IntType,
						λ.NewStr("comment_count"): λ.IntType,
						λ.NewStr("age_limit"):     λ.NewInt(18),
						λ.NewStr("tags"):          λ.ListType,
						λ.NewStr("categories"):    λ.ListType,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.pornhub.com/view_video.php?viewkey=1331683002"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):            λ.NewStr("1331683002"),
						λ.NewStr("ext"):           λ.NewStr("mp4"),
						λ.NewStr("title"):         λ.NewStr("重庆婷婷女王足交"),
						λ.NewStr("uploader"):      λ.NewStr("Unknown"),
						λ.NewStr("upload_date"):   λ.NewStr("20150213"),
						λ.NewStr("duration"):      λ.NewInt(1753),
						λ.NewStr("view_count"):    λ.IntType,
						λ.NewStr("like_count"):    λ.IntType,
						λ.NewStr("dislike_count"): λ.IntType,
						λ.NewStr("comment_count"): λ.IntType,
						λ.NewStr("age_limit"):     λ.NewInt(18),
						λ.NewStr("tags"):          λ.ListType,
						λ.NewStr("categories"):    λ.ListType,
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.pornhub.com/view_video.php?viewkey=ph5af5fef7c2aa7"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):            λ.NewStr("ph5af5fef7c2aa7"),
						λ.NewStr("ext"):           λ.NewStr("mp4"),
						λ.NewStr("title"):         λ.NewStr("BFFS - Cute Teen Girls Share Cock On the Floor"),
						λ.NewStr("uploader"):      λ.NewStr("BFFs"),
						λ.NewStr("duration"):      λ.NewInt(622),
						λ.NewStr("view_count"):    λ.IntType,
						λ.NewStr("like_count"):    λ.IntType,
						λ.NewStr("dislike_count"): λ.IntType,
						λ.NewStr("comment_count"): λ.IntType,
						λ.NewStr("age_limit"):     λ.NewInt(18),
						λ.NewStr("tags"):          λ.ListType,
						λ.NewStr("categories"):    λ.ListType,
						λ.NewStr("subtitles"): λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("en"): λ.NewList(λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("ext"): λ.NewStr("srt"),
							})),
						}),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.pornhub.com/view_video.php?viewkey=ph557bbb6676d2d"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://fr.pornhub.com/view_video.php?viewkey=ph55ca2f9760862"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.pornhub.com/view_video.php?viewkey=788152859"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.pornhub.com/view_video.php?viewkey=ph572716d15a111"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.pornhub.com/view_video.php?viewkey=ph56fd731fce6b7"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.thumbzilla.com/video/ph56c6114abd99a/horny-girlfriend-sex"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.pornhub.com/video/show?viewkey=648719015"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.pornhub.net/view_video.php?viewkey=203640933"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			PornHubIE__extract_count = λ.NewFunction("_extract_count",
				[]λ.Param{
					{Name: "self"},
					{Name: "pattern"},
					{Name: "webpage"},
					{Name: "name"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒname    = λargs[3]
						ϒpattern = λargs[1]
						ϒself    = λargs[0]
						ϒwebpage = λargs[2]
					)
					return λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						ϒpattern,
						ϒwebpage,
						λ.Mod(λ.NewStr("%s count"), ϒname),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
				})
			PornHubIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒassignments       λ.Object
						ϒassn              λ.Object
						ϒcategories        λ.Object
						ϒcomment_count     λ.Object
						ϒdefinition        λ.Object
						ϒdislike_count     λ.Object
						ϒdl_webpage        λ.Object
						ϒduration          λ.Object
						ϒerror_msg         λ.Object
						ϒflashvars         λ.Object
						ϒformats           λ.Object
						ϒheight            λ.Object
						ϒhost              λ.Object
						ϒjs_vars           λ.Object
						ϒlike_count        λ.Object
						ϒmedia_definitions λ.Object
						ϒmobj              λ.Object
						ϒpage_params       λ.Object
						ϒparse_js_value    λ.Object
						ϒself              = λargs[0]
						ϒsubtitle_url      λ.Object
						ϒsubtitles         λ.Object
						ϒtags              λ.Object
						ϒtbr               λ.Object
						ϒthumbnail         λ.Object
						ϒtitle             λ.Object
						ϒtv_webpage        λ.Object
						ϒupload_date       λ.Object
						ϒurl               = λargs[1]
						ϒvalue             λ.Object
						ϒvideo_id          λ.Object
						ϒvideo_uploader    λ.Object
						ϒvideo_url         λ.Object
						ϒvideo_urls        λ.Object
						ϒvideo_urls_set    λ.Object
						ϒview_count        λ.Object
						ϒvname             λ.Object
						ϒwebpage           λ.Object
						τmp0               λ.Object
						τmp1               λ.Object
						τmp2               λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒhost = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("host")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewStr("pornhub.com")
						}
					}()
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					λ.Cal(λ.GetAttr(ϒself, "_set_cookie", nil), ϒhost, λ.NewStr("age_verified"), λ.NewStr("1"))
					ϒdl_webpage = λ.NewFunction("dl_webpage",
						[]λ.Param{
							{Name: "platform"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒplatform = λargs[0]
							)
							λ.Cal(λ.GetAttr(ϒself, "_set_cookie", nil), ϒhost, λ.NewStr("platform"), ϒplatform)
							return λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), λ.Mod(λ.NewStr("http://www.%s/view_video.php?viewkey=%s"), λ.NewTuple(
								ϒhost,
								ϒvideo_id,
							)), ϒvideo_id, λ.Mod(λ.NewStr("Downloading %s webpage"), ϒplatform))
						})
					ϒwebpage = λ.Cal(ϒdl_webpage, λ.NewStr("pc"))
					ϒerror_msg = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("(?s)<div[^>]+class=([\"\\'])(?:(?!\\1).)*\\b(?:removed|userMessageSection)\\b(?:(?!\\1).)*\\1[^>]*>(?P<error>.+?)</div>"),
						ϒwebpage,
						λ.NewStr("error message"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
						{Name: "group", Value: λ.NewStr("error")},
					})
					if λ.IsTrue(ϒerror_msg) {
						ϒerror_msg = λ.Cal(Ωre.ϒsub, λ.NewStr("\\s+"), λ.NewStr(" "), ϒerror_msg)
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("PornHub said: %s"), ϒerror_msg)), λ.KWArgs{
							{Name: "expected", Value: λ.True},
							{Name: "video_id", Value: ϒvideo_id},
						})))
					}
					ϒtitle = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
							λ.NewStr("twitter:title"),
							ϒwebpage,
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.NewTuple(
									λ.NewStr("<h1[^>]+class=[\"\\']title[\"\\'][^>]*>(?P<title>[^<]+)"),
									λ.NewStr("<div[^>]+data-video-title=([\"\\'])(?P<title>.+?)\\1"),
									λ.NewStr("shareTitle\\s*=\\s*([\"\\'])(?P<title>.+?)\\1"),
								),
								ϒwebpage,
								λ.NewStr("title"),
							), λ.KWArgs{
								{Name: "group", Value: λ.NewStr("title")},
							})
						}
					}()
					ϒvideo_urls = λ.NewList()
					ϒvideo_urls_set = λ.Cal(λ.SetType)
					ϒsubtitles = λ.NewDictWithTable(map[λ.Object]λ.Object{})
					ϒflashvars = λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("var\\s+flashvars_\\d+\\s*=\\s*({.+?});"),
						ϒwebpage,
						λ.NewStr("flashvars"),
					), λ.KWArgs{
						{Name: "default", Value: λ.NewStr("{}")},
					}), ϒvideo_id)
					if λ.IsTrue(ϒflashvars) {
						ϒsubtitle_url = λ.Cal(ϒurl_or_none, λ.Cal(λ.GetAttr(ϒflashvars, "get", nil), λ.NewStr("closedCaptionsFile")))
						if λ.IsTrue(ϒsubtitle_url) {
							λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒsubtitles, "setdefault", nil), λ.NewStr("en"), λ.NewList()), "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("url"): ϒsubtitle_url,
								λ.NewStr("ext"): λ.NewStr("srt"),
							}))
						}
						ϒthumbnail = λ.Cal(λ.GetAttr(ϒflashvars, "get", nil), λ.NewStr("image_url"))
						ϒduration = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒflashvars, "get", nil), λ.NewStr("video_duration")))
						ϒmedia_definitions = λ.Cal(λ.GetAttr(ϒflashvars, "get", nil), λ.NewStr("mediaDefinitions"))
						if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒmedia_definitions, λ.ListType)) {
							τmp0 = λ.Cal(λ.BuiltinIter, ϒmedia_definitions)
							for {
								if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
									break
								}
								ϒdefinition = τmp1
								if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒdefinition, λ.DictType)))) {
									continue
								}
								ϒvideo_url = λ.Cal(λ.GetAttr(ϒdefinition, "get", nil), λ.NewStr("videoUrl"))
								if λ.IsTrue(func() λ.Object {
									if λv := λ.NewBool(!λ.IsTrue(ϒvideo_url)); λ.IsTrue(λv) {
										return λv
									} else {
										return λ.NewBool(!λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒvideo_url, ϒcompat_str)))
									}
								}()) {
									continue
								}
								if λ.IsTrue(λ.NewBool(λ.Contains(ϒvideo_urls_set, ϒvideo_url))) {
									continue
								}
								λ.Cal(λ.GetAttr(ϒvideo_urls_set, "add", nil), ϒvideo_url)
								λ.Cal(λ.GetAttr(ϒvideo_urls, "append", nil), λ.NewTuple(
									ϒvideo_url,
									λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒdefinition, "get", nil), λ.NewStr("quality"))),
								))
							}
						}
					} else {
						τmp0 = λ.Mul(λ.NewList(λ.None), λ.NewInt(2))
						ϒthumbnail = λ.GetItem(τmp0, λ.NewInt(0))
						ϒduration = λ.GetItem(τmp0, λ.NewInt(1))
					}
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒvideo_urls))) {
						ϒtv_webpage = λ.Cal(ϒdl_webpage, λ.NewStr("tv"))
						ϒassignments = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("(var.+?mediastring.+?)</script>"), ϒtv_webpage, λ.NewStr("encoded url")), "split", nil), λ.NewStr(";"))
						ϒjs_vars = λ.NewDictWithTable(map[λ.Object]λ.Object{})
						ϒparse_js_value = λ.NewFunction("parse_js_value",
							[]λ.Param{
								{Name: "inp"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒinp  = λargs[0]
									ϒinps λ.Object
								)
								ϒinp = λ.Cal(Ωre.ϒsub, λ.NewStr("/\\*(?:(?!\\*/).)*?\\*/"), λ.NewStr(""), ϒinp)
								if λ.IsTrue(λ.NewBool(λ.Contains(ϒinp, λ.NewStr("+")))) {
									ϒinps = λ.Cal(λ.GetAttr(ϒinp, "split", nil), λ.NewStr("+"))
									return λ.Cal(λ.GetAttr(λ.None, "reduce", nil), λ.None, λ.Cal(λ.MapIteratorType, ϒparse_js_value, ϒinps))
								}
								ϒinp = λ.Cal(λ.GetAttr(ϒinp, "strip", nil))
								if λ.IsTrue(λ.NewBool(λ.Contains(ϒjs_vars, ϒinp))) {
									return λ.GetItem(ϒjs_vars, ϒinp)
								}
								return λ.Cal(λ.None, ϒinp)
							})
						τmp0 = λ.Cal(λ.BuiltinIter, ϒassignments)
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒassn = τmp1
							ϒassn = λ.Cal(λ.GetAttr(ϒassn, "strip", nil))
							if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒassn))) {
								continue
							}
							ϒassn = λ.Cal(Ωre.ϒsub, λ.NewStr("var\\s+"), λ.NewStr(""), ϒassn)
							τmp2 = λ.Cal(λ.GetAttr(ϒassn, "split", nil), λ.NewStr("="), λ.NewInt(1))
							ϒvname = λ.GetItem(τmp2, λ.NewInt(0))
							ϒvalue = λ.GetItem(τmp2, λ.NewInt(1))
							λ.SetItem(ϒjs_vars, ϒvname, λ.Cal(ϒparse_js_value, ϒvalue))
						}
						ϒvideo_url = λ.GetItem(ϒjs_vars, λ.NewStr("mediastring"))
						if λ.IsTrue(λ.NewBool(!λ.Contains(ϒvideo_urls_set, ϒvideo_url))) {
							λ.Cal(λ.GetAttr(ϒvideo_urls, "append", nil), λ.NewTuple(
								ϒvideo_url,
								λ.None,
							))
							λ.Cal(λ.GetAttr(ϒvideo_urls_set, "add", nil), ϒvideo_url)
						}
					}
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfinditer, λ.NewStr("<a[^>]+\\bclass=[\"\\']downloadBtn\\b[^>]+\\bhref=([\"\\'])(?P<url>(?:(?!\\1).)+)\\1"), ϒwebpage))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒmobj = τmp1
						ϒvideo_url = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("url"))
						if λ.IsTrue(λ.NewBool(!λ.Contains(ϒvideo_urls_set, ϒvideo_url))) {
							λ.Cal(λ.GetAttr(ϒvideo_urls, "append", nil), λ.NewTuple(
								ϒvideo_url,
								λ.None,
							))
							λ.Cal(λ.GetAttr(ϒvideo_urls_set, "add", nil), ϒvideo_url)
						}
					}
					ϒupload_date = λ.None
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, ϒvideo_urls)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒvideo_url = λ.GetItem(τmp2, λ.NewInt(0))
						ϒheight = λ.GetItem(τmp2, λ.NewInt(1))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒupload_date))) {
							ϒupload_date = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.NewStr("/(\\d{6}/\\d{2})/"),
								ϒvideo_url,
								λ.NewStr("upload data"),
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							})
							if λ.IsTrue(ϒupload_date) {
								ϒupload_date = λ.Cal(λ.GetAttr(ϒupload_date, "replace", nil), λ.NewStr("/"), λ.NewStr(""))
							}
						}
						ϒtbr = λ.None
						ϒmobj = λ.Cal(Ωre.ϒsearch, λ.NewStr("(?P<height>\\d+)[pP]?_(?P<tbr>\\d+)[kK]"), ϒvideo_url)
						if λ.IsTrue(ϒmobj) {
							if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒheight))) {
								ϒheight = λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("height")))
							}
							ϒtbr = λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("tbr")))
						}
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"): ϒvideo_url,
							λ.NewStr("format_id"): func() λ.Object {
								if λ.IsTrue(ϒheight) {
									return λ.Mod(λ.NewStr("%dp"), ϒheight)
								} else {
									return λ.None
								}
							}(),
							λ.NewStr("height"): ϒheight,
							λ.NewStr("tbr"):    ϒtbr,
						}))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒvideo_uploader = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("(?s)From:&nbsp;.+?<(?:a\\b[^>]+\\bhref=[\"\\']/(?:(?:user|channel)s|model|pornstar)/|span\\b[^>]+\\bclass=[\"\\']username)[^>]+>(.+?)<"),
						ϒwebpage,
						λ.NewStr("uploader"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒview_count = λ.Cal(λ.GetAttr(ϒself, "_extract_count", nil), λ.NewStr("<span class=\"count\">([\\d,\\.]+)</span> views"), ϒwebpage, λ.NewStr("view"))
					ϒlike_count = λ.Cal(λ.GetAttr(ϒself, "_extract_count", nil), λ.NewStr("<span class=\"votesUp\">([\\d,\\.]+)</span>"), ϒwebpage, λ.NewStr("like"))
					ϒdislike_count = λ.Cal(λ.GetAttr(ϒself, "_extract_count", nil), λ.NewStr("<span class=\"votesDown\">([\\d,\\.]+)</span>"), ϒwebpage, λ.NewStr("dislike"))
					ϒcomment_count = λ.Cal(λ.GetAttr(ϒself, "_extract_count", nil), λ.NewStr("All Comments\\s*<span>\\(([\\d,.]+)\\)"), ϒwebpage, λ.NewStr("comment"))
					ϒpage_params = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
						λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("page_params\\.zoneDetails\\[([\\'\"])[^\\'\"]+\\1\\]\\s*=\\s*(?P<data>{[^}]+})"),
							ϒwebpage,
							λ.NewStr("page parameters"),
						), λ.KWArgs{
							{Name: "group", Value: λ.NewStr("data")},
							{Name: "default", Value: λ.NewStr("{}")},
						}),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "transform_source", Value: ϒjs_to_json},
						{Name: "fatal", Value: λ.False},
					})
					τmp0 = λ.None
					ϒtags = τmp0
					ϒcategories = τmp0
					if λ.IsTrue(ϒpage_params) {
						ϒtags = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒpage_params, "get", nil), λ.NewStr("tags"), λ.NewStr("")), "split", nil), λ.NewStr(","))
						ϒcategories = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒpage_params, "get", nil), λ.NewStr("categories"), λ.NewStr("")), "split", nil), λ.NewStr(","))
					}
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):            ϒvideo_id,
						λ.NewStr("uploader"):      ϒvideo_uploader,
						λ.NewStr("upload_date"):   ϒupload_date,
						λ.NewStr("title"):         ϒtitle,
						λ.NewStr("thumbnail"):     ϒthumbnail,
						λ.NewStr("duration"):      ϒduration,
						λ.NewStr("view_count"):    ϒview_count,
						λ.NewStr("like_count"):    ϒlike_count,
						λ.NewStr("dislike_count"): ϒdislike_count,
						λ.NewStr("comment_count"): ϒcomment_count,
						λ.NewStr("formats"):       ϒformats,
						λ.NewStr("age_limit"):     λ.NewInt(18),
						λ.NewStr("tags"):          ϒtags,
						λ.NewStr("categories"):    ϒcategories,
						λ.NewStr("subtitles"):     ϒsubtitles,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):         PornHubIE__TESTS,
				λ.NewStr("_VALID_URL"):     PornHubIE__VALID_URL,
				λ.NewStr("_extract_count"): PornHubIE__extract_count,
				λ.NewStr("_real_extract"):  PornHubIE__real_extract,
			})
		}())
		PornHubPlaylistBaseIE = λ.Cal(λ.TypeType, λ.NewStr("PornHubPlaylistBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {

			return λ.NewDictWithTable(map[λ.Object]λ.Object{})
		}())
		PornHubPlaylistIE = λ.Cal(λ.TypeType, λ.NewStr("PornHubPlaylistIE"), λ.NewTuple(PornHubPlaylistBaseIE), func() λ.Dict {
			var (
				PornHubPlaylistIE__VALID_URL λ.Object
			)
			PornHubPlaylistIE__VALID_URL = λ.NewStr("https?://(?:[^/]+\\.)?(?P<host>pornhub\\.(?:com|net))/playlist/(?P<id>\\d+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): PornHubPlaylistIE__VALID_URL,
			})
		}())
		PornHubUserVideosIE = λ.Cal(λ.TypeType, λ.NewStr("PornHubUserVideosIE"), λ.NewTuple(PornHubPlaylistBaseIE), func() λ.Dict {
			var (
				PornHubUserVideosIE__VALID_URL λ.Object
			)
			PornHubUserVideosIE__VALID_URL = λ.NewStr("https?://(?:[^/]+\\.)?(?P<host>pornhub\\.(?:com|net))/(?:(?:user|channel)s|model|pornstar)/(?P<id>[^/]+)/videos")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): PornHubUserVideosIE__VALID_URL,
			})
		}())
	})
}
