package sql

type Clause string

const (
	SELECT Clause = "SELECT"
	INSERT Clause = "INSERT"
	UPDATE Clause = "UPDATE"
	DELETE Clause = "DELETE"
)
