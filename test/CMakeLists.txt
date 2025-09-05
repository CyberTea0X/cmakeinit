cmake_minimum_required(VERSION 3.10)
project(project)

# Generate compile_commands.json
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)

# Dependencies

# Создаем список всех исходных файлов
file(GLOB_RECURSE SOURCES 
    "${CMAKE_CURRENT_SOURCE_DIR}/src/*.c"
)

# Добавляем заголовочные файлы в include пути
include_directories(
    ${CMAKE_CURRENT_SOURCE_DIR}/include
)

# Наш проект
add_executable(${PROJECT_NAME} ${SOURCES})