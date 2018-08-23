package storage

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/reddec/state-machine/machine"
	"os"
	"strconv"
	"strings"
	"time"
)

type DbStorage struct {
	db *sqlx.DB
}

func (ms *DbStorage) OldestContextIDInState(state machine.State) (string, error) {
	st, err := OldestInState(ms.db, int64(state))
	if err != nil {
		return "", err
	}
	return st.ContextID, nil
}

func (ms *DbStorage) NumNotInStates(state ...machine.State) (int64, error) {
	if len(state) == 0 {
		return -1, errors.New("no states provides")
	}

	var opts []string
	var params []interface{}
	for i, val := range state {
		opts = append(opts, "$"+strconv.Itoa(i+1))
		params = append(params, val)
	}
	var queryBase = `SELECT count(1) FROM "state" WHERE "state"."state" NOT IN (` + strings.Join(opts, ",") + `)`
	var count int64
	rs := ms.db.QueryRow(queryBase, params...)
	err := rs.Scan(&count)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return count, err
}

func (ms *DbStorage) Alias(originalId string, alias string) error {
	var item = Alias{
		Alias:     alias,
		ContextID: originalId,
	}
	return item.Insert(ms.db)
}

func (ms *DbStorage) Append(ctx *machine.StateContext, state machine.State, e error) error {
	str := sql.NullString{}
	if e != nil {
		str.Valid = true
		str.String = e.Error()
	}
	tx, err := ms.db.Beginx()
	if err != nil {
		return err
	}

	record := State{
		State:           int(state),
		Event:           ctx.Event,
		Data:            ctx.Storage,
		ContextID:       ctx.ID,
		CreatedAt:       time.Now(),
		ProcessingError: str,
	}
	err = record.Insert(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, alias := range ctx.Aliases {
		al := Alias{
			Alias:     alias,
			ContextID: ctx.ID,
		}
		err = al.Insert(tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (ms *DbStorage) Last(id string) (*machine.IncompleteStateContext, error) {
	item, err := LastState(ms.db, id, id)
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

func NewDbStorage(db *sqlx.DB) (*DbStorage, error) {
	_, err := db.Exec(string(MustAsset("init.sql")))
	if err != nil {
		return nil, errors.Wrap(err, "db storage - initialize db structure")
	}
	return &DbStorage{
		db: db,
	}, nil
}
