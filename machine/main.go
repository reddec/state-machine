package machine

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"os"
)

type State int
type SmartData []byte

func (sd *SmartData) GetJSON(dest interface{}) error {
	return json.Unmarshal(*sd, dest)
}

func (sd *SmartData) SetJSON(obj interface{}) error {
	v, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return err
	}
	*sd = v
	return nil
}

func (sd *SmartData) MustSetJSON(obj interface{}) {
	err := sd.SetJSON(obj)
	if err != nil {
		panic(err)
	}
}

func (sd *SmartData) MustGetJSON(obj interface{}) {
	err := sd.GetJSON(obj)
	if err != nil {
		panic(err)
	}
}

type IncompleteStateContext struct {
	ID      string
	Storage SmartData
	Current State
}

type StateContext struct {
	IncompleteStateContext
	Event SmartData
}

type StateHandler func(ctx context.Context, state *StateContext) (State, error)

type StateStorage interface {
	Append(stateContext *StateContext, state State, processingErr error) error
	Last(id string) (*IncompleteStateContext, error)
}

type StateMachine struct {
	state          map[State]StateHandler
	init           StateHandler
	storage        StateStorage
	defaultHandler StateHandler
}

func New(storage StateStorage, init StateHandler) *StateMachine {
	return &StateMachine{
		storage: storage,
		state:   make(map[State]StateHandler),
		init:    init,
	}
}

func (st *StateMachine) WithStorage(storage StateStorage) *StateMachine {
	st.storage = storage
	return st
}

func (st *StateMachine) Register(state State, handler StateHandler) *StateMachine {
	if st.state == nil {
		st.state = make(map[State]StateHandler)
	}
	st.state[state] = handler
	return st
}

func (st *StateMachine) WithDefaultHandler(handler StateHandler) *StateMachine {
	st.defaultHandler = handler
	return st
}

func (st *StateMachine) EmitString(ctx context.Context, contextID string, event string) error {
	return st.Emit(ctx, contextID, []byte(event))
}

func (st *StateMachine) Emit(ctx context.Context, contextID string, event []byte) error {
	prevStateContext, err := st.storage.Last(contextID)

	var stateHandler StateHandler

	if err == os.ErrNotExist {
		stateHandler = st.init
	} else if err != nil {
		return errors.Wrapf(err, "state-machine: fetch context by with id '%v'", contextID)
	} else if handler, ok := st.state[prevStateContext.Current]; ok {
		stateHandler = handler
	} else if st.defaultHandler != nil {
		stateHandler = st.defaultHandler
	} else {
		return errors.Errorf("state-machine: state '%v' unknown and no default state handler defined", prevStateContext.Current)
	}

	if prevStateContext == nil {
		prevStateContext = &IncompleteStateContext{ID: contextID}
	}

	stateContext := &StateContext{
		Event: event,
		IncompleteStateContext: *prevStateContext,
	}

	nextState, err := stateHandler(ctx, stateContext)
	err = st.storage.Append(stateContext, nextState, err)
	if err != nil {
		return errors.Wrapf(err, "state-machine: %v -> %v: save state context", stateContext.Current, nextState)
	}
	if err != nil {
		return errors.Wrapf(err, "state-machine: process handle for state '%v'", stateContext.Current)
	}
	return nil
}
