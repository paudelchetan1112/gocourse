package main 
import "fmt"
type userRegistration interface{
	signup()
	login()
	logout()
}
type user struct{
	username string
	email string
	password string 
}
func (u *user) signup(){
fmt.Print("Enter your full Name:");
fmt.Scanf("%s", &u.username)	
fmt.Print("Enter email:")
fmt.Scanf("%s",&u.email)
fmt.Print("Enter password:")
fmt.Scanf("%s",&u.password)
fmt.Printf("signup Successfully!.. with username: %s and email:%s", u.username, u.email)
}
func (u *user) login(){
fmt.Print("Enter email:")
fmt.Scanf("%s",&u.email)
fmt.Print("Enter password:")
fmt.Scanf("%s",&u.password)
fmt.Printf("login successfully, welcome %s", u.email)
}
func (u *user) logout(){
fmt.Printf("logout")
}
func methodRunner(a userRegistration){
	a.signup()
	a.login()
	a.logout()
}
func main(){
user1:=user{
	username: "chetan",
	email: "paudelchetan1112@gmail.com",
	password: "12345",

}
methodRunner(&user1)

}