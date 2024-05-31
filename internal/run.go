package internal

import (
	"DiffTrad/internal/config"
)

func Run(c config.Config) error {
	err := RetrieveFiles(&c)
	if err != nil {
		return err
	}

	err = GenerateTranslations(&c)
	if err != nil {
		return err
	}

	err = WriteDiffs(c)
	if err != nil {
		return err
	}

	err = CompileUI(c)
	if err != nil {
		return err
	}
	
	return nil
}