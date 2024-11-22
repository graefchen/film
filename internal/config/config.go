package config

type Config interface {
	GetSiteData() SiteData
	GetTemplateData() TemplateData
	GetPhotos() []Photo
}

type SiteData struct {
	Name      string
	Author    string
	Root      string
	OutputDir string
}

type TemplateData struct {
	Archive string
	Site    string
}

type Photo struct {
	File   string
	Name   string
	Resize bool
}

func FilmConfig() Config {
	return NewFileConfig("film.yaml")
}
