# Ztorage

## Compile

`go get github.com/peterbourgon/diskv`
`cd C:/code/go/src/github.com/fredrik-hjarner/ztorage/ztorage`
`go install`

## Run

`C:/code/go/bin/ztorage.exe`

## Use

`http://localhost:8080/value?key=alpha%2Fbeta`

to read the value of 'alpha/beta'

## Test

`go test ./.. -v`
