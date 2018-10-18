package web

import (
    "net/http"

    "github.com/gorilla/mux"
    "git.eju-inc.com/ops/go-common/log"

    "github.com/bypdhu/ssh-executor/conf"
)

type users struct {
    UserName string
    Password string
}

var (
    C *conf.Config
    UserMap map[string]users
)

func Run(c *conf.Config) {
    C = c
    initUser()

    for _, u := range c.Serv.Users {
        UserMap[u.Type] = users{UserName:u.UserName, Password:u.Password}
    }

    router := mux.NewRouter()
    router.HandleFunc("/job", RunJob).Methods("POST")
    router.HandleFunc("/hello", Hello).Methods("GET", "POST", "HEAD")

    log.Infof("Server now launched.")
    log.Fatal(http.ListenAndServe(c.Serv.Web.ListenAddress, router))
}
