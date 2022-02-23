package adapters

import "context"

type UserHandler struct {
	service UserService
}
type UserService interface {
	Create
	GetAll
	Update
	Delete
}

type GroupHandler struct {
	service GroupService
}
type GroupService interface {
	Create
	GetAll
	Update
	Delete
	CreateSub
}

type Create interface {
	Create(context.Context, []byte)
}

type GetAll interface {
	GetAll(context.Context, []byte) []byte
}

type Update interface {
	Update(context.Context, []byte)
}

type Delete interface {
	Delete(context.Context, []byte)
}

type CreateSub interface {
	CreateSub(context.Context, []byte)
}

//...
func (u *UserHandler) Create(context.Context, []byte) {

}

func (u *UserHandler) GetAll(context.Context, []byte) []byte {
	var a []byte // заглушка
	return a
}

func (u *UserHandler) Update(context.Context, []byte) {

}

func (u *UserHandler) Delete(context.Context, []byte) {

}

func (g *GroupHandler) Create(context.Context, []byte) {

}

func (g *GroupHandler) GetAll(context.Context, []byte) []byte {
	var a []byte // заглушка
	return a
}

func (g *GroupHandler) Update(context.Context, []byte) {

}

func (g *GroupHandler) Delete(context.Context, []byte) {

}

func (g *GroupHandler) CreateSub(context.Context, []byte) {

}
