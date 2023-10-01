package main

import (
	"net/http"
)

func Default(data interface{}) {
	http.Get("https://example.com")
}
