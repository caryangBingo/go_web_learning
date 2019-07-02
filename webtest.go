/*
* @Author   : caryangBingo
* @Date     : 2019-06-24 10:49:31
* @Filename : web_go.py
* @Version  : Version 1.0
*/

package main

import (
    "fmt"
    "html/template"
    "net/http"
    "strings"
    "log"
)

func sayhelloName(w http.ResponseWriter,r *http.Request){
    r.ParseForm()
    fmt.Println(r.Form)
    fmt.Println("path",r.URL.Path)
    fmt.Println("scheme",r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k,v :=range r.Form{
        fmt.Println("key:",k)
        fmt.Println("val:",strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello astaxie!")
}

func login(w http.ResponseWriter,r *http.Request) {
    fmt.Print("method:", r.Method)
    if r.Method == "GET"{
        t,_ := template.ParseFiles("login.gtpl")
        t.Execute(w,nil)
    } else {
        r.ParseForm()
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
    }
}

func main() {
    http.HandleFunc("/", sayhelloName)
    http.HandleFunc("/login", login)
    err:=http.ListenAndServe(":9090", nil)
    if err!=nil{
        log.Fatal("ListenAndServe: ",err)
    }
}

/*
import (
    "fmt"
    "net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter,r *http.Request) {
    if r.URL.Path =="/" {
        sayhelloName(w,r)
        return
    }
    http.NotFound(w, r)
    return
}

func sayhelloName(w http.ResponseWriter,r *http.Request) {
    fmt.Fprintf(w, "Hello myroute!")
}

func main() {
    mux:=&MyMux{}
    http.ListenAndServe(":9090", mux)
}
*/
