package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var all_headers = map[string][]string(r.Header)
		all_headers["Host"] = []string{r.Host}
		jstr, err := json.MarshalIndent(all_headers, "", " ")
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
		} else {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, string(jstr), "\n")
		}
	})
	http.ListenAndServe(":8000", nil)
}
