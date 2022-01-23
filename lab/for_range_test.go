package lab

import "fmt"

func (s *LabSuite) TestFor() {
	sl := []int32{1, 2, 3}
	for i, it := range sl {
		fmt.Printf("%d it addr: %p ele addr: %p\n", i, &it, &sl[i])
	}
	for i, it := range sl {
		fmt.Println(i, " it addr: ", &it, " ele addr: ", &sl[i])
	}

	// 永动机测试
	fmt.Printf("%p\n", sl)
	for i, it := range sl { // it 是值复制, sl也是值复制？
		fmt.Println(i, " it addr: ", &it, " ele addr: ", &sl[i])
		sl = append(sl, it) // sl发生了扩容
	}
	fmt.Println("slice array: ", sl)
}
