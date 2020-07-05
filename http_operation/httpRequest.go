package http_operation

import (
    "fmt"
    "strings"
    "net/http"
    //"io/ioutil"
    "net/url"
)

func HttpRequest(method, urlVal, data string)  *http.Response {
    
    client := &http.Client{}
    var req *http.Request
 
    if data == "" {
        urlArr := strings.Split(urlVal,"?")
        if len(urlArr)  == 2 {
            urlVal = urlArr[0] + "?" + getParseParam(urlArr[1])
        }
        req, _ = http.NewRequest(method, urlVal, nil)
    }else {
        req, _ = http.NewRequest(method, urlVal, strings.NewReader(data))
    }
 
    //添加header，key为X-Xsrftoken，value为b6d695bbdcd111e8b681002324e63af81
    //req.Header.Add("X-Xsrftoken","b6d695bbdcd111e8b681002324e63af81")
 
    resp, err := client.Do(req)
 
    if err != nil {
        fmt.Println("something wrong")
    }

    return resp
}


func getParseParam(param string) string  {
    return url.PathEscape(param)
}

/*func setCookie(req) {

    //可以添加多个cookie
    cookie := &http.Cookie{
        Name: "X-Xsrftoken",
        Value: "df41ba54db5011e89861002324e63af81", 
        HttpOnly: true}
    req.AddCookie(cookie)
}*/


