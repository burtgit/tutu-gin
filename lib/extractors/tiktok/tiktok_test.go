package tiktok

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
	"tutu-gin/lib/extractors"

	"github.com/iawia002/lux/test"
)

func TestDownload(t *testing.T) {
	tests := []struct {
		name string
		args test.Args
	}{
		{
			name: "normal test 1",
			args: test.Args{
				URL:   "https://www.tiktok.com/@fatos.naturais/video/7277951260928560390",
				Title: "イケすぎたXOXO#xoxo #repezenfoxx #背中男 #kfam #yoshikiさんを泣かせたチーム @K fam @【Repezen Foxx】🦊",
				Size:  4356253,
			},
		},
		{
			name: "normal test 2",
			args: test.Args{
				URL:   "https://www.tiktok.com/@customize_extensions/video/7234132978287267090?is_from_webapp=1&sender_device=pc&web_id=7256777820008269355",
				Title: "깜짝 퇴장 👋 #ENHYPEN #SUNGHOON #NI_KI #Make_the_change",
				Size:  3848307,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := New().Extract(tt.args.URL, extractors.Options{})
			test.CheckError(t, err)
			t.Log(data)
		})
	}
}

func TestReuqest(t *testing.T) {
	url := "https://api16-normal-c-useast1a.tiktokv.com/aweme/v1/feed/?aweme_id=7344364446921313569&version_name=26.1.3&version_code=260103&build_number=26.1.3&manifest_version_code=260103&update_version_code=260103&openudid=3d111e550a57eb1a&uuid=3774239574718649&_rticket=1710078903000&ts=1710078903&device_brand=Google&device_type=Pixel+7&device_platform=android&resolution=1080%2A2400&dpi=420&os_version=13&os_api=29&carrier_region=US&sys_region=US&region=US&app_name=trill&app_language=en&language=en&timezone_name=America%2FNew_York&timezone_offset=-14400&channel=googleplay&ac=wifi&mcc_mnc=310260&is_my_cn=0&aid=1180&ssmix=a&as=a1qwert123&cp=cbfhckdckkde1"

	// 创建 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// 创建 GET 请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 设置请求的 User-Agent 和 Accept 头部
	req.Header.Set("User-Agent", "com.ss.android.ugc.trill/260103 (Linux; U; Android 13; en_US; Pixel 7; Build/TD1A.220804.031; Cronet/58.0.2991.0)")
	req.Header.Set("Accept", "application/json")

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求发送失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	// 输出响应状态码
	fmt.Println("响应状态码:", resp.StatusCode)

	// 输出响应内容
	fmt.Println("响应内容:", string(body))
}
