package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Receiving File...")

	r.ParseMultipartForm(32 << 20) // 32MB is the default used by FormFile
	
	fhs := r.MultipartForm.File["myFile"] // Get list of FileHeader
	for _, fh := range fhs {
		f, err := fh.Open()
		// f is one of the files
		if err != nil {
			fmt.Println("Error Retrieving the File: "+fh.Filename)
			fmt.Println(err)
			return
		}
		defer f.Close()
		fmt.Printf("Uploaded File: %+v\n", fh.Filename)
		fmt.Printf("File Size: %+v\n", fh.Size)
		fmt.Printf("MIME Header: %+v\n", fh.Header)

		// Create a temporary file within "uploaded-dir" directory that follows
		// a particular naming pattern
		tempFile, err := ioutil.TempFile("uploaded-dir", "*-"+fh.Filename)
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()

		// read all of the contents of our uploaded file into a byte array
		fileBytes, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println(err)
		}

		// write this byte array to our temporary file
		tempFile.Write(fileBytes)
	}

	// return that we have successfully uploaded our file!
    fmt.Fprintf(w, "Successfully Uploaded Files\n")
}

func setupRoutes() {
    http.HandleFunc("/upload", uploadFile)
    http.ListenAndServe(":8080", nil)
}

func main() {
    fmt.Println("file-directory Server Application")
    setupRoutes()
}