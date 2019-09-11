package model

import "fmt"

type Student struct {
	Sname string  `json:"sname"`
	Sid   string  `json:"sid"`
	Age   uint8   `json:"age"`
	Score float32 `json:"score"`
}

func (this *Student) String() string {
	return fmt.Sprintf("Sname : %s, Sid : %s, Age : %d, Score : %f", this.Sname, this.Sid, this.Age, this.Score)
}
