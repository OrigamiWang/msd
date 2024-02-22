package transfer

import (
	"github.com/OrigamiWang/msd/micro/model"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/mitchellh/mapstructure"
)

func FacadeRespToMap(m map[string]interface{}, resp interface{}) (map[string]interface{}, error) {
	var modelResp *model.Response
	err := mapstructure.Decode(m, &modelResp)
	if err != nil {
		logutil.Error("unmarshal register config failed, err: %v", err)
		panic(err.Error())
	}
	result := modelResp.Data.(map[string]interface{})
	return result, nil
}
