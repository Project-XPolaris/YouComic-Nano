package datasource

import (
	"YouComic-Nano/config"
	"YouComic-Nano/entity"
	"strings"
)

func InitDataSource() {
	InitTagDataSource()
	for idx := range config.Current.Books {
		config.Current.Books[idx].ID = idx + 1
	}
}

func GetBookById(id int) (*entity.Book, error) {
	for _, book := range config.Current.Books {
		if book.ID == id {
			return &book, nil
		}
	}
	return nil, nil
}

func GetBookCoverCacheById(id int) (string, error) {
	book, err := GetBookById(id)
	if err != nil {
		return "", err
	}
	rawThumbnails := config.CacheConfig.GetStringMapString("thumbnail")
	return rawThumbnails[book.UUID], nil
}

type BookReader struct {
	Page     int
	PageSize int
	Filters  []DataFilter
}

func (b *BookReader) Read() (interface{}, int, error) {
	if b.Filters == nil {
		b.Filters = make([]DataFilter, 0)
	}
	source := config.Current.Books
	for _, dataFilter := range b.Filters {
		result := make([]entity.Book, 0)
		for _, book := range source {
			ok, err := dataFilter.filter(book)
			if err != nil {
				return make([]entity.Tag, 0), 0, err
			}
			if ok {
				result = append(result, book)
			}
		}
		source = result
	}
	start := (b.Page - 1) * b.PageSize
	if start >= len(source) {
		return make([]entity.Book, 0), 0, nil
	}
	end := b.Page * b.PageSize
	if end > len(source) {
		end = len(source)
	}
	result := source[start:end]
	return result, len(source), nil
}

type BookTagFilter struct {
	TagIds []int
}

func (b *BookTagFilter) filter(data interface{}) (bool, error) {
	book := data.(entity.Book)
	for _, bookTag := range book.Tags {
		for _, tagId := range b.TagIds {
			if bookTag.ID == tagId {
				return true, nil
			}
		}
	}
	return false, nil
}

type BookNameSearchFilter struct {
	Key string
}

func (b *BookNameSearchFilter) filter(data interface{}) (bool, error) {
	return strings.Contains((data.(entity.Book)).Name, b.Key), nil
}
