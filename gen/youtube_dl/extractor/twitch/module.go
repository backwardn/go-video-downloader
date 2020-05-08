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
 * twitch/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/twitch.py
 */

package twitch

import (
	Ωjson "github.com/tenta-browser/go-video-downloader/gen/json"
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError                 λ.Object
	InfoExtractor                  λ.Object
	TwitchAllVideosIE              λ.Object
	TwitchBaseIE                   λ.Object
	TwitchChapterIE                λ.Object
	TwitchClipsIE                  λ.Object
	TwitchHighlightsIE             λ.Object
	TwitchItemBaseIE               λ.Object
	TwitchPastBroadcastsIE         λ.Object
	TwitchPlaylistBaseIE           λ.Object
	TwitchProfileIE                λ.Object
	TwitchStreamIE                 λ.Object
	TwitchUploadsIE                λ.Object
	TwitchVideoIE                  λ.Object
	TwitchVideosBaseIE             λ.Object
	TwitchVodIE                    λ.Object
	ϒclean_html                    λ.Object
	ϒcompat_kwargs                 λ.Object
	ϒcompat_parse_qs               λ.Object
	ϒcompat_str                    λ.Object
	ϒcompat_urllib_parse_urlencode λ.Object
	ϒcompat_urllib_parse_urlparse  λ.Object
	ϒint_or_none                   λ.Object
	ϒparse_duration                λ.Object
	ϒparse_iso8601                 λ.Object
	ϒtry_get                       λ.Object
	ϒunified_timestamp             λ.Object
	ϒupdate_url_query              λ.Object
	ϒurl_or_none                   λ.Object
	ϒurljoin                       λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_kwargs = Ωcompat.ϒcompat_kwargs
		ϒcompat_parse_qs = Ωcompat.ϒcompat_parse_qs
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒcompat_urllib_parse_urlencode = Ωcompat.ϒcompat_urllib_parse_urlencode
		ϒcompat_urllib_parse_urlparse = Ωcompat.ϒcompat_urllib_parse_urlparse
		ϒclean_html = Ωutils.ϒclean_html
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒtry_get = Ωutils.ϒtry_get
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		ϒurl_or_none = Ωutils.ϒurl_or_none
		ϒurljoin = Ωutils.ϒurljoin
		TwitchBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitchBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TwitchBaseIE__CLIENT_ID       λ.Object
				TwitchBaseIE__NETRC_MACHINE   λ.Object
				TwitchBaseIE__VALID_URL_BASE  λ.Object
				TwitchBaseIE__login           λ.Object
				TwitchBaseIE__real_initialize λ.Object
			)
			TwitchBaseIE__VALID_URL_BASE = λ.StrLiteral("https?://(?:(?:www|go|m)\\.)?twitch\\.tv")
			TwitchBaseIE__CLIENT_ID = λ.StrLiteral("kimne78kx3ncx6brgo4mv6wki5h1ko")
			TwitchBaseIE__NETRC_MACHINE = λ.StrLiteral("twitch")
			TwitchBaseIE__real_initialize = λ.NewFunction("_real_initialize",
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
			TwitchBaseIE__login = λ.NewFunction("_login",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒfail          λ.Object
						ϒhandle        λ.Object
						ϒlogin_page    λ.Object
						ϒlogin_step    λ.Object
						ϒpassword      λ.Object
						ϒredirect_page λ.Object
						ϒself          = λargs[0]
						ϒtfa_token     λ.Object
						ϒusername      λ.Object
						τmp0           λ.Object
					)
					τmp0 = λ.Calm(ϒself, "_get_login_info")
					ϒusername = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒpassword = λ.GetItem(τmp0, λ.IntLiteral(1))
					if ϒusername == λ.None {
						return λ.None
					}
					ϒfail = λ.NewFunction("fail",
						[]λ.Param{
							{Name: "message"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒmessage = λargs[0]
							)
							panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.StrLiteral("Unable to login. Twitch said: %s"), ϒmessage)), λ.KWArgs{
								{Name: "expected", Value: λ.True},
							})))
							return λ.None
						})
					ϒlogin_step = λ.NewFunction("login_step",
						[]λ.Param{
							{Name: "page"},
							{Name: "urlh"},
							{Name: "note"},
							{Name: "data"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒdata         = λargs[3]
								ϒerror        λ.Object
								ϒform         λ.Object
								ϒheaders      λ.Object
								ϒnote         = λargs[2]
								ϒpage         = λargs[0]
								ϒpage_url     λ.Object
								ϒpost_url     λ.Object
								ϒredirect_url λ.Object
								ϒresponse     λ.Object
								ϒurlh         = λargs[1]
							)
							ϒform = λ.Calm(ϒself, "_hidden_inputs", ϒpage)
							λ.Calm(ϒform, "update", ϒdata)
							ϒpage_url = λ.Calm(ϒurlh, "geturl")
							ϒpost_url = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("<form[^>]+action=([\"\\'])(?P<url>.+?)\\1"),
								ϒpage,
								λ.StrLiteral("post url"),
							), λ.KWArgs{
								{Name: "default", Value: λ.GetAttr(ϒself, "_LOGIN_POST_URL", nil)},
								{Name: "group", Value: λ.StrLiteral("url")},
							})
							ϒpost_url = λ.Cal(ϒurljoin, ϒpage_url, ϒpost_url)
							ϒheaders = λ.DictLiteral(map[string]λ.Object{
								"Referer":      ϒpage_url,
								"Origin":       ϒpage_url,
								"Content-Type": λ.StrLiteral("text/plain;charset=UTF-8"),
							})
							ϒresponse = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
								ϒpost_url,
								λ.None,
								ϒnote,
							), λ.KWArgs{
								{Name: "data", Value: λ.Calm(λ.Cal(Ωjson.ϒdumps, ϒform), "encode")},
								{Name: "headers", Value: ϒheaders},
								{Name: "expected_status", Value: λ.IntLiteral(400)},
							})
							ϒerror = func() λ.Object {
								if λv := λ.Calm(ϒresponse, "get", λ.StrLiteral("error_description")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Calm(ϒresponse, "get", λ.StrLiteral("error_code"))
								}
							}()
							if λ.IsTrue(ϒerror) {
								λ.Cal(ϒfail, ϒerror)
							}
							if λ.Contains(λ.Calm(ϒresponse, "get", λ.StrLiteral("message"), λ.StrLiteral("")), λ.StrLiteral("Authenticated successfully")) {
								return λ.NewTuple(
									λ.None,
									λ.None,
								)
							}
							ϒredirect_url = λ.Cal(ϒurljoin, ϒpost_url, func() λ.Object {
								if λv := λ.Calm(ϒresponse, "get", λ.StrLiteral("redirect")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.GetItem(ϒresponse, λ.StrLiteral("redirect_path"))
								}
							}())
							return λ.Call(λ.GetAttr(ϒself, "_download_webpage_handle", nil), λ.NewArgs(
								ϒredirect_url,
								λ.None,
								λ.StrLiteral("Downloading login redirect page"),
							), λ.KWArgs{
								{Name: "headers", Value: ϒheaders},
							})
						})
					τmp0 = λ.Calm(ϒself, "_download_webpage_handle", λ.GetAttr(ϒself, "_LOGIN_FORM_URL", nil), λ.None, λ.StrLiteral("Downloading login page"))
					ϒlogin_page = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒhandle = λ.GetItem(τmp0, λ.IntLiteral(1))
					if λ.Contains(ϒlogin_page, λ.StrLiteral("blacklist_message")) {
						λ.Cal(ϒfail, λ.Cal(ϒclean_html, ϒlogin_page))
					}
					τmp0 = λ.Cal(ϒlogin_step, ϒlogin_page, ϒhandle, λ.StrLiteral("Logging in"), λ.DictLiteral(map[string]λ.Object{
						"username":  ϒusername,
						"password":  ϒpassword,
						"client_id": λ.GetAttr(ϒself, "_CLIENT_ID", nil),
					}))
					ϒredirect_page = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒhandle = λ.GetItem(τmp0, λ.IntLiteral(1))
					if !λ.IsTrue(ϒredirect_page) {
						return λ.None
					}
					if λ.Cal(Ωre.ϒsearch, λ.StrLiteral("(?i)<form[^>]+id=\"two-factor-submit\""), ϒredirect_page) != λ.None {
						ϒtfa_token = λ.Calm(ϒself, "_get_tfa_info", λ.StrLiteral("two-factor authentication token"))
						λ.Cal(ϒlogin_step, ϒredirect_page, ϒhandle, λ.StrLiteral("Submitting TFA token"), λ.DictLiteral(map[string]λ.Object{
							"authy_token":  ϒtfa_token,
							"remember_2fa": λ.StrLiteral("true"),
						}))
					}
					return λ.None
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_CLIENT_ID":       TwitchBaseIE__CLIENT_ID,
				"_NETRC_MACHINE":   TwitchBaseIE__NETRC_MACHINE,
				"_VALID_URL_BASE":  TwitchBaseIE__VALID_URL_BASE,
				"_login":           TwitchBaseIE__login,
				"_real_initialize": TwitchBaseIE__real_initialize,
			})
		}())
		TwitchItemBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitchItemBaseIE"), λ.NewTuple(TwitchBaseIE), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		TwitchVideoIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitchVideoIE"), λ.NewTuple(TwitchItemBaseIE), func() λ.Dict {
			var (
				TwitchVideoIE__VALID_URL λ.Object
			)
			TwitchVideoIE__VALID_URL = λ.Mod(λ.StrLiteral("%s/[^/]+/b/(?P<id>\\d+)"), λ.GetAttr(TwitchBaseIE, "_VALID_URL_BASE", nil))
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": TwitchVideoIE__VALID_URL,
			})
		}())
		TwitchChapterIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitchChapterIE"), λ.NewTuple(TwitchItemBaseIE), func() λ.Dict {
			var (
				TwitchChapterIE__VALID_URL λ.Object
			)
			TwitchChapterIE__VALID_URL = λ.Mod(λ.StrLiteral("%s/[^/]+/c/(?P<id>\\d+)"), λ.GetAttr(TwitchBaseIE, "_VALID_URL_BASE", nil))
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": TwitchChapterIE__VALID_URL,
			})
		}())
		TwitchVodIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitchVodIE"), λ.NewTuple(TwitchItemBaseIE), func() λ.Dict {
			var (
				TwitchVodIE__VALID_URL λ.Object
			)
			TwitchVodIE__VALID_URL = λ.StrLiteral("(?x)\n                    https?://\n                        (?:\n                            (?:(?:www|go|m)\\.)?twitch\\.tv/(?:[^/]+/v(?:ideo)?|videos)/|\n                            player\\.twitch\\.tv/\\?.*?\\bvideo=v?\n                        )\n                        (?P<id>\\d+)\n                    ")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": TwitchVodIE__VALID_URL,
			})
		}())
		TwitchPlaylistBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitchPlaylistBaseIE"), λ.NewTuple(TwitchBaseIE), func() λ.Dict {
			var (
				TwitchPlaylistBaseIE__PLAYLIST_PATH λ.Object
			)
			TwitchPlaylistBaseIE__PLAYLIST_PATH = λ.StrLiteral("kraken/channels/%s/videos/?offset=%d&limit=%d")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_PLAYLIST_PATH": TwitchPlaylistBaseIE__PLAYLIST_PATH,
			})
		}())
		TwitchProfileIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitchProfileIE"), λ.NewTuple(TwitchPlaylistBaseIE), func() λ.Dict {
			var (
				TwitchProfileIE__VALID_URL λ.Object
			)
			TwitchProfileIE__VALID_URL = λ.Mod(λ.StrLiteral("%s/(?P<id>[^/]+)/profile/?(?:\\#.*)?$"), λ.GetAttr(TwitchBaseIE, "_VALID_URL_BASE", nil))
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": TwitchProfileIE__VALID_URL,
			})
		}())
		TwitchVideosBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitchVideosBaseIE"), λ.NewTuple(TwitchPlaylistBaseIE), func() λ.Dict {
			var (
				TwitchVideosBaseIE__PLAYLIST_PATH         λ.Object
				TwitchVideosBaseIE__VALID_URL_VIDEOS_BASE λ.Object
			)
			TwitchVideosBaseIE__VALID_URL_VIDEOS_BASE = λ.Mod(λ.StrLiteral("%s/(?P<id>[^/]+)/videos"), λ.GetAttr(TwitchBaseIE, "_VALID_URL_BASE", nil))
			TwitchVideosBaseIE__PLAYLIST_PATH = λ.Add(λ.GetAttr(TwitchPlaylistBaseIE, "_PLAYLIST_PATH", nil), λ.StrLiteral("&broadcast_type="))
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_PLAYLIST_PATH":         TwitchVideosBaseIE__PLAYLIST_PATH,
				"_VALID_URL_VIDEOS_BASE": TwitchVideosBaseIE__VALID_URL_VIDEOS_BASE,
			})
		}())
		TwitchAllVideosIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitchAllVideosIE"), λ.NewTuple(TwitchVideosBaseIE), func() λ.Dict {
			var (
				TwitchAllVideosIE__VALID_URL λ.Object
			)
			TwitchAllVideosIE__VALID_URL = λ.Mod(λ.StrLiteral("%s/all"), λ.GetAttr(TwitchVideosBaseIE, "_VALID_URL_VIDEOS_BASE", nil))
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": TwitchAllVideosIE__VALID_URL,
			})
		}())
		TwitchUploadsIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitchUploadsIE"), λ.NewTuple(TwitchVideosBaseIE), func() λ.Dict {
			var (
				TwitchUploadsIE__VALID_URL λ.Object
			)
			TwitchUploadsIE__VALID_URL = λ.Mod(λ.StrLiteral("%s/uploads"), λ.GetAttr(TwitchVideosBaseIE, "_VALID_URL_VIDEOS_BASE", nil))
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": TwitchUploadsIE__VALID_URL,
			})
		}())
		TwitchPastBroadcastsIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitchPastBroadcastsIE"), λ.NewTuple(TwitchVideosBaseIE), func() λ.Dict {
			var (
				TwitchPastBroadcastsIE__VALID_URL λ.Object
			)
			TwitchPastBroadcastsIE__VALID_URL = λ.Mod(λ.StrLiteral("%s/past-broadcasts"), λ.GetAttr(TwitchVideosBaseIE, "_VALID_URL_VIDEOS_BASE", nil))
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": TwitchPastBroadcastsIE__VALID_URL,
			})
		}())
		TwitchHighlightsIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitchHighlightsIE"), λ.NewTuple(TwitchVideosBaseIE), func() λ.Dict {
			var (
				TwitchHighlightsIE__VALID_URL λ.Object
			)
			TwitchHighlightsIE__VALID_URL = λ.Mod(λ.StrLiteral("%s/highlights"), λ.GetAttr(TwitchVideosBaseIE, "_VALID_URL_VIDEOS_BASE", nil))
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": TwitchHighlightsIE__VALID_URL,
			})
		}())
		TwitchStreamIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitchStreamIE"), λ.NewTuple(TwitchBaseIE), func() λ.Dict {
			var (
				TwitchStreamIE__VALID_URL λ.Object
				TwitchStreamIE_suitable   λ.Object
			)
			TwitchStreamIE__VALID_URL = λ.StrLiteral("(?x)\n                    https?://\n                        (?:\n                            (?:(?:www|go|m)\\.)?twitch\\.tv/|\n                            player\\.twitch\\.tv/\\?.*?\\bchannel=\n                        )\n                        (?P<id>[^/#?]+)\n                    ")
			TwitchStreamIE_suitable = λ.NewFunction("suitable",
				[]λ.Param{
					{Name: "cls"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcls = λargs[0]
						ϒurl = λargs[1]
					)
					return func() λ.Object {
						if λ.IsTrue(λ.Cal(λ.BuiltinAny, λ.Cal(λ.NewFunction("<generator>",
							nil,
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
									var (
										ϒie  λ.Object
										τmp0 λ.Object
										τmp1 λ.Object
									)
									τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
										TwitchVideoIE,
										TwitchChapterIE,
										TwitchVodIE,
										TwitchProfileIE,
										TwitchAllVideosIE,
										TwitchUploadsIE,
										TwitchPastBroadcastsIE,
										TwitchHighlightsIE,
										TwitchClipsIE,
									))
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										ϒie = τmp1
										λgy.Yield(λ.Calm(ϒie, "suitable", ϒurl))
									}
									return λ.None
								})
							})))) {
							return λ.False
						} else {
							return λ.Calm(λ.Cal(λ.SuperType, TwitchStreamIE, ϒcls), "suitable", ϒurl)
						}
					}()
				})
			TwitchStreamIE_suitable = λ.Cal(λ.ClassMethodType, TwitchStreamIE_suitable)
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": TwitchStreamIE__VALID_URL,
				"suitable":   TwitchStreamIE_suitable,
			})
		}())
		TwitchClipsIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitchClipsIE"), λ.NewTuple(TwitchBaseIE), func() λ.Dict {
			var (
				TwitchClipsIE_IE_NAME       λ.Object
				TwitchClipsIE__VALID_URL    λ.Object
				TwitchClipsIE__real_extract λ.Object
			)
			TwitchClipsIE_IE_NAME = λ.StrLiteral("twitch:clips")
			TwitchClipsIE__VALID_URL = λ.StrLiteral("(?x)\n                    https?://\n                        (?:\n                            clips\\.twitch\\.tv/(?:embed\\?.*?\\bclip=|(?:[^/]+/)*)|\n                            (?:(?:www|go|m)\\.)?twitch\\.tv/[^/]+/clip/\n                        )\n                        (?P<id>[^/?#&]+)\n                    ")
			TwitchClipsIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒclip          λ.Object
						ϒformats       λ.Object
						ϒmobj          λ.Object
						ϒoption        λ.Object
						ϒself          = λargs[0]
						ϒsource        λ.Object
						ϒthumb         λ.Object
						ϒthumbnail_id  λ.Object
						ϒthumbnail_url λ.Object
						ϒthumbnails    λ.Object
						ϒurl           = λargs[1]
						ϒvideo_id      λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒclip = λ.GetItem(λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.StrLiteral("https://gql.twitch.tv/gql"),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "data", Value: λ.Calm(λ.Cal(Ωjson.ϒdumps, λ.DictLiteral(map[string]λ.Object{
							"query": λ.Mod(λ.StrLiteral("{\n  clip(slug: \"%s\") {\n    broadcaster {\n      displayName\n    }\n    createdAt\n    curator {\n      displayName\n      id\n    }\n    durationSeconds\n    id\n    tiny: thumbnailURL(width: 86, height: 45)\n    small: thumbnailURL(width: 260, height: 147)\n    medium: thumbnailURL(width: 480, height: 272)\n    title\n    videoQualities {\n      frameRate\n      quality\n      sourceURL\n    }\n    viewCount\n  }\n}"), ϒvideo_id),
						})), "encode")},
						{Name: "headers", Value: λ.DictLiteral(map[string]λ.Object{
							"Client-ID": λ.GetAttr(ϒself, "_CLIENT_ID", nil),
						})},
					}), λ.StrLiteral("data")), λ.StrLiteral("clip"))
					if !λ.IsTrue(ϒclip) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.StrLiteral("This clip is no longer available")), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒclip, "get", λ.StrLiteral("videoQualities"), λ.NewList()))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒoption = τmp1
						if !λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒoption, λ.DictType)) {
							continue
						}
						ϒsource = λ.Cal(ϒurl_or_none, λ.Calm(ϒoption, "get", λ.StrLiteral("sourceURL")))
						if !λ.IsTrue(ϒsource) {
							continue
						}
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"url":       ϒsource,
							"format_id": λ.Calm(ϒoption, "get", λ.StrLiteral("quality")),
							"height":    λ.Cal(ϒint_or_none, λ.Calm(ϒoption, "get", λ.StrLiteral("quality"))),
							"fps":       λ.Cal(ϒint_or_none, λ.Calm(ϒoption, "get", λ.StrLiteral("frameRate"))),
						}))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒthumbnails = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
						λ.StrLiteral("tiny"),
						λ.StrLiteral("small"),
						λ.StrLiteral("medium"),
					))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒthumbnail_id = τmp1
						ϒthumbnail_url = λ.Calm(ϒclip, "get", ϒthumbnail_id)
						if !λ.IsTrue(ϒthumbnail_url) {
							continue
						}
						ϒthumb = λ.DictLiteral(map[string]λ.Object{
							"id":  ϒthumbnail_id,
							"url": ϒthumbnail_url,
						})
						ϒmobj = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("-(\\d+)x(\\d+)\\."), ϒthumbnail_url)
						if λ.IsTrue(ϒmobj) {
							λ.Calm(ϒthumb, "update", λ.DictLiteral(map[string]λ.Object{
								"height": λ.Cal(λ.IntType, λ.Calm(ϒmobj, "group", λ.IntLiteral(2))),
								"width":  λ.Cal(λ.IntType, λ.Calm(ϒmobj, "group", λ.IntLiteral(1))),
							}))
						}
						λ.Calm(ϒthumbnails, "append", ϒthumb)
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id": func() λ.Object {
							if λv := λ.Calm(ϒclip, "get", λ.StrLiteral("id")); λ.IsTrue(λv) {
								return λv
							} else {
								return ϒvideo_id
							}
						}(),
						"title": func() λ.Object {
							if λv := λ.Calm(ϒclip, "get", λ.StrLiteral("title")); λ.IsTrue(λv) {
								return λv
							} else {
								return ϒvideo_id
							}
						}(),
						"formats":    ϒformats,
						"duration":   λ.Cal(ϒint_or_none, λ.Calm(ϒclip, "get", λ.StrLiteral("durationSeconds"))),
						"views":      λ.Cal(ϒint_or_none, λ.Calm(ϒclip, "get", λ.StrLiteral("viewCount"))),
						"timestamp":  λ.Cal(ϒunified_timestamp, λ.Calm(ϒclip, "get", λ.StrLiteral("createdAt"))),
						"thumbnails": ϒthumbnails,
						"creator": λ.Cal(ϒtry_get, ϒclip, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("broadcaster")), λ.StrLiteral("displayName"))
							}), ϒcompat_str),
						"uploader": λ.Cal(ϒtry_get, ϒclip, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("curator")), λ.StrLiteral("displayName"))
							}), ϒcompat_str),
						"uploader_id": λ.Cal(ϒtry_get, ϒclip, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("curator")), λ.StrLiteral("id"))
							}), ϒcompat_str),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":       TwitchClipsIE_IE_NAME,
				"_VALID_URL":    TwitchClipsIE__VALID_URL,
				"_real_extract": TwitchClipsIE__real_extract,
			})
		}())
	})
}
