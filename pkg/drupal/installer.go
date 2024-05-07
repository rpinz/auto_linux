//
// pkg/drupal/installer.go
//
// 💧 auto_linux
//
// gionnec@gmail.com
// iXOM@prosversusjoes.net
//

package drupal

import (
	"flag"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

// Installer data
type Installer struct {
	jar         *cookiejar.Jar
	client      *http.Client
	formBuildID string
	url         string
	dbUser      string
	dbName      string
	dbPass      string
	dbHost      string
	dbPort      string
	admin       string
	siteMail    string
	siteName    string
	adminPass   string
}

// Initialize struct
func (i *Installer) Initialize() error {
	var err error
	// define command-line arguments
	var url = flag.String("url", "http://127.0.0.1", "Drupal host url.")
	var dbUser = flag.String("dbuser", "drupal", "Drupal database user.")
	var dbName = flag.String("dbname", "drupal", "Drupal database name.")
	var dbPass = flag.String("dbpass", "", "Drupal database pass.")
	var dbHost = flag.String("dbhost", "localhost", "Drupal database host.")
	var dbPort = flag.String("dbport", "3306", "Drupal database port.")
	var admin = flag.String("admin", "admin", "Drupal admin user.")
	var adminPass = flag.String("admin_pass", "password", "Drupal admin password.")
	var siteMail = flag.String("site_mail", "drupal@example.com", "Drupal site email.")
	var siteName = flag.String("site_name", "drupal", "Drupal site name.")
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
	i.siteMail = *siteMail
	i.siteName = *siteName
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

// GetFormBuildID url
func (i *Installer) GetFormBuildID(url string) error {
	var err error
	response, err := i.client.Get(url)
	if err != nil {
		return err
	}
	log.Printf("🔗 %d\n", response.StatusCode)
	i.formBuildID, err = findBuildFormID(response)
	if err != nil {
		return err
	}

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

// PostFormBuildID url
func (i *Installer) PostFormBuildID(url string, data url.Values) error {
	var err error
	response, err := i.client.PostForm(url, data)
	if err != nil {
		return err
	}
	log.Printf("🔗 %d\n", response.StatusCode)
	i.formBuildID, err = findBuildFormID(response)
	if err != nil {
		return err
	}

	return response.Body.Close()
}

// PostProgress url
func (i *Installer) PostProgress(url string) (string, error) {
	var err error
	var buf io.Reader
	var progress string
	response, err := i.client.Post(url, "application/json", buf)
	if err != nil {
		return progress, err
	}
	log.Printf("🔗 %d\n", response.StatusCode)
	progress, err = getInstallProgress(response)
	if err != nil {
		return progress, err
	}

	return progress, response.Body.Close()
}

func findBuildFormID(response *http.Response) (string, error) {
	var err error
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	bodyStr := string(body[:])
	formIndexStart := strings.Index(bodyStr, "form_build_id\" value=\"") + 22
	formIndexEnd := strings.Index(bodyStr[formIndexStart:], "\"") + formIndexStart
	formBuildID := bodyStr[formIndexStart:formIndexEnd]
	log.Println(formBuildID)
	return formBuildID, err
}

func getInstallProgress(response *http.Response) (string, error) {
	var err error
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	bodyStr := string(body[:])
	jsonIndexStart := strings.Index(bodyStr, "percentage") + 13
	jsonIndexEnd := strings.Index(bodyStr[jsonIndexStart:], "\"") + jsonIndexStart
	percentage := bodyStr[jsonIndexStart:jsonIndexEnd]
	log.Printf("💯 Progress: %s%%\n", percentage)
	return percentage, err
}
