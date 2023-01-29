package main
// curr change <------------
import (
	logic "TestCi/src"
	"fmt"
	"net/http"
	"sync"
)

var one sync.Once

var instance *http.Client

// State Enum in GO
type State string

// State Enum in GO
const (
	CANCELLED = "CANCELLED"
	OK        = "OK"
	ERROR     = "ERROR"
)

func GetHttpClient() *http.Client {

	one.Do(func() {
		var client http.Client
		instance = &client
	})
	return instance
}

func main() {
	// comment
	var myState State
	myState = CANCELLED
	fmt.Println(myState)
	myState = OK
	fmt.Println(myState)

	fmt.Println("Hello CircleIC")
	fmt.Printf("Foo parameter is: %d", 100)
	logic.WhichInput(100)
}
