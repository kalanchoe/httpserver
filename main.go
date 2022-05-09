package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		for k, v := range request.Header {
			for _, s := range v {
				writer.Header().Add(k, s)
			}
		}

		writer.Header().Set("VERSION", os.Getenv("VERSION"))

		ip, err := getIP(request)
		if err != nil {
			log.Printf("fail to get ip, err: %v", err)
			return
		}

		log.Printf("ip: %s, status: %d.", ip, 200)
	})

	http.HandleFunc("/healthz", func(writer http.ResponseWriter, request *http.Request) {

	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// getIP returns request real ip.
func getIP(r *http.Request) (string, error) {
	ip := r.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	ip = r.Header.Get("X-Forward-For")
	for _, i := range strings.Split(ip, ",") {
		if net.ParseIP(i) != nil {
			return i, nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	return "", errors.New("no valid ip found")
}
