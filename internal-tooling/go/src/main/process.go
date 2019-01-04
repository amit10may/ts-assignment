package main

import ( 
   "fmt"
   "os"
   "os/user"
)

func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }

    // Current User
    fmt.Println("User ID: " + user.Uid )
    //fmt.Println("Username: " + user.Username)
    fmt.Println("Process ID: ", os.Getpid())

    // Get "Real" User under sudo.
    // More Info: https://stackoverflow.com/q/29733575/402585
    //fmt.Println("Real User: " + os.Getenv("SUDO_USER"))
}
