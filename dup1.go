package main

import(
	"fmt"
	"bufio"
	"os"
)

func main(){
	map := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		map[input.Text()]++
	}
	for k,v := range map{
		if v>1 {
		fmt.Println(k)
	   }
	}
}