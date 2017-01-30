package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    // get the absolute directory of this .go file
    root, _ := filepath.Abs(".")

    fmt.Println("Processing path", root)

    // call our function for each item in the tree
    err := filepath.Walk(root, processPath)

    if err != nil {
        fmt.Println(err)
    }

    // SAMPLE OUTPUT:
    //
    // Processing path D:\VinceGo\src\GoLearningSession
    // Directory: D:\VinceGo\src\GoLearningSession
    // Directory: D:\VinceGo\src\GoLearningSession\.idea
    // File: D:\VinceGo\src\GoLearningSession\.idea\codeStyleSettings.xml
    // Directory: D:\VinceGo\src\GoLearningSession\.idea\libraries
    // File: D:\VinceGo\src\GoLearningSession\.idea\libraries\GOPATH__GoLearningSession_.xml
    // File: D:\VinceGo\src\GoLearningSession\.idea\misc.xml
    // File: D:\VinceGo\src\GoLearningSession\.idea\modules.xml
    // File: D:\VinceGo\src\GoLearningSession\.idea\workspace.xml
    // File: D:\VinceGo\src\GoLearningSession\00SetupInsdtuctions.txt
}

func processPath(path string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }

    // make sure we're not in the root directory
    if path != "." {        
        if info.IsDir() {
            // if we're working with a directory
            fmt.Println("Directory:", path)
        } else {
            // if we're working with a file
            fmt.Println("File:", path)
        }
    }

    return nil
}
