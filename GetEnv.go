/*
	This is my first ever Golang project. I used the following websites to get started and how to use environment strings:
	https://medium.com/rungo/how-to-write-a-simple-go-program-13fd104f3018
	https://yourbasic.org/golang/environment-variables/
	
	I used the hint to run code from the browser. I used https://play.golang.org/ to run my script.
	
	I reach out to a few developer friends but none of them used Golang. Did an online search, took me a while to get hold of it. 
	This script lists the environment variables and I used '=' to delimit key/value pairs. This project is part of the job pre-request
	assignment.
*/
package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {

    for _, m := range os.Environ() {

        pair := strings.SplitN(m, "=", 2)
        fmt.Printf("%s=%s\n", pair[0], pair[1])
    }
}