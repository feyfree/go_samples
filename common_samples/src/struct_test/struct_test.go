package struct_test

import (
	"fmt"
	"testing"
)

import (
	"encoding/json"
)

type Person struct {
	ID        int
	FirstName string `json:"name"`
	LastName  string
	Address   string `json:"address,omitempty"`
}

type Employee struct {
	Person
	ManagerID int
}

type Contractor struct {
	Person
	CompanyID int
}

func TestStruct(t *testing.T) {
	employees := []Employee{
		{
			Person: Person{
				LastName: "Doe", FirstName: "John",
			},
		},
		{
			Person: Person{
				LastName: "Campbell", FirstName: "David",
			},
		},
	}

	data, _ := json.Marshal(employees)
	fmt.Printf("%s\n", data)

	var decoded []Employee
	_ = json.Unmarshal(data, &decoded)
	fmt.Printf("%v", decoded)
}
