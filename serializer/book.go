package serializer

import (
	"YouComic-Nano/entity"
	"fmt"
	"github.com/jinzhu/copier"
	"time"
)

type BaseBookTemplate struct {
	ID    uint        `json:"id"`
	Name  string      `json:"name"`
	Cover string      `json:"cover"`
	Tags  interface{} `json:"tags"`
}

func (b *BaseBookTemplate) Serializer(dataModel interface{}, context map[string]interface{}) error {
	serializerModel := dataModel.(entity.Book)
	err := copier.Copy(b, serializerModel)
	if err != nil {
		return err
	}
	if len(b.Cover) != 0 {
		b.Cover = fmt.Sprintf("/content/book/%d%s?t=%d", serializerModel.ID, serializerModel.Cover, time.Now().Unix())
	}
	tags := b.Tags
	serializedTags := SerializeMultipleTemplate(tags, &BaseTagTemplate{}, nil)
	b.Tags = serializedTags
	return nil
}
