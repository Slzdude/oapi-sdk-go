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
	"github.com/larksuite/oapi-sdk-go/v3/service/corehr/v1"
)

// PATCH /open-apis/corehr/v1/employments/:employment_id
func main() {
	// 创建 Client
	client := lark.NewClient("appID", "appSecret")
	// 创建请求对象
	req := larkcorehr.NewPatchEmploymentReqBuilder().
		EmploymentId("1616161616").
		ClientToken("12454646").
		UserIdType("people_corehr_id").
		DepartmentIdType("people_corehr_department_id").

		Employment(larkcorehr.NewEmploymentBuilder().
			SeniorityDate("2020-01-01").
			EmployeeNumber("1000000").
			EmploymentType(larkcorehr.NewEnumBuilder().Build()).
			PersonId("6919733936050406926").
			PrimaryEmployment(true).
			CustomFields([]*larkcorehr.ObjectFieldData{larkcorehr.NewObjectFieldDataBuilder().Build()}).
			WorkEmailList([]*larkcorehr.Email{larkcorehr.NewEmailBuilder().Build()}).
			ReasonForOffboarding(larkcorehr.NewEnumBuilder().Build()).
			AtsApplicationId("6838119494196871234").
			Build()).
		Build()
	// 发起请求
	resp, err := client.Corehr.V1.Employment.Patch(context.Background(), req)

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
