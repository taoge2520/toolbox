/*
 * Copyright (c) 2017,厦门团队
 * ALL Rights reserved
 *
 * 包    名：basic_op
 * 文 件 名：XXX.go
 * 作    者：LinTao
 * 时    间：2017/11/16 11:35
 *===版本修订记录===
 *定位具体修订位置方式: 文档内搜索"序号"
 *序号命名规则:[ADD|DEL|MOD][XX],XX代表序号,eg:MOD01 代表第1处修改的位置
 *序号-作者-时间-修订事项描述
 *
**/
package basic_op

import (
	"testing"
	"fmt"
)

func TestSliceOp(t *testing.T){
	data:=[]string{"one","two","","three"}
	data=Nonempty(data)
	fmt.Printf("new data is :%v\n",data)


	s := []int{5, 6, 7, 8, 9}
	s=Remove(s,2)
	fmt.Printf("new s is :%v",s)
}