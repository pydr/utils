package utils

import "reflect"

const MAX_DEPTH = 32

func JsonMerge(dst, src map[string]interface{}) map[string]interface{} {
	return jsMerge(dst, src, 0)
}

func jsMerge(dst, src map[string]interface{}, depth int) map[string]interface{} {
	if depth > MAX_DEPTH {
		return nil
	}

	if dst == nil {
		return src
	}

	for key, srcVal := range src {
		if dstVal, ok := dst[key]; ok {
			srcMap, srcMapOk := jsMapify(srcVal)
			dstMap, dstMapOk := jsMapify(dstVal)
			if srcMapOk && dstMapOk {
				srcVal = jsMerge(dstMap, srcMap, depth+1)
			}
		}
		dst[key] = srcVal
	}

	return dst
}

func jsMapify(i interface{}) (map[string]interface{}, bool) {
	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Map {
		m := map[string]interface{}{}
		for _, k := range value.MapKeys() {
			m[k.String()] = value.MapIndex(k).Interface()
		}
		return m, true
	}

	return map[string]interface{}{}, false
}
