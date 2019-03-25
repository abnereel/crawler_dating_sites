package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is ccmouse1@gmail.com ,hello
email1 is abc@def.org
email2 is 408666378@qq.com
email3 is ddd@abc.com.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
	for _, m := range match {
		fmt.Println(m)
		/*for _, n := range m {
			 fmt.Println(n)
		}*/
	}
}
