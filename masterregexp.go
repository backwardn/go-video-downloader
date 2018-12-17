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
 * masterregexp.go: Automatically generated regexp by combining URL validation expressions of all extractors
 */

// Code generated by transpiler. DO NOT EDIT.

package downloader

func init() {
	masterRegexps = []string{
		`(?:https?://abcnews\.go\.com/(?:[^/]+/video/(?:[0-9a-z-]+)-|video/embed\?.*?\bid` +
			`=)(?:\d+))|(?:https?://(?:(?:(?:embed|www)\.)?acast\.com/|play\.acast\.com/s/)(?` +
			`:[^/]+)/(?:[^/#?]+))|(?:https?://tv\.adobe\.com/(?:(?:fr|de|es|jp)/)?watch/(?:[^` +
			`/]+)/(?:[^/]+))|(?:https?://video\.tv\.adobe\.com/v/(?:\d+))|(?:(?:aol-video:|ht` +
			`tps?://(?:(?:www|on)\.)?aol\.com/(?:[^/]+/)*(?:[^/?#&]+-)?)(?:[^/?#&]+))|(?:http` +
			`s?://(?:www\.)?allocine\.fr/(?:article|video|film)/(?:fichearticle_gen_carticle=` +
			`|player_gen_cmedia=|fichefilm_gen_cfilm=|video-)(?:[0-9]+)(?:\.html)?)|(?:https?` +
			`://(?:www\.)?aparat\.com/(?:v/|video/video/embed/videohash/)(?:[a-zA-Z0-9]+))|(?` +
			`:(?:https?://(www\.)?daserste\.de/[^?#]+/videos/(?:[^/?#]+)-(?:[0-9]+))\.html)|(` +
			`?:^https?://(?:(?:(?:www|classic)\.)?ardmediathek\.de|mediathek\.(?:daserste|rbb` +
			`-online)\.de|one\.ard\.de)/(?:.*/)(?:[0-9]+|[^0-9][^/\?]+)[^/\?]*(?:\?.*)?)|(?:h` +
			`ttps?://(?:www\.)?audi-mediacenter\.com/(?:en|de)/audimediatv/(?:video/)?(?:[^/?` +
			`#]+))|(?:https?://(?:www\.)?audioboom\.com/(?:boos|posts)/(?:[0-9]+))|(?:https?:` +
			`//(?:www\.)?audiomack\.com/song/(?:[\w/-]+))|(?:https?://[^/]+\.bandcamp\.com/tr` +
			`ack/(?:[^/?#&]+))|(?:https?://(?:www\.)?bandcamp\.com/?\?(?:.*?&)?show=(?:\d+))|` +
			`(?:https?://(?:www\.|pro\.)?beatport\.com/track/(?:[^/]+)/(?:[0-9]+))|(?:https?:` +
			`//(?:www\.)?bigflix\.com/.+/(?:[0-9]+))|(?:https?://(?:www\.)?bild\.de/(?:[^/]+/` +
			`)+(?:[^/]+)-(?:\d+)(?:,auto=true)?\.bild\.html)|(?:https?://(?:www\.)?bitchute\.` +
			`com/(?:video|embed|torrent/[^/]+)/(?:[^/?#&]+))|(?:https?://(?:www\.)?bleacherre` +
			`port\.com/video_embed\?id=(?:[0-9a-f-]{36}))|(?:https?://(?:www\.)?bpb\.de/media` +
			`thek/(?:[0-9]+)/)|(?:(?:https?://(?:www\.)?br(?:-klassik)?\.de)/(?:[a-z0-9\-_]+/` +
			`)+(?:[a-z0-9\-_]+)\.html)|(?:(?:https?://.*brightcove\.com/(services|viewer).*?\` +
			`?|brightcove:)(?:.*))|(?:https?://players\.brightcove\.net/(?:\d+)/(?:[^/]+)_(?:` +
			`[^/]+)/index\.html\?.*videoId=(?:\d+|ref:[^&]+))|(?:https?://(?:[^/]+\.)?busines` +
			`sinsider\.(?:com|nl)/(?:[^/]+/)*(?:[^/?#&]+))|(?:https?://(?:www\.)?camdemy\.com` +
			`/media/(?:\d+))|(?:https?://(?:www\.)?(?:mycanal|piwiplus)\.fr/(?:[^/]+/)*(?:[^?` +
			`/]+)(?:\.html\?.*\bvid=|/p/)(?:\d+))|(?:https?://(?:(?:www\.)?canalc2\.tv/video/` +
			`|archives-canalc2\.u-strasbg\.fr/video\.asp\?.*\bidVideo=)(?:\d+))|(?:https?://(` +
			`?:www\.)?media\.ccc\.de/v/(?:[^/?#&]+))|(?:https?://(?:www\.)?charlierose\.com/(` +
			`?:video|episode)(?:s|/player)/(?:\d+))|(?:https?://(?:www\.)?chirb\.it/(?:(?:wp|` +
			`pl)/|fb_chirbit_player\.swf\?key=)?(?:[\da-zA-Z]+))|(?:https?://player\.cinchcas` +
			`t\.com/.*?(?:assetId|show_id)=(?:[0-9]+))|(?:https?://(?:www\.)?clippituser\.tv/` +
			`c/(?:[a-z]+))|(?:https?://(?:chic|www)\.clipsyndicate\.com/video/play(list/\d+)?` +
			`/(?:\d+))|(?:https?://(?:www\.)?clyp\.it/(?:[a-z0-9]+))|(?:https?://(?:(?:editio` +
			`n|www|money)\.)?cnn\.com/(?:video/(?:data/.+?|\?)/)?videos?/(?:.+?/(?:[^/]+?)(?:` +
			`\.(?:[a-z\-]+)|(?=&))))|(?:https?://[^\.]+\.blogs\.cnn\.com/.+)|(?:https?://(?:(` +
			`?:edition|www)\.)?cnn\.com/(?!videos?/))|(?:(?:coub:|https?://(?:coub\.com/(?:vi` +
			`ew|embed|coubs)/|c-cdn\.coub\.com/fb-player\.swf\?.*\bcoub(?:ID|id)=))(?:[\da-z]` +
			`+))|(?:https?://(?:www\.)?cracked\.com/video_(?:\d+)_[\da-z-]+\.html)|(?:https?:` +
			`//(?:www\.)?c-span\.org/video/\?(?:[0-9a-f]+))|(?:https?://news\.cts\.com\.tw/[a` +
			`-z]+/[a-z]+/\d+/(?:\d+)\.html)|(?:https?://app\.curiositystream\.com/video/(?:\d` +
			`+))|(?:https?://(?:www\.)?dailymail\.co\.uk/(?:video/[^/]+/video-|embed/video/)(` +
			`?:[0-9]+))|(?:(?i)https?://(?:(www|touch)\.)?dailymotion\.[a-z]{2,3}/(?:(?:(?:em` +
			`bed|swf|#)/)?video|swf)/(?:[^/?_]+))|(?:https?://(?:www\.)?digg\.com/video/(?:[^` +
			`/?#&]+))|(?:https?://(?:www\.)?dotsub\.com/view/(?:[^/]+))|(?:https?://(?:www\.)` +
			`?dr\.dk/(?:tv/se|nyheder|radio/ondemand)/(?:[^/]+/)*(?:[\da-z-]+)(?:[/#?]|$))|(?` +
			`:https?://(?:www\.)?(?:discovery|tlc|animalplanet|dmax)\.de/(?:.*\#(?:\d+)|(?:[^` +
			`/]+/)*videos/(?:[^/?#]+)|programme/(?:[^/]+)/video/(?:[^/]+)))|(?:https?://(?:(?` +
			`:[^/]+\.)?(?:disney\.[a-z]{2,3}(?:\.[a-z]{2})?|disney(?:(?:me|latino)\.com|turki` +
			`ye\.com\.tr|channel\.de)|(?:starwars|marvelkids)\.com))/(?:(?:embed/|(?:[^/]+/)+` +
			`[\w-]+-)(?:[a-z0-9]{24})|(?:[^/]+/)?(?:[^/?#]+)))|(?:https?://(?:s?evt\.dispeak|` +
			`events\.digitallyspeaking)\.com/(?:[^/]+/)+xml/(?:[^.]+)\.xml)|(?:https?://(?:ww` +
			`w\.)?dw\.com/(?:[^/]+/)+(?:av|e)-(?:\d+))|(?:https?://(?:www\.)?ebaumsworld\.com` +
			`/videos/[^/]+/(?:\d+))|(?:https?://(?:www\.)?echo\.msk\.ru/sounds/(?:\d+))|(?:ht` +
			`tps?://(?:[^.]+\.)?elpais\.com/.*/(?:[^/#?]+)\.html(?:$|[?#]))|(?:https?://(?:ww` +
			`w\.)?engadget\.com/video/(?:[^/?#]+))|(?:https?://ec\.europa\.eu/avservices/(?:v` +
			`ideo/player|audio/audioDetails)\.cfm\?.*?\bref=(?:[A-Za-z0-9-]+))|(?:https?://(?` +
			`:www\.)?extremetube\.com/(?:[^/]+/)?video/(?:[^/#?&]+))|(?:(?:https?://(?:[\w-]+` +
			`\.)?(?:facebook\.com|facebookcorewwwi\.onion)/(?:[^#]*?\#!/)?(?:(?:video/video\.` +
			`php|photo\.php|video\.php|video/embed|story\.php)\?(?:.*?)(?:v|video_id|story_fb` +
			`id)=|[^/]+/videos/(?:[^/]+/)?|[^/]+/posts/|groups/[^/]+/permalink/)|facebook:)(?` +
			`:[0-9]+))|(?:https?://(?:[\w-]+\.)?facebook\.com/plugins/video\.php\?.*?\bhref=(` +
			`?:https.+))|(?:https?://(?:www\.)?faz\.net/(?:[^/]+/)*.*?-(?:\d+)\.html)|(?:(?:5` +
			`min:|https?://(?:[^/]*?5min\.com/|delivery\.vidible\.tv/aol)(?:(?:Scripts/Player` +
			`Seed\.js|playerseed/?)?\?.*?playList=)?)(?:\d+))|(?:http://(?:www\.)?5-tv\.ru/(?` +
			`:(?:[^/]+/)+(?:\d+)|(?:[^/?#]+)(?:[/?#])?))|(?:https?://(?:www\.|secure\.)?flick` +
			`r\.com/photos/[\w\-_@]+/(?:\d+))|(?:https?://(?:www\.)?formula1\.com/(?:content/` +
			`fom-website/)?en/video/\d{4}/\d{1,2}/(?:.+?)\.html)|(?:https?://(?:video\.(?:ins` +
			`ider\.)?fox(?:news|business)\.com)/v/(?:video-embed\.html\?video_id=)?(?:\d+))|(` +
			`?:https?://(?:www\.)?(?:insider\.)?foxnews\.com/(?!v)([^/]+/)+(?:[a-z-]+))|(?:ht` +
			`tps?://(?:www\.)?franceculture\.fr/emissions/(?:[^/]+/)*(?:[^/?#&]+))|(?:https?:` +
			`//(?:www\.)?freesound\.org/people/[^/]+/sounds/(?:[^/]+))|(?:https?://(?:www\.)?` +
			`gamespot\.com/(?:video|article|review)s/(?:[^/]+/\d+-|embed/)(?:\d+))|(?:https?:` +
			`//(?:www\.)?gaskrank\.tv/tv/(?:[^/]+)/(?:[^/]+)\.htm)|(?:https?://(?:www\.)?gdcv` +
			`ault\.com/play/(?:\d+)/(?:(\w|-)+)?)|(?:https?://(?:www\.)?gfycat\.com/(?:ifr/|g` +
			`ifs/detail/)?(?:[^/?#]+))|(?:https?://(?:www\.)?godtube\.com/watch/\?v=(?:[\da-z` +
			`A-Z]+))|(?:^https?://video\.golem\.de/.+?/(?:.+?)/)|(?:https?://on-demand\.gpute` +
			`chconf\.com/gtc/2015/video/S(?:\d+)\.html)|(?:^https?://hentai\.animestigma\.com` +
			`/(?:[^/]+))|(?:https?://(?:www\.)?historicfilms\.com/(?:tapes/|play)(?:\d+))|(?:` +
			`https?://(?:www\.)?hitrecord\.org/records/(?:\d+))|(?:http?://(?:www\.)?hornbunn` +
			`y\.com/videos/(?:[a-z-]+)-(?:\d+)\.html)|(?:https?://[\da-z-]+\.(?:howstuffworks` +
			`|stuff(?:(?:youshould|theydontwantyouto)know|toblowyourmind|momnevertoldyou)|(?:` +
			`brain|car)stuffshow|fwthinking|geniusstuff)\.com/(?:[^/]+/)*(?:\d+-)?(?:.+?)-vid` +
			`eo\.htm)|(?:https?://.+?\.ign\.com/(?:[^/]+/)?(?:videos|show_videos|articles|fea` +
			`ture|(?:[^/]+/\d+/video))(/.+)?/(?:.+))|(?:https?://(?:www\.)?pcmag\.com/(?:vide` +
			`os|article2)(/.+)?/(?:.+))|(?:https?://(?:www|m)\.imdb\.com/(?:video|title|list)` +
			`.+?[/-]vi(?:\d+))|(?:https?://(?:i\.)?imgur\.com/(?!(?:a|gallery|(?:t(?:opic)?|r` +
			`)/[^/]+)/)(?:[a-zA-Z0-9]+))|(?:https?://(?:i\.)?imgur\.com/(?:gallery|(?:t(?:opi` +
			`c)?|r)/[^/]+)/(?:[a-zA-Z0-9]+))|(?:https?://(?:www\.)?ina\.fr/video/(?:I?[A-Z0-9` +
			`]+))|(?:https?://(?:www\.)?infoq\.com/(?:[^/]+/)+(?:[^/]+))|(?:(?:https?://(?:ww` +
			`w\.)?instagram\.com/p/(?:[^/?#&]+)))|(?:https?://(?:www\.)?90tv\.ir/video/(?:[0-` +
			`9]+)/.*)|(?:https?://(?:www\.)?ivideon\.com/tv/(?:[^/]+/)*camera/(?:\d+-[\da-f]+` +
			`)/(?:\d+))|(?:https?://(?:(?:www|m)\.)?izlesene\.com/(?:video|embedplayer)/(?:[^` +
			`/]+/)?(?:[0-9]+))|(?:https?://(?:licensing\.jamendo\.com/[^/]+|(?:www\.)?jamendo` +
			`\.com)/track/(?:[0-9]+)/(?:[^/?#&]+))|(?:https?://.*?\.jeuxvideo\.com/.*/(.*?)\.` +
			`htm)|(?:(?:joj:|https?://media\.joj\.sk/embed/)(?:[^/?#^]+))|(?:(?:https?://cont` +
			`ent\.jwplatform\.com/(?:feeds|players|jw6)/|jwplatform:)(?:[a-zA-Z0-9]{8}))|(?:h` +
			`ttps?://(?:www\.)?keezmovies\.com/video/(?:(?:[^/]+)-)?(?:\d+))|(?:https?://(?:w` +
			`ww\.)?kickstarter\.com/projects/(?:[^/]*)/.*)|(?:https?://krasview\.ru/(?:video|` +
			`embed)/(?:\d+))|(?:https?://(?:www\.)?loc\.gov/(?:item/|today/cyberlc/feature_wd` +
			`esc\.php\?.*\brec=)(?:[0-9a-z_.]+))|(?:(?:limelight:media:|https?://(?:link\.vid` +
			`eoplatform\.limelight\.com/media/|assets\.delvenetworks\.com/player/loader\.swf)` +
			`\?.*?\bmediaId=)(?:[a-z0-9]{32}))|(?:https?://(?:new\.)?livestream\.com/(?:accou` +
			`nts/(?:\d+)|(?:[^/]+))/(?:events/(?:\d+)|(?:[^/]+))(?:/videos/(?:\d+))?)|(?:http` +
			`s?://(?:www\.)?lovehomeporn\.com/video/(?:\d+)(?:/(?:[^/?#&]+))?)|(?:https?://(?` +
			`:www\.)?macgamestore\.com/mediaviewer\.php\?trailer=(?:\d+))|(?:https?://(?:www\` +
			`.)?massengeschmack\.tv/play/(?:[^?&#]+))|(?:https://player\.megaphone\.fm/(?:[A-` +
			`Z0-9]+))|(?:https?://(?:www\.)?metacafe\.com/watch/(?:[^/]+)/(?:[^/?#]+))|(?:htt` +
			`ps?://(?:www\.)?metacritic\.com/.+?/trailers/(?:\d+))|(?:(?:mva:|https?://(?:mva` +
			`\.microsoft|(?:www\.)?microsoftvirtualacademy)\.com/[^/]+/training-courses/[^/?#` +
			`&]+-)(?:\d+)(?::|\?l=)(?:[\da-zA-Z]+_\d+))|(?:https?://(?:[\da-z_-]+\.)*(?:mlb)\` +
			`.com/(?:(?:(?:[^/]+/)*c-|(?:shared/video/embed/(?:embed|m-internal-embed)\.html|` +
			`(?:[^/]+/)+(?:play|index)\.jsp|)\?.*?\bcontent_id=)(?:\d+)))|(?:https?://(?:www\` +
			`.)?mofosex\.com/videos/(?:\d+)/(?:[^/?#&.]+)\.html)|(?:https?://(?:www\.)?mojvid` +
			`eo\.com/video-(?:[^/]+)/(?:[a-f0-9]+))|(?:(?:https?://(?:www\.)?myvi\.(?:(?:ru/p` +
			`layer|tv)/(?:(?:embed/html|flash|api/Video/Get)/|content/preloader\.swf\?.*\bid=` +
			`)|ru/watch/)|myvi:)(?:[\da-zA-Z_-]+))|(?:https?://(?:watch\.|www\.)?nba\.com/(?:` +
			`(?:[^/]+/)+(?:[^?]*?))/?(?:/index\.html)?(?:\?.*)?$)|(?:https?://(?:www\.)?ndr\.` +
			`de/(?:[^/]+/)*(?:[^/?#]+),[\da-z]+\.html)|(?:https?://(?:www\.)?ndr\.de/(?:[^/]+` +
			`/)*(?:[\da-z]+)-(?:player|externalPlayer)\.html)|(?:https?://(?:www\.)?n-joy\.de` +
			`/(?:[^/]+/)*(?:[\da-z]+)-(?:player|externalPlayer)_[^/]+\.html)|(?:https?://(?:[` +
			`^/]+\.)?ndtv\.com/(?:[^/]+/)*videos?/?(?:[^/]+/)*[^/?^&]+-(?:\d+))|(?:https?://(` +
			`?:www\.)?netzkino\.de/\#!/(?:[^/]+)/(?:[^/]+))|(?:https?://(?:www\.)?newgrounds\` +
			`.com/(?:audio/listen|portal/view)/(?:[0-9]+))|(?:https?://hk\.dv\.nextmedia\.com` +
			`/actionnews/[^/]+/(?:\d+)/(?:\d+)/\d+)|(?:https?://(?:www\.)?nexttv\.com\.tw/(?:` +
			`[^/]+/)+(?:\d+))|(?:(?:https?://api\.nexx(?:\.cloud|cdn\.com)/v3/(?:\d+)/videos/` +
			`byid/|nexx:(?:(?:\d+):)?|https?://arc\.nexx\.cloud/api/video/)(?:\d+))|(?:https?` +
			`://(?:www\.)?(?:nhl|wch2016)\.com/(?:[^/]+/)*c-(?:\d+))|(?:https?://media\.cms\.` +
			`nova\.cz/embed/(?:[^/?#&]+))|(?:https?://(?:(?:www|cn)\.)?nowness\.com/(?:story|` +
			`(?:series|category)/[^/]+)/(?:[^/]+?)(?:$|[?#]))|(?:https?://(?:www\.)?ntv\.ru/(` +
			`?:[^/]+/)*(?:[^/?#&]+))|(?:https?://(?:(?:www|m|mobile)\.)?(?:odnoklassniki|ok)\` +
			`.ru/(?:video(?:embed)?/|web-api/video/moviePlayer/|live/|dk\?.*?st\.mvId=)(?:[\d` +
			`-]+))|(?:(?:ooyala:|https?://.+?\.ooyala\.com/.*?(?:embedCode|ec)=)(?:.+?)(&|$))` +
			`|(?:https?://(?:www\.)?packtpub\.com/mapt/video/[^/]+/(?:\d+)/(?:\d+)/(?:\d+))|(` +
			`?:https?://(?:(?:www\.)?pandora\.tv/view/(?:[^/]+)/(?:\d+)|(?:.+?\.)?channel\.pa` +
			`ndora\.tv/channel/video\.ptv\?|m\.pandora\.tv/?\?))|(?:https?://(?:www\.)?patreo` +
			`n\.com/(?:creation\?hid=|posts/(?:[\w-]+-)?)(?:\d+))|(?:https?://(?:www\.)?pearv` +
			`ideo\.com/video_(?:\d+))|(?:https?://(?:www\.)?people\.com/people/videos/0,,(?:\` +
			`d+),00\.html)|(?:https?://(?:[a-z0-9]+\.)?photobucket\.com/.*(([\?\&]current=)|_` +
			`)(?:.*)\.(?:(flv)|(mp4)))|(?:https?://(?:www\.)?play\.fm/(?:(?:[^/]+/)+(?:[^/]+)` +
			`)/?(?:$|[?#]))|(?:https?://(?:www\.)?plays\.tv/(?:video|embeds)/(?:[0-9a-f]{18})` +
			`)|(?:https?://(?:www\.)?playvid\.com/watch(\?v=|/)(?:.+?)(?:#|$))|(?:(?:https?):` +
			`//(?:(?:[^.]+)\.podomatic\.com/entry|(?:www\.)?podomatic\.com/podcasts/(?:[^/]+)` +
			`/episodes)/(?:[^/?#&]+))|(?:https?://(?:www\.)?pokemon\.com/[a-z]{2}(?:.*?play=(` +
			`?:[a-z0-9]{32})|/(?:[^/]+/)+(?:[^/?#&]+)))|(?:(?:https?://)(?:www\.|)91porn\.com` +
			`/.+?\?viewkey=(?:[\w\d]+))|(?:https?://(?:[a-zA-Z]+\.)?porn\.com/videos/(?:(?:[^` +
			`/]+)-)?(?:\d+))|(?:https?://(?:www\.)?pornhd\.com/(?:[a-z]{2,4}/)?videos/(?:\d+)` +
			`(?:/(?:.+))?)|(?:https?://(?:(?:[^/]+\.)?(?:pornhub\.(?:com|net))/(?:(?:view_vid` +
			`eo\.php|video/show)\?viewkey=|embed/)|(?:www\.)?thumbzilla\.com/video/)(?:[\da-z` +
			`]+))|(?:https?://(?:\w+\.)?pornotube\.com/(?:[^?#]*?)/video/(?:[0-9]+))|(?:https` +
			`?://(?:www\.)?puhutv\.com/(?:[^/?#&]+)-izle)|(?:https?://(?:www\.)?presstv\.ir/[` +
			`^/]+/(?:\d+)/(?:\d+)/(?:\d+)/(?:\d+)/(?:[^/]+)?)|(?:https?://(?:.+?)\.(?:radio\.` +
			`(?:de|at|fr|pt|es|pl|it)|rad\.io))|(?:https?://(?:www\.)?radiojavan\.com/videos/` +
			`video/(?:[^/]+)/?)|(?:https?://[^/]+\.(?:rai\.(?:it|tv)|rainews\.it)/dl/.+?-(?:[` +
			`\da-f]{8}-[\da-f]{4}-[\da-f]{4}-[\da-f]{4}-[\da-f]{12})(?:-.+?)?\.html)|(?:https` +
			`?://(?:videos\.raywenderlich\.com/courses|(?:www\.)?raywenderlich\.com)/(?:[^/]+` +
			`)/lessons/(?:\d+))|(?:https?://(?:www\.)?(?:rbmaradio|redbullradio)\.com/shows/(` +
			`?:[^/]+)/episodes/(?:[^/?#&]+))|(?:https?://(?:(?:www\.)?redtube\.com/|embed\.re` +
			`dtube\.com/\?.*?\bid=)(?:[0-9]+))|(?:(?:rentv:|https?://(?:www\.)?ren\.tv/(?:pla` +
			`yer|video/epizod)/)(?:\d+))|(?:^https?://(?:www\.)?reverbnation\.com/.*?/song/(?` +
			`:\d+).*?$)|(?:https?://(?:www\.)?rockstargames\.com/videos(?:/video/|#?/?\?.*\bv` +
			`ideo=)(?:\d+))|(?:https?://(?:www\.)?prehravac\.rozhlas\.cz/audio/(?:[0-9]+))|(?` +
			`:rts:(?:\d+)|https?://(?:.+?\.)?rts\.ch/(?:[^/]+/){2,}(?:[0-9]+)-(?:.+?)\.html)|` +
			`(?:https?://(?:www\.)?rtvs\.sk/(?:radio|televizia)/archiv/\d+/(?:\d+))|(?:https?` +
			`://(?:test)?player\.(?:rutv\.ru|vgtrk\.com)/(?:flash\d+v/container\.swf\?id=|ifr` +
			`ame/(?:swf|video|live)/id/|index/iframe/cast_id/)(?:\d+))|(?:https?://(?:(?:v2|w` +
			`ww)\.)?videos\.sapo\.(?:pt|cv|ao|mz|tl)/(?:[\da-zA-Z]{20}))|(?:https?://(?:www\.` +
			`)?screencast\.com/t/(?:[a-zA-Z0-9]+))|(?:https?://(?:www\.)?senate\.gov/isvp/?\?` +
			`(?:.+))|(?:https?://(?:www\.)?seznamzpravy\.cz/iframe/player\?.*\bsrc=)|(?:https` +
			`?://vivo\.sx/(?:[\da-z]{10}))|(?:https?://(?:www\.)?skysports\.com/watch/video/(` +
			`?:[0-9]+))|(?:https?://(?:www\.)?slideshare\.net/[^/]+?/(?:.+?)($|\?))|(?:https?` +
			`://(?:\w+\.)?slutload\.com/(?:video/[^/]+|embed_player|watch)/(?:[^/]+))|(?:http` +
			`s?://(?:www\.)?soundgasm\.net/u/(?:[0-9a-zA-Z_-]+)/(?:[0-9a-zA-Z_-]+))|(?:https?` +
			`://(?:(?:www|m|[a-z]{2})\.)?spankbang\.com/(?:[\da-z]+)/video)|(?:https?://(?:ww` +
			`w\.)?stitcher\.com/podcast/(?:[^/]+/)+e/(?:(?:[^/#?&]+?)-)?(?:\d+)(?:[/#?&]|$))|` +
			`(?:https?://sportdeutschland\.tv/(?:[^/?#]+)/(?:[^?#/]+)(?:$|[?#]))|(?:https?://` +
			`(?:(?:www|play)\.)?(?:srf|rts|rsi|rtr|swissinfo)\.ch/play/(?:tv|radio)/[^/]+/(?:` +
			`video|audio)/[^?]+\?id=(?:[0-9a-f\-]{36}|\d+))|(?:https?://openclassroom\.stanfo` +
			`rd\.edu(?:/?|(/MainFolder/(?:HomePage|CoursePage|VideoPage)\.php([?]course=(?:[^` +
			`&]+)(&video=(?:[^&]+))?(&.*)?)?))$)|(?:https?://streamable\.com/(?:[es]/)?(?:\w+` +
			`))|(?:https?://(?:(?:www\.)?sunporno\.com/videos|embeds\.sunporno\.com/embed)/(?` +
			`:\d+))|(?:https?://(?:www\.)?swrmediathek\.de/(?:content/)?player\.htm\?show=(?:` +
			`[\da-f]{8}-[\da-f]{4}-[\da-f]{4}-[\da-f]{4}-[\da-f]{12}))|(?:https?://(?:www\.)?` +
			`tagesschau\.de/(?:[^/]+/(?:[^/]+/)*?(?:[^/#?]+?(?:-?[0-9]+)?))(?:~_?[^/#?]+?)?\.` +
			`html)|(?:https?://(?:www\.)?tastytrade\.com/tt/shows/[^/]+/episodes/(?:[^/?#&]+)` +
			`)|(?:https?://tds\.lifeway\.com/v1/trainingdeliverysystem/courses/(?:\d+)/index\` +
			`.html)|(?:https?://(?:www\.)?teachertube\.com/(viewVideo\.php\?video_id=|music\.` +
			`php\?music_id=|video/(?:[\da-z-]+-)?|audio/)(?:\d+))|(?:(?:https?://)(?:www|embe` +
			`d(?:-ssl)?)(?:\.ted\.com/((?:playlists(?:/\d+)?)|((?:talks))|(?:watch)/[^/]+/[^/` +
			`]+)(/lang/(.*?))?/(?:[\w-]+).*)$)|(?:^https?://(?:www\.)?t13\.cl/videos(?:/[^/]+` +
			`)+/(?:[\w-]+))|(?:https?://(?:www\.)?tenta\.com/how-to-download-videos)|(?:https` +
			`?://(?:www\.)?thisoldhouse\.com/(?:watch|how-to|tv-episode)/(?:[^/?#]+))|(?:http` +
			`s?://(?:m\.)?tiktok\.com/v/(?:\d+))|(?:https?://(?:.+?\.)?tinypic\.com/player\.p` +
			`hp\?v=(?:[^&]+)&s=\d+)|(?:https?://player\.(?:tna|emp)flix\.com/video/(?:\d+))|(` +
			`?:https?://(?:www\.)?tnaflix\.com/[^/]+/(?:[^/]+)/video(?:\d+))|(?:https?://(?:w` +
			`ww\.)?empflix\.com/(?:videos/(?:.+?)-|[^/]+/(?:[^/]+)/video)(?:[0-9]+))|(?:https` +
			`?://(?:www\.)?moviefap\.com/videos/(?:[0-9a-f]+)/(?:[^/]+)\.html)|(?:https?://vi` +
			`deos\.toypics\.net/view/(?:[0-9]+))|(?:https?://(?:(?:www|m)\.)?trilulilu\.ro/(?` +
			`:[^/]+/)?(?:[^/#\?]+))|(?:https?://(?:www\.)?tube8\.com/(?:[^/]+/)+(?:[^/]+)/(?:` +
			`\d+))|(?:https?://(?:[^/?#&]+)\.tumblr\.com/(?:post|video)/(?:[0-9]+)(?:$|[/?#])` +
			`)|(?:https?://(?:www\.)?tunein\.com/station/.*?audioClipId\=(?:\d+))|(?:https?:/` +
			`/(?:www\.)?tunein\.com/(?:radio/.*?-s|station/.*?StationId=|embed/player/s)(?:\d` +
			`+))|(?:https?://(?:www\.)?tunein\.com/(?:radio/.*?-p|program/.*?ProgramId=|embed` +
			`/player/p)(?:\d+))|(?:https?://tun\.in/(?:[A-Za-z0-9]+))|(?:https?://(?:www\.)?t` +
			`vanouvelles\.ca/videos/(?:\d+))|(?:https?://(?:www\.)?tvc\.ru/video/iframe/id/(?` +
			`:\d+))|(?:https?://(?:www\.)?tvc\.ru/(?!video/iframe/id/)(?:[^?#]+))|(?:https?:/` +
			`/(?:(?:[^/]+)\.)?tvn24(?:bis)?\.pl/(?:[^/]+/)*(?:[^/]+))|(?:https?://(?:www\.)?t` +
			`vnoe\.cz/video/(?:[0-9]+))|(?:https?://[^/]+\.tvp\.(?:pl|info)/(?:video/(?:[^,\s` +
			`]*,)*|(?:(?!\d+/)[^/]+/)*)(?:\d+))|(?:https?://tvplay\.(?:tv3\.lt|skaties\.lv|tv` +
			`3\.ee)/[^/]+/[^/?#&]+-(?:\d+))|(?:https?://(?:www\.)?20min\.ch/(?:videotv/*\?.*?` +
			`\bvid=|videoplayer/videoplayer\.html\?.*?\bvideoId@)(?:\d+))|(?:https?://(?:clip` +
			`s\.twitch\.tv/(?:[^/]+/)*|(?:www\.)?twitch\.tv/[^/]+/clip/)(?:[^/?#&]+))|(?:http` +
			`s?://amp\.twimg\.com/v/(?:[0-9a-f\-]{36}))|(?:https?://video\.udn\.com/(?:embed|` +
			`play)/news/(?:\d+))|(?:https?://(?:www\.)?(?:digiteka\.net|ultimedia\.com)/(?:de` +
			`liver/(?:generic|musique)(?:/[^/]+)*/(?:src|article)|default/index/video(?:gener` +
			`ic|music)/id)/(?:[\d+a-z]+))|(?:https?://utv\.unistra\.fr/(?:index|video)\.php\?` +
			`id_video\=(?:\d+))|(?:https?://(?:.+?\.)?uol\.com\.br/.*?(?:(?:mediaId|v)=|view/` +
			`(?:[a-z0-9]+/)?|video(?:=|/(?:\d{4}/\d{2}/\d{2}/)?))(?:\d+|[\w-]+-[A-Z0-9]+))|(?` +
			`:https?://(?:www\.)?usatoday\.com/(?:[^/]+/)*(?:[^?/#]+))|(?:https?://(?:(?:app|` +
			`embed)\.)?ustudio\.com/embed/(?:[^/]+)/(?:[^/]+))|(?:https?://(?:www\.)?video\.v` +
			`arzesh3\.com/(?:[^/]+/)+(?:[^/]+)/?)|(?:https?://(?:[^/]+\.)?vbox7\.com/(?:play:` +
			`|(?:emb/external\.php|player/ext\.swf)\?.*?\bvid=)(?:[\da-fA-F]+))|(?:https?://v` +
			`eehd\.com/video/(?:\d+))|(?:https?://(?:www\.)?veoh\.com/(?:watch|embed|iphone/#` +
			`_Watch)/(?:(?:v|e|yapi-)[\da-zA-Z]+))|(?:https?://(?:.+?\.)?vesti\.ru/(?:.+))|(?` +
			`:(?:https?://(?:www\.)?vevo\.com/watch/(?!playlist|genre)(?:[^/]+/(?:[^/]+/)?)?|` +
			`https?://cache\.vevo\.com/m/html/embed\.html\?video=|https?://videoplayer\.vevo\` +
			`.com/embed/embedded\?videoId=|vevo:)(?:[^&?#]+))|(?:(?:https?://(?:www\.)?(?:vgt` +
			`v.no|bt.no/tv|aftenbladet.no/tv|fvn.no/fvntv|aftenposten.no/webtv|ap.vgtv.no/web` +
			`tv|tv.aftonbladet.se/abtv|www.aftonbladet.se/tv)/?(?:(?:\#!/)?(?:video|live)/|em` +
			`bed?.*id=|a(?:rticles)?/)|(?:vgtv|bttv|satv|fvntv|aptv|abtv):)(?:\d+))|(?:https?` +
			`://(?:www\.)?viddler\.com/(?:v|embed|player)/(?:[a-z0-9]+))|(?:https?://videopre` +
			`ss\.com/embed/(?:[\da-zA-Z]+))|(?:https?://(?:www\.|m\.|)videosz\.com/[a-z]+/[a-` +
			`z]+/(?:[0-9]+)_)|(?:https?://(?:www\.)?vidlii\.com/(?:watch|embed)\?.*?\bv=(?:[0` +
			`-9A-Za-z_-]{11}))|(?:https?://(?:(?:www|(?:player))\.)?vimeo(?:pro)?\.com/(?!(?:` +
			`channels|album)/[^/?#]+/?(?:$|[?#])|[^/]+/review/|ondemand/)(?:.*?/)?(?:(?:play_` +
			`redirect_hls|moogaloop\.swf)\?clip_id=)?(?:videos?/)?(?:[0-9]+)(?:/[\da-f]+)?/?(` +
			`?:[?&].*)?(?:[#].*)?$)|(?:https?://(?:www\.)?vimeo\.com/ondemand/(?:[^/?#&]+))|(` +
			`?:https://vimeo\.com/[^/]+/review/(?:[^/]+))|(?:https?://(?:player\.vimple\.(?:r` +
			`u|co)/iframe|vimple\.(?:ru|co))/(?:[\da-f-]{32,36}))|(?:https?://(?:www\.)?vine\` +
			`.co/(?:v|oembed)/(?:\w+))|(?:(?:viqeo:|https?://cdn\.viqeo\.tv/embed/*\?.*?\bvid` +
			`=|https?://api\.viqeo\.tv/v\d+/data/startup?.*?\bvideo(?:%5B%5D|\[\])=)(?:[\da-f` +
			`]+))|(?:https?://(?:(?:(?:(?:m|new)\.)?vk\.com/video_|(?:www\.)?daxab.com/)ext\.` +
			`php\?(?:.*?\boid=(?:-?\d+).*?\bid=(?:\d+).*)|(?:(?:(?:m|new)\.)?vk\.com/(?:.+?\?` +
			`.*?z=)?video|(?:www\.)?daxab.com/embed/)(?:-?\d+_\d+)(?:.*\blist=(?:[\da-f]+))?)` +
			`)|(?:https?://(?:(?:www|m)\.)?vlive\.tv/video/(?:[0-9]+))|(?:https?://(?:m\.)?vu` +
			`clip\.com/w\?.*?cid=(?:[0-9]+))|(?:https?://(?:(?:www|view)\.)?vzaar\.com/(?:vid` +
			`eos/)?(?:\d+))|(?:(?:washingtonpost:|https?://(?:www\.)?washingtonpost\.com/vide` +
			`o/(?:[^/]+/)*)(?:[\da-f]{8}-[\da-f]{4}-[\da-f]{4}-[\da-f]{4}-[\da-f]{12}))|(?:(?` +
			`:wat:|https?://(?:www\.)?wat\.tv/video/.*-)(?:[0-9a-z]+))|(?:https?://m\.weibo\.` +
			`cn/status/(?:[0-9]+)(\?.+)?)|(?:https?://(?:www|m)\.worldstar(?:candy|hiphop)\.c` +
			`om/(?:videos|android)/video\.php\?.*?\bv=(?:[^&]+))|(?:https?://(?:.+?\.)?xhamst` +
			`er\.com/(?:movies/(?:\d+)/(?:[^/]*)\.html|videos/(?:[^/]*)-(?:\d+)))|(?:https?:/` +
			`/(?:.+?\.)?xhamster\.com/xembed\.php\?video=(?:\d+))|(?:https?://(?:www\.)?xiami` +
			`\.com/song/(?:[^/?#&]+))|(?:https?://(?:video|www)\.xnxx\.com/video-?(?:[0-9a-z]` +
			`+)/)|(?:(?:xtube:|https?://(?:www\.)?xtube\.com/(?:watch\.php\?.*\bv=|video-watc` +
			`h/(?:embedded/)?(?:[^/]+)-))(?:[^/?&#]+))|(?:https?://(?:(?:www\.)?xvideos\.com/` +
			`video|flashservice\.xvideos\.com/embedframe/|static-hw\.xvideos\.com/swf/xv-play` +
			`er\.swf\?.*?\bid_video=)(?:[0-9]+))|(?:https?://(?:www\.)?xxxymovies\.com/videos` +
			`/(?:\d+)/(?:[^/]+))|(?:https?://v\.yinyuetai\.com/video(?:/h5)?/(?:[0-9]+))|(?:h` +
			`ttps?://(?:\w+\.)?youjizz\.com/videos/(?:[^/#?]*-(?:\d+)\.html|embed/(?:\d+)))|(` +
			`?:https?://(?:www\.)?youporn\.com/watch/(?:\d+)/(?:[^/?#&]+))|(?:https?://(?:www` +
			`\.)?yourporn\.sexy/post/(?:[^/?#&.]+))|(?:https?://(?:www\.)?(?:yourupload\.com/` +
			`(?:watch|embed)|embed\.yourupload\.com)/(?:[A-Za-z0-9]+))|(?:^((?:https?://|//)(` +
			`?:(?:(?:(?:\w+\.)?[yY][oO][uU][tT][uU][bB][eE](?:-nocookie)?\.com/|(?:www\.)?det` +
			`url\.com/www\.youtube\.com/|(?:www\.)?pwnyoutube\.com/|(?:www\.)?hooktube\.com/|` +
			`(?:www\.)?yourepeat\.com/|tube\.majestyc\.net/|(?:www\.)?invidio\.us/|youtube\.g` +
			`oogleapis\.com/)(?:.*?\#/)?(?:(?:(?:v|embed|e)/(?!videoseries))|(?:(?:(?:watch|m` +
			`ovie)(?:_popup)?(?:\.php)?/?)?(?:\?|\#!?)(?:.*?[&;])??v=)))|(?:youtu\.be|vid\.pl` +
			`us|zwearz\.com/watch|)/|(?:www\.)?cleanvideosearch\.com/media/action/yt/watch\?v` +
			`ideoId=))?([0-9A-Za-z_-]{11})(?!.*?\blist=(?:(?:PL|LL|EC|UU|FL|RD|UL|TL)[0-9A-Za` +
			`-z-_]{10,}|WL))$)`,
	}
}
