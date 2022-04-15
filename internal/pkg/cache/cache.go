package cache

type TokenCacheInfo struct {
	Token  string
	UserID string
	Auth   int64
}

type CacheInterface interface {
	SetPoolSize(poolSize int)
	Init() error
	Destroy()
}

type TokenCacheInterface interface {
	CacheInterface
	GetToken(tokeninfo *TokenCacheInfo) error
	SetToken(tokeninfo *TokenCacheInfo) error
	DelToken(tokeninfo *TokenCacheInfo) error
}
