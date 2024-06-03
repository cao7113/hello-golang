package main

func main() {
	defer println("defer in top-nested")
	defer func() {
		defer func() {
			panic("panic again and again")
		}()
		panic("panic again")
	}()
	panic("top panic once")
}
