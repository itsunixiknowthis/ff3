package ff3

import ("testing"; "github.com/itsunixiknowthis/ff1")

func TestA(t *testing.T) {
	z := ff1.T1{}
	if "T1.F" != z.F() { t.Errorf("!") }
}

