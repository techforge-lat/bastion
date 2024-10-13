package di

import "github.com/techforge-lat/linkit"

type Container struct {
	*linkit.DependencyContainer
}

func NewContainer() *Container {
	return &Container{
		DependencyContainer: linkit.New(),
	}
}
