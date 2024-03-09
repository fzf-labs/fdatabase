import (
    "context"
	"encoding/json"
	"errors"

    "github.com/fzf-labs/fdatabase/orm/custom"
    "github.com/fzf-labs/fdatabase/orm/gen/cache"
    "{{.daoPkgPath}}"
    "{{.modelPkgPath}}"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)
