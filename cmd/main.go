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

	//MONGODB INTERFACE
	mongoService := &timetable.MongoService{}
	////////////////////////

	// TELEGRAM INTERFACES
	tgService := &tg.ServiceTg{}
	////////////////

	// SETUP CONFIG
	ctx, cancel := chromedp.NewExecAllocator(context.Background(),
		chromedp.ExecPath(os.Getenv("BROWSER_PATH")),
		// chromedp.Flag("headless", false), // uncomment if you want to see how it looks like in real time
		chromedp.Flag("headless", true),
		chromedp.WindowSize(1280, 1024),
	)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 240*time.Second)
	defer cancel()

	if err := authService.MicrosoftLogin(ctx); err != nil {
		log.Fatal(err)
	}

	if err := timetableService.TimetableAuth(ctx); err != nil {
		log.Fatal(err)
	}

	mu.Lock()
	scheduleInfo := timetableService.TimetableRetrieve(ctx)
	defer mu.Unlock()

	mu.Lock()
	badgesInfo := badgeService.BadgeRetrieve(ctx)
	defer mu.Unlock()

	mu.Lock()
	modulesInfo := badgeService.ModulesRetrieve(ctx)
	defer mu.Unlock()

	// fmt.Println("Информация о расписании:", scheduleInfo)
	// fmt.Println("Информация о значках:", badgesInfo)

	// Uncomment if you want to add data to MongoDB, prepare login data into .env file
	// if err := mongoService.MongoSend(scheduleInfo); err != nil {
	// 	log.Fatal(err)
	// }
	// if err := mongoService.MongoSend(badgesInfo); err != nil {
	// 	log.Fatal(err)
	// }

	// заглушки
	_ = mongoService.MongoSend(scheduleInfo)
	_ = mongoService.MongoSend(badgesInfo)

	if err := json.Unmarshal([]byte(modulesInfo), &data.Badge); err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Информация о модулях: ", data.Badge)

	if err := json.Unmarshal([]byte(scheduleInfo), &data.Schedule); err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Информация о расписании:", data.Schedule)

	data.BadgeAmount = badgesInfo

	// telegram api launch
	go tgService.TgConnection()

	select {}
}
