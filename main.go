package main

import (
	"ginblog_backend/global"
	"ginblog_backend/internal/model"
	"ginblog_backend/internal/routers"
	"ginblog_backend/pkg/setting"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting() error = %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine() error = %v", err)
	}
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:              ":" + global.ServerSetting.HttpPort,
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       global.ServerSetting.ReadTimeout,
		ReadHeaderTimeout: 0,
		WriteTimeout:      global.ServerSetting.WriteTimeout,
		IdleTimeout:       0,
		MaxHeaderBytes:    1 << 20,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	s.ListenAndServe()
}

func setupSetting() error {
	ASetting, err := setting.NewSettings()
	if err != nil {
		return err
	}
	err = ASetting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = ASetting.ReadSection("APP", &global.AppSetting)
	if err != nil {
		return err
	}
	err = ASetting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(&global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}
