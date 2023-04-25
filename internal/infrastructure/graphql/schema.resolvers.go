package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"

	"github.com/ac2393921/bigami-sheet-api/internal/infrastructure/graphql/model"
)

// Schools is the resolver for the schools field.
func (r *queryResolver) Schools(ctx context.Context) ([]*model.School, error) {
	return []*model.School{
		{
			ID:    "SCHOOL-1",
			Name:  "斜歯忍軍",
			Style: "他の流派の「奥義の内容」を集める。",
			Enemy: "鞍馬神流",
		},
		{
			ID:    "SCHOOL-2",
			Name:  "鞍馬神流",
			Style: "シノビガミの復活を阻止する。",
			Enemy: "隠忍の血統",
		},
	}, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
