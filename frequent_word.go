package main

import (
	"fmt"
	"strings"
)
func frequent(str string) (string , int) {
	sentence:=strings.Fields(str)
	m:=make(map[string]int)

	for _, word := range sentence{
		m[word]=m[word]+1
	}
	max:=0
    word:=""
	for key, value := range m {
     if value>max{
		max=value
		word=key
	 }
	}
return word, max

}
func main(){
	sent1 := "qwe asd qwe asd zxc yui qwe"
	// for key , value := range frequent(sent1){
	// 	fmt.Println(key,"--",value)
	key, value:=frequent(sent1)
	fmt.Println(key,"word has highest frequency of",value)
}
