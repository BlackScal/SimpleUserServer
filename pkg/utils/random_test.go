package utils

import "testing"

func TestGenRandomStringValidOutput(t *testing.T) {
	ranstrlen := 12
	ranstr := GenRandomString(ranstrlen)
	if len(ranstr) != ranstrlen {
		t.Errorf("GenRandomString failed: len=%d, except=%d", len(ranstr), ranstrlen)
	}

	for _, r := range ranstr {
		ok := false
		for _, a := range Alnum {
			if r == a {
				ok = true
				break
			}
		}
		if !ok {
			t.Error("GenRandomString failed: ranstr has invalid chars: ", ranstr)
			break
		}
	}
}

func TestGenRandomStringUnique(t *testing.T) {
	ranstrlen := 12
	ranstr1 := GenRandomString(ranstrlen)
	ranstr2 := GenRandomString(ranstrlen)
	if ranstr1 == ranstr2 {
		t.Error("GenRandomString failed: not unique")
	}
}

func BenchmarkGenRandomStringLen8(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GenRandomString(8)
	}
}
