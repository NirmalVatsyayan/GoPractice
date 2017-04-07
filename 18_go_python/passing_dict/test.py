from ctypes import cdll
import json
#load go library
lib = cdll.LoadLibrary('./libdata.so')

print("Loaded go generated SO library")
print(dir(lib))
result = lib.data(json.dumps({"Nirmal":"Vatsyayan"}))
print(result)

# go build -buildmode=c-shared -o libdata.so data.go
