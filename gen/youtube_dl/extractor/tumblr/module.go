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
 * tumblr/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/tumblr.py
 */

package tumblr

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
	TumblrIE            λ.Object
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
		TumblrIE = λ.Cal(λ.TypeType, λ.NewStr("TumblrIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TumblrIE__NETRC_MACHINE   λ.Object
				TumblrIE__TESTS           λ.Object
				TumblrIE__VALID_URL       λ.Object
				TumblrIE__login           λ.Object
				TumblrIE__real_extract    λ.Object
				TumblrIE__real_initialize λ.Object
			)
			TumblrIE__VALID_URL = λ.NewStr("https?://(?P<blog_name>[^/?#&]+)\\.tumblr\\.com/(?:post|video)/(?P<id>[0-9]+)(?:$|[/?#])")
			TumblrIE__NETRC_MACHINE = λ.NewStr("tumblr")
			TumblrIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://tatianamaslanydaily.tumblr.com/post/54196191430/orphan-black-dvd-extra-behind-the-scenes"),
					λ.NewStr("md5"): λ.NewStr("479bb068e5b16462f5176a6828829767"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("54196191430"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("tatiana maslany news, Orphan Black || DVD extra - behind the scenes ↳..."),
						λ.NewStr("description"): λ.NewStr("md5:37db8211e40b50c7c44e95da14f630b7"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:http://.*\\.jpg"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://5sostrum.tumblr.com/post/90208453769/yall-forgetting-the-greatest-keek-of-them-all"),
					λ.NewStr("md5"): λ.NewStr("bf348ef8c0ef84fbf1cbd6fa6e000359"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("90208453769"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("5SOS STRUM ;]"),
						λ.NewStr("description"): λ.NewStr("md5:dba62ac8639482759c8eb10ce474586a"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:http://.*\\.jpg"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://hdvideotest.tumblr.com/post/130323439814/test-description-for-my-hd-video"),
					λ.NewStr("md5"): λ.NewStr("7ae503065ad150122dc3089f8cf1546c"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("130323439814"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("HD Video Testing — Test description for my HD video"),
						λ.NewStr("description"): λ.NewStr("md5:97cc3ab5fcd27ee4af6356701541319c"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:http://.*\\.jpg"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("format"): λ.NewStr("hd"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://naked-yogi.tumblr.com/post/118312946248/naked-smoking-stretching"),
					λ.NewStr("md5"): λ.NewStr("de07e5211d60d4f3a2c3df757ea9f6ab"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("Wmur"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("naked smoking & stretching"),
						λ.NewStr("upload_date"): λ.NewStr("20150506"),
						λ.NewStr("timestamp"):   λ.NewInt(1430931613),
						λ.NewStr("age_limit"):   λ.NewInt(18),
						λ.NewStr("uploader_id"): λ.NewStr("1638622"),
						λ.NewStr("uploader"):    λ.NewStr("naked-yogi"),
					}),
					λ.NewStr("add_ie"): λ.NewList(λ.NewStr("Vidme")),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://camdamage.tumblr.com/post/98846056295/"),
					λ.NewStr("md5"): λ.NewStr("a9e0c8371ea1ca306d6554e3fecf50b6"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("105463834"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Cam Damage-HD 720p"),
						λ.NewStr("uploader"):    λ.NewStr("John Moyer"),
						λ.NewStr("uploader_id"): λ.NewStr("user32021558"),
					}),
					λ.NewStr("add_ie"): λ.NewList(λ.NewStr("Vimeo")),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://sutiblr.tumblr.com/post/139638707273"),
					λ.NewStr("md5"): λ.NewStr("2dd184b3669e049ba40563a7d423f95c"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):            λ.NewStr("ir7qBEIKqvq"),
						λ.NewStr("ext"):           λ.NewStr("mp4"),
						λ.NewStr("title"):         λ.NewStr("Vine by sutiblr"),
						λ.NewStr("alt_title"):     λ.NewStr("Vine by sutiblr"),
						λ.NewStr("uploader"):      λ.NewStr("sutiblr"),
						λ.NewStr("uploader_id"):   λ.NewStr("1198993975374495744"),
						λ.NewStr("upload_date"):   λ.NewStr("20160220"),
						λ.NewStr("like_count"):    λ.IntType,
						λ.NewStr("comment_count"): λ.IntType,
						λ.NewStr("repost_count"):  λ.IntType,
					}),
					λ.NewStr("add_ie"): λ.NewList(λ.NewStr("Vine")),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://vitasidorkina.tumblr.com/post/134652425014/joskriver-victoriassecret-invisibility-or"),
					λ.NewStr("md5"): λ.NewStr("01c12ceb82cbf6b2fe0703aa56b3ad72"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("-7LnUPGlSo"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Video by victoriassecret"),
						λ.NewStr("description"): λ.NewStr("Invisibility or flight…which superpower would YOU choose? #VSFashionShow #ThisOrThat"),
						λ.NewStr("uploader_id"): λ.NewStr("victoriassecret"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.jpg"),
					}),
					λ.NewStr("add_ie"): λ.NewList(λ.NewStr("Instagram")),
				}),
			)
			TumblrIE__real_initialize = λ.NewFunction("_real_initialize",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
					)
					λ.Cal(λ.GetAttr(ϒself, "_login", nil))
					return λ.None
				})
			TumblrIE__login = λ.NewFunction("_login",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒlogin_errors λ.Object
						ϒlogin_form   λ.Object
						ϒlogin_page   λ.Object
						ϒpassword     λ.Object
						ϒresponse     λ.Object
						ϒself         = λargs[0]
						ϒurlh         λ.Object
						ϒusername     λ.Object
						τmp0          λ.Object
					)
					τmp0 = λ.Cal(λ.GetAttr(ϒself, "_get_login_info", nil))
					ϒusername = λ.GetItem(τmp0, λ.NewInt(0))
					ϒpassword = λ.GetItem(τmp0, λ.NewInt(1))
					if λ.IsTrue(λ.NewBool(ϒusername == λ.None)) {
						return λ.None
					}
					ϒlogin_page = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), λ.GetAttr(ϒself, "_LOGIN_URL", nil), λ.None, λ.NewStr("Downloading login page"))
					ϒlogin_form = λ.Cal(λ.GetAttr(ϒself, "_hidden_inputs", nil), ϒlogin_page)
					λ.Cal(λ.GetAttr(ϒlogin_form, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("user[email]"):    ϒusername,
						λ.NewStr("user[password]"): ϒpassword,
					}))
					τmp0 = λ.Call(λ.GetAttr(ϒself, "_download_webpage_handle", nil), λ.NewArgs(
						λ.GetAttr(ϒself, "_LOGIN_URL", nil),
						λ.None,
						λ.NewStr("Logging in"),
					), λ.KWArgs{
						{Name: "data", Value: λ.Cal(ϒurlencode_postdata, ϒlogin_form)},
						{Name: "headers", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("Content-Type"): λ.NewStr("application/x-www-form-urlencoded"),
							λ.NewStr("Referer"):      λ.GetAttr(ϒself, "_LOGIN_URL", nil),
						})},
					})
					ϒresponse = λ.GetItem(τmp0, λ.NewInt(0))
					ϒurlh = λ.GetItem(τmp0, λ.NewInt(1))
					if λ.IsTrue(λ.NewBool(λ.Contains(λ.Cal(λ.GetAttr(ϒurlh, "geturl", nil)), λ.NewStr("/dashboard")))) {
						return λ.None
					}
					ϒlogin_errors = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
						λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("RegistrationForm\\.errors\\s*=\\s*(\\[.+?\\])\\s*;"),
							ϒresponse,
							λ.NewStr("login errors"),
						), λ.KWArgs{
							{Name: "default", Value: λ.NewStr("[]")},
						}),
						λ.None,
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					if λ.IsTrue(ϒlogin_errors) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("Unable to login: %s"), λ.GetItem(ϒlogin_errors, λ.NewInt(0)))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					λ.Cal(λ.GetAttr(ϒself, "report_warning", nil), λ.NewStr("Login has probably failed"))
					return λ.None
				})
			TumblrIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒblog         λ.Object
						ϒduration     λ.Object
						ϒformats      λ.Object
						ϒhd_url       λ.Object
						ϒiframe       λ.Object
						ϒiframe_url   λ.Object
						ϒm_url        λ.Object
						ϒoptions      λ.Object
						ϒredirect_url λ.Object
						ϒsd_url       λ.Object
						ϒself         = λargs[0]
						ϒsources      λ.Object
						ϒurl          = λargs[1]
						ϒurlh         λ.Object
						ϒvideo_id     λ.Object
						ϒvideo_title  λ.Object
						ϒwebpage      λ.Object
						τmp0          λ.Object
					)
					ϒm_url = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒm_url, "group", nil), λ.NewStr("id"))
					ϒblog = λ.Cal(λ.GetAttr(ϒm_url, "group", nil), λ.NewStr("blog_name"))
					ϒurl = λ.Mod(λ.NewStr("http://%s.tumblr.com/post/%s/"), λ.NewTuple(
						ϒblog,
						ϒvideo_id,
					))
					τmp0 = λ.Cal(λ.GetAttr(ϒself, "_download_webpage_handle", nil), ϒurl, ϒvideo_id)
					ϒwebpage = λ.GetItem(τmp0, λ.NewInt(0))
					ϒurlh = λ.GetItem(τmp0, λ.NewInt(1))
					ϒredirect_url = λ.Cal(ϒcompat_str, λ.Cal(λ.GetAttr(ϒurlh, "geturl", nil)))
					if λ.IsTrue(func() λ.Object {
						if λv := λ.NewBool(λ.Contains(ϒredirect_url, λ.NewStr("tumblr.com/safe-mode"))); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒredirect_url, "startswith", nil), λ.NewStr("/safe-mode"))
						}
					}()) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.NewStr("This Tumblr may contain sensitive media. Disable safe mode in your account settings at https://www.tumblr.com/settings/account#safe_mode")), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒiframe_url = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("src=\\'(https?://www\\.tumblr\\.com/video/[^\\']+)\\'"),
						ϒwebpage,
						λ.NewStr("iframe url"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(λ.NewBool(ϒiframe_url == λ.None)) {
						return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), ϒredirect_url, λ.NewStr("Generic"))
					}
					ϒiframe = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒiframe_url, ϒvideo_id, λ.NewStr("Downloading iframe page"))
					ϒduration = λ.None
					ϒsources = λ.NewList()
					ϒsd_url = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("<source[^>]+src=([\"\\'])(?P<url>.+?)\\1"),
						ϒiframe,
						λ.NewStr("sd video url"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
						{Name: "group", Value: λ.NewStr("url")},
					})
					if λ.IsTrue(ϒsd_url) {
						λ.Cal(λ.GetAttr(ϒsources, "append", nil), λ.NewTuple(
							ϒsd_url,
							λ.NewStr("sd"),
						))
					}
					ϒoptions = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
						λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("data-crt-options=([\"\\'])(?P<options>.+?)\\1"),
							ϒiframe,
							λ.NewStr("hd video url"),
						), λ.KWArgs{
							{Name: "default", Value: λ.NewStr("")},
							{Name: "group", Value: λ.NewStr("options")},
						}),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					if λ.IsTrue(ϒoptions) {
						ϒduration = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒoptions, "get", nil), λ.NewStr("duration")))
						ϒhd_url = λ.Cal(λ.GetAttr(ϒoptions, "get", nil), λ.NewStr("hdUrl"))
						if λ.IsTrue(ϒhd_url) {
							λ.Cal(λ.GetAttr(ϒsources, "append", nil), λ.NewTuple(
								ϒhd_url,
								λ.NewStr("hd"),
							))
						}
					}
					ϒformats = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒformat_id λ.Object
									ϒquality   λ.Object
									ϒvideo_url λ.Object
									τmp0       λ.Object
									τmp1       λ.Object
									τmp2       λ.Object
									τmp3       λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.EnumerateIteratorType, ϒsources))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									τmp2 = τmp1
									ϒquality = λ.GetItem(τmp2, λ.NewInt(0))
									τmp3 = λ.GetItem(τmp2, λ.NewInt(1))
									ϒvideo_url = λ.GetItem(τmp3, λ.NewInt(0))
									ϒformat_id = λ.GetItem(τmp3, λ.NewInt(1))
									λgen.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("url"):       ϒvideo_url,
										λ.NewStr("ext"):       λ.NewStr("mp4"),
										λ.NewStr("format_id"): ϒformat_id,
										λ.NewStr("height"): λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
											λ.NewStr("/(\\d{3,4})$"),
											ϒvideo_url,
											λ.NewStr("height"),
										), λ.KWArgs{
											{Name: "default", Value: λ.None},
										})),
										λ.NewStr("quality"): ϒquality,
									}))
								}
								return λ.None
							})
						})))
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒvideo_title = λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewStr("(?s)<title>(?P<title>.*?)(?: \\| Tumblr)?</title>"), ϒwebpage, λ.NewStr("title"))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):    ϒvideo_id,
						λ.NewStr("title"): ϒvideo_title,
						λ.NewStr("description"): λ.Call(λ.GetAttr(ϒself, "_og_search_description", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}),
						λ.NewStr("thumbnail"): λ.Call(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}),
						λ.NewStr("duration"): ϒduration,
						λ.NewStr("formats"):  ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_NETRC_MACHINE"):   TumblrIE__NETRC_MACHINE,
				λ.NewStr("_TESTS"):           TumblrIE__TESTS,
				λ.NewStr("_VALID_URL"):       TumblrIE__VALID_URL,
				λ.NewStr("_login"):           TumblrIE__login,
				λ.NewStr("_real_extract"):    TumblrIE__real_extract,
				λ.NewStr("_real_initialize"): TumblrIE__real_initialize,
			})
		}())
	})
}