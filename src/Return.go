package main

import "fmt"

func main() {

}

func Register() {
	var userName string = "admin"
	var userPwd string = "123"
	var userEmail string = "xxx@oo.com"

	CheckInfo(userName, userPwd, userEmail)

	var isValid, message = CheckInfo2(userName, userPwd, userEmail);
	if (isValid) {
		fmt.Printf(message)
	}
}

func CheckInfo(name string, pwd string, email string) bool {
	return name != "" && pwd != "" && email != ""
}

func CheckInfo3(name string, pwd string, email string) (isValid bool) {
	isValid = name != "" && pwd != "" && email != ""
	return isValid
}

func CheckInfo2(name string, pwd string, email string) (bool,string) {
	var isValid = name != "" && pwd != "" && email != ""
	var message string

	if(isValid){
		message = "Valid Message"
	}else{
		message = "Invalid Message"
	}

	return isValid,message
}