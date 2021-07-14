package nodejs

// CurrentArch and CurrentURL describe how the URL of the nodejs download will look like
const (
	CurrentArch      string        = "darwin-x64" //until nodejs provides a specific image for arm64 we'll use x64
	CurrentURL       DownloadURL   = Stable
	CurrentExtension FileExtension = TarXz
)
