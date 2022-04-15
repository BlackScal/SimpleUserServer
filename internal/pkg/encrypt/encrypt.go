package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"userserver/pkg/utils"
)

var (
	defaultSaltLen int = 8
)

func GenSalt() string {
	return utils.GenRandomString(defaultSaltLen)
}

func GenEncryptPasswd(passwd, salt string) string {
	m5 := md5.New()
	m5.Write([]byte(passwd))
	m5.Write([]byte(salt))
	st := m5.Sum(nil)
	return hex.EncodeToString(st)
}
