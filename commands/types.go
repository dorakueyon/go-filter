package commands

import "regexp"

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
	ALLREGEXREPLACE
)

type filterConfig struct {
	filterType   filterType
	from         string
	to           string
	regexPattern *regexp.Regexp
}

type flags struct {
	debug bool
}
