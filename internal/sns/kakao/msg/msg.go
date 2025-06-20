package msg

const (
	sendToMeURL = "https://kapi.kakao.com/v2/api/talk/memo/default/send"
)

func Send(token, target, message string) error {
	if target == "me" {
		return requestSendToMe(token, message)
	}
	// TODO: send to specific user
	return nil
}
