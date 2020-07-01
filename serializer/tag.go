package serializer

import (
	"github.com/jinzhu/copier"
)

type BaseTagTemplate struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func (t *BaseTagTemplate) Serializer(model interface{}, context map[string]interface{}) error {
	var err error
	err = copier.Copy(t, model)
	if err != nil {
		return err
	}
	return nil
}
