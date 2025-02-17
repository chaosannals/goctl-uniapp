package generate

import (
	"fmt"

	"github.com/gookit/color"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/tools/goctl/api/parser"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

func UniAppCommand(p *plugin.Plugin) error {
	api, e := parser.Parse(p.ApiFilePath)
	if e != nil {
		return e
	}

	if err := api.Validate(); err != nil {
		return err
	}

	logx.Must(pathx.MkdirIfNotExist(p.Dir))
	logx.Must(genMessages(p.Dir, api))
	logx.Must(genClient(p.Dir, api))

	fmt.Println(color.Green.Render("Done."))
	return nil
}
