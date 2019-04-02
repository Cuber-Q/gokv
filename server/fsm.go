package server

import (
	"log"
	"encoding/json"
	"io"
	"github.com/hashicorp/raft"
)

// need to implement raft.fsm interface to build our own Finite State Machine
type FSM struct {
	Ctx *Context
	Log *log.Logger
}

type logEntryData struct {
	Key   string
	Value string
}

// Apply applies a Raft log entry to the key-value store.
func (f *FSM) Apply(logEntry *raft.Log) interface{} {
	e := logEntryData{}
	if err := json.Unmarshal(logEntry.Data, &e); err != nil {
		panic("Failed unmarshaling Raft log entry. This is a bug.")
	}
	//ret := f.ctx.st.cm.Set(e.Key, e.Value)
	//f.log.Printf("fms.Apply(), logEntry:%s, ret:%v\n", logEntry.Data, ret)
	return true
}

// Snapshot returns a latest snapshot
func (f *FSM) Snapshot() (raft.FSMSnapshot, error) {
	//return &snapshot{cm: f.ctx.st.cm}, nil
	return nil, nil
}

// Restore stores the key-value store to a previous state.
func (f *FSM) Restore(serialized io.ReadCloser) error {
	//return f.ctx.st.cm.UnMarshal(serialized)
	return nil
}
