package tests

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 对于具备相同的测试逻辑的场景，使用简洁紧凑的 表组测试 来编写测试
func TestAllPages(t *testing.T) {

	baseURL := "http://localhost:3000"

	// 定义一个用于表组测试的结构体，其中要包含测试所需的输入与期望的输出
	// 1. 声明加初始化测试数据
	var tests = []struct {
		method   string // 请求方法
		url      string // URI
		expected int    // 状态码
	}{
		{"GET", "/", 200},
		{"GET", "/about", 200},
		{"GET", "/notfound", 404},
		{"GET", "/articles", 200},
		{"GET", "/articles/create", 200},
		{"GET", "/articles/3", 200},
		{"GET", "/articles/3/edit", 200},
		{"POST", "/articles/3", 200},
		{"POST", "/articles", 200},
		{"POST", "/articles/1/delete", 404},
	}

	// 2. 遍历所有测试
	for _, test := range tests {
		// 打印日志，支持 Printf 格式化打印
		t.Logf("当前请求 URL: %v \n", test.url)
		var (
			resp *http.Response
			err  error
		)
		// 2.1 请求以获取响应
		switch {
		case test.method == "POST":
			data := make(map[string][]string)
			resp, err = http.PostForm(baseURL+test.url, data)
		default:
			resp, err = http.Get(baseURL + test.url)
		}
		// 2.2 断言
		assert.NoError(t, err, "请求 "+test.url+" 时报错")
		assert.Equal(t, test.expected, resp.StatusCode, test.url+" 应返回状态码 "+strconv.Itoa(test.expected))
	}
}
