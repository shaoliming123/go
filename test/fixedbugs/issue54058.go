package main

type T[Pad any] struct {
	_ Pad
	X
}

type X string

func (t X) M() { defer println(t) }

func f[Pad any]() {
	f1 := T[[1]Pad].M
	f1(T[[1]Pad]{X: "a"})

}

func main() {
	f[[1]int]()
	//f[[2]int]()
}
