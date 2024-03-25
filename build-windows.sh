GOOS=windows GOARCH=amd64 go build -o emlconvert.windows.jing ./main.go
zip -re emlconvert.windows2.zip emlconvert.windows.jing