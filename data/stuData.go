package data

import "study/stuMange/model"

var Students = map[string]*model.Student{
	"1001": {
		Sid:   "1001",
		Sname: "zhangsan",
		Age:   20,
		Score: 98.6,
	},
	"1002": {
		Sid:   "1002",
		Sname: "lisi",
		Age:   22,
		Score: 90.6,
	},
	"1003": {
		Sid:   "1003",
		Sname: "wangwu",
		Age:   28,
		Score: 94.6,
	},
}
