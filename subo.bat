git add .
git commit -m "Ultimo commit"
git push
set GOOS=linux
set GOARCH=amd64
go build main.go
del main.zip
tar -a -cf main.zip main