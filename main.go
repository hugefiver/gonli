package main

import (
    "encoding/json"
    "flag"
    "io/ioutil"
    "log"

    . "gonli/loginmod"
)

var (
    user *string
    password *string
    userFile *string
    ua *string
    conf *string
    help *bool
)

func init()  {
    user = flag.String("u", "", "user")
    password = flag.String("p", "", "password")
    userFile = flag.String("U", "", "users file")
    ua = flag.String("ua", "", "user agent")
    conf = flag.String("conf", "", "")
    help = flag.Bool("h", false, "help")
    flag.Parse()
}


func loginOnce(user, password, ua string) (ok bool) {
    ok = Login(user, password, ua)
    if ok {
        log.Printf("Login Successfully. Id: <%s>.\n", user)
    } else {
        log.Printf("Login Failed. Id:<%s>.\n", user)
    }
    return
}

func loginUntilSuccess(users []User) (ok bool) {
    for _, u := range users {
        ok := loginOnce(u.Id, u.Password, *ua)
        if ok {
            return true
        }
    }
    return false
}

func main()  {
    if *help {
        flag.Usage()
        return
    }

    if *user != "" {
        loginOnce(*user, *password, *ua)
    } else if *userFile != "" {
        data, err := ioutil.ReadFile(*userFile)
        if err != nil {
            log.Fatalln(err)
            return
        }

        users := make([]User, 0)
        err = json.Unmarshal(data, user)
        if err != nil {
            log.Fatalln(err)
            return
        }
        loginUntilSuccess(users)
    } else if *conf != "" {
        log.Fatal("Not Support Config Currently.")
    } else {
        flag.Usage()
    }
}