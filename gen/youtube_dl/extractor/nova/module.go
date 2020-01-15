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
 * nova/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/nova.py
 */

package nova

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor    λ.Object
	NovaEmbedIE      λ.Object
	NovaIE           λ.Object
	ϒclean_html      λ.Object
	ϒint_or_none     λ.Object
	ϒjs_to_json      λ.Object
	ϒqualities       λ.Object
	ϒunified_strdate λ.Object
	ϒurl_or_none     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒclean_html = Ωutils.ϒclean_html
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒjs_to_json = Ωutils.ϒjs_to_json
		ϒqualities = Ωutils.ϒqualities
		ϒunified_strdate = Ωutils.ϒunified_strdate
		ϒurl_or_none = Ωutils.ϒurl_or_none
		NovaEmbedIE = λ.Cal(λ.TypeType, λ.StrLiteral("NovaEmbedIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				NovaEmbedIE__VALID_URL λ.Object
			)
			NovaEmbedIE__VALID_URL = λ.StrLiteral("https?://media\\.cms\\.nova\\.cz/embed/(?P<id>[^/?#&]+)")
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL": NovaEmbedIE__VALID_URL,
			})
		}())
		NovaIE = λ.Cal(λ.TypeType, λ.StrLiteral("NovaIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				NovaIE__VALID_URL    λ.Object
				NovaIE__real_extract λ.Object
			)
			NovaIE__VALID_URL = λ.StrLiteral("https?://(?:[^.]+\\.)?(?P<site>tv(?:noviny)?|tn|novaplus|vymena|fanda|krasna|doma|prask)\\.nova\\.cz/(?:[^/]+/)+(?P<id>[^/]+?)(?:\\.html|/|$)")
			NovaIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						DEFAULT_SITE_ID λ.Object
						SITES           λ.Object
						ϒconfig         λ.Object
						ϒconfig_params  λ.Object
						ϒconfig_url     λ.Object
						ϒdescription    λ.Object
						ϒdisplay_id     λ.Object
						ϒembed_id       λ.Object
						ϒformats        λ.Object
						ϒm              λ.Object
						ϒmediafile      λ.Object
						ϒmobj           λ.Object
						ϒparams         λ.Object
						ϒplayer         λ.Object
						ϒself           = λargs[0]
						ϒsite           λ.Object
						ϒsite_id        λ.Object
						ϒthumbnail      λ.Object
						ϒtitle          λ.Object
						ϒupload_date    λ.Object
						ϒurl            = λargs[1]
						ϒvideo_id       λ.Object
						ϒvideo_url      λ.Object
						ϒwebpage        λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒdisplay_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("id"))
					ϒsite = λ.Calm(ϒmobj, "group", λ.StrLiteral("site"))
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒdisplay_id)
					ϒembed_id = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("<iframe[^>]+\\bsrc=[\"\\'](?:https?:)?//media\\.cms\\.nova\\.cz/embed/([^/?#&]+)"),
						ϒwebpage,
						λ.StrLiteral("embed url"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒembed_id) {
						return λ.Call(λ.GetAttr(ϒself, "url_result", nil), λ.NewArgs(λ.Mod(λ.StrLiteral("https://media.cms.nova.cz/embed/%s"), ϒembed_id)), λ.KWArgs{
							{Name: "ie", Value: λ.Calm(NovaEmbedIE, "ie_key")},
							{Name: "video_id", Value: ϒembed_id},
						})
					}
					ϒvideo_id = λ.Calm(ϒself, "_search_regex", λ.NewList(
						λ.StrLiteral("(?:media|video_id)\\s*:\\s*'(\\d+)'"),
						λ.StrLiteral("media=(\\d+)"),
						λ.StrLiteral("id=\"article_video_(\\d+)\""),
						λ.StrLiteral("id=\"player_(\\d+)\""),
					), ϒwebpage, λ.StrLiteral("video id"))
					ϒconfig_url = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("src=\"(https?://(?:tn|api)\\.nova\\.cz/bin/player/videojs/config\\.php\\?[^\"]+)\""),
						ϒwebpage,
						λ.StrLiteral("config url"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					ϒconfig_params = λ.DictLiteral(map[λ.Object]λ.Object{})
					if !λ.IsTrue(ϒconfig_url) {
						ϒplayer = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
							λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("(?s)Player\\s*\\(.+?\\s*,\\s*({.+?\\bmedia\\b[\"\\']?\\s*:\\s*[\"\\']?\\d+.+?})\\s*\\)"),
								ϒwebpage,
								λ.StrLiteral("player"),
							), λ.KWArgs{
								{Name: "default", Value: λ.StrLiteral("{}")},
							}),
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "transform_source", Value: ϒjs_to_json},
							{Name: "fatal", Value: λ.False},
						})
						if λ.IsTrue(ϒplayer) {
							ϒconfig_url = λ.Cal(ϒurl_or_none, λ.Calm(ϒplayer, "get", λ.StrLiteral("configUrl")))
							ϒparams = λ.Calm(ϒplayer, "get", λ.StrLiteral("configParams"))
							if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒparams, λ.DictType)) {
								ϒconfig_params = ϒparams
							}
						}
					}
					if !λ.IsTrue(ϒconfig_url) {
						DEFAULT_SITE_ID = λ.StrLiteral("23000")
						SITES = λ.DictLiteral(map[string]λ.Object{
							"tvnoviny": DEFAULT_SITE_ID,
							"novaplus": DEFAULT_SITE_ID,
							"vymena":   DEFAULT_SITE_ID,
							"krasna":   DEFAULT_SITE_ID,
							"fanda":    λ.StrLiteral("30"),
							"tn":       λ.StrLiteral("30"),
							"doma":     λ.StrLiteral("30"),
						})
						ϒsite_id = func() λ.Object {
							if λv := λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("site=(\\d+)"),
								ϒwebpage,
								λ.StrLiteral("site id"),
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							}); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(SITES, "get", ϒsite, DEFAULT_SITE_ID)
							}
						}()
						ϒconfig_url = λ.StrLiteral("https://api.nova.cz/bin/player/videojs/config.php")
						ϒconfig_params = λ.DictLiteral(map[string]λ.Object{
							"site":    ϒsite_id,
							"media":   ϒvideo_id,
							"quality": λ.IntLiteral(3),
							"version": λ.IntLiteral(1),
						})
					}
					ϒconfig = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						ϒconfig_url,
						ϒdisplay_id,
						λ.StrLiteral("Downloading config JSON"),
					), λ.KWArgs{
						{Name: "query", Value: ϒconfig_params},
						{Name: "transform_source", Value: λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "s"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒs = λargs[0]
								)
								return λ.GetItem(ϒs, λ.NewSlice(λ.Calm(ϒs, "index", λ.StrLiteral("{")), λ.Add(λ.Calm(ϒs, "rindex", λ.StrLiteral("}")), λ.IntLiteral(1)), λ.None))
							})},
					})
					ϒmediafile = λ.GetItem(ϒconfig, λ.StrLiteral("mediafile"))
					ϒvideo_url = λ.GetItem(ϒmediafile, λ.StrLiteral("src"))
					ϒm = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("^(?P<url>rtmpe?://[^/]+/(?P<app>[^/]+?))/&*(?P<playpath>.+)$"), ϒvideo_url)
					if λ.IsTrue(ϒm) {
						ϒformats = λ.NewList(λ.DictLiteral(map[string]λ.Object{
							"url":         λ.Calm(ϒm, "group", λ.StrLiteral("url")),
							"app":         λ.Calm(ϒm, "group", λ.StrLiteral("app")),
							"play_path":   λ.Calm(ϒm, "group", λ.StrLiteral("playpath")),
							"player_path": λ.StrLiteral("http://tvnoviny.nova.cz/static/shared/app/videojs/video-js.swf"),
							"ext":         λ.StrLiteral("flv"),
						}))
					} else {
						ϒformats = λ.NewList(λ.DictLiteral(map[string]λ.Object{
							"url": ϒvideo_url,
						}))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒtitle = func() λ.Object {
						if λv := λ.Calm(λ.Calm(ϒmediafile, "get", λ.StrLiteral("meta"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("title")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒself, "_og_search_title", ϒwebpage)
						}
					}()
					ϒdescription = λ.Cal(ϒclean_html, λ.Call(λ.GetAttr(ϒself, "_og_search_description", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
						{Name: "default", Value: λ.None},
					}))
					ϒthumbnail = λ.Calm(ϒconfig, "get", λ.StrLiteral("poster"))
					if λ.IsTrue(λ.Eq(ϒsite, λ.StrLiteral("novaplus"))) {
						ϒupload_date = λ.Cal(ϒunified_strdate, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.StrLiteral("(\\d{1,2}-\\d{1,2}-\\d{4})$"),
							ϒdisplay_id,
							λ.StrLiteral("upload date"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}))
					} else {
						if λ.IsTrue(λ.Eq(ϒsite, λ.StrLiteral("fanda"))) {
							ϒupload_date = λ.Cal(ϒunified_strdate, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("<span class=\"date_time\">(\\d{1,2}\\.\\d{1,2}\\.\\d{4})"),
								ϒwebpage,
								λ.StrLiteral("upload date"),
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							}))
						} else {
							ϒupload_date = λ.None
						}
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id":          ϒvideo_id,
						"display_id":  ϒdisplay_id,
						"title":       ϒtitle,
						"description": ϒdescription,
						"upload_date": ϒupload_date,
						"thumbnail":   ϒthumbnail,
						"formats":     ϒformats,
					})
				})
			return λ.DictLiteral(map[string]λ.Object{
				"_VALID_URL":    NovaIE__VALID_URL,
				"_real_extract": NovaIE__real_extract,
			})
		}())
	})
}
