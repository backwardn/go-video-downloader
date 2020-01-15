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
 * facebook/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/facebook.py
 */

package facebook

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωsocket "github.com/tenta-browser/go-video-downloader/gen/socket"
	Ωerror "github.com/tenta-browser/go-video-downloader/gen/urllib/error"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError                    λ.Object
	FacebookIE                        λ.Object
	FacebookPluginsVideoIE            λ.Object
	InfoExtractor                     λ.Object
	ϒclean_html                       λ.Object
	ϒcompat_etree_fromstring          λ.Object
	ϒcompat_http_client               λ.Object
	ϒcompat_urllib_parse_unquote      λ.Object
	ϒcompat_urllib_parse_unquote_plus λ.Object
	ϒerror_to_compat_str              λ.Object
	ϒget_element_by_id                λ.Object
	ϒint_or_none                      λ.Object
	ϒjs_to_json                       λ.Object
	ϒlimit_length                     λ.Object
	ϒparse_count                      λ.Object
	ϒsanitized_Request                λ.Object
	ϒtry_get                          λ.Object
	ϒurlencode_postdata               λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_etree_fromstring = Ωcompat.ϒcompat_etree_fromstring
		ϒcompat_http_client = Ωcompat.ϒcompat_http_client
		ϒcompat_urllib_parse_unquote = Ωcompat.ϒcompat_urllib_parse_unquote
		ϒcompat_urllib_parse_unquote_plus = Ωcompat.ϒcompat_urllib_parse_unquote_plus
		ϒclean_html = Ωutils.ϒclean_html
		ϒerror_to_compat_str = Ωutils.ϒerror_to_compat_str
		ExtractorError = Ωutils.ExtractorError
		ϒget_element_by_id = Ωutils.ϒget_element_by_id
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒjs_to_json = Ωutils.ϒjs_to_json
		ϒlimit_length = Ωutils.ϒlimit_length
		ϒparse_count = Ωutils.ϒparse_count
		ϒsanitized_Request = Ωutils.ϒsanitized_Request
		ϒtry_get = Ωutils.ϒtry_get
		ϒurlencode_postdata = Ωutils.ϒurlencode_postdata
		FacebookIE = λ.Cal(λ.TypeType, λ.StrLiteral("FacebookIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				FacebookIE_IE_NAME                    λ.Object
				FacebookIE__CHROME_USER_AGENT         λ.Object
				FacebookIE__NETRC_MACHINE             λ.Object
				FacebookIE__VALID_URL                 λ.Object
				FacebookIE__VIDEO_PAGE_TAHOE_TEMPLATE λ.Object
				FacebookIE__VIDEO_PAGE_TEMPLATE       λ.Object
				FacebookIE__extract_from_url          λ.Object
				FacebookIE__login                     λ.Object
				FacebookIE__real_extract              λ.Object
				FacebookIE__real_initialize           λ.Object
			)
			FacebookIE__VALID_URL = λ.StrLiteral("(?x)\n                (?:\n                    https?://\n                        (?:[\\w-]+\\.)?(?:facebook\\.com|facebookcorewwwi\\.onion)/\n                        (?:[^#]*?\\#!/)?\n                        (?:\n                            (?:\n                                video/video\\.php|\n                                photo\\.php|\n                                video\\.php|\n                                video/embed|\n                                story\\.php\n                            )\\?(?:.*?)(?:v|video_id|story_fbid)=|\n                            [^/]+/videos/(?:[^/]+/)?|\n                            [^/]+/posts/|\n                            groups/[^/]+/permalink/\n                        )|\n                    facebook:\n                )\n                (?P<id>[0-9]+)\n                ")
			FacebookIE__NETRC_MACHINE = λ.StrLiteral("facebook")
			FacebookIE_IE_NAME = λ.StrLiteral("facebook")
			FacebookIE__CHROME_USER_AGENT = λ.StrLiteral("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.97 Safari/537.36")
			FacebookIE__VIDEO_PAGE_TEMPLATE = λ.StrLiteral("https://www.facebook.com/video/video.php?v=%s")
			FacebookIE__VIDEO_PAGE_TAHOE_TEMPLATE = λ.StrLiteral("https://www.facebook.com/video/tahoe/async/%s/?chain=true&isvideo=true&payloadtype=primary")
			FacebookIE__login = λ.NewFunction("_login",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcheck_form     λ.Object
						ϒcheck_req      λ.Object
						ϒcheck_response λ.Object
						ϒerror          λ.Object
						ϒfb_dtsg        λ.Object
						ϒh              λ.Object
						ϒlgnrnd         λ.Object
						ϒlogin_form     λ.Object
						ϒlogin_page     λ.Object
						ϒlogin_page_req λ.Object
						ϒlogin_results  λ.Object
						ϒlsd            λ.Object
						ϒpassword       λ.Object
						ϒrequest        λ.Object
						ϒself           = λargs[0]
						ϒuseremail      λ.Object
						τmp0            λ.Object
						τmp1            λ.Object
					)
					τmp0 = λ.Calm(ϒself, "_get_login_info")
					ϒuseremail = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒpassword = λ.GetItem(τmp0, λ.IntLiteral(1))
					if ϒuseremail == λ.None {
						return λ.None
					}
					ϒlogin_page_req = λ.Cal(ϒsanitized_Request, λ.GetAttr(ϒself, "_LOGIN_URL", nil))
					λ.Calm(ϒself, "_set_cookie", λ.StrLiteral("facebook.com"), λ.StrLiteral("locale"), λ.StrLiteral("en_US"))
					ϒlogin_page = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
						ϒlogin_page_req,
						λ.None,
					), λ.KWArgs{
						{Name: "note", Value: λ.StrLiteral("Downloading login page")},
						{Name: "errnote", Value: λ.StrLiteral("Unable to download login page")},
					})
					ϒlsd = λ.Calm(ϒself, "_search_regex", λ.StrLiteral("<input type=\"hidden\" name=\"lsd\" value=\"([^\"]*)\""), ϒlogin_page, λ.StrLiteral("lsd"))
					ϒlgnrnd = λ.Calm(ϒself, "_search_regex", λ.StrLiteral("name=\"lgnrnd\" value=\"([^\"]*?)\""), ϒlogin_page, λ.StrLiteral("lgnrnd"))
					ϒlogin_form = λ.DictLiteral(map[string]λ.Object{
						"email":              ϒuseremail,
						"pass":               ϒpassword,
						"lsd":                ϒlsd,
						"lgnrnd":             ϒlgnrnd,
						"next":               λ.StrLiteral("http://facebook.com/home.php"),
						"default_persistent": λ.StrLiteral("0"),
						"legacy_return":      λ.StrLiteral("1"),
						"timezone":           λ.StrLiteral("-60"),
						"trynum":             λ.StrLiteral("1"),
					})
					ϒrequest = λ.Cal(ϒsanitized_Request, λ.GetAttr(ϒself, "_LOGIN_URL", nil), λ.Cal(ϒurlencode_postdata, ϒlogin_form))
					λ.Calm(ϒrequest, "add_header", λ.StrLiteral("Content-Type"), λ.StrLiteral("application/x-www-form-urlencoded"))
					τmp0, τmp1 = func() (λexit λ.Object, λret λ.Object) {
						defer λ.CatchMulti(
							nil,
							&λ.Catcher{λ.NewTuple(
								Ωerror.URLError,
								λ.GetAttr(ϒcompat_http_client, "HTTPException", nil),
								Ωsocket.ϒerror,
							), func(λex λ.BaseException) {
								var ϒerr λ.Object = λex
								λ.Calm(λ.GetAttr(ϒself, "_downloader", nil), "report_warning", λ.Mod(λ.StrLiteral("unable to log in: %s"), λ.Cal(ϒerror_to_compat_str, ϒerr)))
								λexit, λret = λ.BlockExitReturn, λ.None
								return
							}},
						)
						ϒlogin_results = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
							ϒrequest,
							λ.None,
						), λ.KWArgs{
							{Name: "note", Value: λ.StrLiteral("Logging in")},
							{Name: "errnote", Value: λ.StrLiteral("unable to fetch login page")},
						})
						if λ.Cal(Ωre.ϒsearch, λ.StrLiteral("<form(.*)name=\"login\"(.*)</form>"), ϒlogin_results) != λ.None {
							ϒerror = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("(?s)<div[^>]+class=([\"\\']).*?login_error_box.*?\\1[^>]*><div[^>]*>.*?</div><div[^>]*>(?P<error>.+?)</div>"),
								ϒlogin_results,
								λ.StrLiteral("login error"),
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
								{Name: "group", Value: λ.StrLiteral("error")},
							})
							if λ.IsTrue(ϒerror) {
								panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.StrLiteral("Unable to login: %s"), ϒerror)), λ.KWArgs{
									{Name: "expected", Value: λ.True},
								})))
							}
							λ.Calm(λ.GetAttr(ϒself, "_downloader", nil), "report_warning", λ.StrLiteral("unable to log in: bad username/password, or exceeded login rate limit (~3/min). Check credentials or wait."))
							λexit, λret = λ.BlockExitReturn, λ.None
							return
						}
						ϒfb_dtsg = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.StrLiteral("name=\"fb_dtsg\" value=\"(.+?)\""),
							ϒlogin_results,
							λ.StrLiteral("fb_dtsg"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
						ϒh = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.StrLiteral("name=\"h\"\\s+(?:\\w+=\"[^\"]+\"\\s+)*?value=\"([^\"]+)\""),
							ϒlogin_results,
							λ.StrLiteral("h"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
						if λ.IsTrue(func() λ.Object {
							if λv := λ.NewBool(!λ.IsTrue(ϒfb_dtsg)); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewBool(!λ.IsTrue(ϒh))
							}
						}()) {
							λexit, λret = λ.BlockExitReturn, λ.None
							return
						}
						ϒcheck_form = λ.DictLiteral(map[string]λ.Object{
							"fb_dtsg":              ϒfb_dtsg,
							"h":                    ϒh,
							"name_action_selected": λ.StrLiteral("dont_save"),
						})
						ϒcheck_req = λ.Cal(ϒsanitized_Request, λ.GetAttr(ϒself, "_CHECKPOINT_URL", nil), λ.Cal(ϒurlencode_postdata, ϒcheck_form))
						λ.Calm(ϒcheck_req, "add_header", λ.StrLiteral("Content-Type"), λ.StrLiteral("application/x-www-form-urlencoded"))
						ϒcheck_response = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
							ϒcheck_req,
							λ.None,
						), λ.KWArgs{
							{Name: "note", Value: λ.StrLiteral("Confirming login")},
						})
						if λ.Cal(Ωre.ϒsearch, λ.StrLiteral("id=\"checkpointSubmitButton\""), ϒcheck_response) != λ.None {
							λ.Calm(λ.GetAttr(ϒself, "_downloader", nil), "report_warning", λ.StrLiteral("Unable to confirm login, you have to login in your browser and authorize the login."))
						}
						return λ.BlockExitNormally, nil
					}()
					if τmp0 == λ.BlockExitReturn {
						return τmp1
					}
					return λ.None
				})
			FacebookIE__real_initialize = λ.NewFunction("_real_initialize",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
					)
					λ.Calm(ϒself, "_login")
					return λ.None
				})
			FacebookIE__extract_from_url = λ.NewFunction("_extract_from_url",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
					{Name: "video_id"},
					{Name: "fatal_if_no_video", Def: λ.True},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdash_manifest                 λ.Object
						ϒextract_from_jsmods_instances λ.Object
						ϒextract_video_data            λ.Object
						ϒf                             λ.Object
						ϒfatal_if_no_video             = λargs[3]
						ϒformat_id                     λ.Object
						ϒformats                       λ.Object
						ϒinfo_dict                     λ.Object
						ϒm_msg                         λ.Object
						ϒpreference                    λ.Object
						ϒquality                       λ.Object
						ϒreq                           λ.Object
						ϒself                          = λargs[0]
						ϒserver_js_data                λ.Object
						ϒsrc                           λ.Object
						ϒsrc_type                      λ.Object
						ϒsubtitles                     λ.Object
						ϒsubtitles_src                 λ.Object
						ϒtahoe_data                    λ.Object
						ϒtahoe_js_data                 λ.Object
						ϒthumbnail                     λ.Object
						ϒtimestamp                     λ.Object
						ϒuploader                      λ.Object
						ϒurl                           = λargs[1]
						ϒvideo_data                    λ.Object
						ϒvideo_id                      = λargs[2]
						ϒvideo_title                   λ.Object
						ϒview_count                    λ.Object
						ϒwebpage                       λ.Object
						τmp0                           λ.Object
						τmp1                           λ.Object
						τmp2                           λ.Object
						τmp3                           λ.Object
						τmp4                           λ.Object
						τmp5                           λ.Object
						τmp6                           λ.Object
					)
					ϒreq = λ.Cal(ϒsanitized_Request, ϒurl)
					λ.Calm(ϒreq, "add_header", λ.StrLiteral("User-Agent"), λ.GetAttr(ϒself, "_CHROME_USER_AGENT", nil))
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒreq, ϒvideo_id)
					ϒvideo_data = λ.None
					ϒextract_video_data = λ.NewFunction("extract_video_data",
						[]λ.Param{
							{Name: "instances"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒinstances  = λargs[0]
								ϒitem       λ.Object
								ϒvideo_item λ.Object
								τmp0        λ.Object
								τmp1        λ.Object
							)
							τmp0 = λ.Cal(λ.BuiltinIter, ϒinstances)
							for {
								if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
									break
								}
								ϒitem = τmp1
								if λ.IsTrue(λ.Eq(λ.GetItem(λ.GetItem(ϒitem, λ.IntLiteral(1)), λ.IntLiteral(0)), λ.StrLiteral("VideoConfig"))) {
									ϒvideo_item = λ.GetItem(λ.GetItem(ϒitem, λ.IntLiteral(2)), λ.IntLiteral(0))
									if λ.IsTrue(λ.Calm(ϒvideo_item, "get", λ.StrLiteral("video_id"))) {
										return λ.GetItem(ϒvideo_item, λ.StrLiteral("videoData"))
									}
								}
							}
							return λ.None
						})
					ϒserver_js_data = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
						λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.StrLiteral("handleServerJS\\(({.+})(?:\\);|,\")"),
							ϒwebpage,
							λ.StrLiteral("server js data"),
						), λ.KWArgs{
							{Name: "default", Value: λ.StrLiteral("{}")},
						}),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					if λ.IsTrue(ϒserver_js_data) {
						ϒvideo_data = λ.Cal(ϒextract_video_data, λ.Calm(ϒserver_js_data, "get", λ.StrLiteral("instances"), λ.NewList()))
					}
					ϒextract_from_jsmods_instances = λ.NewFunction("extract_from_jsmods_instances",
						[]λ.Param{
							{Name: "js_data"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒjs_data = λargs[0]
							)
							if λ.IsTrue(ϒjs_data) {
								return λ.Cal(ϒextract_video_data, func() λ.Object {
									if λv := λ.Cal(ϒtry_get, ϒjs_data, λ.NewFunction("<lambda>",
										[]λ.Param{
											{Name: "x"},
										},
										0, false, false,
										func(λargs []λ.Object) λ.Object {
											var (
												ϒx = λargs[0]
											)
											return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("jsmods")), λ.StrLiteral("instances"))
										}), λ.ListType); λ.IsTrue(λv) {
										return λv
									} else {
										return λ.NewList()
									}
								}())
							}
							return λ.None
						})
					if !λ.IsTrue(ϒvideo_data) {
						ϒserver_js_data = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
							λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("bigPipe\\.onPageletArrive\\(({.+?})\\)\\s*;\\s*}\\s*\\)\\s*,\\s*[\"\\']onPageletArrive\\s+(?:pagelet_group_mall|permalink_video_pagelet|hyperfeed_story_id_\\d+)"),
								ϒwebpage,
								λ.StrLiteral("js data"),
							), λ.KWArgs{
								{Name: "default", Value: λ.StrLiteral("{}")},
							}),
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "transform_source", Value: ϒjs_to_json},
							{Name: "fatal", Value: λ.False},
						})
						ϒvideo_data = λ.Cal(ϒextract_from_jsmods_instances, ϒserver_js_data)
					}
					if !λ.IsTrue(ϒvideo_data) {
						if !λ.IsTrue(ϒfatal_if_no_video) {
							return λ.NewTuple(
								ϒwebpage,
								λ.False,
							)
						}
						ϒm_msg = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("class=\"[^\"]*uiInterstitialContent[^\"]*\"><div>(.*?)</div>"), ϒwebpage)
						if ϒm_msg != λ.None {
							panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.StrLiteral("The video is not available, Facebook said: \"%s\""), λ.Calm(ϒm_msg, "group", λ.IntLiteral(1)))), λ.KWArgs{
								{Name: "expected", Value: λ.True},
							})))
						} else {
							if λ.Contains(ϒwebpage, λ.StrLiteral(">You must log in to continue")) {
								λ.Calm(ϒself, "raise_login_required")
							}
						}
						ϒtahoe_data = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
							λ.Mod(λ.GetAttr(ϒself, "_VIDEO_PAGE_TAHOE_TEMPLATE", nil), ϒvideo_id),
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "data", Value: λ.Cal(ϒurlencode_postdata, λ.DictLiteral(map[string]λ.Object{
								"__a": λ.IntLiteral(1),
								"__pc": λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
									λ.StrLiteral("pkg_cohort[\"\\']\\s*:\\s*[\"\\'](.+?)[\"\\']"),
									ϒwebpage,
									λ.StrLiteral("pkg cohort"),
								), λ.KWArgs{
									{Name: "default", Value: λ.StrLiteral("PHASED:DEFAULT")},
								}),
								"__rev": λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
									λ.StrLiteral("client_revision[\"\\']\\s*:\\s*(\\d+),"),
									ϒwebpage,
									λ.StrLiteral("client revision"),
								), λ.KWArgs{
									{Name: "default", Value: λ.StrLiteral("3944515")},
								}),
								"fb_dtsg": λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
									λ.StrLiteral("\"DTSGInitialData\"\\s*,\\s*\\[\\]\\s*,\\s*{\\s*\"token\"\\s*:\\s*\"([^\"]+)\""),
									ϒwebpage,
									λ.StrLiteral("dtsg token"),
								), λ.KWArgs{
									{Name: "default", Value: λ.StrLiteral("")},
								}),
							}))},
							{Name: "headers", Value: λ.DictLiteral(map[string]string{
								"Content-Type": "application/x-www-form-urlencoded",
							})},
						})
						ϒtahoe_js_data = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
							λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("for\\s+\\(\\s*;\\s*;\\s*\\)\\s*;(.+)"),
								ϒtahoe_data,
								λ.StrLiteral("tahoe js data"),
							), λ.KWArgs{
								{Name: "default", Value: λ.StrLiteral("{}")},
							}),
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "fatal", Value: λ.False},
						})
						ϒvideo_data = λ.Cal(ϒextract_from_jsmods_instances, ϒtahoe_js_data)
					}
					if !λ.IsTrue(ϒvideo_data) {
						panic(λ.Raise(λ.Cal(ExtractorError, λ.StrLiteral("Cannot parse data"))))
					}
					ϒsubtitles = λ.DictLiteral(map[λ.Object]λ.Object{})
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, ϒvideo_data)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒf = τmp1
						ϒformat_id = λ.GetItem(ϒf, λ.StrLiteral("stream_type"))
						if λ.IsTrue(func() λ.Object {
							if λv := ϒf; !λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(λ.BuiltinIsInstance, ϒf, λ.DictType)
							}
						}()) {
							ϒf = λ.NewList(ϒf)
						}
						if λ.IsTrue(func() λ.Object {
							if λv := λ.NewBool(!λ.IsTrue(ϒf)); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewBool(!λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒf, λ.ListType)))
							}
						}()) {
							continue
						}
						τmp2 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
							λ.StrLiteral("sd"),
							λ.StrLiteral("hd"),
						))
						for {
							if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
								break
							}
							ϒquality = τmp3
							τmp4 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
								λ.StrLiteral("src"),
								λ.StrLiteral("src_no_ratelimit"),
							))
							for {
								if τmp5 = λ.NextDefault(τmp4, λ.AfterLast); τmp5 == λ.AfterLast {
									break
								}
								ϒsrc_type = τmp5
								ϒsrc = λ.Calm(λ.GetItem(ϒf, λ.IntLiteral(0)), "get", λ.Mod(λ.StrLiteral("%s_%s"), λ.NewTuple(
									ϒquality,
									ϒsrc_type,
								)))
								if λ.IsTrue(ϒsrc) {
									ϒpreference = func() λ.Object {
										if λ.IsTrue(λ.Eq(ϒformat_id, λ.StrLiteral("progressive"))) {
											return λ.Neg(λ.IntLiteral(10))
										} else {
											return λ.IntLiteral(0)
										}
									}()
									if λ.IsTrue(λ.Eq(ϒquality, λ.StrLiteral("hd"))) {
										τmp6 = λ.IAdd(ϒpreference, λ.IntLiteral(5))
										ϒpreference = τmp6
									}
									λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
										"format_id": λ.Mod(λ.StrLiteral("%s_%s_%s"), λ.NewTuple(
											ϒformat_id,
											ϒquality,
											ϒsrc_type,
										)),
										"url":        ϒsrc,
										"preference": ϒpreference,
									}))
								}
							}
						}
						ϒdash_manifest = λ.Calm(λ.GetItem(ϒf, λ.IntLiteral(0)), "get", λ.StrLiteral("dash_manifest"))
						if λ.IsTrue(ϒdash_manifest) {
							λ.Calm(ϒformats, "extend", λ.Calm(ϒself, "_parse_mpd_formats", λ.Cal(ϒcompat_etree_fromstring, λ.Cal(ϒcompat_urllib_parse_unquote_plus, ϒdash_manifest))))
						}
						ϒsubtitles_src = λ.Calm(λ.GetItem(ϒf, λ.IntLiteral(0)), "get", λ.StrLiteral("subtitles_src"))
						if λ.IsTrue(ϒsubtitles_src) {
							λ.Calm(λ.Calm(ϒsubtitles, "setdefault", λ.StrLiteral("en"), λ.NewList()), "append", λ.DictLiteral(map[string]λ.Object{
								"url": ϒsubtitles_src,
							}))
						}
					}
					if !λ.IsTrue(ϒformats) {
						panic(λ.Raise(λ.Cal(ExtractorError, λ.StrLiteral("Cannot find video formats"))))
					}
					τmp0 = λ.Cal(λ.BuiltinIter, ϒformats)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒf = τmp1
						λ.SetItem(λ.Calm(ϒf, "setdefault", λ.StrLiteral("http_headers"), λ.DictLiteral(map[λ.Object]λ.Object{})), λ.StrLiteral("User-Agent"), λ.StrLiteral("facebookexternalhit/1.1"))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒvideo_title = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("<h2\\s+[^>]*class=\"uiHeaderTitle\"[^>]*>([^<]*)</h2>"),
						ϒwebpage,
						λ.StrLiteral("title"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if !λ.IsTrue(ϒvideo_title) {
						ϒvideo_title = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
							λ.StrLiteral("(?s)<span class=\"fbPhotosPhotoCaption\".*?id=\"fbPhotoPageCaption\"><span class=\"hasCaption\">(.*?)</span>"),
							ϒwebpage,
							λ.StrLiteral("alternative title"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
					}
					if !λ.IsTrue(ϒvideo_title) {
						ϒvideo_title = λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
							λ.StrLiteral("description"),
							ϒwebpage,
							λ.StrLiteral("title"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
					}
					if λ.IsTrue(ϒvideo_title) {
						ϒvideo_title = λ.Cal(ϒlimit_length, ϒvideo_title, λ.IntLiteral(80))
					} else {
						ϒvideo_title = λ.Mod(λ.StrLiteral("Facebook video #%s"), ϒvideo_id)
					}
					ϒuploader = func() λ.Object {
						if λv := λ.Cal(ϒclean_html, λ.Cal(ϒget_element_by_id, λ.StrLiteral("fbPhotoPageAuthorName"), ϒwebpage)); λ.IsTrue(λv) {
							return λv
						} else if λv := λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.StrLiteral("ownerName\\s*:\\s*\"([^\"]+)\""),
							ϒwebpage,
							λ.StrLiteral("uploader"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Call(λ.GetAttr(ϒself, "_og_search_title", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
								{Name: "fatal", Value: λ.False},
							})
						}
					}()
					ϒtimestamp = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("<abbr[^>]+data-utime=[\"\\'](\\d+)"),
						ϒwebpage,
						λ.StrLiteral("timestamp"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					}))
					ϒthumbnail = λ.Calm(ϒself, "_html_search_meta", λ.NewList(
						λ.StrLiteral("og:image"),
						λ.StrLiteral("twitter:image"),
					), ϒwebpage)
					ϒview_count = λ.Cal(ϒparse_count, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("\\bviewCount\\s*:\\s*[\"\\']([\\d,.]+)"),
						ϒwebpage,
						λ.StrLiteral("view count"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					}))
					ϒinfo_dict = λ.DictLiteral(map[string]λ.Object{
						"id":         ϒvideo_id,
						"title":      ϒvideo_title,
						"formats":    ϒformats,
						"uploader":   ϒuploader,
						"timestamp":  ϒtimestamp,
						"thumbnail":  ϒthumbnail,
						"view_count": ϒview_count,
						"subtitles":  ϒsubtitles,
					})
					return λ.NewTuple(
						ϒwebpage,
						ϒinfo_dict,
					)
				})
			FacebookIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒentries   λ.Object
						ϒinfo_dict λ.Object
						ϒreal_url  λ.Object
						ϒself      = λargs[0]
						ϒurl       = λargs[1]
						ϒvideo_id  λ.Object
						ϒwebpage   λ.Object
						τmp0       λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒreal_url = func() λ.Object {
						if λ.IsTrue(λ.Calm(ϒurl, "startswith", λ.StrLiteral("facebook:"))) {
							return λ.Mod(λ.GetAttr(ϒself, "_VIDEO_PAGE_TEMPLATE", nil), ϒvideo_id)
						} else {
							return ϒurl
						}
					}()
					τmp0 = λ.Call(λ.GetAttr(ϒself, "_extract_from_url", nil), λ.NewArgs(
						ϒreal_url,
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "fatal_if_no_video", Value: λ.False},
					})
					ϒwebpage = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒinfo_dict = λ.GetItem(τmp0, λ.IntLiteral(1))
					if λ.IsTrue(ϒinfo_dict) {
						return ϒinfo_dict
					}
					if λ.Contains(ϒurl, λ.StrLiteral("/posts/")) {
						ϒentries = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
							nil,
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
									var (
										ϒvid λ.Object
										τmp0 λ.Object
										τmp1 λ.Object
									)
									τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒself, "_parse_json", λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
										λ.StrLiteral("([\"\\'])video_ids\\1\\s*:\\s*(?P<ids>\\[.+?\\])"),
										ϒwebpage,
										λ.StrLiteral("video ids"),
									), λ.KWArgs{
										{Name: "group", Value: λ.StrLiteral("ids")},
									}), ϒvideo_id))
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										ϒvid = τmp1
										λgy.Yield(λ.Calm(ϒself, "url_result", λ.Mod(λ.StrLiteral("facebook:%s"), ϒvid), λ.Calm(FacebookIE, "ie_key")))
									}
									return λ.None
								})
							})))
						return λ.Calm(ϒself, "playlist_result", ϒentries, ϒvideo_id)
					} else {
						τmp0 = λ.Call(λ.GetAttr(ϒself, "_extract_from_url", nil), λ.NewArgs(
							λ.Mod(λ.GetAttr(ϒself, "_VIDEO_PAGE_TEMPLATE", nil), ϒvideo_id),
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "fatal_if_no_video", Value: λ.True},
						})
						_ = λ.GetItem(τmp0, λ.IntLiteral(0))
						ϒinfo_dict = λ.GetItem(τmp0, λ.IntLiteral(1))
						return ϒinfo_dict
					}
					return λ.None
				})
			return λ.DictLiteral(map[string]λ.Object{
				"IE_NAME":                    FacebookIE_IE_NAME,
				"_CHROME_USER_AGENT":         FacebookIE__CHROME_USER_AGENT,
				"_NETRC_MACHINE":             FacebookIE__NETRC_MACHINE,
				"_VALID_URL":                 FacebookIE__VALID_URL,
				"_VIDEO_PAGE_TAHOE_TEMPLATE": FacebookIE__VIDEO_PAGE_TAHOE_TEMPLATE,
				"_VIDEO_PAGE_TEMPLATE":       FacebookIE__VIDEO_PAGE_TEMPLATE,
				"_extract_from_url":          FacebookIE__extract_from_url,
				"_login":                     FacebookIE__login,
				"_real_extract":              FacebookIE__real_extract,
				"_real_initialize":           FacebookIE__real_initialize,
			})
		}())
		FacebookPluginsVideoIE = λ.Cal(λ.TypeType, λ.StrLiteral("FacebookPluginsVideoIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				FacebookPluginsVideoIE__VALID_URL    λ.Object
				FacebookPluginsVideoIE__real_extract λ.Object
			)
			FacebookPluginsVideoIE__VALID_URL = λ.StrLiteral("https?://(?:[\\w-]+\\.)?facebook\\.com/plugins/video\\.php\\?.*?\\bhref=(?P<id>https.+)")
			FacebookPluginsVideoIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
						ϒurl  = λargs[1]
					)
					return λ.Calm(ϒself, "url_result", λ.Cal(ϒcompat_urllib_parse_unquote, λ.Calm(ϒself, "_match_id", ϒurl)), λ.Calm(FacebookIE, "ie_key"))
				})
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL":    FacebookPluginsVideoIE__VALID_URL,
				"_real_extract": FacebookPluginsVideoIE__real_extract,
			})
		}())
	})
}
