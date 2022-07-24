# Go-Persist

This repo is for building a Windows persistence toolkit in order to learn Golang.   
At the moment, the only supported persistence techniques are registry edits under HKEY_LOCAL_MACHINE and HKEY_CURRENT_USER, file creation under the system or user-specific startup folders, and service creation.    
Inspiration taken from [Mandiant's SharPersist](https://github.com/mandiant/SharPersist)   
  
```
Usage of .\go-persist.exe:
  -a string
        Set the arguments for the command
  -c string
        Set the command or binary to be used for persistence (default "C:\\Windows\\System32\\cmd.exe")
  -f string
        Specify the filename to be written (currently only used for startupfolder persistence) (default "update.bat")
  -k string
        Specify the registry key to be edited for persistence, options include:
        - hklmrunonce: RunOnce key under HKEY LOCAL_MACHINE
        - hklmrun: Run key under HKEY LOCAL_MACHINE
        - hkcurunonce: RunOnce key under HKEY CURRENT_USER
        - hkcurun: Run key under HKEY CURRENT_USER
         (default "hklmrunonce")
  -m string
        add or remove the persistence technique (add will overwrite files/registry values if they already exist) (default "add")
  -n string
        Specify the service name to be created/removed for service persistence (default "updater")
  -t string
        Specify a persistence technique to use, valid options are:
        - registry
        - startupfolder
        - service
         (default "registry")
  -u string
        Specify the user for startupfolder persistence, or specify SYSTEM for the system startup folder (default "Current User")
  -v string
        Set the registry value to be set for persistence (default "*a")
```
   
## Examples  
`.\go-persist.exe -t registry -k hklmrun -v regTest -c "cmd.exe" -a "/c echo 'test' > C:\test.txt" -m add`  
`.\go-persist.exe -t registry -k hklmrun -v regTest -m remove`  
`.\go-persist.exe -t startupfolder -c "cmd.exe" -a "/c echo 'test' > C:\test.txt" -m add`  
`.\go-persist.exe -t startupfolder -m remove`  
`.\go-persist.exe -t service -c C:\Windows\System32\cmd.exe -a "/c echo test > C:\Users\Public\test.txt" -m add`    
`.\go-persist.exe -t service -m remove`   