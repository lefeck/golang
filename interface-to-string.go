package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func Strval( value interface{}) string  {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		value := value.(float64)
		key = strconv.FormatFloat(value,'f', -1,64)
	case float32:
		value := value.(float32)
		key = strconv.FormatFloat(float64(value),'f',-1,64)
	case int:
		it := value.(int)
		key=strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key=strconv.Itoa(int(it))
	case int8:
		it:= value.(int8)
		key=strconv.Itoa(int(it))
	case uint8:
		it:= value.(uint8)
		key=strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key=strconv.Itoa(int(it))
	case uint16:
		it:= value.(uint16)
		key=strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key=strconv.Itoa(int(it))
	case uint32:
		it:= value.(uint32)
		key=strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key=strconv.Itoa(int(it))
	case uint64:
		it := value.(uint64)
		key=strconv.Itoa(int(it))
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}

func main()  {
	s := Strval(19.2)

	fmt.Printf("out type: %v\n",s)
	fmt.Printf("out type: %T\n",s)
}
