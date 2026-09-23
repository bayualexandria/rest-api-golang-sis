package scheduler

import (
	"log"
	"time"

	"backend-api/services"
)

func StartAcademicPeriodScheduler(
	academicPeriodService *services.AcademicPeriodService,
) {
	go func() {

		// Cek langsung ketika scheduler dijalankan
		check := func() {
			_, _, err := academicPeriodService.EnsureCurrentPeriod()

			if err != nil {
				log.Printf(
					"[ACADEMIC PERIOD] gagal memperbarui: %v",
					err,
				)
				return
			}

			log.Println(
				"[ACADEMIC PERIOD] tahun ajaran dan semester sudah diperiksa",
			)
		}

		check()

		// Cek setiap 1 jam
		ticker := time.NewTicker(4 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			check()
		}
	}()
}
