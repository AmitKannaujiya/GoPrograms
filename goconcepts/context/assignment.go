package context

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"
)

func ContextWithTimeOut() {
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()

	urls := []string{
		"https://api.restful-api.dev/objects/7",
		"https://api.restful-api.dev/objects/5",
		"https://api.restful-api.dev/objects/9",
		"https://api.restful-api.dev/objects/11",
		"https://api.restful-api.dev/objects/15",
	}

	results := make(chan string)
	for _, url := range urls {
		go callAPI(ctx, url, results)
	}
loop:
	for {
		select {
		case res, ok := <-results:
			if !ok {
				fmt.Println(res)
			} else {
				fmt.Println("channel is closed", res)
			}
		case <-ctx.Done():
			fmt.Println("All job done")
			break loop
		}
	}
	close(results)
	fmt.Println("All url are readed")
}

func callAPI(ctx context.Context, url string, results chan<- string) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		results <- fmt.Sprintf("Error creating request obj :%s for url: %s", err.Error(), url)
		return
	}
	client := http.DefaultClient
	response, er := client.Do(req)
	if er != nil {
		results <- fmt.Sprintf("error making request : %s , for url : %s ", er.Error(), url)
		return
	}
	defer response.Body.Close()
	res, e := io.ReadAll(response.Body)
	if e != nil {
		results <- fmt.Sprintf("error reading response : %s , for url : %s ", e.Error(), url)
		return
	}
	results <- fmt.Sprintf("Response :%s, from url : %s", url, string(res))
}

func ContextWithValueExample(uid string) bool {
	ctx := context.WithValue(context.TODO(), "userId", uid)
	result := make(chan bool)
	go performUserIdValidation(ctx, result)
	return <-result
}

func performUserIdValidation(ctx context.Context, result chan bool) {
	uid := ctx.Value("userId")
	r, _ := regexp.Compile("[a-z]+")
	d, ok := uid.(string)
	if ok {
		b := []byte(d)
		result <- r.Match(b)
	} else {
		result <- false
	}
}

func ContextWithCancel() {
	ctx, cancel := context.WithCancel(context.TODO())
	go performLongTask(ctx)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func performLongTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context is cancelled")
			return
		default:
			fmt.Println("Peforming tasks")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func ContextWithDeadline() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2 * time.Second))
	defer cancel()
	go performLongTask(ctx)
	time.Sleep(3 * time.Second)
}