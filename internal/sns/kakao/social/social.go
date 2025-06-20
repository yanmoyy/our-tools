package social

const (
	getFriendsURL = "https://kapi.kakao.com/v1/api/talk/friends"
)

func GetFriends(token string) ([]friend, error) {
	resp, err := requestGetFriends(token)
	if err != nil {
		return nil, err
	}
	return resp.Elements, nil
}
