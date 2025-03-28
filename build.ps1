# Build service
go build -o ./bin/gital-service.exe ./service/

# Build fat-client
Set-Location .\fat-client 
wails build -o gital.exe
Set-Location ..\
Copy-Item -Path "fat-client\build\bin\gital.exe" -Destination "bin\gital.exe"