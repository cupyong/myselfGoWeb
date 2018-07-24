///*
//获取 token (GET)
//https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=$corpid&corpsecret=$corpsecret
//
//发送信息 (POST) application/json
//https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=qDWmOsaH5QAleB
//{
//   "touser": "UserID1|UserID2|UserID3",
//   "toparty": " PartyID1 | PartyID2 ",
//   "totag": " TagID1 | TagID2 ",
//   "msgtype": "text",
//   "agentid": 3,
//   "text": {
//       "content": "Holiday Request For Pony(http://xxxxx)"
//   },
//   "safe":"0"
//}
//*/
//
package common
//
//import "fmt"
//import "log"
//import "net/http"
//import "io/ioutil"
//import "strings"
//import "encoding/json"
//import "flag"
//import "os"
//
//var corpid = "wx71a8339b7df6c04c"
//var corpsecret = "IZUitB3fnJzaPAugUCKGCUvCSfqEogXR2WySWJCJg1a81qFTYHMfgf1-bmDPTmel"
//
//func main() {
//
//    // 参数解析
//    user_list := flag.String("u", "", "to user list: 'user1,user2', 'user1|user2', '@all'")
//    msg := flag.String("m", "", "'text Message'")
//    flag.Parse()
//    if len(*user_list) == 0 || len(*msg) == 0 {
//        fmt.Println("Usage of", os.Args[0])
//        flag.PrintDefaults()
//        os.Exit(2)
//    }
//    *user_list = strings.Replace(*user_list, ",", "|", -1)
//
//    var err error
//
//    // 获取token
//    token_api := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", corpid, corpsecret)
//
//    res, err := http.Get(token_api)
//    if err != nil {
//        log.Fatal("http error:", err)
//        return
//    }
//    data, err := ioutil.ReadAll(res.Body)
//    res.Body.Close()
//    if err != nil {
//        log.Fatal(err)
//        return
//    }
//
//    type Token_t struct {
//        Errcode      int
//        Errmsg       string
//        Access_token string
//        Expires_in   int
//    }
//    var token Token_t
//
//    err = json.Unmarshal(data, &token)
//
//    if 0 < token.Errcode {
//        log.Fatal(token.Errmsg)
//        return
//    }
//
//    // 发送信息
//    msg_api := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", token.Access_token)
//
//    //msg_json := fmt.Sprintf("{ \"touser\": \"cnzhangquan\", \"msgtype\":\"text\", \"agentid\":3, \"text\":{ \"content\": \"test @all msg_json\"}}")
//    //msg_json := fmt.Sprintf("{ \"touser\": \"huangyiqiang\", \"msgtype\":\"text\", \"agentid\":3, \"text\":{ \"content\": \"test @all msg_json\"}}")
//    //msg_json := fmt.Sprintf("{ \"touser\": \"cnzhangquan|huangyiqiang\", \"msgtype\":\"text\", \"agentid\":3, \"text\":{ \"content\": \"test @all msg_json\"}}")
//    //msg_json := fmt.Sprintf("{ \"touser\": \"@all\", \"msgtype\":\"text\", \"agentid\":3, \"text\":{ \"content\": \"test @all msg_json\"}}")
//    msg_json := fmt.Sprintf("{ \"touser\": \"%s\", \"msgtype\":\"text\", \"agentid\":3, \"text\":{ \"content\": \"[秒针] %s\"}}", *user_list, *msg)
//
//    res, err = http.Post(msg_api, "application/json", strings.NewReader(msg_json))
//    if err != nil {
//        log.Fatal(err)
//        return
//    }
//    data, err = ioutil.ReadAll(res.Body)
//    res.Body.Close()
//    if err != nil {
//        log.Fatal(err)
//        return
//    }
//
//    type result_t struct {
//        Errcode      int
//        Errmsg       string
//        Invaliduser  string
//        Invalidparty string
//        Invalidtag   string
//    }
//    var result result_t
//
//    err = json.Unmarshal(data, &result)
//    if 0 < result.Errcode {
//        log.Fatal(result.Errmsg, result.Invaliduser)
//        return
//    }
//
//    if len(result.Invaliduser) > 0 {
//        log.Fatal("Invalid User:", result.Invaliduser)
//        return
//    }
//
//    fmt.Println("Success")
//    return
//}

