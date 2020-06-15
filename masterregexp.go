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
 * masterregexp.go: Automatically generated regexp by combining URL validation expressions of all extractors
 */

package downloader

func init() {
	masterRegexps = []string{
		`(?:https?://(?:abcnews\.go\.com/(?:[^/]+/video/(?:[0-9a-z-]+)-|video/embed\?.*?\` +
			`bid=)|fivethirtyeight\.abcnews\.go\.com/video/embed/\d+/)(?:\d+))|(?:https?://(?` +
			`:abc(?:7(?:news|ny|chicago)?|11|13|30)|6abc)\.com(?:(?:/[^/]+)*/(?:[^/]+))?/(?:\` +
			`d+))|(?:https?://(?:(?:(?:embed|www)\.)?acast\.com/|play\.acast\.com/s/)(?:[^/]+` +
			`)/(?:[^/#?]+))|(?:https?://video\.tv\.adobe\.com/v/(?:\d+))|(?:(?:aol-video:|htt` +
			`ps?://(?:www\.)?aol\.(?:com|ca|co\.uk|de|jp)/video/(?:[^/]+/)*)(?:[0-9a-f]+))|(?` +
			`:https?://(?:www\.)?allocine\.fr/(?:article|video|film)/(?:fichearticle_gen_cart` +
			`icle=|player_gen_cmedia=|fichefilm_gen_cfilm=|video-)(?:[0-9]+)(?:\.html)?)|(?:h` +
			`ttps://(?:(?:beta|www)\.)?ardmediathek\.de/(?:[^/]+)/(?:player|live|video)/(?:(?` +
			`:[^/]+/)*)(?:[a-zA-Z0-9]+))|(?:^https?://(?:(?:(?:www|classic)\.)?ardmediathek\.` +
			`de|mediathek\.(?:daserste|rbb-online)\.de|one\.ard\.de)/(?:.*/)(?:[0-9]+|[^0-9][` +
			`^/\?]+)[^/\?]*(?:\?.*)?)|(?:https?://(?:www\.)?arte\.tv/(?:fr|de|en|es|it|pl)/vi` +
			`deos/(?:\d{6}-\d{3}-[AF]))|(?:https?://(?:www\.)?(?:(?:(?:asiancrush|yuyutv|midn` +
			`ightpulp)\.com|cocoro\.tv))/video/(?:[^/]+/)?0+(?:\d+)v\b)|(?:https?://(?:www\.)` +
			`?audioboom\.com/(?:boos|posts)/(?:[0-9]+))|(?:https?://(?:www\.)?audiomack\.com/` +
			`song/(?:[\w/-]+))|(?:https?://(?:www\.)?(?:telezueri\.ch|telebaern\.tv|telem1\.c` +
			`h)/[^/]+/(?:[^/]+-(?:\d+))(?:\#video=(?:[_0-9a-z]+))?)|(?:https?://[^/]+\.bandca` +
			`mp\.com/track/(?:[^/?#&]+))|(?:https?://(?:www\.)?bandcamp\.com/?\?(?:.*?&)?show` +
			`=(?:\d+))|(?:https?://(?:www\.|pro\.)?beatport\.com/track/(?:[^/]+)/(?:[0-9]+))|` +
			`(?:https?://(?:www\.)?bigflix\.com/.+/(?:[0-9]+))|(?:https?://(?:www\.)?bild\.de` +
			`/(?:[^/]+/)+(?:[^/]+)-(?:\d+)(?:,auto=true)?\.bild\.html)|(?:https?://(?:www\.)?` +
			`bilibili\.com/audio/au(?:\d+))|(?:https?://(?:www\.)?biqle\.(?:com|org|ru)/watch` +
			`/(?:-?\d+_\d+))|(?:https?://(?:www\.)?bleacherreport\.com/articles/(?:\d+))|(?:h` +
			`ttps?://(?:www\.)?bleacherreport\.com/video_embed\?id=(?:[0-9a-f-]{36}|\d{5}))|(` +
			`?:https?://union\.bokecc\.com/playvideo\.bo\?(?:.*))|(?:https?://(?:www\.)?bpb\.` +
			`de/mediathek/(?:[0-9]+)/)|(?:(?:https?://(?:www\.)?br(?:-klassik)?\.de)/(?:[a-z0` +
			`-9\-_]+/)+(?:[a-z0-9\-_]+)\.html)|(?:https?://(?:www\.)?bravotv\.com/(?:[^/]+/)+` +
			`(?:[^/?#]+))|(?:https?://(?:[^/]+\.)?businessinsider\.(?:com|nl)/(?:[^/]+/)*(?:[` +
			`^/?#&]+))|(?:https?://(?:www\.)?camdemy\.com/media/(?:\d+))|(?:https?://(?:www\.` +
			`)?(?:mycanal|piwiplus)\.fr/(?:[^/]+/)*(?:[^?/]+)(?:\.html\?.*\bvid=|/p/)(?:\d+))` +
			`|(?:https?://(?:(?:www\.)?canalc2\.tv/video/|archives-canalc2\.u-strasbg\.fr/vid` +
			`eo\.asp\?.*\bidVideo=)(?:\d+))|(?:(?:cbcplayer:|https?://(?:www\.)?cbc\.ca/(?:pl` +
			`ayer/play/|i/caffeine/syndicate/\?mediaId=))(?:\d+))|(?:https?://(?:www\.)?media` +
			`\.ccc\.de/v/(?:[^/?#&]+))|(?:https?://(?:www\.)?ccma\.cat/(?:[^/]+/)*?(?:video|a` +
			`udio)/(?:\d+))|(?:https?://(?:(?:[^/]+)\.(?:cntv|cctv)\.(?:com|cn)|(?:www\.)?ncp` +
			`a-classic\.com)/(?:[^/]+/)*?(?:[^/?#&]+?)(?:/index)?(?:\.s?html|[?#&]|$))|(?:htt` +
			`ps?://(?:www\.)?charlierose\.com/(?:video|episode)(?:s|/player)/(?:\d+))|(?:http` +
			`s?://(?:www\.)?chirb\.it/(?:(?:wp|pl)/|fb_chirbit_player\.swf\?key=)?(?:[\da-zA-` +
			`Z]+))|(?:https?://(?:www\.)?cinemax\.com/(?:[^/]+/video/[0-9a-z-]+-(?:\d+)))|(?:` +
			`https?://(?:www\.)?clippituser\.tv/c/(?:[a-z]+))|(?:https?://(?:www\.)?clip\.rs/` +
			`(?:[^/]+)/\d+)|(?:https?://(?:chic|www)\.clipsyndicate\.com/video/play(list/\d+)` +
			`?/(?:\d+))|(?:https?://(?:www\.)?closertotruth\.com/(?:[^/]+/)*(?:[^/?#&]+))|(?:` +
			`https?://(?:www\.)?clyp\.it/(?:[a-z0-9]+))|(?:https?://(?:(?:edition|www|money)\` +
			`.)?cnn\.com/(?:video/(?:data/.+?|\?)/)?videos?/(?:.+?/(?:[^/]+?)(?:\.(?:[a-z\-]+` +
			`)|(?=&))))|(?:https?://[^\.]+\.blogs\.cnn\.com/.+)|(?:https?://(?:(?:edition|www` +
			`)\.)?cnn\.com/(?!videos?/))|(?:(?:coub:|https?://(?:coub\.com/(?:view|embed|coub` +
			`s)/|c-cdn\.coub\.com/fb-player\.swf\?.*\bcoub(?:ID|id)=))(?:[\da-z]+))|(?:https?` +
			`://(?:video|www|player(?:-backend)?)\.(?:allure|architecturaldigest|arstechnica|` +
			`bonappetit|brides|cnevids|cntraveler|details|epicurious|glamour|golfdigest|gq|ne` +
			`wyorker|self|teenvogue|vanityfair|vogue|wired|wmagazine)\.com/(?:(?:embed(?:js)?` +
			`|(?:script|inline)/video)/(?:[0-9a-f]{24})(?:/(?:[0-9a-f]{24}))?(?:.+?\btarget=(` +
			`?:[^&]+))?|(?:watch|series|video)/(?:[^/?#]+)))|(?:https?://(?:www\.)?contv\.com` +
			`/details-movie/(?:[^/]+))|(?:https?://(?:www\.)?cracked\.com/video_(?:\d+)_[\da-` +
			`z-]+\.html)|(?:https?://news\.cts\.com\.tw/[a-z]+/[a-z]+/\d+/(?:\d+)\.html)|(?:h` +
			`ttps?://(?:app\.)?curiositystream\.com/video/(?:\d+))|(?:https?://(?:www\.)?dail` +
			`ymail\.co\.uk/(?:video/[^/]+/video-|embed/video/)(?:[0-9]+))|(?:(?i)https?://(?:` +
			`(?:(?:www|touch)\.)?dailymotion\.[a-z]{2,3}/(?:(?:(?:embed|swf|\#)/)?video|swf)|` +
			`(?:www\.)?lequipe\.fr/video)/(?:[^/?_]+)(?:.+?\bplaylist=(?:x[0-9a-z]+))?)|(?:ht` +
			`tps?://(?:(?:m\.)?tvpot\.daum\.net/v/|videofarm\.daum\.net/controller/player/Vod` +
			`Player\.swf\?vid=)(?:[^?#&]+))|(?:https?://(?:m\.)?tvpot\.daum\.net/(?:clip/Clip` +
			`View.(?:do|tv)|mypot/View.do)\?.*?clipid=(?:\d+))|(?:https?://(?:m\.)?tvpot\.dau` +
			`m\.net/mypot/(?:View\.do|Top\.tv)\?.*?playlistid=(?:[0-9]+))|(?:https?://(?:m\.)` +
			`?tvpot\.daum\.net/mypot/(?:View|Top)\.(?:do|tv)\?.*?ownerid=(?:[0-9a-zA-Z]+))|(?` +
			`:https?://(?:www\.)?dctp\.tv/(?:#/)?filme/(?:[^/?#&]+))|(?:https?://(?:www\.)?do` +
			`tsub\.com/view/(?:[^/]+))|(?:https?://(?:(?:www|m)\.)?drtuber\.com/(?:video|embe` +
			`d)/(?:\d+)(?:/(?:[\w-]+))?)|(?:https?://video\.aktualne\.cz/(?:[^/]+/)+r~(?:[0-9` +
			`a-f]{32}))|(?:(?:https?)://(?:(?:www|legacy)\.)?dumpert\.nl/(?:mediabase|embed|i` +
			`tem)/(?:[0-9]+[/_][0-9a-zA-Z]+))|(?:https?://(?:(?:[^/]+\.)?(?:disney\.[a-z]{2,3` +
			`}(?:\.[a-z]{2})?|disney(?:(?:me|latino)\.com|turkiye\.com\.tr|channel\.de)|(?:st` +
			`arwars|marvelkids)\.com))/(?:(?:embed/|(?:[^/]+/)+[\w-]+-)(?:[a-z0-9]{24})|(?:[^` +
			`/]+/)?(?:[^/?#]+)))|(?:https?://(?:s?evt\.dispeak|events\.digitallyspeaking)\.co` +
			`m/(?:[^/]+/)+xml/(?:[^.]+)\.xml)|(?:https?://(?:www\.)?dw\.com/(?:[^/]+/)+(?:av|` +
			`e)-(?:\d+))|(?:(?:eagleplatform:(?:[^/]+):|https?://(?:.+?\.media\.eagleplatform` +
			`\.com)/index/player\?.*\brecord_id=)(?:\d+))|(?:https?://(?:www\.)?ebaumsworld\.` +
			`com/videos/[^/]+/(?:\d+))|(?:https?://(?:www\.)?echo\.msk\.ru/sounds/(?:\d+))|(?` +
			`:https?://(?:[^.]+\.)?elpais\.com/.*/(?:[^/#?]+)\.html(?:$|[?#]))|(?:https?://(?` +
			`:www\.)?engadget\.com/video/(?:[^/?#]+))|(?:https?://(?:www\.)?eporner\.com/(?:h` +
			`d-porn|embed)/(?:\w+)(?:/(?:[\w-]+))?)|(?:https?://?(?:(?:www|v1)\.)?escapistmag` +
			`azine\.com/videos/view/[^/]+/(?:[0-9]+))|(?:https?://(?:www\.)?fivethirtyeight\.` +
			`com/features/(?:[^/?#]+))|(?:https?://(?:www\.)?extremetube\.com/(?:[^/]+/)?vide` +
			`o/(?:[^/#?&]+))|(?:(?:https?://(?:[\w-]+\.)?(?:facebook\.com|facebookcorewwwi\.o` +
			`nion)/(?:[^#]*?\#!/)?(?:(?:video/video\.php|photo\.php|video\.php|video/embed|st` +
			`ory\.php)\?(?:.*?)(?:v|video_id|story_fbid)=|[^/]+/videos/(?:[^/]+/)?|[^/]+/post` +
			`s/|groups/[^/]+/permalink/)|facebook:)(?:[0-9]+))|(?:https?://(?:[\w-]+\.)?faceb` +
			`ook\.com/plugins/video\.php\?.*?\bhref=(?:https.+))|(?:https?://(?:www\.)?faz\.n` +
			`et/(?:[^/]+/)*.*?-(?:\d+)\.html)|(?:https?://(?:www\.)?fc-zenit\.ru/video/(?:[0-` +
			`9]+))|(?:https?://(?:www\.)?filmweb\.no/(?:trailere|filmnytt)/article(?:\d+)\.ec` +
			`e)|(?:(?:5min:|https?://(?:[^/]*?5min\.com/|delivery\.vidible\.tv/aol)(?:(?:Scri` +
			`pts/PlayerSeed\.js|playerseed/?)?\?.*?playList=)?)(?:\d+))|(?:https?://(?:www\.)` +
			`?5-tv\.ru/(?:(?:[^/]+/)+(?:\d+)|(?:[^/?#]+)(?:[/?#])?))|(?:https?://(?:www\.|sec` +
			`ure\.)?flickr\.com/photos/[\w\-_@]+/(?:\d+))|(?:https?://(?:www\.)?formula1\.com` +
			`/(?:content/fom-website/)?en/video/\d{4}/\d{1,2}/(?:.+?)\.html)|(?:https?://(?:v` +
			`ideo\.(?:insider\.)?fox(?:news|business)\.com)/v/(?:video-embed\.html\?video_id=` +
			`)?(?:\d+))|(?:https?://(?:www\.)?(?:insider\.)?foxnews\.com/(?!v)([^/]+/)+(?:[a-` +
			`z-]+))|(?:https?://(?:www\.)?franceculture\.fr/emissions/(?:[^/]+/)*(?:[^/?#&]+)` +
			`)|(?:https?://(?:www\.)?franceinter\.fr/emissions/(?:[^?#]+))|(?:https?://genera` +
			`tion-what\.francetv\.fr/[^/]+/video/(?:[^/?#&]+))|(?:https?://(?:www\.)?freesoun` +
			`d\.org/people/[^/]+/sounds/(?:[^/]+))|(?:https?://(?:www\.)?freespeech\.org/stor` +
			`ies/(?:.+))|(?:https?://(?:www\.)?funk\.net/(?:channel|playlist)/[^/]+/(?:[0-9a-` +
			`z-]+)-(?:\d+))|(?:https?://(?:www\.)?gameinformer\.com/(?:[^/]+/)*(?:[^.?&#]+))|` +
			`(?:https?://(?:www\.)?gaskrank\.tv/tv/(?:[^/]+)/(?:[^/]+)\.htm)|(?:(?:https?://(` +
			`?:www\.)?gazeta\.ru/(?:[^/]+/)?video/(?:main/)*(?:\d{4}/\d{2}/\d{2}/)?(?:[A-Za-z` +
			`0-9-_.]+)\.s?html))|(?:https?://(?:(?:www|giant|thumbs)\.)?gfycat\.com/(?:ru/|if` +
			`r/|gifs/detail/)?(?:[^-/?#\.]+))|(?:https?://(?:www\.)?giantbomb\.com/(?:videos|` +
			`shows)/(?:[^/]+)/(?:\d+-\d+))|(?:https?://(?:www\.)?godtube\.com/watch/\?v=(?:[\` +
			`da-zA-Z]+))|(?:^https?://video\.golem\.de/.+?/(?:.+?)/)|(?:https?://(?:(?:docs|d` +
			`rive)\.google\.com/(?:(?:uc|open)\?.*?id=|file/d/)|video\.google\.com/get_player` +
			`\?.*?docid=)(?:[a-zA-Z0-9_-]{28,}))|(?:https?://on-demand\.gputechconf\.com/gtc/` +
			`2015/video/S(?:\d+)\.html)|(?:https?://(?:www\.)?hbo\.com/(?:video|embed)(?:/[^/` +
			`]+)*/(?:[^/?#]+))|(?:https?://(?:www\.)?heise\.de/(?:[^/]+/)+[^/]+-(?:[0-9]+)\.h` +
			`tml)|(?:https?://(?:www\.)?hellporno\.(?:com/videos|net/v)/(?:[^/]+))|(?:https?:` +
			`//(?:www\.)?historicfilms\.com/(?:tapes/|play)(?:\d+))|(?:https?://(?:www\.)?hit` +
			`record\.org/records/(?:\d+))|(?:https?://(?:www\.)?hypem\.com/track/(?:[0-9a-z]{` +
			`5}))|(?:https?://.+?\.ign\.com/(?:[^/]+/)?(?:videos|show_videos|articles|feature` +
			`|(?:[^/]+/\d+/video))(/.+)?/(?:.+))|(?:https?://(?:www|m)\.imdb\.com/(?:video|ti` +
			`tle|list).*?[/-]vi(?:\d+))|(?:https?://(?:i\.)?imgur\.com/(?!(?:a|gallery|(?:t(?` +
			`:opic)?|r)/[^/]+)/)(?:[a-zA-Z0-9]+))|(?:https?://(?:i\.)?imgur\.com/(?:gallery|(` +
			`?:t(?:opic)?|r)/[^/]+)/(?:[a-zA-Z0-9]+))|(?:https?://(?:www\.)?ina\.fr/(?:video|` +
			`audio)/(?:[A-Z0-9_]+))|(?:https?://(?:www\.)?infoq\.com/(?:[^/]+/)+(?:[^/]+))|(?` +
			`:https?://(?:www\.)?ivideon\.com/tv/(?:[^/]+/)*camera/(?:\d+-[\da-f]+)/(?:\d+))|` +
			`(?:https?://(?:(?:www|m)\.)?izlesene\.com/(?:video|embedplayer)/(?:[^/]+/)?(?:[0` +
			`-9]+))|(?:https?://(?:licensing\.jamendo\.com/[^/]+|(?:www\.)?jamendo\.com)/trac` +
			`k/(?:[0-9]+)(?:/(?:[^/?#&]+))?)|(?:https?://.*?\.jeuxvideo\.com/.*/(.*?)\.htm)|(` +
			`?:(?:joj:|https?://media\.joj\.sk/embed/)(?:[^/?#^]+))|(?:(?:https?://(?:content` +
			`\.jwplatform|cdn\.jwplayer)\.com/(?:(?:feed|player|thumb|preview)s|jw6|v2/media)` +
			`/|jwplatform:)(?:[a-zA-Z0-9]{8}))|(?:https?://(?:play-)?tv\.kakao\.com/(?:channe` +
			`l/\d+|embed/player)/cliplink/(?:\d+|[^?#&]+@my))|(?:https?://(?:www\.)?keezmovie` +
			`s\.com/video/(?:(?:[^/]+)-)?(?:\d+))|(?:^https?://(?:(?:www|api)\.)?khanacademy\` +
			`.org/(?:[^/]+)/(?:[^/]+/){,2}(?:[^?#/]+)(?:$|[?#]))|(?:https?://(?:www\.)?kickst` +
			`arter\.com/projects/(?:[^/]*)/.*)|(?:https?://(?:www\.)?lenta\.ru/[^/]+/\d+/\d+/` +
			`\d+/(?:[^/?#&]+))|(?:https?://(?:www\.)?loc\.gov/(?:item/|today/cyberlc/feature_` +
			`wdesc\.php\?.*\brec=)(?:[0-9a-z_.]+))|(?:(?:https?://html5-player\.libsyn\.com/e` +
			`mbed/episode/id/(?:[0-9]+)))|(?:(?:limelight:media:|https?://(?:link\.videoplatf` +
			`orm\.limelight\.com/media/|assets\.delvenetworks\.com/player/loader\.swf)\?.*?\b` +
			`mediaId=)(?:[a-z0-9]{32}))|(?:https?://(?:new\.)?livestream\.com/(?:accounts/(?:` +
			`\d+)|(?:[^/]+))/(?:events/(?:\d+)|(?:[^/]+))(?:/videos/(?:\d+))?)|(?:https?://(?` +
			`:www\.)?lovehomeporn\.com/video/(?:\d+)(?:/(?:[^/?#&]+))?)|(?:https?://(?:www\.)` +
			`?(?:lynda\.com|educourse\.ga)/(?:(?:[^/]+/){2,3}(?:\d+)|player/embed)/(?:\d+))|(` +
			`?:https?://my\.mail\.ru/+music/+songs/+[^/?#&]+-(?:[\da-f]+))|(?:(?i)https?://(?` +
			`:www\.)?manyvids\.com/video/(?:\d+))|(?:(?:mediaset:|https?://(?:(?:www|static3)` +
			`\.)?mediasetplay\.mediaset\.it/(?:(?:video|on-demand)/(?:[^/]+/)+[^/]+_|player/i` +
			`ndex\.html\?.*?\bprogramGuid=))(?:[0-9A-Z]{16,}))|(?:https://player\.megaphone\.` +
			`fm/(?:[A-Z0-9]+))|(?:https?://(?:www\.)?metacritic\.com/.+?/trailers/(?:\d+))|(?` +
			`:(?:mva:|https?://(?:mva\.microsoft|(?:www\.)?microsoftvirtualacademy)\.com/[^/]` +
			`+/training-courses/[^/?#&]+-)(?:\d+)(?::|\?l=)(?:[\da-zA-Z]+_\d+))|(?:https?://(` +
			`?:[\da-z_-]+\.)*(?:mlb)\.com/(?:(?:(?:[^/]+/)*c-|(?:shared/video/embed/(?:embed|` +
			`m-internal-embed)\.html|(?:[^/]+/)+(?:play|index)\.jsp|)\?.*?\bcontent_id=)(?:\d` +
			`+)))|(?:https?://(?:www\.)?mojvideo\.com/video-(?:[^/]+)/(?:[a-f0-9]+))|(?:https` +
			`?://myspace\.com/[^/]+/(?:video/[^/]+/(?:\d+)|music/song/[^/?#&]+-(?:\d+)-\d+(?:` +
			`[/?#&]|$)))|(?:https?://(?:www\.)?myspass\.de/([^/]+/)*(?:\d+))|(?:https?://vide` +
			`o\.nationalgeographic\.com/.*?)|(?:https?://(?:m\.)?tv(?:cast)?\.naver\.com/(?:v` +
			`|embed)/(?:\d+))|(?:https?://(?:watch\.|www\.)?nba\.com/(?:(?:[^/]+/)+(?:[^?]*?)` +
			`)/?(?:/index\.html)?(?:\?.*)?$)|(?:https?://(?:www\.)?csnne\.com/video/(?:[0-9a-` +
			`z-]+))|(?:https?://(?:www\.)?(?:nbcnews|today|msnbc)\.com/([^/]+/)*(?:.*-)?(?:[^` +
			`/?]+))|(?:https?://vplayer\.nbcsports\.com/(?:[^/]+/)+(?:[0-9a-zA-Z_]+))|(?:http` +
			`s?://(?:www\.)?ndr\.de/(?:[^/]+/)*(?:[^/?#]+),[\da-z]+\.html)|(?:https?://(?:www` +
			`\.)?ndr\.de/(?:[^/]+/)*(?:[\da-z]+)-(?:player|externalPlayer)\.html)|(?:https?:/` +
			`/(?:www\.)?n-joy\.de/(?:[^/]+/)*(?:[\da-z]+)-(?:player|externalPlayer)_[^/]+\.ht` +
			`ml)|(?:https?://(?:[^/]+\.)?ndtv\.com/(?:[^/]+/)*videos?/?(?:[^/]+/)*[^/?^&]+-(?` +
			`:\d+))|(?:https?://(?:www\.)?netzkino\.de/\#!/(?:[^/]+)/(?:[^/]+))|(?:https?://(` +
			`?:www\.)?newgrounds\.com/(?:audio/listen|portal/view)/(?:[0-9]+))|(?:(?:https?:/` +
			`/api\.nexx(?:\.cloud|cdn\.com)/v3/(?:\d+)/videos/byid/|nexx:(?:(?:\d+):)?|https?` +
			`://arc\.nexx\.cloud/api/video/)(?:\d+))|(?:https?://embed\.nexx(?:\.cloud|cdn\.c` +
			`om)/\d+/(?:video/)?(?:[^/?#&]+))|(?:https?://(?:www\.)?(?:nhl|wch2016)\.com/(?:[` +
			`^/]+/)*c-(?:\d+))|(?:https?://(?:www\.)?nonktube\.com/(?:(?:video|embed)/|media/` +
			`nuevo/embed\.php\?.*?\bid=)(?:\d+))|(?:https?://(?:[^.]+\.)?(?:tv(?:noviny)?|tn|` +
			`novaplus|vymena|fanda|krasna|doma|prask)\.nova\.cz/(?:[^/]+/)+(?:[^/]+?)(?:\.htm` +
			`l|/|$))|(?:https?://(?:(?:www|cn)\.)?nowness\.com/(?:story|(?:series|category)/[` +
			`^/]+)/(?:[^/]+?)(?:$|[?#]))|(?:https?://(?:www\.)?ntv\.ru/(?:[^/]+/)*(?:[^/?#&]+` +
			`))|(?:https?://(?:(?:www|m|mobile)\.)?(?:odnoklassniki|ok)\.ru/(?:video(?:embed)` +
			`?/|web-api/video/moviePlayer/|live/|dk\?.*?st\.mvId=)(?:[\d-]+))|(?:https?://(?:` +
			`www\.)?onionstudios\.com/(?:video(?:s/[^/]+-|/)|embed\?.*\bid=)(?:\d+)(?!-))|(?:` +
			`https?://(?:www\.)?outsidetv\.com/(?:[^/]+/)*?play/[a-zA-Z0-9]{8}/\d+/\d+/(?:[a-` +
			`zA-Z0-9]{8}))|(?:https?://(?:(?:www\.)?packtpub\.com/mapt|subscription\.packtpub` +
			`\.com)/video/[^/]+/(?:\d+)/(?:[^/]+)/(?:[^/]+)(?:/(?:[^/?&#]+))?)|(?:https?://(?` +
			`:(?:www\.)?pandora\.tv/view/(?:[^/]+)/(?:\d+)|(?:.+?\.)?channel\.pandora\.tv/cha` +
			`nnel/video\.ptv\?|m\.pandora\.tv/?\?))|(?:https?://(?:(?:(?:video|www|player)\.p` +
			`bs\.org|video\.aptv\.org|video\.gpb\.org|video\.mpbonline\.org|video\.wnpt\.org|` +
			`video\.wfsu\.org|video\.wsre\.org|video\.wtcitv\.org|video\.pba\.org|video\.alas` +
			`kapublic\.org|video\.azpbs\.org|portal\.knme\.org|video\.vegaspbs\.org|watch\.ae` +
			`tn\.org|video\.ket\.org|video\.wkno\.org|video\.lpb\.org|videos\.oeta\.tv|video\` +
			`.optv\.org|watch\.wsiu\.org|video\.keet\.org|pbs\.kixe\.org|video\.kpbs\.org|vid` +
			`eo\.kqed\.org|vids\.kvie\.org|video\.pbssocal\.org|video\.valleypbs\.org|video\.` +
			`cptv\.org|watch\.knpb\.org|video\.soptv\.org|video\.rmpbs\.org|video\.kenw\.org|` +
			`video\.kued\.org|video\.wyomingpbs\.org|video\.cpt12\.org|video\.kbyueleven\.org` +
			`|video\.thirteen\.org|video\.wgbh\.org|video\.wgby\.org|watch\.njtvonline\.org|w` +
			`atch\.wliw\.org|video\.mpt\.tv|watch\.weta\.org|video\.whyy\.org|video\.wlvt\.or` +
			`g|video\.wvpt\.net|video\.whut\.org|video\.wedu\.org|video\.wgcu\.org|video\.wpb` +
			`t2\.org|video\.wucftv\.org|video\.wuft\.org|watch\.wxel\.org|video\.wlrn\.org|vi` +
			`deo\.wusf\.usf\.edu|video\.scetv\.org|video\.unctv\.org|video\.pbshawaii\.org|vi` +
			`deo\.idahoptv\.org|video\.ksps\.org|watch\.opb\.org|watch\.nwptv\.org|video\.wil` +
			`l\.illinois\.edu|video\.networkknowledge\.tv|video\.wttw\.com|video\.iptv\.org|v` +
			`ideo\.ninenet\.org|video\.wfwa\.org|video\.wfyi\.org|video\.mptv\.org|video\.wni` +
			`n\.org|video\.wnit\.org|video\.wpt\.org|video\.wvut\.org|video\.weiu\.net|video\` +
			`.wqpt\.org|video\.wycc\.org|video\.wipb\.org|video\.indianapublicmedia\.org|watc` +
			`h\.cetconnect\.org|video\.thinktv\.org|video\.wbgu\.org|video\.wgvu\.org|video\.` +
			`netnebraska\.org|video\.pioneer\.org|watch\.sdpb\.org|video\.tpt\.org|watch\.ksm` +
			`q\.org|watch\.kpts\.org|watch\.ktwu\.org|watch\.easttennesseepbs\.org|video\.wct` +
			`e\.tv|video\.wljt\.org|video\.wosu\.org|video\.woub\.org|video\.wvpublic\.org|vi` +
			`deo\.wkyupbs\.org|video\.kera\.org|video\.mpbn\.net|video\.mountainlake\.org|vid` +
			`eo\.nhptv\.org|video\.vpt\.org|video\.witf\.org|watch\.wqed\.org|video\.wmht\.or` +
			`g|video\.deltabroadcasting\.org|video\.dptv\.org|video\.wcmu\.org|video\.wkar\.o` +
			`rg|wnmuvideo\.nmu\.edu|video\.wdse\.org|video\.wgte\.org|video\.lptv\.org|video\` +
			`.kmos\.org|watch\.montanapbs\.org|video\.krwg\.org|video\.kacvtv\.org|video\.kco` +
			`stv\.org|video\.wcny\.org|video\.wned\.org|watch\.wpbstv\.org|video\.wskg\.org|v` +
			`ideo\.wxxi\.org|video\.wpsu\.org|on-demand\.wvia\.org|video\.wtvi\.org|video\.we` +
			`sternreservepublicmedia\.org|video\.ideastream\.org|video\.kcts9\.org|video\.bas` +
			`inpbs\.org|video\.houstonpbs\.org|video\.klrn\.org|video\.klru\.tv|video\.wtjx\.` +
			`org|video\.ideastations\.org|video\.kbtc\.org)/(?:(?:vir|port)alplayer|video)/(?` +
			`:[0-9]+)(?:[?/]|$)|(?:www\.)?pbs\.org/(?:[^/]+/){1,5}(?:[^/]+?)(?:\.html)?/?(?:$` +
			`|[?\#])|(?:video|player)\.pbs\.org/(?:widget/)?partnerplayer/(?:[^/]+)/))|(?:htt` +
			`ps?://(?:www\.)?pearvideo\.com/video_(?:\d+))|(?:https?://(?:www\.)?people\.com/` +
			`people/videos/0,,(?:\d+),00\.html)|(?:https?://player\.performgroup\.com/eplayer` +
			`(?:/eplayer\.html|\.js)#/?(?:[0-9a-f]{26})\.(?:[0-9a-z]{26}))|(?:https?://(?:[a-` +
			`z0-9]+\.)?photobucket\.com/.*(([\?\&]current=)|_)(?:.*)\.(?:(flv)|(mp4)))|(?:htt` +
			`ps?://player\.piksel\.com/v/(?:refid/[^/]+/prefid/)?(?:[a-z0-9_]+))|(?:https?://` +
			`(?:www\.)?play\.fm/(?:(?:[^/]+/)+(?:[^/]+))/?(?:$|[?#]))|(?:https?://(?:www\.)?p` +
			`layvid\.com/watch(\?v=|/)(?:.+?)(?:#|$))|(?:(?:https?)://(?:(?:[^.]+)\.podomatic` +
			`\.com/entry|(?:www\.)?podomatic\.com/podcasts/(?:[^/]+)/episodes)/(?:[^/?#&]+))|` +
			`(?:https?://(?:www\.)?pokemon\.com/[a-z]{2}(?:.*?play=(?:[a-z0-9]{32})|/(?:[^/]+` +
			`/)+(?:[^/?#&]+)))|(?:https?://(?:www\.)?pornhd\.com/(?:[a-z]{2,4}/)?videos/(?:\d` +
			`+)(?:/(?:.+))?)|(?:https?://(?:(?:[^/]+\.)?(?:pornhub(?:premium)?\.(?:com|net))/` +
			`(?:(?:view_video\.php|video/show)\?viewkey=|embed/)|(?:www\.)?thumbzilla\.com/vi` +
			`deo/)(?:[\da-z]+))|(?:https?://(?:\w+\.)?pornotube\.com/(?:[^?#]*?)/video/(?:[0-` +
			`9]+))|(?:https?://(?:www\.)?radiojavan\.com/videos/video/(?:[^/]+)/?)|(?:https?:` +
			`//[^/]+\.(?:rai\.(?:it|tv)|rainews\.it)/.+?-(?:[\da-f]{8}-[\da-f]{4}-[\da-f]{4}-` +
			`[\da-f]{4}-[\da-f]{12})(?:-.+?)?\.html)|(?:https?://(?:videos\.raywenderlich\.co` +
			`m/courses|(?:www\.)?raywenderlich\.com)/(?:[^/]+)/lessons/(?:\d+))|(?:https?://(` +
			`?:(?:www\.)?redtube\.com/|embed\.redtube\.com/\?.*?\bid=)(?:[0-9]+))|(?:^https?:` +
			`//(?:www\.)?reverbnation\.com/.*?/song/(?:\d+).*?$)|(?:https?://(?:www\.)?rockst` +
			`argames\.com/videos(?:/video/|#?/?\?.*\bvideo=)(?:\d+))|(?:https?://(?:www\.)?pr` +
			`ehravac\.rozhlas\.cz/audio/(?:[0-9]+))|(?:rts:(?:\d+)|https?://(?:.+?\.)?rts\.ch` +
			`/(?:[^/]+/){2,}(?:[0-9]+)-(?:.+?)\.html)|(?:https?://(?:test)?player\.(?:rutv\.r` +
			`u|vgtrk\.com)/(?:flash\d+v/container\.swf\?id=|iframe/(?:swf|video|live)/id/|ind` +
			`ex/iframe/cast_id/)(?:\d+))|(?:https?://(?:www\.)?(?:ruutu|supla)\.fi/(?:video|s` +
			`upla)/(?:\d+))|(?:https?://(?:www\.)?(?:safaribooksonline|(?:learning\.)?oreilly` +
			`)\.com/(?:library/view/[^/]+/(?:[^/]+)/(?:[^/?\#&]+)\.html|videos/[^/]+/[^/]+/(?` +
			`:[^-]+-[^/?\#&]+)))|(?:https?://(?:(?:v2|www)\.)?videos\.sapo\.(?:pt|cv|ao|mz|tl` +
			`)/(?:[\da-zA-Z]{20}))|(?:https?://(?:www\.)?sbs\.com\.au/(?:ondemand|news)/video` +
			`/(?:single/)?(?:[0-9]+))|(?:https?://(?:www\.)?screencast\.com/t/(?:[a-zA-Z0-9]+` +
			`))|(?:https?://screencast-o-matic\.com/watch/(?:[0-9a-zA-Z]+))|(?:https?://(?:ww` +
			`w\.)?senate\.gov/isvp/?\?(?:.+))|(?:https?://(?:www\.)?seznamzpravy\.cz/iframe/p` +
			`layer\?.*\bsrc=)|(?:https?://(?:.*?\.)?video\.sina\.com\.cn/(?:(?:view/|.*\#)(?:` +
			`\d+)|.+?/(?:[^/?#]+)(?:\.s?html)|api/sinawebApi/outplay.php/(?:.+?)\.swf))|(?:ht` +
			`tps?://news\.sky\.com/video/[0-9a-z-]+-(?:[0-9]+))|(?:https?://(?:www\.)?slidesh` +
			`are\.net/[^/]+?/(?:.+?)($|\?))|(?:https?://slideslive\.com/(?:[0-9]+))|(?:https?` +
			`://(?:www\.)?sonyliv\.com/details/[^/]+/(?:\d+))|(?:https?://(?:www\.)?soundgasm` +
			`\.net/u/(?:[0-9a-zA-Z_-]+)/(?:[0-9a-zA-Z_-]+))|(?:https?://(?:[^/]+\.)?spankbang` +
			`\.com/(?:[\da-z]+)/(?:video|play|embed)\b)|(?:https?://(?:www\.)?spankwire\.com/` +
			`(?:[^/]+/video|EmbedPlayer\.aspx/?\?.*?\bArticleId=)(?:\d+))|(?:https?://(?:www\` +
			`.)?stitcher\.com/podcast/(?:[^/]+/)+e/(?:(?:[^/#?&]+?)-)?(?:\d+)(?:[/#?&]|$))|(?` +
			`:https?://(?:(?:www|play)\.)?(?:srf|rts|rsi|rtr|swissinfo)\.ch/play/(?:tv|radio)` +
			`/(?:[^/]+/(?:video|audio)/[^?]+|popup(?:video|audio)player)\?id=(?:[0-9a-f\-]{36` +
			`}|\d+))|(?:https?://openclassroom\.stanford\.edu(?:/?|(/MainFolder/(?:HomePage|C` +
			`oursePage|VideoPage)\.php([?]course=(?:[^&]+)(&video=(?:[^&]+))?(&.*)?)?))$)|(?:` +
			`https?://streamable\.com/(?:[es]/)?(?:\w+))|(?:https?://portal\.stretchinternet\` +
			`.com/[^/]+/(?:portal|full)\.htm\?.*?\beventId=(?:\d+))|(?:https?://(?:www\.)?sve` +
			`rigesradio\.se/(?:sida/)?avsnitt/(?:[0-9]+))|(?:https?://(?:www\.)?sverigesradio` +
			`\.se/sida/(?:artikel|gruppsida)\.aspx\?.*?\bartikel=(?:[0-9]+))|(?:https?://(?:w` +
			`ww\.)?tagesschau\.de/(?:[^/]+/(?:[^/]+/)*?(?:[^/#?]+?(?:-?[0-9]+)?))(?:~_?[^/#?]` +
			`+?)?\.html)|(?:https?://(?:www\.)?tastytrade\.com/tt/shows/[^/]+/episodes/(?:[^/` +
			`?#&]+))|(?:https?://tds\.lifeway\.com/v1/trainingdeliverysystem/courses/(?:\d+)/` +
			`index\.html)|(?:https?://(?:\w+\.)?teamcoco\.com/(?:([^/]+/)*[^/?#]+))|(?:https?` +
			`://(?:www\.)?teamtreehouse\.com/library/(?:[^/]+))|(?:(?:https?://)(?:www|embed(` +
			`?:-ssl)?)(?:\.ted\.com/((?:playlists(?:/(?:\d+))?)|((?:talks))|(?:watch)/[^/]+/[` +
			`^/]+)(/lang/(.*?))?/(?:[\w-]+).*)$)|(?:https?://(?:www\.)?telegraaf\.nl/video/(?` +
			`:\d+))|(?:https?://(?:www\.)?tenta\.com/how-to-download-videos)|(?:https?://(?:w` +
			`ww\.)?tfo\.org/(?:en|fr)/(?:[^/]+/){2}(?:\d+))|(?:https?://(?:www\.)?tmz\.com/vi` +
			`deos/(?:[^/?#]+))|(?:https?://player\.(?:tna|emp)flix\.com/video/(?:\d+))|(?:htt` +
			`ps?://(?:www\.)?tnaflix\.com/[^/]+/(?:[^/]+)/video(?:\d+))|(?:https?://(?:www\.)` +
			`?moviefap\.com/videos/(?:[0-9a-f]+)/(?:[^/]+)\.html)|(?:https?://(?:www\.)?toong` +
			`oggles\.com/shows/(?:\d+)(?:/[^/]+/episodes/(?:\d+))?)|(?:https?://videos\.toypi` +
			`cs\.net/view/(?:[0-9]+))|(?:https?://(?:www\.)?tube8\.com/(?:[^/]+/)+(?:[^/]+)/(` +
			`?:\d+))|(?:https?://(?:www\.)?tvanouvelles\.ca/videos/(?:\d+))|(?:https?://(?:ww` +
			`w\.)?tvc\.ru/video/iframe/id/(?:\d+))|(?:https?://(?:www\.)?tvc\.ru/(?!video/ifr` +
			`ame/id/)(?:[^?#]+))|(?:https?://(?:(?:[^/]+)\.)?tvn24(?:bis)?\.pl/(?:[^/]+/)*(?:` +
			`[^/]+))|(?:(?:tvp:|https?://[^/]+\.tvp\.(?:pl|info)/sess/tvplayer\.php\?.*?objec` +
			`t_id=)(?:\d+))|(?:https?://[^/]+\.tvp\.(?:pl|info)/(?:video/(?:[^,\s]*,)*|(?:(?!` +
			`\d+/)[^/]+/)*)(?:\d+))|(?:https?://(?:www\.)?20min\.ch/(?:videotv/*\?.*?\bvid=|v` +
			`ideoplayer/videoplayer\.html\?.*?\bvideoId@)(?:\d+))|(?:https?://video\.(?:twent` +
			`ythree\.net|23video\.com|filmweb\.no)/v\.ihtml/player\.html\?(?:.*?\bphoto(?:_|%` +
			`5f)id=(?:\d+).*))|(?:https?://(?:clips\.twitch\.tv/(?:embed\?.*?\bclip=|(?:[^/]+` +
			`/)*)|(?:(?:www|go|m)\.)?twitch\.tv/[^/]+/clip/)(?:[^/?#&]+))|(?:https?://amp\.tw` +
			`img\.com/v/(?:[0-9a-f\-]{36}))|(?:https?://video\.udn\.com/(?:embed|play)/news/(` +
			`?:\d+))|(?:https?://(?:www\.)?(?:digiteka\.net|ultimedia\.com)/(?:deliver/(?:gen` +
			`eric|musique)(?:/[^/]+)*/(?:src|article)|default/index/video(?:generic|music)/id` +
			`)/(?:[\d+a-z]+))|(?:https?://utv\.unistra\.fr/(?:index|video)\.php\?id_video\=(?` +
			`:\d+))|(?:https?://(?:.+?\.)?uol\.com\.br/.*?(?:(?:mediaId|v)=|view/(?:[a-z0-9]+` +
			`/)?|video(?:=|/(?:\d{4}/\d{2}/\d{2}/)?))(?:\d+|[\w-]+-[A-Z0-9]+))|(?:https?://(?` +
			`:(?:app|embed)\.)?ustudio\.com/embed/(?:[^/]+)/(?:[^/]+))|(?:https?://(?:[^/]+\.` +
			`)?vbox7\.com/(?:play:|(?:emb/external\.php|player/ext\.swf)\?.*?\bvid=)(?:[\da-f` +
			`A-F]+))|(?:https?://(?:www\.)?veoh\.com/(?:watch|embed|iphone/#_Watch)/(?:(?:v|e` +
			`|yapi-)[\da-zA-Z]+))|(?:https?://(?:.+?\.)?vesti\.ru/(?:.+))|(?:(?:https?://(?:w` +
			`ww\.)?vevo\.com/watch/(?!playlist|genre)(?:[^/]+/(?:[^/]+/)?)?|https?://cache\.v` +
			`evo\.com/m/html/embed\.html\?video=|https?://videoplayer\.vevo\.com/embed/embedd` +
			`ed\?videoId=|https?://embed\.vevo\.com/.*?[?&]isrc=|vevo:)(?:[^&?#]+))|(?:(?:htt` +
			`ps?://(?:www\.)?(?:vgtv.no|bt.no/tv|aftenbladet.no/tv|fvn.no/fvntv|aftenposten.n` +
			`o/webtv|ap.vgtv.no/webtv|tv.aftonbladet.se/abtv|www.aftonbladet.se/tv)/?(?:(?:\#` +
			`!/)?(?:video|live)/|embed?.*id=|a(?:rticles)?/)|(?:vgtv|bttv|satv|fvntv|aptv|abt` +
			`v):)(?:\d+))|(?:https://(?:www\.)?vice\.com/(?:[^/]+)/article/(?:[0-9a-z]{6}/)?(` +
			`?:[^?#]+))|(?:https?://(?:www\.)?viddler\.com/(?:v|embed|player)/(?:[a-z0-9]+)(?` +
			`:.+?\bsecret=(\d+))?)|(?:https?://videopress\.com/embed/(?:[\da-zA-Z]+))|(?:http` +
			`s?://(?:www\.|m\.|)videosz\.com/[a-z]+/[a-z]+/(?:[0-9]+)_)|(?:https?://(?:www\.)` +
			`?vidlii\.com/(?:watch|embed)\?.*?\bv=(?:[0-9A-Za-z_-]{11}))|(?:https?://(?:www\.` +
			`)?(?:(?:(?:main\.)?snagfilms|snagxtreme|funnyforfree|kiddovid|winnersview|(?:mon` +
			`umental|lax)sportsnetwork|vayafilm|failarmy|ftfnext|lnppass\.legapallacanestro|m` +
			`oviespree|app\.myoutdoortv|neoufitness|pflmma|theidentitytb)\.com|(?:hoichoi|app` +
			`\.horseandcountry|kronon|marquee|supercrosslive)\.tv)(?:(?:/(?:films/title|show|` +
			`(?:news/)?videos?|watch))?/(?:[^?#]+)))|(?:https?://(?:(?:www|embed)\.)?(?:(?:(?` +
			`:main\.)?snagfilms|snagxtreme|funnyforfree|kiddovid|winnersview|(?:monumental|la` +
			`x)sportsnetwork|vayafilm|failarmy|ftfnext|lnppass\.legapallacanestro|moviespree|` +
			`app\.myoutdoortv|neoufitness|pflmma|theidentitytb)\.com|(?:hoichoi|app\.horseand` +
			`country|kronon|marquee|supercrosslive)\.tv)/embed/player\?.*\bfilmId=(?:[\da-f]{` +
			`8}-(?:[\da-f]{4}-){3}[\da-f]{12}))|(?:https?://(?:(?:www|player)\.)?vimeo(?:pro)` +
			`?\.com/(?!(?:channels|album|showcase)/[^/?#]+/?(?:$|[?#])|[^/]+/review/|ondemand` +
			`/)(?:.*?/)?(?:(?:play_redirect_hls|moogaloop\.swf)\?clip_id=)?(?:videos?/)?(?:[0` +
			`-9]+)(?:/[\da-f]+)?/?(?:[?&].*)?(?:[#].*)?$)|(?:https?://(?:www\.)?vimeo\.com/on` +
			`demand/([^/]+/)?(?:[^/?#&]+))|(?:(?:https://vimeo\.com/[^/]+/review/(?:[^/]+)/[0` +
			`-9a-f]{10}))|(?:https?://(?:player\.vimple\.(?:ru|co)/iframe|vimple\.(?:ru|co))/` +
			`(?:[\da-f-]{32,36}))|(?:https?://(?:www\.)?vine\.co/(?:v|oembed)/(?:\w+))|(?:(?:` +
			`viqeo:|https?://cdn\.viqeo\.tv/embed/*\?.*?\bvid=|https?://api\.viqeo\.tv/v\d+/d` +
			`ata/startup?.*?\bvideo(?:%5B%5D|\[\])=)(?:[\da-f]+))|(?:https?://(?:(?:(?:(?:m|n` +
			`ew)\.)?vk\.com/video_|(?:www\.)?daxab.com/)ext\.php\?(?:.*?\boid=(?:-?\d+).*?\bi` +
			`d=(?:\d+).*)|(?:(?:(?:m|new)\.)?vk\.com/(?:.+?\?.*?z=)?video|(?:www\.)?daxab.com` +
			`/embed/)(?:-?\d+_\d+)(?:.*\blist=(?:[\da-f]+))?))|(?:https?://(?:(?:www|m)\.)?vl` +
			`ive\.tv/video/(?:[0-9]+))|(?:https?://(?:(?:www|m)\.)?vlive\.tv/video/(?:[0-9]+)` +
			`/playlist/(?:[0-9]+))|(?:https?://voicerepublic\.com/(?:talks|embed)/(?:[0-9a-z-` +
			`]+))|(?:https?://(?:(?:www|view)\.)?vzaar\.com/(?:videos/)?(?:\d+))|(?:https?://` +
			`m\.weibo\.cn/status/(?:[0-9]+)(\?.+)?)|(?:https?://(?:www|m)\.worldstar(?:candy|` +
			`hiphop)\.com/(?:videos|android)/video\.php\?.*?\bv=(?:[^&]+))|(?:(?:https?://vid` +
			`eo-api\.wsj\.com/api-video/player/iframe\.html\?.*?\bguid=|https?://(?:www\.)?(?` +
			`:wsj|barrons)\.com/video/(?:[^/]+/)+|wsj:)(?:[a-fA-F0-9-]{36}))|(?:(?i)https?://` +
			`(?:www\.)?wsj\.com/articles/(?:[^/?#&]+))|(?:https?://(?:.+?\.)?(?:xhamster\.(?:` +
			`com|one|desi)|xhms\.pro|xhamster[27]\.com)/(?:movies/(?:\d+)/(?:[^/]*)\.html|vid` +
			`eos/(?:[^/]*)-(?:\d+)))|(?:https?://(?:.+?\.)?(?:xhamster\.(?:com|one|desi)|xhms` +
			`\.pro|xhamster[27]\.com)/xembed\.php\?video=(?:\d+))|(?:https?://(?:video|www)\.` +
			`xnxx\.com/video-?(?:[0-9a-z]+)/)|(?:(?:xtube:|https?://(?:www\.)?xtube\.com/(?:w` +
			`atch\.php\?.*\bv=|video-watch/(?:embedded/)?(?:[^/]+)-))(?:[^/?&#]+))|(?:https?:` +
			`//vlog\.xuite\.net/(?:play|embed)/(?:(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|` +
			`[A-Za-z0-9+/]{3}=)?))|(?:https?://(?:(?:[^/]+\.)?xvideos2?\.com/video|(?:www\.)?` +
			`xvideos\.es/video|flashservice\.xvideos\.com/embedframe/|static-hw\.xvideos\.com` +
			`/swf/xv-player\.swf\?.*?\bid_video=)(?:[0-9]+))|(?:https?://(?:www\.)?xxxymovies` +
			`\.com/videos/(?:\d+)/(?:[^/]+))|(?:(?:https?://(?:(?:[a-zA-Z]{2}(?:-[a-zA-Z]{2})` +
			`?|malaysia)\.)?(?:[\da-zA-Z_-]+\.)?yahoo\.com/(?:[^/]+/)*(?:[^?&#]*-[0-9]+(?:-[a` +
			`-z]+)?)\.html))|(?:https?://v\.yinyuetai\.com/video(?:/h5)?/(?:[0-9]+))|(?:https` +
			`?://(?:\w+\.)?youjizz\.com/videos/(?:[^/#?]*-(?:\d+)\.html|embed/(?:\d+)))|(?:ht` +
			`tps?://(?:www\.)?youporn\.com/(?:watch|embed)/(?:\d+)(?:/(?:[^/?#&]+))?)|(?:http` +
			`s?://(?:www\.)?sxyprn\.com/post/(?:[^/?#&.]+))|(?:https?://(?:www\.)?(?:youruplo` +
			`ad\.com/(?:watch|embed)|embed\.yourupload\.com)/(?:[A-Za-z0-9]+))|(?:^((?:https?` +
			`://|//)(?:(?:(?:(?:\w+\.)?[yY][oO][uU][tT][uU][bB][eE](?:-nocookie|kids)?\.com/|` +
			`(?:www\.)?deturl\.com/www\.youtube\.com/|(?:www\.)?pwnyoutube\.com/|(?:www\.)?ho` +
			`oktube\.com/|(?:www\.)?yourepeat\.com/|tube\.majestyc\.net/|(?:(?:www|dev)\.)?in` +
			`vidio\.us/|(?:(?:www|no)\.)?invidiou\.sh/|(?:(?:www|fi|de)\.)?invidious\.snopyta` +
			`\.org/|(?:www\.)?invidious\.kabi\.tk/|(?:www\.)?invidious\.13ad\.de/|(?:www\.)?i` +
			`nvidious\.mastodon\.host/|(?:www\.)?invidious\.nixnet\.xyz/|(?:www\.)?invidious\` +
			`.drycat\.fr/|(?:www\.)?tube\.poal\.co/|(?:www\.)?vid\.wxzm\.sx/|(?:www\.)?yewtu\` +
			`.be/|(?:www\.)?yt\.elukerio\.org/|(?:www\.)?yt\.lelux\.fi/|(?:www\.)?invidious\.` +
			`ggc-project\.de/|(?:www\.)?yt\.maisputain\.ovh/|(?:www\.)?invidious\.13ad\.de/|(` +
			`?:www\.)?invidious\.toot\.koeln/|(?:www\.)?invidious\.fdn\.fr/|(?:www\.)?watch\.` +
			`nettohikari\.com/|(?:www\.)?kgg2m7yk5aybusll\.onion/|(?:www\.)?qklhadlycap4cnod\` +
			`.onion/|(?:www\.)?axqzx4s6s54s32yentfqojs3x5i7faxza6xo3ehd4bzzsg2ii4fv2iid\.onio` +
			`n/|(?:www\.)?c7hqkpkpemu6e7emz5b4vyz7idjgdvgaaa3dyimmeojqbgpea3xqjoid\.onion/|(?` +
			`:www\.)?fz253lmuao3strwbfbmx46yu7acac2jz27iwtorgmbqlkurlclmancad\.onion/|(?:www\` +
			`.)?invidious\.l4qlywnpwqsluw65ts7md3khrivpirse744un3x7mlskqauz5pyuzgqd\.onion/|(` +
			`?:www\.)?owxfohz4kjyv25fvlqilyxast7inivgiktls3th44jhk3ej3i7ya\.b32\.i2p/|(?:www\` +
			`.)?4l2dgddgsrkf2ous66i6seeyi6etzfgrue332grh2n7madpwopotugyd\.onion/|youtube\.goo` +
			`gleapis\.com/)(?:.*?\#/)?(?:(?:(?:v|embed|e)/(?!videoseries))|(?:(?:(?:watch|mov` +
			`ie)(?:_popup)?(?:\.php)?/?)?(?:\?|\#!?)(?:.*?[&;])??v=)))|(?:youtu\.be|vid\.plus` +
			`|zwearz\.com/watch|)/|(?:www\.)?cleanvideosearch\.com/media/action/yt/watch\?vid` +
			`eoId=))?([0-9A-Za-z_-]{11})(?!.*?\blist=(?:(?:PL|LL|EC|UU|FL|RD|UL|TL|PU|OLAK5uy` +
			`_)[0-9A-Za-z-_]{10,}|WL)))|(?:(?:(?:https?://)?(?:\w+\.)?(?:(?:youtube(?:kids)?\` +
			`.com|invidio\.us)/(?:(?:course|view_play_list|my_playlists|artist|playlist|watch` +
			`|embed/(?:videoseries|[0-9A-Za-z_-]{11}))\?(?:.*?[&;])*?(?:p|a|list)=|p/)|youtu\` +
			`.be/[0-9A-Za-z_-]{11}\?.*?\blist=)((?:PL|LL|EC|UU|FL|RD|UL|TL|PU|OLAK5uy_)?[0-9A` +
			`-Za-z-_]{10,}|(?:MC)[\w\.]*).*|((?:PL|LL|EC|UU|FL|RD|UL|TL|PU|OLAK5uy_)[0-9A-Za-` +
			`z-_]{10,})))`,
	}
}
