package main

import (
	"context"
	"flag"

	corev1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var log = ctrl.Log.WithName("Maneger")

func main() {
	opts := zap.Options{
		Development: true,
	}
	opts.BindFlags(flag.CommandLine)
	flag.Parse()
	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))
	cfg, _ := config.GetConfig()
	mgr, _ := manager.New(cfg, manager.Options{})

	Reconciler := reconcile.Func(Reconciler)

	ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Pod{}).
		Complete(Reconciler)

	ctx := context.Background()
	if err := mgr.Start(ctx); err != nil {
		log.Error(err, "can't start manager")
	}
}

func Reconciler(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	log.Info("Reconciler is called", "Pod", req.Name)
	return reconcile.Result{}, nil
}
