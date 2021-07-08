// Code generated by goa v3.4.3, DO NOT EDIT.
//
// calc client
//
// Command:
// $ goa gen calc/design

package calc

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "calc" service client.
type Client struct {
	AddEndpoint      goa.Endpoint
	DivEndpoint      goa.Endpoint
	RedirectEndpoint goa.Endpoint
}

// NewClient initializes a "calc" service client given the endpoints.
func NewClient(add, div, redirect goa.Endpoint) *Client {
	return &Client{
		AddEndpoint:      add,
		DivEndpoint:      div,
		RedirectEndpoint: redirect,
	}
}

// Add calls the "add" endpoint of the "calc" service.
func (c *Client) Add(ctx context.Context, p *AddPayload) (res int, err error) {
	var ires interface{}
	ires, err = c.AddEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}

// Div calls the "div" endpoint of the "calc" service.
// Div may return the following errors:
//	- "zero_division" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Div(ctx context.Context, p *DivPayload) (res int, err error) {
	var ires interface{}
	ires, err = c.DivEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}

// Redirect calls the "redirect" endpoint of the "calc" service.
func (c *Client) Redirect(ctx context.Context) (err error) {
	_, err = c.RedirectEndpoint(ctx, nil)
	return
}
