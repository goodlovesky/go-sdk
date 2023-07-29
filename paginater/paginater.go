/*
 * @Description:
 * @Author: goodlovesky
 * @Date: 2023-07-29 11:18:08
 */

package paginater

import (
	"gorm.io/gorm"
	"math"
)

type PageData struct {
	Page      int // 当前页
	PageCount int // 总页数
	Data      interface{}
}

// PageOperation 普通分页查询
// data是指针,page从1开始,_db调用model指定表和where指定条件
func PageOperation(page int, _db *gorm.DB, limit int, data interface{}) (PageData, error) {
	var (
		count    int64
		pageData PageData
		err      error
	)

	db := _db

	// 总页数
	err = db.Count(&count).Error
	if err != nil {
		return pageData, err
	}
	pageCount := int(math.Ceil(float64(count) / float64(limit)))

	// 获取分页数据
	err = db.Offset((page - 1) * limit).Limit(limit).Find(data).Error
	if err != nil {
		return pageData, err
	}

	return PageData{
		Data:      data,
		Page:      page,
		PageCount: pageCount,
	}, nil
}
