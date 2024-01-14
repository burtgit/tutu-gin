prod:
	@GOOS=linux GOARCH=amd64 go build main.go
	@mv main tutu-gin
	@ssh root@paqi2 "rm -rf /www/wwwroot/www.danaqsy.com/tutu-gin/tutu-gin"
	@scp ./tutu-gin root@paqi2:/www/wwwroot/www.danaqsy.com/tutu-gin/tutu-gin
	@ssh root@paqi2 "/www/server/panel/pyenv/bin/supervisorctl restart tutu-gin:tutu-gin_00"

master:
	@GOOS=linux GOARCH=amd64 go build main.go
	@mv main tutu-gin
	@ssh root@admin "rm -rf /data/go/tutu-gin/tutu-gin"
	@scp ./tutu-gin root@admin:/data/go/tutu-gin/tutu-gin
	@#ssh root@admin "/www/server/panel/pyenv/bin/supervisorctl restart tutu-gin:tutu-gin_00"


upload_resorce:
	scp -r ./http/resource root@paqi2:/www/wwwroot/www.danaqsy.com/tutu-gin/http


upload_resorce_over:
	scp -r ./http/resource root@byn-dida:/usr/local/src/tutu-gin/http/resource

prod-over:
	@GOOS=linux GOARCH=amd64 go build main.go
	@mv main tutu-gin
	@ssh root@ove "rm -rf /data/tutu-gin/tutu-gin"
	@scp ./tutu-gin root@ove:/data/tutu-gin/tutu-gin
	@ssh root@ove "/www/server/panel/pyenv/bin/supervisorctl restart tutu-gin:tutu-gin_00"