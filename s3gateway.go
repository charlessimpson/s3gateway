package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	svc := s3.New(session.New(nil))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
			Bucket: &os.Args[1],
			Key:    &r.URL.Path,
		})

		url, err := req.Presign(300 * time.Second)
		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, url, 302)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
