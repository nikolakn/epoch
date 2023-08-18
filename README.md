# epoch

Is a console app for creating timelines that can be exported to HTML or other formats  
it also has an interactive console to interact with events and get various information  
epoch internally uses Julian days for date time storing so it supports negative years for BC  

## build

```bash
go build -o epoch cmd/main.go
```

## use

epoch works on JSON document files and a file must be provided at the start
to print a document on the console:

```bash
./epoch -f test.json -p  
```

to work on file:

```bash
./epoch -f test.json
```

after opening the file, you can type '?' for more information on how to use it

### exporting document

to HTML

```bash
./epoch -f test.json -o test.html
```

to new JSON file:

```bash
./epoch -f test.json -o test.json
```

## Example

You can try examples from the test_data folder
print to console

```bash
./epoch -f test_data/ww2.json -p
```

export to HTML

```bash
./epoch -f test_data/ww2.json -o ww2.html 
```

edit document

```bash
./epoch -f test_data/ww2.json
```


inside a document, you can use commands:

+ p - to print timeline  
+ s - save  
+ q - exit
+ help,? - to look for other options  
