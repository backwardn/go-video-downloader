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
 * theplatform/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/theplatform.py
 */

package theplatform

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωadobepass "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/adobepass"
	Ωonce "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/once"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	AdobePassIE                   λ.Object
	ExtractorError                λ.Object
	OnceIE                        λ.Object
	ThePlatformBaseIE             λ.Object
	ThePlatformFeedIE             λ.Object
	ThePlatformIE                 λ.Object
	ϒ_x                           λ.Object
	ϒcompat_parse_qs              λ.Object
	ϒcompat_urllib_parse_urlparse λ.Object
	ϒdefault_ns                   λ.Object
	ϒdetermine_ext                λ.Object
	ϒfind_xpath_attr              λ.Object
	ϒfloat_or_none                λ.Object
	ϒint_or_none                  λ.Object
	ϒmimetype2ext                 λ.Object
	ϒsanitized_Request            λ.Object
	ϒunsmuggle_url                λ.Object
	ϒupdate_url_query             λ.Object
	ϒxpath_with_ns                λ.Object
)

func init() {
	λ.InitModule(func() {
		OnceIE = Ωonce.OnceIE
		AdobePassIE = Ωadobepass.AdobePassIE
		ϒcompat_parse_qs = Ωcompat.ϒcompat_parse_qs
		ϒcompat_urllib_parse_urlparse = Ωcompat.ϒcompat_urllib_parse_urlparse
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒsanitized_Request = Ωutils.ϒsanitized_Request
		ϒunsmuggle_url = Ωutils.ϒunsmuggle_url
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		ϒxpath_with_ns = Ωutils.ϒxpath_with_ns
		ϒmimetype2ext = Ωutils.ϒmimetype2ext
		ϒfind_xpath_attr = Ωutils.ϒfind_xpath_attr
		ϒdefault_ns = λ.StrLiteral("http://www.w3.org/2005/SMIL21/Language")
		ϒ_x = λ.NewFunction("<lambda>",
			[]λ.Param{
				{Name: "p"},
			},
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				var (
					ϒp = λargs[0]
				)
				return λ.Cal(ϒxpath_with_ns, ϒp, λ.DictLiteral(map[string]λ.Object{
					"smil": ϒdefault_ns,
				}))
			})
		ThePlatformBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("ThePlatformBaseIE"), λ.NewTuple(OnceIE), func() λ.Dict {
			var (
				ThePlatformBaseIE__TP_TLD                        λ.Object
				ThePlatformBaseIE__download_theplatform_metadata λ.Object
				ThePlatformBaseIE__extract_theplatform_metadata  λ.Object
				ThePlatformBaseIE__extract_theplatform_smil      λ.Object
				ThePlatformBaseIE__parse_theplatform_metadata    λ.Object
			)
			ThePlatformBaseIE__TP_TLD = λ.StrLiteral("com")
			ThePlatformBaseIE__extract_theplatform_smil = λ.NewFunction("_extract_theplatform_smil",
				[]λ.Param{
					{Name: "self"},
					{Name: "smil_url"},
					{Name: "video_id"},
					{Name: "note", Def: λ.StrLiteral("Downloading SMIL data")},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒ_format       λ.Object
						ϒerror_element λ.Object
						ϒexception     λ.Object
						ϒformats       λ.Object
						ϒhdnea2        λ.Object
						ϒmedia_url     λ.Object
						ϒmeta          λ.Object
						ϒnote          = λargs[3]
						ϒself          = λargs[0]
						ϒsmil_formats  λ.Object
						ϒsmil_url      = λargs[1]
						ϒsubtitles     λ.Object
						ϒvideo_id      = λargs[2]
						τmp0           λ.Object
						τmp1           λ.Object
					)
					ϒmeta = λ.Call(λ.GetAttr(ϒself, "_download_xml", nil), λ.NewArgs(
						ϒsmil_url,
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "note", Value: ϒnote},
						{Name: "query", Value: λ.DictLiteral(map[string]string{
							"format": "SMIL",
						})},
						{Name: "headers", Value: λ.Calm(ϒself, "geo_verification_headers")},
					})
					ϒerror_element = λ.Cal(ϒfind_xpath_attr, ϒmeta, λ.Cal(ϒ_x, λ.StrLiteral(".//smil:ref")), λ.StrLiteral("src"))
					if ϒerror_element != λ.None {
						ϒexception = λ.Cal(ϒfind_xpath_attr, ϒerror_element, λ.Cal(ϒ_x, λ.StrLiteral(".//smil:param")), λ.StrLiteral("name"), λ.StrLiteral("exception"))
						if ϒexception != λ.None {
							if λ.IsTrue(λ.Eq(λ.Calm(ϒexception, "get", λ.StrLiteral("value")), λ.StrLiteral("GeoLocationBlocked"))) {
								λ.Calm(ϒself, "raise_geo_restricted", λ.GetItem(λ.GetAttr(ϒerror_element, "attrib", nil), λ.StrLiteral("abstract")))
							} else {
								if λ.IsTrue(λ.Calm(λ.GetItem(λ.GetAttr(ϒerror_element, "attrib", nil), λ.StrLiteral("src")), "startswith", λ.Mod(λ.StrLiteral("http://link.theplatform.%s/s/errorFiles/Unavailable."), λ.GetAttr(ϒself, "_TP_TLD", nil)))) {
									panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.GetItem(λ.GetAttr(ϒerror_element, "attrib", nil), λ.StrLiteral("abstract"))), λ.KWArgs{
										{Name: "expected", Value: λ.True},
									})))
								}
							}
						}
					}
					ϒsmil_formats = λ.Call(λ.GetAttr(ϒself, "_parse_smil_formats", nil), λ.NewArgs(
						ϒmeta,
						ϒsmil_url,
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "namespace", Value: ϒdefault_ns},
						{Name: "f4m_params", Value: λ.DictLiteral(map[string]string{
							"g":      "UXWGVKRWHFSP",
							"hdcore": "3.0.3",
						})},
						{Name: "transform_rtmp_url", Value: λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "streamer"},
								{Name: "src"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒsrc      = λargs[1]
									ϒstreamer = λargs[0]
								)
								return λ.NewTuple(
									ϒstreamer,
									λ.Add(λ.StrLiteral("mp4:"), ϒsrc),
								)
							})},
					})
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, ϒsmil_formats)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒ_format = τmp1
						if λ.IsTrue(λ.Calm(OnceIE, "suitable", λ.GetItem(ϒ_format, λ.StrLiteral("url")))) {
							λ.Calm(ϒformats, "extend", λ.Calm(ϒself, "_extract_once_formats", λ.GetItem(ϒ_format, λ.StrLiteral("url"))))
						} else {
							ϒmedia_url = λ.GetItem(ϒ_format, λ.StrLiteral("url"))
							if λ.IsTrue(λ.Eq(λ.Cal(ϒdetermine_ext, ϒmedia_url), λ.StrLiteral("m3u8"))) {
								ϒhdnea2 = λ.Calm(λ.Calm(ϒself, "_get_cookies", ϒmedia_url), "get", λ.StrLiteral("hdnea2"))
								if λ.IsTrue(ϒhdnea2) {
									λ.SetItem(ϒ_format, λ.StrLiteral("url"), λ.Cal(ϒupdate_url_query, ϒmedia_url, λ.DictLiteral(map[string]λ.Object{
										"hdnea3": λ.GetAttr(ϒhdnea2, "value", nil),
									})))
								}
							}
							λ.Calm(ϒformats, "append", ϒ_format)
						}
					}
					ϒsubtitles = λ.Calm(ϒself, "_parse_smil_subtitles", ϒmeta, ϒdefault_ns)
					return λ.NewTuple(
						ϒformats,
						ϒsubtitles,
					)
				})
			ThePlatformBaseIE__download_theplatform_metadata = λ.NewFunction("_download_theplatform_metadata",
				[]λ.Param{
					{Name: "self"},
					{Name: "path"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒinfo_url λ.Object
						ϒpath     = λargs[1]
						ϒself     = λargs[0]
						ϒvideo_id = λargs[2]
					)
					ϒinfo_url = λ.Mod(λ.StrLiteral("http://link.theplatform.%s/s/%s?format=preview"), λ.NewTuple(
						λ.GetAttr(ϒself, "_TP_TLD", nil),
						ϒpath,
					))
					return λ.Calm(ϒself, "_download_json", ϒinfo_url, ϒvideo_id)
				})
			ThePlatformBaseIE__parse_theplatform_metadata = λ.NewFunction("_parse_theplatform_metadata",
				[]λ.Param{
					{Name: "self"},
					{Name: "info"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒ_add_chapter λ.Object
						ϒcaption      λ.Object
						ϒcaptions     λ.Object
						ϒchapter      λ.Object
						ϒchapters     λ.Object
						ϒduration     λ.Object
						ϒinfo         = λargs[1]
						ϒlang         λ.Object
						ϒmime         λ.Object
						ϒself         = λargs[0]
						ϒsrc          λ.Object
						ϒsubtitles    λ.Object
						ϒtp_chapters  λ.Object
						τmp0          λ.Object
						τmp1          λ.Object
						τmp2          λ.Object
					)
					_ = ϒself
					ϒsubtitles = λ.DictLiteral(map[λ.Object]λ.Object{})
					ϒcaptions = λ.Calm(ϒinfo, "get", λ.StrLiteral("captions"))
					if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒcaptions, λ.ListType)) {
						τmp0 = λ.Cal(λ.BuiltinIter, ϒcaptions)
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒcaption = τmp1
							τmp2 = λ.NewTuple(
								λ.Calm(ϒcaption, "get", λ.StrLiteral("lang"), λ.StrLiteral("en")),
								λ.Calm(ϒcaption, "get", λ.StrLiteral("src")),
								λ.Calm(ϒcaption, "get", λ.StrLiteral("type")),
							)
							ϒlang = λ.GetItem(τmp2, λ.IntLiteral(0))
							ϒsrc = λ.GetItem(τmp2, λ.IntLiteral(1))
							ϒmime = λ.GetItem(τmp2, λ.IntLiteral(2))
							λ.Calm(λ.Calm(ϒsubtitles, "setdefault", ϒlang, λ.NewList()), "append", λ.DictLiteral(map[string]λ.Object{
								"ext": λ.Cal(ϒmimetype2ext, ϒmime),
								"url": ϒsrc,
							}))
						}
					}
					ϒduration = λ.Calm(ϒinfo, "get", λ.StrLiteral("duration"))
					ϒtp_chapters = λ.Calm(ϒinfo, "get", λ.StrLiteral("chapters"), λ.NewList())
					ϒchapters = λ.NewList()
					if λ.IsTrue(ϒtp_chapters) {
						ϒ_add_chapter = λ.NewFunction("_add_chapter",
							[]λ.Param{
								{Name: "start_time"},
								{Name: "end_time"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒend_time   = λargs[1]
									ϒstart_time = λargs[0]
								)
								ϒstart_time = λ.Cal(ϒfloat_or_none, ϒstart_time, λ.IntLiteral(1000))
								ϒend_time = λ.Cal(ϒfloat_or_none, ϒend_time, λ.IntLiteral(1000))
								if λ.IsTrue(func() λ.Object {
									if λv := λ.NewBool(ϒstart_time == λ.None); λ.IsTrue(λv) {
										return λv
									} else {
										return λ.NewBool(ϒend_time == λ.None)
									}
								}()) {
									return λ.None
								}
								λ.Calm(ϒchapters, "append", λ.DictLiteral(map[string]λ.Object{
									"start_time": ϒstart_time,
									"end_time":   ϒend_time,
								}))
								return λ.None
							})
						τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(ϒtp_chapters, λ.NewSlice(λ.None, λ.Neg(λ.IntLiteral(1)), λ.None)))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒchapter = τmp1
							λ.Cal(ϒ_add_chapter, λ.Calm(ϒchapter, "get", λ.StrLiteral("startTime")), λ.Calm(ϒchapter, "get", λ.StrLiteral("endTime")))
						}
						λ.Cal(ϒ_add_chapter, λ.Calm(λ.GetItem(ϒtp_chapters, λ.Neg(λ.IntLiteral(1))), "get", λ.StrLiteral("startTime")), func() λ.Object {
							if λv := λ.Calm(λ.GetItem(ϒtp_chapters, λ.Neg(λ.IntLiteral(1))), "get", λ.StrLiteral("endTime")); λ.IsTrue(λv) {
								return λv
							} else {
								return ϒduration
							}
						}())
					}
					return λ.DictLiteral(map[string]λ.Object{
						"title":       λ.GetItem(ϒinfo, λ.StrLiteral("title")),
						"subtitles":   ϒsubtitles,
						"description": λ.GetItem(ϒinfo, λ.StrLiteral("description")),
						"thumbnail":   λ.GetItem(ϒinfo, λ.StrLiteral("defaultThumbnailUrl")),
						"duration":    λ.Cal(ϒfloat_or_none, ϒduration, λ.IntLiteral(1000)),
						"timestamp": func() λ.Object {
							if λv := λ.Cal(ϒint_or_none, λ.Calm(ϒinfo, "get", λ.StrLiteral("pubDate")), λ.IntLiteral(1000)); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.None
							}
						}(),
						"uploader": λ.Calm(ϒinfo, "get", λ.StrLiteral("billingCode")),
						"chapters": ϒchapters,
					})
				})
			ThePlatformBaseIE__extract_theplatform_metadata = λ.NewFunction("_extract_theplatform_metadata",
				[]λ.Param{
					{Name: "self"},
					{Name: "path"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒinfo     λ.Object
						ϒpath     = λargs[1]
						ϒself     = λargs[0]
						ϒvideo_id = λargs[2]
					)
					ϒinfo = λ.Calm(ϒself, "_download_theplatform_metadata", ϒpath, ϒvideo_id)
					return λ.Calm(ϒself, "_parse_theplatform_metadata", ϒinfo)
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_TP_TLD":                        ThePlatformBaseIE__TP_TLD,
				"_download_theplatform_metadata": ThePlatformBaseIE__download_theplatform_metadata,
				"_extract_theplatform_metadata":  ThePlatformBaseIE__extract_theplatform_metadata,
				"_extract_theplatform_smil":      ThePlatformBaseIE__extract_theplatform_smil,
				"_parse_theplatform_metadata":    ThePlatformBaseIE__parse_theplatform_metadata,
			})
		}())
		ThePlatformIE = λ.Cal(λ.TypeType, λ.StrLiteral("ThePlatformIE"), λ.NewTuple(
			ThePlatformBaseIE,
			AdobePassIE,
		), func() λ.Dict {
			var (
				ThePlatformIE__VALID_URL    λ.Object
				ThePlatformIE__real_extract λ.Object
			)
			ThePlatformIE__VALID_URL = λ.StrLiteral("(?x)\n        (?:https?://(?:link|player)\\.theplatform\\.com/[sp]/(?P<provider_id>[^/]+)/\n           (?:(?:(?:[^/]+/)+select/)?(?P<media>media/(?:guid/\\d+/)?)?|(?P<config>(?:[^/\\?]+/(?:swf|config)|onsite)/select/))?\n         |theplatform:)(?P<id>[^/\\?&]+)")
			ThePlatformIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcombined_subtitles λ.Object
						ϒconfig             λ.Object
						ϒconfig_url         λ.Object
						ϒfeed_id            λ.Object
						ϒfeed_script        λ.Object
						ϒformats            λ.Object
						ϒheaders            λ.Object
						ϒmobj               λ.Object
						ϒpath               λ.Object
						ϒprovider_id        λ.Object
						ϒqs_dict            λ.Object
						ϒrelease_url        λ.Object
						ϒrequest            λ.Object
						ϒret                λ.Object
						ϒscript             λ.Object
						ϒscripts            λ.Object
						ϒself               = λargs[0]
						ϒsig                λ.Object
						ϒsmil_url           λ.Object
						ϒsmuggled_data      λ.Object
						ϒsource_url         λ.Object
						ϒsubtitles          λ.Object
						ϒurl                = λargs[1]
						ϒvideo_id           λ.Object
						ϒwebpage            λ.Object
						τmp0                λ.Object
						τmp1                λ.Object
					)
					τmp0 = λ.Cal(ϒunsmuggle_url, ϒurl, λ.DictLiteral(map[λ.Object]λ.Object{}))
					ϒurl = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒsmuggled_data = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒprovider_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("provider_id"))
					ϒvideo_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("id"))
					if !λ.IsTrue(ϒprovider_id) {
						ϒprovider_id = λ.StrLiteral("dJ5BDC")
					}
					ϒpath = λ.Add(ϒprovider_id, λ.StrLiteral("/"))
					if λ.IsTrue(λ.Calm(ϒmobj, "group", λ.StrLiteral("media"))) {
						τmp0 = λ.IAdd(ϒpath, λ.Calm(ϒmobj, "group", λ.StrLiteral("media")))
						ϒpath = τmp0
					}
					τmp0 = λ.IAdd(ϒpath, ϒvideo_id)
					ϒpath = τmp0
					ϒqs_dict = λ.Cal(ϒcompat_parse_qs, λ.GetAttr(λ.Cal(ϒcompat_urllib_parse_urlparse, ϒurl), "query", nil))
					if λ.Contains(ϒqs_dict, λ.StrLiteral("guid")) {
						ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒvideo_id)
						ϒscripts = λ.Cal(Ωre.ϒfindall, λ.StrLiteral("<script[^>]+src=\"([^\"]+)\""), ϒwebpage)
						ϒfeed_id = λ.None
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.ReversedIteratorType, ϒscripts))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒscript = τmp1
							ϒfeed_script = λ.Calm(ϒself, "_download_webpage", λ.Calm(ϒself, "_proto_relative_url", ϒscript, λ.StrLiteral("http:")), ϒvideo_id, λ.StrLiteral("Downloading feed script"))
							ϒfeed_id = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("defaultFeedId\\s*:\\s*\"([^\"]+)\""),
								ϒfeed_script,
								λ.StrLiteral("default feed id"),
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							})
							if ϒfeed_id != λ.None {
								break
							}
						}
						if ϒfeed_id == λ.None {
							panic(λ.Raise(λ.Cal(ExtractorError, λ.StrLiteral("Unable to find feed id"))))
						}
						return λ.Calm(ϒself, "url_result", λ.Mod(λ.StrLiteral("http://feed.theplatform.com/f/%s/%s?byGuid=%s"), λ.NewTuple(
							ϒprovider_id,
							ϒfeed_id,
							λ.GetItem(λ.GetItem(ϒqs_dict, λ.StrLiteral("guid")), λ.IntLiteral(0)),
						)))
					}
					if λ.IsTrue(λ.Calm(ϒsmuggled_data, "get", λ.StrLiteral("force_smil_url"), λ.False)) {
						ϒsmil_url = ϒurl
					} else {
						if λ.Contains(ϒurl, λ.StrLiteral("/guid/")) {
							ϒheaders = λ.DictLiteral(map[λ.Object]λ.Object{})
							ϒsource_url = λ.Calm(ϒsmuggled_data, "get", λ.StrLiteral("source_url"))
							if λ.IsTrue(ϒsource_url) {
								λ.SetItem(ϒheaders, λ.StrLiteral("Referer"), ϒsource_url)
							}
							ϒrequest = λ.Call(ϒsanitized_Request, λ.NewArgs(ϒurl), λ.KWArgs{
								{Name: "headers", Value: ϒheaders},
							})
							ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒrequest, ϒvideo_id)
							ϒsmil_url = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("<link[^>]+href=([\"\\'])(?P<url>.+?)\\1[^>]+type=[\"\\']application/smil\\+xml"),
								ϒwebpage,
								λ.StrLiteral("smil url"),
							), λ.KWArgs{
								{Name: "group", Value: λ.StrLiteral("url")},
							})
							ϒpath = λ.Calm(ϒself, "_search_regex", λ.StrLiteral("link\\.theplatform\\.com/s/((?:[^/?#&]+/)+[^/?#&]+)"), ϒsmil_url, λ.StrLiteral("path"))
							τmp0 = λ.IAdd(ϒsmil_url, func() λ.Object {
								if !λ.Contains(ϒsmil_url, λ.StrLiteral("?")) {
									return λ.StrLiteral("?")
								} else {
									return λ.Add(λ.StrLiteral("&"), λ.StrLiteral("formats=m3u,mpeg4"))
								}
							}())
							ϒsmil_url = τmp0
						} else {
							if λ.IsTrue(λ.Calm(ϒmobj, "group", λ.StrLiteral("config"))) {
								ϒconfig_url = λ.Add(ϒurl, λ.StrLiteral("&form=json"))
								ϒconfig_url = λ.Calm(ϒconfig_url, "replace", λ.StrLiteral("swf/"), λ.StrLiteral("config/"))
								ϒconfig_url = λ.Calm(ϒconfig_url, "replace", λ.StrLiteral("onsite/"), λ.StrLiteral("onsite/config/"))
								ϒconfig = λ.Calm(ϒself, "_download_json", ϒconfig_url, ϒvideo_id, λ.StrLiteral("Downloading config"))
								if λ.Contains(ϒconfig, λ.StrLiteral("releaseUrl")) {
									ϒrelease_url = λ.GetItem(ϒconfig, λ.StrLiteral("releaseUrl"))
								} else {
									ϒrelease_url = λ.Mod(λ.StrLiteral("http://link.theplatform.com/s/%s?mbr=true"), ϒpath)
								}
								ϒsmil_url = λ.Add(ϒrelease_url, λ.StrLiteral("&formats=MPEG4&manifest=f4m"))
							} else {
								ϒsmil_url = λ.Mod(λ.StrLiteral("http://link.theplatform.com/s/%s?mbr=true"), ϒpath)
							}
						}
					}
					ϒsig = λ.Calm(ϒsmuggled_data, "get", λ.StrLiteral("sig"))
					if λ.IsTrue(ϒsig) {
						ϒsmil_url = λ.Calm(ϒself, "_sign_url", ϒsmil_url, λ.GetItem(ϒsig, λ.StrLiteral("key")), λ.GetItem(ϒsig, λ.StrLiteral("secret")))
					}
					τmp0 = λ.Calm(ϒself, "_extract_theplatform_smil", ϒsmil_url, ϒvideo_id)
					ϒformats = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒsubtitles = λ.GetItem(τmp0, λ.IntLiteral(1))
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒret = λ.Calm(ϒself, "_extract_theplatform_metadata", ϒpath, ϒvideo_id)
					ϒcombined_subtitles = λ.Calm(ϒself, "_merge_subtitles", λ.Calm(ϒret, "get", λ.StrLiteral("subtitles"), λ.DictLiteral(map[λ.Object]λ.Object{})), ϒsubtitles)
					λ.Calm(ϒret, "update", λ.DictLiteral(map[string]λ.Object{
						"id":        ϒvideo_id,
						"formats":   ϒformats,
						"subtitles": ϒcombined_subtitles,
					}))
					return ϒret
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    ThePlatformIE__VALID_URL,
				"_real_extract": ThePlatformIE__real_extract,
			})
		}())
		ThePlatformFeedIE = λ.Cal(λ.TypeType, λ.StrLiteral("ThePlatformFeedIE"), λ.NewTuple(ThePlatformBaseIE), func() λ.Dict {
			var (
				ThePlatformFeedIE__VALID_URL λ.Object
			)
			ThePlatformFeedIE__VALID_URL = λ.StrLiteral("https?://feed\\.theplatform\\.com/f/(?P<provider_id>[^/]+)/(?P<feed_id>[^?/]+)\\?(?:[^&]+&)*(?P<filter>by(?:Gui|I)d=(?P<id>[^&]+))")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": ThePlatformFeedIE__VALID_URL,
			})
		}())
	})
}
