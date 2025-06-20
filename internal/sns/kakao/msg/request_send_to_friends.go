package msg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/yanmoyy/our-tools/internal/sns/kakao/auth"
	"github.com/yanmoyy/our-tools/internal/sns/kakao/header"
)

func requestSendToFriends(token string, uuids []string, message string) error {
	reqBody := url.Values{}
	textObject := textObject{
		ObjectType: "text",
		Text:       message,
	}
	textObjectJson, err := json.Marshal(textObject)
	if err != nil {
		return err
	}
	reqBody.Add("template_object", string(textObjectJson))
	uuidsJson, err := json.Marshal(uuids)
	if err != nil {
		return err
	}
	reqBody.Add("receiver_uuids", string(uuidsJson))
	req, err := http.NewRequest("POST", sentToFriendsURL, strings.NewReader(reqBody.Encode()))

	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", header.UrlEncoded)
	req.Header.Set("Authorization", auth.GetBearerToken(token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message (status: %s)", resp.Status)
	}
	return nil
}
