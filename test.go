package main
import (
    "fmt"
    "gen-go/src"
    "github.com/golang/protobuf/proto"
)
func main() {
    test := example.AllDataType{
        Fint32:      257,
        Fint64:      -2,
        Fuint32:     1,
        Fuint64:     1025,
        Fsint32:     0,
        Fsint64:     -2,
        Ffixed32:    17,
        Ffixed64:    2049,
        Fdouble:     -0.1,
        Ffloat:      0.6,
        Fbool:       true,
        Fenum:       example.DayOfWeek_SUNDAY,
        Fmessage:    &example.Child{Fsint64: 3},
        Fmap:        map[uint32]float64{3: -0.1, 0: 2.12},
        Frepeatbool: []bool{true, false, true},
        Fstring:     "Hello World",
        Fbytes:      []byte{129, 0, 19, 56},
        Fsfixed32:   12345,
        Fsfixed64:   54321,
    }
    data, _ := proto.Marshal(&test) // protobuf将结构体序列化为二进制串
    fmt.Println(data) // 打印AllDataType类型的数据序列化后的二进制串
}