package encrypt

import (
	"testing"
	"unicode"
)

func TestGenSaltDoubleDiff(t *testing.T) {
	salt1, salt2 := GenSalt(), GenSalt()
	if salt1 == salt2 {
		t.Error("TestGenSaltDoubleDiff failed: salt is same")
	}
}

func TestGenPasswordValidOutput(t *testing.T) {
	passwd := "1c2537c456"
	crypasswd := GenEncryptPasswd(passwd, GenSalt())
	if len(crypasswd) != 32 {
		t.Error("TestGenPassword failed: crypt passwd len error")
	}

	for _, c := range crypasswd {
		if !unicode.IsNumber(c) && !unicode.IsLetter(c) {
			t.Error("TestGenPassword failed: crypt passwd has invalid char(s)")
			break
		}
	}
}

func TestGenPasswordDoubleDiff(t *testing.T) {
	passwd := "1c2537c456"
	crypasswd1 := GenEncryptPasswd(passwd, GenSalt())
	crypasswd2 := GenEncryptPasswd(passwd, GenSalt())
	if crypasswd1 == crypasswd2 {
		t.Error("TestGenPasswordDoubleDiff failed: double gens same crypt passwd")
	}
}
