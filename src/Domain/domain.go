package Domain
import (
  "strings"
  "fmt"
  // "strconv"
)
type Domain struct{
  Name string
  Privileges []string
}

func (domain *Domain) Make(){
  domain.Privileges = make([]string, 10)
  for i:=0; i<10;i++{
    domain.Privileges[i]=""
  }
}
func (domain *Domain) AddPrivilege(name string,index int) bool{
  domain.Privileges[index]+=name+"|"
  return true
  //fmt.Printf("ADDES %s in %d\n",name, index)

}

func (domain *Domain) Unset(name string, index int) {
  privilegeIndex := domain.Has(name, index)
  if(privilegeIndex<0){
    panic(fmt.Sprintf("Privilege "+name+" not found in "+domain.Name))
  }
  var array []string
  array = strings.Split(domain.Privileges[index],"|")
  array = append(array[:privilegeIndex],array[privilegeIndex+1:]...);
  domain.Privileges[index] = strings.Join(array, "|")
}

func (domain *Domain) Has(n string, index int) (i int) {
  // fmt.Println("Entro en "+n+" con "+strconv.Itoa(index))
  array := strings.Split(domain.Privileges[index],"|")
  for i= 0; i<len(array); i++{
    var p = strings.Trim(array[i], "*")
    if(strings.Compare(p,n)==0){
      return
    }
  }
  i=-1
  return
}

func getAllSwitchablesPrivileges(privileges string) []string {
  array := strings.Split(privileges,"|")
  var returnValue []string
  for i:=0; i<len(array);i++{
    if(strings.HasSuffix(array[i], "*")){
      returnValue=append(returnValue,strings.TrimSuffix(array[i], "*"))
    }
  }
  return returnValue
}

func (domain *Domain) UnsetSwitchablePrivileges(privileges []string)  {
  for i:=0; i<len(privileges); i++{
    array := getAllSwitchablesPrivileges(privileges[i])
    for index := 0; index < len(array); index++ {
      domain.Unset(array[index], i)
    }
  }
}
func (domain *Domain) SetSwitchablePrivileges(privileges []string){
  for i:=0; i<len(privileges); i++{
    array := getAllSwitchablesPrivileges(privileges[i])
    for index := 0; index < len(array); index++ {
      domain.AddPrivilege(array[index], i)
    }
  }
}
