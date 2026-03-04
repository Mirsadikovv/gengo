package {{ toSnake .EntityName }}_model

import "time"

type {{ toCamel .EntityName }} struct {
	Id        int64      `json:"id"        gorm:"primaryKey"`
	Name      string     `json:"name"      gorm:"not null"`
	CreatedAt *time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"softDelete"`
}

func ({{ toCamel .EntityName }}) TableName() string {
	return "{{ toSnake .EntityName }}s"
}
