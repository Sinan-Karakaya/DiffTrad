package internal

import (
	"DiffTrad/internal/config"
	"bufio"
	"fmt"
	"os"

	"github.com/xyproto/ollamaclient"
)

const prompt = `
Translate the following text in %s.
You must NOT include any text before or after the translation.

%s
`

func GenerateTranslations(c *config.Config) error {
	oc := ollamaclient.NewWithModelAndAddr(c.AI_Model, c.AI_URL)
	oc.Verbose = c.Verbose

	if err := oc.PullIfNeeded(); err != nil {
		return fmt.Errorf("could not pull the model: %s", err.Error())
	}

	writer := bufio.NewWriter(os.Stdout)
	count := 0
	for idx, lang := range c.TargetLanguages {
		if c.Verbose {
			writer.Flush()
			fmt.Printf("Progression: %d/%d", count, len(c.TargetLanguages))
		}

		formattedPrompt := fmt.Sprintf(prompt, c.RefLang, lang.Content)
		translation, err := oc.GetOutput(formattedPrompt)
		if err != nil {
			return fmt.Errorf("could not get output: %s", err.Error())
		}

		c.TargetLanguages[idx].Translation = translation
	}

	return nil
}
