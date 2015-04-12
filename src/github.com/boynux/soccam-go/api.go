package soccam

import (
    "fmt"
    "net/http"
    "net/url"

    "appengine"
)

var soccam Soccam

func init() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/tweet", tweet)
    http.HandleFunc("/followers", followers)
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

func followers(w http.ResponseWriter, r *http.Request) {
    soccam = *initialize(appengine.NewContext(r))

    followersPageChannel := soccam.followers(url.Values {})
    for followersPage := range followersPageChannel {
        for _, user := range followersPage.Followers {
            fmt.Fprint(w, user.ScreenName)
        }
    }
}
