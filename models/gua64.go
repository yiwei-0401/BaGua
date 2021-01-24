package model

import (
	"bytes"
)

type Ben_gua struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Guaci string `json:"guaci"`
	GuaciFanyi string `json:"guaci_fanyi"`
	Yao_1 string `json:"yao_1"`
	Yao_2 string `json:"yao_2"`
	Yao_3 string `json:"yao_3"`
	Yao_4 string `json:"yao_4"`
	Yao_5 string `json:"yao_5"`
	Yao_6 string `json:"yao_6"`
	Code string `json:"code"`
	Alias string `json:"alias"`
	Location string `json:"location"`
	Form string `json:"form"`
	BenguaCode bytes.Buffer
	BianYaos []int
}

//定义变卦结构体
type Bian_gua struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Guaci string `json:"guaci"`
	GuaciFanyi string `json:"guaci_fanyi"`
	Yao_1 string `json:"yao_1"`
	Yao_2 string `json:"yao_2"`
	Yao_3 string `json:"yao_3"`
	Yao_4 string `json:"yao_4"`
	Yao_5 string `json:"yao_5"`
	Yao_6 string `json:"yao_6"`
	Code string `json:"code"`
	Alias string `json:"alias"`
	Location string `json:"location"`
	Form string `json:"form"`
	BianguaCode bytes.Buffer
}
