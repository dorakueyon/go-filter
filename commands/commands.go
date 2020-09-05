package commands

import (
	"fmt"
	"path/filepath"
)

func newConfig() (config, error) {
	fltCfgs := []filterConfig{

		filterConfig{
			filterType: allReplace,
			from:       "#",
			to:         "yooooooooooooooooo",
		},
		filterConfig{
			filterType: prefixReplace,
			from:       "#",
			to:         "to",
		},
	}

	return config{
		filterConfigs: fltCfgs,
	}, nil
}

// NewApp is ...
func NewApp() (*App, error) {
	cfg, err := newConfig()
	if err != nil {
		return nil, err
	}
	return &App{config: cfg}, nil
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
			case allReplace:
				from := filterCfg.from
				to := filterCfg.to
				newLines := filterReplaceStrings(processFile.bufLines, from, to)
				processFile.bufLines = newLines
				processFile.newLines = newLines

			case prefixReplace:
				from := filterCfg.from
				to := filterCfg.to
				newLines := filterReplacePrefixString(processFile.bufLines, from, to)
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

func (app *App) debug() {
	for _, processFile := range app.processFiles {
		fmt.Printf("FILENAME:%s\n", processFile.filename)
		fmt.Println("===============")
		for _, line := range processFile.newLines {
			fmt.Println(line)
		}
		fmt.Println("===============")
	}

}

func (app *App) createOutputFiles() {
	for _, processFile := range app.processFiles {

		outFileName := fmt.Sprintf("./output/%s", filepath.Base(processFile.filename))
		createFile(
			outFileName,
			processFile.newLines,
		)

	}

}

// Run is ...
func (app *App) Run(debug bool) error {
	filename := "./input/hoge.md"
	processFile := processFile{
		filename: filename,
	}
	app.processFiles = append(app.processFiles, processFile)

	err := app.execute()
	if err != nil {
		return err
	}
	app.debug()
	app.createOutputFiles()

	return nil
}
