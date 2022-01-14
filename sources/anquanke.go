package sources

import (
	"errors"
	"github.com/BaizeSec/Pusher/pkg/logger"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func GetAnquanke() string {
	s, _ := getSeebugPaper()
	if len(s) == 0 {
		logger.Info("正在获取 安全客\t[0]")
		logger.Warning("安全客 今天又拉了..")
		return ""
	}
	logger.Info("正在获取 安全客\t[" + strconv.Itoa(len(s)) + "]")
	body := "安全客:\n"
	for _, i := range s {
		text := "标题：" + i[0] + "\\n地址：" + i[1] + "\\n-----------------------\\n"
		body = body + text
	}
	return body
}

func getAnquanke() ([][]string, error) {
	client := &http.Client{
		Timeout: time.Duration(4) * time.Second,
	}
	req, err := http.NewRequest("GET", "https://www.anquanke.com/knowledge", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	bodyString := string(body)

	//去除STYLE
	re, _ := regexp.Compile(`\<style[\S\s]+?\</style\>`)
	bodyString = re.ReplaceAllString(bodyString, "")

	//去除SCRIPT
	re, _ = regexp.Compile(`\<script[\S\s]+?\</script\>`)
	bodyString = re.ReplaceAllString(bodyString, "")

	//去除head
	re, _ = regexp.Compile(`\<head[\S\s]+?\</head\>`)
	bodyString = re.ReplaceAllString(bodyString, "")

	//去除header
	re, _ = regexp.Compile(`\<header[\S\s]+?\</header\>`)
	bodyString = re.ReplaceAllString(bodyString, "")

	//去除footer
	re, _ = regexp.Compile(`\<footer[\S\s]+?\</footer\>`)
	bodyString = re.ReplaceAllString(bodyString, "")

	//去除sidebar
	re, _ = regexp.Compile(`\<div class="load-more[\S\s]+?\</html\>`)
	bodyString = re.ReplaceAllString(bodyString, "")

	//去除tag及desp
	re, _ = regexp.Compile(`\<div class="tags  hide-in-mobile-device[\S\s]+?class="fa fa-clock-o"\>`)
	bodyString = re.ReplaceAllString(bodyString, "")

	//去除多余的信息
	re, _ = regexp.Compile(`\<div class="article-item common-item">[\S\s]+?common-item-right"\>`)
	bodyString = re.ReplaceAllString(bodyString, "")

	//去除连续的换行符
	re, _ = regexp.Compile(`\s{2,}`)
	bodyString = re.ReplaceAllString(bodyString, "")

	//去除href中可能存在的class
	re, _ = regexp.Compile(`class="red-title"`)
	bodyString = re.ReplaceAllString(bodyString, "")

	re = regexp.MustCompile(`<div class="title"><a target="_blank" rel="noopener noreferrer"href="(.*?)"> (.*?)</a></div></i>(.*?)</span>`)
	result := re.FindAllStringSubmatch(strings.TrimSpace(bodyString), -1)

	var resultSlice [][]string
	for _, match := range result {
		match[1:][0] = "https://www.anquanke.com" + match[1:][0]
		time_zone := time.FixedZone("CST", 8*3600)
		t, err := time.ParseInLocation("2006-01-02 15:04:05", match[1:][2], time_zone)
		if err != nil {
			return nil, err
		}

		if !isIn24Hours(t) {
			// 默认时间顺序是从近到远
			break
		}
		resultSlice = append(resultSlice, match[1:][0:2])
	}
	if len(resultSlice) == 0 {
		return nil, errors.New("no records in the last 24 hours")
	}
	return resultSlice, nil
}
