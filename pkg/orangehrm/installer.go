//
// pkg/orangehrm/installer.go
//
// 🍊 auto_linux
//
// gionnec@gmail.com
// iXOM@prosversusjoes.net
//

package orangehrm

import (
	"flag"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

// Installer data
type Installer struct {
	jar       *cookiejar.Jar
	client    *http.Client
	url       string
	dbUser    string
	dbName    string
	dbPass    string
	dbHost    string
	dbPort    string
	admin     string
	adminPass string
}

// Initialize struct
func (i *Installer) Initialize() error {
	var err error
	// define command-line arguments
	var url = flag.String("url", "http://127.0.0.1", "OrangeHRM host url.")
	var dbUser = flag.String("dbuser", "ohrm", "OrangeHRM database user.")
	var dbName = flag.String("dbname", "ohrm", "OrangeHRM database name.")
	var dbPass = flag.String("dbpass", "", "OrangeHRM database pass.")
	var dbHost = flag.String("dbhost", "localhost", "OrangeHRM database host.")
	var dbPort = flag.String("dbport", "3306", "OrangeHRM database port.")
	var admin = flag.String("admin", "admin", "OrangeHRM admin user.")
	var adminPass = flag.String("admin_pass", "password", "OrangeHRM admin password.")
	// parse command-line arguments
	flag.Parse()
	// populate struct from command-line arguments
	i.url = *url
	i.dbUser = *dbUser
	i.dbName = *dbName
	i.dbPass = *dbPass
	i.dbHost = *dbHost
	i.dbPort = *dbPort
	i.admin = *admin
	i.adminPass = *adminPass
	// new cookiejar
	i.jar, err = cookiejar.New(nil)
	if err != nil {
		return err
	}
	// new http client
	i.client = &http.Client{Jar: i.jar}
	return err
}

// Get url
func (i *Installer) Get(url string) error {
	var err error
	response, err := i.client.Get(url)
	if err != nil {
		return err
	}
	log.Printf("🔗 %d\n", response.StatusCode)

	return response.Body.Close()
}

// PostForm url
func (i *Installer) PostForm(url string, data url.Values) error {
	var err error
	response, err := i.client.PostForm(url, data)
	if err != nil {
		return err
	}
	log.Printf("🔗 %d\n", response.StatusCode)

	return response.Body.Close()
}
