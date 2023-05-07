package config

import (
	"ids/snowflake"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

var (
	ServerListenAddr string          // 监听端口地址
	SnowflakeNodeId  int64           // 雪花算法节点id
	SnowflakeNode    *snowflake.Node // 雪花算法节点
)

// InitConfig init config
func InitConfig() {
	v := viper.New()
	v.AutomaticEnv()
	v.BindEnv("listen_addr", "node_id")
	v.SetDefault("listen_addr", ":8080")

	ServerListenAddr = v.GetString("listen_addr")

	// init snowflake node
	nodeId := getNodeId(v)
	node, err := snowflake.NewNode(nodeId)
	if err != nil {
		log.Fatalf("snowflake init fatal: %+v", err)
	}
	SnowflakeNodeId = nodeId
	SnowflakeNode = node
}

func getNodeId(v *viper.Viper) int64 {
	key := "node_id"
	if v.InConfig(key) {
		return v.GetInt64(key)
	}
	nodeId := getNodeIdViaKubernetesHostname()
	if nodeId > 0 {
		log.Printf("snowflake node id(k8s): %d", nodeId)
		return nodeId
	}
	// rand node id
	nodeMax := -1 ^ (-1 << snowflake.NodeBits)
	nodeId = rand.Int63n(int64(nodeMax + 1))
	log.Printf("snowflake node id(rand): %d", nodeId)
	return nodeId
}

func getNodeIdViaKubernetesHostname() int64 {
	hostname := os.Getenv("HOSTNAME")
	if hostname == "" {
		return -1
	}
	ary := strings.Split(hostname, "-")
	if len(ary) < 2 {
		return -1
	}
	nodeId, err := strconv.ParseInt(ary[len(ary)-1], 10, 64)
	if err != nil {
		return -1
	}
	return nodeId
}
