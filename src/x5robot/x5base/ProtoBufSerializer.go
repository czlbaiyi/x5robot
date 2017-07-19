package x5base

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

//x5的ProtoBuf格式
//1.从子类到父类按序号序列化
//2. 1）先写入数据类型，2）如果是不定长度的，在1）之后追加长度，3）写入数据

const InheritFlag uint16 = 0xFEFE

const USE_VAR_INT32 = true
const USE_VAR_INT64 = true

const (
	PT_NotSupported int32 = iota
	PT_UInt32
	PT_Int32
	PT_UInt64
	PT_Int64
	PT_Float32
	PT_Float64
	PT_String
	PT_Array
	PT_Slice
	PT_Map
	PT_Struct
)

const (
	PT_LengthDelimited int32 = iota
	PT_VarInt                = 1
	PT_Data32                = 2
	PT_Data64                = 3
)

func Serialize(netMsg interface{}) []byte {
	allBuffs := bytes.NewBuffer([]byte{})

	fmt.Println("---------------------")
	t := reflect.TypeOf(netMsg)
	v := reflect.ValueOf(netMsg)
	// 如果是指针，则获取其所指向的元素
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	fmt.Println("---------------------")
	to_var_uint64(99999999999, allBuffs)
	fmt.Println(from_var_uint64(allBuffs))
	fmt.Println("---------------------")

	// //1 p.CLSID
	CLSID := v.FieldByName("CLSID")
	binary.Write(allBuffs, binary.LittleEndian, CLSID.Interface())

	// //2 p.Serial
	Serial := v.FieldByName("Serial")
	binary.Write(allBuffs, binary.LittleEndian, Serial.Interface())

	// //3 p.SeqOrAck
	SeqOrAck := v.FieldByName("SeqOrAck")
	binary.Write(allBuffs, binary.LittleEndian, SeqOrAck.Interface())

	serialize(netMsg, allBuffs)

	return allBuffs.Bytes()
}

func serialize(netMsg interface{}, allBuffs *bytes.Buffer) {
	t := reflect.TypeOf(netMsg)
	v := reflect.ValueOf(netMsg)
	// 如果是指针，则获取其所指向的元素
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	// 进一步获取 i 的方法信息
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name)
		tag := t.Field(i).Tag.Get("x5tag")
		if tag != "" {
			fmt.Println(tag)
			tempT := t.Field(i).Type
			tempV := v.Field(i)
			if tempT.Kind() == reflect.Ptr {
				tempT = tempT.Elem()
				tempV = tempV.Elem()
			}

			if tag == "inherit" {
				if i != t.NumField()-1 {
					panic("被继承的类必须是放在最后一个，此条仅限于x5tag的协议标记")
				} else {
					serialize(tempV.Interface(), allBuffs)
				}
			} else {
				if tempT.Kind() == reflect.Struct {
					serialize(tempV.Interface(), allBuffs)
				} else {
					encode_value_to_buf(i, tempV, allBuffs)
				}
			}
		}
	}
}

func get_wire_type(vt int32) int32 {
	if vt == PT_Int32 || vt == PT_UInt32 {
		if USE_VAR_INT32 {
			return PT_VarInt
		} else {
			return PT_Data32
		}
	} else if vt == PT_Int64 || vt == PT_Time64 || vt == PT_UInt64 {
		if USE_VAR_INT64 {
			return PT_VarInt
		} else {
			return PT_Data64
		}
	} else if vt == PT_Float32 {
		return PT_Data32
	} else if vt == PT_Float64 {
		return PT_Data64
	} else {
		return PT_LengthDelimited
	}
}

func get_value_type(ft reflect.Kind) int32 {
	if ft == reflect.Int32 || ft == reflect.Int8 || ft == reflect.Int16 || ft == reflect.Bool {
		return PT_Int32
	} else if ft == reflect.Uint32 || ft == reflect.Uint8 || ft == reflect.Uint16 {
		return PT_UInt32
	} else if ft == reflect.Int64 {
		return PT_Int64
	} else if ft == reflect.Uint64 {
		return PT_UInt64
	} else if ft == reflect.Float32 {
		return PT_Float32
	} else if ft == reflect.Float64 {
		return PT_Float64
	} else if ft == reflect.String {
		return PT_String
	} else if ft == reflect.Array {
		return PT_Array
	} else if ft == reflect.Slice {
		return PT_Slice
	} else if ft == reflect.Map {
		return PT_Map
	} else if ft == reflect.Struct {
		return PT_Struct
	} else {
		panic("不支持的数据类型")
	}
}

func format_type_value(v reflect.Value) (int32, interface{}) {
	if !v.CanInterface() {
		panic("不支持无法Interface()的数据")
	}

	vr := v.Interface()
	fmt.Println(v.Kind())
	switch v.Kind() {
	case reflect.Bool:
		if vr == true {
			return PT_Int32, int32(1)
		} else {
			return PT_Int32, int32(0)
		}
	case reflect.Int8:
		return PT_Int32, vr.(int32)
	case reflect.Int16:
		return PT_Int32, vr.(int32)
	case reflect.Int32:
		return PT_Int32, vr.(int32)
	case reflect.Uint8:
		return PT_Int32, vr.(uint32)
	case reflect.Uint16:
		return PT_Int32, vr.(uint32)
	case reflect.Uint32:
		return PT_Int32, vr.(uint32)
	case reflect.Int64:
		return PT_Int64, vr.(int64)
	case reflect.Uint64:
		return PT_UInt64, vr.(uint64)
	default:
		panic("不支持的序列化类型")
	}
}

func encode_value_to_buf(idx int, v reflect.Value, allBuffs *bytes.Buffer) {
	if !v.CanInterface() {
		panic("不支持无法Interface()的数据")
	}

	t, iv := format_type_value(v)
	fmt.Println(t, iv)
}

func to_var_int64(n int64, allBuffs *bytes.Buffer) {
	v := (n << 1) ^ (n >> 63)
	for {
		abyte := byte(v) & 0x7F
		v = v >> 7
		if v == 0 {
			binary.Write(allBuffs, binary.LittleEndian, &abyte)
			break
		} else {
			abyte |= 0x80
			binary.Write(allBuffs, binary.LittleEndian, &abyte)
		}
	}
}

func to_var_uint64(n uint64, allBuffs *bytes.Buffer) {
	v := (n << 1) ^ (n >> 63)
	for {
		abyte := byte(v) & 0x7F
		v = v >> 7
		if v == 0 {
			binary.Write(allBuffs, binary.LittleEndian, &abyte)
			break
		} else {
			abyte |= 0x80
			binary.Write(allBuffs, binary.LittleEndian, &abyte)
		}
	}
}

func from_var_int64(allBuffs *bytes.Buffer) int64 {
	var n int64 = 0
	var shift uint64 = 0
	for {
		var abyte byte
		binary.Read(allBuffs, binary.LittleEndian, &abyte)
		bvar := (int64)(abyte & 0x7f)
		n += bvar << shift
		shift += 7
		if (abyte & 0x80) == 0 {
			break
		}

	}
	return (n >> 1) ^ (-(n & 1))
}

func from_var_uint64(allBuffs *bytes.Buffer) uint64 {
	var n uint64 = 0
	var shift uint64 = 0
	for {
		var abyte byte
		binary.Read(allBuffs, binary.LittleEndian, &abyte)
		bvar := (uint64)(abyte & 0x7f)
		n += bvar << shift
		shift += 7
		if (abyte & 0x80) == 0 {
			break
		}

	}
	return (n >> 1) ^ (-(n & 1))
}
