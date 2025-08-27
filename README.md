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

Interactive mode (default)
```bash
cmakeinit
```
You will be prompted to enter project settings.

Non-interactive mode
```bash
cmakeinit -p MyProject -e myapp -s ./src -i ./include -m 3.16 -it=false
```

## Generated Output

The generated `CMakeLists.txt` includes:

```
cmake_minimum_required(VERSION ...)
project(...)
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)
file(GLOB_RECURSE SOURCES "src/*.c")
include_directories(include)
add_executable(${PROJECT_NAME} ${SOURCES})
```
