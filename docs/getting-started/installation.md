# Installation

## Requirements

- Go 1.21 or later
- Git (for cloning the repository)

## Installing the Library

Add structured-goals to your Go project:

```bash
go get github.com/grokify/structured-goals
```

## Installing CLI Tools

### genschema - JSON Schema Generator

```bash
go install github.com/grokify/structured-goals/cmd/genschema@latest
```

### genpandoc - Pandoc Markdown Generator

```bash
go install github.com/grokify/structured-goals/cmd/genpandoc@latest
```

## Verifying Installation

Verify the library is installed correctly:

```bash
go list -m github.com/grokify/structured-goals
```

Verify CLI tools:

```bash
genschema --help
genpandoc --help
```

## Building from Source

Clone and build the project:

```bash
git clone https://github.com/grokify/structured-goals.git
cd structured-goals
go build ./...
go test ./...
```

## Optional Dependencies

For PDF generation with Pandoc renderer:

- [Pandoc](https://pandoc.org/installing.html) - Document converter
- [LuaLaTeX](https://www.latex-project.org/get/) - LaTeX distribution (TeX Live or MiKTeX)

For Marp slide generation:

- [Marp CLI](https://github.com/marp-team/marp-cli) - Markdown presentation tool
