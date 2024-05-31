# DiffTrad

A Golang tool that translates a list of files to a given language, and generate a diff of the original language file, and the translated one to check for semantic differences.

## How to use

```
Usage of ./DiffTrad:
  -ai-auth string
        Authorization token (leave empty if not needed)
  -ai-model string
        Model to use (default "mistral")
  -ai-url string
        URL to the AI service (default "http://localhost:11434")
  -output string
        Output path (default "output.html")
  -ref string
        Path to the reference file
  -ref-lang string
        Regular language (default "french")
  -targets string
        Path to the target files
  -verbose
        Verbose output
```

`./DiffTrad --targets="DiffTrad/examples/*.txt" --ref="DiffTrad/examples/fr.txt" --ref-lang=french --output="test.html"`