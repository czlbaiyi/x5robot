package tools

import(
	"bytes"
	"encoding/binary"
)

//GenTokenID ...
func GenTokenID(token_string string) int64 {
	temp_pstid := int64(0)
	i_buffs := bytes.NewBuffer([]byte{})
	binary.Write(i_buffs, binary.BigEndian, temp_pstid)
	i_bytes := i_buffs.Bytes()

	all_bytes := []byte(token_string)
	i_length := len(i_bytes);
	a_lenth := len(all_bytes);
	for i := 0; i< a_lenth; i++{
		i_bytes[i_length - i - 1] = all_bytes[i];
	}

	all_buffs := bytes.NewBuffer(i_bytes)
	binary.Read(all_buffs, binary.LittleEndian, &temp_pstid)
	return temp_pstid;
}