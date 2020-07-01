package entity

type PageFile struct {
	Path  string `json:"path"`
	Order int
}
type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
type Book struct {
	ID    int        `json:"id"`
	UUID  string     `json:"uuid"`
	Name  string     `json:"name"`
	Path  string     `json:"path"`
	Cover string     `json:"cover"`
	Pages []PageFile `json:"pages"`
	Tags  []Tag      `json:"tags"`
}
type CoverThumbnail struct {
	SourcePath string `json:"source_path"`
	Name       string `json:"name"`
}
