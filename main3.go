package main

import "fmt"

type a struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func main() {
	//a := map[string]string{"name": "jixingxing", "age": "18"}
	var c a
	c.Name = "jixingxing"
	c.Age = "18"
	for {
		b := make(map[string]string)
		b[c.Name] = c.Age
		fmt.Println(b)
	}

}
