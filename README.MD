# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

It can output your final numerical grade in the class and determine whether you have passed or failed.

Runtime requirements:
Golang 1.24+

To download and install Golang, navigate to the official Golang website at: go.dev/doc/install then download the appropriate installer package. Run the installer and verify installation by running the following command in the terminal:
```
go version
```

If Go was installed correctly, you will see the appropriate Go version displayed as output on your terminal.



Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```