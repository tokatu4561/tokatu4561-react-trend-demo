package drivers

import (
	"context"
	"myapp/adapter/controllers"
)

// func NewInputFactory() controllers.InputFactory {

// }

func InitializeTaskDriver() (User, error) {
	wire.Build(, NewOutputFactory, NewInputFactory, NewRepositoryFactory, controllers.NewUserController, NewUserDriver)
	return &UserDriver{}, nil
}
