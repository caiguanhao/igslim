package igslim

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var (
	ErrNoSuchUser = errors.New("no such user")
	ErrFailed     = errors.New("failed to get user")
)

const (
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36"

	scriptOpen  = "window._sharedData = "
	scriptClose = "</script>"
)

type (
	Client struct {
		SessionId string
	}

	user struct {
		Id           string `json:"id"`
		FbId         string `json:"fbid"`
		FullName     string `json:"full_name"`
		Verified     bool   `json:"is_verified"`
		UserName     string `json:"username"`
		Picture      string `json:"profile_pic_url"`
		PictureHd    string `json:"profile_pic_url_hd"`
		Biography    string `json:"biography"`
		CategoryName string `json:"category_name"`
		Followings   struct {
			Count int `json:"count"`
		} `json:"edge_follow"`
		Followers struct {
			Count int `json:"count"`
		} `json:"edge_followed_by"`
		Posts struct {
			Count int `json:"count"`
		} `json:"edge_owner_to_timeline_media"`
	}

	User struct {
		Id              int
		FbId            int64
		UserName        string
		FullName        string
		Verified        bool
		Picture         string
		Biography       string
		CategoryName    string
		FollowingsCount int
		FollowersCount  int
		PostsCount      int
	}
)

// Create new client with session id.
func NewClient(sessionId string) *Client {
	return &Client{
		SessionId: sessionId,
	}
}

// GetUser wraps GetUserWithContext using context.Background.
func (client Client) GetUser(username string) (*User, error) {
	return client.GetUserWithContext(context.Background(), username)
}

// GetUserWithContext returns user info or error.
func (client Client) GetUserWithContext(ctx context.Context, username string) (*User, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://www.instagram.com/"+username+"/", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Cookie", "sessionid="+client.SessionId)
	req.Header.Add("User-Agent", userAgent)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		return nil, ErrNoSuchUser
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	user := getProfileFromHtml(string(body))
	if user == nil {
		return nil, ErrFailed
	}
	return user.User(), nil
}

func (user *user) User() *User {
	picture := user.PictureHd
	if picture == "" {
		picture = user.Picture
	}
	id, _ := strconv.Atoi(user.Id)
	fbid, _ := strconv.ParseInt(user.FbId, 10, 64)
	return &User{
		Id:              id,
		FbId:            fbid,
		UserName:        user.UserName,
		FullName:        user.FullName,
		Verified:        user.Verified,
		Picture:         picture,
		Biography:       user.Biography,
		CategoryName:    user.CategoryName,
		FollowingsCount: user.Followings.Count,
		FollowersCount:  user.Followers.Count,
		PostsCount:      user.Posts.Count,
	}
}

func getProfileFromHtml(html string) *user {
	i := strings.Index(html, scriptOpen)
	if i < 0 {
		return nil
	}
	html = html[i+len(scriptOpen):]
	i = strings.Index(html, scriptClose)
	if i < 0 {
		return nil
	}
	html = strings.TrimSuffix(html[:i], ";")
	bytes := []byte(html)
	var obj map[string]map[string][]map[string]map[string]*user
	json.Unmarshal(bytes, &obj)
	data := obj["entry_data"]["ProfilePage"]
	if len(data) > 0 {
		return data[0]["graphql"]["user"]
	}
	return nil
}
