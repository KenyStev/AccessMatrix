package MatrixManager

import (
  "fmt"
  "strings"
  //"strconv"
  "../Domain"
)
type MatrixManager struct{
  objects []string
  domains []Domain.Domain
  objectCounter int
  domainCounter int
  tableSize int
  currentDomain *Domain.Domain
}

func (matrix *MatrixManager) findDomain(name string) *Domain.Domain{
  for i:=0; i< matrix.domainCounter;i++{
    //fmt.Printf("%s y %s",matrix.domains[i].name )
    if (strings.Compare(matrix.domains[i].Name,name) ==0){
      return &matrix.domains[i]
    }
  }
  return nil
}

func (matrix *MatrixManager) Verify(privilege, object string)  {
  defer func() {
        if r := recover(); r != nil {
            fmt.Println(r)
        }
	}()
  objectIndex:= matrix.findObject(object)
  if(objectIndex<0){
    thrownPanic("Object "+object+" not found.")
  }
  if(matrix.currentDomain.Has(privilege, objectIndex) >=0){
    fmt.Println("Access right "+privilege+" available.")
  }else{
    fmt.Println("Unvalid access right.")
  }
}
func (matrix *MatrixManager) findObject(name string) int{
  for i:=0; i< matrix.objectCounter;i++{
    //fmt.Printf("%s y %s ",matrix.objects[i],name)
    if (strings.Compare(matrix.objects[i],name)==0){
      return i
    }
  }
  return -1
}
func (matrix *MatrixManager) indexOfDomain(name string) int{
  for i:=0; i< matrix.domainCounter;i++{
    //fmt.Printf("%s y %s ",matrix.objects[i],name)
    if (strings.Compare(matrix.domains[i].Name,name)==0){
      return i
    }
  }
  return -1
}

func thrownPanic(message string){
  panic(fmt.Sprintf(message))
}
func (matrix *MatrixManager)setAccess(domain,privilege, object string) bool{
    d := matrix.findDomain(domain)
    objectIndex:= matrix.findObject(object)
    if(d==nil){
      thrownPanic("domain "+domain+" not found")
    }
    if(objectIndex < 0){
      //fmt.Println("aqui2")
      thrownPanic("object "+object+" not found")
    }
    return d.AddPrivilege(privilege,objectIndex)
}
func (matrix *MatrixManager) Reverse(old **MatrixManager) {

  matrix.currentDomain.UnsetSwitchablePrivileges((*old).currentDomain.Privileges)
  name:= (*old).currentDomain.Name
  *old = matrix
  (*old).setCurrentDomain((*old).findDomain(name))
}

func (matrix *MatrixManager) MakeObject(object string, isDomain bool){
  objectIndex:= matrix.findObject(object)
  if(objectIndex>=0){
    //fmt.Println("aqui")
    thrownPanic("object "+object+" already exist")
  }
  if(matrix.objectCounter < matrix.tableSize){
    objectIndex =  matrix.objectCounter
    matrix.objectCounter++
    matrix.objects[objectIndex] = object
    if(!isDomain){
      matrix.currentDomain.AddPrivilege("owner", objectIndex)
      fmt.Println("Object "+object+" created")
    }else{
      matrix.currentDomain.AddPrivilege("control", objectIndex)
    }
    //fmt.Println("object counter dentro "+strconv.Itoa(matrix.objectCounter))
  }
}

func (matrix *MatrixManager)InitTable(){
  defer func() {
        if r := recover(); r != nil {
            fmt.Println(r)
        }
	}()
  matrix.tableSize = 10;
  matrix.objects = make([]string,10)
  matrix.objectCounter = 0
  matrix.domains = make([]Domain.Domain, 10)
  for i:= 0; i<10;i++{
    matrix.domains[i].Make()
  }
  matrix.domains[0].Name ="admin"
  matrix.domainCounter=1
  matrix.setCurrentDomain(&matrix.domains[0])

  matrix.MakeObject("domains",false)
  //fmt.Println("object counter "+strconv.Itoa(matrix.objectCounter))
  matrix.setAccess("admin","make","domains")
  matrix.setAccess("admin","delete","domains")

  matrix.MakeDominio("kenystev")
  matrix.MakeDominio("lisaula")
  matrix.MakeDominio("hola")
  matrix.setAccess("kenystev","make*","domains")
  matrix.setAccess("admin","switch","kenystev")
  matrix.setAccess("kenystev","switch","lisaula")
  matrix.setAccess("lisaula","switch","hola")
  // matrix.setAccess("admin","owner","domains")
  //matrix.PrintTable()
  //matrix.GetCurrentDomain()
//fmt.Println(matrix.stack.Top())
//fmt.Println(matrix.stack.Pop(), matrix.stack.Pop())
}

