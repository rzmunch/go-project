package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

/*
	func listFiles(w io.Writer, bucket string) error {
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS",
			"/home/rzmunch/Vict/Google Cloud/Buckets/vic-sistemas-vnetbuckets-dev-b8c9172ce811.json")
		//bucket := "bucket_vic_pf"
		ctx := context.Background()
		client, err := storage.NewClient(ctx)
		if err != nil {
			return fmt.Errorf("storage.NewClient: %w", err)
		}
		defer client.Close()

		ctx, cancel := context.WithTimeout(ctx, time.Second*10)
		defer cancel()

		it := client.Bucket(bucket).Objects(ctx, nil)
		for {
			attrs, err := it.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return fmt.Errorf("Bucket(%q).Objects: %w", bucket, err)
			}
			fmt.Fprintln(w, attrs.Name)
		}
		return nil
	}
*/
func main() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS",
		"sa")
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Printf("Error")
	}
	bucketName := "bucket"
	bkt := client.Bucket(bucketName)

	query := &storage.Query{Prefix: ""}

	var names []string
	it := bkt.Objects(ctx, query)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		names = append(names, attrs.Name)
		fmt.Println(attrs.Name)
	}
}
