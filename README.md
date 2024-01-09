# Enterprise Automation LMS Check Service Discovery #
### Overview <br>
This application is designed to streamline the process of discovering LMS (Leveraged Management System) check services within the 'Enterprise Automation' organization on GitHub. It follows a naming standard where LMS check services are expected to be named as 'lms-check-<Name>-ms'. The application retrieves a list of these services, extracts relevant check functionality files from the /handlers folder in the root directory, and analyzes individual helper files to compile a structured list of required arguments for each check.

## Functionality <br>
**Service Discovery:**<br>
Searches the 'Enterprise Automation' organization on GitHub for LMS check services following the naming standard 'lms-check-<"ServceName">-ms', It then compiles a list of identified services.

<br>**File Retrieval:**<br>
Retrieves all relevant check functionality files from the /handlers folder in the root directory of each identified service, stores the folder names and performs regex to create a variable for the function name.

<br>**Argument Analysis:**<br>
Looks inside individual helper .go files to extract a list of arguments required to run each check. It then organizes all of the Service Names, File Names, Function Names and Args required for those functions into a structured format.


<br>**Output Structure**<br>
The application outputs a structured JSON representation of the discovered LMS check services, their associated files, and the required arguments in a sorted manner.

Example Respons
```
[
   {
      "lms-check-github-ms": {
         "file_contains.go FileContains": [
            "CHECK_ORG",
            "CHECK_REPO",
            "CHECK_FILE",
            "CHECK_STRING"
         ]
      }
   },
   {
      "lms-check-github-ms": {
         "file_exists.go FileExists": [
            "CHECK_ORG",
            "CHECK_REPO"
         ]
      }
   },
   {
      "lms-check-github-ms": {
         "repo_exists.go RepoExists": [
            "CHECK_ORG",
            "CHECK_REPO"
         ]
      }
   },
   {
      "lms-check-vmware-ms": {
         "disks.go Disks": [
            "CHECK_HOST",
            "CHECK_USER",
            "CHECK_PASSWORD",
            "CHECK_VM"
         ]
      }
   }
]
``````

**Note**
Ensure that you have the necessary permissions to access the 'Enterprise Automation' organization on GitHub and that the required authentication tokens are provided if the repository is private.