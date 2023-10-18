package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

// writeToConsole is a custom middleware
func writeToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(writer, request)
	})
}

// NoSurf generate the csrf token and protection to all post request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   appConfiguration.InProduction, //Now we are using http, so this value is false. in production we will use https then the value will be true.
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad this load the session and save for each request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
