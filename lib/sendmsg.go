package lib

import (
	"github.com/BaizeSec/Pusher/common"
	"github.com/BaizeSec/Pusher/utils"
	"strconv"
)

func SendMsg(groupId int, message string) {

	body := ""
	body += "{\"group_id\": "
	body += strconv.Itoa(groupId)
	body += ",\"message\": \""
	body += message
	body += "\",\"auto_escape\": false}"

	groupUrl := common.SERVER_URL + "/send_group_msg"

	utils.PostJson(groupUrl, body)
}
