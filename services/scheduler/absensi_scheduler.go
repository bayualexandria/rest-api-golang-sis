package scheduler

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"

	"backend-api/config"
	"backend-api/services"
)

func StartAbsensiScheduler() {

	location, err := time.LoadLocation("Asia/Jayapura")

	if err != nil {
		log.Fatal(err)
	}

	absensiService := services.NewAbsensiService(config.DB)

	cronJob := cron.New(
		cron.WithLocation(location),
	)

	// Jalankan setiap 5 menit
	_, err = cronJob.AddFunc("*/5 * * * *", func() {

		now := time.Now()

		// Belum mencapai batas absensi
		if now.Hour() < 8 {
			return
		}

		log.Println(
			"Menjalankan pengecekan ALPA:",
			now.Format("2006-01-02 15:04:05"),
		)

		if err := absensiService.GenerateAlpa(); err != nil {

			log.Println(
				"Gagal generate ALPA:",
				err,
			)

			return
		}

		log.Println(
			"Generate ALPA selesai",
		)
	})

	if err != nil {
		log.Fatal(err)
	}

	cronJob.Start()

	log.Println(
		"Scheduler absensi aktif - timezone Asia/Jayapura",
	)
}