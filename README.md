# cmakeinit — CMakeLists.txt Generator for C/C++ Projects

A lightweight CLI tool written in Go that generates `CMakeLists.txt` files for C/C++ projects.

## Features

- Generates standard `CMakeLists.txt`
- Recursively collects `.c` source files
- Adds include directories
- Enables `compile_commands.json` output
- Interactive and non-interactive modes
- Configurable via command-line flags

## Requirements

- Go

## Installation

### Using go install

```bash
go install github.com/CyberTea0X/cmakeinit@latest
```

Make sure ~/go/bin is in your PATH. 

### Manual build 

```bash
git clone https://github.com/CyberTea0X/cmakeinit.git
cd cmakeinit
go build -o cmakeinit
```

## Usage

Interactive mode
```bash
cmakeinit .
```
You will be prompted to enter project settings.

Non-interactive mode
```bash
Usage of cmakeinit:
  -ac
        Create folders and main.c file (default notset)
  -i string
        Include directory (default "./include")
  -m string
        Minimum version of cmake (default "3.10")
  -name string
        CMakeLists.txt filename (default "CMakeLists.txt")
  -p string
        Project name (default "project")
  -s string
        Source directory (default "./src")
```

## Generated Output

The generated `CMakeLists.txt` includes (see [cmake.tmpl](cmake.tmpl)):
