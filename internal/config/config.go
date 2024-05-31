package config

type Config struct {
	Verbose    bool
	OutputPath string

	TargetsPath string
	RefPath     string
	RefLang     string

	AI_URL   string
	AI_Model string
	AI_Auth  string

	// not cmd flags
	ReferenceLanguage Language
	TargetLanguages   []Language

	OutputDirPath string
	OutputFilename string
}

type Language struct {
	Content string
	Translation string

	Filename string
	Path     string
	FullPath string
}
