package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title   string
	Year    int  `json:"release"`         //成员标签，定时是结构体在成员在编译期间关联的一些元信息
	Color   bool `json:"color,omitempty"` // 而外选项 omitempty,如果成成员的值是0或者零值，不输出到JSON中
	testing string
	Actors  []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false, testing: "不可导出成员不会被转为json字段",
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Casablanca", Year: 1942, Color: false, testing: "不可导出成员不会被转为json字段",
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Casablanca", Year: 1942},
}

func main() {

	// ===============
	// 1. Go通过标准库 encoding/json, encoding.xml, encoding/asn1 和其他库提供相应的数据编码与解码支持

	// 把 Go数据结构转换为JSON 称为 marshal，
	// Year 被转换成 release, 是通过成员标签实现的
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// MarshalIndent转码为 json并格式化
	data2, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data2)

	// 2. JSON 到 Go数据结构，这个过程叫做  unmarshal
	// 合理的定义数据结构， 可以选择性的将json数据解码到结构体对象中
	var titles []struct{ Title string }
	// 只解析title
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshal failed: %s", err)
	}
	fmt.Println(titles)
	fmt.Printf("%#v\n", titles)

	// 2. 另外还有流式解码器 json.Decoder， 可以用来依次从字节流里面解码出多个json实体，
	// 还有个 json.Encoder流式编码器
	// json.NewDecoder(data).Decode(&result);

}
