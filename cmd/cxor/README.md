go build -o cxor main.go

./cxor e -f filename key generatedFile
./cxor d key generatedFile

./cxor e 'msg' key generatedFile
./cxor d key generatedFile