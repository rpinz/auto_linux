//
// pkg/orangehrm/requests.go
//
// 🍊 auto_linux
//
// gionnec@gmail.com
// iXOM@prosversusjoes.net
//

package orangehrm

import (
	"net/url"
	"time"
)

// Send initial request to get form_build_id
func (i *Installer) getFormBuildID() error {
	err := i.Get(i.url)
	if err != nil {
		return err
	}
	return err
}

// Accept license
func (i *Installer) getAcceptLicense() error {
	data := url.Values{}
	data.Set("txtScreen", "1")
	data.Set("actionResponse", "LICENSEOK")
	url := i.url + "/install.php"
	err := i.PostForm(url, data)
	if err != nil {
		return err
	}
	return err
}

// Send database info request
func (i *Installer) getDbInfo() error {
	var err error
	data := url.Values{}
	data.Set("txtScreen", "2")
	data.Set("actionResponse", "DBINFO")
	data.Set("dbmethods", "")
	data.Set("dbHostName", i.dbHost)
	data.Set("dbHostPortModifier", "port")
	data.Set("dbHostPort", i.dbPort)
	data.Set("dbName", i.dbName)
	data.Set("dbUserName", i.dbUser)
	data.Set("dbPassword", "")
	data.Set("chkSameUser", "1")
	data.Set("cMethod", "existing")
	data.Set("dbCreateMethod", "new")
	err = i.PostForm(i.url, data)
	if err != nil {
		return err
	}
	return err
}

// System check OK request
func (i *Installer) getSystemCheck() error {
	var err error
	data := url.Values{}
	data.Set("txtScreen", "3")
	data.Set("actionResponse", "SYSCHECKOK")
	err = i.PostForm(i.url, data)
	if err != nil {
		return err
	}
	return err
}

// Send Admin account creation request
func (i *Installer) getAdminCreate() error {
	var err error
	data := url.Values{}
	data.Set("txtScreen", "4")
	data.Set("actionResponse", "DEFUSERINFO")
	data.Set("OHRMAdminUserName", i.admin)
	data.Set("OHRMAdminPassword", i.adminPass)
	data.Set("OHRMAdminPasswordConfirm", i.adminPass)
	err = i.PostForm(i.url, data)
	if err != nil {
		return err
	}
	return err
}

// Confirm install request
func (i *Installer) getConfirmInstall() error {
	var err error
	data := url.Values{}
	data.Set("txtScreen", "5")
	data.Set("actionResponse", "CONFIRMED")
	err = i.PostForm(i.url, data)
	if err != nil {
		return err
	}
	return err
}

// Send five requests
func (i *Installer) getFive() error {
	var err error
	for j := 0; j < 4; j++ {
		err := i.Get(i.url)
		if err != nil {
			return err
		}
	}
	time.Sleep(5 * time.Second)
	return err
}

// REGISTER request
func (i *Installer) getRegister() error {
	var err error
	data := url.Values{}
	data.Set("txtScreen", "6")
	data.Set("actionResponse", "REGISTER")
	err = i.PostForm(i.url, data)
	if err != nil {
		return err
	}
	return err
}

// Skip user registration request
func (i *Installer) getSkipUserRegistration() error {
	var err error
	data := url.Values{}
	data.Set("txtScreen", "7")
	data.Set("actionResponse", "NOREG")
	data.Set("firstName", "")
	data.Set("userName", "")
	data.Set("company", "")
	data.Set("userEmail", "")
	data.Set("userTp", "")
	data.Set("userComments", "")
	err = i.PostForm(i.url, data)
	if err != nil {
		return err
	}
	return err
}
