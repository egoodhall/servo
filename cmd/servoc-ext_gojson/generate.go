package main

import (
	"fmt"
	"github.com/egoodhall/servo/pkg/ipc"
)

func (x *GoJsonPlugin) Generate(req *ipc.GenerateRequest, res *ipc.GenerateResponse) error {
	fmt.Printf("%+v\n", req)
	return nil
}
