import (
    "context"
	"errors"
	"github.com/fzf-labs/fdatabase/orm/condition"
	"github.com/fzf-labs/fdatabase/orm/dbcache"
	"github.com/fzf-labs/fdatabase/orm/encoding"
	"github.com/fzf-labs/fdatabase/orm/gen/config"
    "{{.daoPkgPath}}"
    "{{.modelPkgPath}}"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)
