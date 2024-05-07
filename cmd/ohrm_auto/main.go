package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rpinz/auto_linux/pkg/orangehrm"
)

func usage() int {
	fmt.Println("Usage:")
	fmt.Printf("%s:\n", os.Args[0])
	fmt.Printf("  --dbuser=\"drupal\"\n")
	fmt.Printf("  --dbname=\"drupal\"\n")
	fmt.Printf("  --dbhost=\"localhost\"\n")
	fmt.Printf("  --dbport=\"3306\"\n")
	fmt.Printf("  --admin=\"admin\"\n")
	fmt.Printf("  --account_pass=\"ubuntu\"\n")
	return 1
}

func main() {
	// usage if no args
	if len(os.Args) < 6 {
		os.Exit(usage())
	}
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
	os.Exit(orangehrm.Run())
}
