package greeting

import "github.com/nkien0204/lets-go/internal/domain/entity/greeting"

func (u *usecase) Greeting() (greeting.GreetingResponseEntity, error) {
	return u.repo.SayHello()
}
