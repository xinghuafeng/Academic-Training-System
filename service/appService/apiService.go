package appService

import (
	"sauth/model/appModel"
	"sauth/util"
)

func FindOptions() ([]map[string]interface{}, error) {
	res, err := appModel.FindOptions()
	if nil != err {
		return nil, err
	}
	data := util.DataTrans(res)
	return data, nil
}
