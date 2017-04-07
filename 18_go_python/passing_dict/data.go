package main

import "C"
import "fmt"
import "encoding/json"


type Value struct {
	Name string
}

/*
-buildmode=c-shared
    Build the listed main packages, plus all packages that they
    import, into C shared libraries. The only callable symbols will
    be those functions exported using a cgo //export comment.
    Non-main packages are ignored.

   //export func_name is very important

*/

//export data
func data() int {

     var v Value
     bs := []byte(`{"Name":"Nirmal"}`)
     json.Unmarshal(bs, &v)
     fmt.Println(">>>>>>>>>>>>>>  ",v.Name)
     return 1
}

func main(){
}
