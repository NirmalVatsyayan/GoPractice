package main

import "fmt"
import "errors"
import "reflect"

// native error method
func testFunc(arg int)(int, error){
  if arg == 42 {
      return -1, errors.New("Native error --> Cannot work on meaning of Life !!")
  }
  return arg+1, nil
}

// custom type error method
// will have to implement error method for them

type argError struct{
    arg int
    problem string
}

func (e *argError) Error() string{
     return fmt.Sprintf("%d - %s", e.arg, e.problem)
}

func testFuncCustom(arg int)(int, error){
  if arg == 42 {
      return -1, &argError{arg, "Custom error --> Cannot work on meaning of Life !!"}
  }
  return arg+1, nil
}

func main(){
    fmt.Println("Learning errors")

    for _, i := range []int{1,2,3,42,4}{
      if ret, e := testFunc(i); e != nil{
          fmt.Println("Flow failed --> ",e)
      }else{
        fmt.Println("Flow worked --> ", ret)
      }
    }

    fmt.Println("\n\n")
    // using custom error
    for _, i := range []int{1,2,3,42,4}{
        if ret, e := testFuncCustom(i); e != nil{
            fmt.Println("Flow failed --> ",e)
        }else{
            fmt.Println("Flow worked --> ", ret)
          }
        }


    // if you need to access variables of custom error
    // then after getting error , you’ll need to get the error as an instance
    // of the custom error type via type assertion.
    x, y := testFuncCustom(42)
    fmt.Println(">>>> ",x)
    fmt.Println(y)

    fmt.Println(reflect.TypeOf(y))

    //type assertion to access internal variable
    ae, ok := y.(*argError) // passing pointer is necessary as it is
    // type of receiver in error function
    // if dont want to use pointer chane error receiver to non pointer and
    // return plain object of argError from testFuncCustom
    fmt.Println(ae)
    fmt.Println(ok)

    if ok{
      fmt.Println(">>>> ",ae.arg)
      fmt.Println(ae.problem)
    }

}
