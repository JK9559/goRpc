package ggob

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
)

// 准备编码的数据结构
type P struct {
	X, Y, Z int
	Name    string
}

// 接收解码结果的数据结构
type Q struct {
	X, Y *int32
	Name string
}

func encode(data interface{}) *bytes.Buffer {
	//Buffer类型实现了io.Writer接口
	var buf bytes.Buffer
	//得到编码器
	enc := gob.NewEncoder(&buf)
	//编码器调Encode编码数据data
	if err := enc.Encode(data); err != nil {
		fmt.Println(err)
	}
	//结果放在buf里
	return &buf
}

func decode(data interface{}) *Q {
	d := data.(io.Reader)
	//获取一个解码器，参数需要实现io.Reader接口
	dec := gob.NewDecoder(d)
	var q Q
	//调解码器的Decode方法解码数据，用Q类型的q接收
	if err := dec.Decode(&q); err != nil {
		fmt.Println(err)
	}
	return &q
}

func Gob1test() {
	// 初始化数据
	data := P{3, 4, 5, "Kawin"}
	//
	buf := encode(data)
	fmt.Println(buf)
	var q *Q
	q = decode(buf)
	fmt.Println(*q.X, *q.Y, q.Name)
}
