package answer

import (
	"errors"
	"sync/atomic"
)

type Answer struct {
	Id     int64
	Answer int
	UserId int
	PollId int64
}

func init() {
	answers = append(answers, New(0, 1, 1))
}

var answers = []Answer{}

func Find(id int64) (Answer, error) {
	for _, v := range answers {
		if v.Id == id {
			return v, nil
		}
	}

	return Answer{}, errors.New("answer not found")
}

func FindPoll(pollId int64) []Answer {
	r := []Answer{}
	for _, v := range answers {
		if v.PollId == pollId {
			r = append(r, v)
		}
	}

	return r
}

var counter int64 = 0

func New(answer int, userId int, pollId int64) Answer {
	id := atomic.AddInt64(&counter, 1)
	ans := Answer{
		Id:     id,
		Answer: answer,
		UserId: userId,
		PollId: pollId,
	}

	return ans
}
