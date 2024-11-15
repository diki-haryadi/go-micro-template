package articleJob

import (
	"context"

	"go.uber.org/zap"

	"github.com/robfig/cron/v3"

	articleDomain "github.com/diki-haryadi/go-micro-template/internal/article/domain"
	"github.com/diki-haryadi/ztools/wrapper"
	wrapperErrorhandler "github.com/diki-haryadi/ztools/wrapper/handlers/error_handler"
	wrapperRecoveryHandler "github.com/diki-haryadi/ztools/wrapper/handlers/recovery_handler"
	wrapperSentryHandler "github.com/diki-haryadi/ztools/wrapper/handlers/sentry_handler"

	cronJob "github.com/diki-haryadi/ztools/cron"
)

type job struct {
	cron   *cron.Cron
	logger *zap.Logger
}

func NewJob(logger *zap.Logger) articleDomain.Job {
	newCron := cron.New(cron.WithChain(
		cron.SkipIfStillRunning(cronJob.NewLogger()),
	))
	return &job{cron: newCron, logger: logger}
}

func (j *job) StartJobs(ctx context.Context) {
	j.logArticleJob(ctx)
	go j.cron.Start()
}

func (j *job) logArticleJob(ctx context.Context) {
	worker := wrapper.BuildChain(j.logArticleWorker(),
		wrapperSentryHandler.SentryHandler,
		wrapperRecoveryHandler.RecoveryHandler,
		wrapperErrorhandler.ErrorHandler,
	)

	entryId, _ := j.cron.AddFunc("*/1 * * * *",
		worker.ToWorkerFunc(ctx, nil),
	)

	j.logger.Sugar().Infof("Article Job Started: %v", entryId)
}
