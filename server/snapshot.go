package server

import (
	"github.com/hashicorp/raft"
	"gokv/core"
)

type snapshot struct {
	storage *core.Storage
}

// Persist saves the FSM snapshot out to the given sink.
func (s *snapshot) Persist(sink raft.SnapshotSink) error {
	snapshotBytes, err := s.storage.Marshal()
	if err != nil {
		sink.Cancel()
		return err
	}

	if _, err := sink.Write(snapshotBytes); err != nil {
		sink.Cancel()
		return err
	}

	if err := sink.Close(); err != nil {
		sink.Cancel()
		return err
	}
	return nil
}

func (f *snapshot) Release() {}
