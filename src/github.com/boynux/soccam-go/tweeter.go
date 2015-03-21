package soccam

import (
    "github.com/boynux/anaconda"

    "appengine"
    "appengine/urlfetch"
)

func initialize(ctx appengine.Context) *Soccam {
    configuration := GetConfig()

    anaconda.SetConsumerKey(configuration.TwitterConsumerKey)
    anaconda.SetConsumerSecret(configuration.TwitterConsumerSecret)

    anaconda.SetHttpTransport(&urlfetch.Transport{ Context: ctx })
    api := &Soccam {
        TwitterApi: anaconda.NewTwitterApi(configuration.TwitterAccessToken, configuration.TwitterAccessTokenSecret),
    }

    return api
}

func (s Soccam) tweet(message string) (err error) {
    _, err = s.TwitterApi.PostTweet(message, nil)

    return err
}

