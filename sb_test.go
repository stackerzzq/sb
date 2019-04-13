package sb

import (
	"testing"
	"time"
)

func Test_sb_KeyVals(t *testing.T) {
	type Employment struct {
		JobNumber int     `sb:"job_number"`
		Salary    float64 `sb:"salary"`
	}

	type person struct {
		Name       string     `sb:"name"`
		Age        int        `sb:"age"`
		Email      string     `sb:"email"`
		StuOrNot   bool       `sb:"stu_or_not"`
		Birthday   time.Time  `sb:"birthday"`
		Hobby      []string   `sb:"hobby"`
		Employment Employment `sb:"employment"`
	}

	s := New()

	ab, _ := time.Parse("2006-01-02", "2010-10-10")
	s.KeyVals(person{
		Name:       "aaaa",
		Age:        29,
		Email:      "aaa@qq.net",
		StuOrNot:   false,
		Birthday:   ab,
		Hobby:      []string{"music11111", "movie22222"},
		Employment: Employment{JobNumber: 11111, Salary: 99.99},
	})
	if len(s.kvs) != 1 {
		t.Error("s.KeyVals failed")
	}

	bb, _ := time.Parse("2006-01-02", "2010-10-10")
	p := []person{
		{
			Name:       "aaaa",
			Age:        29,
			Email:      "aaa@qq.net",
			StuOrNot:   false,
			Birthday:   ab,
			Hobby:      []string{"music11111", "movie22222"},
			Employment: Employment{JobNumber: 11111, Salary: 99.99},
		},
		{
			Name:       "bbb",
			Age:        28,
			Email:      "bbb@qq.net",
			StuOrNot:   false,
			Birthday:   bb,
			Hobby:      []string{"music", "movie"},
			Employment: Employment{JobNumber: 22222, Salary: 99.99},
		},
	}

	s.KeyVals(p)
	if len(s.kvs) != 2 {
		t.Error("KeyVals failed")
	}

	m := map[string]int{"aaa": 111, "bbb": 222}
	s.KeyVals(m)
	if len(s.kvs) != 1 {
		t.Error("KeyVals failed")
	}

	ms := []map[string]int{
		map[string]int{"aaa": 111, "bbb": 222, "ccc": 333},
		map[string]int{"aaa": 111, "bbb": 222, "ddd": 444},
	}
	s.KeyVals(ms)
	if len(s.kvs) != 2 {
		t.Error("KeyVals failed")
	}
}
