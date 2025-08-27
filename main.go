package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

type CmakeLists struct {
	CmakeMinimumRequired string
	ProjectName          string
	ExecutableName       string
	MainFile             string
	SrcDir               string
	IncludeDir           string
}

//go:embed cmake.tmpl
var cmakeTmpl string

func main() {
	cmakeMinimumRequired := flag.String("m", "3.10", "Minimum version of cmake")
	projectName := flag.String("p", "project", "Project name")
	executableName := flag.String("e", "main", "Executable name")
	srcDir := flag.String("s", "./src", "Source directory")
	includeDir := flag.String("i", "./include", "Include directory")
	isInteractive := flag.Bool("i", true, "Interactive mode")
	flag.Parse()
	var cmakeLists CmakeLists
	if isInteractive == nil {
		cmakeLists.CmakeMinimumRequired = *cmakeMinimumRequired
		cmakeLists.ProjectName = *projectName
		cmakeLists.ExecutableName = *executableName
		cmakeLists.SrcDir = *srcDir
		cmakeLists.IncludeDir = *includeDir
	} else {
		cmakeLists = interactiveCmake()
	}

	// Парсим шаблон из эмбеднутой строки
	tmpl, err := template.New("cmake").Parse(cmakeTmpl)
	if err != nil {
		log.Fatal("Ошибка парсинга шаблона:", err)
	}

	// Выполняем шаблон
	err = tmpl.Execute(os.Stdout, cmakeLists)
	if err != nil {
		log.Fatal("Ошибка выполнения шаблона:", err)
	}
}

func scanOrDefault(defaultValue string) string {
	var value string
	fmt.Scanln(&value)
	if strings.TrimSpace(value) == "" {
		value = defaultValue
	}
	return value
}

func interactiveCmake() CmakeLists {
	cmakeLists := CmakeLists{}
	fmt.Print("Название проекта: (по умолчанию project)")
	cmakeLists.ProjectName = scanOrDefault("project")
	fmt.Print("Имя исполняемого файла: (по умолчанию main)")
	cmakeLists.ExecutableName = scanOrDefault("main")
	fmt.Print("Путь к исходному коду: (по умолчанию ./src)")
	cmakeLists.SrcDir = scanOrDefault("./src")
	fmt.Print("Путь к заголовочным файлам: (по умолчанию ./include)")
	cmakeLists.IncludeDir = scanOrDefault("./include")
	fmt.Print("Версия cmake для сборки: (по умолчанию 3.10)")
	cmakeLists.CmakeMinimumRequired = scanOrDefault("3.10")
	return cmakeLists
}
