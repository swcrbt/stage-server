package cmd

func Run() error {
	conf, err := GetConfig("configs/config.yaml")
	if err != nil {
		return err
	}

	aliSer, err := NewAliSer(conf.Aliyun)
	if err != nil {
		return err
	}

	// _, err = aliSer.CreateUploadVideo("UploadTest", "video_01.mp4")
	err = aliSer.DeleteVideo([]string{
		"0005b220bcee71eebfab7511979f0102",
	})
	if err != nil {
		return err
	}

	return nil
}