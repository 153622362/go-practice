package model

import "encoding/json"

//数据模型结构
type Profile struct {
	Name 		string
	Gender 		string
	Age 		int
	Height 		int
	Income		string
	Marriage	string
	Education	string
	Occupation	string
	Hokou		string
	Xinzuo		string
	House		string
	Car			string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o) //将数据转成json
	if  err != nil{
		return profile, err
	}

	err = json.Unmarshal(s, &profile) //将json成数据结构上
	return  profile, err
}