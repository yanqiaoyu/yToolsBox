package dspm

import (
	"context"
	"io"
	"net/http"
	"time"
)

const defaultHttpTimeout = 3 * time.Second

/**
 * HttpGet
 * @Description: http get请求封装
 * @param:       url  请求url
 * @param:       jar 请求参数
 * @return:      http请求数据  error错误信息
 */
func HttpGet(url string) ([]byte, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultHttpTimeout)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, 0, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, 0, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, response.StatusCode, err
	}

	return body, response.StatusCode, nil
}

/**
 * HttpPatch
 * @Description: http patch请求封装
 * @param:       url  请求url
 * @return:      http请求数据  error错误信息
 */
func HttpPatch(url string) ([]byte, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultHttpTimeout)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodPatch, url, nil)
	if err != nil {
		return nil, 0, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, 0, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, response.StatusCode, nil
}
