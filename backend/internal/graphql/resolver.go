package graphql

import "github.com/example/fullstack-template/internal/service"

type Resolver struct {
	Auth *service.AuthService
}
