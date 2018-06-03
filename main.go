package main

import "github.com/Racherom/grap/Ridoku"

func main() {
	b, err := Ridoku.NewBoard(Ridoku.Level3)
	if err != nil {
		panic(err)
	}
	b.Set(0, 0, 76)
	b.Set(0, 7, 84)
	b.Set(1, 0, 74)
	b.Set(1, 6, 55)
	b.Set(2, 2, 69)
	b.Set(2, 3, 64)
	b.Set(2, 5, 56)
	b.Set(3, 8, 88)
	b.Set(4, 0, 8)
	b.Set(6, 3, 1)
	b.Set(7, 5, 19)
	b.Set(7, 8, 92)
	b.Set(8, 4, 47)
	b.Set(8, 4, 47)
	b.Set(9, 1, 35)
	b.Successive(0, 4, 1, 4)
	b.Successive(1, 2, 1)
	b.Successive(4, 4, 2)
	b.Successive(5, 6, 3)
	b.Successive(6, 0, 1)
	b.Successive(6, 2, 2, 5)
	b.Successive(8, 2, 1, 3)
	b.Successive(9, 6, 1)
	b.Successive(10, 2, 1)
	//b.SetCorrecct()
	b.Print()
	b.Slove().Print()
}
