package main

import "flag"

// https://harveyli.me/go%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0%E4%B9%8B1/

type hero struct {
	name       string
	superpower string
}

//go:noinline
func createSuperMan() hero {
	h := hero{
		name:       "Superman",
		superpower: "X-ray vision",
	}
	println("Superman ", &h)
	return h
}

//阻止编译器将方法强行inline, 这样会让我们解读EA分析变得更简单些
//go:noinline
func createTheFlash() *hero {
	h := hero{
		name:       "The Flash",
		superpower: "Super speed",
	}
	println("The Flash ", &h)
	return &h
}

func createAnotherHero() *hero {
	h := &hero{
		name:       "Mao",
		superpower: "people",
	}
	println("h pointer: ", h, " addr: ", &h)
	return h
}

// go:noinline
func createSomeHeros(n *int) []hero {
	heros := make([]hero, *n, *n)
	for i := 0; i < *n; i++ {
		var h hero
		if i <= (*n / 2) {
			h = hero{
				name:       "Superman",
				superpower: "X-ray vision",
			}
		} else {
			h = hero{
				name:       "The Flash",
				superpower: "Super speed",
			}
		}
		heros = append(heros, h)
	}
	println(&heros)
	return heros
}

//go:generate go run -gcflags "-m -m" escape_analysis.go
func main() {
	h1 := createSuperMan() // 值copy

	// h will escape to heap, after createTheFlash returns its value, no matter if h2 is declared or not
	h2 := createTheFlash() // 指针传递，对应变量被存储在heap上（动态内存上）

	println("Superman ", &h1)
	println("The Flash pointer: ", h2, " addr of pointer: ", &h2)

	h3 := createAnotherHero()
	println("The Mao", h3, " addr of h3: ", &h3)

	// view console: "moved to heap: h"

	// add n heros where n is specified by users at runtime
	n := flag.Int("n", 0, "specify num of heros")
	flag.Parse()
	hs := createSomeHeros(n)
	println("Heros ", *n, &hs)

	/*
		./escape_analysis.go:34:15: make([]hero, *n, *n) escapes to heap
		./escape_analysis.go:34:15:     from make([]hero, *n, *n) (non-constant size) at ./escape_analysis.go:34:15
		这是因为要创建的hero slice的大小并不是一个常量，也就时说在编译时，编译器无法确切了解slice的大小。由于担心slice会过大stack容纳不小，编译器又一次“保守地”将这个变量放到了容量更大的heap上。
	*/
}
