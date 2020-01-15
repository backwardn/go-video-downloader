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
 * vice/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/vice.py
 */

package vice

import (
	Ωjson "github.com/tenta-browser/go-video-downloader/gen/json"
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωtime "github.com/tenta-browser/go-video-downloader/gen/time"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωadobepass "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/adobepass"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωyoutube "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/youtube"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	AdobePassIE      λ.Object
	ExtractorError   λ.Object
	InfoExtractor    λ.Object
	ViceArticleIE    λ.Object
	ViceBaseIE       λ.Object
	ViceIE           λ.Object
	ViceShowIE       λ.Object
	YoutubeIE        λ.Object
	ϒclean_html      λ.Object
	ϒcompat_str      λ.Object
	ϒint_or_none     λ.Object
	ϒparse_age_limit λ.Object
	ϒstr_or_none     λ.Object
	ϒtry_get         λ.Object
)

func init() {
	λ.InitModule(func() {
		AdobePassIE = Ωadobepass.AdobePassIE
		InfoExtractor = Ωcommon.InfoExtractor
		YoutubeIE = Ωyoutube.YoutubeIE
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒclean_html = Ωutils.ϒclean_html
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_age_limit = Ωutils.ϒparse_age_limit
		ϒstr_or_none = Ωutils.ϒstr_or_none
		ϒtry_get = Ωutils.ϒtry_get
		ViceBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("ViceBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ViceBaseIE__call_api λ.Object
			)
			ViceBaseIE__call_api = λ.NewFunction("_call_api",
				[]λ.Param{
					{Name: "self"},
					{Name: "resource"},
					{Name: "resource_key"},
					{Name: "resource_id"},
					{Name: "locale"},
					{Name: "fields"},
					{Name: "args", Def: λ.StrLiteral("")},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒargs         = λargs[6]
						ϒfields       = λargs[5]
						ϒlocale       = λargs[4]
						ϒresource     = λargs[1]
						ϒresource_id  = λargs[3]
						ϒresource_key = λargs[2]
						ϒself         = λargs[0]
					)
					return λ.GetItem(λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.StrLiteral("https://video.vice.com/api/v1/graphql"),
						ϒresource_id,
					), λ.KWArgs{
						{Name: "query", Value: λ.DictLiteral(map[string]λ.Object{
							"query": λ.Mod(λ.StrLiteral("{\n  %s(locale: \"%s\", %s: \"%s\"%s) {\n    %s\n  }\n}"), λ.NewTuple(
								ϒresource,
								ϒlocale,
								ϒresource_key,
								ϒresource_id,
								ϒargs,
								ϒfields,
							)),
						})},
					}), λ.StrLiteral("data")), ϒresource)
				})
			return λ.DictLiteral(map[string]λ.Object{
				"_call_api": ViceBaseIE__call_api,
			})
		}())
		ViceIE = λ.Cal(λ.TypeType, λ.StrLiteral("ViceIE"), λ.NewTuple(
			ViceBaseIE,
			AdobePassIE,
		), func() λ.Dict {
			var (
				ViceIE_IE_NAME       λ.Object
				ViceIE__VALID_URL    λ.Object
				ViceIE__extract_url  λ.Object
				ViceIE__extract_urls λ.Object
				ViceIE__real_extract λ.Object
			)
			ViceIE_IE_NAME = λ.StrLiteral("vice")
			ViceIE__VALID_URL = λ.StrLiteral("https?://(?:(?:video|vms)\\.vice|(?:www\\.)?vice(?:land|tv))\\.com/(?P<locale>[^/]+)/(?:video/[^/]+|embed)/(?P<id>[\\da-f]{24})")
			ViceIE__extract_urls = λ.NewFunction("_extract_urls",
				[]λ.Param{
					{Name: "webpage"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒwebpage = λargs[0]
					)
					return λ.Cal(Ωre.ϒfindall, λ.StrLiteral("<iframe\\b[^>]+\\bsrc=[\"\\']((?:https?:)?//video\\.vice\\.com/[^/]+/embed/[\\da-f]{24})"), ϒwebpage)
				})
			ViceIE__extract_urls = λ.Cal(λ.StaticMethodType, ViceIE__extract_urls)
			ViceIE__extract_url = λ.NewFunction("_extract_url",
				[]λ.Param{
					{Name: "webpage"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒurls    λ.Object
						ϒwebpage = λargs[0]
					)
					ϒurls = λ.Calm(ViceIE, "_extract_urls", ϒwebpage)
					return func() λ.Object {
						if λ.IsTrue(ϒurls) {
							return λ.GetItem(ϒurls, λ.IntLiteral(0))
						} else {
							return λ.None
						}
					}()
				})
			ViceIE__extract_url = λ.Cal(λ.StaticMethodType, ViceIE__extract_url)
			ViceIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcc_url        λ.Object
						ϒchannel       λ.Object
						ϒepisode       λ.Object
						ϒerror         λ.Object
						ϒerror_message λ.Object
						ϒexp           λ.Object
						ϒformats       λ.Object
						ϒlanguage_code λ.Object
						ϒlocale        λ.Object
						ϒpreplay       λ.Object
						ϒquery         λ.Object
						ϒrating        λ.Object
						ϒresource      λ.Object
						ϒseason        λ.Object
						ϒself          = λargs[0]
						ϒsubtitle      λ.Object
						ϒsubtitles     λ.Object
						ϒtitle         λ.Object
						ϒurl           = λargs[1]
						ϒvideo         λ.Object
						ϒvideo_data    λ.Object
						ϒvideo_id      λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
					)
					τmp0 = λ.Calm(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups")
					ϒlocale = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒvideo_id = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒvideo = λ.GetItem(λ.Calm(ϒself, "_call_api", λ.StrLiteral("videos"), λ.StrLiteral("id"), ϒvideo_id, ϒlocale, λ.StrLiteral("body\n    locked\n    rating\n    thumbnail_url\n    title")), λ.IntLiteral(0))
					ϒtitle = λ.Calm(λ.GetItem(ϒvideo, λ.StrLiteral("title")), "strip")
					ϒrating = λ.Calm(ϒvideo, "get", λ.StrLiteral("rating"))
					ϒquery = λ.DictLiteral(map[λ.Object]λ.Object{})
					if λ.IsTrue(λ.Calm(ϒvideo, "get", λ.StrLiteral("locked"))) {
						ϒresource = λ.Calm(ϒself, "_get_mvpd_resource", λ.StrLiteral("VICELAND"), ϒtitle, ϒvideo_id, ϒrating)
						λ.SetItem(ϒquery, λ.StrLiteral("tvetoken"), λ.Calm(ϒself, "_extract_mvpd_auth", ϒurl, ϒvideo_id, λ.StrLiteral("VICELAND"), ϒresource))
					}
					ϒexp = λ.Add(λ.Cal(λ.IntType, λ.Cal(Ωtime.ϒtime)), λ.IntLiteral(1440))
					λ.Calm(ϒquery, "update", λ.DictLiteral(map[string]λ.Object{
						"exp": ϒexp,
						"sign": λ.Calm(λ.Calm(λ.None, "sha512", λ.Calm(λ.Mod(λ.StrLiteral("%s:GET:%d"), λ.NewTuple(
							ϒvideo_id,
							ϒexp,
						)), "encode")), "hexdigest"),
						"skipadstitching": λ.IntLiteral(1),
						"platform":        λ.StrLiteral("desktop"),
						"rn":              λ.Cal(λ.None, λ.IntLiteral(10000), λ.IntLiteral(100000)),
					}))
					τmp0, τmp1 = func() (λexit λ.Object, λret λ.Object) {
						defer λ.CatchMulti(
							nil,
							&λ.Catcher{ExtractorError, func(λex λ.BaseException) {
								var ϒe λ.Object = λex
								if λ.IsTrue(func() λ.Object {
									if λv := λ.Cal(λ.BuiltinIsInstance, λ.GetAttr(ϒe, "cause", nil), λ.None); !λ.IsTrue(λv) {
										return λv
									} else {
										return λ.NewBool(λ.Contains(λ.NewTuple(
											λ.IntLiteral(400),
											λ.IntLiteral(401),
										), λ.GetAttr(λ.GetAttr(ϒe, "cause", nil), "code", nil)))
									}
								}()) {
									ϒerror = λ.Cal(Ωjson.ϒloads, λ.Calm(λ.Calm(λ.GetAttr(ϒe, "cause", nil), "read"), "decode"))
									ϒerror_message = func() λ.Object {
										if λv := λ.Calm(ϒerror, "get", λ.StrLiteral("error_description")); λ.IsTrue(λv) {
											return λv
										} else {
											return λ.GetItem(ϒerror, λ.StrLiteral("details"))
										}
									}()
									panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.StrLiteral("%s said: %s"), λ.NewTuple(
										λ.GetAttr(ϒself, "IE_NAME", nil),
										ϒerror_message,
									))), λ.KWArgs{
										{Name: "expected", Value: λ.True},
									})))
								}
								panic(λ.Raise(λex))
							}},
						)
						ϒpreplay = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
							λ.Mod(λ.StrLiteral("https://vms.vice.com/%s/video/preplay/%s"), λ.NewTuple(
								ϒlocale,
								ϒvideo_id,
							)),
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "query", Value: ϒquery},
						})
						return λ.BlockExitNormally, nil
					}()
					ϒvideo_data = λ.GetItem(ϒpreplay, λ.StrLiteral("video"))
					ϒformats = λ.Calm(ϒself, "_extract_m3u8_formats", λ.GetItem(ϒpreplay, λ.StrLiteral("playURL")), ϒvideo_id, λ.StrLiteral("mp4"), λ.StrLiteral("m3u8_native"))
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒepisode = func() λ.Object {
						if λv := λ.Calm(ϒvideo_data, "get", λ.StrLiteral("episode")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.DictLiteral(map[λ.Object]λ.Object{})
						}
					}()
					ϒchannel = func() λ.Object {
						if λv := λ.Calm(ϒvideo_data, "get", λ.StrLiteral("channel")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.DictLiteral(map[λ.Object]λ.Object{})
						}
					}()
					ϒseason = func() λ.Object {
						if λv := λ.Calm(ϒvideo_data, "get", λ.StrLiteral("season")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.DictLiteral(map[λ.Object]λ.Object{})
						}
					}()
					ϒsubtitles = λ.DictLiteral(map[λ.Object]λ.Object{})
					τmp1 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒpreplay, "get", λ.StrLiteral("subtitleURLs"), λ.NewList()))
					for {
						if τmp0 = λ.NextDefault(τmp1, λ.AfterLast); τmp0 == λ.AfterLast {
							break
						}
						ϒsubtitle = τmp0
						ϒcc_url = λ.Calm(ϒsubtitle, "get", λ.StrLiteral("url"))
						if !λ.IsTrue(ϒcc_url) {
							continue
						}
						ϒlanguage_code = func() λ.Object {
							if λv := λ.Cal(ϒtry_get, ϒsubtitle, λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("languages")), λ.IntLiteral(0)), λ.StrLiteral("language_code"))
								}), ϒcompat_str); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.StrLiteral("en")
							}
						}()
						λ.Calm(λ.Calm(ϒsubtitles, "setdefault", ϒlanguage_code, λ.NewList()), "append", λ.DictLiteral(map[string]λ.Object{
							"url": ϒcc_url,
						}))
					}
					return λ.DictLiteral(map[string]λ.Object{
						"formats":     ϒformats,
						"id":          ϒvideo_id,
						"title":       ϒtitle,
						"description": λ.Cal(ϒclean_html, λ.Calm(ϒvideo, "get", λ.StrLiteral("body"))),
						"thumbnail":   λ.Calm(ϒvideo, "get", λ.StrLiteral("thumbnail_url")),
						"duration":    λ.Cal(ϒint_or_none, λ.Calm(ϒvideo_data, "get", λ.StrLiteral("video_duration"))),
						"timestamp":   λ.Cal(ϒint_or_none, λ.Calm(ϒvideo_data, "get", λ.StrLiteral("created_at")), λ.IntLiteral(1000)),
						"age_limit": λ.Cal(ϒparse_age_limit, func() λ.Object {
							if λv := λ.Calm(ϒvideo_data, "get", λ.StrLiteral("video_rating")); λ.IsTrue(λv) {
								return λv
							} else {
								return ϒrating
							}
						}()),
						"series": λ.Cal(ϒtry_get, ϒvideo_data, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("show")), λ.StrLiteral("base")), λ.StrLiteral("display_title"))
							}), ϒcompat_str),
						"episode_number": λ.Cal(ϒint_or_none, λ.Calm(ϒepisode, "get", λ.StrLiteral("episode_number"))),
						"episode_id": λ.Cal(ϒstr_or_none, func() λ.Object {
							if λv := λ.Calm(ϒepisode, "get", λ.StrLiteral("id")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(ϒvideo_data, "get", λ.StrLiteral("episode_id"))
							}
						}()),
						"season_number": λ.Cal(ϒint_or_none, λ.Calm(ϒseason, "get", λ.StrLiteral("season_number"))),
						"season_id": λ.Cal(ϒstr_or_none, func() λ.Object {
							if λv := λ.Calm(ϒseason, "get", λ.StrLiteral("id")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(ϒvideo_data, "get", λ.StrLiteral("season_id"))
							}
						}()),
						"uploader":    λ.Calm(ϒchannel, "get", λ.StrLiteral("name")),
						"uploader_id": λ.Cal(ϒstr_or_none, λ.Calm(ϒchannel, "get", λ.StrLiteral("id"))),
						"subtitles":   ϒsubtitles,
					})
				})
			return λ.DictLiteral(map[string]λ.Object{
				"IE_NAME":       ViceIE_IE_NAME,
				"_VALID_URL":    ViceIE__VALID_URL,
				"_extract_url":  ViceIE__extract_url,
				"_extract_urls": ViceIE__extract_urls,
				"_real_extract": ViceIE__real_extract,
			})
		}())
		ViceShowIE = λ.Cal(λ.TypeType, λ.StrLiteral("ViceShowIE"), λ.NewTuple(ViceBaseIE), func() λ.Dict {
			var (
				ViceShowIE__VALID_URL λ.Object
			)
			ViceShowIE__VALID_URL = λ.StrLiteral("https?://(?:video\\.vice|(?:www\\.)?vice(?:land|tv))\\.com/(?P<locale>[^/]+)/show/(?P<id>[^/?#&]+)")
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL": ViceShowIE__VALID_URL,
			})
		}())
		ViceArticleIE = λ.Cal(λ.TypeType, λ.StrLiteral("ViceArticleIE"), λ.NewTuple(ViceBaseIE), func() λ.Dict {
			var (
				ViceArticleIE_IE_NAME       λ.Object
				ViceArticleIE__VALID_URL    λ.Object
				ViceArticleIE__real_extract λ.Object
			)
			ViceArticleIE_IE_NAME = λ.StrLiteral("vice:article")
			ViceArticleIE__VALID_URL = λ.StrLiteral("https://(?:www\\.)?vice\\.com/(?P<locale>[^/]+)/article/(?:[0-9a-z]{6}/)?(?P<id>[^?#]+)")
			ViceArticleIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒ_url_res    λ.Object
						ϒarticle     λ.Object
						ϒbody        λ.Object
						ϒdisplay_id  λ.Object
						ϒembed_code  λ.Object
						ϒlocale      λ.Object
						ϒself        = λargs[0]
						ϒurl         = λargs[1]
						ϒvice_url    λ.Object
						ϒvideo_url   λ.Object
						ϒyoutube_url λ.Object
						τmp0         λ.Object
					)
					τmp0 = λ.Calm(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups")
					ϒlocale = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒdisplay_id = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒarticle = λ.GetItem(λ.Calm(ϒself, "_call_api", λ.StrLiteral("articles"), λ.StrLiteral("slug"), ϒdisplay_id, ϒlocale, λ.StrLiteral("body\n    embed_code")), λ.IntLiteral(0))
					ϒbody = λ.GetItem(ϒarticle, λ.StrLiteral("body"))
					ϒ_url_res = λ.NewFunction("_url_res",
						[]λ.Param{
							{Name: "video_url"},
							{Name: "ie_key"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒie_key    = λargs[1]
								ϒvideo_url = λargs[0]
							)
							return λ.DictLiteral(map[string]λ.Object{
								"_type":      λ.StrLiteral("url_transparent"),
								"url":        ϒvideo_url,
								"display_id": ϒdisplay_id,
								"ie_key":     ϒie_key,
							})
						})
					ϒvice_url = λ.Calm(ViceIE, "_extract_url", ϒbody)
					if λ.IsTrue(ϒvice_url) {
						return λ.Cal(ϒ_url_res, ϒvice_url, λ.Calm(ViceIE, "ie_key"))
					}
					ϒembed_code = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("embedCode=([^&\\'\"]+)"),
						ϒbody,
						λ.StrLiteral("ooyala embed code"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒembed_code) {
						return λ.Cal(ϒ_url_res, λ.Mod(λ.StrLiteral("ooyala:%s"), ϒembed_code), λ.StrLiteral("Ooyala"))
					}
					ϒyoutube_url = λ.Calm(YoutubeIE, "_extract_url", ϒbody)
					if λ.IsTrue(ϒyoutube_url) {
						return λ.Cal(ϒ_url_res, ϒyoutube_url, λ.Calm(YoutubeIE, "ie_key"))
					}
					ϒvideo_url = λ.Calm(ϒself, "_html_search_regex", λ.StrLiteral("data-video-url=\"([^\"]+)\""), λ.GetItem(ϒarticle, λ.StrLiteral("embed_code")), λ.StrLiteral("video URL"))
					return λ.Cal(ϒ_url_res, ϒvideo_url, λ.Calm(ViceIE, "ie_key"))
				})
			return λ.DictLiteral(map[string]λ.Object{
				"IE_NAME":       ViceArticleIE_IE_NAME,
				"_VALID_URL":    ViceArticleIE__VALID_URL,
				"_real_extract": ViceArticleIE__real_extract,
			})
		}())
	})
}
