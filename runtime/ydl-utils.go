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
 * ydl-utils.go: Re-implementation of various functions from
 *               https://github.com/rg3/youtube-dl/blob/master/youtube_dl/utils.py
 */

package runtime

import (
	"fmt"
	htmlLib "html"
	"net/http"
	urlLib "net/url"
	"strconv"
	"strings"
	"unicode"

	"github.com/tenta-browser/go-pcre-matcher/replacer"

	"github.com/tenta-browser/go-pcre-matcher"
	"github.com/tenta-browser/go-video-downloader/utils"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var cleanHTMLRe1, cleanHTMLRe2, cleanHTMLRe3 matcher.Regexp

// CleanHTML implements utils.py/clean_html
func CleanHTML(html OptString) OptString {
	// these expressions must be compiled lazily, since the regexp engine isn't set at init time
	if cleanHTMLRe1 == nil {
		cleanHTMLRe1 = ReMustCompile(`\s*<\s*br\s*/?\s*>\s*`, 0)
	}
	if cleanHTMLRe2 == nil {
		cleanHTMLRe2 = ReMustCompile(`<\s*/\s*p\s*>\s*<\s*p[^>]*>`, 0)
	}
	if cleanHTMLRe3 == nil {
		cleanHTMLRe3 = ReMustCompile(`<.*?>`, 0)
	}
	if !html.IsSet() {
		return html
	}
	_html := html.Get()
	_html = strings.Replace(_html, "\n", " ", -1)
	_html = cleanHTMLRe1.Replace(_html, "\n")
	_html = cleanHTMLRe2.Replace(_html, "\n")
	_html = cleanHTMLRe3.Replace(_html, "")
	_html = htmlLib.UnescapeString(_html)
	_html = strings.TrimSpace(_html)
	return AsOptString(_html)
}

// UnescapeHTML implements utils.py/unescapeHTML
func UnescapeHTML(html OptString) OptString {
	if !html.IsSet() {
		return html
	}
	return AsOptString(htmlLib.UnescapeString(html.Get()))
}

var knownExtensions = utils.MakeSet([]string{
	"mp4", "m4a", "m4p", "m4b", "m4r", "m4v", "aac",
	"flv", "f4v", "f4a", "f4b",
	"webm", "ogg", "ogv", "oga", "ogx", "spx", "opus",
	"mkv", "mka", "mk3d",
	"avi", "divx",
	"mov",
	"asf", "wmv", "wma",
	"3gp", "3g2",
	"mp3",
	"flac",
	"ape",
	"wav",
	"f4f", "f4m", "m3u8", "smil",
})

// DetermineExt implements utils.py/determine_ext
func DetermineExt(url OptString, defaultExt string) string {
	if !url.IsSet() {
		return defaultExt
	}

	// get it from before '?' and after the last '.'
	guess := url.Get()
	pos := utils.IndexRune(guess, '?')
	if pos < 0 {
		pos = len(guess)
	}
	guess = guess[:pos]
	guess = guess[utils.LastIndexRune(guess, '.')+1:]

	if match := ReMatch("^[A-Za-z0-9]+$", guess, 0); match != nil {
		return guess
	}

	guess = strings.TrimRight(guess, "/")
	if _, ok := knownExtensions[guess]; ok {
		return guess
	}

	return defaultExt
}

// taken from https://stackoverflow.com/a/26722698
func removeDiacritics(s string) string {
	isMn := func(r rune) bool {
		return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
	}
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	rs, _, _ := transform.String(t, s)
	return rs
}

// SanitizeFilename implements utils.py/sanitize_filename
func SanitizeFilename(s string, restricted bool, isID bool) string {
	// TODO handle timestamps (needs replace with func callback)

	if restricted {
		s = removeDiacritics(s)
	}

	replaceInsaneChar := func(char rune) string {
		if char == '?' || int(char) < 32 || int(char) == 127 {
			return ""
		} else if char == '"' {
			if restricted {
				return ""
			}
			return "'"
		} else if char == ':' {
			if restricted {
				return "_-"
			}
			return " -"
		} else if utils.IndexRune("\\/|*<>", char) >= 0 {
			return "_"
		} else if restricted && (utils.IndexRune("!&'()[]{}$;`^,#", char) >= 0 || unicode.IsSpace(char)) {
			return "_"
		} else if restricted && int(char) > 127 {
			return "_"
		} else {
			return string(char)
		}
	}

	var sparts []string
	for _, char := range s {
		sparts = append(sparts, replaceInsaneChar(char))
	}
	s = strings.Join(sparts, "")

	if !isID {
		for strings.Index(s, "__") >= 0 {
			s = strings.Replace(s, "__", "_", -1)
		}
		s = strings.Trim(s, "_")
		if restricted && strings.HasPrefix(s, "-_") {
			s = s[2:]
		}
		if strings.HasPrefix(s, "-") {
			s = "_" + s[1:]
		}
		s = strings.TrimLeft(s, ".")
		if s == "" {
			s = "_"
		}
	}

	return s
}

// StripOrNone implements utils.py/strip_or_none
func StripOrNone(s OptString) OptString {
	if !s.IsSet() {
		return s
	}
	return AsOptString(StrStrip(s.Get(), ""))
}

// PyExtractorError implements utils.py/ExtractorError
func PyExtractorError(msg string, expected bool, videoID OptString) error {
	if videoID.IsSet() {
		msg = videoID.Get() + ": " + msg
	}
	return newExtractorError(msg)
}

// IntOrNone implements utils.py/int_or_none
func IntOrNone(val interface{}, scale int, def OptInt, invScale int) OptInt {
	if val == nil || val == "" {
		return def
	} else if oint, ok := val.(OptInt); ok {
		if !oint.IsSet() {
			return def
		}
		val = oint.Get()
	} else if ostr, ok := val.(OptString); ok {
		if !ostr.IsSet() {
			return def
		}
		val = ostr.Get()
	}
	switch v := val.(type) {
	case int:
		return AsOptInt(v * invScale / scale)
	case float32:
		return AsOptInt(int(v) * invScale / scale)
	case float64:
		return AsOptInt(int(v) * invScale / scale)
	case string:
		if iv, err := strconv.Atoi(v); err == nil {
			return AsOptInt(iv * invScale / scale)
		}
		return def
	default:
		panic(newExtractorError(fmt.Sprintf("Cannot convert to int: %T", v)))
	}
}

// StrToInt implements utils.py/str_to_int
func StrToInt(s OptString) OptInt {
	if !s.IsSet() {
		return OptInt{}
	}
	ss := ReSub(`[,\.\+]`, "", s.Get(), 0, 0)
	i, err := strconv.Atoi(ss)
	if err != nil {
		panic(newExtractorError(fmt.Sprintf("Failed to convert to int: '%s'", s.Get())))
	}
	return AsOptInt(i)
}

// UtilDictGet implements utils.py/dict_get
func UtilDictGet(dict map[string]interface{}, key interface{}, def interface{}, skipFalseValues bool) interface{} {
	keys, ok := key.([]string)
	if !ok {
		keys = []string{key.(string)}
	}
	for _, key := range keys {
		val, ok := dict[key]
		if !ok || val == nil {
			continue
		}
		if skipFalseValues {
			// TODO possibly handle empty slices and maps too
			if val == 0 || val == "" || val == nil {
				continue
			}
		}
		return val
	}
	return def
}

// UnifiedStrDate implements utils.py/unified_strdate
func UnifiedStrDate(dateStr OptString, dayFirst bool) OptString {
	// TODO implement me
	return OptString{}
}

// ParseDuration implements utils.py/parse_duration
func ParseDuration(s OptString) OptString {
	// TODO implement me
	return OptString{}
}

// DetermineProtocol implements utils.py/determine_protocol
func DetermineProtocol(infoDict map[string]interface{}) string {
	if protocol := GetStringField(infoDict, "protocol", false, ""); protocol != "" {
		return protocol
	}

	url := GetStringField(infoDict, "url", true, "")

	if strings.HasPrefix(url, "rtmp") {
		return "rtmp"
	} else if strings.HasPrefix(url, "mms") {
		return "mms"
	} else if strings.HasPrefix(url, "rtsp") {
		return "rtsp"
	}

	ext := DetermineExt(AsOptString(url), "unknown_video")

	if ext == "m3u8" {
		return "m3u8"
	} else if ext == "f4m" {
		return "f4m"
	}

	purl, err := urlLib.Parse(url)
	if err != nil {
		panic(newExtractorError("Invalid URL: " + url))
	}

	return purl.Scheme
}

// RemoveQuotes implements utils.py/remove_quotes
func RemoveQuotes(s OptString) OptString {
	if !s.IsSet() {
		return s
	}
	ss := s.Get()
	if len(ss) < 2 {
		return s
	}
	for _, quote := range []byte{'\'', '"'} {
		if ss[0] == quote && ss[len(ss)-1] == quote {
			return AsOptString(ss[1 : len(ss)-1])
		}
	}
	return s
}

// JsToJSON implements utils.py/js_to_json
func JsToJSON(code string) string {
	commentRe := `/\*(?:(?!\*/).)*?\*/|//[^\n]*`
	skipRe := fmt.Sprintf(`\s*(?:%s)?\s*`, commentRe)
	integerTable := []struct {
		regex string
		base  int
	}{
		{fmt.Sprintf(`(?s)^(0[xX][0-9a-fA-F]+)%s:?$`, skipRe), 16},
		{fmt.Sprintf(`(?s)^(0+[0-7]+)%s:?$`, skipRe), 8},
	}

	fixKv := func(m matcher.Match) string {
		v := m.GroupByIdx(0)
		if v == "true" || v == "false" || v == "null" {
			return v
		} else if strings.HasPrefix(v, "/*") || strings.HasPrefix(v, "//") || v == "," {
			return ""
		}

		if v[0] == '\'' || v[0] == '"' {
			v = ReMustCompile(`(?s)\\.|"`, 0).ReplaceFunc(v[1:len(v)-1], replacer.NewReplacer(func(m matcher.Match) string {
				m0 := m.GroupByIdx(0)
				switch m0 {
				case "\"":
					return "\\\""
				case "\\'":
					return "'"
				case "\\\n":
					return ""
				case "\\x":
					return "\\u0"
				default:
					return m0
				}
			}))
		}

		for _, rb := range integerTable {
			im := ReMatch(rb.regex, v, 0)
			if im != nil {
				i64, err := strconv.ParseInt(im.GroupByIdx(1), rb.base, 0)
				if err != nil {
					panic(newExtractorError("Int conversion failed: " + err.Error()))
				}
				if strings.HasSuffix(v, ":") {
					return fmt.Sprintf(`"%d"`, i64)
				}
				return string(i64)
			}
		}

		return fmt.Sprintf(`"%s"`, v)
	}

	return ReMustCompile(fmt.Sprintf(`(?sx)
		"(?:[^"\\]*(?:\\\\|\\['"nurtbfx/\n]))*[^"\\]*"|
		'(?:[^'\\]*(?:\\\\|\\['"nurtbfx/\n]))*[^'\\]*'|
		%[1]s|,(?=%[2]s[\]}}])|
		[a-zA-Z_][.a-zA-Z_0-9]*|
		\b(?:0[xX][0-9a-fA-F]+|0+[0-7]+)(?:%[2]s:)?|
		[0-9]+(?=%[2]s:)`,
		commentRe, skipRe), 0).ReplaceFunc(code, replacer.NewReplacer(fixKv))
}

// sanitizeURL implements utils.py/sanitize_url
func sanitizeURL(url string) string {
	if strings.HasPrefix(url, "//") {
		return "http:" + url
	}
	return url
}

// SanitizedRequest implements utils.py/sanitized_Request
func SanitizedRequest(url string) *http.Request {
	req, err := http.NewRequest("GET", sanitizeURL(url), nil)
	if err != nil {
		panic(newExtractorError(err.Error()))
	}
	return req
}

// RequestAddHeader implements python/Request.add_header
func RequestAddHeader(req *http.Request, key, val string) {
	req.Header.Set(key, val)
}
