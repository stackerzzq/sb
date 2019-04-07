package sb

const (
	tmplForInsert = `INSERT INTO :table (field1, field2...) VALUES (:field1, :field2...)`
)

func (s *sb) BuildInsertSql() {

}
