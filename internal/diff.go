package internal

import (
	"DiffTrad/internal/config"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"text/template"
)

func WriteDiffs(c config.Config) error {
	compiledPath := fmt.Sprintf("%s/%s", c.OutputDirPath, "diffsCompiled")
	_, err := os.Create(compiledPath)
	if err != nil {
		return fmt.Errorf("could not create the compiled diffs file: %s", err.Error())
	}

	for _, lang := range c.TargetLanguages {
		path := fmt.Sprintf("/tmp/%s", lang.Filename)
		err := os.WriteFile(path, []byte(lang.Translation), 0644)
		if err != nil {
			return fmt.Errorf("could not write the translation to a file: %s", err.Error())
		}

		cmd := exec.Command("diff", "-u", path, c.RefPath)
		output, _ := cmd.CombinedOutput()		// we ignore the error because of how diff handles returns codes
		output = bytes.ReplaceAll(output, []byte("\\ No newline at end of file"), []byte(""))

		f, err := os.OpenFile(compiledPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("could not open the compiled diffs file: %s", err.Error())
		}
		defer f.Close()

		_, err = f.WriteString(string(output)+"\n")
		if err != nil {
			return fmt.Errorf("could not write the diff output to the compiled diffs file: %s", err.Error())
		}
	}
	
	return nil
}

func CompileUI(c config.Config) error {
	uiPath := fmt.Sprintf("%s/%s", c.OutputDirPath, c.OutputFilename)

	err := os.MkdirAll(c.OutputDirPath, 0755)
	if err != nil {
		return fmt.Errorf("could not create the output directory: %s", err.Error())
	}

	_, err = os.Create(uiPath)
	if err != nil {
		return fmt.Errorf("could not create the UI file: %s", err.Error())
	}

	exec, err := os.Executable()
	if err != nil {
		return fmt.Errorf("could not get the executable path: %s", err.Error())
	}

	templatePath := fmt.Sprintf("%s/ui/template.html", filepath.Dir(exec))
	uiTemplateRaw, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("could not read the UI template: %s", err.Error())
	}

	tmpl, err := template.New("template").Parse(string(uiTemplateRaw))
	if err != nil {
		return fmt.Errorf("could not parse the UI template: %s", err.Error())
	}

	diffCompiled, err := os.ReadFile(fmt.Sprintf("%s/%s", c.OutputDirPath, "diffsCompiled"))
	if err != nil {
		return fmt.Errorf("could not read the compiled diffs: %s", err.Error())
	}

	var uiCompiled bytes.Buffer
	err = tmpl.Execute(&uiCompiled, compiledTemplate{DiffCompiled: string(diffCompiled)})
	if err != nil {
		return fmt.Errorf("could not execute the UI template: %s", err.Error())
	}

	err = os.WriteFile(uiPath, uiCompiled.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("could not write the UI file: %s", err.Error())
	}

	Open(uiPath)
	
	return nil
}

type compiledTemplate struct {
	DiffCompiled string
}

func Open(path string) {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open", path}
	case "windows":
		args = []string{"cmd", "/c", "start", path}
	default:
		args = []string{"xdg-open", path}
	}

	cmd := exec.Command(args[0], args[1:]...)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to open the output file: "+path)
	}
}