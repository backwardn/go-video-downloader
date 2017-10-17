/**
 * Go Video Downloader
 *
 *    Copyright 2017 Tenta, LLC
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
 * xnxxie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/xnxx.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type XNXXIE struct {
	*rnt.CommonIE
}

func NewXNXXIE() rnt.InfoExtractor {
	ret := &XNXXIE{}
	ret.CommonIE = rnt.NewCommonIE()
	ret.VALIDURL = "https?://(?:video|www)\\.xnxx\\.com/video-?(?P<id>[0-9a-z]+)/"
	return ret
}

func (self *XNXXIE) Key() string {
	return "XNXX"
}

func (self *XNXXIE) Name() string {
	return "XNXX extractor"
}

func (self *XNXXIE) _real_extract(url string) map[string]interface{} {
	video_id := (self).MatchID(url)
	webpage := (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	video_url := (self).SearchRegexOne("flv_url=(.*?)&amp;", webpage, "video URL", rnt.NoDefault, true, 0, nil)
	video_url = rnt.AsOptString(rnt.ParseUnquote(video_url.Get()))
	video_title := (self).HTMLSearchRegexOne("<title>(.*?)\\s+-\\s+XNXX.COM", webpage, "title", rnt.NoDefault, true, 0, nil)
	video_thumbnail := (self).SearchRegexOne("url_bigthumb=(.*?)&amp;", webpage, "thumbnail", rnt.NoDefault, false, 0, nil)
	return map[string]interface{}{"id": video_id,
		"url":       video_url,
		"title":     video_title,
		"ext":       "flv",
		"thumbnail": video_thumbnail,
		"age_limit": 18}
}

func (self *XNXXIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory("XNXX", NewXNXXIE)
}
