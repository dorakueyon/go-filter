package commands

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

func loadFilterConfig() (config, error) {
	fltCfgCsvFilePath := "config.csv"
	csvLines, err := readCfgCsv(fltCfgCsvFilePath)
	if err != nil {
		return config{}, err
	}
	var fltCfgs []filterConfig
	for _, row := range csvLines {
		var fltCfg filterConfig
		if strings.HasPrefix(row[0], "//") {
			continue
		}

		switch row[0] {
		case "allReplace":
			fltCfg.filterType = ALLREPLACE
			fltCfg.from = row[1]
			fltCfg.to = row[2]
		case "prefixReplace":
			fltCfg.filterType = PREFIXREPLACE
			fltCfg.from = row[1]
			fltCfg.to = row[2]
		case "allRegexReplace":
			fltCfg.filterType = ALLREGEXREPLACE
			patternStr := row[1]
			var wordReg = regexp.MustCompile(patternStr)
			fltCfg.regexPattern = wordReg
			fltCfg.to = row[2]
		default:
			continue
		}
		fltCfgs = append(fltCfgs, fltCfg)
	}

	return config{
		filterConfigs: fltCfgs,
	}, nil
}

func parseFlags() (flags, error) {
	var (
		debug = flag.Bool("d", false, "debug mode")
	)
	flag.Parse()
	return flags{
		debug: *debug,
	}, nil
}

// NewApp is ...
func NewApp() (*App, error) {
	cfg, err := loadFilterConfig()
	if err != nil {
		return nil, err
	}
	flgs, err := parseFlags()
	if err != nil {
		return nil, err
	}

	return &App{config: cfg, flags: flgs}, nil
}

func (app *App) execute() error {
	for i, processFile := range app.processFiles {
		rowLines, err := getFileLines(processFile.filename)
		if err != nil {
			panic(err)
		}
		processFile.rowLines = rowLines
		processFile.bufLines = rowLines
		for _, filterCfg := range app.config.filterConfigs {
			switch filterCfg.filterType {
			case ALLREPLACE:
				from := filterCfg.from
				to := filterCfg.to
				newLines := filterReplaceStrings(processFile.bufLines, from, to)
				processFile.bufLines = newLines
				processFile.newLines = newLines

			case PREFIXREPLACE:
				from := filterCfg.from
				to := filterCfg.to
				newLines := filterReplacePrefixString(processFile.bufLines, from, to)
				processFile.bufLines = newLines
				processFile.newLines = newLines

			case ALLREGEXREPLACE:
				regexPattern := filterCfg.regexPattern
				to := filterCfg.to
				newLines := filterAllRegexReplace(processFile.bufLines, regexPattern, to)
				processFile.bufLines = newLines
				processFile.newLines = newLines

			default:
				fmt.Println("default")
			}
		}
		app.processFiles[i] = processFile
	}

	return nil
}

func (fcf *filterConfig) displayFilterType() string {
	switch fcf.filterType {
	case ALLREPLACE:
		return "allDisplay"
	case PREFIXREPLACE:
		return "prefixDisplay"
	case ALLREGEXREPLACE:
		return "allRegexReplace"
	default:
		return "not defined"
	}
}

func (app *App) debug() {

	for _, fltCfg := range app.config.filterConfigs {
		fmt.Printf("type=%s, from=%s, to=%s, regexPattern=%s\n",
			fltCfg.displayFilterType(),
			fltCfg.from,
			fltCfg.to,
			fltCfg.regexPattern,
		)
	}

	for _, processFile := range app.processFiles {
		fmt.Printf("FILENAME:%s\n", processFile.filename)
		fmt.Println("===============")
		for _, line := range processFile.newLines {
			fmt.Println(line)
		}
		fmt.Println("===============")
	}

}

func (app *App) createOutputFiles(dirName string) {
	for _, processFile := range app.processFiles {

		name := filepath.Join(dirName, filepath.Base(processFile.filename))
		createFile(
			name,
			processFile.newLines,
		)

	}

}

func (app *App) getFileNames(dirName string) error {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		return err
	}
	var processFiles []processFile
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		processFiles = append(processFiles, processFile{
			filename: filepath.Join(dirName, file.Name()),
		})
	}
	app.processFiles = processFiles
	return nil
}

// Run is ...
func (app *App) Run() error {
	inputDirName := "input"
	outputDirName := "output"
	err := app.getFileNames(inputDirName)
	if err != nil {
		return err
	}
	err = app.execute()
	if err != nil {
		return err
	}
	if app.flags.debug {
		app.debug()
	} else {
		app.createOutputFiles(outputDirName)
	}

	return nil
}
