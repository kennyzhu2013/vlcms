/*
@Time : 2018/9/3 15:41 
@Author : kenny zhu
@File : base.go
@Software: GoLand
@Others:
*/
package config

type Service interface {
	C(keys string) interface{}

}

// direct call..
func C(keys string) interface{}{
	if c == nil {
		return nil
	}
	return c.dataMap[keys]
}

var c *configService

// NewMovieService returns the default movie service.
func Instance() Service {
	if c == nil {
		c = &configService{
			dataMap: make(map[string]interface{}),
		}
	}
	return  c
}

type configService struct {
	dataMap map[string]interface{}

}

func (c *configService) C(keys string) interface{}{
	return c.dataMap[keys]
}

