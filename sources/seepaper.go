package sources

import (
	"errors"
	"fmt"
	"github.com/BaizeSec/Pusher/pkg/logger"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func GetSeePaper() string {
	s, _ := getAnquanke()
	if len(s) == 0 {
		logger.Info("正在获取 Seebug Paper\t[0]")
		logger.Warning("Seebug 今天又拉了..")
		return ""
	}
	logger.Info("正在获取 Seebug Paper\t[" + strconv.Itoa(len(s)) + "]")
	body := "Seebug Paper:\n"
	for _, i := range s {
		text := "标题：" + i[0] + "\\n地址：" + i[1] + "\\n-----------------------\\n"
		body = body + text
	}
	return body
}

func getSeebugPaper() ([][]string, error) {
	client := &http.Client{
		Timeout: time.Duration(4) * time.Second,
	}
	req, err := http.NewRequest("GET", "https://paper.seebug.org/rss/", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
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

	re := regexp.MustCompile(`<item><title>([\w\W]*?)</title><link>([\w\W]*?)</link><description>[\w\W]*?</description><pubDate>([\w\W]*?)</pubDate><guid>[\w\W]*?</guid><category>[\w\W]*?/category></item>`)
	result := re.FindAllStringSubmatch(strings.TrimSpace(bodyString), -1)

	var resultSlice [][]string
	for _, match := range result {
		utc, _ := time.LoadLocation("UTC")
		t, err := time.ParseInLocation(time.RFC1123Z, match[1:][2], utc)
		if err != nil {
			return nil, err
		}

		time_zone := time.FixedZone("CST", 8*3600)
		if !isIn24Hours(t.In(time_zone)) {
			// 默认时间顺序是从近到远
			break
		}

		// 去除title中的换行符
		re, _ = regexp.Compile(`\s{1,}`)
		match[1:][0] = re.ReplaceAllString(match[1:][0], "")

		resultSlice = append(resultSlice, match[1:][0:2])
		// slice中title和url调换位置，以符合统一的format
		for _, item := range resultSlice {
			item[0], item[1] = item[1], item[0]
		}
	}
	if len(resultSlice) == 0 {
		return nil, errors.New("no records in the last 24 hours")
	}
	return resultSlice, nil
}

func currentTime() string {
	time_zone := time.FixedZone("CST", 8*3600) // 8*3600 = 8h
	n := time.Now().In(time_zone)
	// 获取时间，格式如2006/01/02 15:04:05
	t := n.Format("2006/01/02 15:04:05")
	weekMap := map[time.Weekday]string{0: "星期日", 1: "星期一", 2: "星期二", 3: "星期三", 4: "星期四", 5: "星期五", 6: "星期六"}
	formatTime := fmt.Sprintf("%s %s", t, weekMap[n.Weekday()])
	return formatTime
}

func isIn24Hours(t time.Time) bool {
	time_zone := time.FixedZone("CST", 8*3600) // 8*3600 = 8h
	now := time.Now().In(time_zone)
	// 根据config生成每日整点时间
	cronTime := time.Date(now.Year(), now.Month(), now.Day(), int(10), 0, 0, 0, time_zone)
	subTime := cronTime.Sub(t)
	if subTime > time.Duration(24)*time.Hour || subTime < time.Duration(0) {
		return false
	}
	return true
}
