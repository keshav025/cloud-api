package domain

import "context"

type Command interface {
	Create(ctx context.Context, obj interface{}) error
	Update(ctx context.Context, obj interface{}, objId int32) error
	Delete(ctx context.Context, obj interface{}, objId int32) error
}

type Query interface {
	Get(ctx context.Context, obj interface{}, objId int32) error
	List(ctx context.Context, objList interface{}, filters ...interface{}) error
}