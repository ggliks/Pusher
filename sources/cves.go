package sources

import (
	"github.com/BaizeSec/Pusher/pkg/logger"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

func GetCves() string {
	s := getcvesdaily()

	if len(s) == 0 {
		logger.Warning("麻了, 找不到CVE..")
		return ""
	}
	body := "# Today CVEs\n\n"
	body += "This document is produced by Little Cookie\n\n"
	for _, i := range s {
		text := "## " + i[0] + "\n\n地址: [" + i[1] + "](" + i[1] + ")\n\n描述: " + i[2] + "\n\n"
		body = body + text
	}
	return body
}

func getcvesdaily() [][]string {
	cveUrl := "https://cassandra.cerias.purdue.edu/CVE_changes/today.html"
	resp, err := http.Get(cveUrl)

	if err != nil {
		return nil
	}

	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	//fmt.Println(string(respBody))
	re, _ := regexp.Compile(`<A HREF = '(.*?)'>(.*?)</A><br />\n`)
	cves := re.FindAllStringSubmatch(string(respBody), -1)
	re, _ = regexp.Compile(`<HTML><BODY><BR>date: (.*?)<BR>New entries:<br />\n`)

	var resultSlice [][]string

	logger.Info("正在获取 CVE\t[" + strconv.Itoa(len(cves)) + "]")

	for i, cve := range cves {
		cveName := "CVE-" + cve[2]
		cvea := cve[1]
		logger.Info("[" + strconv.Itoa(i+1) + "/" + strconv.Itoa(len(cves)) + "] 正在获取 " + cveName + " URL: " + cvea)
		cveResp, err := http.Get(cvea)
		if err != nil {
			continue
		}
		defer cveResp.Body.Close()
		cveBody, err := ioutil.ReadAll(cveResp.Body)
		if err != nil {
			continue
		}
		//fmt.Println(string(cveBody))
		re, err = regexp.Compile(`<tr>
		<td colspan="2">(.*?)

</td>`)
		if err != nil {
			continue
		}
		desc := re.FindAllStringSubmatch(string(cveBody), -1)[0][1]
		//break
		var s []string
		s = append(s, cveName, cvea, desc)
		resultSlice = append(resultSlice, s)

	}

	return resultSlice
}
