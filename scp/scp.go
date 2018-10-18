package scp

import (
    "golang.org/x/crypto/ssh"
    "os/user"
    "strings"
    "path/filepath"
    "os"
    "io"
    "fmt"
)

type SCPClient struct {
    Connection *ssh.Client
    Session    *ssh.Session
    Permission bool
}

func unset(s []string, i int) []string {
    if i >= len(s) {
        return s
    }
    return append(s[:i], s[i + 1:]...)
}

func getFullPath(path string) (string, error) {
    _user, err := user.Current()
    if err != nil {
        return path, err
    }
    fullPath := strings.Replace(path, "~", _user.HomeDir, 1)
    fullPath, err = filepath.Abs(fullPath)
    if err != nil {
        return path, err
    }
    return fullPath, nil
}

func walkDir(root string) (files []string, err error) {
    err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if info.IsDir() {
            path = path + "/"
        }
        files = append(files, path)
        return nil
    })
    return
}

func pushDirData(w io.WriteCloser, baseDir string, path string, toName string, perm bool) {
    baseDirSlice := strings.Split(baseDir, "/")
    baseDirSlice = unset(baseDirSlice, len(baseDirSlice) - 1)
    baseDir = strings.Join(baseDirSlice, "/")

    relPath, _ := filepath.Rel(baseDir, path)
    dir := filepath.Dir(relPath)

    if len(dir) > 0 && dir != "." {
        dirList := strings.Split(dir, "/")
        dirPath := baseDir
        for _, dirName := range dirList {
            dirPath = dirPath + "/" + dirName
            dInfo, _ := os.Stat(dirPath)
            dPerm := fmt.Sprintf("%04o", dInfo.Mode().Perm())

            fmt.Fprintln(w, "D" + dPerm, 0, dirName)
        }
    }

    fInfo, _ := os.Stat(path)

    if !fInfo.IsDir() {
        pushFileData(w, path, toName, perm)
    }

    if len(dir) > 0 && dir != "." {
        dirList := strings.Split(dir, "/")
        end_str := strings.Repeat("E\n", len(dirList))
        fmt.Fprintf(w, end_str)
    }
    return

}
func pushFileData(w io.WriteCloser, path string, toName string, perm bool) {
    content, err := os.Open(path)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }

    stat, _ := content.Stat()
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }

    fInfo, _ := os.Stat(path)

    fPerm := "0644"
    if perm == true {
        fPerm = fmt.Sprintf("%04o", fInfo.Mode())
    }

    fmt.Fprintln(w, "C" + fPerm, stat.Size(), toName)
    io.Copy(w, content)
    fmt.Fprint(w, "\x00")

    return
}