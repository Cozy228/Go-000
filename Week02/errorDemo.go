package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	business(20200607010463)
	//another search
	business(-1)
}

func dao(id int) (string, error) {
	if id < 0 {
		//404
		//change comment
		//这里毛老师建议可以定义出自己的错误码如 code. ErrNoRows 解耦掉 依赖
		return "", errors.Wrapf(sql.ErrNoRows, "404 no data found with id %d", id)
	}
	//normal logic
	return "Ziyu", nil
}

//service
func findNameById(id int) (string, error) {
	return dao(id)
}

func business(id int) {
	name, err := findNameById(id)
	if err != nil {
		if errors.Is(errors.Cause(err), sql.ErrNoRows) {
			// handle error
			//Change comment:
			//errors.is 里面不用嵌套 errors.Cause(err) 了， 直接errors.is(err,sql.ErrNoRows)就行
			//errors.is能递归找根因
			fmt.Printf("Error：%T %v\n", errors.Cause(err), errors.Cause(err))
			fmt.Printf("Stacktrace：\n%+v\n", err)
		} else {
			// other errors
			fmt.Printf("some other error: %+v\n", err)
		}
		return
	}
	//handle normal logic
	fmt.Println("handle normal logic")
	fmt.Printf("Name: %s\n", name)
}
