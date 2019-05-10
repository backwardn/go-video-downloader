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
 * xfileshare/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/xfileshare.py
 */

package xfileshare

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError      λ.Object
	InfoExtractor       λ.Object
	NO_DEFAULT          λ.Object
	XFileShareIE        λ.Object
	ϒdetermine_ext      λ.Object
	ϒint_or_none        λ.Object
	ϒurlencode_postdata λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		NO_DEFAULT = Ωutils.NO_DEFAULT
		ϒurlencode_postdata = Ωutils.ϒurlencode_postdata
		XFileShareIE = λ.Cal(λ.TypeType, λ.NewStr("XFileShareIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				XFileShareIE__SITES     λ.Object
				XFileShareIE__VALID_URL λ.Object
			)
			XFileShareIE__SITES = λ.NewTuple(
				λ.NewTuple(
					λ.NewStr("daclips\\.(?:in|com)"),
					λ.NewStr("DaClips"),
				),
				λ.NewTuple(
					λ.NewStr("filehoot\\.com"),
					λ.NewStr("FileHoot"),
				),
				λ.NewTuple(
					λ.NewStr("gorillavid\\.(?:in|com)"),
					λ.NewStr("GorillaVid"),
				),
				λ.NewTuple(
					λ.NewStr("movpod\\.in"),
					λ.NewStr("MovPod"),
				),
				λ.NewTuple(
					λ.NewStr("powerwatch\\.pw"),
					λ.NewStr("PowerWatch"),
				),
				λ.NewTuple(
					λ.NewStr("rapidvideo\\.ws"),
					λ.NewStr("Rapidvideo.ws"),
				),
				λ.NewTuple(
					λ.NewStr("thevideobee\\.to"),
					λ.NewStr("TheVideoBee"),
				),
				λ.NewTuple(
					λ.NewStr("vidto\\.(?:me|se)"),
					λ.NewStr("Vidto"),
				),
				λ.NewTuple(
					λ.NewStr("streamin\\.to"),
					λ.NewStr("Streamin.To"),
				),
				λ.NewTuple(
					λ.NewStr("xvidstage\\.com"),
					λ.NewStr("XVIDSTAGE"),
				),
				λ.NewTuple(
					λ.NewStr("vidabc\\.com"),
					λ.NewStr("Vid ABC"),
				),
				λ.NewTuple(
					λ.NewStr("vidbom\\.com"),
					λ.NewStr("VidBom"),
				),
				λ.NewTuple(
					λ.NewStr("vidlo\\.us"),
					λ.NewStr("vidlo"),
				),
				λ.NewTuple(
					λ.NewStr("rapidvideo\\.(?:cool|org)"),
					λ.NewStr("RapidVideo.TV"),
				),
				λ.NewTuple(
					λ.NewStr("fastvideo\\.me"),
					λ.NewStr("FastVideo.me"),
				),
			)
			XFileShareIE__VALID_URL = λ.Mod(λ.NewStr("https?://(?P<host>(?:www\\.)?(?:%s))/(?:embed-)?(?P<id>[0-9a-zA-Z]+)"), λ.Cal(λ.GetAttr(λ.NewStr("|"), "join", nil), λ.Cal(λ.NewFunction("<generator>",
				nil,
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
						var (
							ϒsite λ.Object
							τmp0  λ.Object
							τmp1  λ.Object
						)
						τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(λ.Cal(λ.ListType, λ.Cal(λ.ZipIteratorType, λ.Unpack(λ.AsStarred(XFileShareIE__SITES))...)), λ.NewInt(0)))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒsite = τmp1
							λgy.Yield(ϒsite)
						}
						return λ.None
					})
				}))))
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_SITES"):     XFileShareIE__SITES,
				λ.NewStr("_VALID_URL"): XFileShareIE__VALID_URL,
			})
		}())
	})
}
