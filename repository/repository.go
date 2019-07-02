package repository

import (
	"strconv"
)

type Contact struct {
	Name  string
	Phone string
}
type PhoneBook map[string]Contact

var AI = 0

func (pb *PhoneBook) FindById(key string) bool {
	_, ok := (*pb)[key]
	return ok
}
func (pb *PhoneBook) GetContact(key string) Contact {
	return (*pb)[key]
}
func (pb *PhoneBook) GetContactList() []Contact {
	var list []Contact
	for _, item := range *pb {
		list = append(list, item)
	}
	return list
}
func (pb *PhoneBook) AddContact(c Contact) {
	(*pb)[strconv.Itoa(AI)] = c
	AI++
}
func (pb *PhoneBook) DeleteContact(key string) {
	delete((*pb), key)
}
func (pb *PhoneBook) UpdateContact(key string, c Contact) {
	(*pb)[key] = c
}

func (pb *PhoneBook) CountContacts() string {
	s := len(*pb)
	return strconv.Itoa(s)
} // метод возвращяет количества в репозитории

func (pb *PhoneBook) CheckName(key string) bool {

	for _, item := range *pb {
		if item.Name == key {
			return true
		}
	}
	return false
}

func (pb *PhoneBook) CheckPhone(key string) bool {
	for _, item := range *pb {
		if item.Phone == key {
			return true
		}
	}
	return false
}
