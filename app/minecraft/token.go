package minecraft

import (
	"github.com/restartfu/gophig"
	"github.com/sandertv/gophertunnel/minecraft/auth"
	"golang.org/x/oauth2"
)

var src oauth2.TokenSource

func init() {
	var err error
	src, err = resolveToken()
	if err != nil {
		panic(err)
	}
}

func resolveToken() (oauth2.TokenSource, error) {
	token := new(oauth2.Token)
	g := gophig.NewGophig("assets/token", "json", 699)
	err := g.GetConf(token)
	if err != nil {
		token, err = auth.RequestLiveToken()
		if err != nil {
			return nil, err
		}
	}

	src := auth.RefreshTokenSource(token)
	_, err = src.Token()
	if err != nil {
		// The cached refresh token expired and can no longer be used to obtain a new token. We require the
		// user to log in again and use that token instead.
		token, err = auth.RequestLiveToken()
		src = auth.RefreshTokenSource(token)
	}

	err = g.SetConf(token)
	return src, err
}
