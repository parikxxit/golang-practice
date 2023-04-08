# golang-practice
practicing golang using ardanlabs practicle golang foundations 

## Strings and Formatting

Files to look -> banner/banner.go

If you run banner fuction with emoji whitespace of function call with emoji will not get printed because go uses utf-8 encoding ~= rune type in go and utf-8 differ from 1 byte to 4 byte so english words are of 1 byte and emojis are of 3 bytes due to which padding calucation will get differ, and len() method on go returns number of bytes not number of char.

Range on string go on char by char that is on of the mechanical what difference b/w range and len

So strings in go can be of byte i.e utf-8 or int32 i.e rune 

Basic type in go we do not have any method associated to them we do not have string.someMethod but we have  string library which can perform string operation like substr

to fix the formatting bug we can utf8 package RuneCountInString methdo to count all rune instead of going by bytes

Use #v while logging/debugging as it mention its type Ex: fmt.Print("#%v) it will log value with type ex: "1" for string instead of 1 

If we are doing a string operation which need to access index convert that string to rune and return  string its a good practice

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


