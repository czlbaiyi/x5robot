package x5base

const(
	PT_NotSupported = iota 
    PT_UInt32
    PT_Int32
    PT_UInt64
    PT_Int64
    PT_Float32
    PT_Float64
    PT_String
    PT_ByteArray
    PT_Time64
    PT_List
    PT_H3DDictionary
    PT_Compound
    PT_HashSet
)

func get_value_type(ft interface{}) int{
    switch ft.(type){
        case bool:
        case int32:
        case int16:
            return PT_Int32
    }
        return 0
}