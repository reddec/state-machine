package cmd

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/reddec/state-machine/machine"
	"github.com/reddec/state-machine/storage"
	"testing"
)

func TestEvent(t *testing.T) {
	conn, err := sqlx.Connect("postgres", "host=localhost user=postgres password=postgres sslmode=disable")
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.Close()
	stor, err := storage.NewDbStorage(conn)
	if err != nil {
		t.Error(err)
		return
	}
	vm := machine.New(stor, func(ctx context.Context, state *machine.StateContext) (machine.State, error) {
		state.Storage = []byte("hell")
		return 1, nil
	})
	vm.Register(1, func(ctx context.Context, state *machine.StateContext) (machine.State, error) {
		data := state.Storage
		if string(data) != "hell" {
			t.Error("bad storage")
		}
		state.Aliases = append(state.Aliases, "abd")
		return 2, nil
	})
	alias := false
	vm.Register(2, func(ctx context.Context, state *machine.StateContext) (machine.State, error) {
		fmt.Println("AAAAAA")
		alias = true
		return 3, nil
	})
	err = vm.EmitString(context.Background(), "abc", "aaaaa")
	if err != nil {
		t.Error(err)
		return
	}
	err = vm.EmitString(context.Background(), "abc", "aaaaa")
	if err != nil {
		t.Error(err)
		return
	}

	err = vm.EmitString(context.Background(), "abd", "aaaaa")
	if err != nil {
		t.Error(err)
		return
	}
	if !alias {
		t.Error("alias not working")
	}
}
func TestEvent(t *testing.T) {
	conn, err := sqlx.Connect("postgres", "host=localhost user=postgres password=postgres sslmode=disable")
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.Close()
	stor, err := storage.NewDbStorage(conn)
	if err != nil {
		t.Error(err)
		return
	}
	vm := machine.New(stor, func(ctx context.Context, state *machine.StateContext) (machine.State, error) {
		fmt.Println("hell 1", state.ID)
		return 1, nil
	})

	vm := machine.New()

}
