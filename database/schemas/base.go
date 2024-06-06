package schemas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID `json:"id" gorm:"column:id;type:uuid;primarykey"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp"`
	IsActive  bool      `json:"is_active" gorm:"column:is_active;default:false"`
}

// BeforeCreate - generates a new UUID for the persisted entity when not provided and a timeNow for created_at
func (i *Base) BeforeCreate(tx *gorm.DB) (err error) {
	if i.ID == uuid.Nil {
		i.ID, _ = uuid.NewRandom()
	}

	i.CreatedAt = time.Now()

	return
}

// BeforeUpdate - sets a timeNow for updated_at
func (i *Base) BeforeUpdate(tx *gorm.DB) (err error) {
	i.UpdatedAt = time.Now()

	return
}

func (i *Base) BaseSortableFields() []string {
	return []string{"id", "created_at", "updated_at", "is_active"}
}
