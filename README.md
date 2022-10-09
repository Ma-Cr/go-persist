# Go-Persist

---     

This repo is for building a Windows persistence toolkit in order to learn Golang.   
At the moment, the only supported persistence techniques are registry edits under HKEY_LOCAL_MACHINE and HKEY_CURRENT_USER, file creation under the system or user-specific startup folders, and service creation.    
Inspiration taken from [Mandiant's SharPersist](https://github.com/mandiant/SharPersist)   
  
```
Go-Persist is a simple Windows persistence toolkit

Usage:
  go-persist [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  registry    Persistence using the run or runonce registry key (T1547.001)
  service     Persistence by creating a Windows service (T1543.003)
  startup     Persistence using the startup folder for either a user or system (T1547.001)

Flags:
  -h, --help   help for go-persist

Use "go-persist [command] --help" for more information about a command.
```
    
---   
## Registry     
```
Persistence using the run or runonce registry key (T1547.001)

Usage:
  go-persist registry [flags]

Flags:
  -c, --command string   Command to execute for persistence (required with the add method)
  -h, --help             help for registry
  -k, --key string       Specify the registry key to be edited for persistence, options include:
                         - hklmrunonce: RunOnce key under HKEY LOCAL_MACHINE
                         - hklmrun: Run key under HKEY LOCAL_MACHINE
                         - hkcurunonce: RunOnce key under HKEY CURRENT_USER
                         - hkcurun: Run key under HKEY CURRENT_USER
                          (default "hklmrunonce")
  -m, --method string    Add or remove the persistence technique (add will overwrite the registry key if it already exists) (required)
  -v, --value string     Set the registry value to be set for persistence (default "*a")
```
### Example   
```
.\go-persist.exe registry -k hklmrun -v regTest -c "cmd.exe /c echo test > C:\test.txt" -m add
.\go-persist.exe registry -k hklmrun -v regTest -m remove
```
   
---
## Service Creation
```
Persistence by creating a Windows service (T1543.003)

Usage:
  go-persist service [flags]

Flags:
  -a, --arguments string   Arguments for the command
  -c, --command string     Command to execute for persistence (required with the add method)
  -h, --help               help for service
  -m, --method string      Add or remove the persistence technique (required)
  -n, --name string        Name of the service to be created (default "updater")
```
### Example
```
.\go-persist.exe service -c C:\Windows\System32\cmd.exe -a "/c echo test > C:\test.txt" -m add
.\go-persist.exe service -m remove
```
   
---
## Startup Folder
```
Persistence using the startup folder for either a user or system (T1547.001)

Usage:
  go-persist startup [flags]

Flags:
  -c, --command string    Command to execute for persistence (required with the add method)
  -f, --filename string   Filename to be written (default "update.bat")
  -h, --help              help for startup
  -m, --method string     Add or remove the persistence technique (add will overwrite the file if it already exists) (required)
  -u, --user string       Specify the user for startupfolder persistence, or specify SYSTEM for the system startup folder
```
### Example
```
.\go-persist.exe startup -c "cmd.exe /c echo test > C:\test.txt" -m add
.\go-persist.exe startup -m remove
```
