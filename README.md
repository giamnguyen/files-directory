# files-directory
This app has basic function for upload files to server and download files from server.

## Backend:
1.  Run server: `go run server.go`
2.  Run test: 
    * `go run server-test.go` 
    * then go to http://localhost:3000/ and try to upload some files.
    * those files will be upload to `uploaded-dir` folder.

## Frontend:
1. Run the app in the simulator of Xcode (Note: iOS simulators of Iphone 11 ProMax work best).
2. Click "Upload" button to upload file to `uploaded-dir` folder.
3. Click "Multiple Files" button to upload file to `uploaded-dir` folder.

