package cli

import "github.com/090809/marlow/examples/library/models"

// Command represents a type that used by the example app cli.
type Command func(*models.Stores, []string) error
