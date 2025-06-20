package msg

const (
	sendToMeURL      = "https://kapi.kakao.com/v2/api/talk/memo/default/send"
	sentToFriendsURL = "https://kapi.kakao.com/v1/api/talk/friends/message/default/send"
)

type textObject struct {
	ObjectType string `json:"object_type"`
	Text       string `json:"text"`
	Link       struct {
		WebURL       string `json:"web_url"`
		MobileWebURL string `json:"mobile_web_url"`
	} `json:"link"`
	ButtonTitle string `json:"button_title"`
}

func SendToMe(token, message string) error {
	return requestSendToMe(token, message)
}

func SendToFriend(token, receiverUUID, message string) error {
	return requestSendToFriends(token, []string{receiverUUID}, message)
}
