package core

// interface of core operation
// it could be implement as in-memory operation,
// or by http/binary protocol
type StoreOperation interface {

	// set a k-v in gokv
	Set(k, v string)

	// get value by specific key
	Get(k string) string

	// check is the key exists in gokv
	Exist(k string) bool

	// remove the value of specific key
	Remove(k string) bool
}

type ClusterOperation interface {

	// a newNode join in a running cluster
	// only in cluster mode and only the leader node
	// is able to call this function.
	// parameter 'newNode' should be the new node's cluster address
	// as <ip>:<port>
	Join(newNode string) string
}
