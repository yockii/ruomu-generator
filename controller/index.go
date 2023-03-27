package controller

import (
	"encoding/json"

	"github.com/yockii/ruomu-generator/constant"
)

func Dispatch(code string, headers map[string]string, value []byte) ([]byte, error) {
	switch code {
	// HTTP 注入点
	case constant.InjectCodeGenerateAll:
		return wrapCall(value, GeneratorController.generateAll)

	}
	return nil, nil
}

func wrapCall(v []byte, f func([]byte) (any, error)) ([]byte, error) {
	r, err := f(v)
	if err != nil {
		return nil, err
	}
	bs, err := json.Marshal(r)
	return bs, err
}
