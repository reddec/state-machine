package storage

import (
	"database/sql"
	"github.com/pkg/errors"
	"github.com/reddec/state-machine/machine"
	"os"
	"strconv"
	"strings"
	"time"
)

type DbStorage struct {
	db XODB
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
	rs, err := ms.db.Query(queryBase, params...)
	if err == sql.ErrNoRows {
		return 0, nil
	} else if err != nil {
		return -1, err
	}
	defer rs.Close()
	err = rs.Scan(&count)
	return count, err
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
