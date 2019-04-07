package sb

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

type sb struct {
	action string
}

func New() *sb {
	return new(sb)
}

func (s *sb) SetAction(action string) *sb {
	s.action = strings.ToLower(action)

	return s
}

func (s *sb) GetAction() string {
	return s.action
}

func (s *sb) Bind(stub interface{}) *sb {
	switch s.action {
	case "insert":
	case "update":
	case "delete":
	case "select":
	default:
		log.Fatal("unknown action, please specify the right sql action")
	}

	return s
}

/*
type Person struct {
	Name       string    `sb:name`
	Age        int       `sb:age`
	Email      string    `sb:email`
	StuOrNot   bool      `sb:stu_or_not`
	Birthday   time.Time `sb:birthday`
	Hobby      []string  `sb:hobby`
	Employment struct {
		JobNumber int     `sb:job_number`
		Salary    float64 `sb:salary`
	} `sb:employment`
}
*/
func (s *sb) Insert(x interface{}) *sb {
	switch val := reflect.ValueOf(x); val.Kind() {
	case reflect.Map:
		ks := make([]string, 0, 0)
		vs := make([]interface{}, 0, 0)
		for _, key := range val.MapKeys() {
			fmt.Println(key, val.MapIndex(key))
			ks = append(ks, key)
		}
	case reflect.Struct:

	}
	return s
}

func (s *sb) Update(stub interface{}) *sb {
	return s
}

func (s *sb) Delete(stub interface{}) *sb {
	return s
}

func (s *sb) Select(stub interface{}) *sb {
	return s
}

func (s *sb) EchoSQL() string {
	return ""
}
