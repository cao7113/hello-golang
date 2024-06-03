package lab

func (s *LabSuite) TestSwitch() {
	s.Equal(-1, switchFn(0))
	s.Equal(1, switchFn(1))
	s.Equal(2, switchFn(3))
	s.Equal(0, switchFn(4))
}

func switchFn(n int) int {
	switch n {
	case 0: // nothing do
	case 1:
		return 1
	case 2, 3:
		return 2
	default:
		return 0
	}
	return -1
}
