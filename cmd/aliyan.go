package cmd

import (
	"strings"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	vod20170321 "github.com/alibabacloud-go/vod-20170321/v3/client"
)

type AliSerice struct {
	client *vod20170321.Client
}

func NewAliSer(conf *AliConfig) (*AliSerice, error) {
	config := &openapi.Config{
		// Endpoint 请参考 https://api.aliyun.com/product/vod
		Endpoint: tea.String("vod.cn-shenzhen.aliyuncs.com"),
		// 必填，您的 AccessKey ID
		AccessKeyId: tea.String(conf.AccessKeyID),
		// 必填，您的 AccessKey Secret
		AccessKeySecret: tea.String(conf.AccessKeySecret),
	}

	client, err := vod20170321.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &AliSerice{
		client: client,
	}, nil
}

func (s *AliSerice) CreateUploadVideo(title, filename string) (*vod20170321.CreateUploadVideoResponseBody, error) {
	resp, err := s.client.CreateUploadVideo(&vod20170321.CreateUploadVideoRequest{
		Title: tea.String(title),
		FileName: tea.String(filename),
	})
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func (s *AliSerice) DeleteVideo(videoIds []string) error {
	_, err := s.client.DeleteVideo(&vod20170321.DeleteVideoRequest{
		VideoIds: tea.String(strings.Join(videoIds, ",")),
	})

	return err
}