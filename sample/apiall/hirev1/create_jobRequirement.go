// Package hire code generated by oapi sdk gen
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
	"github.com/larksuite/oapi-sdk-go/v3/service/hire/v1"
)

// POST /open-apis/hire/v1/job_requirements
func main() {
	// 创建 Client
	client := lark.NewClient("appID", "appSecret")
	// 创建请求对象
	req := larkhire.NewCreateJobRequirementReqBuilder().
		UserIdType("open_id").
		DepartmentIdType("open_department_id").

		JobRequirement(larkhire.NewJobRequirementBuilder().
			ShortCode("xx1").
			Name("test").
			DisplayProgress(1).
			HeadCount(11).
			RecruitmentTypeId("1618209327096").
			EmployeeTypeId("6807409776231254285").
			MaxLevelId("123").
			MinLevelId("11").
			SequenceId("111").
			Category(1).
			DepartmentId("1111").
			RecruiterIdList([]string{}).
			JrHiringManagerIdList([]string{}).
			DirectLeaderIdList([]string{}).
			StartTime("1625729379000").
			Deadline("1625729379000").
			Priority(1).
			RequiredDegree(1).
			MaxSalary("123").
			MinSalary("11").
			AddressId("11").
			Description("11").
			CustomizedDataList([]*larkhire.JobRequirementCustomizedData{larkhire.NewJobRequirementCustomizedDataBuilder().Build()}).
			ProcessType(1).
			JobTypeId("6930815272790114324").
			Build()).
		Build()
	// 发起请求
	resp, err := client.Hire.V1.JobRequirement.Create(context.Background(), req)

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
