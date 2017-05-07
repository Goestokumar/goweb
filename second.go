package main

import (
	"fmt"
	//"os"
)

func main() {
	name := "Kumar"
	tpl := `
	<!DOCTYPE html>
	<html lang = "en">
	<head>
	<meta charset = "UTF-8">
	<title> "Learning Golang"</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`
	fmt.Println(tpl)

}
