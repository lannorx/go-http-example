example-go-service
==================

[Install the Go tools](http://golang.org/doc/install)

A example install go-service (Ubuntu)
-------------------------------------

```bash
$ git clone https://github.com/lannor-v2/example-go-service.git ~/myservice
$ cd ~/myservice
$ go build http.go
$
$ # Replace a string in a file on "exec fullpath-your-go-service"
$ pwd
/home/user/myservice
$ sed -e "s/exec \/path\/to\/http/exec \/home\/user\/myservice\/http/g" ./go-http.conf > ./go-myservice.conf
$ sudo cp -fv ./go-myservice.conf /etc/init/go-myservice.conf
$
$ # Start/Stop service
$ sudo service go-myservice start
$ sudo service go-myservice stop
```

Example install Apache (Ubuntu)
-------------------------------

```bash
$ sudo apt-get install apache2
$ sudo a2enmod rewrite
$ sudo a2enmod proxy
$ sudo a2enmod proxy_http
$
$ sudo cp -fv ./httpd-go-http.conf /etc/apache2/sites-available/go-http
$ sudo a2ensite go-http
$ sudo a2dissite default
$
$ sudo service apache2 restart
```

