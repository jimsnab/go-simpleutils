package simpleutils

//DeepCopy generates a separate copy of a source object
func DeepCopy(src interface{}) (dest interface{}) {

	switch val := src.(type) {
	case map[string]interface{}:
		m := make(map[string]interface{})
		for k, v := range val {
			m[k] = DeepCopy(v)
		}
		dest = m

	case []interface{}:
		a := make([]interface{}, 0, len(val))
		for _, v := range val {
			a = append(a, DeepCopy(v))
		}
		dest = a

	default:
		dest = val
	}

	return
}
