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
	MainFile             string
	SrcDir               string
	IncludeDir           string
	Filename             string
}

//go:embed cmake.tmpl
var cmakeTmpl string

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cmakeMinimumRequired := flag.String("m", "3.10", "Minimum version of cmake")
	projectName := flag.String("p", "project", "Project name")
	srcDir := flag.String("s", "./src", "Source directory")
	includeDir := flag.String("i", "./include", "Include directory")
	listsFilename := flag.String("name", "CMakeLists.txt", "CMakeLists.txt filename")
	autoCreate := flag.Bool("ac", false, "Create folders and main.c file (default notset)")
	flag.Parse()
	var cmakeLists CmakeLists
	if flag.NFlag() > 0 {
		cmakeLists.CmakeMinimumRequired = *cmakeMinimumRequired
		cmakeLists.ProjectName = *projectName
		cmakeLists.SrcDir = *srcDir
		cmakeLists.IncludeDir = *includeDir
		cmakeLists.Filename = *listsFilename
		if autoCreate != nil && *autoCreate {
			createFilesAndFolders(cmakeLists)
		}
		cmakeLists.SrcDir = strings.TrimPrefix(cmakeLists.SrcDir, ".")
		cmakeLists.IncludeDir = strings.TrimPrefix(cmakeLists.IncludeDir, ".")
	} else {
		cmakeLists = interactiveCmake()
	}

	// Парсим шаблон из эмбеднутой строки
	tmpl, err := template.New("cmake").Parse(cmakeTmpl)
	if err != nil {
		log.Fatal("Ошибка парсинга шаблона:", err)
	}
	file, err := os.OpenFile(cmakeLists.Filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Ошибка открытия файла:", err)
	}
	defer file.Close()

	// Выполняем шаблон
	err = tmpl.Execute(file, cmakeLists)
	if err != nil {
		log.Fatal("Ошибка выполнения шаблона:", err)
	}
	if err := file.Close(); err != nil {
		log.Fatal("Ошибка закрытия файла:", err)
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

func createFilesAndFolders(cmakeLists CmakeLists) {
	os.MkdirAll(cmakeLists.SrcDir, 0755)
	os.MkdirAll(cmakeLists.IncludeDir, 0755)
	file, err := os.OpenFile(cmakeLists.SrcDir+"/main.c", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Ошибка открытия файла:", err)
	}
	defer file.Close()
	_, err = file.WriteString("int main(int argc, char **argv) { return 0; }")
	if err != nil {
		log.Fatal("Ошибка записи в файл:", err)
	}
}

func interactiveCmake() CmakeLists {
	cmakeLists := CmakeLists{}
	fmt.Print("Название проекта: (по умолчанию project)")
	cmakeLists.ProjectName = scanOrDefault("project")
	fmt.Print("Путь к исходному коду: (по умолчанию ./src)")
	cmakeLists.SrcDir = scanOrDefault("./src")
	fmt.Print("Путь к заголовочным файлам: (по умолчанию ./include)")
	cmakeLists.IncludeDir = scanOrDefault("./include")
	fmt.Print("Версия cmake для сборки: (по умолчанию 3.10)")
	cmakeLists.CmakeMinimumRequired = scanOrDefault("3.10")
	fmt.Print("Название файла (по умолчанию CMakeLists.txt)")
	cmakeLists.Filename = scanOrDefault("CMakeLists.txt")
	fmt.Print("Создать указанные файлы и папки автоматически? (по умолчанию yes)")
	if scanOrDefault("yes") == "yes" {
		createFilesAndFolders(cmakeLists)
	}

	cmakeLists.SrcDir = strings.TrimPrefix(cmakeLists.SrcDir, ".")
	cmakeLists.IncludeDir = strings.TrimPrefix(cmakeLists.IncludeDir, ".")

	return cmakeLists
}