func (matrix *MatrixManager) Make(objects []string, domains []Domain.Domain, objectCounter, domainCounter, tableSize int)  {
  matrix.objects = make([]string, tableSize)
  copy(matrix.objects,objects)
  matrix.domains = make([]Domain.Domain, tableSize)
  for i:= 0; i<tableSize;i++{
    matrix.domains[i].Make()
  }
  for i:= 0; i<tableSize;i++{
    matrix.domains[i].Name = domains[i].Name
    copy(matrix.domains[i].Privileges, domains[i].Privileges)
  }
  matrix.objectCounter = objectCounter
  matrix.domainCounter = domainCounter
  matrix.tableSize = tableSize
}
func (matrix *MatrixManager)setCurrentDomain(domain *Domain.Domain) {
  matrix.currentDomain = domain
}

func (matrix *MatrixManager)SetAccessRight(domain,privilege, object string) {
  defer func() {
        if r := recover(); r != nil {
            fmt.Println(r)
        }
	}()
  objectIndex := matrix.findObject(object)
  if(objectIndex <0){
    thrownPanic("Object "+object+" not found")
  }
  domainIndex := matrix.findObject(domain)
  if(matrix.currentDomain.Has("owner", objectIndex)>=0 || strings.Compare(matrix.currentDomain.Name,"admin")==0 ||matrix.currentDomain.Has("control", domainIndex)>=0){
    if (matrix.setAccess(domain, privilege, object)){
      fmt.Println("Access right "+privilege+" set to "+domain+" in object "+object)
    }
  }else{
    thrownPanic("Unauthorized domamin can't set access right in "+object)
  }
}
func (matrix *MatrixManager) RmDomain(name string)  {
  defer func() {
        if r := recover(); r != nil {
            fmt.Println(r)
        }
  }()
  if(strings.Compare(matrix.currentDomain.Name,"admin")==0){
      domainIndex := matrix.indexOfDomain(name)
      if(domainIndex<0){
        thrownPanic("Domain "+name+" not found.")
      }
      matrix.RmObject(name)
      matrix.domains = append(matrix.domains[:domainIndex],matrix.domains[domainIndex+1:]...);
      matrix.appendDomain();
      matrix.domainCounter--;
      fmt.Println("Domain "+name+" removed")
  }else{
    thrownPanic("Unauthorized command")
  }
}

func (matrix *MatrixManager) appendDomain()  {
  domain :=  new(Domain.Domain)
  domain.Make()
  matrix.domains = append(matrix.domains, *domain)

}

func(matrix *MatrixManager) RmObjectAfterValidation(name string)  {
  defer func() {
        if r := recover(); r != nil {
            fmt.Println(r)
        }
  }()
  objectIndex := matrix.findObject(name)
  if(objectIndex <0){
    thrownPanic("Object "+name+" not found")
  }
  if(matrix.currentDomain.Has("owner", objectIndex)>=0){
    thrownPanic("Unauthorized domain can't remove object")
  }
  domain := matrix.findDomain(name)
  if(domain != nil){
    thrownPanic("Invalid action, can't remove a domain")
  }
  matrix.RmObject(name)
}
func (matrix *MatrixManager) UnsetAccess(domainName, privilege, object string)  {
  defer func() {
        if r := recover(); r != nil {
            fmt.Println(r)
        }
  }()
  objectIndex:= matrix.findObject(object)
  if(objectIndex<0){
    thrownPanic("Object "+object+" not found.")
  }
  domain:= matrix.findDomain(domainName)
  if (domain == nil){
    thrownPanic("Domain "+domainName+" not found.")
  }

  domainIndex := matrix.findObject(domain.Name)

  if(matrix.currentDomain.Has("owner", objectIndex) >=0 || strings.Compare(matrix.currentDomain.Name, "admin")==0 ||matrix.currentDomain.Has("control", domainIndex)>=0){
    domain.Unset(privilege,objectIndex)
    fmt.Println("Access right "+privilege+" was unset in "+domainName+" on "+object)
  }else{
    thrownPanic("Unauthorized domain can't unset access right")
  }
}

