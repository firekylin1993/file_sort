package handle

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"io"
	"mytest/file_sort/model"
	"os"
	"path/filepath"
	"time"
)

var fetchByte = []byte("{\"name\": \"张三\"")

type MyReader struct {
	readDir string
	dm      *model.DateModel
}

func NewMyReader(rd string, dm *model.DateModel) *MyReader {
	return &MyReader{
		readDir: rd,
		dm:      dm,
	}
}

// ReadFile 读取文件内容，存储张三数据到db
func (mr *MyReader) ReadFile(ctx context.Context, fileName string) error {
	fmt.Println(fileName + "开始读取")
	filePath := filepath.Join(mr.readDir, fileName)
	f, err := os.Open(filePath)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("%s 文件打开失败", mr.readDir))
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	fmt.Println(fileName + "开始写入")

	for {
		line, _, c := reader.ReadLine()
		if c == io.EOF {
			break
		}

		//抓取数据前17个字节，对比数据是否为张山
		i := line[0:17]
		if !bytes.Equal(i, fetchByte) {
			continue
		}
		result := gjson.Get(string(line), "timestamp")
		if result.Int() == 0 {
			continue
		}
		day := time.Unix(0, result.Int()*int64(time.Millisecond)).Format("2006-02-01")
		mr.dm.Add(ctx, &model.ContextModel{ // 存db
			Name:      "张三",
			Date:      day,
			Msg:       line,
			Timestamp: result.Int(),
		})
	}
	fmt.Println(fileName + "写入完毕")
	return nil
}
