package utils

import "github.com/bwmarrin/snowflake"

var node *snowflake.Node

func InitSnowflake(machineId int64) {
	n, err := snowflake.NewNode(machineId)
	if err != nil {
		panic(err)
	}

	node = n
}

func GenerateId() int64 {
	return node.Generate().Int64()
}
