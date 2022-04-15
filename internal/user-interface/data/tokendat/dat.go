package tokendat

import (
	"userserver/internal/pkg/cache"

	"github.com/pkg/errors"
)

type TokenDatInterface interface {
	Init() error
	SetToken(token, userid string, auth int64) error
	GetToken(token string) (string, int64, error)
	DelToken(token string) error
}

type TokenDat struct {
	cache cache.TokenCacheInterface
}

func NewTokenDat(cache cache.TokenCacheInterface) *TokenDat {
	return &TokenDat{cache: cache}
}

func (dat *TokenDat) Init() error {
	if dat.cache == nil {
		return errors.New("TokenDat Init failed: cache is unset.")
	}
	if err := dat.cache.Init(); err != nil {
		return errors.WithMessage(err, "dat cache init failed")
	}
	return nil
}

func (dat *TokenDat) SetToken(token, userid string, auth int64) error {
	tokencache := cache.TokenCacheInfo{
		Token:  token,
		UserID: userid,
		Auth:   auth,
	}
	if err := dat.cache.SetToken(&tokencache); err != nil {
		return err
	}
	return nil
}

func (dat *TokenDat) GetToken(token string) (string, int64, error) {
	tokencache := cache.TokenCacheInfo{
		Token: token,
	}
	if err := dat.cache.GetToken(&tokencache); err != nil {
		return "", 0, err
	}
	return tokencache.UserID, tokencache.Auth, nil
}

func (dat *TokenDat) DelToken(token string) error {
	tokencache := cache.TokenCacheInfo{
		Token: token,
	}
	if err := dat.cache.DelToken(&tokencache); err != nil {
		return err
	}
	return nil
}
