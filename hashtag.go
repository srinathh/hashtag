/*
   Copyright 2014 Hariharan Srinath

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

/*
Package hashtag implements the logic of Twitter's official Java hashtag
extractor in Go correctly handling both # and \uFF03 as hashtag markers
and correctly omitting hashtags follwed immediately by hash or url markers.
As an additional feature, this package also extracts hashtags correctly
from multiline strings. Twitter's original java implementation can be
found at their repo https://github.com/twitter/twitter-text
*/
package hashtag

import (
	"regexp"
	"strings"
)

const (
	hashtag_letters       = "\\pL\\pM"
	hashtag_numerals      = "\\p{Nd}"
	hashtag_special_chars = "_" + //underscore
		"\\x{200c}" + // ZERO WIDTH NON-JOINER (ZWNJ)
		"\\x{200d}" + // ZERO WIDTH JOINER (ZWJ)
		"\\x{a67e}" + // CYRILLIC KAVYKA
		"\\x{05be}" + // HEBREW PUNCTUATION MAQAF
		"\\x{05f3}" + // HEBREW PUNCTUATION GERESH
		"\\x{05f4}" + // HEBREW PUNCTUATION GERSHAYIM
		"\\x{309b}" + // KATAKANA-HIRAGANA VOICED SOUND MARK
		"\\x{309c}" + // KATAKANA-HIRAGANA SEMI-VOICED SOUND MARK
		"\\x{30a0}" + // KATAKANA-HIRAGANA DOUBLE HYPHEN
		"\\x{30fb}" + // KATAKANA MIDDLE DOT
		"\\x{3003}" + // DITTO MARK
		"\\x{0f0b}" + // TIBETAN MARK INTERSYLLABIC TSHEG
		"\\x{0f0c}" + // TIBETAN MARK DELIMITER TSHEG BSTAR
		"\\x{0f0d}" // TIBETAN MARK SHAD

	hashtag_letters_numerals     = hashtag_letters + hashtag_numerals + hashtag_special_chars
	hashtag_letters_numerals_set = "[" + hashtag_letters_numerals + "]"
	hashtag_letters_set          = "[" + hashtag_letters + "]"
)

var hashtag_start = regexp.MustCompile("(?m)(?:^|[^&" + hashtag_letters_numerals + "])(?:#|\\x{FF03})(" +
	hashtag_letters_numerals_set + "*" + hashtag_letters_set + hashtag_letters_numerals_set +
	"*)")

/*
ExtractHashtags extracts hashtag texts without hash markers contained in the string provided
the input parameter and returns them as a slice of strings. For example passing the following string
    " #user1 mention #user2 here #user3 "
will cause the following string slice to be returned by the function. Look at hashtag_test.go for more examples
    []string{"user1","user2","user3"}

*/
func ExtractHashtags(s string) []string {

	//Performance optimization. If text doesn't contain # or \xFF03 , return
	if !strings.ContainsAny(s, "#\uFF03") {
		return []string{}
	}

	tags := hashtag_start.FindAllStringSubmatchIndex(s, -1)
	ret := []string{}

	//we need to loop through each tag and check if the chars following tag are not # or ://
	for _, match := range tags {
		tagtext := s[match[2]:match[3]] //the first group is the full string
		suffix := s[match[3]:]          //the last char of the third group is the start of suffix

		if strings.HasPrefix(suffix, "#") || strings.HasPrefix(suffix, "\\xFF03") || strings.HasPrefix(suffix, "://") {
			continue
		}

		ret = append(ret, tagtext)
	}

	return ret
}
