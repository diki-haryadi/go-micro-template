package wrapperSentryhandler

import (
	"context"

	"github.com/getsentry/sentry-go"

	"github.com/diki-haryadi/go-micro-template/pkg/config"
	sentryUtils "github.com/diki-haryadi/go-micro-template/pkg/sentry/sentry_utils"
	"github.com/diki-haryadi/go-micro-template/pkg/wrapper"
)

var SentryHandler = func(f wrapper.HandlerFunc) wrapper.HandlerFunc {
	return func(ctx context.Context, args ...interface{}) (interface{}, error) {
		hub := sentry.GetHubFromContext(ctx)
		if hub == nil {
			hub = sentry.CurrentHub().Clone()
			ctx = sentry.SetHubOnContext(ctx, hub)
		}
		hub.Scope().SetExtra("args", args)
		hub.Scope().SetTag("application", config.BaseConfig.App.AppName)
		hub.Scope().SetTag("AppEnv", config.BaseConfig.App.AppEnv)

		opts := &sentryUtils.Options{
			Repanic: true,
		}
		defer sentryUtils.RecoverWithSentry(hub, ctx, opts)

		return f(ctx, args)
	}
}
