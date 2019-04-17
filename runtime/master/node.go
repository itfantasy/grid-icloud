package main

import (
	"strings"

	"github.com/itfantasy/gonode-icloud/icloud/behaviors/lobby"
	"github.com/itfantasy/gonode-icloud/icloud/logics/master"
	"github.com/itfantasy/gonode/behaviors/gen_server"
)

type GridNode struct {
	nodeInfo *gen_server.NodeInfo
}

func NewGridNode(nodeInfo *gen_server.NodeInfo) *GridNode {
	this := new(GridNode)
	this.nodeInfo = nodeInfo
	return this
}

func (this *GridNode) Setup() *gen_server.NodeInfo {

	redisConf, _ := this.nodeInfo.UserDatas["redisConf"]
	if err := lobby.RegisterCoreRedis(redisConf); err != nil {
		return nil
	}

	defaultRoomUrl, _ := this.nodeInfo.UserDatas["defaultRoomUrl"]
	master.SetDefaultRoomUrl(defaultRoomUrl)

	return this.nodeInfo
}

func (this *GridNode) Start() {

}
func (this *GridNode) OnConn(id string) {

}
func (this *GridNode) OnMsg(id string, msg []byte) {
	if strings.Index(id, "room") == 0 {
		master.HandleServerMsg(id, msg)
	} else {
		master.HandleMsg(id, msg)
	}
}
func (this *GridNode) OnClose(id string) {

}
