//
// pkg/drupal/run.go
//
// 💧 auto_linux
//
// gionnec@gmail.com
// iXOM@prosversusjoes.net
//

package drupal

import (
	"log"
)

// Run the install
func Run() int {
	var err error
	var installer = new(Installer)

	// initialize struct
	err = installer.Initialize()
	log.Println("💧 Initializing.")
	if err != nil {
		log.Printf("🐛Failed to initialize! %s\n", err)
		return 1
	}

	// Send initial request to get form_build_id
	log.Println("💧 Send initial request to get form_build_id")
	formBuildID, err := installer.getFormBuildID()
	if err != nil {
		log.Printf("💧 Home request failed! %s\n", err)
		return 1
	}

	// Send Select Language request
	log.Println("💧 Send Select Language request")
	formBuildID, err = installer.getSelectLanguage(formBuildID)
	if err != nil {
		log.Printf("💧 Language request failed! %s\n", err)
		return 1
	}

	// Send install profile request
	log.Println("💧 Send install profile request")
	formBuildID, err = installer.getInstallProfile(formBuildID)
	if err != nil {
		log.Printf("💧 Install profile request failed! %s\n", err)
		return 1
	}

	// Send Configure DB request
	log.Println("💧 Send Configure DB request")
	err = installer.getConfigureDb(formBuildID)
	if err != nil {
		log.Printf("💧 Configure DB request failed! %s\n", err)
		return 1
	}

	// Send start install request
	log.Println("💧 Send start install request")
	err = installer.getStartInstall()
	if err != nil {
		log.Printf("💧 Start install request failed! %s\n", err)
		return 1
	}

	// Poll install progress
	log.Println("💧 Polling install progress")
	err = installer.getPollInstallProgress()
	if err != nil {
		log.Printf("💧 polling install progress request failed! %s\n", err)
		return 1
	}

	// Send install finished request
	log.Println("💧 Send install finished request")
	formBuildID, err = installer.getFinishInstall()
	if err != nil {
		log.Printf("💧 install finished failed! %s\n", err)
		return 1
	}

	// Send final request
	log.Println("💧 Send final request")
	err = installer.getFinal(formBuildID)
	if err != nil {
		log.Printf("💧 final request failed! %s\n", err)
		return 1
	}

	log.Println("💧 !!Done!! 😊:)")

	return 0
}
