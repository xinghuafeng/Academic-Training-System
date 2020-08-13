package util

/*
	结果集数据类型转换
	[]map[string][]byte >> []map[string]interface{}
*/
func DataTrans(in []map[string][]byte) []map[string]interface{} {
	data := make([]map[string]interface{}, 0)
	for _, m := range in {
		e := make(map[string]interface{})
		for k, v := range m {
			str := string(v[:])
			e[k] = str
		}
		data = append(data, e)
		e = nil // 清除 map
	}
	return data
}
