package launcher

//spoo:generate protoc -I../proto -I ${GOPATH}/src -I ${GOPATH}/src/github.com/golang --go_out=plugins=grpc:. ../proto/fsme.proto

import (
	fsm "github.com/chickenandpork/dsv/gorm"
)

// LaunchStore defines the extended key/value store use to hold jobs that are shelved for later.
// In theory, the backend for a LaunchStore can be any database.  Eventually there will need to be
// a scheduler.
type LaunchStore interface {
	// Checkout is used to atomically get a job and assign it to the client so that other
	// engines don't run it in split-brain parallel
	Checkout(uuid string) (fsm.Context, fsm.MutableData)

	// Commit is used to push back changes in the MutableData at the end of a step/stage run
	Commit(uuid string, mu fsm.MutableData) error

	// Create is used to create the constant Context and initial Mutable data
	Create(ctx fsm.Context, mu fsm.MutableData) error

	// CreateBlankMutable is used to create the constant Context of a FSM without MutableData
	//CreateBlankMutable(ctx fsm.Context) error

	// CreateQuick is used to instantiate a new process with as minimal data as possible, returning process ID (uuid)
	CreateQuick(className, start string) string

	// GetJobs -- as the name implies -- retrieves a list of runnable jobs.  Through locality,
	// the query can be focused on jobs that require an engine running on a particular host, or
	// a cluster, or a datacenter.  Within the limit of a few jobs retrieved, the client should
	// also randomize the results to choose a runnable job.  Allowing the client to avoid being
	// pegged forever behind a specific stuck job, this is stolen from a corruption of the
	// basic RR DNS hack.
	GetJobs(classes []string, locality map[string]string, limit int) (
		jobs []fsm.Context, // returned context on query is "short": no parameters
		err error,
	)
}
