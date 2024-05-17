/*

" UNIX is very simple, it just needs a genius to understand its simplicity. " - Dennis Ritchie 

*/

package main 

import (
  "fmt"
  "os"
  "os/exec"
  "bufio"
  // "errors"
  "strings"
)

func execInput(input string) error{

  // Removing the newline character at the end of line 
  input = strings.TrimSuffix(input , "\n")

  // Prepare the command to execute
  cmd := exec.Command(input)

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
  }
}
