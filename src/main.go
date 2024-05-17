/*

" UNIX is very simple, it just needs a genius to understand its simplicity. " - Dennis Ritchie 

*/
package main 

import (
  "fmt"
  "os"
  "os/exec"
  "bufio"
  "errors"
  "strings"
  "path/filepath"
 
  "github.com/fatih/color"
)

func execInput(input string) error{

  // Removing the newline character at the end of line 
  input = strings.TrimSuffix(input , "\n")

  // Split the input to separate the command and the arguments 
  args := strings.Split(input, " ")

  switch args[0]{
    case "cd":
      
      // 'cd' with no home dir is not supported
      if len(args) < 2{
        return errors.New("path required")
      }
      
      // Change the directory and return the error
      return os.Chdir(args[1])
    
    case "exit":
      os.Exit(0)
  }

  // Prepare the command to execute
  cmd := exec.Command(args[0], args[1:]...)

  cmd.Stderr = os.Stderr
  cmd.Stdout = os.Stdout

  // Excecute the command or return the error
  return cmd.Run()
}

func hostName() (string) {
    host, err1 := os.Hostname()
    if err1 != nil {
        return ""
    }

    user, err2 := os.Getwd()
    if err2 != nil {
        return ""
    }

    name := host + ">> " + filepath.Base(user)
    return name
}

func main(){
   
  reader := bufio.NewReader(os.Stdin)
  
  for{

    // This adds color to the output buffer
    color.New(color.FgRed).Printf("%s> ", hostName())

    // Read the keyboard input here. 
    input, err := reader.ReadString('\n')
    if err != nil{
      fmt.Fprintln(os.Stderr , err)
    }

    // Handling the errors in input
    if err = execInput(input); err != nil{
      fmt.Fprintln(os.Stderr , err)
    }

    fmt.Println(" ")
  }
}
