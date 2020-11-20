package nodejs

// FileExtension defines all possible file extensions for nodejs downloads
type FileExtension string

// Stable and Unofficial are the two values for DownloadURL
const (
	TarGz FileExtension = "tar.gz"
	TarXz FileExtension = "tar.xz"
	Zip   FileExtension = "zip"
)
