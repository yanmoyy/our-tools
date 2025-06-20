package msg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/yanmoyy/our-tools/internal/sns/kakao/auth"
	"github.com/yanmoyy/our-tools/internal/sns/kakao/header"
)

type TextObject struct {
	ObjectType string `json:"object_type"`
	Text       string `json:"text"`
	Link       struct {
		WebURL       string `json:"web_url"`
		MobileWebURL string `json:"mobile_web_url"`
	} `json:"link"`
	ButtonTitle string `json:"button_title"`
}

func requestSendToMe(token, message string) error {
	reqBody := url.Values{}
	textObject := TextObject{
		ObjectType: "text",
		Text:       message,
	}
	data, err := json.Marshal(textObject)
	if err != nil {
		return err
	}
	reqBody.Add("template_object", string(data))
	req, err := http.NewRequest("POST", sendToMeURL, strings.NewReader(reqBody.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create request: %s", err)
	}

	req.Header.Set("Content-Type", header.UrlEncoded)
	req.Header.Set("Authorization", auth.GetBearerToken(token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send message: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// print body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %s", err)
		}
		fmt.Println(string(body))
		return fmt.Errorf("failed to send message (status: %s)", resp.Status)
	}
	return nil
}
