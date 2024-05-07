//
// cmd/orangehrm/main.go
//
// 🍊 auto_linux
//
// gionnec@gmail.com
// iXOM@prosversusjoes.net
//

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rpinz/auto_linux/pkg/orangehrm"
)

func usage() int {
	fmt.Println("🍊 Usage:")
	fmt.Printf("%s:\n", os.Args[0])
	fmt.Printf("  --url=\"http://localhost\"  OrangeHRM host URL.\n")
	fmt.Printf("  --dbuser=\"ohrmuser\"       OrangeHRM database user.\n")
	fmt.Printf("  --dbname=\"ohrm\"           OrangeHRM database name.\n")
	fmt.Printf("  --dbhost=\"localhost\"      OrangeHRM database host.\n")
	fmt.Printf("  --dbport=\"3306\"           OrangeHRM database port.\n")
	fmt.Printf("  --admin=\"admin\"           OrangeHRM admin user.\n")
	fmt.Printf("  --admin_pass=\"ubuntu\"     OrangeHRM admin account password.\n")
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
