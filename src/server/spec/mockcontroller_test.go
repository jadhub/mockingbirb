package spec

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"testing"

	"flamingo.me/flamingo/v3/framework/flamingo"
	"flamingo.me/flamingo/v3/framework/web"
	configApplication "go.aoe.com/mockingbirb/src/config/application"
	configDomain "go.aoe.com/mockingbirb/src/config/domain"
	"go.aoe.com/mockingbirb/src/server/interfaces/controller"
)

type (
	DemoConfigProvider struct {
		JsonConfigProvider configApplication.JsonConfigProvider
	}

	testRequestInfo struct {
		URI    string
		Method string
		GET    map[string]string
		POST   map[string][]string
	}
)

func (p *DemoConfigProvider) GetConfigTree() configDomain.ConfigTree {
	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	}
	//exPath := filepath.Dir(dir)

	return p.LoadConfig(dir + "/test_json_config/config.json")
}

func (p *DemoConfigProvider) LoadConfig(path string) configDomain.ConfigTree {
	configTree := p.JsonConfigProvider.LoadConfig(path)

	return configTree
}

func TestMockController_MockAction(t *testing.T) {
	type fields struct {
		logger         flamingo.Logger
		responder      *web.Responder
		configProvider configDomain.ConfigProvider
	}
	type args struct {
		ctx context.Context
		req *web.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   web.Result
	}{
		{
			name: "greet me",
			fields: fields{
				logger:         flamingo.NullLogger{},
				responder:      &web.Responder{},
				configProvider: &DemoConfigProvider{},
			},
			args: args{
				ctx: context.Background(),
				req: createRequest(&testRequestInfo{
					URI:    "/greet/me",
					Method: "get",
					GET:    nil,
					POST:   nil,
				}),
			},
			want: &web.DataResponse{
				Response: web.Response{
					Status: uint(200),
					Body:   io.Reader(nil),
					Header: map[string][]string{
						"Content-Type": {"", "text/plain; charset=utf-8"},
					},
				},
				Data: map[string]interface{}{
					"message": "Mockingbirb says hi!",
				},
			},
		}, {
			name: "get json",
			fields: fields{
				logger:         flamingo.NullLogger{},
				responder:      &web.Responder{},
				configProvider: &DemoConfigProvider{},
			},
			args: args{
				ctx: context.Background(),
				req: createRequest(&testRequestInfo{
					URI:    "/get/json",
					Method: "get",
					GET:    nil,
					POST:   nil,
				}),
			},
			want: &web.DataResponse{
				Response: web.Response{
					Status: uint(200),
					Body:   io.Reader(nil),
					Header: map[string][]string{
						"Content-Type": {"", "application/json; charset=utf-8"},
					},
				},
				Data: map[string]interface{}{
					"foo": map[string]interface{}{
						"key1": float64(1),
						"key2": true,
						"key3": []interface{}{
							map[string]interface{}{
								"foo": "foo",
								"bar": true,
								"baz": []interface{}{
									float64(1),
									float64(2),
									"3",
								},
							},
						},
					},
				},
			},
		}, {
			name: "get xml",
			fields: fields{
				logger:         flamingo.NullLogger{},
				responder:      &web.Responder{},
				configProvider: &DemoConfigProvider{},
			},
			args: args{
				ctx: context.Background(),
				req: createRequest(&testRequestInfo{
					URI:    "/get/xml",
					Method: "get",
					GET:    nil,
					POST:   nil,
				}),
			},
			want: &web.Response{
				Status: uint(200),
				Header: map[string][]string{
					"Content-Type": {"", "application/xml; charset=utf-8"},
				},
				Body: bytes.NewBufferString("<test><default><a>1</a><b>1</b></default></test>"),
			},
		}, {
			name: "specific get param",
			fields: fields{
				logger:         flamingo.NullLogger{},
				responder:      &web.Responder{},
				configProvider: &DemoConfigProvider{},
			},
			args: args{
				ctx: context.Background(),
				req: createRequest(&testRequestInfo{
					URI:    "/get/specific_get_param",
					Method: "get",
					GET: map[string]string{
						"test123": "123",
					},
					POST: nil,
				}),
			},
			want: &web.Response{
				Status: uint(200),
				Header: map[string][]string{
					"Content-Type": {"", "text/plain; charset=utf-8"},
				},
				Body: bytes.NewBufferString("successful"),
			},
		}, {
			name: "specific post param",
			fields: fields{
				logger:         flamingo.NullLogger{},
				responder:      &web.Responder{},
				configProvider: &DemoConfigProvider{},
			},
			args: args{
				ctx: context.Background(),
				req: createRequest(&testRequestInfo{
					URI:    "/post/specific_post_param",
					Method: "post",
					GET:    nil,
					POST: map[string][]string{
						"test123": {"123"},
					},
				}),
			},
			want: &web.Response{
				Status: uint(200),
				Header: map[string][]string{
					"Content-Type": {"", "text/plain; charset=utf-8"},
				},
				Body: bytes.NewBufferString("successful"),
			},
		}, {
			name: "config not found",
			fields: fields{
				logger:         flamingo.NullLogger{},
				responder:      &web.Responder{},
				configProvider: &DemoConfigProvider{},
			},
			args: args{
				ctx: context.Background(),
				req: createRequest(&testRequestInfo{
					URI:    "/non/existing/uri",
					Method: "get",
					GET:    nil,
					POST: map[string][]string{
						"test123": {"123"},
					},
				}),
			},
			want: &web.Response{
				Status: uint(404),
				Header: map[string][]string{
					"Content-Type": {"", "text/plain; charset=utf-8"},
				},
				Body: bytes.NewBufferString("Mockingbirb config not found for this request"),
			},
		}, {
			name: "parameter mismatch",
			fields: fields{
				logger:         flamingo.NullLogger{},
				responder:      &web.Responder{},
				configProvider: &DemoConfigProvider{},
			},
			args: args{
				ctx: context.Background(),
				req: createRequest(&testRequestInfo{
					URI:    "/post/specific_post_param",
					Method: "post",
					GET:    nil,
					POST: map[string][]string{
						"not_configured_param": {"123"},
					},
				}),
			},
			want: &web.Response{
				Status: uint(404),
				Header: map[string][]string{
					"Content-Type": {"", "text/plain; charset=utf-8"},
				},
				Body: bytes.NewBufferString("Mockingbirb config found, but param mismatch"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &controller.MockController{
				Logger:         tt.fields.logger,
				Responder:      tt.fields.responder,
				ConfigProvider: tt.fields.configProvider,
			}

			got := c.MockAction(tt.args.ctx, tt.args.req)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockController.MockAction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createRequest(requestInfo *testRequestInfo) *web.Request {
	fakeSession := &web.Session{}

	query := ""
	for key, value := range requestInfo.GET {
		query = query + key + "=" + value
	}

	request := &http.Request{
		Method: requestInfo.Method,
		URL: &url.URL{
			Path:     requestInfo.URI,
			RawQuery: query,
		},
	}

	request.PostForm = url.Values{}
	for key, values := range requestInfo.POST {
		for _, value := range values {
			request.PostForm.Add(key, value)
		}
	}

	return web.CreateRequest(request, fakeSession)
}
