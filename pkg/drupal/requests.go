//
// pkg/drupal/requests.go
//
// 💧 auto_linux
//
// gionnec@gmail.com
// iXOM@prosversusjoes.net
//

package drupal

import (
	"net/url"
)

// Send initial request to get form_build_id
func (i *Installer) getFormBuildID() (string, error) {
	err := i.GetFormBuildID(i.url)
	if err != nil {
		return "", err
	}
	return i.formBuildID, err
}

// Send Select Language request
func (i *Installer) getSelectLanguage(formBuildID string) (string, error) {
	var err error
	data := url.Values{}
	data.Set("langcode", "en")
	data.Set("form_build_id", formBuildID)
	data.Set("form_id", "install_select_language_form")
	data.Set("op", "Save and continue")
	drupalURL := i.url + "/core/install.php"
	err = i.PostFormBuildID(drupalURL, data)
	if err != nil {
		return "", err
	}
	return i.formBuildID, err
}

// Send install profile request
func (i *Installer) getInstallProfile(formBuildID string) (string, error) {
	var err error
	data := url.Values{}
	data.Set("profile", "standard")
	data.Set("form_build_id", formBuildID)
	data.Set("form_id", "install_select_profile_form")
	data.Set("op", "Save and continue")
	url := i.url + "/core/install.php?rewrite=ok&langcode=en"
	err = i.PostFormBuildID(url, data)
	if err != nil {
		return "", err
	}
	return i.formBuildID, err
}

// Send Configure DB request
func (i *Installer) getConfigureDb(formBuildID string) error {
	data := url.Values{}
	data.Set("mysql[database]", i.dbName)
	data.Set("mysql[username]", i.dbUser)
	data.Set("mysql[password]", i.dbPass)
	data.Set("mysql[host]", i.dbHost)
	data.Set("mysql[port]", i.dbPort)
	data.Set("mysql[prefix]", "")
	data.Set("form_build_id", formBuildID)
	data.Set("form_id", "install_settings_form")
	data.Set("op", "Save and continue")
	drupalURL := i.url + "/core/install.php?rewrite=ok&langcode=en&profile=standard"
	return i.PostForm(drupalURL, data)
}

// Send start install request
func (i *Installer) getStartInstall() error {
	drupalURL := i.url + "/core/install.php?rewrite=ok&langcode=en&profile=standard&id=1&op=start"
	return i.Get(drupalURL)
}

// Poll install progress
func (i *Installer) getPollInstallProgress() error {
	var err error
	var progress string
	drupalURL := i.url + "/core/install.php?rewrite=ok&langcode=en&profile=standard&id=1&op=do_nojs&op=do&_format=json"
	for {
		progress, err = i.PostProgress(drupalURL)
		if err != nil {
			return err
		}

		if progress == "100" {
			break
		}
	}
	return err
}

// Send install finished request
func (i *Installer) getFinishInstall() (string, error) {
	var err error
	drupalURL := i.url + "/core/install.php?rewrite=ok&langcode=en&profile=standard&id=1&op=do_nojs&op=finished"
	err = i.GetFormBuildID(drupalURL)
	if err != nil {
		return "", err
	}
	return i.formBuildID, err
}

// Send final request
func (i *Installer) getFinal(formBuildID string) error {
	data := url.Values{}
	data.Set("site_name", i.siteName)
	data.Set("site_mail", i.siteMail)
	data.Set("account[name]", i.admin)
	data.Set("account[pass][pass1]", i.adminPass)
	data.Set("account[pass][pass2]", i.adminPass)
	data.Set("account[mail]", i.siteMail)
	data.Set("site_default_country", "")
	data.Set("date_default_timezone", "America/New_York")
	data.Set("enable_update_status_emails", "1")
	data.Set("form_id", "install_configure_form")
	data.Set("form_build_id", formBuildID)
	data.Set("op", "Save and continue")
	drupalURL := i.url + "/core/install.php?rewrite=ok&langcode=en&profile=standard"
	return i.PostForm(drupalURL, data)
}
