package testkv

import (
	"testing"

	dbm "github.com/tendermint/tm-db"
	"gotest.tools/v3/assert"

	"github.com/blockgenx/blockgen-sdk/orm/model/ormtable"
)

func NewGoLevelDBBackend(t testing.TB) ormtable.Backend {
	db, err := dbm.NewGoLevelDB("test", t.TempDir())
	assert.NilError(t, err)
	return ormtable.NewBackend(ormtable.BackendOptions{
		CommitmentStore: db,
	})
}
