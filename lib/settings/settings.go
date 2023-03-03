package settings

import (
	"github.com/getevo/evo-min/lib/generic"
	"github.com/getevo/evo-min/lib/settings/yml"
)

var instances = []Interface{&yml.Yaml{}}
var Instance = proxy{}

type Interface interface {
	Get(key string) generic.Value
	Has(key string) (bool, generic.Value)
	All() map[string]generic.Value
	Set(key string, value interface{}) error
	SetMulti(data map[string]interface{}) error
	Register(settings ...interface{}) error
	Init(params ...string) error
}

type Setting struct {
	Domain      string `gorm:"column:domain;primaryKey" json:"domain"`
	Name        string `gorm:"column:name;primaryKey" json:"name"`
	Title       string `gorm:"column:title" json:"title"`
	Description string `gorm:"column:description" json:"description"`
	Value       string `gorm:"column:value" json:"value"`
	Type        string `gorm:"column:type" json:"type"`
	Params      string `gorm:"column:params" json:"params"`
	ReadOnly    bool   `gorm:"column:read_only" json:"read_only"`
	Visible     bool   `gorm:"column:visible" json:"visible"`
}

func SetInterface(i Interface) {
	instances = append(instances, i)
}

func Get(key string) generic.Value {
	return Instance.Get(key)
}

func Has(key string) (bool, generic.Value) {
	return Instance.Has(key)
}

func All() map[string]generic.Value {
	return Instance.All()
}

func Set(key string, value interface{}) error {
	return Instance.Set(key, value)
}

func SetMulti(data map[string]interface{}) error {
	return Instance.SetMulti(data)
}

func Register(settings ...interface{}) {
	Instance.Register(settings...)
}

func Init(params ...string) error {
	return Instance.Init(params...)
}
