package hello

import (
    "fmt"
    "net/http"

    "appengine"
)

var soccam Soccam

func init() {

    http.HandleFunc("/", handler)
    http.HandleFunc("/tweet", tweet)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, Boynux!")
}

func tweet(w http.ResponseWriter, r *http.Request) {
    soccam = *initialize(appengine.NewContext(r))

    err := soccam.tweet("test")

    if err == nil {
        fmt.Fprint(w, "Successful")
    } else {
        fmt.Fprint(w, err.Error())
    }

}
