// Package main provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.3 DO NOT EDIT.
package main

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns all Todos
	// (GET /todos)
	GetTodos(ctx echo.Context, params GetTodosParams) error

	// (POST /todos)
	CreateTodo(ctx echo.Context) error

	// (DELETE /todos/{todoId})
	DeleteTodo(ctx echo.Context, todoId int32) error

	// (PUT /todos/{todoId})
	UpdateTodo(ctx echo.Context, todoId int32) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetTodos converts echo context to params.
func (w *ServerInterfaceWrapper) GetTodos(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetTodosParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTodos(ctx, params)
	return err
}

// CreateTodo converts echo context to params.
func (w *ServerInterfaceWrapper) CreateTodo(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateTodo(ctx)
	return err
}

// DeleteTodo converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteTodo(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "todoId" -------------
	var todoId int32

	err = runtime.BindStyledParameterWithLocation("simple", false, "todoId", runtime.ParamLocationPath, ctx.Param("todoId"), &todoId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter todoId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteTodo(ctx, todoId)
	return err
}

// UpdateTodo converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateTodo(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "todoId" -------------
	var todoId int32

	err = runtime.BindStyledParameterWithLocation("simple", false, "todoId", runtime.ParamLocationPath, ctx.Param("todoId"), &todoId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter todoId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateTodo(ctx, todoId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/todos", wrapper.GetTodos)
	router.POST(baseURL+"/todos", wrapper.CreateTodo)
	router.DELETE(baseURL+"/todos/:todoId", wrapper.DeleteTodo)
	router.PUT(baseURL+"/todos/:todoId", wrapper.UpdateTodo)

}
