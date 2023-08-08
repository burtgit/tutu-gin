package web

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strconv"
	"strings"
	"time"

	"tutu-gin/record/recordDomain/values"

	"tutu-gin/core/global"

	"github.com/gin-gonic/gin"
	"github.com/juju/errors"

	"tutu-gin/core/api"
	"tutu-gin/core/exception"
	"tutu-gin/http/validator/webValidator"
	"tutu-gin/parser/parserApplicaition"
)

type WebParse struct{}

func (w *WebParse) Parse(c *gin.Context) {
	var requestData webValidator.WebParseValidator

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.Error(exception.ValidatorError(errors.Annotate(err, exception.API_PARAMETER_CHECK_FAIL)))
		return
	}

	day, _ := strconv.ParseInt(time.Now().Format("20060102"), 10, 64)
	getInt64 := c.GetInt64("userId")

	var apply values.Apply

	global.DB.Where("user_id = ?", getInt64).Where("platform = ?", "zhishuzhan").Where("date = ?", day).Get(&apply)

	if apply.Id <= 0 {
		apply.Platform = "zhishuzhan"
		apply.UserId = getInt64
		apply.Date = day
		apply.TotalTimes = 1
		_, _ = global.DB.Insert(&apply)
	} else {
		global.DB.ID(apply.Id).Incr("total_times").Cols("total_times").Update(&apply)
	}

	parserService := parserApplicaition.NewParserService()
	result, err := parserService.Parse(requestData.PageUrl, c.ClientIP(), getInt64)
	if err != nil {
		c.Error(exception.ValidatorError(errors.Annotate(err, exception.API_PARSER_FAIL)))
		return
	}
	global.DB.ID(apply.Id).Incr("success_times").Cols("success_times").Update(&apply)

	if result.IsVideo {
		// 加密
		// 计算 SHA256 哈希值
		secret := sha256.Sum256([]byte("lihuanjiehahha"))
		// 初始化向量
		iv := []byte("loserLiHuanJieaq")

		// 创建 AES-256-CBC 加密器
		block, err := aes.NewCipher(secret[:])
		if err != nil {
			panic(err)
		}
		mode := cipher.NewCBCEncrypter(block, iv)

		// 加密数据
		data := []byte(result.VideoUrls)
		padded := pkcs7Pad(data, aes.BlockSize)
		encrypted := make([]byte, len(padded))
		mode.CryptBlocks(encrypted, padded)

		// 将加密后的数据转换成 base64 编码
		encoded := base64.StdEncoding.EncodeToString(encrypted)

		// 替换 + 号为 ! 号
		encoded = replacePlus(encoded)

		result.EncodeUrl = "http://www.zanqianba.com/xzb/video_" + strconv.FormatInt(time.Now().Unix(), 10) + ".mp4?s=" + encoded
	}

	c.JSON(http.StatusOK, api.ApiSuccessResponse(result))
}

func (w *WebParse) Api(c *gin.Context) {
	var requestData webValidator.ApiParseValidator

	if err := c.ShouldBind(&requestData); err != nil {
		c.Error(exception.ValidatorError(errors.Annotate(err, exception.API_PARAMETER_CHECK_FAIL)))
		return
	}

	if requestData.Token != "9788678dced821353c6d881b3fde18cb" {
		c.Error(exception.ValidatorError(errors.Annotate(errors.New("token 有误"), exception.API_PARAMETER_CHECK_FAIL)))
		return
	}

	parserService := parserApplicaition.NewParserService()
	result, err := parserService.Parse(requestData.PageUrl, c.ClientIP(), 1)
	if err != nil {
		c.Error(exception.ValidatorError(errors.Annotate(err, exception.API_PARSER_FAIL)))
		return
	}

	c.JSON(http.StatusOK, api.ApiSuccessResponse(result))
}

func (w *WebParse) Agent(c *gin.Context) {
	var requestData webValidator.ApiParseValidator

	if err := c.ShouldBind(&requestData); err != nil {
		c.Error(exception.ValidatorError(errors.Annotate(err, exception.API_PARAMETER_CHECK_FAIL)))
		return
	}

	if requestData.Token != "9788678dced821353c6d881b3fde18cb" {
		c.Error(exception.ValidatorError(errors.Annotate(errors.New("token 有误"), exception.API_PARAMETER_CHECK_FAIL)))
		return
	}

	parserService := parserApplicaition.NewParserService()
	result, err := parserService.Agent(requestData.PageUrl)
	if err != nil {
		c.Error(exception.ValidatorError(errors.Annotate(err, exception.API_PARSER_FAIL)))
		return
	}

	c.JSON(http.StatusOK, api.ApiSuccessResponse(result))
}

// 使用 PKCS7 填充方式对数据进行填充
func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// 将字符串中的 + 号替换为 ! 号
func replacePlus(str string) string {
	return strings.ReplaceAll(str, "+", "!")
}

func NewWebParse() *WebParse {
	return &WebParse{}
}
