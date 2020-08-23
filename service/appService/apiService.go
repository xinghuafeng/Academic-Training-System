package appService

import (
	"Academic-Training-System/model/appModel"
	"Academic-Training-System/util"
)

func FindOptions() ([]map[string]interface{}, error) {
	res, err := appModel.FindOptions()
	if nil != err {
		return nil, err
	}
	data := util.DataTrans(res)
	return data, nil
}
