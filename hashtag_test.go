package hashtag

import (
	"reflect"
	"testing"
)

func TestExtract(t *testing.T) {
	tests := map[string][]string{
		"#hashtag":                                                []string{"hashtag"},
		"\uff03hashtag":                                           []string{"hashtag"},
		"#Azərbaycanca":                                           []string{"Azərbaycanca"},
		"#mûǁae":                                                  []string{"mûǁae"},
		"#Čeština":                                                []string{"Čeština"},
		"#Ċaoiṁín":                                                []string{"Ċaoiṁín"},
		"#Caoiṁín":                                                []string{"Caoiṁín"},
		"#ta\u0301im":                                             []string{"ta\u0301im"},
		"#hag\u0303ua":                                            []string{"hag\u0303ua"},
		"#caf\u00E9":                                              []string{"caf\u00E9"},
		"#\u05e2\u05d1\u05e8\u05d9\u05ea":                         []string{"\u05e2\u05d1\u05e8\u05d9\u05ea"},                         // "#Hebrew"
		"#\u05d0\u05b2\u05e9\u05b6\u05c1\u05e8":                   []string{"\u05d0\u05b2\u05e9\u05b6\u05c1\u05e8"},                   // with marks
		"#\u05e2\u05b7\u05dc\u05be\u05d9\u05b0\u05d3\u05b5\u05d9": []string{"\u05e2\u05b7\u05dc\u05be\u05d9\u05b0\u05d3\u05b5\u05d9"}, // with maqaf 05be
		"#\u05d5\u05db\u05d5\u05f3":                               []string{"\u05d5\u05db\u05d5\u05f3"},                               // with geresh 05f3
		"#\u05de\u05f4\u05db":                                     []string{"\u05de\u05f4\u05db"},                                     // with gershayim 05f4
		"#\u0627\u0644\u0639\u0631\u0628\u064a\u0629":             []string{"\u0627\u0644\u0639\u0631\u0628\u064a\u0629"},             // "#Arabic"
		"#\u062d\u0627\u0644\u064a\u0627\u064b":                   []string{"\u062d\u0627\u0644\u064a\u0627\u064b"},                   // with mark
		"#\u064a\u0640\ufbb1\u0640\u064e\u0671":                   []string{"\u064a\u0640\ufbb1\u0640\u064e\u0671"},                   // with pres. form
		"#ประเทศไทย":                                              []string{"ประเทศไทย"},
		"#ฟรี":                                                    []string{"ฟรี"}, // with mark
		"#日本語ハッシュタグ":                                              []string{"日本語ハッシュタグ"},
		"＃日本語ハッシュタグ":                                              []string{"日本語ハッシュタグ"},
		"#1": []string{},
		"#0": []string{},
		"これはOK #ハッシュタグ": []string{"ハッシュタグ"},
		"これもOK。#ハッシュタグ": []string{"ハッシュタグ"},
		"これはダメ#ハッシュタグ":  []string{},
		"#hashtag mention": []string{
			"hashtag",
		},
		" #hashtag mention": []string{
			"hashtag",
		},
		"mention #hashtag here": []string{
			"hashtag",
		},
		"text #hashtag1 #hashtag2": []string{
			"hashtag1",
			"hashtag2",
		},
		" #user1 mention #user2 here #user3 ": []string{
			"user1",
			"user2",
			"user3",
		},
	}

	for k, v := range tests {
		tags := ExtractHashtags(k)
		if !reflect.DeepEqual(tags, v) {
			t.Errorf("Mismatch in %s: Want %v : Got %v", k, v, tags)
		}
	}

}
