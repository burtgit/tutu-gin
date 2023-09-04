/*!
 * site.css
 * (c) 2015-2018 iiiLab.com
 */
function parseSuffix(url) {
    var parser = document.createElement('a');
    parser.href = url;
    return parser.pathname.split('.').pop();
}

function isMP4File(url) {
    var suffix = parseSuffix(url);
    return suffix.toUpperCase() === "MP4";
}

//字符串转base64
function encode(str){
// 对字符串进行编码
    var encode = encodeURI(str);
// 对编码的字符串转化base64
    var base64 = btoa(encode);
    return base64;
}

Vue.use(VueClipboard);
var app = new Vue({
    delimiters: ["{", "}"],
    el: '#app',
    data: {
        link: '6.12 qrR:/ # 被好风景收买  https://v.douyin.com/ie2HkymX/ 复制此链接，打开Dou音搜索，直接观看视频！',
        errorTip: '',
        isMP4File: false,
        requestSuccess: false,
        showAllSupportLink: false,
        showClearBtn: false,
        showWxProfileKey:false,
        showWxProfileMinKey:false,
        message:"",
        token:"",
        vip_type:1,
        vip_times:0,
        vip_end_time:"",
        requestResult: {
            VideoUrls: '',
            EncodeUrl: '',
            Title: '',
            CoverUrls: '',
            Pics: [],
            IsVideo: true,
        },
        loginModal:{
            isReg: false,
            title: "登录",
            switchBtn: "没有账号？去注册",
            errorTip: '',
            phone: '',
            pwd: '',
            pwd2: ''
        },
        activationModal:{
            code: '',
            errorTip: ''
        },
        activationCode:{
            num: '',
            type: '',
            errorTip: ''
        },
        copyText:"复制视频下载链接",
        copyTimes:1,
        downloadProcess:0,
        downloadUrl :"",
        downloadHtml:""
    },
    watch: {
        'link': function(newVal,oldVal) {
            if (newVal.length > 0) {
                this.showClearBtn = true;
                $(".input-group-lg .link-input").css('padding-right', '32px');
            }else {
                this.showClearBtn = false;
                $(".input-group-lg .link-input").css('padding-right', '16px');
            }
        }
    },
    mounted: function() {
    },
    methods: {
        showWxProfile: function() {
            this.showWxProfileKey = !this.showWxProfileKey;
        },
        showWxProfileMin: function() {
            this.showWxProfileMinKey = !this.showWxProfileMinKey;
        },
        videoDownloadLink: function(t, e) {
            return void 0 === e && (e = "video"),
            "https://service0.iiilab.com/video/iiilab/" + e + "_iiilab_" + (new Date).getTime() + Math.random().toString(10).substring(2, 4) + ".mp4?source=" + encode(t)
        },
        _initGeeTest: function (captchaObj) {
            var that = this;

            captchaObj.onReady(function () {
                that.submitBtnClass.disabled = false;
                captchaObj.verify();
            }).onSuccess(function () {
                var result = captchaObj.getValidate();
                if (!result) {
                    console.log(1);
                    return;
                }

                that.parseVideo(result);
            });
            //window.gt = captchaObj;
        },
        buildVideoDownloadUrl: function (videoUrl) {
            return "/xzbs/video_" + new Date().getTime() + ".mp4?s=" + encodeURIComponent(CryptoJS.enc.Base64.stringify(CryptoJS.enc.Utf8.parse(videoUrl)));
        },
        submit: function (event) {
            console.log("13123123")
            //清空并隐藏之前的错误提示
            this.errorTip = "";

            $("#loadingModal").modal('show');

            //解析视频地址
            this.parseVideo({});
        },
        parseVideo: function(code) {
            //服务端处理
            this.requestSuccess = false;
            this.copyText = '复制视频下载链接';
            this.copyTimes = 1;

            var vm = this;
            $.ajax({
                type: 'POST',
                url: '/v1/parse/dana',
                // xhrFields:{withCredentials: true},
                // crossDomain: true,
                async: true,
                data: {
                    "pageUrl": vm.link,
                },
                dataType: "json",
                success: function(data) {
                    console.log(data)
                    if(data.Code == 201) {
                        vm.errorTip = data.Msg;
                    }else {
                        vm.requestResult = data.Data
                        vm.requestSuccess = true
                    }
                    console.log(vm.requestResult)
                },
                error: function () {
                    vm.errorTip = "处理失败,请重试!";
                },
                complete: function () {
                    $("#loadingModal").modal('hide');
                }
            });
        },
        clear: function() {
            this.requestSuccess = false;
            this.link = "";
        },
        doCopy: function() {
            var that = this;
            this.$copyText(this.requestResult.video).then(function (e) {
                that.copyText = "已复制链接(" + that.copyTimes + ")";
                console.log(e)
                that.copyTimes++;
            }, function (e) {
                alert('复制失败')
                console.log(e)
            }).catch(function(err) {
                console.log(12213);
                console.log(err);
            })
        },
        download: function (e) {
            var that = this
            var url = e.target.dataset.url,
                filename = e.target.dataset.name;
            $('#downloadModal').modal('show');
            that.downloadUrl = e.target.dataset.build
            that.downloadProcess = 0
            req = new XMLHttpRequest()
            req.open("get", url, true);
            req.setRequestHeader("Content-type", "application/x-www-form-urlencoded; charset=UTF-8");
            req.setRequestHeader("origin", "");
            //监听进度事件
            req.addEventListener("progress", function(evt) {
                if (evt.lengthComputable) {
                    var percentComplete = parseInt(evt.loaded / evt.total * 100);
                    console.log(percentComplete);
                    if (percentComplete == 100) {
                        that.downloadProcess = 100
                        that.downloadHtml = "<b style='font-size: 16px; color: #50bfff'>下载成功</b>，未弹窗请点击手动保存"
                    } else {
                        that.downloadProcess = percentComplete
                    }
                } else {
                    that.downloadHtml = "<b>下载失败</b>，未弹窗请点击手动保存"
                }
            }, false);

            req.addEventListener('error', function (evt) {
                console.log(evt)
                that.downloadHtml = "<b style='font-size: 16px; color: red'>下载失败</b>，请点击手动下载视频"
            });
            req.responseType = "blob";
            req.onreadystatechange = function() {
                if (req.readyState === 4 && req.status === 200) {
                    if (typeof window.chrome !== 'undefined') {
                        // Chrome version
                        var link = document.createElement('a');
                        link.href = window.URL.createObjectURL(req.response);
                        link.download = filename;
                        link.click();
                    } else if (typeof window.navigator.msSaveBlob !== 'undefined') {
                        // IE version
                        var blob = new Blob([req.response], {
                            type: 'application/force-download'
                        });
                        window.navigator.msSaveBlob(blob, filename);
                    } else {
                        // Firefox version
                        var file = new File([req.response], filename, {
                            type: 'application/force-download'
                        });
                        window.open(URL.createObjectURL(file));
                    }
                }
            };
            req.send();
        }
    }
});
