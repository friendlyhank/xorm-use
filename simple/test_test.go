package test

import (
	"github.com/friendlyhank/xorm-use/foundation/db"
	"github.com/xormplus/xorm"
	"testing"
)

func TestFindAndCount(t *testing.T){
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	db.Engine().MapCacher(&db.Store{}, cacher)

	var store =&db.Store{}
	has,err := db.Engine().ID(1).Get(store)

	if err != nil || !has{
		t.Error("%v",err)
		return
	}

	t.Logf("%v",store)
}



