package profile

import (
	"github.com/kelleygo/clashcore/common/atomic"
)

// StoreSelected is a global switch for storing selected proxy to cache
var StoreSelected = atomic.NewBool(true)
