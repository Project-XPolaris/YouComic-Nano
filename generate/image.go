package generate

import (
	"YouComic-Nano/config"
	"YouComic-Nano/entity"
	"fmt"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
	"github.com/sirupsen/logrus"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func GenerateLibraryCoverThumbnail(books []entity.Book) error {
	storeRoot := filepath.Join(config.RootPath, ".cache", "thumbnail")
	err := os.MkdirAll(storeRoot, os.ModePerm)

	cacheCoverItems := make(map[string]string, 0)
	if err != nil {
		return err
	}
	for idx, book := range books {
		logrus.Debug(fmt.Sprintf("generate book cover thumbnail %d of %d", idx, len(books)))
		sourcePath := filepath.Join(config.RootPath, book.Path, book.Cover)
		coverExt := filepath.Ext(book.Cover)
		fileUUID, err := uuid.NewUUID()
		coverThumbnailFileName := fmt.Sprintf("%s%s", fileUUID.String(), coverExt)
		if err != nil {
			return err
		}
		_, err = GenerateCoverThumbnail(sourcePath, filepath.Join(storeRoot, coverThumbnailFileName))
		if err != nil {
			return err
		}
		cacheCoverItems[book.UUID] = filepath.Base(coverThumbnailFileName)
	}
	config.CacheConfig.Set("thumbnail", cacheCoverItems)
	err = config.CacheConfig.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

//generate thumbnail image
func GenerateCoverThumbnail(coverImageFilePath string, storePath string) (string, error) {
	// setup image decoder
	fileExt := filepath.Ext(coverImageFilePath)
	thumbnailImageFile, err := os.Open(coverImageFilePath)
	if err != nil {
		return "", err
	}
	var thumbnailImage image.Image
	if strings.ToLower(fileExt) == ".png" {
		thumbnailImage, err = png.Decode(thumbnailImageFile)
	}
	if strings.ToLower(fileExt) == ".jpg" {
		thumbnailImage, err = jpeg.Decode(thumbnailImageFile)
	}
	if err != nil {
		return "", err
	}

	// make thumbnail
	resizeImage := resize.Thumbnail(480, 480, thumbnailImage, resize.Lanczos3)

	output, err := os.Create(storePath)
	if err != nil {
		return "", err
	}

	defer thumbnailImageFile.Close()
	defer output.Close()

	// save result
	err = jpeg.Encode(output, resizeImage, nil)
	if err != nil {
		return "", err
	}
	return storePath, nil
}
