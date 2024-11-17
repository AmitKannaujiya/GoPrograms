package reviewrating

import (
	"errors"
	"fmt"
)

// create/add user
func(r *ReviewManager) AddUser(userId int, name string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.users[userId] = User{Id: userId, Name: name}
}

// Add Product
func(r *ReviewManager) AddProduct(productId int, name, productDesc string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.products[productId] = Product{Id: productId, Name: name, Description: productDesc}
}

// post review
func (r *ReviewManager) PostReview(userId, productId, rating int, reviewText string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[userId]; !exists {
		return errors.New("USER DOES NOT EXISTS")
	}
	if _, exists := r.products[productId]; !exists {
		return errors.New("PRODUCT DOES NOT EXISTS")
	}
	// upload video to s3

	// check if already reviwed
	if _, exists := r.userProductReview[userId];!exists {
		r.userProductReview[userId] = make(map[int]Review)
	}
	if review, exists := r.userProductReview[userId][productId]; exists {
		review.Rating = rating
		review.ReviewText = reviewText
		r.userProductReview[userId][productId]= review
	} else {
		review := Review{
			Id: len(r.productReviews[productId])+1,
			UserId: userId,
			ProductId: productId,
			Rating: rating,
			ReviewText: reviewText,
		}
		r.productReviews[productId] =append(r.productReviews[productId], review)
		r.userProductReview[userId][productId]= review
	}
	// notify all followers
	r.SendNotification(userId, fmt.Sprintf("New reiviewhas been posted by %d on product %d", userId, productId))
	return nil
}

func(r *ReviewManager) Upvotes(reviewId int) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for productId, reviews :=range r.productReviews {
		for i, review := range reviews {
			if review.Id == reviewId {
				r.productReviews[productId][i].Upvotes++
				r.SendNotification(review.UserId, "Your review has been upvoted !")
				return
			}
		}
	}
}

func(r *ReviewManager) Downvotes(reviewId int) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for productId, reviews :=range r.productReviews {
		for i, review := range reviews {
			if review.Id == reviewId {
				r.productReviews[productId][i].Downvotes++
				r.SendNotification(review.UserId, "Your review has been downvoted !")
				return
			}
		}
	}
}
// get all reviews of product
func(r *ReviewManager) GetReviews(productId int) []Review {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.productReviews[productId]
}

func(r *ReviewManager) SendNotification(userId int, message string) {
	// notify all followers
	r.Wg.Add(1)
	go func ()  {
		defer r.Wg.Done()
		r.notification <- Notification{
			UserId: userId,
			Message: message,
		}
	}()
}

func(r *ReviewManager) StartNotificationWorker() {
	go func(){
	  for notification := range r.notification{
		fmt.Printf("Notification sent to user %d : %s \n", notification.UserId,notification.Message)
	  }
	}()
}