package cli

import (
	"github.com/dhruvvadoliya1/movie-app-backend/config"

	_ "github.com/lib/pq"

	"github.com/spf13/cobra"
)

func GetSeedCommandDef(cfg config.AppConfig) cobra.Command {

	seederCommand := cobra.Command{
		Use:   "seed",
		Short: "seed the data in the models from the resource models",
		Long:  `seed the in the models from the resource csv files`,
		// RunE: func(cmd *cobra.Command, args []string) error {

		// 	// db, err := database.Connect(cfg.DB)
		// 	// if err != nil {
		// 	// 	return err
		// 	// }

		// 	// // Initialize Seeder
		// 	// seeder := database.NewSeeder(db)

		// 	// JobCSVPath := os.Getenv("COMMON_RESOURCE_PATH") + os.Getenv("JOB_PATH")
		// 	// err = seeder.SeedJobs(JobCSVPath)
		// 	// if err != nil {
		// 	// 	return err
		// 	// }

		// 	// CompanyPath := os.Getenv("COMMON_RESOURCE_PATH") + os.Getenv("COMPANY_PATH")
		// 	// err = seeder.SeedCompanies(CompanyPath)
		// 	// return err
		// },
	}
	return seederCommand
}
