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
 * generic/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/generic.py
 */

package generic

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωanvato "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/anvato"
	Ωapa "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/apa"
	Ωarkena "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/arkena"
	Ωbrightcove "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/brightcove"
	Ωchannel9 "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/channel9"
	Ωcloudflarestream "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/cloudflarestream"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωcommonprotocols "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/commonprotocols"
	Ωcondenast "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/condenast"
	Ωdailymail "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/dailymail"
	Ωdailymotion "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/dailymotion"
	Ωdbtv "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/dbtv"
	Ωdigiteka "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/digiteka"
	Ωdrtuber "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/drtuber"
	Ωeagleplatform "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/eagleplatform"
	Ωexpressen "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/expressen"
	Ωfacebook "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/facebook"
	Ωfoxnews "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/foxnews"
	Ωgoogledrive "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/googledrive"
	Ωindavideo "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/indavideo"
	Ωinstagram "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/instagram"
	Ωjoj "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/joj"
	Ωjwplatform "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/jwplatform"
	Ωkaltura "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/kaltura"
	Ωlimelight "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/limelight"
	Ωliveleak "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/liveleak"
	Ωmediaset "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/mediaset"
	Ωmediasite "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/mediasite"
	Ωmegaphone "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/megaphone"
	Ωmtv "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/mtv"
	Ωmyvi "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/myvi"
	Ωnbc "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/nbc"
	Ωnexx "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/nexx"
	Ωonionstudios "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/onionstudios"
	Ωooyala "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/ooyala"
	Ωopenload "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/openload"
	Ωpeertube "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/peertube"
	Ωpiksel "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/piksel"
	Ωpladform "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/pladform"
	Ωpornhub "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/pornhub"
	Ωredtube "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/redtube"
	Ωrutube "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/rutube"
	Ωrutv "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/rutv"
	Ωsenateisvp "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/senateisvp"
	Ωsmotri "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/smotri"
	Ωsoundcloud "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/soundcloud"
	Ωsportbox "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/sportbox"
	Ωspringboardplatform "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/springboardplatform"
	Ωsvt "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/svt"
	Ωtheplatform "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/theplatform"
	Ωthreeqsdn "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/threeqsdn"
	Ωtnaflix "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/tnaflix"
	Ωtube8 "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/tube8"
	Ωtunein "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/tunein"
	Ωtvc "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/tvc"
	Ωtwentymin "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/twentymin"
	Ωudn "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/udn"
	Ωustream "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/ustream"
	Ωvbox7 "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/vbox7"
	Ωvessel "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/vessel"
	Ωvice "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/vice"
	Ωvidea "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/videa"
	Ωvideomore "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/videomore"
	Ωvideopress "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/videopress"
	Ωviewlift "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/viewlift"
	Ωvimeo "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/vimeo"
	Ωviqeo "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/viqeo"
	Ωvshare "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/vshare"
	Ωvzaar "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/vzaar"
	Ωwashingtonpost "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/washingtonpost"
	Ωwebcaster "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/webcaster"
	Ωwistia "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/wistia"
	Ωxfileshare "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/xfileshare"
	Ωxhamster "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/xhamster"
	Ωyapfiles "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/yapfiles"
	Ωyoutube "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/youtube"
	Ωzype "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/zype"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	APAIE                        λ.Object
	AnvatoIE                     λ.Object
	ArkenaIE                     λ.Object
	BrightcoveLegacyIE           λ.Object
	BrightcoveNewIE              λ.Object
	Channel9IE                   λ.Object
	CloudflareStreamIE           λ.Object
	CondeNastIE                  λ.Object
	DBTVIE                       λ.Object
	DailyMailIE                  λ.Object
	DailymotionIE                λ.Object
	DigitekaIE                   λ.Object
	DrTuberIE                    λ.Object
	EaglePlatformIE              λ.Object
	ExpressenIE                  λ.Object
	ExtractorError               λ.Object
	FacebookIE                   λ.Object
	FoxNewsIE                    λ.Object
	GoogleDriveIE                λ.Object
	IndavideoEmbedIE             λ.Object
	InfoExtractor                λ.Object
	InstagramIE                  λ.Object
	JWPlatformIE                 λ.Object
	JojIE                        λ.Object
	KNOWN_EXTENSIONS             λ.Object
	KalturaIE                    λ.Object
	LimelightBaseIE              λ.Object
	LiveLeakIE                   λ.Object
	MTVServicesEmbeddedIE        λ.Object
	MediasetIE                   λ.Object
	MediasiteIE                  λ.Object
	MegaphoneIE                  λ.Object
	MyviIE                       λ.Object
	NBCSportsVPlayerIE           λ.Object
	NexxEmbedIE                  λ.Object
	NexxIE                       λ.Object
	OnionStudiosIE               λ.Object
	OoyalaIE                     λ.Object
	OpenloadIE                   λ.Object
	PeerTubeIE                   λ.Object
	PikselIE                     λ.Object
	PladformIE                   λ.Object
	PornHubIE                    λ.Object
	RUTVIE                       λ.Object
	RedTubeIE                    λ.Object
	RtmpIE                       λ.Object
	RutubeIE                     λ.Object
	SVTIE                        λ.Object
	SenateISVPIE                 λ.Object
	SmotriIE                     λ.Object
	SoundcloudIE                 λ.Object
	SportBoxIE                   λ.Object
	SpringboardPlatformIE        λ.Object
	TNAFlixNetworkEmbedIE        λ.Object
	TVCIE                        λ.Object
	ThePlatformIE                λ.Object
	ThreeQSDNIE                  λ.Object
	Tube8IE                      λ.Object
	TuneInBaseIE                 λ.Object
	TwentyMinutenIE              λ.Object
	UDNEmbedIE                   λ.Object
	UstreamIE                    λ.Object
	VShareIE                     λ.Object
	Vbox7IE                      λ.Object
	VesselIE                     λ.Object
	ViceIE                       λ.Object
	VideaIE                      λ.Object
	VideoPressIE                 λ.Object
	VideomoreIE                  λ.Object
	ViewLiftEmbedIE              λ.Object
	VimeoIE                      λ.Object
	ViqeoIE                      λ.Object
	VzaarIE                      λ.Object
	WashingtonPostIE             λ.Object
	WebcasterFeedIE              λ.Object
	WistiaIE                     λ.Object
	XFileShareIE                 λ.Object
	XHamsterEmbedIE              λ.Object
	YapFilesIE                   λ.Object
	YoutubeIE                    λ.Object
	ZypeIE                       λ.Object
	ϒcompat_etree_fromstring     λ.Object
	ϒcompat_str                  λ.Object
	ϒcompat_urllib_parse_unquote λ.Object
	ϒcompat_xml_parse_error      λ.Object
	ϒdetermine_ext               λ.Object
	ϒfloat_or_none               λ.Object
	ϒjs_to_json                  λ.Object
	ϒmerge_dicts                 λ.Object
	ϒmimetype2ext                λ.Object
	ϒorderedSet                  λ.Object
	ϒsanitized_Request           λ.Object
	ϒsmuggle_url                 λ.Object
	ϒunescapeHTML                λ.Object
	ϒunified_strdate             λ.Object
	ϒunsmuggle_url               λ.Object
	ϒxpath_text                  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		YoutubeIE = Ωyoutube.YoutubeIE
		ϒcompat_etree_fromstring = Ωcompat.ϒcompat_etree_fromstring
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒcompat_urllib_parse_unquote = Ωcompat.ϒcompat_urllib_parse_unquote
		ϒcompat_xml_parse_error = Ωcompat.ϒcompat_xml_parse_error
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒjs_to_json = Ωutils.ϒjs_to_json
		KNOWN_EXTENSIONS = Ωutils.KNOWN_EXTENSIONS
		ϒmerge_dicts = Ωutils.ϒmerge_dicts
		ϒmimetype2ext = Ωutils.ϒmimetype2ext
		ϒorderedSet = Ωutils.ϒorderedSet
		ϒsanitized_Request = Ωutils.ϒsanitized_Request
		ϒsmuggle_url = Ωutils.ϒsmuggle_url
		ϒunescapeHTML = Ωutils.ϒunescapeHTML
		ϒunified_strdate = Ωutils.ϒunified_strdate
		ϒunsmuggle_url = Ωutils.ϒunsmuggle_url
		ϒxpath_text = Ωutils.ϒxpath_text
		RtmpIE = Ωcommonprotocols.RtmpIE
		BrightcoveLegacyIE = Ωbrightcove.BrightcoveLegacyIE
		BrightcoveNewIE = Ωbrightcove.BrightcoveNewIE
		NexxIE = Ωnexx.NexxIE
		NexxEmbedIE = Ωnexx.NexxEmbedIE
		NBCSportsVPlayerIE = Ωnbc.NBCSportsVPlayerIE
		OoyalaIE = Ωooyala.OoyalaIE
		RUTVIE = Ωrutv.RUTVIE
		TVCIE = Ωtvc.TVCIE
		SportBoxIE = Ωsportbox.SportBoxIE
		SmotriIE = Ωsmotri.SmotriIE
		MyviIE = Ωmyvi.MyviIE
		CondeNastIE = Ωcondenast.CondeNastIE
		UDNEmbedIE = Ωudn.UDNEmbedIE
		SenateISVPIE = Ωsenateisvp.SenateISVPIE
		SVTIE = Ωsvt.SVTIE
		PornHubIE = Ωpornhub.PornHubIE
		XHamsterEmbedIE = Ωxhamster.XHamsterEmbedIE
		TNAFlixNetworkEmbedIE = Ωtnaflix.TNAFlixNetworkEmbedIE
		DrTuberIE = Ωdrtuber.DrTuberIE
		RedTubeIE = Ωredtube.RedTubeIE
		Tube8IE = Ωtube8.Tube8IE
		VimeoIE = Ωvimeo.VimeoIE
		DailymotionIE = Ωdailymotion.DailymotionIE
		DailyMailIE = Ωdailymail.DailyMailIE
		OnionStudiosIE = Ωonionstudios.OnionStudiosIE
		ViewLiftEmbedIE = Ωviewlift.ViewLiftEmbedIE
		MTVServicesEmbeddedIE = Ωmtv.MTVServicesEmbeddedIE
		PladformIE = Ωpladform.PladformIE
		VideomoreIE = Ωvideomore.VideomoreIE
		WebcasterFeedIE = Ωwebcaster.WebcasterFeedIE
		GoogleDriveIE = Ωgoogledrive.GoogleDriveIE
		JWPlatformIE = Ωjwplatform.JWPlatformIE
		DigitekaIE = Ωdigiteka.DigitekaIE
		ArkenaIE = Ωarkena.ArkenaIE
		InstagramIE = Ωinstagram.InstagramIE
		LiveLeakIE = Ωliveleak.LiveLeakIE
		ThreeQSDNIE = Ωthreeqsdn.ThreeQSDNIE
		ThePlatformIE = Ωtheplatform.ThePlatformIE
		VesselIE = Ωvessel.VesselIE
		KalturaIE = Ωkaltura.KalturaIE
		EaglePlatformIE = Ωeagleplatform.EaglePlatformIE
		FacebookIE = Ωfacebook.FacebookIE
		SoundcloudIE = Ωsoundcloud.SoundcloudIE
		TuneInBaseIE = Ωtunein.TuneInBaseIE
		Vbox7IE = Ωvbox7.Vbox7IE
		DBTVIE = Ωdbtv.DBTVIE
		PikselIE = Ωpiksel.PikselIE
		VideaIE = Ωvidea.VideaIE
		TwentyMinutenIE = Ωtwentymin.TwentyMinutenIE
		UstreamIE = Ωustream.UstreamIE
		OpenloadIE = Ωopenload.OpenloadIE
		VideoPressIE = Ωvideopress.VideoPressIE
		RutubeIE = Ωrutube.RutubeIE
		LimelightBaseIE = Ωlimelight.LimelightBaseIE
		AnvatoIE = Ωanvato.AnvatoIE
		WashingtonPostIE = Ωwashingtonpost.WashingtonPostIE
		WistiaIE = Ωwistia.WistiaIE
		MediasetIE = Ωmediaset.MediasetIE
		JojIE = Ωjoj.JojIE
		MegaphoneIE = Ωmegaphone.MegaphoneIE
		VzaarIE = Ωvzaar.VzaarIE
		Channel9IE = Ωchannel9.Channel9IE
		VShareIE = Ωvshare.VShareIE
		MediasiteIE = Ωmediasite.MediasiteIE
		SpringboardPlatformIE = Ωspringboardplatform.SpringboardPlatformIE
		YapFilesIE = Ωyapfiles.YapFilesIE
		ViceIE = Ωvice.ViceIE
		XFileShareIE = Ωxfileshare.XFileShareIE
		CloudflareStreamIE = Ωcloudflarestream.CloudflareStreamIE
		PeerTubeIE = Ωpeertube.PeerTubeIE
		IndavideoEmbedIE = Ωindavideo.IndavideoEmbedIE
		APAIE = Ωapa.APAIE
		FoxNewsIE = Ωfoxnews.FoxNewsIE
		ViqeoIE = Ωviqeo.ViqeoIE
		ExpressenIE = Ωexpressen.ExpressenIE
		ZypeIE = Ωzype.ZypeIE
	})
}
