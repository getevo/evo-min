package database

import (
	"fmt"
	"github.com/getevo/evo-min"
	"github.com/getevo/evo-min/lib/args"
	"github.com/getevo/evo-min/lib/generic"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
	"sync"
)

var db *gorm.DB

type Database struct {
	mu   sync.Mutex
	data map[string]map[string]generic.Value
}

func (config *Database) Get(key string) generic.Value {
	key = strings.ToUpper(key)
	var chunks = strings.SplitN(key, ".", 2)
	if len(chunks) != 2 {
		return generic.Parse("")
	}
	if domain, ok := config.data[chunks[0]]; ok {
		if v, exists := domain[chunks[1]]; exists {
			return v
		}
	}
	return generic.Parse("")
}
func (config *Database) Has(key string) (bool, generic.Value) {
	key = strings.ToUpper(key)
	var chunks = strings.SplitN(key, ".", 2)
	if len(chunks) != 2 {
		return false, generic.Parse("")
	}
	if domain, ok := config.data[chunks[0]]; ok {
		if v, exists := domain[chunks[1]]; exists {
			return true, v
		}
	}
	return false, generic.Parse("")
}
func (config *Database) All() map[string]generic.Value {
	var m = map[string]generic.Value{}
	for domain, inner := range config.data {
		for name, value := range inner {
			m[domain+"."+name] = value
		}
	}
	return m
}
func (config *Database) Set(key string, value interface{}) error {
	key = strings.ToUpper(key)
	var chunks = strings.SplitN(key, ".", 2)
	if len(chunks) == 2 {
		db.Where("domain = ? AND name = ?", chunks[0], chunks[1]).Model(Setting{}).Update("value", value)
	}
	return nil
}
func (config *Database) SetMulti(data map[string]interface{}) error {
	for key, value := range data {
		key = strings.ToUpper(key)
		var chunks = strings.SplitN(key, ".", 2)
		if len(chunks) == 2 {
			db.Where("domain = ? AND name = ?", chunks[0], chunks[1]).Model(Setting{}).Update("value", value)
		}
	}
	return nil
}
func (config *Database) Register(settings ...interface{}) error {
	for _, s := range settings {
		var v = generic.Parse(s)
		var setting = Setting{}
		var err = v.Cast(&setting)
		if err == nil {
			if !v.Is("settings.Setting") {
				return fmt.Errorf("invalid settings")
			}
			if ok, _ := config.Has(setting.Domain + "." + setting.Name); !ok {
				config.Set(setting.Domain+"."+setting.Name, setting.Value)
				db.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&setting)
			}
		} else {
			return err
		}
	}
	return nil
}
func (config *Database) Init(params ...string) error {
	config.mu.Lock()
	var items []Setting
	db = evo.GetDBO()
	if args.Exists("migrate") {
		db.AutoMigrate(&Setting{}, &SettingDomain{})
	}
	db.Find(&items)
	for _, item := range items {
		item.Domain = strings.ToUpper(item.Domain)
		item.Name = strings.ToUpper(item.Name)
		if _, ok := config.data[item.Domain]; !ok {
			config.data[item.Domain] = map[string]generic.Value{}
		}
		config.data[item.Domain][item.Name] = generic.Parse(item.Value)
	}
	config.mu.Unlock()
	return nil
}
