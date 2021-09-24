package conf

type Repository struct {
	Workspace string    `yaml:"workspace"`
	Projects  []Project `yaml:"projects"`
}

type Project struct {
	Name     string    `yaml:"name"`
	Url      string    `yaml:"url"`
	Versions []Version `yaml:"versions"`
}

type Version struct {
	Tag    string `yaml:"tag"`
	Digest string `yaml:"digest"`
}
