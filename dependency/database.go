package dependency

type Database struct {
	Name string
}

// tipe data alias
type PostgreeSQL Database
type MongoDB Database

func NewDatabasePgsql() *PostgreeSQL {
	return (*PostgreeSQL)(&Database{Name: "PostgreeSQL"})
}

func NewDatabaseMongo() *MongoDB {
	return (*MongoDB)(&Database{Name: "MongoDB"})
}

type DatabaseRepository struct {
	DatatasePgsql *PostgreeSQL
	DatabaseMongo *MongoDB
}

func NewDatabaseRepository(pgsql *PostgreeSQL, mongo *MongoDB) *DatabaseRepository {
	return &DatabaseRepository{
		DatatasePgsql: pgsql,
		DatabaseMongo: mongo,
	}
}
