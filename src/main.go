package main

import (
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
    "os"
    "crypto/tls"
    "fmt"
)

func main() {
    http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

    Host := os.Getenv("WAC_TARGET_HOST")
    Port := os.Getenv("WAC_TARGET_PORT")
    Scheme := os.Getenv("WAC_TARGET_SCHEME")

    origin, _ := url.Parse(fmt.Sprintf("%s://%s:%s", Scheme, Host, Port))

    director := func(req *http.Request) {
        req.Header.Add("Authorization", "Apikey " + os.Getenv("WA_API_KEY"))
        req.URL.Scheme = Scheme
        req.URL.Host = origin.Host
    }

    proxy := &httputil.ReverseProxy{Director: director}

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        proxy.ServeHTTP(w, r)
    })

    log.Fatal(http.ListenAndServe(":" + os.Getenv("WAC_SOURCE_PORT"), nil))
}
