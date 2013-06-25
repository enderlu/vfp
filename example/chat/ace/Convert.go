// 消息解析 Create By Ace 2013-2-1
package ace

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type AceBuffer struct {
	message []byte
	length  int
	index   int
}

const (
	INT_SIZE    = 4
	FLOAT_SIZE  = 4
	DOUBLE_SIZE = 8
	BYTE_SIZE   = 1
)

func NewBuffer() *AceBuffer {
	var in []byte
	return &AceBuffer{in, len(in), 0}
}

/*读取int*/
func (ab *AceBuffer) ReadInt() int {
	in := ab.message[ab.index : ab.index+INT_SIZE]
	result := int(in[3]) | (int(in[2]) << 8) | (int(in[1]) << 16) | (int(in[0]) << 24)
	ab.index += INT_SIZE
	return int(result)
}

/*读取float*/
func (ab *AceBuffer) ReadFloat() float32 {
	in := ab.message[ab.index : ab.index+FLOAT_SIZE]
	var result float32
	buf := bytes.NewBuffer(in)
	err := binary.Read(buf, binary.BigEndian, &result)
	if err != nil {
		fmt.Println("float解析失败", err)
	}
	ab.index += FLOAT_SIZE
	return result
}

/*读取double*/
func (ab *AceBuffer) ReadDouble() float64 {
	in := ab.message[ab.index : ab.index+DOUBLE_SIZE]
	var result float64
	buf := bytes.NewBuffer(in)
	err := binary.Read(buf, binary.BigEndian, &result)
	if err != nil {
		fmt.Println("double解析失败", err)
	}
	ab.index += DOUBLE_SIZE
	return result
}

func (ab *AceBuffer) ReadByte() byte {
	result := ab.message[ab.index]
	ab.message = ab.message[1:ab.length]
	ab.length = len(ab.message)
	ab.index += BYTE_SIZE
	return result
}

/*读取string*/
func (ab *AceBuffer) ReadString() string {
	length := ab.ReadInt()
	in := ab.message[ab.index:length]
	ab.index += length
	return string(in)
}

/*写入int*/
func (ab *AceBuffer) WriteInt(value int) {
	b := make([]byte, 4)
	b[0] = byte(value >> 24)
	b[1] = byte(value >> 16)
	b[2] = byte(value >> 8)
	b[3] = byte(value)
	ab.message = append(ab.message, b...)
	ab.length = len(ab.message)
}

/*写入byte*/
func (ab *AceBuffer) WriteByte(value byte) {
	ab.message = append(ab.message, value)
	ab.length = len(ab.message)
}

/*写入float*/
func (ab *AceBuffer) WriteFloat(value float32) {
	b := make([]byte, 0)
	buf := bytes.NewBuffer(b)
	err := binary.Write(buf, binary.BigEndian, &value)
	if err != nil {
		fmt.Println("float写入失败", err)
	}
	ab.message = append(ab.message, buf.Bytes()...)
	ab.length = len(ab.message)
}

/*写入double*/
func (ab *AceBuffer) WriteDouble(value float64) {
	b := make([]byte, 0)
	buf := bytes.NewBuffer(b)
	err := binary.Write(buf, binary.BigEndian, &value)
	if err != nil {
		fmt.Println("double写入失败", err)
	}
	ab.message = append(ab.message, buf.Bytes()...)
	ab.length = len(ab.message)
}

/*写入String*/
func (ab *AceBuffer) WriteString(value string) {
	b := []byte(value)
	ab.WriteInt(len(b))
	ab.message = append(ab.message, b...)
	ab.length = len(ab.message)
}

/*写入bytes*/
func (ab *AceBuffer) WriteBytes(value []byte) {
	ab.message = append(ab.message, value...)
	ab.length = len(ab.message)
}

/*读取byteArray*/
func (ab *AceBuffer) ReadBytes() []byte {
	result := ab.message[ab.index:ab.length]
	ab.index = ab.length
	return result
}

/*重置读取序**/
func (ab *AceBuffer) Reset() {
	ab.index = 0
}

/*清空buffer数据**/
func (ab *AceBuffer) Clear() {
	ab.message = []byte{}
	ab.index = 0
	ab.length = 0
}

func (ab *AceBuffer) Length() int {
	return ab.length
}

/*获取对象的byteArray值*/
func (ab *AceBuffer) Bytes() []byte {
	return ab.message
}
