package main

import (
	"testing"
)

func TestInfo(t *testing.T) {
	s1 := Student{
		name: "Harry",
		age:  20,
		Contacts: Contacts{
			phone: "98793729200", email: "harry.potter@sample.com",
		},
	}
	ans := s1.age
	if ans != 20 {
		t.Errorf("Student age = %d; want 20", ans)
	}

	ansEmail := s1.email
	ansMobile := s1.phone
	if (ansEmail == "" || ansMobile == "") {
		t.Error("Student info is incomplete.")
	}
}