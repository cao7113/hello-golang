package lab

import (
	"fmt"
	"unsafe"
)

func (s *LabSuite) TestRune() {
	str := "你好世界你"
	// byte = uint8, rune=int32
	s.Equal("你好", string([]rune(str)[:2]))
	// 1个中文字符=3个字节
	s.Equal("你", string([]byte(str)[:3]))

	fmt.Printf("%b\n", []byte(str))
	fmt.Println([]byte(str))
	fmt.Println([]rune(str))

	fmt.Println(unsafe.Sizeof(str))
}
