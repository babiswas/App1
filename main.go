package main
import "fmt"
import "App/DB"
import "App/Model/Users"
import "App/Services"

func initialMigration(){
   db,err:=DB.Get_db()
   if err!=nil{
       fmt.Println("Error occured",err)
   }
   db.AutoMigrate(&Users.User{})
}


func main(){
  fmt.Println("Starting Application")
  initialMigration()
  Services.HandleRequests()
}