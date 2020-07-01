package serializer

import (
	"YouComic-Nano/config"
	"YouComic-Nano/datasource"
	"YouComic-Nano/entity"
	"YouComic-Nano/utils"
	"fmt"
	"github.com/jinzhu/copier"
	"os"
	"path"
	"path/filepath"
	"time"
)

type BasePageTemplate struct {
	CreatedAt time.Time `json:"created_at"`
	Order     int       `json:"order"`
	Path      string    `json:"path"`
}
type PageTemplateWithSize struct {
	CreatedAt time.Time `json:"created_at"`
	Order     int       `json:"order"`
	Path      string    `json:"path"`
	Width     int       `json:"width"`
	Height    int       `json:"height"`
}

func (t *BasePageTemplate) Serializer(dataModel interface{}, context map[string]interface{}) error {
	var err error
	serializerModel := dataModel.(entity.PageFile)
	err = copier.Copy(t, serializerModel)
	bookIdInterface := context["bookId"]
	bookId := bookIdInterface.(int)
	t.Path = fmt.Sprintf("/content/book/%d/%s?t=%d", bookId, serializerModel.Path, time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}

func (t *PageTemplateWithSize) Serializer(dataModel interface{}, context map[string]interface{}) error {
	var err error
	serializerModel := dataModel.(entity.PageFile)
	err = copier.Copy(t, serializerModel)
	bookIdInterface := context["bookId"]
	bookId := bookIdInterface.(int)
	t.Path = fmt.Sprintf("%s/?t=%d", path.Join("/content/book", fmt.Sprintf("%d", bookId), serializerModel.Path), time.Now().Unix())

	book, err := datasource.GetBookById(bookId)
	if err != nil {
		return err
	}



	filePath := filepath.Join(config.RootPath, book.Path, serializerModel.Path)
	if _, err := os.Stat(filePath); err == nil {
		width, height, _ := utils.GetImageDimension(filePath)
		t.Width = width
		t.Height = height
	} else {
		return err
	}

	return nil
}
