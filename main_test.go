package state_machine

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/reddec/state-machine/machine"
	"github.com/reddec/state-machine/storage"
	"strconv"
	"testing"
	"time"
)

const (
	OneValueAdded  = 1
	ResultComplete = 2
	InitFailed     = 3
)

func TestStateMachine_memstorage(t *testing.T) {
	st := machine.New(&storage.MemStorage{}, func(ctx context.Context, state *machine.StateContext) (machine.State, error) {
		var val int
		err := state.Event.GetJSON(&val)
		if err != nil {
			return InitFailed, err
		}
		state.Storage.MustSetJSON(val)
		return OneValueAdded, nil
	})
	baseTest(t, st)
}

func TestStateMachine_dbstorage(t *testing.T) {
	db, err := sqlx.Connect("postgres", "host=localhost dbname=postgres sslmode=disable user=postgres")
	if err != nil {
		t.Error("connect to db:", err)
		return
	}
	defer db.Close()
	stor, err := storage.NewDbStorage(db)
	if err != nil {
		t.Error("initialize db storage:", err)
		return
	}
	st := machine.New(stor, func(ctx context.Context, state *machine.StateContext) (machine.State, error) {
		var val int
		err := state.Event.GetJSON(&val)
		if err != nil {
			return InitFailed, err
		}
		state.Storage.MustSetJSON(val)
		return OneValueAdded, nil
	})
	baseTest(t, st)
}

func baseTest(t *testing.T, st *machine.StateMachine) {

	st.Register(OneValueAdded, func(ctx context.Context, state *machine.StateContext) (machine.State, error) {
		var val int
		err := state.Event.GetJSON(&val)
		if err != nil {
			return state.Current, err
		}
		var prev int
		state.Storage.MustGetJSON(&prev)

		res := prev + val
		fmt.Println("result:", prev+val)
		state.Storage.MustSetJSON(res)
		return ResultComplete, nil
	})

	id := "test-" + strconv.FormatInt(time.Now().Unix(), 10)

	ctx := context.Background()
	if err := st.EmitString(ctx, id, "1"); err != nil {
		t.Error("failed emit 1:", err)
		return
	}
	if err := st.EmitString(ctx, id, "2"); err != nil {
		t.Error("failed emit 2:", err)
		return
	}
	if err := st.EmitString(ctx, id, "3"); err != nil {
		t.Error("failed emit 3:", err)
		return
	}
}
