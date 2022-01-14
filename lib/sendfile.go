package lib

import (
	"github.com/BaizeSec/Pusher/common"
	"github.com/BaizeSec/Pusher/utils"
	"os"
	"strconv"
	"time"
)

func SendFile(groupId int) {

	t := time.Now()
	filename := "CVE-" + t.Format("2006-01-02") + ".md"
	wd, _ := os.Getwd()
	fullfile := wd + "/cves/" + filename

	body := ""
	body += "{\"group_id\": "
	body += strconv.Itoa(groupId)
	body += ",\"file\": \""
	body += fullfile
	body += "\",\"name\": \""
	body += filename
	body += "\"}"

	groupUrl := common.SERVER_URL + "/upload_group_file"

	utils.PostJson(groupUrl, body)
}
