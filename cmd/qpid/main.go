package main

import "fmt"
import "golang.org/x/exp/io/i2c"

func main() {
	fmt.Println("hello PI!")
	d, err := i2c.Open(&i2c.Devfs{Dev: "/dev/i2c-1"}, 0x4d)
	if err != nil {
		panic(err)
	}
	b := make([]byte, 3)
	err = d.Read(b)
	if err != nil {
		panic(err)
	}
	temp := b[1]
	celcius := temp / 5 + b[2]
	f := celcius*9/5 + 32

	fmt.Println(temp, celcius, f)
}
