package controllers

import "lenslocked.com/views"

// NewStatic instantiates the Static type
func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "static/home"),
		Contact: views.NewView("bootstrap", "static/contact"),
	}
}

// Static represents the static pages
type Static struct {
	Home    *views.View
	Contact *views.View
}
