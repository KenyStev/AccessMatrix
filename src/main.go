package main

import (
  "fmt"
  "./MatrixManager"
)

func commandLine(current *MatrixManager.MatrixManager, old **MatrixManager.MatrixManager){
  for{
    var command1, command2,command3,command4,command5 string
    fmt.Printf("[%s@Host]$ ",current.GetCurrentDomain().Name)
    fmt.Scanf("%s %s %s %s %s", &command1, &command2, &command3,&command4,&command5)
    switch(command1){
    case "su":
      current.SetDomain(command2)
    break;
    case "set":
      switch command2 {
      case "access":
        current.SetAccessRight(command3,command4, command5)
        //fmt.Println("set access"+command3+" -> "+command4)
        break;
      default:
        fmt.Println("set access <domain name> <privilege> <object name>")
      }
      break;
    case "get":
      switch command2 {
      case "domain":
        current.PrintCurrentDomain()
        //fmt.Println("get domain"+command3)
        break;
      default:
        fmt.Println("get domain")
      }
      break;
    case "make":
      switch command2 {
      case "domain":
        current.MakeDominio(command3)
        //fmt.Println("make domain"+command3)
        break;
      case "object":
        current.MakeObject(command3,false)
        //fmt.Println("make object"+command3)
        break;
      default:
        fmt.Println("make domain <domain name>")
        fmt.Println("make object <object name>")
      }
      break;
    case "rm":
      switch command2 {
      case "domain":
        current.RmDomain(command3)
        break;
      case "object":
        current.RmObjectAfterValidation(command3)
        break;
      default:
        fmt.Println("rm domain <domain name>")
        fmt.Println("rm object <object name>")
      }
      break;
    case "unset":
      switch command2 {
      case "access":
        current.UnsetAccess(command3, command4, command5)
        break;
      default:
        fmt.Println("unset access <domain name> <privilege> <object name>")
      }
      break;
    case "verify":
      if len(command2)>0 && len(command3)>0 {
        current.Verify(command2, command3)
      }else{
        fmt.Println("verify <privilege name> <object name>")
      }
      break;
    case "switch":
      var nuevaMatrix *MatrixManager.MatrixManager
      nuevaMatrix = current.Switch(command2, *(*old))
      if(nuevaMatrix != nil){
        commandLine(nuevaMatrix, &current)
      }
      current.SetSwitchablePrivileges((*old).GetCurrentDomain())
      break;
    case "exit":
      current.Reverse(old)
      fmt.Println("Sayounara")
      return
      break;
    case "show":
      switch command2{
        case "matrix":
          current.PrintTable();
        break;
        default:
          fmt.Println("show matrix")
      }
      break;
    default:
      fmt.Println("Command not found")
    }
  }
}

func main() {
    fmt.Printf("Access Matrix\n")
    //fmt.Println("You are ADMIN")
    var matrix *MatrixManager.MatrixManager
    matrix = new(MatrixManager.MatrixManager)
    matrix.InitTable();
    commandLine(matrix, &matrix)
}
