package main

import (
	"fmt"
	"net/http"
)


func homeHandler( w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-type", "text/html")
	fmt.Fprint(w, "<h1> Hello go </h1>")
}

func contactHandler( w http.ResponseWriter ,  r *http.Request){
    fmt.Fprint(w, "<h1>contact us</h1>")  
}
 

func main() { 
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact" , contactHandler)
	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", nil)
}