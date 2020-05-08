// Code generated by transpiler. DO NOT EDIT.
// +build !restricted

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
 * youtube/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/youtube.py
 */

package youtube

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωtime "github.com/tenta-browser/go-video-downloader/gen/time"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError                    λ.Object
	InfoExtractor                     λ.Object
	SearchInfoExtractor               λ.Object
	YoutubeBaseInfoExtractor          λ.Object
	YoutubeChannelIE                  λ.Object
	YoutubeEntryListBaseInfoExtractor λ.Object
	YoutubeFavouritesIE               λ.Object
	YoutubeFeedsInfoExtractor         λ.Object
	YoutubeHistoryIE                  λ.Object
	YoutubeIE                         λ.Object
	YoutubeLiveIE                     λ.Object
	YoutubePlaylistBaseInfoExtractor  λ.Object
	YoutubePlaylistIE                 λ.Object
	YoutubePlaylistsBaseInfoExtractor λ.Object
	YoutubePlaylistsIE                λ.Object
	YoutubeRecommendedIE              λ.Object
	YoutubeSearchBaseInfoExtractor    λ.Object
	YoutubeSearchIE                   λ.Object
	YoutubeSearchURLIE                λ.Object
	YoutubeShowIE                     λ.Object
	YoutubeSubscriptionsIE            λ.Object
	YoutubeTruncatedIDIE              λ.Object
	YoutubeTruncatedURLIE             λ.Object
	YoutubeUserIE                     λ.Object
	YoutubeWatchLaterIE               λ.Object
	ϒclean_html                       λ.Object
	ϒcompat_HTTPError                 λ.Object
	ϒcompat_chr                       λ.Object
	ϒcompat_kwargs                    λ.Object
	ϒcompat_parse_qs                  λ.Object
	ϒcompat_str                       λ.Object
	ϒcompat_urllib_parse_unquote      λ.Object
	ϒcompat_urllib_parse_unquote_plus λ.Object
	ϒcompat_urllib_parse_urlencode    λ.Object
	ϒcompat_urllib_parse_urlparse     λ.Object
	ϒerror_to_compat_str              λ.Object
	ϒextract_attributes               λ.Object
	ϒfloat_or_none                    λ.Object
	ϒget_element_by_attribute         λ.Object
	ϒget_element_by_id                λ.Object
	ϒint_or_none                      λ.Object
	ϒmimetype2ext                     λ.Object
	ϒparse_codecs                     λ.Object
	ϒparse_duration                   λ.Object
	ϒremove_start                     λ.Object
	ϒsmuggle_url                      λ.Object
	ϒstr_or_none                      λ.Object
	ϒstr_to_int                       λ.Object
	ϒtry_get                          λ.Object
	ϒunescapeHTML                     λ.Object
	ϒunified_strdate                  λ.Object
	ϒunsmuggle_url                    λ.Object
	ϒurl_or_none                      λ.Object
	ϒurlencode_postdata               λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		SearchInfoExtractor = Ωcommon.SearchInfoExtractor
		ϒcompat_chr = Ωcompat.ϒcompat_chr
		ϒcompat_HTTPError = Ωcompat.ϒcompat_HTTPError
		ϒcompat_kwargs = Ωcompat.ϒcompat_kwargs
		ϒcompat_parse_qs = Ωcompat.ϒcompat_parse_qs
		ϒcompat_urllib_parse_unquote = Ωcompat.ϒcompat_urllib_parse_unquote
		ϒcompat_urllib_parse_unquote_plus = Ωcompat.ϒcompat_urllib_parse_unquote_plus
		ϒcompat_urllib_parse_urlencode = Ωcompat.ϒcompat_urllib_parse_urlencode
		ϒcompat_urllib_parse_urlparse = Ωcompat.ϒcompat_urllib_parse_urlparse
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒclean_html = Ωutils.ϒclean_html
		ϒerror_to_compat_str = Ωutils.ϒerror_to_compat_str
		ϒextract_attributes = Ωutils.ϒextract_attributes
		ExtractorError = Ωutils.ExtractorError
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒget_element_by_attribute = Ωutils.ϒget_element_by_attribute
		ϒget_element_by_id = Ωutils.ϒget_element_by_id
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒmimetype2ext = Ωutils.ϒmimetype2ext
		ϒparse_codecs = Ωutils.ϒparse_codecs
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒremove_start = Ωutils.ϒremove_start
		ϒsmuggle_url = Ωutils.ϒsmuggle_url
		ϒstr_or_none = Ωutils.ϒstr_or_none
		ϒstr_to_int = Ωutils.ϒstr_to_int
		ϒtry_get = Ωutils.ϒtry_get
		ϒunescapeHTML = Ωutils.ϒunescapeHTML
		ϒunified_strdate = Ωutils.ϒunified_strdate
		ϒunsmuggle_url = Ωutils.ϒunsmuggle_url
		ϒurl_or_none = Ωutils.ϒurl_or_none
		ϒurlencode_postdata = Ωutils.ϒurlencode_postdata
		YoutubeBaseInfoExtractor = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubeBaseInfoExtractor"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				YoutubeBaseInfoExtractor__PLAYLIST_ID_RE  λ.Object
				YoutubeBaseInfoExtractor__real_initialize λ.Object
				YoutubeBaseInfoExtractor__set_language    λ.Object
			)
			YoutubeBaseInfoExtractor__PLAYLIST_ID_RE = λ.StrLiteral("(?:PL|LL|EC|UU|FL|RD|UL|TL|PU|OLAK5uy_)[0-9A-Za-z-_]{10,}")
			YoutubeBaseInfoExtractor__set_language = λ.NewFunction("_set_language",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
					)
					λ.Call(λ.GetAttr(ϒself, "_set_cookie", nil), λ.NewArgs(
						λ.StrLiteral(".youtube.com"),
						λ.StrLiteral("PREF"),
						λ.StrLiteral("f1=50000000&hl=en"),
					), λ.KWArgs{
						{Name: "expire_time", Value: λ.Add(λ.Cal(Ωtime.ϒtime), λ.Mul(λ.Mul(λ.Mul(λ.IntLiteral(2), λ.IntLiteral(30)), λ.IntLiteral(24)), λ.IntLiteral(3600)))},
					})
					return λ.None
				})
			YoutubeBaseInfoExtractor__real_initialize = λ.NewFunction("_real_initialize",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
					)
					if λ.GetAttr(ϒself, "_downloader", nil) == λ.None {
						return λ.None
					}
					λ.Calm(ϒself, "_set_language")
					if !λ.IsTrue(λ.Calm(ϒself, "_login")) {
						return λ.None
					}
					return λ.None
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_PLAYLIST_ID_RE":  YoutubeBaseInfoExtractor__PLAYLIST_ID_RE,
				"_real_initialize": YoutubeBaseInfoExtractor__real_initialize,
				"_set_language":    YoutubeBaseInfoExtractor__set_language,
			})
		}())
		YoutubeEntryListBaseInfoExtractor = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubeEntryListBaseInfoExtractor"), λ.NewTuple(YoutubeBaseInfoExtractor), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		YoutubePlaylistBaseInfoExtractor = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubePlaylistBaseInfoExtractor"), λ.NewTuple(YoutubeEntryListBaseInfoExtractor), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		YoutubePlaylistsBaseInfoExtractor = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubePlaylistsBaseInfoExtractor"), λ.NewTuple(YoutubeEntryListBaseInfoExtractor), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		YoutubeIE = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubeIE"), λ.NewTuple(YoutubeBaseInfoExtractor), func() λ.Dict {
			var (
				YoutubeIE__GEO_BYPASS   λ.Object
				YoutubeIE__VALID_URL    λ.Object
				YoutubeIE___init__      λ.Object
				YoutubeIE__extract_url  λ.Object
				YoutubeIE__extract_urls λ.Object
			)
			YoutubeIE__VALID_URL = λ.Mod(λ.StrLiteral("(?x)^\n                     (\n                         (?:https?://|//)                                    # http(s):// or protocol-independent URL\n                         (?:(?:(?:(?:\\w+\\.)?[yY][oO][uU][tT][uU][bB][eE](?:-nocookie|kids)?\\.com/|\n                            (?:www\\.)?deturl\\.com/www\\.youtube\\.com/|\n                            (?:www\\.)?pwnyoutube\\.com/|\n                            (?:www\\.)?hooktube\\.com/|\n                            (?:www\\.)?yourepeat\\.com/|\n                            tube\\.majestyc\\.net/|\n                            # Invidious instances taken from https://github.com/omarroth/invidious/wiki/Invidious-Instances\n                            (?:(?:www|dev)\\.)?invidio\\.us/|\n                            (?:(?:www|no)\\.)?invidiou\\.sh/|\n                            (?:(?:www|fi|de)\\.)?invidious\\.snopyta\\.org/|\n                            (?:www\\.)?invidious\\.kabi\\.tk/|\n                            (?:www\\.)?invidious\\.13ad\\.de/|\n                            (?:www\\.)?invidious\\.mastodon\\.host/|\n                            (?:www\\.)?invidious\\.nixnet\\.xyz/|\n                            (?:www\\.)?invidious\\.drycat\\.fr/|\n                            (?:www\\.)?tube\\.poal\\.co/|\n                            (?:www\\.)?vid\\.wxzm\\.sx/|\n                            (?:www\\.)?yt\\.elukerio\\.org/|\n                            (?:www\\.)?yt\\.lelux\\.fi/|\n                            (?:www\\.)?kgg2m7yk5aybusll\\.onion/|\n                            (?:www\\.)?qklhadlycap4cnod\\.onion/|\n                            (?:www\\.)?axqzx4s6s54s32yentfqojs3x5i7faxza6xo3ehd4bzzsg2ii4fv2iid\\.onion/|\n                            (?:www\\.)?c7hqkpkpemu6e7emz5b4vyz7idjgdvgaaa3dyimmeojqbgpea3xqjoid\\.onion/|\n                            (?:www\\.)?fz253lmuao3strwbfbmx46yu7acac2jz27iwtorgmbqlkurlclmancad\\.onion/|\n                            (?:www\\.)?invidious\\.l4qlywnpwqsluw65ts7md3khrivpirse744un3x7mlskqauz5pyuzgqd\\.onion/|\n                            (?:www\\.)?owxfohz4kjyv25fvlqilyxast7inivgiktls3th44jhk3ej3i7ya\\.b32\\.i2p/|\n                            youtube\\.googleapis\\.com/)                        # the various hostnames, with wildcard subdomains\n                         (?:.*?\\#/)?                                          # handle anchor (#/) redirect urls\n                         (?:                                                  # the various things that can precede the ID:\n                             (?:(?:v|embed|e)/(?!videoseries))                # v/ or embed/ or e/\n                             |(?:                                             # or the v= param in all its forms\n                                 (?:(?:watch|movie)(?:_popup)?(?:\\.php)?/?)?  # preceding watch(_popup|.php) or nothing (like /?v=xxxx)\n                                 (?:\\?|\\#!?)                                  # the params delimiter ? or # or #!\n                                 (?:.*?[&;])??                                # any other preceding param (like /?s=tuff&v=xxxx or ?s=tuff&amp;v=V36LpHqtcDY)\n                                 v=\n                             )\n                         ))\n                         |(?:\n                            youtu\\.be|                                        # just youtu.be/xxxx\n                            vid\\.plus|                                        # or vid.plus/xxxx\n                            zwearz\\.com/watch|                                # or zwearz.com/watch/xxxx\n                         )/\n                         |(?:www\\.)?cleanvideosearch\\.com/media/action/yt/watch\\?videoId=\n                         )\n                     )?                                                       # all until now is optional -> you can pass the naked ID\n                     ([0-9A-Za-z_-]{11})                                      # here is it! the YouTube video ID\n                     (?!.*?\\blist=\n                        (?:\n                            %(playlist_id)s|                                  # combined list/video URLs are handled by the playlist IE\n                            WL                                                # WL are handled by the watch later IE\n                        )\n                     )\n                     #(?(1).+)?                                                # if we found the ID, everything can follow\n                     "), λ.DictLiteral(map[string]λ.Object{
				"playlist_id": λ.GetAttr(YoutubeBaseInfoExtractor, "_PLAYLIST_ID_RE", nil),
			}))
			YoutubeIE__GEO_BYPASS = λ.False
			YoutubeIE___init__ = λ.NewFunction("__init__",
				[]λ.Param{
					{Name: "self"},
				},
				0, true, true,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒargs   = λargs[1]
						ϒkwargs = λargs[2]
						ϒself   = λargs[0]
					)
					λ.Call(λ.GetAttr(λ.Cal(λ.SuperType, YoutubeIE, ϒself), "__init__", nil), λ.NewArgs(λ.Unpack(λ.AsStarred(ϒargs))...), λ.KWArgs{
						{Name: "", Value: ϒkwargs},
					})
					λ.SetAttr(ϒself, "_player_cache", λ.DictLiteral(map[λ.Object]λ.Object{}))
					return λ.None
				})
			YoutubeIE__extract_urls = λ.NewFunction("_extract_urls",
				[]λ.Param{
					{Name: "webpage"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒentries λ.Object
						ϒmatches λ.Object
						ϒwebpage = λargs[0]
					)
					ϒentries = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒmobj λ.Object
									τmp0  λ.Object
									τmp1  λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfinditer, λ.StrLiteral("(?x)\n            (?:\n                <iframe[^>]+?src=|\n                data-video-url=|\n                <embed[^>]+?src=|\n                embedSWF\\(?:\\s*|\n                <object[^>]+data=|\n                new\\s+SWFObject\\(\n            )\n            ([\"\\'])\n                (?P<url>(?:https?:)?//(?:www\\.)?youtube(?:-nocookie)?\\.com/\n                (?:embed|v|p)/[0-9A-Za-z_-]{11}.*?)\n            \\1"), ϒwebpage))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒmobj = τmp1
									λgy.Yield(λ.Cal(ϒunescapeHTML, λ.Calm(ϒmobj, "group", λ.StrLiteral("url"))))
								}
								return λ.None
							})
						})))
					λ.Calm(ϒentries, "extend", λ.Cal(λ.ListType, λ.Cal(λ.MapIteratorType, ϒunescapeHTML, λ.Cal(Ωre.ϒfindall, λ.StrLiteral("class=\"lazyYT\" data-youtube-id=\"([^\"]+)\""), ϒwebpage))))
					ϒmatches = λ.Cal(Ωre.ϒfindall, λ.StrLiteral("(?x)<div[^>]+\n            class=(?P<q1>[\\'\"])[^\\'\"]*\\byvii_single_video_player\\b[^\\'\"]*(?P=q1)[^>]+\n            data-video_id=(?P<q2>[\\'\"])([^\\'\"]+)(?P=q2)"), ϒwebpage)
					λ.Calm(ϒentries, "extend", λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒm   λ.Object
									τmp0 λ.Object
									τmp1 λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, ϒmatches)
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒm = τmp1
									λgy.Yield(λ.GetItem(ϒm, λ.Neg(λ.IntLiteral(1))))
								}
								return λ.None
							})
						})))
					return ϒentries
				})
			YoutubeIE__extract_urls = λ.Cal(λ.StaticMethodType, YoutubeIE__extract_urls)
			YoutubeIE__extract_url = λ.NewFunction("_extract_url",
				[]λ.Param{
					{Name: "webpage"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒurls    λ.Object
						ϒwebpage = λargs[0]
					)
					ϒurls = λ.Calm(YoutubeIE, "_extract_urls", ϒwebpage)
					return func() λ.Object {
						if λ.IsTrue(ϒurls) {
							return λ.GetItem(ϒurls, λ.IntLiteral(0))
						} else {
							return λ.None
						}
					}()
				})
			YoutubeIE__extract_url = λ.Cal(λ.StaticMethodType, YoutubeIE__extract_url)
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_GEO_BYPASS":   YoutubeIE__GEO_BYPASS,
				"_VALID_URL":    YoutubeIE__VALID_URL,
				"__init__":      YoutubeIE___init__,
				"_extract_url":  YoutubeIE__extract_url,
				"_extract_urls": YoutubeIE__extract_urls,
			})
		}())
		YoutubePlaylistIE = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubePlaylistIE"), λ.NewTuple(YoutubePlaylistBaseInfoExtractor), func() λ.Dict {
			var (
				YoutubePlaylistIE__VIDEO_RE_TPL λ.Object
			)
			YoutubePlaylistIE__VIDEO_RE_TPL = λ.StrLiteral("href=\"\\s*/watch\\?v=%s(?:&amp;(?:[^\"]*?index=(?P<index>\\d+))?(?:[^>]+>(?P<title>[^<]+))?)?")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VIDEO_RE_TPL": YoutubePlaylistIE__VIDEO_RE_TPL,
			})
		}())
		YoutubeChannelIE = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubeChannelIE"), λ.NewTuple(YoutubePlaylistBaseInfoExtractor), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		YoutubeUserIE = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubeUserIE"), λ.NewTuple(YoutubeChannelIE), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		YoutubeLiveIE = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubeLiveIE"), λ.NewTuple(YoutubeBaseInfoExtractor), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		YoutubePlaylistsIE = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubePlaylistsIE"), λ.NewTuple(YoutubePlaylistsBaseInfoExtractor), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		YoutubeSearchBaseInfoExtractor = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubeSearchBaseInfoExtractor"), λ.NewTuple(YoutubePlaylistBaseInfoExtractor), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		YoutubeSearchIE = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubeSearchIE"), λ.NewTuple(
			SearchInfoExtractor,
			YoutubeSearchBaseInfoExtractor,
		), func() λ.Dict {
			var (
				YoutubeSearchIE_IE_NAME λ.Object
			)
			YoutubeSearchIE_IE_NAME = λ.StrLiteral("youtube:search")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME": YoutubeSearchIE_IE_NAME,
			})
		}())
		YoutubeSearchURLIE = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubeSearchURLIE"), λ.NewTuple(YoutubeSearchBaseInfoExtractor), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		YoutubeShowIE = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubeShowIE"), λ.NewTuple(YoutubePlaylistsBaseInfoExtractor), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		YoutubeFeedsInfoExtractor = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubeFeedsInfoExtractor"), λ.NewTuple(YoutubeBaseInfoExtractor), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		YoutubeWatchLaterIE = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubeWatchLaterIE"), λ.NewTuple(YoutubePlaylistIE), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		YoutubeFavouritesIE = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubeFavouritesIE"), λ.NewTuple(YoutubeBaseInfoExtractor), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		YoutubeRecommendedIE = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubeRecommendedIE"), λ.NewTuple(YoutubeFeedsInfoExtractor), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		YoutubeSubscriptionsIE = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubeSubscriptionsIE"), λ.NewTuple(YoutubeFeedsInfoExtractor), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		YoutubeHistoryIE = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubeHistoryIE"), λ.NewTuple(YoutubeFeedsInfoExtractor), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		YoutubeTruncatedURLIE = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubeTruncatedURLIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		YoutubeTruncatedIDIE = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubeTruncatedIDIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
	})
}
