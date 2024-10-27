package main

import (
	"fmt"
	"log"

	"github.com/colinmarc/hdfs/v2"
)

func main() {
	client, err := hdfs.NewClient(hdfs.ClientOptions{
		Addresses: []string{"localhost:8020"},
		User:      "hadoop",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// create dir
	// if err = client.Mkdir("/user", 0777); err != nil {
	// 	log.Fatal(err)
	// }

	files, err := client.ReadDir("/user")
	if err != nil {
		log.Fatalf("Failed to list directory: %v", err)
	}

	for _, file := range files {
		log.Printf("Name: %s, Size: %d bytes, Directory: %v\n", file.Name(), file.Size(), file.IsDir())
		content, err := client.ReadFile(fmt.Sprintf("/user/%s", file.Name()))
		if err != nil {
			log.Println("failed to read file", file.Name(), err)
			continue
		}
		log.Printf("%s\n", string(content))
	}

	// dir, err := client.ReadDir("user")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// for _, f := range dir {
	// 	log.Println(f.Name())
	// }
	//
	// file, err := client.Open("/user/text.txt")
	// if err != nil {
	// 	var pathErr *os.PathError
	// 	if errors.As(err, &pathErr) {
	// 		log.Println("file /user/text.txt is empty", err)
	// 	} else {
	// 		log.Fatal("failed to open file", err)
	// 	}
	// }
	//
	// _ = file
	//
	// file, err = client.Open("/user/text2.txt")
	// if err != nil {
	// 	var pathErr *os.PathError
	// 	if errors.As(err, &pathErr) {
	// 		log.Println("file /user/text2.txt is empty", err)
	// 	} else {
	// 		log.Fatal("failed to open file", err)
	// 	}
	// }
}
