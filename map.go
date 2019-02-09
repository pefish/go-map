package p_map

import (
	"gitee.com/pefish/p-go-slice"
	"reflect"
)

type MapClass struct {
}

var Map = MapClass{}

func (this *MapClass) Keys(in_ interface{}) (keys []string) {
	reflectKeys := reflect.ValueOf(in_).MapKeys()
	for _, v := range reflectKeys {
		keys = append(keys, v.String())
	}
	return
}

func (this *MapClass) Values(map_ map[string]interface{}) (values []interface{}) {
	for _, v := range map_ {
		values = append(values, v)
	}
	return
}

func (this *MapClass) Entries(map_ map[string]interface{}) (entries [][2]interface{}) {
	for k, v := range map_ {
		entries = append(entries, [2]interface{}{k, v})
	}
	return
}

func (this *MapClass) Length(map_ map[string]interface{}) int64 {
	return int64(len(map_))
}

func (this *MapClass) IsEmpty(map_ interface{}) bool {
	return len(map_.(map[string]interface{})) == 0
}

func (this *MapClass) Append(maps_ ...map[string]interface{}) (out map[string]interface{}) {
	out = map[string]interface{}{}
	for _, map_ := range maps_ {
		for key, val := range map_ {
			out[key] = val
		}
	}
	return
}

func (this *MapClass) Remain(map_ map[string]interface{}, fields []string) (out map[string]interface{}) {
	for key, _ := range map_ {
		if !p_slice.Slice.IncludesBySliceString(fields, key) {
			delete(map_, key)
		}
	}
	out = map_
	return
}
