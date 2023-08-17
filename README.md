# epoch
is a console app for creating timelines that can be exported to HTML or other formats
it also has an interactive console to interact with events and get various information
epoch internally uses Julian days for date time storing so it supports a negative year for BC

## build
```
go build -o epoch cmd/main.go
```
## use

epoch forks on JSON document files and a file must be provided on the start
to print a document on the console:
```
./epoch -f test.json -p  
```
to work on file:
```
./epoch -f test.json
```
after opening the file, you can type '?' for more information on how to use it

### exporting document
to HTML
```
./epoch -f test.json -o test.html
```
to new JSON file:
```
./epoch -f test.json -o test.json
```

## Example
You can try examples from the test_data folder
print to console
```
./epoch -f test_data/ww2.json -p
```
export to HTML
```
./epoch -f test_data/ww2.json -o ww2.html 
```
edit document
```
./epoch -f test_data/ww2.json
```
inside document, you can use commands<br />
p - to print timeline<br />
s - save<br />
q - exit<br />
help - to look for other options<br />
