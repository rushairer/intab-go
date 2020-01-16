package commands

import (
	"fmt"
	m "intab-core/models"
	"os"
)

//MigrateDB 迁移数据库
func MigrateDB() {
	m.Migrate()
	os.Exit(0)
}

//ShowHelp 显示命令行帮助
func ShowHelp() {
	fmt.Println("Restful API command tool.")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("    ./app [options]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("    migrate                 Automatically migrate your schema")
	fmt.Println("    help                    Show this help message and exit")
	fmt.Println()
	os.Exit(0)
}
