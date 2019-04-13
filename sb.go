package sb

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	tplForInsert  = `INSERT INTO :TABLE (%s) VALUES (%s)`
	tplForUpdate  = `UPDATE :TABLE SET %s :WHERE :LIMIT :OFFSET`
	tplForDelete  = `DELETE FROM :TABLE :WHERE :LIMIT :OFFSET`
	tplForSelect  = `SELECT %s FROM :TABLE :WHERE :GROUPBY :ORDERBY :LIMIT :OFFSET`
	tplForWhere   = `WHERE %s`
	tplForGroupBy = `GROUP BY %s`
	tplForHaving  = `HAVING %s`
	tplForOrderBy = `ORDER BY %s`
	tplForLimit   = `LIMIT %d`
	tplForOffset  = `OFFSET %d`
)

type sb struct {
	table   string
	kvs     []map[string]interface{}
	columns string
	where   map[string]interface{}
	groupBy string
	having  string
	orderBy string
	limit   string
	offset  string
	tag     string
}

func New() *sb {
	return &sb{
		tag: "sb",
	}
}

//SetTag change the default tag `sb` to other string, please make sure exists in your struct
func (s *sb) SetTag(tag string) *sb {
	s.tag = tag
	return s
}

func (s *sb) GetTag() string {
	return s.tag
}

func (s *sb) Table(table string) *sb {
	s.table = table
	return s
}

func (s *sb) KeyVals(kvs interface{}) *sb {
	s.kvs = bind(kvs, s.tag)
	return s
}

func bind(kvs interface{}, stag string) []map[string]interface{} {
	ms := make([]map[string]interface{}, 0, 0)
	t := reflect.TypeOf(kvs)
	v := reflect.ValueOf(kvs)
	switch kind := v.Kind(); kind {
	case reflect.Struct:
		m := make(map[string]interface{}, 0)
		for i := 0; i < v.NumField(); i++ {
			if tag := t.Field(i).Tag.Get(stag); tag != "" && tag != "_" && reflect.ValueOf(v.Field(i).Interface()).Kind() != reflect.Struct {
				m[tag] = v.Field(i).Interface()
			}
		}
		ms = append(ms, m)
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			m := bind(v.Index(i).Interface(), stag)
			ms = append(ms, m...)
		}
	case reflect.Map:
		m := make(map[string]interface{}, 0)
		for _, k := range v.MapKeys() {
			if k.Type().Kind() == reflect.String {
				m[k.String()] = v.MapIndex(k).Interface()
			}
		}
		ms = append(ms, m)
	}

	return ms
}

func (s *sb) Columns(columns []string) *sb {
	s.columns = strings.Join(columns, ",")
	return s
}

func (s *sb) Where() *sb {
	return s
}

func (s *sb) GroupBy() *sb {
	return s
}

func (s *sb) Having() *sb {
	return s
}

func (s *sb) OrderBy() *sb {
	return s
}

func (s *sb) Limit(limit int64) *sb {
	s.limit = fmt.Sprintf(tplForLimit, limit)
	return s
}

func (s *sb) Offset(offset int64) *sb {
	s.offset = fmt.Sprintf(tplForOffset, offset)
	return s
}
