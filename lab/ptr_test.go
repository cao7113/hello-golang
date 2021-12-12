package lab

type ByteSlice []byte

func (s *ByteSlice) Write(b []byte) {
	*s = b
}

func (s *TrySuite) TestPointer() {
	var b ByteSlice
	b.Write([]byte{0x00, 0x01})
	s.Equal([]byte{0x00, 0x01}, []byte(b))
}
