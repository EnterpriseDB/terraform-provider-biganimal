package models

type KeyValues []KeyValue

func (kvs KeyValues) PropList() PropList {
	l := PropList{}
	for _, kv := range kvs {
		l = append(l, kv.Prop())
	}
	return l
}

type KeyValue struct {
	Name  string `json:"name" mapstructure:"name"`
	Value string `json:"value" mapstructure:"value"`
}

func (kv KeyValue) Prop() Prop {
	m := Prop{}
	m["name"] = kv.Name
	m["value"] = kv.Value
	return m
}
