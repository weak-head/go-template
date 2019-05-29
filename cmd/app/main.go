package app

import "fmt"
import "rsc.io/quote"

// HelloModule returns welcome message
func HelloModule() string {
	// return "Hello, Module"
	return quote.Hello()
}

func main() {
	fmt.Printf(HelloModule())
}
