package loginmod

import (
    "errors"
    "log"
    "strings"

    "github.com/imroc/req"
)

const LOGIN_URL = "http://192.168.7.221:801/eportal/?" +
    "c=ACSetting&a=Login&protocol=http:&iTermType=1&" +
    "enAdvert=0&queryACIP=0&loginMethod=1"


func Login(user, password, ua string) (ok bool) {
    if ua == "" {
        ua = "Mozilla/5.0 (Windows NT 10.0; WOW64; rv:56.0) Gecko/20100101 Firefox/56.0"
    }

    err := connect(user, password, ua)
    if err != nil {
        log.Println(err)
        return false
    } else {
        return true
    }
}

func connect(user, password, ua string) error {
    client := req.New()

    header := req.Header{
        "User-Agent": ua,
        "Content-Type": "application/x-www-form-urlencoded",
    }

    body := req.Param{
        "DDDDD": ",0," + user,
        "upass": password,
    }

    resp, err := client.Post(LOGIN_URL, header, body)
    if err != nil {
        return err
    }


    r := resp.Response()
    status := r.StatusCode
    if status != 200 {
        return errors.New("request error")
    }

    text := resp.String()
    ok := strings.Contains(text, "认证成功页") || strings.Contains(text, "login again")
    if ok {
        return nil
    } else {
        return errors.New("failed to login")
    }
}