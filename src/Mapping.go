package src

var (
	FieldTypeMapping = map[string]string{
		"varchar":   "string",
		"text":      "string",
		"timestamp": "*LocalTime",
		"bigint":    "int64",
		"int":       "int",
		"tinyint":   "int8",
		"float":     "float64",
		"decimal":   "string",
	}
)
