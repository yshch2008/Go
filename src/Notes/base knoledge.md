# base knowledge of Go Language  

## declare  

- use default value
> var identifier type
> var v_1,v_2,v_3
- judge by value without declaring type
> var v_name = value
> var vname1, vname2, vname3 = v1, v2, v3
- even without var,but using ":"
> v_name:=value

> vname1, vname2, vname3 := v1, v2, v3
- parallel assignment
> a,b:= Func_A("hello world")

## characters  

.
,
:
...

## types

### bool  

 >var b bool =true  

### numbers

    1. arithmetic within complement
    2. support complex
1. int(unit8,unit16,unit32,unit64,int8.int16,int32.int64)
2. float (float32,float64)
3. complex(complex64,complex128)
4. byte ,equals unit 8
5. rune,equals unit 32
6. uintptr

### string

> code within UTF-8,unicode
- derived types
 1. pointer (*ptr= &identifier ,looks like a number,but Point)
 2. array(a:=[...]int{},even you need it auto calculate count of this array ,you should not let the [] empty)
 3. struct
 4. channel
 5. function
 6. slice (a:=[]int{},if the [] is empty ,it is slice)
 7. interface
 8. map

## get memory addres 

> Pointer v_address  
> *v_address=&identifier  

## priority

|level|character|
|-|-|
|0|() [] -> . ++ --|
|1|! ^ type *& sizeof|
|2|* / %|
|3|+ -|
|4|>> <<|
|5|>= > < <=|
|6|!= ==|
|7|&|
|8|^|
|9|\||
|10|&&|
|11|\|\||
|12|? :|
|13|= += -= *= /= %= >>= <<= &= ^= \|=\

## key words

|index|key word|paraphrace|
|-|-|-|
|0| break|end current loop |
|1| case|used with switch or selection key-word,as the alternative additions|
|3| chan|channel identifier ,located in front of the virable value ,use "<-" instead of "=" to assignment ,always one side waiting for the another side get/set the value|
|4| const|constant identifier ,readonly and one-off assignment|
|5| continue|skip this time loop|
|6| default|the default case of switch or selection|
|7| defer|defer excecuting the rest part until the super lump ends|
|8| else|alternative case when the above if-else cases neither works|
|9| fallthrough|make the hide "break" disabled ,so it can continue until meet another hide "break"|
|10| for|contains while & for|
|11| func|difine a method|
|12| go|identifier for a |
|13| goto|jump to a position ,need a displayed identifier|
|14| if|the head of conditional structor, when follow an else,means additional cases|
|15| import|import packages , short-cut as "name : package_name"|
|16| interface|interface api|
|17| map|as same as hashmap or dictionary type|
|18| package|orgnize multiple files as a packge ,when name as "main", it will be compiled as a runable program|
|19| range|used like "in"|
|20| return|return values and error information ,then break this method|
|21| select|if there is no case according to the condition and there is default case ,just using default  case;if no default case ,waiting till it appears,then choose one case to oprate randomly|
|22| struct|identifier a struct,as some fields assemblied to a community |
|23| switch|conditional for multi cases ,can transfer to if-else or in return,can let param be empty for write a long if-else struct|
|24| type|identifier for a customed type|
|25| var|identifier for a variable value|

## identifier

- "_"
> empty character

|index|identifier|paraphrace|
|-|-|-|
|1|append||
|2|bool||
|3|byte||
|4|cap||
|5|close||
|6|complex||
|7|complex64||
|8|complex128||
|9|uint16||
|10|copy||
|11|false||
|12|float32||
|13|float64||
|14|imag||
|15|int||
|16|int8||
|17|int16||
|18|uint32||
|19|int32||
|20|int64||
|21|iota||
|22|len||
|23|make||
|24|new||
|25|nil||
|26|panic||
|27|uint64||
|28|print||
|29|println||
|30|real||
|31|recover||
|32|string||
|33|true||
|34|uint||
|35|uint8||
|36|uintptr||