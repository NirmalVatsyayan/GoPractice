from ctypes import cdll

#load go library
lib = cdll.LoadLibrary('./libadd.so')

print("Loaded go generated SO library")
result = lib.add(2, 3)
print(result)

#go build -buildmode=c-shared -o libadd.so libadd.go
