package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"
	dbmodel "typebeast-service/database/model"
	"typebeast-service/graph/model"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"go.uber.org/zap"
)

// CreateWritingSample is the resolver for the createWritingSample field.
func (r *mutationResolver) CreateWritingSample(ctx context.Context, input model.CreateWritingSampleInput) (*model.WritingSample, error) {
	logger := r.Resolver.Logger
	logger.Info("creating sample", zap.Reflect("input", input))

	claims, ok := clerk.SessionFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("invalid session claims")
	}

	sample := dbmodel.WritingSample{
		Title:   input.Title,
		Content: input.Content,
		UserID:  claims.Subject,
	}
	r.DB.Create(&sample)

	return &model.WritingSample{
		ID:      sample.ID,
		Title:   sample.Title,
		Content: sample.Content,
	}, nil
}

// RecordUserPerformance is the resolver for the recordUserPerformance field.
func (r *mutationResolver) RecordUserPerformance(ctx context.Context, input model.PerformanceInput) (bool, error) {
	r.Resolver.Logger.Info("logging performance", zap.Reflect("input", input))
	return true, nil
}

// GetSamples is the resolver for the getSamples field.
func (r *queryResolver) GetSamples(ctx context.Context) ([]*model.WritingSample, error) {
	logger := r.Resolver.Logger
	logger.Info("getting samples")
	claims, ok := clerk.SessionFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("invalid session claims")
	}

	var samples []dbmodel.WritingSample
	r.DB.Where("user_id = ?", claims.Subject).Find(&samples)

	var gqlSamples []*model.WritingSample

	for _, sample := range samples {
		gqlSamples = append(gqlSamples, &model.WritingSample{
			ID:      sample.ID,
			Title:   sample.Title,
			Content: sample.Content,
		})
	}

	return gqlSamples, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
