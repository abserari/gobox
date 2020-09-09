package main

import "net/http"

var HomeFolder = "./"

func main(){
	  http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(HomeFolder + "images/"))))
}
