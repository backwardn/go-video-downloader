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
 * gdcvault/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/gdcvault.py
 */

package gdcvault

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωkaltura "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/kaltura"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	GDCVaultIE          λ.Object
	HEADRequest         λ.Object
	InfoExtractor       λ.Object
	KalturaIE           λ.Object
	ϒsanitized_Request  λ.Object
	ϒsmuggle_url        λ.Object
	ϒurlencode_postdata λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		KalturaIE = Ωkaltura.KalturaIE
		HEADRequest = Ωutils.HEADRequest
		ϒsanitized_Request = Ωutils.ϒsanitized_Request
		ϒsmuggle_url = Ωutils.ϒsmuggle_url
		ϒurlencode_postdata = Ωutils.ϒurlencode_postdata
		GDCVaultIE = λ.Cal(λ.TypeType, λ.StrLiteral("GDCVaultIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				GDCVaultIE__NETRC_MACHINE λ.Object
				GDCVaultIE__VALID_URL     λ.Object
				GDCVaultIE__login         λ.Object
				GDCVaultIE__real_extract  λ.Object
			)
			GDCVaultIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?gdcvault\\.com/play/(?P<id>\\d+)(?:/(?P<name>[\\w-]+))?")
			GDCVaultIE__NETRC_MACHINE = λ.StrLiteral("gdcvault")
			GDCVaultIE__login = λ.NewFunction("_login",
				[]λ.Param{
					{Name: "self"},
					{Name: "webpage_url"},
					{Name: "display_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id  = λargs[2]
						ϒlogin_form  λ.Object
						ϒlogin_url   λ.Object
						ϒlogout_url  λ.Object
						ϒmobj        λ.Object
						ϒpassword    λ.Object
						ϒrequest     λ.Object
						ϒself        = λargs[0]
						ϒstart_page  λ.Object
						ϒusername    λ.Object
						ϒwebpage_url = λargs[1]
						τmp0         λ.Object
					)
					τmp0 = λ.Calm(ϒself, "_get_login_info")
					ϒusername = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒpassword = λ.GetItem(τmp0, λ.IntLiteral(1))
					if λ.IsTrue(func() λ.Object {
						if λv := λ.NewBool(ϒusername == λ.None); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewBool(ϒpassword == λ.None)
						}
					}()) {
						λ.Calm(ϒself, "report_warning", λ.Add(λ.Add(λ.StrLiteral("It looks like "), ϒwebpage_url), λ.StrLiteral(" requires a login. Try specifying a username and password and try again.")))
						return λ.None
					}
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.StrLiteral("(?P<root_url>https?://.*?/).*"), ϒwebpage_url)
					ϒlogin_url = λ.Add(λ.Calm(ϒmobj, "group", λ.StrLiteral("root_url")), λ.StrLiteral("api/login.php"))
					ϒlogout_url = λ.Add(λ.Calm(ϒmobj, "group", λ.StrLiteral("root_url")), λ.StrLiteral("logout"))
					ϒlogin_form = λ.DictLiteral(map[string]λ.Object{
						"email":    ϒusername,
						"password": ϒpassword,
					})
					ϒrequest = λ.Cal(ϒsanitized_Request, ϒlogin_url, λ.Cal(ϒurlencode_postdata, ϒlogin_form))
					λ.Calm(ϒrequest, "add_header", λ.StrLiteral("Content-Type"), λ.StrLiteral("application/x-www-form-urlencoded"))
					λ.Calm(ϒself, "_download_webpage", ϒrequest, ϒdisplay_id, λ.StrLiteral("Logging in"))
					ϒstart_page = λ.Calm(ϒself, "_download_webpage", ϒwebpage_url, ϒdisplay_id, λ.StrLiteral("Getting authenticated video page"))
					λ.Calm(ϒself, "_download_webpage", ϒlogout_url, ϒdisplay_id, λ.StrLiteral("Logging out"))
					return ϒstart_page
				})
			GDCVaultIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						PLAYER_REGEX λ.Object
						ϒdirect_url  λ.Object
						ϒdisplay_id  λ.Object
						ϒembed_url   λ.Object
						ϒie_key      λ.Object
						ϒlogin_res   λ.Object
						ϒname        λ.Object
						ϒself        = λargs[0]
						ϒstart_page  λ.Object
						ϒtitle       λ.Object
						ϒurl         = λargs[1]
						ϒvideo_id    λ.Object
						ϒvideo_url   λ.Object
						ϒwebpage_url λ.Object
						ϒxml_name    λ.Object
						ϒxml_root    λ.Object
						τmp0         λ.Object
					)
					τmp0 = λ.Calm(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups")
					ϒvideo_id = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒname = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒdisplay_id = func() λ.Object {
						if λv := ϒname; λ.IsTrue(λv) {
							return λv
						} else {
							return ϒvideo_id
						}
					}()
					ϒwebpage_url = λ.Add(λ.StrLiteral("http://www.gdcvault.com/play/"), ϒvideo_id)
					ϒstart_page = λ.Calm(ϒself, "_download_webpage", ϒwebpage_url, ϒdisplay_id)
					ϒdirect_url = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("s1\\.addVariable\\(\"file\",\\s*encodeURIComponent\\(\"(/[^\"]+)\"\\)\\);"),
						ϒstart_page,
						λ.StrLiteral("url"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒdirect_url) {
						ϒtitle = λ.Calm(ϒself, "_html_search_regex", λ.StrLiteral("<td><strong>Session Name:?</strong></td>\\s*<td>(.*?)</td>"), ϒstart_page, λ.StrLiteral("title"))
						ϒvideo_url = λ.Add(λ.StrLiteral("http://www.gdcvault.com"), ϒdirect_url)
						ϒvideo_url = λ.Calm(λ.Calm(ϒself, "_request_webpage", λ.Cal(HEADRequest, ϒvideo_url), ϒvideo_id), "geturl")
						return λ.DictLiteral(map[string]λ.Object{
							"id":         ϒvideo_id,
							"display_id": ϒdisplay_id,
							"url":        ϒvideo_url,
							"title":      ϒtitle,
						})
					}
					ϒembed_url = λ.Calm(KalturaIE, "_extract_url", ϒstart_page)
					if λ.IsTrue(ϒembed_url) {
						ϒembed_url = λ.Cal(ϒsmuggle_url, ϒembed_url, λ.DictLiteral(map[string]λ.Object{
							"source_url": ϒurl,
						}))
						ϒie_key = λ.StrLiteral("Kaltura")
					} else {
						PLAYER_REGEX = λ.StrLiteral("<iframe src=\"(?P<xml_root>.+?)/(?:gdc-)?player.*?\\.html.*?\".*?</iframe>")
						ϒxml_root = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
							PLAYER_REGEX,
							ϒstart_page,
							λ.StrLiteral("xml root"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
						if ϒxml_root == λ.None {
							ϒlogin_res = λ.Calm(ϒself, "_login", ϒwebpage_url, ϒdisplay_id)
							if ϒlogin_res == λ.None {
								λ.Calm(ϒself, "report_warning", λ.StrLiteral("Could not login."))
							} else {
								ϒstart_page = ϒlogin_res
								ϒxml_root = λ.Calm(ϒself, "_html_search_regex", PLAYER_REGEX, ϒstart_page, λ.StrLiteral("xml root"))
							}
						}
						ϒxml_name = λ.Calm(ϒself, "_html_search_regex", λ.StrLiteral("<iframe src=\".*?\\?xml(?:=|URL=xml/)(.+?\\.xml).*?\".*?</iframe>"), ϒstart_page, λ.StrLiteral("xml filename"))
						ϒembed_url = λ.Mod(λ.StrLiteral("%s/xml/%s"), λ.NewTuple(
							ϒxml_root,
							ϒxml_name,
						))
						ϒie_key = λ.StrLiteral("DigitallySpeaking")
					}
					return λ.DictLiteral(map[string]λ.Object{
						"_type":      λ.StrLiteral("url_transparent"),
						"id":         ϒvideo_id,
						"display_id": ϒdisplay_id,
						"url":        ϒembed_url,
						"ie_key":     ϒie_key,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_NETRC_MACHINE": GDCVaultIE__NETRC_MACHINE,
				"_VALID_URL":     GDCVaultIE__VALID_URL,
				"_login":         GDCVaultIE__login,
				"_real_extract":  GDCVaultIE__real_extract,
			})
		}())
	})
}
