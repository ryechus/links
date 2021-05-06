package main

import (
	"fmt"
	"strings"

	anchor "github.com/ryechus/link"
)

var exampleHtml = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <a href="/other-page">A link to 
	
	
	<p>me</p>
	another page</a>
    <a href="/other-page">A link to a third page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := anchor.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
