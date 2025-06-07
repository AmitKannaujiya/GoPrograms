package gorutines

import (
	"bytes"
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func MultiURLTime(urls []string) {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go func(u string) {
			defer wg.Done()
			Urltime(u)
		}(url)
	}
	wg.Wait()
}

func Urltime(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error %q - %s", url, err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("error : %q - bad status - %s \n", url, resp.Status)
		return
	}
	// read the body
	_, err = io.Copy(io.Discard, resp.Body)
	if err != nil {
		log.Printf("error %q - %s", url, err)
		return
	}
	duration := time.Since(start)
	log.Printf("info %q - %v\n", url, duration)
}

// sha1sig return SHA1 signature in the format "35aabcd5a32e01d18a5ef688111624f3c547e13b"
func sha1Sig(data []byte) (string, error) {
	w := sha1.New()
	r := bytes.NewReader(data)
	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	sig := fmt.Sprintf("%x", w.Sum(nil))
	return sig, nil
}

type file struct {
	Name      string
	Content   []byte
	Signature string
}

type result struct {
	match    bool
	fileName string
	err      error
}

func WorkerSign(file file, ch chan<- result) {
	sig, err := sha1Sig(file.Content)
	ch <- result{
		match:    file.Signature == sig,
		err:      err,
		fileName: file.Name,
	}
}

// ValidateSigs return slice of OK files and slice of mismatched files
func ValidateSigs(files []file) ([]string, []string, error) {
	var okFiles []string
	var badFiles []string
	ch := make(chan result)
	for _, file := range files {
		go WorkerSign(file, ch)
	}
	for range files {
		res := <-ch
		if !res.match || res.err != nil {
			badFiles = append(badFiles, res.fileName)
		} else {
			okFiles = append(okFiles, res.fileName)
		}
	}
	close(ch)
	return okFiles, badFiles, nil
}

// 3: example for timeouts & cancellation
type Movie struct {
	ID    string
	Title string
}

var (
	defaultMovie = Movie{
		ID:    "tt0093779",
		Title: "The Princess Bride",
	}

	bmvTime = 50 * time.Millisecond
)

// BestNextMovie return the best move recommendation for a user
func BestNextMovie(user string) Movie {
	time.Sleep(bmvTime) // Simulate work

	// Don't change this, otherwise the test will fail
	return Movie{
		ID:    "tt0083658",
		Title: "Blade Runner",
	}
}

// NextMovie return BestNextMovie result if it finished before ctx expires, otherwise defaultMovie
func NextMovie(ctx context.Context, user string) Movie {
	ch := make(chan Movie)

	go func() {
		ch <- BestNextMovie(user)
	}()

	select {
	case movie, ok := <-ch:
		if !ok {
			return defaultMovie
		}
		return movie
	case <-ctx.Done():
		return defaultMovie
	}
}

// Cuncurrency : Limit No of goroutines
type input struct {
	src string
	dst string
}

// Center creates destFile which is the center of image encode in data.
func Center(srcFile, destFile string) error {
	file, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer file.Close()

	src, err := jpeg.Decode(file)
	if err != nil {
		return err
	}

	x, y := src.Bounds().Max.X, src.Bounds().Max.Y
	r := image.Rect(0, 0, x/2, y/2)
	dest := image.NewRGBA(r)
	draw.Draw(dest, dest.Bounds(), src, image.Point{x / 4, y / 4}, draw.Over)

	out, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer out.Close()

	return jpeg.Encode(out, dest, nil)
}

func worker1(ctx context.Context, in <-chan input, out chan<- error) {
	for {
		select {
		case i, ok := <-in:
			if !ok {
				return
			}
			out <- Center(i.src, i.dst)

		case <-ctx.Done():
			return
		}

	}
}

func producer1(ctx context.Context, in chan input, srcFiles []string, destDir string) {
	defer close(in)

	for _, src := range srcFiles {
		destFile := fmt.Sprintf("%s/%s", destDir, src)
		select {
		case in <- input{src: src, dst: destFile}:
		case <-ctx.Done():
			return
		}
	}
}

// CenterDir calls Center on every image in srcDir. n is the maximal number of goroutines.
func CenterDir(ctx context.Context, srcDir []string, destDir string, n int) error {
	if err := os.Mkdir(destDir, 0750); err != nil && !errors.Is(err, fs.ErrExist) {
		return err
	}

	matches, err := filepath.Glob(fmt.Sprintf("%s/*.jpg", srcDir))
	if err != nil {
		return err
	}
	in, out := make(chan input), make(chan error, n)

	for i := 0; i < n; i++ {
		go worker1(ctx, in, out)
	}
	go producer1(ctx, in, matches, destDir)

	for range matches {
		select {
		case err := <-out:
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return nil
}
