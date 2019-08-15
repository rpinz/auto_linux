package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

func main() {
	var ohrmHost = flag.String("orangehrm_host", "http://127.0.0.1", "Orangehrm host address.")
	var dbUser = flag.String("dbuser", "orangehrm", "Orangehrm database user.")
	var dbName = flag.String("dbname", "orangehrm", "Orangehrm database user.")
	//var dbPass = flag.String("dbpass", "", "Orangehrm database user.")
	var dbHost = flag.String("dbhost", "localhost", "Orangehrm database host.")
	var dbPort = flag.String("dbport", "3306", "Orangehrm database port.")
	var admin = flag.String("admin", "admin", "Orangehrm Admin username.")
	//var siteMail = flag.String("site_mail", "orangehrm@example.com", "Orangehrm site email.")
	//var siteName = flag.String("site_name", "orangehrm", "Orangehrm site name.")
	var accountPass = flag.String("account_pass", "password", "Orangehrm admin account password.")
	flag.Parse()
	jar, _ := cookiejar.New(nil)
	ohrmClient := http.Client{
		Jar: jar,
	}

	// Send initial request to get form_build_id
	log.Println("Send initial request to get form_build_id")
	resp, err := ohrmClient.Get(*ohrmHost)
	if err != nil {
		panic("Orangehrm Home request failed!")
	}
	resp.Body.Close()

	// Accept license
	log.Println("Accept license")
	data := url.Values{}
	data.Set("txtScreen", "1")
	data.Set("actionResponse", "LICENSEOK")
	ohrmURL := *ohrmHost + "/install.php"
	resp, err = ohrmClient.PostForm(ohrmURL, data)
	if err != nil {
		panic("Orangehrm Accept license failed!")
	}
	log.Println(resp.StatusCode)
	resp.Body.Close()

	// Send database info request
	log.Println("Send database info request")
	data = url.Values{}
	data.Set("txtScreen", "2")
	data.Set("actionResponse", "DBINFO")
	data.Set("dbmethods", "")
	data.Set("dbHostName", *dbHost)
	data.Set("dbHostPortModifier", "port")
	data.Set("dbHostPort", *dbPort)
	data.Set("dbName", *dbName)
	data.Set("dbUserName", *dbUser)
	data.Set("dbPassword", "")
	data.Set("chkSameUser", "1")
	data.Set("cMethod", "existing")
	data.Set("dbCreateMethod", "new")
	resp, err = ohrmClient.PostForm(ohrmURL, data)
	if err != nil {
		panic("Orangehrm database info request failed!")
	}
	log.Println(resp.StatusCode)
	resp.Body.Close()

	// System check OK request
	log.Println("Send System check OK request")
	data = url.Values{}
	data.Set("txtScreen", "3")
	data.Set("actionResponse", "SYSCHECKOK")
	resp, err = ohrmClient.PostForm(ohrmURL, data)
	if err != nil {
		panic("Orangehrm System check OK request failed!")
	}
	log.Println(resp.StatusCode)
	resp.Body.Close()

	// Send Admin account creation request
	log.Println("Send Admin account creation request")
	data = url.Values{}
	data.Set("txtScreen", "4")
	data.Set("actionResponse", "DEFUSERINFO")
	data.Set("OHRMAdminUserName", *admin)
	data.Set("OHRMAdminPassword", *accountPass)
	data.Set("OHRMAdminPasswordConfirm", *accountPass)
	resp, err = ohrmClient.PostForm(ohrmURL, data)
	if err != nil {
		panic("Orangehrm Language request failed!")
	}
	log.Println(resp.StatusCode)
	resp.Body.Close()

	// Confirm install request
	log.Println("Send Confirm install request")
	data = url.Values{}
	data.Set("txtScreen", "5")
	data.Set("actionResponse", "CONFIRMED")
	resp, err = ohrmClient.PostForm(ohrmURL, data)
	if err != nil {
		panic("Orangehrm Confirm install request failed!")
	}
	log.Println(resp.StatusCode)
	resp.Body.Close()

	// Send five requests
	log.Println("Send five request for the install")
	for i := 0; i < 4; i++ {
		resp, err = ohrmClient.Get(ohrmURL)
		if err != nil {
			panic("Orangehrm install request failed!")
		}
		resp.Body.Close()
	}
	time.Sleep(5 * time.Second)

	// REGISTER request
	log.Println("Send REGISTER request")
	data = url.Values{}
	data.Set("txtScreen", "6")
	data.Set("actionResponse", "REGISTER")
	resp, err = ohrmClient.PostForm(ohrmURL, data)
	if err != nil {
		panic("Orangehrm REGISTER request failed!")
	}
	log.Println(resp.StatusCode)
	resp.Body.Close()

	// Skip user registration request
	log.Println("Send Skip user registration request")
	data = url.Values{}
	data.Set("txtScreen", "7")
	data.Set("actionResponse", "NOREG")
	data.Set("firstName", "")
	data.Set("userName", "")
	data.Set("company", "")
	data.Set("userEmail", "")
	data.Set("userTp", "")
	data.Set("userComments", "")
	resp, err = ohrmClient.PostForm(ohrmURL, data)
	if err != nil {
		panic("Orangehrm Skip user registration request failed!")
	}
	log.Println(resp.StatusCode)
	resp.Body.Close()

	log.Println("Done!! :)")
}
