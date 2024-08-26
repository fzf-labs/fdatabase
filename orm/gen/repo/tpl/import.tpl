import (
    "context"
	"encoding/json"
	"errors"
    "{{.daoPkgPath}}"
    "{{.modelPkgPath}}"
	"github.com/fzf-labs/fdatabase/orm"
	"github.com/fzf-labs/fdatabase/orm/dbcache"
	"github.com/fzf-labs/fdatabase/orm/condition"
	"github.com/fzf-labs/fdatabase/orm/encoding"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)
