package sql

type Operator string

const (
	Equals    Operator = "="
	Unequals  Operator = "!="
	Like      Operator = "LIKE"
	Less      Operator = "<"
	Greater   Operator = ">"
	LessEq    Operator = "<="
	GreaterEq Operator = ">="
	In        Operator = "IN"
)
