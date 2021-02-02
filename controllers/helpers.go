package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

func parseForm(r *http.Request, dst interface{}) error {
	// causes the PostForm to be filled
	if err := r.ParseForm(); err != nil {
		return err
	}

	// decode the PostForm into our dst handled by gorilla/schema
	dec := schema.NewDecoder()
	if err := dec.Decode(dst, r.PostForm); err != nil {
		return err
	}
	return nil
}
