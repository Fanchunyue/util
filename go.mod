module github.com/Fanchunyue/util

go 1.12

replace (
	golang.org/x/net => github.com/golang/net v0.0.0-20190301231341-16b79f2e4e95
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190309122539-980fc434d28e

)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.4.0
	github.com/opentracing/opentracing-go v1.1.0
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.4.2
)
