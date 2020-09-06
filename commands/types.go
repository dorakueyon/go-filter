package commands

// App is ...
type App struct {
	config       config
	flags        flags
	processFiles []processFile
}

type processFile struct {
	filename string
	rowLines []string
	bufLines []string
	newLines []string
}

type config struct {
	filterConfigs []filterConfig
}

type filterType int

const (
	_ filterType = iota
	ALLREPLACE
	PREFIXREPLACE
)

type filterConfig struct {
	filterType filterType
	from       string
	to         string
}

type flags struct {
	debug bool
}
