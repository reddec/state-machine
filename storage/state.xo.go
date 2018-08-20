// Package storage contains the types for schema 'public'.
package storage

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// State represents a row from 'public.state'.
type State struct {
	ID              int64          `json:"id"`               // id
	ContextID       string         `json:"context_id"`       // context_id
	CreatedAt       time.Time      `json:"created_at"`       // created_at
	State           int            `json:"state"`            // state
	Data            []byte         `json:"data"`             // data
	Event           []byte         `json:"event"`            // event
	ProcessingError sql.NullString `json:"processing_error"` // processing_error

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the State exists in the database.
func (s *State) Exists() bool {
	return s._exists
}

// Deleted provides information if the State has been deleted from the database.
func (s *State) Deleted() bool {
	return s._deleted
}

// Insert inserts the State to the database.
func (s *State) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if s._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.state (` +
		`"context_id", "created_at", "state", "data", "event", "processing_error"` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`) RETURNING "id"`

	// run query
	XOLog(sqlstr, s.ContextID, s.CreatedAt, s.State, s.Data, s.Event, s.ProcessingError)
	err = db.QueryRow(sqlstr, s.ContextID, s.CreatedAt, s.State, s.Data, s.Event, s.ProcessingError).Scan(&s.ID)
	if err != nil {
		return err
	}

	// set existence
	s._exists = true

	return nil
}

// Update updates the State in the database.
func (s *State) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !s._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if s._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.state SET (` +
		`"context_id", "created_at", "state", "data", "event", "processing_error"` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6` +
		`) WHERE "id" = $7`

	// run query
	XOLog(sqlstr, s.ContextID, s.CreatedAt, s.State, s.Data, s.Event, s.ProcessingError, s.ID)
	_, err = db.Exec(sqlstr, s.ContextID, s.CreatedAt, s.State, s.Data, s.Event, s.ProcessingError, s.ID)
	return err
}

// Save saves the State to the database.
func (s *State) Save(db XODB) error {
	if s.Exists() {
		return s.Update(db)
	}

	return s.Insert(db)
}

// Upsert performs an upsert for State.
//
// NOTE: PostgreSQL 9.5+ only
func (s *State) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if s._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.state (` +
		`"id", "context_id", "created_at", "state", "data", "event", "processing_error"` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`) ON CONFLICT ("id") DO UPDATE SET (` +
		`"id", "context_id", "created_at", "state", "data", "event", "processing_error"` +
		`) = (` +
		`EXCLUDED."id", EXCLUDED."context_id", EXCLUDED."created_at", EXCLUDED."state", EXCLUDED."data", EXCLUDED."event", EXCLUDED."processing_error"` +
		`)`

	// run query
	XOLog(sqlstr, s.ID, s.ContextID, s.CreatedAt, s.State, s.Data, s.Event, s.ProcessingError)
	_, err = db.Exec(sqlstr, s.ID, s.ContextID, s.CreatedAt, s.State, s.Data, s.Event, s.ProcessingError)
	if err != nil {
		return err
	}

	// set existence
	s._exists = true

	return nil
}

// Delete deletes the State from the database.
func (s *State) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !s._exists {
		return nil
	}

	// if deleted, bail
	if s._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.state WHERE "id" = $1`

	// run query
	XOLog(sqlstr, s.ID)
	_, err = db.Exec(sqlstr, s.ID)
	if err != nil {
		return err
	}

	// set deleted
	s._deleted = true

	return nil
}

// StateByID retrieves a row from 'public.state' as a State.
//
// Generated from index 'state_pkey'.
func StateByID(db XODB, id int64) (*State, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`"id", "context_id", "created_at", "state", "data", "event", "processing_error" ` +
		`FROM public.state ` +
		`WHERE "id" = $1`

	// run query
	XOLog(sqlstr, id)
	s := State{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&s.ID, &s.ContextID, &s.CreatedAt, &s.State, &s.Data, &s.Event, &s.ProcessingError)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

// LastState runs a custom query, returning results as State.
func LastState(db XODB, contextId string) (*State, error) {
	var err error

	// sql query
	const sqlstr = `SELECT * FROM state WHERE context_id = $1 ORDER BY created_at DESC LIMIT 1`

	// run query
	XOLog(sqlstr, contextId)
	var s State
	err = db.QueryRow(sqlstr, contextId).Scan(&s.ID, &s.ContextID, &s.CreatedAt, &s.State, &s.Data, &s.Event, &s.ProcessingError)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

// OldestInState runs a custom query, returning results as State.
func OldestInState(db XODB, state int64) (*State, error) {
	var err error

	// sql query
	const sqlstr = `SELECT * FROM state WHERE state = $1 ORDER BY created_at DESC LIMIT 1`

	// run query
	XOLog(sqlstr, state)
	var s State
	err = db.QueryRow(sqlstr, state).Scan(&s.ID, &s.ContextID, &s.CreatedAt, &s.State, &s.Data, &s.Event, &s.ProcessingError)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
