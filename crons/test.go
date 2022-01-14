package crons

import (
	"github.com/BaizeSec/Pusher/lib"
	"github.com/BaizeSec/Pusher/pkg/logger"
	"github.com/BaizeSec/Pusher/sources"
	"strconv"
	"strings"
	"time"
)

func Test() {
	daily := "现在是 "
	daily += time.Now().Format("2006-01-02 15:04")
	daily += " 开始搜集今日推送.."
	logger.Info(daily)

	source := ""
	source += sources.GetEdge()
	source += sources.GetSeePaper()
	source += sources.GetAnquanke()

	file := ""
	if !lib.Cve {
		file = sources.GetCves()
	} else {
		file = "test"
	}

	logger.Info("正在写 CVEs 文件...")
	writeToFile(file)

	dailyMessage := ""

	sp := strings.Split(lib.GroupID, ",")
	dailyMessage = "Bingan Pusher, License: " + sp[0] + "\n"
	dailyMessage += source
	dailyMessage += "小饼干推送已经整合棱角日报、seebug、安全客\n食用愉快么么叽～"
	group, _ := strconv.Atoi(sp[1])
	lib.SendMsg(group, dailyMessage)
	lib.SendFile(group)
	logger.Success("对于授权群 [" + sp[0] + "]: " + sp[1] + ", 的推送已经完成")

	logger.Success("已经成功推送今日所有消息！")
}
