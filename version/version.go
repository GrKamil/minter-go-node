package version

// Version components
const (
	Maj = "1"
	Min = "0"
	Fix = "5"

	AppVer = 5
)

var (
	// Must be a string because scripts like dist.sh read this file.
	Version = "1.0.5"

	// GitCommit is the current HEAD set using ldflags.
	GitCommit string
)

func init() {
	if GitCommit != "" {
		Version += "-" + GitCommit
	}
}
