package web

import (
    "net/http"
    "encoding/json"
    "fmt"

    "git.eju-inc.com/ops/go-common/log"

    "github.com/bypdhu/ssh-executor/module"
    "github.com/bypdhu/ssh-executor/conf"
    "github.com/bypdhu/ssh-executor/result"
)

var (
    cs map[string]*conf.Config
)

func RunJob(w http.ResponseWriter, r *http.Request) {

    log.Infof("Now Run Job.\n")

    defer r.Body.Close()
    buf := make([]byte, 1024)
    var content string
    for {
        n, err := r.Body.Read(buf)
        content += string(buf[:n])
        if err != nil {
            break
        }
    }

    log.Infof("The request body is:%s\n", content)

    wb := &WebBody{}

    err := json.Unmarshal([]byte(content), wb)
    if err != nil {
        buildNewResponse(w, 500, result.Result{Success:false, Msg:err.Error()})
        return
    }

    log.Infof("The hosts is:%s\n", wb.Hosts)

    if wb.UserFlag != "" {
        wb.SSHConfig.UserName = UserMap[wb.UserFlag].UserName
        wb.SSHConfig.Password = UserMap[wb.UserFlag].Password
    }

    C.Tasks = wb.Tasks

    conf.CopySSHConfig(&wb.SSHConfig, &C.SSHConfig)

    cs = conf.GetCopiedConfigMap(C, wb.Hosts)

    module.RunAll(cs)

    buildNewResponse(w, 200, conf.GenerateResult(cs))
}

func buildNewResponse(w http.ResponseWriter, code int, result result.Result) {

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)

    formatRes, err := json.Marshal(result)

    if err != nil {
        log.Errorf("Failed to format response str, err is %s", err)
        errByte := []byte(fmt.Sprintf(`{"msg":"%s"}`, err))
        w.Write(errByte)
        return
    }

    w.Write(formatRes)
}

func Hello(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    w.Write([]byte(fmt.Sprintf(`{"msg":"%s"}`, "hello world.")))
    return
}