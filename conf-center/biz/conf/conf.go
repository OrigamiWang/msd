package conf

import (
	"fmt"
	"github.com/OrigamiWang/msd/conf-center/model/dao"
	"github.com/OrigamiWang/msd/micro/confparser"
	httpmethod "github.com/OrigamiWang/msd/micro/const/http"
	"github.com/OrigamiWang/msd/micro/framework/client"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"gopkg.in/yaml.v3"
	"net/http"
)

func LoadConfCenter(svcName string) (*confparser.Config, error) {
	uri := fmt.Sprintf("/config/%s", svcName)
	resp, err := client.RequestWithHead(httpmethod.GET, "localhost:8084", uri, http.Header{}, nil)
	if err != nil {
		logutil.Error("request with head failed, err: %v", err)
		return nil, err
	} else {
		confStr := resp.(*dao.SvcConfig).Conf
		config := &confparser.Config{}
		err = yaml.Unmarshal([]byte(confStr), config)
		if err != nil {
			logutil.Error("unmarshal config failed, err: %v", err)
			return nil, err
		}
		return config, nil
	}
}
