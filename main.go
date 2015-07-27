package main

import "net/http"

const redirect_to = "http://iwata.hidetake.org"

func init () {
  http.HandleFunc("/", handler)
}

func handler (w http.ResponseWriter, r *http.Request) {
  http.Redirect(w, r, redirect_to, http.StatusMovedPermanently)
}
