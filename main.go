package main

import "fmt"

func main() {
	body := `
	<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="path/broken
			<span>Boot.dev</span>
		</a>
	</body>
	</html>
	`
	a, e := getURLsFromHTML(body, "https://blog.boot.dev")
	if e != nil {
		fmt.Println("e")
	}
	fmt.Println(a)
}