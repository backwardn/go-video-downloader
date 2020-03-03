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
 * compat/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/compat.py
 */

package compat

import (
	Ωbase64 "github.com/tenta-browser/go-video-downloader/gen/base64"
	Ωentities "github.com/tenta-browser/go-video-downloader/gen/html/entities"
	Ωparser "github.com/tenta-browser/go-video-downloader/gen/html/parser"
	Ωerror "github.com/tenta-browser/go-video-downloader/gen/urllib/error"
	Ωparse "github.com/tenta-browser/go-video-downloader/gen/urllib/parse"
	ΩElementTree "github.com/tenta-browser/go-video-downloader/gen/xml/etree/ElementTree"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ϒcompat_HTMLParseError            λ.Object
	ϒcompat_HTMLParser                λ.Object
	ϒcompat_HTTPError                 λ.Object
	ϒcompat_b64decode                 λ.Object
	ϒcompat_basestring                λ.Object
	ϒcompat_chr                       λ.Object
	ϒcompat_etree_Element             λ.Object
	ϒcompat_etree_fromstring          λ.Object
	ϒcompat_html_entities_html5       λ.Object
	ϒcompat_http_client               λ.Object
	ϒcompat_integer_types             λ.Object
	ϒcompat_kwargs                    λ.Object
	ϒcompat_numeric_types             λ.Object
	ϒcompat_parse_qs                  λ.Object
	ϒcompat_str                       λ.Object
	ϒcompat_urllib_parse_unquote      λ.Object
	ϒcompat_urllib_parse_unquote_plus λ.Object
	ϒcompat_urllib_parse_urlencode    λ.Object
	ϒcompat_urllib_parse_urlparse     λ.Object
	ϒcompat_xml_parse_error           λ.Object
	ϒcompat_xpath                     λ.Object
)

func init() {
	λ.InitModule(func() {
		ϒcompat_etree_Element = ΩElementTree.Element
		ϒcompat_HTMLParser = Ωparser.HTMLParser
		ϒcompat_HTMLParseError = Ωparser.HTMLParseError
		ϒcompat_str = λ.StrType
		ϒcompat_basestring = λ.StrType
		ϒcompat_integer_types = λ.NewTuple(λ.IntType)
		ϒcompat_numeric_types = λ.NewTuple(
			λ.IntType,
			λ.FloatType,
		)
		ϒcompat_chr = λ.BuiltinChr
		ϒcompat_urllib_parse_urlparse = Ωparse.ϒurlparse
		ϒcompat_urllib_parse_unquote = Ωparse.ϒunquote
		ϒcompat_urllib_parse_unquote_plus = Ωparse.ϒunquote_plus
		ϒcompat_urllib_parse_urlencode = Ωparse.ϒurlencode
		ϒcompat_HTTPError = Ωerror.HTTPError
		ϒcompat_parse_qs = Ωparse.ϒparse_qs
		ϒcompat_html_entities_html5 = Ωentities.ϒhtml5
		ϒcompat_b64decode = Ωbase64.ϒb64decode
		ϒcompat_kwargs = λ.NewFunction("compat_kwargs",
			[]λ.Param{
				{Name: "kwargs"},
			},
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				var (
					ϒkwargs = λargs[0]
				)
				return ϒkwargs
			})
		ϒcompat_http_client = λ.Cal(λ.TypeType, λ.StrLiteral("compat_http_client"), λ.NewTuple(), func() λ.Dict {
			var (
				ϒcompat_http_client_HTTPException  λ.Object
				ϒcompat_http_client_IncompleteRead λ.Object
			)
			ϒcompat_http_client_IncompleteRead = λ.Cal(λ.TypeType, λ.StrLiteral("IncompleteRead"), λ.NewTuple(λ.BaseExceptionType), func() λ.Dict {
				// pass
				return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
			}())
			ϒcompat_http_client_HTTPException = λ.Cal(λ.TypeType, λ.StrLiteral("HTTPException"), λ.NewTuple(λ.BaseExceptionType), func() λ.Dict {
				// pass
				return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
			}())
			return λ.ClassDictLiteral(map[string]λ.Object{
				"HTTPException":  ϒcompat_http_client_HTTPException,
				"IncompleteRead": ϒcompat_http_client_IncompleteRead,
			})
		}())
		ϒcompat_xml_parse_error = ΩElementTree.ParseError
		ϒcompat_etree_fromstring = λ.NewFunction("compat_etree_fromstring",
			[]λ.Param{
				{Name: "text"},
			},
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				var (
					ϒtext = λargs[0]
				)
				return λ.Cal(ΩElementTree.XML, ϒtext)
			})
		ϒcompat_xpath = λ.NewFunction("compat_xpath",
			[]λ.Param{
				{Name: "xpath"},
			},
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				var (
					ϒxpath = λargs[0]
				)
				return ϒxpath
			})
	})
}
