// Package corehr code generated by oapi sdk gen
/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package main

import (
	"context"
	"fmt"
	"github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/corehr/v2"
)

// POST /open-apis/corehr/v2/probation/search
func main() {
	// 创建 Client
	client := lark.NewClient("appID", "appSecret")
	// 创建请求对象
	req := larkcorehr.NewSearchProbationReqBuilder().
		PageSize(100).
		PageToken("6891251722631890445").
		UserIdType("open_id").
		DepartmentIdType("open_department_id").
		Body(larkcorehr.NewSearchProbationReqBodyBuilder().
			EmploymentIds([]string{}).
			DepartmentIds([]string{}).
			ProbationStartDateStart("2022-05-18").
			ProbationStartDateEnd("2022-05-20").
			ProbationExpectedEndDateStart("2022-06-20").
			ProbationExpectedEndDateEnd("2022-07-20").
			ActualProbationEndDateStart("2022-08-20").
			ActualProbationEndDateEnd("2022-09-20").
			InitiatingTimeStart("2022-10-20").
			InitiatingTimeEnd("2022-11-20").
			ProbationStatus("approved").
			FinalAssessmentResult("approved").
			FinalAssessmentGrade("grade_a").
			Build()).
		Build()
	// 发起请求
	resp, err := client.Corehr.V2.Probation.Search(context.Background(), req)

	// 处理错误
	if err != nil {
		fmt.Println(err)
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		return
	}

	// 业务处理
	fmt.Println(larkcore.Prettify(resp))
}
