package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"thchat/pkg/config"
	"thchat/pkg/e"
	"thchat/pkg/util"
)
type GithubUser struct {
	Login	string	`json:"login"`
	Email 	string	`json:"email"`
}

type AccessTokenResponse struct{
	AccessToken		string	`json:"access_token"`
	Scope  			string	`json:"scope"`
	TokenType 		string	`json:"token_type"`
}

func RequestGithubAccessToken(code string) (string, error) {
	url := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		config.AppConfig.Oauth.GithubClientID,
		config.AppConfig.Oauth.GithubClientSecret,
		code)
	fmt.Println(url)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "application/json")

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		return "", err
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	jsonRes := AccessTokenResponse{}
	err = json.Unmarshal(bodyBytes, &jsonRes)
	if err != nil {
		return "", err
	}
	accessToken := jsonRes.AccessToken
	fmt.Println(accessToken)
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
	//fmt.Println("get token:\t"+token)
	userInfo, err := RequestGithubUser(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": e.ErrAuthToken, "message": "Cannot get github user info"})
		return
	}
	username := userInfo.Login
	err = util.SetSession(c, username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": "Setting sessions failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": "get access token ok", "username": username})
}

func GithubOauthCallback(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code":e.SUCCESS, "message": "callback ok"})
}


func RequestGithubUser(token string) (GithubUser, error){
	githubUser := GithubUser{}
	url := "https://api.github.com/user"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return githubUser, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("token %s", token))

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		return githubUser, err
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return githubUser, err
	}
	//fmt.Println(string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &githubUser)
	if err != nil {
		return githubUser, err
	}
	return githubUser, nil
}

