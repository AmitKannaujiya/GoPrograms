package reviewrating

import (
	"fmt"
	"sync"
)

type User struct {
	Id   int
	Name string
}

type Product struct {
	Id          int
	Name        string
	Description string
}

type Review struct {
	Id         int
	UserId     int
	ProductId  int
	Rating     int
	ReviewText string
	VideoUrl   string
	Upvotes    int
	Downvotes  int
}

type Notification struct {
	Message string
	UserId  int
}

type ReviewManager struct {
	users             map[int]User
	products          map[int]Product
	productReviews    map[int][]Review
	userProductReview map[int]map[int]Review
	mutex             sync.RWMutex
	notification      chan Notification
	Wg                sync.WaitGroup
}

func (r Review) ToString() string {
	return fmt.Sprintf("Review id : %d, review text : %s, rating : %d, user id : %d, upvotes : %d, downvotes : %d",
	r.Id, r.ReviewText, r.Rating, r.UserId, r.Upvotes, r.Downvotes)
}

var SingletonReviewManagerInstance *ReviewManager
var once sync.Once

func GetReviewManagerInstance() *ReviewManager {
	if SingletonReviewManagerInstance == nil {
		once.Do(func() {
			SingletonReviewManagerInstance = &ReviewManager{
				users: make(map[int]User),
				products: make(map[int]Product),
				productReviews: make(map[int][]Review),
				userProductReview: make(map[int]map[int]Review),
				notification: make(chan Notification, 100), // 100 notifications can be send at once
			}
		})
	}
	return SingletonReviewManagerInstance
}
