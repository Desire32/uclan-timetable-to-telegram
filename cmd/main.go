package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"time"
	"uclan/internal/services/tg"
	"uclan/internal/services/timetable"
	"uclan/internal/services/timetable/data"

	"github.com/chromedp/chromedp"
	"github.com/joho/godotenv"
)

func main() {

	// MEMORY COUNT PART
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("До выполнения: %v KB\n", mem.Alloc/1024)
	ls := make([]int, 1e6)
	_ = ls
	//////////////////////////

	// channels for coop work
	//scheduleChan := make(chan string)
	//moduleChan := make(chan string)
	//badgesChan := make(chan string)
	//////////////////////////

	// SYNC
	var mu sync.Mutex
	//////////////////////////

	// .env
	_ = godotenv.Load("../internal/config/.env")
	////////////////////

	// CONNECTING INTERFACE
	authService := &timetable.AuthService{}
	//////////////////////

	//TIMETABLE INTERFACE
	timetableService := &timetable.TimeService{}
	///////////////////////

	//BADGES INTERFACE
	badgeService := &timetable.BadgeService{}
	///////////////////////

	// TELEGRAM INTERFACES
	tgService := &tg.ServiceTg{}
	////////////////

	// SETUP CONFIG
	ctx, cancel := chromedp.NewExecAllocator(context.Background(),
		chromedp.ExecPath(os.Getenv("BROWSER_PATH")),
		chromedp.Flag("headless", false),
		chromedp.WindowSize(1280, 1024),
	)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 90*time.Second)
	defer cancel()

	if err := authService.MicrosoftLogin(ctx); err != nil {
		log.Fatal(err)
	}

	if err := timetableService.TimetableAuth(ctx); err != nil {
		log.Fatal(err)
	}

	mu.Lock()
	scheduleInfo := timetableService.TimetableRetrieve(ctx)
	mu.Unlock()
	//scheduleChan <- scheduleInfo

	mu.Lock()
	badgesInfo := badgeService.BadgeRetrieve(ctx)
	mu.Unlock()
	//badgesChan <- badgesInfo

	mu.Lock()
	modulesInfo := badgeService.ModulesRetrieve(ctx)
	mu.Unlock()
	//moduleChan <- modulesInfo

	// Получаем результаты из каналов
	//scheduleInfo = <-scheduleChan
	//badgesInfo = <-badgesChan
	//modulesInfo = <-moduleChan

	// Выводим результаты
	fmt.Println("Информация о расписании:", scheduleInfo)
	fmt.Println("Информация о значках:", badgesInfo)

	if err := json.Unmarshal([]byte(modulesInfo), &data.Badge); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Информация о модулях: ", data.Badge)

	if err := json.Unmarshal([]byte(scheduleInfo), &data.Schedule); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Информация о расписании:", data.Schedule)

	data.BadgeAmount = badgesInfo

	go func() {
		if err := tgService.TgConnection(); err != nil {
			log.Fatal(err)
		}
	}()
	time.Sleep(100 * time.Minute)
}
