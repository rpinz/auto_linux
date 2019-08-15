package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

func findBuildFormID(resp http.Response) string {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("Error reading resp.Body!")
	}
	bodyStr := string(body)
	formIndexStart := strings.Index(bodyStr, "form_build_id\" value=\"") + 22
	formIndexEnd := strings.Index(bodyStr[formIndexStart:], "\"") + formIndexStart
	formBuildID := bodyStr[formIndexStart:formIndexEnd]
	return formBuildID
}

func getInstallProgress(resp http.Response) string {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("Error reading resp.Body!")
	}
	type Property struct {
		Name, Text string
	}
	bodyStr := string(body)
	jsonIndexStart := strings.Index(bodyStr, "percentage") + 13
	jsonIndexEnd := strings.Index(bodyStr[jsonIndexStart:], "\"") + jsonIndexStart
	percentage := bodyStr[jsonIndexStart:jsonIndexEnd]
	return percentage
}

func main() {
	var drupalHost = flag.String("drupal_host", "http://127.0.0.1", "Drupal host address.")
	var dbUser = flag.String("dbuser", "drupal", "Drupal database user.")
	var dbName = flag.String("dbname", "drupal", "Drupal database user.")
	var dbPass = flag.String("dbpass", "", "Drupal database user.")
	var dbHost = flag.String("dbhost", "localhost", "Drupal database host.")
	var dbPort = flag.String("dbport", "3306", "Drupal database port.")
	var admin = flag.String("admin", "admin", "Drupal Admin username.")
	var siteMail = flag.String("site_mail", "drupal@example.com", "Drupal site email.")
	var siteName = flag.String("site_name", "drupal", "Drupal site name.")
	var accountPass = flag.String("account_pass", "password", "Drupal admin account password.")
	flag.Parse()
	jar, _ := cookiejar.New(nil)
	drupalClient := http.Client{
		Jar: jar,
	}

	// Send initial request to get form_build_id
	log.Println("Send initial request to get form_build_id")
	resp, err := drupalClient.Get(*drupalHost)
	if err != nil {
		panic("Drupal Home request failed!")
	}
	formBuildID := findBuildFormID(*resp)
	log.Println(formBuildID)
	resp.Body.Close()

	// Send Select Language request
	log.Println("Send Select Language request")
	data := url.Values{}
	data.Set("langcode", "en")
	data.Set("form_build_id", formBuildID)
	data.Set("form_id", "install_select_language_form")
	data.Set("op", "Save and continue")
	drupalURL := *drupalHost + "/core/install.php"
	resp, err = drupalClient.PostForm(drupalURL, data)
	if err != nil {
		panic("Drupal Language request failed!")
	}
	log.Println(resp.StatusCode)
	formBuildID = findBuildFormID(*resp)
	log.Println(formBuildID)
	resp.Body.Close()

	// Send install profile request
	log.Println("Send install profile request")
	data = url.Values{}
	data.Set("profile", "standard")
	data.Set("form_build_id", formBuildID)
	data.Set("form_id", "install_select_profile_form")
	data.Set("op", "Save and continue")
	drupalURL = *drupalHost + "/core/install.php?rewrite=ok&langcode=en"
	resp, err = drupalClient.PostForm(drupalURL, data)
	if err != nil {
		panic("Drupal install profile request failed!")
	}
	log.Println(resp.StatusCode)
	formBuildID = findBuildFormID(*resp)
	log.Println(formBuildID)
	resp.Body.Close()

	// Send Configure DB request
	log.Println("Send Configure DB request")
	data = url.Values{}
	data.Set("mysql[database]", *dbName)
	data.Set("mysql[username]", *dbUser)
	data.Set("mysql[password]", *dbPass)
	data.Set("mysql[host]", *dbHost)
	data.Set("mysql[port]", *dbPort)
	data.Set("mysql[prefix]", "")
	data.Set("form_build_id", formBuildID)
	data.Set("form_id", "install_settings_form")
	data.Set("op", "Save and continue")
	drupalURL = *drupalHost + "/core/install.php?rewrite=ok&langcode=en&profile=standard"
	resp, err = drupalClient.PostForm(drupalURL, data)
	if err != nil {
		panic("Drupal Configure DB request failed!")
	}
	log.Println(resp.StatusCode)
	resp.Body.Close()

	// Send start install request
	log.Println("Send start install request")
	drupalURL = *drupalHost + "/core/install.php?rewrite=ok&langcode=en&profile=standard&id=1&op=start"
	resp, err = drupalClient.Get(drupalURL)
	if err != nil {
		panic("Drupal Language request failed!")
	}
	log.Println(resp.StatusCode)
	resp.Body.Close()

	// Poll install progress
	log.Println("Polling install progress")
	drupalURL = *drupalHost + "/core/install.php?rewrite=ok&langcode=en&profile=standard&id=1&op=do_nojs&op=do&_format=json"
	for {
		var buf io.Reader
		var progress string
		resp, err = drupalClient.Post(drupalURL, "application/json", buf)
		if err != nil {
			panic("Drupal polling request failed!")
		}
		progress = getInstallProgress(*resp)
		log.Printf("Progress: %s%%\n", progress)
		if progress == "100" {
			break
		}
	}

	// Send start install request
	log.Println("Send install finished request")
	drupalURL = *drupalHost + "/core/install.php?rewrite=ok&langcode=en&profile=standard&id=1&op=do_nojs&op=finished"
	resp, err = drupalClient.Get(drupalURL)
	if err != nil {
		panic("Drupal install finished request failed!")
	}
	log.Println(resp.StatusCode)
	formBuildID = findBuildFormID(*resp)
	log.Println(formBuildID)
	resp.Body.Close()

	// Send final request
	log.Println("Send final request")
	data = url.Values{}
	data.Set("site_name", *siteName)
	data.Set("site_mail", *siteMail)
	data.Set("account[name]", *admin)
	data.Set("account[pass][pass1]", *accountPass)
	data.Set("account[pass][pass2]", *accountPass)
	data.Set("account[mail]", *siteMail)
	data.Set("site_default_country", "")
	data.Set("date_default_timezone", "America/New_York")
	data.Set("enable_update_status_emails", "1")
	data.Set("form_id", "install_configure_form")
	data.Set("form_build_id", formBuildID)
	data.Set("op", "Save and continue")
	drupalURL = *drupalHost + "/core/install.php?rewrite=ok&langcode=en&profile=standard"
	resp, err = drupalClient.PostForm(drupalURL, data)
	if err != nil {
		panic("Drupal install profile request failed!")
	}
	log.Println(resp.StatusCode)
	resp.Body.Close()

	log.Println("Done!! :)")
}
