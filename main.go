// pass the number in URL, and post the number in Corresponding Channel based on even or odd"
package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	//server running statement
	fmt.Println("server running on 4000")

	//creating mux router
	r := mux.NewRouter()

	//calling the func as a go routine
	go ReadFromChannel()

	//hello route
	r.HandleFunc("/", hello)

	//welcome route-we giving some name after the /
	r.HandleFunc("/homepage/{name}", welcome)

	//taking number from url
	r.HandleFunc("/number/{num}", getNum)

	//server running at 4000
	http.ListenAndServe(":4000", r)

}

// hello handler-we hit http://localhost:4000/ it will show a message on browser with hello msg
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

// welcome handler-we hit http://localhost:4000/homepage/{any-name-here}-it will use mux.Vars to read the url data
// and inside params then give response with same name which we passed in url bcz inside params the value is their
// as ey value pair
func welcome(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Fprintln(w, "welcome", params["name"])

}

// create 2 channels outside the func which means global scope otherwise after func call it will destroy
// try to create channel with make otherwise memory wont allocated for each channel
var chOdd = make(chan int)
var chEven = make(chan int)

// getNum handler-take number from url converting to an integer(in url it will be in string format)
// finding that number is odd or even with if else stmt(switch also fine) then passing to corresponding channel
// giving value to channel means writing data into the channel
func getNum(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	num := params["num"]
	//fmt.Println(num)
	convertnum, _ := strconv.Atoi(num)
	fmt.Println(convertnum)

	if convertnum%2 == 0 {
		fmt.Fprintln(w, "number", convertnum, "is even")
		chEven <- convertnum

	} else {
		fmt.Fprintln(w, "number", convertnum, "is odd")

		chOdd <- convertnum
	}

}

// once writing we need to read the value from channel
func ReadFromChannel() {
	for {
		select {
		case <-chOdd:
			fmt.Println("given number is odd")
		case <-chEven:
			fmt.Println("given number is even")

		}
	}
}
