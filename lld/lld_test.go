package lld

import (
	"fmt"
	rr "go-program/lld/reviewrating"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReviewRating(t *testing.T) {
	rm := rr.GetReviewManagerInstance()

	rm.StartNotificationWorker()
	rm.AddUser(1, "ABC")
	rm.AddUser(102, "Xyz")

	rm.AddProduct(123 , "Toy", "Kids toys")
	rm.AddProduct(12 , "t-shirt", "Kids t-shirt")

	er := rm.PostReview(1, 123, 3, "Ok product")

	assert.Nil(t, er)

	rm.Upvotes(1)
	rm.Downvotes(1)

	reviews := rm.GetReviews(123)

	for _, review := range reviews {
		fmt.Println(review.ToString())
	}
	rm.Wg.Wait()
}



