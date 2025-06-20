package msg

const (
	sendToMeURL      = "https://kapi.kakao.com/v2/api/talk/memo/default/send"
	sentToFriendsURL = "https://kapi.kakao.com/v1/api/talk/friends/message/send"
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

func Send(token, target, message string) error {
	if target == "me" {
		return requestSendToMe(token, message)
	}
	// TODO: send to specific user
	return nil
}
