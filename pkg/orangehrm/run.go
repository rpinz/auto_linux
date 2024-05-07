//
// pkg/orangehrm/run.go
//
// 🍊 auto_linux
//
// gionnec@gmail.com
// iXOM@prosversusjoes.net
//

package orangehrm

import (
	"log"
)

// Run the install
func Run() int {
	var err error
	var installer = new(Installer)

	// initialize struct
	err = installer.Initialize()
	log.Println("🍊 Initializing.")
	if err != nil {
		log.Printf("🐛Failed to initialize! %s\n", err)
		return 1
	}

	// Send initial request to get form_build_id
	log.Println("🍊 Send initial request to get form_build_id.")
	err = installer.getFormBuildID()
	if err != nil {
		log.Printf("🍊 Home request failed! %s\n", err)
		return 1
	}

	// Accept license
	log.Println("🍊 Accept license.")
	err = installer.getAcceptLicense()
	if err != nil {
		log.Printf("🍊 Accept license request failed! %s\n", err)
		return 1
	}

	// Send database info request
	log.Println("🍊 Send database info request.")
	err = installer.getDbInfo()
	if err != nil {
		log.Printf("🍊 database info request failed! %s\n", err)
		return 1
	}

	// System check OK request
	log.Println("🍊 Send System check OK request.")
	err = installer.getSystemCheck()
	if err != nil {
		log.Printf("🍊 System check OK request failed! %s\n", err)
		return 1
	}

	// Send Admin account creation request
	log.Println("🍊 Send Admin account creation request.")
	err = installer.getAdminCreate()
	if err != nil {
		log.Printf("🍊 Admin create request failed! %s\n", err)
		return 1
	}

	// Confirm install request
	log.Println("🍊 Send Confirm install request.")
	err = installer.getConfirmInstall()
	if err != nil {
		log.Printf("🍊 Confirm install request failed! %s\n", err)
		return 1
	}

	// Send five requests
	log.Println("🍊 Send five request for the install.")
	err = installer.getFive()
	if err != nil {
		log.Printf("🍊 install request failed! %s\n", err)
		return 1
	}

	// REGISTER request
	log.Println("🍊 Send REGISTER request.")
	err = installer.getRegister()
	if err != nil {
		log.Printf("🍊 REGISTER request failed! %s\n", err)
		return 1
	}

	// Skip user registration request
	log.Println("🍊 Send Skip user registration request.")
	err = installer.getSkipUserRegistration()
	if err != nil {
		log.Printf("🍊 Skip user registration request failed! %s\n", err)
		return 1
	}

	log.Println("🍊 !!Done!! 😊:)")

	return 0
}
