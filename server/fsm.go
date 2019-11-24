package server

import (
	"encoding/json"
	"github.com/hashicorp/raft"
	"gokv/core"
	"io"
	"log"
)

// need to implement raftNode.fsm interface to build our own Finite State Machine
type FSM struct {
	Id      string
	Ctx     *Context
	Log     *log.Logger
	Storage *core.Storage
}

type logEntryData struct {
	K string
	V string
}

// Apply applies a Raft log entry to the key-value store.
func (f *FSM) Apply(logEntry *raft.Log) interface{} {
	log.Printf("%v recive log and ready to APPLY: %v ", f.Id, logEntry)
	e := logEntryData{}
	if err := json.Unmarshal(logEntry.Data, &e); err != nil {
		panic("Failed Unmarshal Raft log entry. This is a bug.")
	}
	//ret := f.ctx.st.cm.Set(e.K, e.V)
	//f.log.Printf("fms.Apply(), logEntry:%s, ret:%v\n", logEntry.Data, ret)
	//f.Storage.SetV2(e.K, e.V)
	cmd := &core.StoreCMD{
		K:      e.K,
		V:      e.V,
		T:      core.SET,
		RespCh: make(chan string, 1),
	}
	f.Storage.SetV3(cmd)
	if r := <-cmd.RespCh; r != "" {
		return true
	}
	return false
}

// Snapshot returns a latest snapshot
func (f *FSM) Snapshot() (raft.FSMSnapshot, error) {
	return &snapshot{storage: f.Storage}, nil
}

// Restore stores the key-value store to a previous state.
func (f *FSM) Restore(serialized io.ReadCloser) error {
	return f.Storage.UnMarshal(serialized)
}
