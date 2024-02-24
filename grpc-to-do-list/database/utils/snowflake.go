package utils

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

type SnowFlake struct {
	SnowFlakeNode *snowflake.Node
}

func NewSnowFlakeNode(machineId int64) *SnowFlake {
	node, err := snowflake.NewNode(machineId)
	if err != nil {
		fmt.Println(err)
	}

	return &SnowFlake{SnowFlakeNode: node}
}

//GetID 取得id編號
func (s *SnowFlake) GetID() int64 {

	id := s.SnowFlakeNode.Generate().Int64()

	return id
}
