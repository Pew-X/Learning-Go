package httpnetworking

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json::"completed"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to my Server")
	fmt.Fprintf(w, "Hello , World from %s\n", r.URL.Path[1:])
}

func SimpleServer() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

const URL = "https://jsonplaceholder.typicode.com/"

func SimpleClient() {
	resp, err := http.Get(URL + "/todos/1")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)

	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		fmt.Println(string(body))

		var item todo

		err = json.Unmarshal(body, &item)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		fmt.Printf("%#v\n", item)

	}
}
