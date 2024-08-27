package graph

//go:generate go run github.com/99designs/gqlgen generate

import "github.com/shunta0213/go-graphql/graph/model"

type Resolver struct {
	todos []*model.Todo
}
