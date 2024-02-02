import (
    "context"
	"encoding/json"
	"errors"

    "github.com/fzf-labs/fdatabase/custom"
    "github.com/fzf-labs/fdatabase/gen/cache"
    "{{.daoPkgPath}}"
    "{{.modelPkgPath}}"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)
