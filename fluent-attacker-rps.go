package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/fluent/fluent-logger-golang/fluent"
	"math/big"
	"runtime"
	"strconv"
	"time"
	"sync"
)

var (
	fluentHost string	// fluentd host
	fluentPort int    	// fluentd port
	fluentTag  string	// fluent tag
	key        int		// 投入するflent Record数
	val        int		// fluent文字列の長さ
	rps        int		// Request Per Second
)


// val用にrandomな文字列を生成する
func random(length int) string {
	const base = 36
	size := big.NewInt(base)
	n := make([]byte, length)
	for i, _ := range n {
		c, _ := rand.Int(rand.Reader, size)
		n[i] = strconv.FormatInt(c.Int64(), base)[0]
	}
	return string(n)
}

// stringからmd5から作る
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// rpsを計算
// x -> sec | y -> 件数
func calcRPS(x float64, y int) float64 {
	return float64(y) / x
}

// パーセンテージを計算
// x - 全体
func calcPer(x int,y float64) (float64) {
        return y / float64(x) * 100
}

// RPSから間隔を計算
func calcInt(x int) int64 {
	return int64(1 / float64(x) *1000 *1000)
}

// メイン処理
func main() {

	// flag処理
	flag.StringVar(&fluentHost, "h","127.0.0.1","Fluentd-Server IP Address")
	flag.IntVar(&fluentPort, "p",24224,"Fluentd-Server Port Address")
	flag.StringVar(&fluentTag, "t","test.tag","Fluentd Tag")
	flag.IntVar(&key, "c",3000,"Fluentd TestData Post Count")
	flag.IntVar(&val, "l",1000,"Fluentd TestData Message Length")
	flag.IntVar(&rps, "r",300,"Request Per Second")

	// いつものおまじない
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	// LogPrint
	fmt.Print("=====================================\n")
	fmt.Printf("Insert Keys=> %d",key)
	fmt.Printf(" RPS Set => %d",rps)
	fmt.Printf(" Message Length => %d\n",val)
	fmt.Print("=====================================\n")

	// Channel宣言
	task := make(chan int)        // TaskQue用
	taskquit := make(chan bool)   // SetTask停止用
	workerquit := make(chan bool) // JobWorker停止用

	// WaitGroup作成
	wg := &sync.WaitGroup{}

	// Timer Set
	start := time.Now()

	// Job Worker
	go func() {
	loop:
		for {
			select {
			case <-taskquit:
				workerquit <- true
				break loop
			case job := <-task:
				wg.Add(1)
				postFluent(job, val)
				wg.Done()
			}
		}
	}()

	// SetTask Worker
	go func() {

		// RPS間隔計算
		interval := calcInt(rps)
		sleepWait := time.Duration(interval) * time.Microsecond

		for i := 0; i < key; i++ {
			task <- i
			time.Sleep(sleepWait)

			// Log Print
			if i%rps == 0 {
				fmt.Printf("%d loop - TaskSet Done.\n",i)
			}
		}
		taskquit <- true
	}()

	// WaitGroup終了まで待つ
	wg.Wait()

	<-workerquit

	// Log Print
	fmt.Printf("%d Task - All Job Done.\n",key)

	// Timer Print
	end := time.Now()
	sec := (end.Sub(start)).Seconds()
	calRPS := calcRPS(sec, key)

	fmt.Print("=====================================\n")
	fmt.Printf("Exec Time => %fSec", sec)
	fmt.Printf(" Real RPS => %frps", calRPS)
	fmt.Printf("(%f%%)\n", (calcPer(rps, calRPS)))
	fmt.Print("=====================================\n")
}

// Post fluentd
func postFluent(k, v int) {
	// fluent Setting
	logger, err := fluent.New(fluent.Config{
		FluentPort: fluentPort,
		FluentHost: fluentHost,
	})

	if err != nil {
		fmt.Println(err)
	}

	defer logger.Close()

	vals := random(v)
	keys := GetMD5Hash(vals)

	var data = map[string]string{
		"ID": strconv.Itoa(k),
		"md5": keys,
		"strings": vals,
	}

	error := logger.Post(fluentTag, data)
	if error != nil {
		panic(error)
	}
}
