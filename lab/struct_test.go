package lab

type T struct{ a int }

func (t T) Mv() {
	println(a)
}

func (t *T) Mp() {
	println(a)
}

func (s *TrySuite) TestF() {
	t := T{1}
	T.Mv(t)
	f := T.Mv
	f(t)
}
