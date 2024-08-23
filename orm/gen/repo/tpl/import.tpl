import (
    "context"
	"encoding/json"
	"errors"
    "{{.daoPkgPath}}"
    "{{.modelPkgPath}}"
	"gitlab.yc345.tv/backend/utils/v2/orm"
	"gitlab.yc345.tv/backend/utils/v2/orm/gen/cache"
	"gitlab.yc345.tv/backend/utils/v2/orm/gen/condition"
	"gitlab.yc345.tv/backend/utils/v2/orm/gen/custom"
	"gitlab.yc345.tv/backend/utils/v2/orm/gen/encoding"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)
