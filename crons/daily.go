package crons

import (
	"fmt"
	"github.com/BaizeSec/Pusher/common"
	"github.com/BaizeSec/Pusher/lib"
	"github.com/BaizeSec/Pusher/pkg/logger"
	"github.com/BaizeSec/Pusher/sources"
	"os"
	"strconv"
	"strings"
	"time"
)

func Daily() {
	daily := "现在是 "
	daily += time.Now().Format("2006-01-02 15:04")
	daily += " 开始搜集今日推送.."
	logger.Info(daily)

	source := ""
	source += sources.GetEdge()
	source += sources.GetSeePaper()
	source += sources.GetAnquanke()

	file := sources.GetCves()

	logger.Info("正在写 CVEs 文件...")
	writeToFile(file)

	dailyMessage := ""

	for _, groupid := range common.GROUP_IDS {
		sp := strings.Split(groupid, ",")
		dailyMessage = "Bingan Pusher, License: " + sp[0] + "\n"
		if len(source) == 0 {
			dailyMessage += "棱角日报、seebug、安全客今天都是懒狗，么么叽～"
		} else {
			dailyMessage += source
			dailyMessage += "小饼干推送已经整合棱角日报、seebug、安全客\n食用愉快么么叽～"
		}
		group, _ := strconv.Atoi(sp[1])
		lib.SendMsg(group, dailyMessage)
		lib.SendFile(group)
		logger.Success("对于授权群 [" + sp[0] + "]: " + sp[1] + ", 的推送已经完成")
	}

	logger.Success("已经成功推送今日所有消息！")
}

func writeToFile(msg string) {
	t := time.Now()
	filename := "CVE-" + t.Format("2006-01-02") + ".md"
	wd, _ := os.Getwd()
	f, err := os.Create("cves/" + filename)
	if err != nil {
		panic(err)
		fmt.Println(wd + "/cves/" + filename)
	} else {
		_, err = f.Write([]byte(msg))
		if err != nil {

		}
	}
	f.Close()
}
