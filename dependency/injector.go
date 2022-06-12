//go:build wireinject
// +build wireinject

package dependency

import "github.com/google/wire"

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository,
		NewSimpleService,
	)
	return nil, nil
}

func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabaseMongo,
		NewDatabasePgsql,
		NewDatabaseRepository,
	)
	return nil
}
