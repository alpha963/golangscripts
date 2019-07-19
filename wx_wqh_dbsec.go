package main

import (
    "bytes"
    "github.com/json-iterator/go"
    "os"
    "fmt"
    "net/http"
    "io/ioutil"
    "github.com/PuerkitoBio/goquery"
    "strings"
    "log"
    "time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type JSON struct {
    Access_token string `json:"access_token"`
}

type MESSAGES struct {
    Touser string `json:"touser"`
    Toparty string `json:"toparty"`
    Msgtype string `json:"msgtype"`
    Agentid int `json:"agentid"`
    Text struct {
        Content string `json:"content"`
    } `json:"text"`
    Safe int `json:"safe"`
}

func Get_AccessToken(corpid,corpsecret string) string {
    gettoken_url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + corpid + "&corpsecret=" + corpsecret

    client := &http.Client{}
    req, _ := client.Get(gettoken_url)
    defer req.Body.Close()
    body, _ := ioutil.ReadAll(req.Body)
 
    var json_str JSON
    json.Unmarshal([]byte(body), &json_str)

    return json_str.Access_token
}

func Send_Message(access_token,msg string) {
    send_url := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + access_token
    client := &http.Client{}
    req, _ := http.NewRequest("POST", send_url, bytes.NewBuffer([]byte(msg)))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("charset","UTF-8")
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
}

func messages(touser string,toparty string,agentid int,content string) string{
    msg := MESSAGES{
        Touser: touser,
        Toparty: toparty,
        Msgtype: "text",
        Agentid: agentid,
        Safe: 0,
        Text: struct {

            Content string `json:"content"`
        }{Content: content},
    }
    sed_msg, _ := json.Marshal(msg)

    return string(sed_msg)
}

func main() {
    url := "http://192.168.7.128/GetKey/"

    res, err := http.Get(url)

    if err != nil {
        fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
        os.Exit(1)
    }

    body, err := ioutil.ReadAll(res.Body)

    res.Body.Close()

    if err != nil {
        fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
        os.Exit(1)
    }

  html := string(body)

	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(html))
	if err!=nil{
		log.Fatalln(err)
	}

    var timeLayoutStr = "2019-07-22 15:09:05" 
	
	  touser := "@all"  
    toparty := "149"       
    agentid:= 1000006 
    corpid := "wwad7173798d430041"  
    corpsecret := "Mi3tb3cAJnMm1goKsU3ls0pn2u8Fapyw9MO0ptph6Ig"  
    accessToken := Get_AccessToken(corpid,corpsecret)
    content := "Today's key and passwd:\n"+time.Now().UTC().Format(timeLayoutStr)
    
  msg := strings.Replace(messages(touser,toparty,agentid,content),"\\\\","\\",-1)
	
	Send_Message(accessToken,msg)
	
	dom.Find(".form-control").Each(func(i int, selection *goquery.Selection) {	

		content=selection.Text();
		msg := strings.Replace(messages(touser,toparty,agentid,content),"\\\\","\\",-1)

    Send_Message(accessToken,msg)
				
		})
	
}
