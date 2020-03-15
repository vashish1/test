package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"test/fileUploading/database"
	"test/fileUploading/helpers"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var cl *mongo.Collection

func ChangeProfileImage(w http.ResponseWriter, r *http.Request) {
	 
	cl,_:=database.Createdb()

	//a function to change a profile image
	vars := mux.Vars(r)
	Id := vars["ID"]
	userId,_:=strconv.Atoi(Id)
	fmt.Println(userId)
	var user database.User

	userData := database.Findfromuserdb(cl,userId)
	//reading the user whose image we want to change from the database
	if !userData {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}

	imageName, err := helpers.FileUpload(r)
	//here we call the function we made to get the image and save it
	if err != nil {
		http.Error(w, "Invalid Data", http.StatusBadRequest)
		return
		//checking whether any error occurred retrieving image
	}
	//if no error we we change the profile image for our user
	ok:=database.Update(cl,imageName,userId)
	if ok{
		user.Id=userId
		user.Img=imageName
	w.Header().Set("Content-Type", "application/json")
	//we then return the new user details to update our user interface
	json.NewEncoder(w).Encode(user)
	}else{
		http.Error(w, "Could not update", http.StatusNotFound)
	}
}
