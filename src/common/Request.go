package common

import (
    "net/http"
    "../model"
    //"io/ioutil"
    //"encoding/json"
)

type Request struct {
    *http.Request
    PathParams map[string]string
    Env        map[string]interface{}
    *model.Permission
    BodyMap    map[string]interface{}
}

//get 请求
func (c *Request) GetQueryArray(key string) ([]string, bool) {
    if values, ok := c.Request.URL.Query()[key]; ok && len(values) > 0 {
        return values, true
    }
    return []string{}, false
}

func (c *Request) GetQuery(key string) (string, bool) {
    if values, ok := c.GetQueryArray(key); ok {
        return values[0], ok
    }
    return "", false
}

func (c *Request) DefaultQuery(key, defaultValue string) string {
    if value, ok := c.GetQuery(key); ok {
        return value
    }
    return defaultValue
}

//post Text  请求
//func (c *Request) GetPostValue(key, defaultValue string) string {
//    _, ok := c.BodyMap["systemAuto"];
//    if ok{
//        value, ok := c.BodyMap[key];
//        if ok {
//            return value.(string)
//        }
//        return defaultValue
//    }
//    result, _ := ioutil.ReadAll(c.Body)
//    c.Body.Close()
//    //未知类型的推荐处理方法
//    var f interface{}
//    json.Unmarshal(result, &f)
//    m := f.(map[string]interface{})
//    m["systemAuto"] = "1"
//    c.BodyMap = m;
//    value, ok := m[key];
//    if ok {
//        return value.(string)
//    }
//    return defaultValue
//}

//post From请求

//func (c *Request) GetPostFormArray(key string) ([]string, bool) {
//    req := c.Request
//    req.ParseForm()
//    req.ParseMultipartForm(c.engine.MaxMultipartMemory)
//    if values := req.PostForm[key]; len(values) > 0 {
//        return values, true
//    }
//    if req.MultipartForm != nil && req.MultipartForm.File != nil {
//        if values := req.MultipartForm.Value[key]; len(values) > 0 {
//            return values, true
//        }
//    }
//    return []string{}, false
//}
//
//
//func (c *Request) GetPostForm(key string) (string, bool) {
//    if values, ok := c.GetPostFormArray(key); ok {
//        return values[0], ok
//    }
//    return "", false
//}
//
//func (c *Request) PostForm(key string) string {
//    value, _ := c.GetPostForm(key)
//    return value
//}
