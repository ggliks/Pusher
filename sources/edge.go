package sources

import (
	"github.com/BaizeSec/Pusher/pkg/logger"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func GetEdge() string {
	s := getedgedaily()
	if len(s) == 0 {
		logger.Info("正在获取棱角日报\t[0]")
		logger.Warning("棱角今天又拉了..")
		return ""
	}
	logger.Info("正在获取棱角日报\t[" + strconv.Itoa(len(s)) + "]")
	body := "棱角日报:\n"
	for _, i := range s {
		text := "标题：" + i[1] + "\\n地址：" + i[0] + "\\n标签：" + i[2] + "\\n-----------------------\\n"
		body = body + text
	}
	return body
}

func getedgedaily() [][]string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://forum.ywhack.com/forumdisplay.php?fid=59&orderby=lastpost&filter=86400", nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.162 Safari/537.36")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	src := string(bodyText)

	//去除STYLE
	re, _ := regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "")

	re = regexp.MustCompile(`<a href="(.*?)" target="_blank">(.*?)</a>.*?<img src="" style="vertical-align: top;margin-top: 2px;"></div><small class="card-subtitle text-muted">.*?<span class="badge badge-tag">(.*?)</span></small>`)
	res := re.FindAllStringSubmatch(strings.TrimSpace(src), -1)
	var result [][]string
	for _, s := range res {
		result = append(result, s[1:])
	}
	if result != nil {
		result = append(result)
	}
	return result
}
