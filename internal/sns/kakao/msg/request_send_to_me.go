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

func requestSendToMe(token, message string) error {
	reqBody := url.Values{}
	textObject := textObject{
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
