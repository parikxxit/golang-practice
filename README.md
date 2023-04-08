# golang-practice
practicing golang using ardanlabs practicle golang foundations 

## Strings and Formatting

Files to look -> banner/banner.go

- If you run banner fuction with emoji whitespace of function call with emoji will not get printed because go uses utf-8 encoding ~= rune type in go and utf-8 differ from 1 byte to 4 byte so english words are of 1 byte and emojis are of 3 bytes due to which padding calucation will get differ, and len() method on go returns number of bytes not number of char.

- Range on string go on char by char that is on of the mechanical what difference b/w range and len

- So strings in go can be of byte i.e utf-8 or int32 i.e rune 

- Basic type in go we do not have any method associated to them we do not have string.someMethod but we have  string library which can perform string operation like substr

- to fix the formatting bug we can utf8 package RuneCountInString methdo to count all rune instead of going by bytes

- Use #v while logging/debugging as it mention its type Ex: fmt.Print("#%v) it will log value with type ex: "1" for string instead of 1 

- If we are doing a string operation which need to access index convert that string to rune and return  string its a good practice

- %x in formatting printing can convert []byte to strign ex: fmt.Sprintf("%x", some[]ByteTypeVar)

## Calling rest API and working with JSON

- Http header is of type http.Header which is a map we are not directly using go's map coz http header can be case insensetive

- log.Fatalf is equivalent to log.Printf() && os.Exit(1)  

- serialization is a process of converting type to bytes and deserialization is process to taking byte and converting it into server side type it is also known as marshelling and unmarshalling respectively, most common serialization and deserialization method is JSON other example are XML, protobuff etc.

### JSON <-> GO type conversion

- true/false -> true/false
- string -> string
- null -> nil
- number -> float64(by default), float32, int64, int32, int8, uint8...
- array -> []any or []interface{} (before go 1.18 when generic was not defined)
- object -> map[string]any, struct

encoding/JSON API

- JSON -> io.Reader -> Go: json.Decoder
- JSON -> []byte -> Go: json.Unmarshal
- GO -> io.Wrtier -> json: json.Encoder
- GO -> []byte -> json: json.Marshal

## Working with file

- file.Open will open the file but we need to close that file also coz any system has limited amount of fileDescriptors so if we open multiple file and forgot to close it then file Descriptor limit will get fulfilled and we will not be able to open more new file you can check your limit via `ulimit -a`

- in go we use file.Close to close the file but we usually call it with defer keyword defer keyword run the line of code at the end of function call if multiple defers are there then it will run in reverse order

- if any interface provide close method its good practice to close it using defer

## Slice

- len -> to get the len of a slice and len is nil safe i.e for nil slice len is 0

- slicing operation slice[start:end] including start excluding end

- for tricks on slice you can check [this](https://ueokande.github.io/go-slice-tricks/)
## Struct

- Go compile does escape analysis and if we are returning a pointer then it will create it on heap

- Go build ``` go build gcflags=-m ``` to see the escape analysis done by compiler

- ``` go clean ``` clean all the previous go build

- we can add receiver to the struct and create a method receiver act similar to this of other langauge but we have to type of recievers value and pointer so pointer receiver change the original value and value receivers do not

- if we want to move the struct around use pointer receivers else use value receivers 

- we should never mix receivers type once it is defined as value type throught the code it should be value type this should be decided during desinging of struct else mix receivers cause integrity issue in the code but there is one exception to it i.e marshalling and unmarshelling shloud be of pointer receivers even tough we desided to go with value type receivers

- struct can embed another struct by doing so we can directly access the mothod/value to the parent struct

- If parent struct and embeded struct have same method then it will be overriden by parent struct but you can  access it via ``` parent.nestedStruct.CommonMethod``` will call the parent one only

- If 2 field are embeded in a struct and both have common key/field then if we try to access parent.CommonKey will give comiple time error and compiler do not konw from which struct we need to pick CommonKey

- Embeding is not inheritance if want to know more look for inheritance vs embedding

## Interface

- To group different type of struct/data by what it does we define an interface for it

- Its convention/good practice in go to make interface small usually 2 or less then 2 methods are there in go standard lib

- to group the interface we need to be explicit i.e if interface is implemented via pointer sematinc grouping will going to have pointer value if implemented via value semantics we need to group it as value not by referance 


