package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"test_2/repository"
)

var Phonebook = make(repository.PhoneBook)

//type User struct {
//	Name  string
//	Phone string
//}

//var m = make(map[string]repository.Contact)
//var inc = repository.AI

func main() {

	router := httprouter.New()

	router.GET("/user/:id", getUser)
	router.GET("/users/list/", getUserList)
	router.POST("/user/", addUser)
	router.PUT("/user/:id", updateUser)
	router.DELETE("/user/:id", deleteUser)

	fmt.Println("Start server")
	log.Fatal(http.ListenAndServe(":8081", router))
}
func addUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user := repository.Contact{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err.Error())
		return
	}

	name := Phonebook.CheckName(user.Name)
	if name {
		fmt.Fprintln(w, "Есть совподения имени.")
	} else {
		fmt.Fprintln(w, "Нет совподения имени.")
	}
	phone := Phonebook.CheckPhone(user.Phone)
	if phone {
		fmt.Fprintln(w, "Есть совподения телефона.")
	} else {
		fmt.Fprintln(w, "Нет совподения телефона.")
	}
	Phonebook.AddContact(user)
}

func getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ok := Phonebook.FindById(ps.ByName("id"))
	if ok {
		m := Phonebook.GetContact(ps.ByName("id"))

		fmt.Fprintf(w, "%s - %s - %s\n", ps.ByName("id"), m.Name, m.Phone)
	} else {
		fmt.Fprintf(w, "Контакт отсутствует.")
	}
}

func getUserList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Список контактов:")
	fmt.Fprintln(w, "Количество контактов: ", Phonebook.CountContacts())
	for _, v := range Phonebook.GetContactList() {
		fmt.Fprintf(w, "Телефон: - %s Имя: - %s\n", v.Phone, v.Name)
	}
	return
	fmt.Fprintf(w, "Список контактов пуст.")
}

func updateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ok := Phonebook.FindById(ps.ByName("id"))
	if !ok {
		fmt.Fprintf(w, "Контакт отсутствуют.")
		return
	}
	user := repository.Contact{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Println(err.Error())
	} else {
		Phonebook.UpdateContact(ps.ByName("id"), user)
		fmt.Fprintf(w, "Контакт измененю")
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ok := Phonebook.FindById(ps.ByName("id"))
	if !ok {
		fmt.Fprintf(w, "Контакт отсутствуют.")
	} else {
		Phonebook.DeleteContact(ps.ByName("id"))
		fmt.Fprintf(w, "Контакт удален.")
	}
}
