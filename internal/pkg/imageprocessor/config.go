package imageprocessor

const UPLOADS_ROOT = "./uploads"

type Config struct {
	Dir string `koanf:"dir"`
	MaxSizeMB int64 `koanf:"max_size_mb"`
	MaxDimension int `koanf:"max_dimension"`
}