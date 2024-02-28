package kafka

import (
	"fmt"
	"net/http"

	"github.com/OrigamiWang/msd/gate/biz"
	"github.com/OrigamiWang/msd/gate/biz/lb"
	"github.com/OrigamiWang/msd/micro/const/mq"
	"github.com/OrigamiWang/msd/micro/mq/kafka"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/gin-gonic/gin"
)

var (
	g *gin.RouterGroup
	c *http.Client
)

func handleGateProxy(arr ...string) {
	svcName := arr[0]
	url := arr[1]

	// add in LB
	logutil.Info("add service %v to load balancer, url: %v", svcName, url)
	lb.LB.AddService(svcName, url)

	// proxy
	prefix := fmt.Sprintf("/%s", svcName)
	biz.Proxy(lb.LB, prefix, g, c)
}

func GateProxyConsumer(group *gin.RouterGroup, client *http.Client) {
	g = group
	c = client
	go kafka.ConsumeMsg(mq.KAFKA_GATE_PROXY, handleGateProxy)
}
