package config

type Config interface {
	GetSiteData() SiteData
	GetPhotos() []Photo
}

type SiteData struct {
	Name     string
	Author   string
	Root     string
	Output   string
	Template string
}

type Photo struct {
	File   string
	Name   string
	Resize bool
}

func FilmConfig() Config {
	return NewFileConfig("film.yaml")
}
