package main

import (
	"Demo/config"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"github.com/levigross/grequests"
	"github.com/saintfish/chardet"
)

func init() {
	// file := "./" + "message" + ".log"
	// logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	// if err != nil {
	// 	panic(err)
	// }
	// log.SetOutput(logFile) // 将文件设置为log输出的文件
	// log.SetPrefix("[qSkipTool]")
	// log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	// return
}

func main() {

	for {
		name := ""
		fmt.Println("输入要搜索小说名称：")
		fmt.Scanf("%s\n", &name)
		if name == "" {
			break
		}
		SearchNovelsPost(name)
	}
}

//SearchNovelsPost 搜索小说
func SearchNovelsPost(novelName string) error {
	//novelName := ""
	if novelName == "" {
		return errors.New("name is nil")
	}
	response, err := requestURLPost("http://www.paoshuzw.com/modules/article/waps.php", novelName)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		s := string(rune(response.StatusCode))
		return errors.New("response.StatusCode:" + s)
	}

	rawHTML := detectBody(response.Bytes())
	if rawHTML == "" {
		return errors.New("Html is nil")
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(rawHTML))
	if err != nil {
		return err
	}

	var resultData []map[string]string
	doc.Find("#checkform > table > tbody > tr").Each(func(i int, s *goquery.Selection) {
		a := make(map[string]string)
		a["name"] = s.Find("td:nth-child(1) > a").Text()
		if a["name"] == "" {
			return
		}
		a["novel_latest_chapter_name"] = s.Find("td:nth-child(2) > a").Text()
		a["novel_author"] = s.Find("td:nth-child(3)").Text()
		a["url"], _ = s.Find("td:nth-child(1) > a").Attr("href")
		fmt.Println(a["url"])
		resultData = append(resultData, a)
	})

	for i, u := range resultData {
		fmt.Printf("%d. %s, %s, %s\n", i, u["name"], u["novel_latest_chapter_name"], u["novel_author"])
	}
	return nil
}

func requestURLPost(url string, value string) (*grequests.Response, error) {
	ro := &grequests.RequestOptions{
		Data: map[string]string{
			"searchkey": value,
		},
		Headers: map[string]string{"User-Agent": config.GetUserAgent()},
	}
	resp, err := grequests.Post(url, ro)
	if err != nil {
		log.Println("Unable to make request: ", err)
		return nil, err
	}
	return resp, err
}

func detectBody(body []byte) string {
	var bodyString string
	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest(body)
	if err != nil {
		return string(body)
	}
	if strings.Contains(strings.ToLower(result.Charset), "utf") {
		bodyString = string(body)
	} else {
		bodyString = mahonia.NewDecoder("gbk").ConvertString(string(body))
	}
	return bodyString
}
