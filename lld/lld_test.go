package lld

import (
	"fmt"
	dms "go-program/lld/dictionarysystem"
	imd "go-program/lld/inmemoryrdbms"
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

func TestInMemoryDataBase1(t *testing.T) {
	databaseManager := imd.GetDatabaseManager()
	assert.Nil(t, databaseManager.CreateDatabase("test_abc"))
	assert.Nil(t, databaseManager.CreateDatabase("test_xyz"))
	assert.Nil(t, databaseManager.SwitchDatabase("test_xyz"))
}

func TestDictionaryManagementSystem(t *testing.T) {
	dictionarysystem := dms.GetDictionaryManager()
	assert.Nil(t, dictionarysystem.CreateDataset("english"))
	assert.Nil(t, dictionarysystem.CreateDataset("hindi"))
	dictionarysystem.AddWordToDataset("english", "Hello")
	dictionarysystem.AddWordToDataset("english", "Hey")
	dictionarysystem.AddWordToDataset("english", "What")
	dictionarysystem.AddWordToDataset("english", "When")
	dictionarysystem.AddWordToDataset("hindi", "When")
	dictionarysystem.AddWordToDataset("hindi", "What")
	res1, err := dictionarysystem.SearchWord("What")
	assert.Nil(t, err)
	assert.Equal(t, []string{"What", "What"},res1)
	assert.Nil(t, dictionarysystem.CreateDataset("dataset"))
	assert.NotNil(t, dictionarysystem.CreateDataset("hindi"))
	assert.Nil(t, dictionarysystem.DeleteDataset("dataset"))

}

