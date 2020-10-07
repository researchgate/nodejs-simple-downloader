package nodejs

// DownloadURL has two values either the stable url or the url for unofficial builds
type DownloadURL string

// Stable and Unofficial are the two values for DownloadURL
const (
	Stable     DownloadURL = "https://nodejs.org/dist/v%s/node-v%s-%s.%s"
	Unofficial DownloadURL = "https://unofficial-builds.nodejs.org/download/release/v%s/node-v%s-%s.%s"
)
