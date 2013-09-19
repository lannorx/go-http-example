package main

import(
    "time"
    "strings"
    http "net/http"
    "fmt"
    "os"
)

const resp = `<html>
    <head>
        <title>Simple Web App</title>
    </head>
    <body>
        <h1>Simple Web App</h1>
        <p>Hello World!</p>
        <p>$stamp</p>
    </body>
</html>`

func handler(w http.ResponseWriter, r *http.Request) {
    stamp := fmt.Sprintf("%s", time.Now())
    data := strings.Replace(resp, "$stamp", stamp, -1)
    w.Write([]byte(data))
}

func main() {
    http.HandleFunc("/", handler)
    err := http.ListenAndServe(":8080", nil)

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}


