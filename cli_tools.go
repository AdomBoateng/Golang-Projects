package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func createFile(filename string) {
	if filename == "" {
		log.Println("filename cannot be empty")
		os.Exit(1)
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	fmt.Println(filename, "successfully created")
	os.Exit(0)
}

func readFile(filename string) {
	if filename == "" {
		log.Println("filename cannot be empty")
		os.Exit(1)
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File successfully read", string(data))
	os.Exit(0)
}

func writeFile(filename string, content string) {
	if filename == "" {
		log.Println("filename cannot be empty")
		os.Exit(1)
	}
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("File successfully written")
	os.Exit(0)
}

func deleteFile(filename string) {
	fmt.Println("Are you sure sure you want to delete", filename, "?(y/n)")
	var response string
	fmt.Scanln(&response)
	if strings.ToLower(response) == "y" {
		err := os.Remove(filename)
		if err != nil {
			fmt.Println("Error deleting file:", err)
			return
		}
		fmt.Println(filename, "successfully deleted")
	} else {
		fmt.Println(filename, "not deleted")
	}
	os.Exit(0)
}

func listDir(dir string) {
	if dir == "" {
		log.Println("directory cannot be empty")
		os.Exit(1)
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
	os.Exit(0)
}

func createFiles(filenames []string) {
	for _, filename := range filenames {
		if _, err := os.Stat(filename); err == nil {
			fmt.Println("File already exists:", filename)
			continue
		} else if os.IsNotExist(err) {
			file, err := os.Create(filename)
			if err != nil {
				fmt.Println("Error creating file:", err)
				continue
			}
			defer file.Close()
			fmt.Println(filename, "created successfully")
		} else {
			fmt.Println("Error checking file:", err)
		}
	}
}

func writeFiles(filenames []string, content string) {
	for _, filename := range filenames {
		err := os.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			continue
		}
		fmt.Println(filename, "written successfully")
	}
}

func deleteFiles(filenames []string) {
	fmt.Println("Are you use you want to delete these files?(y/n)")
	var response string
	fmt.Scanln(&response)
	if strings.ToLower(response) == "y" {
		for _, filename := range filenames {
			err := os.Remove(filename)
			if err != nil {
				fmt.Println("Error deleting file:", err)
				continue
			}
			fmt.Println(filename, "deleted successfully")
		}
	} else {
		fmt.Println("Files not deleted")
	}
}

func createDir(dir string) {
	if dir == "" {
		log.Println("directory cannot be empty")
		os.Exit(1)
	}
	err := os.Mkdir(dir, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	fmt.Println(dir, "created successfully")
	os.Exit(0)
}

func deleteDir(dir string, force bool) {
	fmt.Println("Are you sure you want to delete", dir, "(y/n)")
	var response string
	fmt.Scanln(&response)
	if strings.ToLower(response) == "y" {
		if force {
			err := os.RemoveAll(dir)
			if err != nil {
				fmt.Println("Error deleting directory:", err)
				return
			}
			fmt.Println(dir, "deleted successfully")
			return
		} else {
			err := os.Remove(dir)
			if err != nil {
				fmt.Println("Error deleting directory:", err)
				return
			}
			fmt.Println(dir, "deleted successfully")
		}
	} else {
		fmt.Println(dir, "not deleted")
	}
	os.Exit(0)
}

func main() {
	// Define the flags
	filename := flag.String("filename", "", "Filename to create/read/write/delete")
	content := flag.String("content", "", "Content to write to the file")
	action := flag.String("action", "read", "Action to perform: create-multiple, create, read, write, write-multiple, delete, delete-multiple,create-dir,delete-dir, list")
	dir := flag.String("dir", " ", "Directory to list file")
	filenames := flag.String("filenames", "", "Filenames to create/read/write/delete")
	force := flag.Bool("force", false, "Forceful removal of a directory")

	// Parse flags
	flag.Parse()

	// Perform action based on flag
	switch *action {
	case "create":
		createFile(*filename)
	case "read":
		readFile(*filename)
	case "write":
		writeFile(*filename, *content)
	case "delete":
		deleteFile(*filename)
	case "list":
		listDir(*dir)
	case "create-multiple":
		createFiles(strings.Split(*filenames, ","))
	case "write-multiple":
		writeFiles(strings.Split(*filenames, ","), *content)
	case "delete-multiple":
		deleteFiles(strings.Split(*filenames, ","))
	case "create-dir":
		createDir(*dir)
	case "delete-dir":
		deleteDir(*dir, *force)
	default:
		fmt.Println("Invalid action")
	}
}
