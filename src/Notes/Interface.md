# Interface

## day4

```go

type interfaceName interface{
    methodName(args args'type)(returnTypes,eeror)
}
```

interface means a group of actions should be realized .
so an struct ,can just show the limited actions to its customers.

```go
var myAPI interfaceName // declaration for this api
var mySTRUCT = & myStruct{} //declaration for a struct

myAPI=mySTRUCT //then ,bind this struct entity to this API entity

myAPI() // now ,you can use this the method in this api
```

need attention:
1. like struct ,an interface(A) can be included in another interface(B,C) ,means this interface extend all the method (B,C).
2. 