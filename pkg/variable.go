package pkg

var (
	FieldTypeMapping = map[string]string{
		"varchar":   "string",
		"text":      "string",
		"timestamp": "*LocalTime",
		"bigint":    "int64",
		"int":       "int32",
		"tinyint":   "int8",
		"float":     "float64",
		"decimal":   "string",
		"longtext":  "string",
	}

	FieldDbMapping = map[string]string{
		"varchar": "varchar(255)",
		"text":    "text",
	}
)

type Field struct {
	Name      string `json:"name"`
	Json      string `json:"json"`
	Default   string `json:"default"`
	Describe  string `json:"describe"`
	Primary   bool   `json:"primary"`
	Index     string `json:"index"`
	IsNull    bool   `json:"isNull"`
	Type      string `json:"type"`
	Transform string `json:"Transform"`
}

type TableModal struct {
	Name   string
	Fields []Field
}

type ControllerVar struct {
	TableModal
	Pkg string
}
