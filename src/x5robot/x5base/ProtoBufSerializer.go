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

	// fmt.Println("---------------------")
	// to_var_uint64(99999999999, allBuffs)
	// fmt.Println(from_var_uint64(allBuffs))
	// fmt.Println("---------------------")

	// //1 p.CLSID
	CLSID := v.FieldByName("CLSID")
	binary.Write(allBuffs, binary.LittleEndian, CLSID.Interface())

	// //2 p.Serial
	Serial := v.FieldByName("Serial")
	binary.Write(allBuffs, binary.LittleEndian, Serial.Interface())

	// //3 p.SeqOrAck
	SeqOrAck := v.FieldByName("SeqOrAck")
	binary.Write(allBuffs, binary.LittleEndian, SeqOrAck.Interface())

	// 序列化x5tag数据
	newBufs := bytes.NewBuffer([]byte{})
	save_protobuf_struct(-1, netMsg, newBufs)
	buffs := newBufs.Bytes()
	fmt.Println("buffs :", buffs)
	// 加密
	key := uint32(Serial.Interface().(int32))
	Encrypt(key, &buffs)

	//安全码
	var security_flag uint32 = 0xacabdeaf
	binary.Write(allBuffs, binary.LittleEndian, &security_flag)

	//求和
	checksum := GetCRC32(&buffs, 0, len(buffs))
	binary.Write(allBuffs, binary.LittleEndian, &checksum)

	//写入序列化数据长度
	len := int32(len(buffs))
	binary.Write(allBuffs, binary.LittleEndian, &len)

	//写入写入序列化数据
	binary.Write(allBuffs, binary.LittleEndian, &buffs)

	return allBuffs.Bytes()
}

