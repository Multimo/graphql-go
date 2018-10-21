//go:generate gorunpkg github.com/99designs/gqlgen

package graphql_go

import (
	context "context"
)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateJob(ctx context.Context, input NewJob) (Job, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteJob(ctx context.Context, id string) (string, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateApplication(ctx context.Context, input NewApplication) (Application, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Jobs(ctx context.Context) ([]Job, error) {
	panic("not implemented")
}
func (r *queryResolver) Applications(ctx context.Context, jobID string) ([]Application, error) {
	panic("not implemented")
}
