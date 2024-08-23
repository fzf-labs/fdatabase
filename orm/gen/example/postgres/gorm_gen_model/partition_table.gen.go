// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gorm_gen_model

import (
	"time"
)

const TableNamePartitionTable = "partition_table"

// PartitionTable mapped from table <partition_table>
type PartitionTable struct {
	ID        string    `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID    string    `gorm:"column:user_id;type:uuid;not null" json:"userId"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp with time zone;not null" json:"createdAt"`
}

// TableName PartitionTable's table name
func (*PartitionTable) TableName() string {
	return TableNamePartitionTable
}
