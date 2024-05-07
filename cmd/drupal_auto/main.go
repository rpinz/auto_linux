//
// cmd/drupal/main.go
//
// 💧 auto_linux
//
// gionnec@gmail.com
// iXOM@prosversusjoes.net
//

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rpinz/auto_linux/pkg/drupal"
)

func usage() int {
	fmt.Println("💧 Usage:")
	fmt.Printf("%s:\n", os.Args[0])
	fmt.Printf("  --url=\"http://localhost\"  Drupal host URL.\n")
	fmt.Printf("  --dbuser=\"drupal\"         Drupal database user.\n")
	fmt.Printf("  --dbname=\"drupal\"         Drupal database name.\n")
	fmt.Printf("  --dbhost=\"localhost\"      Drupal database host.\n")
	fmt.Printf("  --dbport=\"3306\"           Drupal database port.\n")
	fmt.Printf("  --admin=\"admin\"           Drupal admin user.\n")
	fmt.Printf("  --admin_pass=\"ubuntu\"     Drupal admin password.\n")
	fmt.Printf("  --site_mail=\"no@no.no\"    Drupal site email.\n")
	fmt.Printf("  --site_name=\"name\"        Drupal site name.\n")
	return 1
}

func main() {
	// usage if no args
	if len(os.Args) < 8 {
		os.Exit(usage())
	}
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
	os.Exit(drupal.Run())
}
