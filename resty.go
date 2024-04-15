package easyresty

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

// Option 定义了配置请求的函数类型
type ClientOption func (*resty.Client)

// WithQueryParam 是一个函数，用来设置请求的查询参数
func WithQueryParam(key, value string) ClientOption {
	return func(c *resty.Client) {
		c.SetQueryParam(key, value)
	}
}

// WithQueryParams 是一个函数，用来批量设置请求的查询参数
func WithQueryParams(params map[string]string) ClientOption {
    return func(c *resty.Client) {
        c.SetQueryParams(params)
    }
}

// WithFormData 是一个函数，用来设置请求的表单数据
func WithFormData(params map[string]string) ClientOption {
    return func(c *resty.Client) {
        c.SetFormData(params)
    }
}

// WithRequestBody 是一个函数，用来设置请求的请求体
func WithRequestBody(body any) ClientOption {
    return func(c *resty.Client) {
        c.R().SetBody(body)
    }
}

// WithHeader 是一个函数，用来设置请求的头部
func WithHeader(key, value string) ClientOption {
	return func(c *resty.Client) {
		c.SetHeader(key, value)
	}
}

// WithHeaders 是一个函数，用来批量设置请求的头部
func WithHeaders(headers map[string]string) ClientOption {
    return func(c *resty.Client) {
        c.SetHeaders(headers)
    }
}

// WithAuthToken 是一个函数，用来设置请求的认证
func WithAuthToken(authToken string) ClientOption {
    return func(c *resty.Client) {
        c.SetAuthToken(authToken)
    }
}

// WithBasicAuth 是一个函数，用来设置请求的基本认证
func WithBasicAuth(username, password string) ClientOption {
    return func(c *resty.Client) {
        c.SetBasicAuth(username, password)
    }
}

// WithProxy 是一个函数，用来设置代理 example: localhost:8080
func WithProxy(proxy string) ClientOption {
    return func(c *resty.Client) {
        c.SetProxy(proxy)
    }
}

// WithTimeout 是一个函数，用来设置请求的超时
func WithTimeout(timeout time.Duration) ClientOption {
    return func(c *resty.Client) {
        c.SetTimeout(timeout)
    }
}

// WithCookie 是一个函数，用来设置请求的Cookie
func WithCookie(cookie *http.Cookie) ClientOption {
    return func(c *resty.Client) {
        c.SetCookie(cookie)
    }
}

// WithCookies 是一个函数，用来批量设置请求的Cookie
func WithCookies(cookies []*http.Cookie) ClientOption {
    return func(c *resty.Client) {
        c.SetCookies(cookies)
    }
}

// NewDefalutRestyClient 是一个函数，用来创建默认的Resty客户端
func NewDefalutRestyClient() *resty.Client {
    client := resty.New()
    client.SetTimeout(10 * time.Second)
    return client
}

func NewRestyClient(opts ...ClientOption) *resty.Client {
    client := NewDefalutRestyClient()
    
    // 应用所有的选项
    for _, opt := range opts {
        opt(client)
    }
    return client
}

// Get 是执行请求的函数，接受一个URL和多个配置选项
func Get(url string, opts ...ClientOption) (*resty.Response, error) {
	client := NewRestyClient(opts...)

	// 发起GET请求
	resp, err := client.R().Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Post 是执行请求的函数，接受一个URL和多个配置选项
func Post(url string, opts ...ClientOption) (*resty.Response, error) {
    client := NewRestyClient(opts...)

    // 发起POST请求
    resp, err := client.R().Post(url)
    if err != nil {
        return nil, err
    }
    return resp, nil
}