package main

import (
	"fmt"

	"github.com/itfantasy/gonode/behaviors/gen_server"
	"github.com/itfantasy/gonode/utils/ini"
	"github.com/itfantasy/gonode/utils/io"

	"github.com/itfantasy/gonode-icloud/icloud/behaviors/lobby"
)

func main() {
	if nodeInfo, err := setupConfig(); err != nil {
		fmt.Println("Launch Error:" + err.Error())
	} else {
		Launch(nodeInfo)
	}
}

func setupConfig() (*gen_server.NodeInfo, error) {
	conf, err := ini.Load(io.CurDir() + ".conf")
	if err != nil {
		return nil, err
	}

	nodeInfo := gen_server.NewNodeInfo()

	nodeInfo.Id = conf.Get("node", "id")
	nodeInfo.Url = conf.Get("node", "url")
	nodeInfo.Pub = conf.GetInt("node", "pub", 0) > 0
	nodeInfo.BackEnds = conf.Get("node", "backends")
	nodeInfo.LogLevel = conf.Get("log", "loglevel")
	nodeInfo.LogComp = conf.Get("log", "logcomp")
	nodeInfo.RegComp = conf.Get("reg", "regcomp")

	redisConf := conf.Get("comps", "redis")
	if err := lobby.RegisterCoreRedis(redisConf); err != nil {
		return nil, err
	}

	return nodeInfo, nil
}
