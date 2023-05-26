package main

import (
	"net/http"
)

func main() {
	// res, err := http.Get("http://www.google.com/robots.txt")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// robots, err := ioutil.ReadAll(res.Body)

	// defer res.Body.Close()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%s", robots)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err.Error())
	}

}
