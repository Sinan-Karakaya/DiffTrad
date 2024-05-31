package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"DiffTrad/internal"
	"DiffTrad/internal/config"
)

func optParse() config.Config {
	c := config.Config{}

	flag.BoolVar(&c.Verbose, "verbose", false, "Verbose output")
	flag.StringVar(&c.OutputPath, "output", "output.html", "Output path")

	flag.StringVar(&c.TargetsPath, "targets", "", "Path to the target files")
	flag.StringVar(&c.RefPath, "ref", "", "Path to the reference file")
	flag.StringVar(&c.RefLang, "ref-lang", "french", "Regular language")

	flag.StringVar(&c.AI_URL, "ai-url", "http://localhost:11434", "URL to the AI service")
	flag.StringVar(&c.AI_Model, "ai-model", "mistral", "Model to use")
	flag.StringVar(&c.AI_Auth, "ai-auth", "", "Authorization token (leave empty if not needed)")

	flag.Parse()

	dirPath, filename := filepath.Split(c.OutputPath)
	if dirPath == "" {
		dirPath = "./"
	}
	c.OutputDirPath = dirPath
	c.OutputFilename = filename

	return c
}

func main() {
	c := optParse()

	err := internal.Run(c)
	if err != nil {
		fmt.Println("An error happened: "+err.Error())
		os.Exit(1)
	}
}
