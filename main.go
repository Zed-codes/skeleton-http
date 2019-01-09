package main

import (
    "fmt"
    "net/http"
    "strings"
    "log"
)

func checkName(w http.ResponseWriter, r *http.Request) {
	
	r.ParseForm() // parse arguments, we have to call this by ourselves
    //fmt.Println(r.Form)  // print form information in server side
    //fmt.Println("path", r.URL.Path)
    //fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	
	var id string = "unknown"

	// Looping on keys/values to find the value for the key 'name'
    for k, v := range r.Form {
		if k == "name" {
			var name string = strings.Join(v, "")
			fmt.Println("name:", name)
			
			// We look for the id linked to the entered value for name
			// This could evolve into a databse search in future examples
			
			switch name {
			case "amir":
				id = "001"
			case "amin":
				id = "003"
			case "amel":
				id = "005"
			default:
				id = "stranger"
			}

			fmt.Println("id:", id)
		}
        //fmt.Println("key:", k)
        //fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Result: %s", id) // display data to client
}

func main() {
    http.HandleFunc("/", checkName) // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
        log.Fatal("ListenAndServe: ", err)
    } 
}

// Browse and play with value name: http://localhost:9090/?name=amin