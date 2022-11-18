package main

import (
	"context"
	"file_sort/db"
	"file_sort/handle"
	"file_sort/kits"
	"file_sort/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/arl/statsviz"
)

var depStr = "./records"
var dstDir = "./result"

func main() {
	// 创建下载目录
	err := os.MkdirAll(dstDir, os.ModePerm)
	if err != nil {
		panic(err)
	}
	dir, err := ioutil.ReadDir(depStr)
	if err != nil {
		panic(err)
	}

	//初始化db
	mysql, err := db.NewMysql()
	if err != nil {
		panic(err)
	}
	dataModel := model.NewDataModel(mysql)
	ctx := context.Background()

	// 来一个reader读取数据
	reader := handle.NewMyReader(depStr, dataModel)
	rPool := kits.NewGoPool(30) // 读数据连接池
	for _, v := range dir {
		rPool.Add(1)
		go func(c context.Context, fileName string) {
			_ = reader.ReadFile(c, fileName)
			rPool.Done()
		}(ctx, v.Name())
	}
	rPool.Wait()

	// 来一个writer写入数据
	writer := handle.NewMyWriter(dstDir, dataModel)
	dateType, err := writer.GetDateType(ctx) //获取所需协程数
	if err != nil {
		panic(err)
	}
	wPool := kits.NewGoPool(len(dateType))
	for _, v := range dateType {
		wPool.Add(1)
		go func(c context.Context, date string) {
			_ = writer.WriteFile(c, date)
			wPool.Done()
		}(ctx, v.Date)
	}
	wPool.Wait()
	//关闭所有writer资源
	//handle.CloseSource()

	statsviz.RegisterDefault()
	log.Println(http.ListenAndServe(":8080", nil))
}
