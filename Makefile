prod:
	@GOOS=linux GOARCH=amd64 go build main.go
	@mv main tutu-gin
	@ssh root@paqi2 "rm -rf /www/wwwroot/www.danaqsy.com/tutu-gin/tutu-gin"
	@scp ./tutu-gin root@paqi2:/www/wwwroot/www.danaqsy.com/tutu-gin/tutu-gin
	@ssh root@paqi2 "/www/server/panel/pyenv/bin/supervisorctl restart tutu-gin:tutu-gin_00"