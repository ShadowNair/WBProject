package main

import "net/http"

func handle(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r, "web/templates/index.html")
}

func main(){
	http.HandleFunc("/", handle)

	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		panic(err)
	}
}
