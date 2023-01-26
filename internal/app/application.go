package app

// import (
// 	postgres "AIDMS_model_toolkit/internal/adapter/repository/postgres"
// 	serviceImage "AIDMS_model_toolkit/internal/app/service/image"
// 	serviceJob "AIDMS_model_toolkit/internal/app/service/job"
// 	"context"
// )

type Application struct {
	// JobService   *serviceJob.JobService
	// ImageService *serviceImage.ImageService
}

// // application 接postgres.Store才能替換成mockDB
// func NewApplication(ctx context.Context, pgRepo postgres.Store) *Application {
// 	// Create application
// 	app := &Application{
// 		JobService: serviceJob.NewJobService(ctx, serviceJob.JobServiceParam{
// 			JobServiceRepo: pgRepo,
// 		}),
// 		ImageService: serviceImage.NewImageService(ctx, serviceImage.ImageServiceParam{
// 			ImageServiceRepo: pgRepo,
// 		}),
// 	}
// 	return app
// }
