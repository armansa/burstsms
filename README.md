install mysql server
creatre burstsms database
I used root/root for mysql username/password

install redis on localhost for message queue
I haven't added user pass authentication!

install go1.8   (I used gvm)

install all packages:
go get github.com/astaxie/beego
go get github.com/beego/bee
go get github.com/astaxie/beego/orm
go get github.com/go-sql-driver/mysql
go get -d github.com/zpnk/go-bitly/...
go get github.com/mvdan/xurls

unfortunately all configurations are hard coded!
please setup api-key, api-sec and bitly auth-token in jobs/main.go

to create mysql table run this command in burstsms directory
./burstsms orm syncdb -db default -force

to run web application run this command in burstsms directory
bee run

to run background worker run this command in jobs directory
go run main.go

open browser on localhost:8080/sms
send a message with url!
