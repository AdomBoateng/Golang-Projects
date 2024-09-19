# CLI Tools Project

This is a simple command-line interface (CLI) tool written in Golang. It allows users to perform basic file and directory operations such as creating, reading, writing, deleting files and directories, and listing the contents of directories. The tool also supports multi-file operations.

## Features

- **Create**: Create single or multiple files.
- **Create-dir**: Create directory.
- **Read**: Read the contents of a file.
- **Write**: Write or append content to a file or multiple files.
- **Delete**: Delete single or multiple files (with user confirmation).
- **Delete-dir**: Delete directory (with user confirmation).
- **Force**: Forcefully delete directory if not empty.
- **List**: List the contents of a directory, including support for recursive listing.

## Prerequisites

Before running this project, ensure you have the following installed:

- [Golang](https://golang.org/dl/) (version 1.16 or later)
- A terminal or command-line interface (CLI) to run the commands.

## Getting Started

### 1. Clone the repository

Clone this repository to your local machine using the command:

```bash
git clone https://github.com/AdomBoateng/Golang-Projects.git
cd Golang-Projects
```

### 2. Build the project

After cloning the project, you can build it by running the following command:

```bash
go build -o Golang-Projects
```

This command will create an executable file named `Golang-Projects` in your project directory.

### 3. Run the project

Once built, you can run the CLI tool by executing the following command in your terminal:

```bash
./Golang-Projects -action <action> -filename <filename> -content <content> -filenames <filenames> -dir <directory>
```

### CLI Commands

| Flag       | Description                                                                      |
|------------|----------------------------------------------------------------------------------|
| `-action` | Action to perform. Supported actions: `create`, `read`, `write`, `delete`, `list`, `create-multiple`, `write-multiple`, `delete-multiple`, `create-dir`,`delete-dir` |
| `-filename` | The filename for single-file actions. (e.g., `test.txt`)                         |
| `-filenames` | Comma-separated filenames for multi-file actions. (e.g., `file1.txt,file2.txt`)   |
| `-content`  | The content to write into the file(s).                                            |
| `-dir`      | The directory to list files from.                                                 |
| `-force`  | Forcefully remove a directory with content(s)                                         |

### Usage Examples

1. **Create a File**

```bash
./Golang-Projects -action create -filename test.txt
```

This will create a file named `test.txt`.

2. **Read a File**

```bash
./Golang-Projects -action read -filename test.txt
```

This will read and display the contents of `test.txt`.

3. **Write to a File**

```bash
./Golang-Projects -action write -filename test.txt --content "Hello, World!"
```

This will write `"Hello, World!"` to `test.txt`.

4. **Delete a File**

```bash
./Golang-Projects -action delete -filename test.txt
```

This will prompt the user to confirm before deleting `test.txt`.

5. **List Files in a Directory**

```bash
./Golang-Projects -action list -dir ./mydirectory
```

This will list all the files in the specified directory (`./mydirectory`).

6. **Create Multiple Files**

```bash
./Golang-Projects -action create-multiple -filenames file1.txt,file2.txt,file3.txt
```

This will create three files: `file1.txt`, `file2.txt`, and `file3.txt`.

7. **Write to Multiple Files**

```bash
./Golang-Projects -action write-multiple -filenames file1.txt,file2.txt --content "Sample text"
```

This will write `"Sample text"` to both `file1.txt` and `file2.txt`.

8. **Delete Multiple Files**

```bash
./Golang-Projects -action delete-multiple -filenames file1.txt,file2.txt
```
This will prompt the user to confirm before deleting `file1.txt` and `file2.txt`.

9. **Create Directory**

```bash
./Golang-Projects -action create-dir -dir ./mydirectory
```

10. **Delete Directory**

```bash
./Golang-Projects -action delete-dir -dir ./mydirectory
```
This will prompt the user to confirm before deleting `./mydirectory`.

11. **Forcefully Delete Directory**

```bash
./Golang-Projects -action delete-dir -dir ./mydirectory -force=true
```
Use this when you want to forcefully delete a non-empty directory.
This will prompt the user to confirm before deleting `./mydirectory`.

### Notes

- The tool will prompt for confirmation before deleting files.
- Ensure that filenames are valid and not empty, as the tool performs basic validation.
- For recursive directory listing, you can modify the `list` action to handle nested directories (see the code).

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---