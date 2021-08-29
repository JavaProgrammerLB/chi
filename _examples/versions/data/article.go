package data

// Article is runtime object, that's not meant to be sent via REST.
// 结构体Article，字段名、字段类型、``里的一些属性（db, json, xml）
type Article struct {
	ID                     int      `db:"id" json:"id" xml:"id"`
	Title                  string   `db:"title" json:"title" xml:"title"`
	Data                   []string `db:"data,stringarray" json:"data" xml:"data"`
	CustomDataForAuthUsers string   `db:"custom_data" json:"-" xml:"-"`
}
