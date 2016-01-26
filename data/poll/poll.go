package poll

import (
	"errors"
	"sync/atomic"

	"github.com/atamis/wt3/data/answer"
)

func init() {
	polls = append(polls, New())
}

type Poll struct {
	Id            int64
	Answers       []answer.Answer
	AnswersLoaded bool
}

var polls = []Poll{}

func Find(id int64) (Poll, error) {
	for _, v := range polls {
		if v.Id == id {
			return v, nil
		}
	}

	return Poll{}, errors.New("poll not found")
}

var counter int64 = 0

func New() Poll {
	id := atomic.AddInt64(&counter, 1)
	ans := Poll{
		Id:            id,
		Answers:       []answer.Answer{},
		AnswersLoaded: false,
	}

	return ans
}

func (p *Poll) LoadAnswers() {
	p.Answers = answer.FindPoll(p.Id)
	p.AnswersLoaded = true
}

func (p Poll) Collated() []int {
	ans := []int{0, 0, 0, 0, 0}
	for _, v := range p.Answers {
		ans[v.Answer]++
	}

	return ans
}

func (p *Poll) AddAnswer(userAnswer int, userId int) {
	ans := answer.New(userAnswer, userId, p.Id)
	ans.Save()
	p.LoadAnswers()
}