func get_wire_type(vt int32) int32 {
	if vt == PT_Int32 || vt == PT_UInt32 {
		if USE_VAR_INT32 {
			return PT_VarInt
		} else {
			return PT_Data32
		}
	} else if vt == PT_Int64 || vt == PT_UInt64 {
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

func encode_value_to_buf(idx int32, v interface{}, allBuffs *bytes.Buffer) {
	fmt.Print("encode_value_to_buf:", v)

	vt := reflect.TypeOf(v)
	t := get_value_type(vt.Kind())
	wt := get_wire_type(t)

	switch vt.Kind() {
	case reflect.Bool:
		va := v.(bool)
		if va {
			save_protobuf_int32(idx, 1, allBuffs, wt)
		} else {
			save_protobuf_int32(idx, 0, allBuffs, wt)
		}
	case reflect.Int8:
		va := int32(v.(int8))
		save_protobuf_int32(idx, va, allBuffs, wt)
	case reflect.Int16:
		va := int32(v.(int16))
		save_protobuf_int32(idx, va, allBuffs, wt)
	case reflect.Int32:
		va := v.(int32)
		save_protobuf_int32(idx, va, allBuffs, wt)
	case reflect.Int64:
		va := v.(int64)
		save_protobuf_int64(idx, va, allBuffs, wt)
	case reflect.Uint8:
		va := uint32(v.(uint8))
		save_protobuf_uint32(idx, va, allBuffs, wt)
	case reflect.Uint16:
		va := uint32(v.(uint16))
		save_protobuf_uint32(idx, va, allBuffs, wt)
	case reflect.Uint32:
		va := v.(uint32)
		save_protobuf_uint32(idx, va, allBuffs, wt)
	case reflect.Uint64:
		va := v.(uint64)
		save_protobuf_uint64(idx, va, allBuffs, wt)
	case reflect.Float32:
		va := v.(float32)
		save_protobuf_float32(idx, va, allBuffs)
	case reflect.Float64:
		va := v.(float64)
		save_protobuf_float64(idx, va, allBuffs)
	case reflect.String:
		va := v.(string)
		save_protobuf_string(idx, []byte(va), allBuffs)
	case reflect.Struct:
		save_protobuf_struct(idx, v, allBuffs)
	default:
		panic("不支持的数据类型")
	}
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

func save_field_desc(idx int32, wt int32, allBuffs *bytes.Buffer) {
	if idx != -1 {
		desc := make_field_desc(idx, wt)
		binary.Write(allBuffs, binary.LittleEndian, &desc)
	}
}

func save_protobuf_lenth_delimited(idx int32, buff []byte, allBuffs *bytes.Buffer) {
	save_field_desc(idx, PT_LengthDelimited, allBuffs)
	len := int32(len(buff))
	binary.Write(allBuffs, binary.LittleEndian, &len)
	binary.Write(allBuffs, binary.LittleEndian, &buff)
}

func parse_field_index(desc uint16) int32 {
	return int32(desc >> 3)
}

func make_field_desc(idx int32, wt int32) uint16 {
	desc := (uint16)(((idx << 3) | wt))
	return desc
}

func parse_wire_type(wt int16) int32 {
	return int32(wt & 0x0007)
}

func save_protobuf_int32(idx int32, va int32, allBuffs *bytes.Buffer, wt int32) {
	if wt == PT_VarInt {
		save_field_desc(idx, PT_VarInt, allBuffs)
		to_var_int64(int64(va), allBuffs)
	} else {
		save_field_desc(idx, PT_Data32, allBuffs)
		binary.Write(allBuffs, binary.LittleEndian, &va)
	}
}

func save_protobuf_int64(idx int32, va int64, allBuffs *bytes.Buffer, wt int32) {
	if wt == PT_VarInt {
		save_field_desc(idx, PT_VarInt, allBuffs)
		to_var_int64(int64(va), allBuffs)
	} else {
		save_field_desc(idx, PT_Data64, allBuffs)
		binary.Write(allBuffs, binary.LittleEndian, &va)
	}
}

func save_protobuf_uint32(idx int32, va uint32, allBuffs *bytes.Buffer, wt int32) {
	if wt == PT_VarInt {
		save_field_desc(idx, PT_VarInt, allBuffs)
		to_var_uint64(uint64(va), allBuffs)
	} else {
		save_field_desc(idx, PT_Data32, allBuffs)
		binary.Write(allBuffs, binary.LittleEndian, &va)
	}
}

func save_protobuf_uint64(idx int32, va uint64, allBuffs *bytes.Buffer, wt int32) {
	if wt == PT_VarInt {
		save_field_desc(idx, PT_VarInt, allBuffs)
		to_var_uint64(uint64(va), allBuffs)
	} else {
		save_field_desc(idx, PT_Data64, allBuffs)
		binary.Write(allBuffs, binary.LittleEndian, &va)
	}
}

func save_protobuf_float32(idx int32, va float32, allBuffs *bytes.Buffer) {
	save_field_desc(idx, PT_Data32, allBuffs)
	binary.Write(allBuffs, binary.LittleEndian, &va)
}

func save_protobuf_float64(idx int32, va float64, allBuffs *bytes.Buffer) {
	save_field_desc(idx, PT_Data64, allBuffs)
	binary.Write(allBuffs, binary.LittleEndian, &va)
}

func save_protobuf_string(idx int32, va []byte, allBuffs *bytes.Buffer) {
	newBuff := bytes.NewBuffer([]byte{})
	binary.Write(newBuff, binary.LittleEndian, &va)
	var eos byte = 0
	binary.Write(newBuff, binary.LittleEndian, &eos)
	save_protobuf_lenth_delimited(idx, newBuff.Bytes(), allBuffs)
	fmt.Println(allBuffs.Bytes())
}

func save_protobuf_struct(idx int32, v interface{}, allBuffs *bytes.Buffer) {
	newBuff := bytes.NewBuffer([]byte{})
	encode_struct_to_buf(v, newBuff)
	save_protobuf_lenth_delimited(idx, newBuff.Bytes(), allBuffs)
}

func encode_struct_to_buf(v interface{}, allBuffs *bytes.Buffer) {
	rt := reflect.TypeOf(v)
	rv := reflect.ValueOf(v)

	// 如果是指针，则获取其所指向的元素
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
		rv = rv.Elem()
	}
	// 进一步获取 i 的方法信息
	for i := 0; i < rt.NumField(); i++ {
		tf := rt.Field(i)
		fmt.Println("tag:", "i = ", tf.Tag)
		tag := tf.Tag.Get("x5tag")
		if tag != "" {
			fmt.Println(tag)
			tft := tf.Type
			vf := rv.Field(i)
			fmt.Println("tft name:", tf.Name)
			if tft.Kind() == reflect.Ptr {
				tft = tft.Elem()
				vf = vf.Elem()
			}
			fmt.Println("tft name:", tft.Name)

			if tag == "inherit" {
				if i != rt.NumField()-1 {
					panic("被继承的类必须是放在最后一个，此条仅限于x5tag的协议标记")
				}
				binary.Write(allBuffs, binary.LittleEndian, InheritFlag)
				encode_struct_to_buf(vf.Interface(), allBuffs)
			} else {
				encode_value_to_buf(int32(i+1), vf.Interface(), allBuffs)
			}

			fmt.Println("allBuffs:", allBuffs.Bytes())
		}
	}
}

func Encrypt(key uint32, buf *[]byte) {
	length := len(*buf)
	for offset := 0; offset < length; offset += 4 {
		// fmt.Println("a", (*buf)[offset+0])
		// fmt.Println("a",(*buf)[offset+1])
		// fmt.Println("a",(*buf)[offset+2])
		// fmt.Println("a",(*buf)[offset+3])
		(*buf)[offset+0] ^= (byte)(key)
		(*buf)[offset+1] ^= (byte)(key >> 8)
		(*buf)[offset+2] ^= (byte)(key >> 16)
		(*buf)[offset+3] ^= (byte)(key >> 24)
		// fmt.Println("b",(*buf)[offset+0])
		// fmt.Println("b",(*buf)[offset+1])
		// fmt.Println("b",(*buf)[offset+2])
		// fmt.Println("b",(*buf)[offset+3])
		key = uint32(GetCRC32(buf, offset, 4))
	}
}

func Decrypt(key uint32, buf *[]byte) {
	length := len(*buf)
	for offset := 0; offset < length; offset += 4 {
		newkey := uint32(GetCRC32(buf, offset, 4))
		(*buf)[offset+0] ^= (byte)(key)
		(*buf)[offset+1] ^= (byte)(key >> 8)
		(*buf)[offset+2] ^= (byte)(key >> 16)
		(*buf)[offset+3] ^= (byte)(key >> 24)
		key = newkey
	}
}
