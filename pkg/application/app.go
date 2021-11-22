package application

import (
	"context"
)

type App struct {
	ctx     context.Context
	appInfo appInfo
}

type appInfo struct {
	id       string
	name     string
	version  string
	metadata map[string]string
}

func (a App) Run() error {
	var (
	//eg, ctx = errgroup.WithContext(a.ctx)
	//wg      = sync.WaitGroup{}
	)

	// 启动服务 cmd or serve

	// 注册 和 处理 signal

	return nil
}

func (a App) Stop() error {
	return nil
}

func (a App) ForceStop() error {
	return nil
}
