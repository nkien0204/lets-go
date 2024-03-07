package greeting

import "github.com/nkien0204/lets-go/internal/domain/entity/greeting"

func (repo *repository) SayHello() (resEntity greeting.GreetingResponseEntity, err error) {
	resEntity.Msg = "hello, world!"
	return
}
