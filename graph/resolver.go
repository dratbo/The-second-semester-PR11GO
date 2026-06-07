package graph

import "example.com/pz11-graphql/internal/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	TaskService *service.TaskService
}
