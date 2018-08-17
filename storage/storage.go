package storage

import (
	"database/sql"
	"github.com/pkg/errors"
	"github.com/reddec/state-machine/machine"
	"os"
	"time"
)

type DbStorage struct {
	db XODB
}

func (ms *DbStorage) Append(ctx *machine.StateContext, state machine.State, e error) error {
	str := sql.NullString{}
	if e != nil {
		str.Valid = true
		str.String = e.Error()
	}
	record := State{
		State:           int(state),
		Event:           ctx.Event,
		Data:            ctx.Storage,
		ContextID:       ctx.ID,
		CreatedAt:       time.Now(),
		ProcessingError: str,
	}
	return record.Insert(ms.db)
}

func (ms *DbStorage) Last(id string) (*machine.IncompleteStateContext, error) {
	item, err := LastState(ms.db, id)
	if err == sql.ErrNoRows {
		return nil, os.ErrNotExist
	} else if err != nil {
		return nil, err
	}
	return &machine.IncompleteStateContext{
		ID:      item.ContextID,
		Storage: item.Data,
		Current: machine.State(item.State),
	}, nil
}

func NewDbStorage(db XODB) (*DbStorage, error) {
	_, err := db.Exec(string(MustAsset("init.sql")))
	if err != nil {
		return nil, errors.Wrap(err, "db storage - initialize db structure")
	}
	return &DbStorage{
		db: db,
	}, nil
}
