package shared

import "gorm.io/gorm"

type Operator string

const (
	OpEq   Operator = "="
	OpGt   Operator = ">"
	OpLt   Operator = "<"
	OpGte  Operator = ">="
	OpLte  Operator = "<="
	OpIn   Operator = "IN"
	OpNot  Operator = "NOT"
	OpLike Operator = "LIKE"
)

type Criterion struct {
	Field    string
	Operator Operator
	Value    any
	Or       bool
}

func ApplyCriterion(db *gorm.DB, c Criterion) *gorm.DB {
	clause := ""
	switch c.Operator {
	case OpEq, OpGt, OpLt, OpGte, OpLte:
		clause = c.Field + " " + string(c.Operator) + " ?"
	case OpLike:
		clause = c.Field + " LIKE ?"
	case OpIn:
		//TODO parser array to string (...values)
		clause = c.Field + " IN ?"
	case OpNot:
		clause = "NOT " + c.Field + " = ?"
	}

	if c.Or {
		return db.Or(clause, c.Value)
	}

	return db.Where(clause, c.Value)
}

func ApplyCriteria(db *gorm.DB, criteria []Criterion) *gorm.DB {
	query := db
	for _, c := range criteria {
		query = ApplyCriterion(query, c)
	}

	return query
}
