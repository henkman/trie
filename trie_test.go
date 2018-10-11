package trie_test

import (
	"fmt"
	"testing"

	"github.com/henkman/trie"
)

type Test struct {
	Domain      string
	ShouldMatch bool
}

func TestSubDomain(t *testing.T) {
	var trie trie.Trie
	for _, domain := range []string{
		"google.com",
		"youtube.com",
		"facebook.com",
		"baidu.com",
		"wikipedia.org",
		"reddit.com",
		"yahoo.com",
		"qq.com",
		"taobao.com",
		"google.co.in",
		"amazon.com",
		"tmall.com",
		"twitter.com",
		"sohu.com",
		"instagram.com",
		"vk.com",
		"live.com",
		"jd.com",
		"sina.com.cn",
		"weibo.com",
		"yandex.ru",
		"360.cn",
		"google.co.jp",
		"google.co.uk",
		"google.ru",
		"google.com.br",
		"netflix.com",
		"google.de",
		"google.com.hk",
		"twitch.tv",
		"google.fr",
		"linkedin.com",
		"yahoo.co.jp",
		"t.co",
		"csdn.net",
		"microsoft.com",
		"bing.com",
		"office.com",
		"ebay.com",
		"alipay.com",
		"google.it",
		"google.ca",
		"mail.ru",
		"ok.ru",
		"google.es",
		"msn.com",
		"google.com.tr",
		"google.com.au",
		"whatsapp.com",
		"spotify.com",
		"google.pl",
		"google.co.id",
		"google.com.ar",
		"google.co.th",
		"naver.com",
		"sogou.com",
		"samsung.com",
		"accuweather.com",
		"goo.gl",
		"sm.cn",
		"googleweblight.com",
	} {
		trie.Insert(Reverse(domain))
	}

	for _, test := range []Test{
		{Domain: "maps.google.com", ShouldMatch: true},
		{Domain: "en.wikipedia.org", ShouldMatch: true},
		{Domain: "accuweather.com", ShouldMatch: true},
		{Domain: "docs.google.com", ShouldMatch: true},
		{Domain: "dev.twitch.tv", ShouldMatch: true},
		{Domain: "google.com", ShouldMatch: true},
		{Domain: "google.co", ShouldMatch: false},
		{Domain: "noogle.com", ShouldMatch: false},
		{Domain: "123", ShouldMatch: false},
		{Domain: "", ShouldMatch: false},
	} {
		n := trie.Lookup(Reverse(test.Domain))
		if n != nil && (len(n.Children) == 0 || n.Leaf) {
			if !test.ShouldMatch {
				t.Log(test.Domain, "should not match",
					fmt.Sprintf("(%c %#v)", n.Character, n.Container))
				t.Fail()
				break
			}
		} else {
			if test.ShouldMatch {
				t.Log(test.Domain, "should match",
					fmt.Sprintf("(%c %#v)", n.Character, n.Container))
				t.Fail()
				break
			}
		}
	}

}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
