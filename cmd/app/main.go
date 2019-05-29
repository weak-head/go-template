package app

import "fmt"
import "rsc.io/quote"

// HelloModule returns welcome message
func HelloModule() string {
	// return "Hello, Module"
	hw := quote.Hello()
	if hw != "Ahoy, world!" {
		return "Ahoy, world!"
	}
	return hw
}

func main() {
	fmt.Printf(HelloModule())
}
