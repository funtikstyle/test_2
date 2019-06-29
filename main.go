package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
	"test_2/repository"
)

var Phonebook = make(repository.PhoneBook)

type User struct {
	Name  string
	Phone string
}

var m = make(map[string]User)
var inc = 1

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
	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err.Error())
		return
	}

	_, ok := m[strconv.Itoa(inc)]
	if ok {
		for i := 0; i < 10; i++ {
			inc++
			_, ok := m[strconv.Itoa(inc)]
			if !ok {
				break
			}
		}
		if ok {
			fmt.Fprintf(w, "Не удалось добавить контакт , мы сделали все что могли!!!")
		}
	}

	if !ok {
		m[strconv.Itoa(inc)] = user
		fmt.Fprintf(w, "Контакт добюавлен.")
		inc++
	}

}

func getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, ok := m[ps.ByName("id")]
	if ok {
		fmt.Fprintf(w, "%s - %s - %s\n", ps.ByName("id"), m[ps.ByName("id")].Name, m[ps.ByName("id")].Phone)
	} else {
		fmt.Fprintf(w, "Контакт отсутствует.")
	}
}

func getUserList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("getUserList")

	if len(m) != 0 {
		for i := range m {
			fmt.Fprintf(w, "%s - %s - %s\n", i, m[i].Name, m[i].Phone)
		}
		return
	}
	fmt.Fprintf(w, "Список контактов пуст.")
}

func updateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, ok := m[ps.ByName("id")]
	if !ok {
		fmt.Fprintf(w, "Контакт отсутствуют.")
		return
	}
	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Println(err.Error())
	} else {
		m[ps.ByName("id")] = user
		fmt.Fprintf(w, "")
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, ok := m[ps.ByName("id")]
	if ok {
		delete(m, ps.ByName("id"))
		fmt.Fprintf(w, "Контакт удален.")
	} else {
		fmt.Fprintf(w, "Kонтакт не найден")
	}
}
