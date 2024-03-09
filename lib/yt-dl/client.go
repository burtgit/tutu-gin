package yt_dl

import (
	"context"
	"github.com/juju/errors"
	"os/exec"
	"regexp"
	"strings"
	"time"
	"tutu-gin/core/global"
)

func Client(pageLink string) (string, error) {

	// 创建一个上下文对象，并设置超时时间为3分钟
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	var cmd *exec.Cmd

	// 测试准备用
	if len(global.SERVICE_CONFIG.YtDl.Path) <= 0 {
		global.SERVICE_CONFIG.YtDl.Path = "yt_dlp"
	}

	if strings.Contains(global.SERVICE_CONFIG.YtDl.Path, "yt_dlp") {
		// 创建一个命令对象来执行Python脚本python3 -m
		cmd = exec.CommandContext(ctx, "/usr/bin/python3", "-m", global.SERVICE_CONFIG.YtDl.Path, "--dump-json", pageLink)
		cmd.Dir = "/Users/lihuanjie/Sites/python/yt-dlp"
	} else {
		// 创建一个命令对象来执行Python脚本
		cmd = exec.CommandContext(ctx, "/usr/local/src/yt-dlp_linux", "--dump-json", pageLink)
	}

	// 设置标准输出和标准错误输出
	output, err := cmd.Output()

	// 检查错误类型，如果是超时错误则不等待
	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		return "", errors.Annotate(errors.New("Python脚本执行超时"), "Python脚本执行超时")
	}

	// 处理其他错误
	if err != nil {
		return "", errors.Annotate(err, "Python执行失败")
	}

	// 定义 JSON 提取的正则表达式
	regex := regexp.MustCompile(`\{.*\}`)

	// 使用正则表达式提取 JSON 内容
	result := regex.FindString(strings.TrimSpace(string(output)))

	return result, nil
}
