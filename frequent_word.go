package main

import (
	"fmt"
	"strings"
)
func frequent(str string) map[string]int {
	sentence:=strings.Fields(str)
	m:=make(map[string]int)

	for _, word := range sentence{
		m[word]=m[word]+1
	}

	return m	

}
func main(){
	sent1 := "qwe asd qwe asd zxc yui qwe"
	for key , value := range frequent(sent1){
		fmt.Println(key,"--",value)
		max:=0
		if value>max{
			max = value
			fmt.Println(key,"word has highest frequency of",max)
		}
}
}