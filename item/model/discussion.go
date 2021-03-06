package item

import (
	"database/sql"
	"errors"

	"github.com/kaaryasthan/kaaryasthan/user/model"
)

// DiscussionRepository to manage items
type DiscussionRepository interface {
	Create(usr *user.User, disc *Discussion) error
	Valid(itm *Discussion) error
}

// Discussion represents a discussion
type Discussion struct {
	ID     string `jsonapi:"primary,discussions"`
	Body   string `jsonapi:"attr,body"`
	ItemID int    `jsonapi:"attr,item_id"`
}

// DiscussionDatastore implements the Repository interface
type DiscussionDatastore struct {
	db *sql.DB
}

// NewDiscussionDatastore constructs a new Repository
func NewDiscussionDatastore(db *sql.DB) *DiscussionDatastore {
	return &DiscussionDatastore{db}
}

// Valid checks the validity of the discussion
func (ds *DiscussionDatastore) Valid(disc *Discussion) error {
	var count int
	err := ds.db.QueryRow(`SELECT count(1) FROM "discussions"
		WHERE id=$1 AND deleted_at IS NULL`,
		disc.ID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Invalid item")
	}
	err = ds.db.QueryRow(`SELECT count(1) FROM "discussions"
		INNER JOIN items ON discussions.item_id = items.id
		INNER JOIN projects ON items.project_id = projects.id
		WHERE discussions.id=$1 AND items.lock_conversation=false
		AND items.deleted_at IS NULL AND projects.archived=false
		AND projects.deleted_at IS NULL`,
		disc.ID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Invalid discussion")
	}
	return nil
}
