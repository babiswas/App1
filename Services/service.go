package Services
import "fmt"
import "net/http"
import "encoding/json"
import "github.com/gorilla/mux"
import "App/Model/Users"
import "App/DB"



func allUsers(w http.ResponseWriter,r *http.Request){
   var users []Users.User
   db,err:=DB.Get_db()
   if err!=nil{
      fmt.Println("Error occured",err)
   }
   db.Find(&users)
   json.NewEncoder(w).Encode(users)
}


func newUser(w http.ResponseWriter,r *http.Request){
   db,err:=DB.Get_db()
   if err!=nil{
     fmt.Println("Error occured",err)
   }
   vars:=mux.Vars(r)
   name:=vars["name"]
   email:=vars["email"]
   db.Create(&Users.User{Name:name,Email:email})
   fmt.Fprintf(w,"New user sucessfully created")
}


func deleteUser(w http.ResponseWriter,r *http.Request){
   db,err:=DB.Get_db()
   if err!=nil{
     fmt.Println("Error occured",err)
   }
   vars:=mux.Vars(r)
   name:=vars["name"]
   var user Users.User
   db.Where("name=?",name).Find(&user)
   db.Delete(&user)
   fmt.Fprintf(w,"Sucessfully deleted user")
}

func updateUser(w http.ResponseWriter,r *http.Request){
   db,err:=DB.Get_db()
   if err!=nil{
      fmt.Println("Error occured",err)
   }
   vars:=mux.Vars(r)
   name:=vars["name"]
   email:=vars["email"]
   var user Users.User
   db.Where("name=?",name).Find(&user)
   user.Email=email
   db.Save(&user)
   fmt.Fprintf(w,"Sucessfully updated user")
}


func HandleRequests(){
   router:=mux.NewRouter().StrictSlash(true)
   router.HandleFunc("/users",allUsers).Methods("GET")
   router.HandleFunc("/user/{name}",deleteUser).Methods("DELETE")
   router.HandleFunc("/user/{name}/{email}",updateUser).Methods("PUT")
   router.HandleFunc("/user/{name}/{email}",newUser).Methods("POST")
   http.ListenAndServe(":8081",router)
}