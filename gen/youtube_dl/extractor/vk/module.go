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
 * vk/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/vk.py
 */

package vk

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωdailymotion "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/dailymotion"
	Ωodnoklassniki "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/odnoklassniki"
	Ωpladform "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/pladform"
	Ωvimeo "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/vimeo"
	Ωyoutube "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/youtube"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	DailymotionIE         λ.Object
	ExtractorError        λ.Object
	InfoExtractor         λ.Object
	OdnoklassnikiIE       λ.Object
	PladformIE            λ.Object
	VKBaseIE              λ.Object
	VKIE                  λ.Object
	VKUserVideosIE        λ.Object
	VKWallPostIE          λ.Object
	VimeoIE               λ.Object
	YoutubeIE             λ.Object
	ϒclean_html           λ.Object
	ϒget_element_by_class λ.Object
	ϒint_or_none          λ.Object
	ϒorderedSet           λ.Object
	ϒstr_or_none          λ.Object
	ϒstr_to_int           λ.Object
	ϒunescapeHTML         λ.Object
	ϒunified_timestamp    λ.Object
	ϒurl_or_none          λ.Object
	ϒurlencode_postdata   λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒclean_html = Ωutils.ϒclean_html
		ExtractorError = Ωutils.ExtractorError
		ϒget_element_by_class = Ωutils.ϒget_element_by_class
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒorderedSet = Ωutils.ϒorderedSet
		ϒstr_or_none = Ωutils.ϒstr_or_none
		ϒstr_to_int = Ωutils.ϒstr_to_int
		ϒunescapeHTML = Ωutils.ϒunescapeHTML
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ϒurl_or_none = Ωutils.ϒurl_or_none
		ϒurlencode_postdata = Ωutils.ϒurlencode_postdata
		DailymotionIE = Ωdailymotion.DailymotionIE
		OdnoklassnikiIE = Ωodnoklassniki.OdnoklassnikiIE
		PladformIE = Ωpladform.PladformIE
		VimeoIE = Ωvimeo.VimeoIE
		YoutubeIE = Ωyoutube.YoutubeIE
		VKBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("VKBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				VKBaseIE__NETRC_MACHINE    λ.Object
				VKBaseIE__download_payload λ.Object
				VKBaseIE__login            λ.Object
				VKBaseIE__real_initialize  λ.Object
			)
			VKBaseIE__NETRC_MACHINE = λ.StrLiteral("vk")
			VKBaseIE__login = λ.NewFunction("_login",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒlogin_form λ.Object
						ϒlogin_page λ.Object
						ϒpassword   λ.Object
						ϒself       = λargs[0]
						ϒurl_handle λ.Object
						ϒusername   λ.Object
						τmp0        λ.Object
					)
					τmp0 = λ.Calm(ϒself, "_get_login_info")
					ϒusername = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒpassword = λ.GetItem(τmp0, λ.IntLiteral(1))
					if ϒusername == λ.None {
						return λ.None
					}
					τmp0 = λ.Calm(ϒself, "_download_webpage_handle", λ.StrLiteral("https://vk.com"), λ.None, λ.StrLiteral("Downloading login page"))
					ϒlogin_page = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒurl_handle = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒlogin_form = λ.Calm(ϒself, "_hidden_inputs", ϒlogin_page)
					λ.Calm(ϒlogin_form, "update", λ.DictLiteral(map[string]λ.Object{
						"email": λ.Calm(ϒusername, "encode", λ.StrLiteral("cp1251")),
						"pass":  λ.Calm(ϒpassword, "encode", λ.StrLiteral("cp1251")),
					}))
					λ.Calm(ϒself, "_apply_first_set_cookie_header", ϒurl_handle, λ.StrLiteral("remixlhk"))
					ϒlogin_page = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
						λ.StrLiteral("https://login.vk.com/?act=login"),
						λ.None,
					), λ.KWArgs{
						{Name: "note", Value: λ.StrLiteral("Logging in")},
						{Name: "data", Value: λ.Cal(ϒurlencode_postdata, ϒlogin_form)},
					})
					if λ.IsTrue(λ.Cal(Ωre.ϒsearch, λ.StrLiteral("onLoginFailed"), ϒlogin_page)) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.StrLiteral("Unable to login, incorrect username and/or password")), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					return λ.None
				})
			VKBaseIE__real_initialize = λ.NewFunction("_real_initialize",
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
			VKBaseIE__download_payload = λ.NewFunction("_download_payload",
				[]λ.Param{
					{Name: "self"},
					{Name: "path"},
					{Name: "video_id"},
					{Name: "data"},
					{Name: "fatal", Def: λ.True},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcode     λ.Object
						ϒdata     = λargs[3]
						ϒfatal    = λargs[4]
						ϒpath     = λargs[1]
						ϒpayload  λ.Object
						ϒself     = λargs[0]
						ϒvideo_id = λargs[2]
						τmp0      λ.Object
					)
					λ.SetItem(ϒdata, λ.StrLiteral("al"), λ.IntLiteral(1))
					τmp0 = λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.Mod(λ.StrLiteral("https://vk.com/%s.php"), ϒpath),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "data", Value: λ.Cal(ϒurlencode_postdata, ϒdata)},
						{Name: "fatal", Value: ϒfatal},
						{Name: "headers", Value: λ.DictLiteral(map[string]string{
							"X-Requested-With": "XMLHttpRequest",
						})},
					}), λ.StrLiteral("payload"))
					ϒcode = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒpayload = λ.GetItem(τmp0, λ.IntLiteral(1))
					if λ.IsTrue(λ.Eq(ϒcode, λ.StrLiteral("3"))) {
						λ.Calm(ϒself, "raise_login_required")
					} else {
						if λ.IsTrue(λ.Eq(ϒcode, λ.StrLiteral("8"))) {
							panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Cal(ϒclean_html, λ.GetItem(λ.GetItem(ϒpayload, λ.IntLiteral(0)), λ.NewSlice(λ.IntLiteral(1), λ.Neg(λ.IntLiteral(1)), λ.None)))), λ.KWArgs{
								{Name: "expected", Value: λ.True},
							})))
						}
					}
					return ϒpayload
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_NETRC_MACHINE":    VKBaseIE__NETRC_MACHINE,
				"_download_payload": VKBaseIE__download_payload,
				"_login":            VKBaseIE__login,
				"_real_initialize":  VKBaseIE__real_initialize,
			})
		}())
		VKIE = λ.Cal(λ.TypeType, λ.StrLiteral("VKIE"), λ.NewTuple(VKBaseIE), func() λ.Dict {
			var (
				VKIE_IE_NAME       λ.Object
				VKIE__VALID_URL    λ.Object
				VKIE__real_extract λ.Object
			)
			VKIE_IE_NAME = λ.StrLiteral("vk")
			VKIE__VALID_URL = λ.StrLiteral("(?x)\n                    https?://\n                        (?:\n                            (?:\n                                (?:(?:m|new)\\.)?vk\\.com/video_|\n                                (?:www\\.)?daxab.com/\n                            )\n                            ext\\.php\\?(?P<embed_query>.*?\\boid=(?P<oid>-?\\d+).*?\\bid=(?P<id>\\d+).*)|\n                            (?:\n                                (?:(?:m|new)\\.)?vk\\.com/(?:.+?\\?.*?z=)?video|\n                                (?:www\\.)?daxab.com/embed/\n                            )\n                            (?P<videoid>-?\\d+_\\d+)(?:.*\\blist=(?P<list_id>[\\da-f]+))?\n                        )\n                    ")
			VKIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ERRORS             λ.Object
						ERROR_COPYRIGHT    λ.Object
						ϒdailymotion_urls  λ.Object
						ϒdata              λ.Object
						ϒerror_message     λ.Object
						ϒerror_msg         λ.Object
						ϒerror_re          λ.Object
						ϒformat_id         λ.Object
						ϒformat_url        λ.Object
						ϒformats           λ.Object
						ϒheight            λ.Object
						ϒinfo_page         λ.Object
						ϒis_live           λ.Object
						ϒlist_id           λ.Object
						ϒm_opts            λ.Object
						ϒm_opts_url        λ.Object
						ϒm_rutube          λ.Object
						ϒmobj              λ.Object
						ϒmv_data           λ.Object
						ϒodnoklassniki_url λ.Object
						ϒopts              λ.Object
						ϒopts_url          λ.Object
						ϒpayload           λ.Object
						ϒpladform_url      λ.Object
						ϒplayer            λ.Object
						ϒrutube_url        λ.Object
						ϒself              = λargs[0]
						ϒtimestamp         λ.Object
						ϒtitle             λ.Object
						ϒurl               = λargs[1]
						ϒvideo_id          λ.Object
						ϒview_count        λ.Object
						ϒvimeo_url         λ.Object
						ϒyoutube_url       λ.Object
						τmp0               λ.Object
						τmp1               λ.Object
						τmp2               λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("videoid"))
					ϒmv_data = λ.DictLiteral(map[λ.Object]λ.Object{})
					if λ.IsTrue(ϒvideo_id) {
						ϒdata = λ.DictLiteral(map[string]λ.Object{
							"act":   λ.StrLiteral("show_inline"),
							"video": ϒvideo_id,
						})
						ϒlist_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("list_id"))
						if λ.IsTrue(ϒlist_id) {
							λ.SetItem(ϒdata, λ.StrLiteral("list"), ϒlist_id)
						}
						ϒpayload = λ.Calm(ϒself, "_download_payload", λ.StrLiteral("al_video"), ϒvideo_id, ϒdata)
						ϒinfo_page = λ.GetItem(ϒpayload, λ.IntLiteral(1))
						ϒopts = λ.GetItem(ϒpayload, λ.Neg(λ.IntLiteral(1)))
						ϒmv_data = func() λ.Object {
							if λv := λ.Calm(ϒopts, "get", λ.StrLiteral("mvData")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.DictLiteral(map[λ.Object]λ.Object{})
							}
						}()
						ϒplayer = func() λ.Object {
							if λv := λ.Calm(ϒopts, "get", λ.StrLiteral("player")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.DictLiteral(map[λ.Object]λ.Object{})
							}
						}()
					} else {
						ϒvideo_id = λ.Mod(λ.StrLiteral("%s_%s"), λ.NewTuple(
							λ.Calm(ϒmobj, "group", λ.StrLiteral("oid")),
							λ.Calm(ϒmobj, "group", λ.StrLiteral("id")),
						))
						ϒinfo_page = λ.Calm(ϒself, "_download_webpage", λ.Add(λ.StrLiteral("http://vk.com/video_ext.php?"), λ.Calm(ϒmobj, "group", λ.StrLiteral("embed_query"))), ϒvideo_id)
						ϒerror_message = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
							λ.NewList(
								λ.StrLiteral("(?s)<!><div[^>]+class=\"video_layer_message\"[^>]*>(.+?)</div>"),
								λ.StrLiteral("(?s)<div[^>]+id=\"video_ext_msg\"[^>]*>(.+?)</div>"),
							),
							ϒinfo_page,
							λ.StrLiteral("error message"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
						if λ.IsTrue(ϒerror_message) {
							panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(ϒerror_message), λ.KWArgs{
								{Name: "expected", Value: λ.True},
							})))
						}
						if λ.IsTrue(λ.Cal(Ωre.ϒsearch, λ.StrLiteral("<!>/login\\.php\\?.*\\bact=security_check"), ϒinfo_page)) {
							panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.StrLiteral("You are trying to log in from an unusual location. You should confirm ownership at vk.com to log in with this IP.")), λ.KWArgs{
								{Name: "expected", Value: λ.True},
							})))
						}
						ERROR_COPYRIGHT = λ.StrLiteral("Video %s has been removed from public access due to rightholder complaint.")
						ERRORS = λ.DictLiteral(map[string]λ.Object{
							">Видеозапись .*? была изъята из публичного доступа в связи с обращением правообладателя.<": ERROR_COPYRIGHT,
							">The video .*? was removed from public access by request of the copyright holder.<":        ERROR_COPYRIGHT,
							"<!>Please log in or <": λ.StrLiteral("Video %s is only available for registered users, use --username and --password options to provide account credentials."),
							"<!>Unknown error":      λ.StrLiteral("Video %s does not exist."),
							"<!>Видео временно недоступно": λ.StrLiteral("Video %s is temporarily unavailable."),
							"<!>Access denied": λ.StrLiteral("Access denied to video %s."),
							"<!>Видеозапись недоступна, так как её автор был заблокирован.":              λ.StrLiteral("Video %s is no longer available, because its author has been blocked."),
							"<!>This video is no longer available, because its author has been blocked.": λ.StrLiteral("Video %s is no longer available, because its author has been blocked."),
							"<!>This video is no longer available, because it has been deleted.":         λ.StrLiteral("Video %s is no longer available, because it has been deleted."),
							"<!>The video .+? is not available in your region.":                          λ.StrLiteral("Video %s is not available in your region."),
						})
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ERRORS, "items"))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							τmp2 = τmp1
							ϒerror_re = λ.GetItem(τmp2, λ.IntLiteral(0))
							ϒerror_msg = λ.GetItem(τmp2, λ.IntLiteral(1))
							if λ.IsTrue(λ.Cal(Ωre.ϒsearch, ϒerror_re, ϒinfo_page)) {
								panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(ϒerror_msg, ϒvideo_id)), λ.KWArgs{
									{Name: "expected", Value: λ.True},
								})))
							}
						}
						ϒplayer = λ.Calm(ϒself, "_parse_json", λ.Calm(ϒself, "_search_regex", λ.StrLiteral("var\\s+playerParams\\s*=\\s*({.+?})\\s*;\\s*\\n"), ϒinfo_page, λ.StrLiteral("player params")), ϒvideo_id)
					}
					ϒyoutube_url = λ.Calm(YoutubeIE, "_extract_url", ϒinfo_page)
					if λ.IsTrue(ϒyoutube_url) {
						return λ.Calm(ϒself, "url_result", ϒyoutube_url, λ.Calm(YoutubeIE, "ie_key"))
					}
					ϒvimeo_url = λ.Calm(VimeoIE, "_extract_url", ϒurl, ϒinfo_page)
					if ϒvimeo_url != λ.None {
						return λ.Calm(ϒself, "url_result", ϒvimeo_url, λ.Calm(VimeoIE, "ie_key"))
					}
					ϒpladform_url = λ.Calm(PladformIE, "_extract_url", ϒinfo_page)
					if λ.IsTrue(ϒpladform_url) {
						return λ.Calm(ϒself, "url_result", ϒpladform_url, λ.Calm(PladformIE, "ie_key"))
					}
					ϒm_rutube = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("\\ssrc=\"((?:https?:)?//rutube\\.ru\\\\?/(?:video|play)\\\\?/embed(?:.*?))\\\\?\""), ϒinfo_page)
					if ϒm_rutube != λ.None {
						ϒrutube_url = λ.Calm(ϒself, "_proto_relative_url", λ.Calm(λ.Calm(ϒm_rutube, "group", λ.IntLiteral(1)), "replace", λ.StrLiteral("\\"), λ.StrLiteral("")))
						return λ.Calm(ϒself, "url_result", ϒrutube_url)
					}
					ϒdailymotion_urls = λ.Calm(DailymotionIE, "_extract_urls", ϒinfo_page)
					if λ.IsTrue(ϒdailymotion_urls) {
						return λ.Calm(ϒself, "url_result", λ.GetItem(ϒdailymotion_urls, λ.IntLiteral(0)), λ.Calm(DailymotionIE, "ie_key"))
					}
					ϒodnoklassniki_url = λ.Calm(OdnoklassnikiIE, "_extract_url", ϒinfo_page)
					if λ.IsTrue(ϒodnoklassniki_url) {
						return λ.Calm(ϒself, "url_result", ϒodnoklassniki_url, λ.Calm(OdnoklassnikiIE, "ie_key"))
					}
					ϒm_opts = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("(?s)var\\s+opts\\s*=\\s*({.+?});"), ϒinfo_page)
					if λ.IsTrue(ϒm_opts) {
						ϒm_opts_url = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("url\\s*:\\s*'((?!/\\b)[^']+)"), λ.Calm(ϒm_opts, "group", λ.IntLiteral(1)))
						if λ.IsTrue(ϒm_opts_url) {
							ϒopts_url = λ.Calm(ϒm_opts_url, "group", λ.IntLiteral(1))
							if λ.IsTrue(λ.Calm(ϒopts_url, "startswith", λ.StrLiteral("//"))) {
								ϒopts_url = λ.Add(λ.StrLiteral("http:"), ϒopts_url)
							}
							return λ.Calm(ϒself, "url_result", ϒopts_url)
						}
					}
					ϒdata = λ.GetItem(λ.GetItem(ϒplayer, λ.StrLiteral("params")), λ.IntLiteral(0))
					ϒtitle = λ.Cal(ϒunescapeHTML, λ.GetItem(ϒdata, λ.StrLiteral("md_title")))
					ϒis_live = λ.Eq(λ.Calm(ϒdata, "get", λ.StrLiteral("live")), λ.IntLiteral(2))
					if λ.IsTrue(ϒis_live) {
						ϒtitle = λ.Calm(ϒself, "_live_title", ϒtitle)
					}
					ϒtimestamp = func() λ.Object {
						if λv := λ.Cal(ϒunified_timestamp, λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
							λ.StrLiteral("class=[\"\\']mv_info_date[^>]+>([^<]+)(?:<|from)"),
							ϒinfo_page,
							λ.StrLiteral("upload date"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(ϒint_or_none, λ.Calm(ϒdata, "get", λ.StrLiteral("date")))
						}
					}()
					ϒview_count = λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("class=[\"\\']mv_views_count[^>]+>\\s*([\\d,.]+)"),
						ϒinfo_page,
						λ.StrLiteral("view count"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					}))
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒdata, "items"))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒformat_id = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒformat_url = λ.GetItem(τmp2, λ.IntLiteral(1))
						ϒformat_url = λ.Cal(ϒurl_or_none, ϒformat_url)
						if λ.IsTrue(func() λ.Object {
							if λv := λ.NewBool(!λ.IsTrue(ϒformat_url)); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewBool(!λ.IsTrue(λ.Calm(ϒformat_url, "startswith", λ.NewTuple(
									λ.StrLiteral("http"),
									λ.StrLiteral("//"),
									λ.StrLiteral("rtmp"),
								))))
							}
						}()) {
							continue
						}
						if λ.IsTrue(func() λ.Object {
							if λv := λ.Calm(ϒformat_id, "startswith", λ.NewTuple(
								λ.StrLiteral("url"),
								λ.StrLiteral("cache"),
							)); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewBool(λ.Contains(λ.NewTuple(
									λ.StrLiteral("extra_data"),
									λ.StrLiteral("live_mp4"),
									λ.StrLiteral("postlive_mp4"),
								), ϒformat_id))
							}
						}()) {
							ϒheight = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("^(?:url|cache)(\\d+)"),
								ϒformat_id,
								λ.StrLiteral("height"),
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							}))
							λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
								"format_id": ϒformat_id,
								"url":       ϒformat_url,
								"height":    ϒheight,
							}))
						} else {
							if λ.IsTrue(λ.Eq(ϒformat_id, λ.StrLiteral("hls"))) {
								λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
									ϒformat_url,
									ϒvideo_id,
									λ.StrLiteral("mp4"),
									λ.StrLiteral("m3u8_native"),
								), λ.KWArgs{
									{Name: "m3u8_id", Value: ϒformat_id},
									{Name: "fatal", Value: λ.False},
									{Name: "live", Value: ϒis_live},
								}))
							} else {
								if λ.IsTrue(λ.Eq(ϒformat_id, λ.StrLiteral("rtmp"))) {
									λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
										"format_id": ϒformat_id,
										"url":       ϒformat_url,
										"ext":       λ.StrLiteral("flv"),
									}))
								}
							}
						}
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					return λ.DictLiteral(map[string]λ.Object{
						"id":        ϒvideo_id,
						"formats":   ϒformats,
						"title":     ϒtitle,
						"thumbnail": λ.Calm(ϒdata, "get", λ.StrLiteral("jpg")),
						"uploader":  λ.Calm(ϒdata, "get", λ.StrLiteral("md_author")),
						"uploader_id": λ.Cal(ϒstr_or_none, func() λ.Object {
							if λv := λ.Calm(ϒdata, "get", λ.StrLiteral("author_id")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(ϒmv_data, "get", λ.StrLiteral("authorId"))
							}
						}()),
						"duration": λ.Cal(ϒint_or_none, func() λ.Object {
							if λv := λ.Calm(ϒdata, "get", λ.StrLiteral("duration")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(ϒmv_data, "get", λ.StrLiteral("duration"))
							}
						}()),
						"timestamp":     ϒtimestamp,
						"view_count":    ϒview_count,
						"like_count":    λ.Cal(ϒint_or_none, λ.Calm(ϒmv_data, "get", λ.StrLiteral("likes"))),
						"comment_count": λ.Cal(ϒint_or_none, λ.Calm(ϒmv_data, "get", λ.StrLiteral("commcount"))),
						"is_live":       ϒis_live,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":       VKIE_IE_NAME,
				"_VALID_URL":    VKIE__VALID_URL,
				"_real_extract": VKIE__real_extract,
			})
		}())
		VKUserVideosIE = λ.Cal(λ.TypeType, λ.StrLiteral("VKUserVideosIE"), λ.NewTuple(VKBaseIE), func() λ.Dict {
			var (
				VKUserVideosIE__VALID_URL λ.Object
			)
			VKUserVideosIE__VALID_URL = λ.StrLiteral("https?://(?:(?:m|new)\\.)?vk\\.com/videos(?P<id>-?[0-9]+)(?!\\?.*\\bz=video)(?:[/?#&](?:.*?\\bsection=(?P<section>\\w+))?|$)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": VKUserVideosIE__VALID_URL,
			})
		}())
		VKWallPostIE = λ.Cal(λ.TypeType, λ.StrLiteral("VKWallPostIE"), λ.NewTuple(VKBaseIE), func() λ.Dict {
			var (
				VKWallPostIE__VALID_URL λ.Object
			)
			VKWallPostIE__VALID_URL = λ.StrLiteral("https?://(?:(?:(?:(?:m|new)\\.)?vk\\.com/(?:[^?]+\\?.*\\bw=)?wall(?P<id>-?\\d+_\\d+)))")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": VKWallPostIE__VALID_URL,
			})
		}())
	})
}
