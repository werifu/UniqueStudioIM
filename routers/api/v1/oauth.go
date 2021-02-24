package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
	"thchat/pkg/config"
	"thchat/pkg/e"
)
type GithubUser struct {
	Login	string	`json:"login"`
	Email 	string	`json:"email"`
}

func RequestGithubAccessToken(code string) (string, error) {
	url := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		config.AppConfig.Oauth.GithubClientID,
		config.AppConfig.Oauth.GithubClientSecret,
		code)
	//fmt.Println(url)
	res, err := http.Post(url, "text/plain", nil)
	if err != nil {
		return "", err
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	accessToken := res2token(bodyBytes)

	return accessToken, nil
}

func OauthGithub(c *gin.Context) {
	code, ok := c.GetPostForm("code")

	if !ok || code == "" {
		c.JSON(http.StatusOK, gin.H{"code": e.ErrAuth, "message": "oauth code error"})
		return
	}
	token, err := RequestGithubAccessToken(code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": e.ErrAuth, "message": "Cannot request a github access token."})
		return
	}
	userInfo, err := RequestGithubUser(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": e.ErrAuthToken, "message": "Cannot get github user info"})
		return
	}
	username := userInfo.Login
	fmt.Println(username)

	c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": "get access token ok"})
}

func GithubOauthCallback(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code":e.SUCCESS, "message": "callback ok"})
}


func RequestGithubUser(token string) (GithubUser, error){
	githubUser := GithubUser{}
	url := fmt.Sprintf("https://api.github.com/user?access_token=%s", token)
	res, err := http.Get(url)
	if err != nil {
		return githubUser, err
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return githubUser, err
	}
	fmt.Println(string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &githubUser)
	if err != nil {
		return githubUser, err
	}
	return githubUser, nil
}


func res2token(body []byte) string {
	bodyString := string(body)
	return strings.Split(strings.Split(bodyString, "&")[0], "=")[1]
}