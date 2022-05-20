package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-graphql-mongodb-api/database"
	"go-graphql-mongodb-api/graph/generated"
	"go-graphql-mongodb-api/graph/model"
)

func (r *myMutationResolver) CreateTodo(ctx context.Context, todo model.TodoInput) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *myMutationResolver) UpdateTodo(ctx context.Context, id string, updatedTodo model.TodoInput) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *myQueryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *myQueryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// MyMutation returns generated.MyMutationResolver implementation.
func (r *Resolver) MyMutation() generated.MyMutationResolver { return &myMutationResolver{r} }

// MyQuery returns generated.MyQueryResolver implementation.
func (r *Resolver) MyQuery() generated.MyQueryResolver { return &myQueryResolver{r} }

type myMutationResolver struct{ *Resolver }
type myQueryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) CreateCourse(ctx context.Context, input model.NewCourse) (*model.Course, error) {
	return db.InsertCourseByID(input), nil
}
func (r *mutationResolver) CreateInstructor(ctx context.Context, input *model.AddInstructor) (*model.Instructor, error) {
	return db.InsertInstructor(input), nil
}
func (r *mutationResolver) UpdateCourse(ctx context.Context, courseID string, name *string, subject *string, instructorID string) (*model.Course, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) DeleteCourse(ctx context.Context, courseID string) (*model.Course, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) Course(ctx context.Context, id string) (*model.Course, error) {
	return db.FindCourseById(id), nil
}
func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	return db.All(), nil
}
func (r *queryResolver) Instructor(ctx context.Context, id string) (*model.Instructor, error) {
	return db.FindInstructorById(id), nil
}
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

var db = database.Connect("mongodb://localhost:27017/")
