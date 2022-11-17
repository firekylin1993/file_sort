package handle

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"mytest/file_sort/model"
	"os"
	"path/filepath"
)

type MyWriter struct {
	writerDir string
	dm        *model.DateModel
}

func NewMyWriter(wd string, dm *model.DateModel) *MyWriter {
	return &MyWriter{
		writerDir: wd,
		dm:        dm,
	}
}

func (mw *MyWriter) GetDateType(ctx context.Context) ([]*model.ContextModel, error) {
	return mw.dm.GetDate(ctx)
}

// WriteFile 数据写入如文件并归类
func (mw *MyWriter) WriteFile(ctx context.Context, date string) error {
	filePath := filepath.Join(mw.writerDir, date+".txt")
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("%s 创建文件失败", date))
	}

	// 先查询所有当天日期的数据
	contextModel, err := mw.dm.GetContextByDate(ctx, date)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("%s 数据读取文件失败", date))
	}
	for _, v := range contextModel {
		bytes := append(v.Msg, []byte("\n")...)
		_, err := file.Write(bytes)
		if err != nil {
			return errors.WithMessage(err, fmt.Sprintf("%d 写入文件失败", v.ID))
		}
	}

	return nil
}
