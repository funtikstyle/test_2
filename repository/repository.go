package repository

import (
	"strconv"
)

type Contact struct {
	Name  string
	Phone string
}
type PhoneBook map[string]Contact

var ai = 0

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
	(*pb)[strconv.Itoa(ai)] = c
	ai++
}
func (pb *PhoneBook) DeleteContact(key string) {
	delete((*pb), key)
}
func (pb *PhoneBook) UpdateContact(key string, c Contact) {
	(*pb)[key] = c
}
