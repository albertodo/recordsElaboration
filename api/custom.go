package main

import "net/http"

func removeIt(id int, ssSlice []item) []item {
	for idx, v := range ssSlice {
		if v.ID == id {
			return append(ssSlice[0:idx], ssSlice[idx+1:]...)
		}
	}
	return ssSlice
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
}
