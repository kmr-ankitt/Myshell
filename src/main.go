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

func main(){
   
  reader := bufio.NewReader(os.Stdin)
  
  for{
    
    fmt.Print("> ")

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
