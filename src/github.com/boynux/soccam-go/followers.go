package soccam

import (
    "github.com/boynux/anaconda"

    "net/url"
)

func (s Soccam) followers(v url.Values) (result chan anaconda.FollowersPage) {
    return s.TwitterApi.GetFollowersListAll(url.Values {})
}
