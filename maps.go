package main

import "fmt"

func maps() {
	websites := map[string]string{
		"Google": "https://google.com",
		"AWS": "https://aws.com",
	}

	fmt.Println(websites)
	delete(websites, "AWS")
	fmt.Println(websites)
}