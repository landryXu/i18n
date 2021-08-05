package i18n

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type I18n struct {
	Data  map[string][]string
	Index int //国际化语言下标
}

//index指定json文件的国际化语言下标
func NewI18n(path, fileName string, index ...int) (*I18n, error) {
	fileName = fmt.Sprintf("%s/%s.json", path, fileName)
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	i18n := I18n{}
	if len(index) == 0 { //默认只有一种外语
		i18n.Index = 0
	} else {
		i18n.Index = index[0]
	}
	if err = json.Unmarshal(bytes, &i18n.Data); err != nil {
		return nil, err
	}
	return &i18n, nil
}

func (i *I18n) T(key string, args ...interface{}) string {
	if _, ok := i.Data[key]; ok {
		key = i.Data[key][i.Index]
	}
	return i.preArgs(key, args...)
}

//use original text
func (i *I18n) TL(key string, args ...interface{}) string {
	return i.preArgs(key, args...)
}

func (i *I18n) preArgs(format string, args ...interface{}) string {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args...)
	}
	return format
}
