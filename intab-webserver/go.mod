module intab-webserver

go 1.13

replace intab-core => ../intab-core

require (
	github.com/asaskevich/govalidator v0.0.0-20200108200545-475eaeb16496
	github.com/bitly/go-simplejson v0.5.0
	github.com/facebookgo/inject v0.0.0-20180706035515-f23751cae28b
	github.com/facebookgo/structtag v0.0.0-20150214074306-217e25fb9691 // indirect
	github.com/gorilla/sessions v1.2.0
	github.com/imdario/mergo v0.3.9
	github.com/jinzhu/configor v1.2.0
	github.com/rushairer/ago v0.0.0-20170905022018-30d89f6a3366
	github.com/segmentio/ksuid v1.0.2
	intab-core v0.0.0-00010101000000-000000000000
)
