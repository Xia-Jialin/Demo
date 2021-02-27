package config

import (
	"math/rand"
	"time"
)

// ItemRuleConfig contains information about novel's rules
type ItemRuleConfig struct {
	NovelName              string
	NovelURL               string
	NovelType              string
	NovelAuthor            string
	NovelCover             string
	NovelAbstract          string
	NovelLatestChapterName string
	NovelLatestChapterURL  string
}

// NovelRule contains information about novle's source
type NovelRule struct {
	Name            string
	HomeURL         string
	SearchURL       string
	Method          string
	Params          map[string]string
	KeywordEncoding string
	TargetItem      string
	ItemRule        ItemRuleConfig
}

var (
	userAgents = []string{
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.101 Safari/537.36",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; Media Center PC 6.0; InfoPath.3; MS-RTC LM 8; Zune 4.7)",
		"Mozilla/4.0 (compatible; MSIE 9.0; Windows NT 5.1; Trident/5.0)",
		"Mozilla/5.0 (X11; Linux x86_64; rv:2.2a1pre) Gecko/20100101 Firefox/4.2a1pre",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:2.0b11pre) Gecko/20110131 Firefox/4.0b11pre",
		"Mozilla/5.0 (X11; U; Linux i686; ru-RU; rv:1.9.2a1pre) Gecko/20090405 Ubuntu/9.04 (jaunty) Firefox/3.6a1pre",
		"Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.2.8) Gecko/20100723 SUSE/3.6.8-0.1.1 Firefox/3.6.8",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; pt-PT; rv:1.9.2.6) Gecko/20100625 Firefox/3.6.6",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; it; rv:1.9.2.6) Gecko/20100625 Firefox/3.6.6 ( .NET CLR 3.5.30729)",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US; rv:1.9.2.6) Gecko/20100625 Firefox/3.6.6 (.NET CLR 3.5.30729)",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; ru; rv:1.9.2.4) Gecko/20100513 Firefox/3.6.4",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; ja; rv:1.9.2.4) Gecko/20100611 Firefox/3.6.4 GTB7.1",
	}

	//ItemRule01 关于笔趣阁的规则
	ItemRule01 = ItemRuleConfig{
		NovelName:              "div.bookinfo h4.bookname a",
		NovelURL:               "div.bookinfo h4.bookname a",
		NovelType:              "div.bookinfo div.cat",
		NovelAuthor:            "div.bookinfo div.author",
		NovelCover:             "div.bookimg a img",
		NovelAbstract:          "div.bookinfo p",
		NovelLatestChapterName: "div.bookinfo div.update a",
		NovelLatestChapterURL:  "div.bookinfo div.update a",
	}
	//ItemRule02 百度第三方网站内容检索的规则
	ItemRule02 = ItemRuleConfig{
		NovelName:              "div.result-game-item-detail a.result-game-item-title-link",
		NovelURL:               "div.result-game-item-detail a.result-game-item-title-link",
		NovelType:              "div.result-game-item-info p:nth-child(2)",
		NovelAuthor:            "div.result-game-item-info p:nth-child(1)",
		NovelCover:             "img.result-game-item-pic-link-img",
		NovelAbstract:          "p.result-game-item-desc",
		NovelLatestChapterName: "p.result-game-item-info-tag a.result-game-item-info-tag-item",
		NovelLatestChapterURL:  "p.result-game-item-info-tag a.result-game-item-info-tag-item",
	}
	//ItemRule03 全本小说网的小说检索规则
	ItemRule03 = ItemRuleConfig{
		NovelName:              "h1.f20h",
		NovelURL:               "div.option span.btopt a",
		NovelType:              "div.box_intro > div.box_info > table > tbody > tr:nth-child(1) > td > div:nth-child(1) > p:nth-child(3) > a",
		NovelAuthor:            "div.box_info > table > tbody > tr:nth-child(1) > td > div:nth-child(1) > p:nth-child(2)",
		NovelCover:             "div.box_intro div.pic img",
		NovelAbstract:          "div.intro",
		NovelLatestChapterName: "div.box_info > table > tbody > tr:nth-child(1) > td > div:nth-child(1) > p:nth-child(2) a",
		NovelLatestChapterURL:  "div.box_info > table > tbody > tr:nth-child(1) > td > div:nth-child(1) > p:nth-child(2) a",
	}
	//ItemRule04 全本小说网的作者信息检索规则 这里自定义为key_1 其中NovelUrl 值获取的是下载地址 客户端可根据最新章节地址提取出path就是目录地址
	ItemRule04 = ItemRuleConfig{
		NovelName:              "ul.info>h2>b>a>font",
		NovelURL:               "",
		NovelType:              "ul.info > li:nth-child(4) > font:nth-child(2)",
		NovelAuthor:            "ul.info > li:nth-child(4) > font:nth-child(1) > b",
		NovelCover:             "ul.info a img",
		NovelAbstract:          "ul.info > li:nth-child(3)",
		NovelLatestChapterName: "ul.info > h2 > a > font > b",
		NovelLatestChapterURL:  "ul.info > h2 > a",
	}
	//ItemRule05 新笔趣阁的解析规则
	ItemRule05 = ItemRuleConfig{
		NovelName:              "h3.result-item-title a span",
		NovelURL:               "h3.result-item-title a",
		NovelType:              "div.result-game-item-info > p:nth-child(2) > span:nth-child(2)",
		NovelAuthor:            "div.result-game-item-info > p:nth-child(1) > span:nth-child(2)",
		NovelCover:             "img.result-game-item-pic-link-img",
		NovelAbstract:          "p.result-game-item-desc",
		NovelLatestChapterName: "div.result-game-item-info p.result-game-item-info-tag a.result-game-item-info-tag-item",
		NovelLatestChapterURL:  "div.result-game-item-info p.result-game-item-info-tag a.result-game-item-info-tag-item",
	}

	//NovelsRulesMap 网址集合
	NovelsRulesMap = map[string]NovelRule{
		"10": {
			Name:            "笔趣阁01",
			HomeURL:         "https://www.bqg99.cc/",
			SearchURL:       "https://www.bqg99.cc/s.php?q=",
			Method:          "Get",
			Params:          make(map[string](string)),
			KeywordEncoding: "",
			TargetItem:      ".bookbox",
			ItemRule:        ItemRule01,
		},
		"11": {
			Name:            "笔趣阁02",
			HomeURL:         "http://www.cdzdgw.com/",
			SearchURL:       "http://www.cdzdgw.com/s.php?q=",
			Method:          "Get",
			Params:          make(map[string](string)),
			KeywordEncoding: "",
			TargetItem:      ".bookbox",
			ItemRule:        ItemRule01,
		},
		"12": {
			Name:            "笔趣阁03",
			HomeURL:         "http://www.biqugex.com/",
			SearchURL:       "http://www.biqugex.com/s.php?q=",
			Method:          "Get",
			Params:          make(map[string](string)),
			KeywordEncoding: "",
			TargetItem:      ".bookbox",
			ItemRule:        ItemRule01,
		},
		"20": {
			Name:            "全本小说网01",
			HomeURL:         "https://www.ybdu.com/",
			SearchURL:       "https://www.ybdu.com/modules/article/search.php?searchtype=keywords&entry=1&searchkey=",
			Method:          "Get",
			Params:          make(map[string](string)),
			KeywordEncoding: "gbk",
			TargetItem:      "#detail-box",
			ItemRule:        ItemRule03,
		},
		"20_1": {
			Name:            "全本小说网01",
			HomeURL:         "https://www.ybdu.com/",
			SearchURL:       "https://www.ybdu.com/modules/article/search.php?searchtype=keywords&entry=1&searchkey=",
			Method:          "Get",
			Params:          make(map[string](string)),
			KeywordEncoding: "gbk",
			TargetItem:      "ul.info",
			ItemRule:        ItemRule04,
		},
		"30": {
			Name:            "新笔趣阁01",
			HomeURL:         "https://www.xxbiquge.com/",
			SearchURL:       "https://www.xxbiquge.com/search.php?keyword=",
			Method:          "Get",
			Params:          make(map[string](string)),
			KeywordEncoding: "",
			TargetItem:      "div.result-game-item",
			ItemRule:        ItemRule05,
		},
		"100": {
			Name:            "新笔趣阁_百度01",
			HomeURL:         "http://www.biqugetv.com/",
			SearchURL:       "http://zhannei.baidu.com/cse/search?click=1&s=16765504158186272814&q=",
			Method:          "Get",
			Params:          make(map[string](string)),
			KeywordEncoding: "",
			TargetItem:      "div.result-list div.result-item",
			ItemRule:        ItemRule02,
		},
		"110": {
			Name:            "笔下文学_百度01",
			HomeURL:         "http://www.xbxwx.net/",
			SearchURL:       "http://so.xbxwx.net/cse/search?click=1&entry=1&s=10874778206555383279&q=",
			Method:          "Get",
			Params:          make(map[string](string)),
			KeywordEncoding: "",
			TargetItem:      "div.result-list div.result-item",
			ItemRule:        ItemRule02,
		},
		"120": {
			Name:            "顶点小说_百度01",
			HomeURL:         "http://www.23wx.cc/du/99/99646/",
			SearchURL:       "http://zhannei.baidu.com/cse/search?s=17788970894453164958&q=",
			Method:          "Get",
			Params:          make(map[string](string)),
			KeywordEncoding: "",
			TargetItem:      "div.result-list div.result-item",
			ItemRule:        ItemRule02,
		},
	}
)

// GetUserAgent returns a random user agent
func GetUserAgent() string {
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(userAgents)
	return userAgents[n]
}