func (matrix * MatrixManager) Switch(name string, old MatrixManager)  *MatrixManager{
  defer func() {
        if r := recover(); r != nil {
            fmt.Println(r)
        }
  }()
  domain := matrix.findDomain(name)
  if(domain == nil){
    thrownPanic("Domain "+name+" not found")
  }
  domainIndex := matrix.findObject(name)
  if(matrix.currentDomain.Has("switch", domainIndex)>=0){
    var nuevaMatrix *MatrixManager
    nuevaMatrix = new(MatrixManager)
    nuevaMatrix.Make(matrix.objects, matrix.domains, matrix.objectCounter, matrix.domainCounter, matrix.tableSize)
    nuevaMatrix.setCurrentDomain(nuevaMatrix.findDomain(domain.Name))

    nuevoDomain :=nuevaMatrix.findDomain(matrix.currentDomain.Name)
    if strings.Compare(nuevoDomain.Name,"admin")!=0 {
      nuevoDomain.UnsetSwitchablePrivileges(old.currentDomain.Privileges)
    }
    nuevaMatrix.SetSwitchablePrivileges(*matrix.currentDomain)
    fmt.Println("Switch completed")
    return nuevaMatrix
  }
  fmt.Println("Unauthorized domain "+matrix.currentDomain.Name+" can't switch to "+name)
  return nil
}

func (matrix *MatrixManager)SetSwitchablePrivileges(domain Domain.Domain){
  matrix.currentDomain.SetSwitchablePrivileges(domain.Privileges)
}
func (matrix *MatrixManager) RmObject(name string)  {
  defer func() {
        if r := recover(); r != nil {
            fmt.Println(r)
        }
  }()
  objectIndex := matrix.findObject(name)
  if(objectIndex <0){
    thrownPanic("Object "+name+" not found")
  }
  matrix.objects = append(matrix.objects[:objectIndex],matrix.objects[objectIndex+1:]...);
  for i:= 0; i<matrix.tableSize; i++{
    privileges:= matrix.domains[i].Privileges
    privileges = append(privileges[:objectIndex],privileges[objectIndex+1:]...);
    matrix.domains[i].Privileges = privileges
  }
  matrix.objectCounter--;
  matrix.appendObjects()
  fmt.Println("Object "+name+" removed.")

}

func (matrix *MatrixManager) appendObjects() {
  matrix.objects = append(matrix.objects, "")
  for i:= 0; i<matrix.tableSize; i++{
    privileges:= matrix.domains[i].Privileges
    privileges = append(privileges,"");
    matrix.domains[i].Privileges = privileges
  }
}

func (matrix *MatrixManager) SetDomain(name string)  {
  defer func() {
        if r := recover(); r != nil {
            fmt.Println(r)
        }
  }()
  domain := matrix.findDomain(name)
  if(domain == nil){
    thrownPanic("Domain "+name+" not found")
  }
  matrix.setCurrentDomain(domain);
  fmt.Println("Domain "+name+" set")
}

func (matrix *MatrixManager)MakeDominio(name string){
  defer func() {
        if r := recover(); r != nil {
            fmt.Println(r)
        }
	}()
  if len(name)==0 {
    thrownPanic("Domain need a name")
  }
  objectIndex := matrix.findObject("domains")
  if(objectIndex <0){
    thrownPanic("Object domains not found")
  }
  if(matrix.currentDomain.Has("make", objectIndex)>=0){
    if(matrix.domainCounter < matrix.tableSize){
      index := matrix.domainCounter;
      matrix.domainCounter++;
      matrix.domains[index].Name = name;
      matrix.MakeObject(name, true)
      fmt.Println("Domain "+name+" created")
    }
  }else{
    thrownPanic("Unauthorized command")
  }
}
func (matrix *MatrixManager)GetCurrentDomain() Domain.Domain {
  return *matrix.currentDomain
}

func (matrix *MatrixManager) PrintCurrentDomain()  {
  format := "|%-20s"
  fmt.Printf(format,"Domain/Object")
  for i:=0; i<matrix.objectCounter;i++{
    fmt.Printf(format,matrix.objects[i])
  }
  fmt.Println("|")

  fmt.Printf(format,matrix.currentDomain.Name)
  for j:=0; j<matrix.objectCounter;j++{
    fmt.Printf(format,matrix.currentDomain.Privileges[j])
  }
  fmt.Println("|")
}

func (matrix *MatrixManager)PrintTable(){
  format := "|%-20s"
  fmt.Printf(format,"Domain/Object")
  for i:=0; i<matrix.objectCounter;i++{
    fmt.Printf(format,matrix.objects[i])
  }
  fmt.Println("|")
  //fmt.Println("Dominios")
  for i:=0; i<matrix.domainCounter;i++{
    fmt.Printf(format,matrix.domains[i].Name)
    for j:=0; j<matrix.objectCounter;j++{
      fmt.Printf(format,matrix.domains[i].Privileges[j])
    }
    fmt.Println("|")
  }
}
