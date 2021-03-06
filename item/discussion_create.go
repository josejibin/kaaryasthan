package controller

import (
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/jsonapi"
	"github.com/kaaryasthan/kaaryasthan/item/model"
	"github.com/kaaryasthan/kaaryasthan/user/model"
)

// CreateDiscussionHandler creates discussion
func (c *DiscussionController) CreateDiscussionHandler(w http.ResponseWriter, r *http.Request) {
	tkn := r.Context().Value("user").(*jwt.Token)
	userID := tkn.Claims.(jwt.MapClaims)["sub"].(string)

	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusCreated)

	usr := &user.User{ID: userID}
	if err := c.uds.Valid(usr); err != nil {
		log.Println("Couldn't validate user: ", err)
		return
	}

	disc := new(item.Discussion)
	if err := jsonapi.UnmarshalPayload(r.Body, disc); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	itm := &item.Item{ID: disc.ItemID}
	if err := c.ids.Valid(itm); err != nil {
		log.Println("Couldn't validate item: ", err)
		return
	}

	err := c.ds.Create(usr, disc)
	if err != nil {
		log.Println("Unable to save data: ", err)
		return
	}

	if err := jsonapi.MarshalPayload(w, disc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
