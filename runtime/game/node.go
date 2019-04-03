package main

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/itfantasy/gonode-icloud/icloud/logics/game"
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
	return this.nodeInfo
}

func (this *GridNode) Start() {

}
func (this *GridNode) OnDetect(id string) bool {
	if strings.Index(id, "lobby") == 0 {
		return true
	}
	return false
}
func (this *GridNode) OnConn(id string) {
	game.HandleConn(id)
}
func (this *GridNode) OnMsg(id string, msg []byte) {
	game.HandleMsg(id, msg)
}
func (this *GridNode) OnClose(id string) {
	game.HandleClose(id)
}
func (this *GridNode) OnShell(channel string, msg string) {

}
func (this *GridNode) OnRanId() string {
	return "cnt-" + strconv.Itoa(rand.Intn(100000))
}
