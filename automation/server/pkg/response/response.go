/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo licensed under the Mulan Permissive Software License, Version 2.
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Wed Sep 27 17:35:12 2023 +0800
 */
package response

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data interface{}, msg string) {
	result(c, http.StatusOK, http.StatusOK, data, msg)
}
func Fail(c *gin.Context, data interface{}, msg string) {
	result(c, http.StatusOK, http.StatusBadRequest, data, msg)
}

// 拼装json 分页数据
func DataPaged(c *gin.Context, list interface{}, total int, query *PagedQuery) {
	c.JSON(http.StatusOK, gin.H{
		"code":         http.StatusOK,
		"data":         list,
		"total":        total,
		"current_page": query.CurrentPage,
		"page_size":    query.PageSize})
}

func result(c *gin.Context, httpStatus int, code int, data interface{}, msg string) {
	c.JSON(httpStatus, gin.H{
		"code":    code,
		"data":    data,
		"message": msg})
}

type PagedQuery struct {
	PageSize    int `form:"page_size" json:"page_size"`
	CurrentPage int `form:"current_page" json:"current_page"`
	Total       int `json:"total"`
}

// 结构体分页查询方法
func StructDataPaged(p *PagedQuery, list interface{}, total int) (interface{}, error) {
	data := make([]interface{}, 0)
	if reflect.TypeOf(list).Kind() == reflect.Slice {
		s := reflect.ValueOf(list)
		for i := 0; i < s.Len(); i++ {
			ele := s.Index(i)
			data = append(data, ele.Interface())
		}
	}
	if p.PageSize < 1 {
		p.PageSize = 10
	}
	if p.CurrentPage < 1 {
		p.CurrentPage = 1
	}
	if total == 0 {
		p.Total = 0
	}
	num := p.PageSize * (p.CurrentPage - 1)
	if num > total {
		return nil, fmt.Errorf("页码超出")
	}
	if p.PageSize*p.CurrentPage > total {
		return data[num:], nil
	} else {
		if p.PageSize*p.CurrentPage < num {
			return nil, fmt.Errorf("读取错误")
		}
		if p.PageSize*p.CurrentPage == 0 {
			return data, nil
		} else {
			return data[num : p.CurrentPage*p.PageSize], nil
		}
	}
}
