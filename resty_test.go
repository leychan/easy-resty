package easyresty

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	// 创建一个临时的HTTP服务器来模拟请求
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, client")
    }))
    defer server.Close()

    // 使用临时服务器的URL进行测试
    url := server.URL

    // 调用Get函数并检查返回的响应
    resp, err := Get(url)
    if err != nil {
        t.Errorf("Expected no error, but got %v", err)
    }

    if string(resp.Body()) != "Hello, client\n" {
        t.Errorf("Expected body to be 'Hello, client', but got %s", string(resp.Body()))
    }
}


func TestNewRestyClientWithSetTimeout(t *testing.T) {
	client := NewRestyClient(WithTimeout(10*time.Second))
	if client == nil {
		t.Error("Expected a non-nil Resty client, but got nil")
	}
	assert.Equal(t, 10*time.Second, client.GetClient().Timeout)
}

func TestNewRestyClientWithAuthToken(t *testing.T) {
	client := NewRestyClient(WithAuthToken("token"))
	if client == nil {
		t.Error("Expected a non-nil Resty client, but got nil")
	}
	authToken := client.Token
	assert.Equal(t, "token", authToken)
} 

func TestNewRestyClientWithCookie(t *testing.T) {
	client := NewRestyClient(WithCookie(&http.Cookie{Name: "name", Value: "value"}))
	if client == nil {
		t.Error("Expected a non-nil Resty client, but got nil")
	}
	cookie := client.Cookies
	exist := false
	for _, c := range cookie {
		if c.Name != "name" {
			continue
		}
		exist = true
		assert.Equal(t, "name", c.Name)
	}

	if !exist {
		t.Error("Expected cookie with name 'name', but got none")
	}
}