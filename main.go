package main

import (
	"fmt"
	"os"
	"net/http"
)

func main(){
	var fileName string
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	fmt.Printf("Opening file %v...\n", fileName)

	file, fileErr := os.Open(fileName)
	if fileErr != nil {
		fmt.Printf("Can't open file: %v\n", file)
		panic(fileErr)
	}
	defer file.Close()

	fileInfo, fileErr := file.Stat()
	if fileErr != nil {
		fmt.Printf("Can't get file info\n")
		panic(fileErr)
	}

	// Check if the file is a directory
    if fileInfo.IsDir() {
        fmt.Println("Error: file is a directory")
        return
    }

	var modTime = fileInfo.ModTime()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "video/mp4")

		http.ServeContent(w, r, fileName, modTime, file)
	})

	fmt.Printf("Serving file on http://localhost:9988\n")

	httpErr := http.ListenAndServe(":9988", nil)
	if httpErr != nil{
		panic(httpErr)
	}
}