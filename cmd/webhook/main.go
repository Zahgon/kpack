package main

import (
	"context"
	"os"
	"strconv"

	"k8s.io/apimachinery/pkg/runtime/schema"
	informersv1 "k8s.io/client-go/informers/storage/v1"
	listersv1 "k8s.io/client-go/listers/storage/v1"
	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/injection"
	"knative.dev/pkg/injection/sharedmain"
	"knative.dev/pkg/signals"
	"knative.dev/pkg/webhook"
	"knative.dev/pkg/webhook/certificates"
	"knative.dev/pkg/webhook/resourcesemantics"

	"github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
)

const defaultWebhookPort = 8443

var types = map[schema.GroupVersionKind]resourcesemantics.GenericCRD{
	v1alpha2.SchemeGroupVersion.WithKind(v1alpha2.ImageKind):            &v1alpha2.Image{},
	v1alpha2.SchemeGroupVersion.WithKind(v1alpha2.BuildKind):            &v1alpha2.Build{},
	v1alpha2.SchemeGroupVersion.WithKind(v1alpha2.BuilderKind):          &v1alpha2.Builder{},
	v1alpha2.SchemeGroupVersion.WithKind(v1alpha2.BuildpackKind):        &v1alpha2.Buildpack{},
	v1alpha2.SchemeGroupVersion.WithKind(v1alpha2.ClusterBuilderKind):   &v1alpha2.ClusterBuilder{},
	v1alpha2.SchemeGroupVersion.WithKind(v1alpha2.ClusterBuildpackKind): &v1alpha2.ClusterBuildpack{},
	v1alpha2.SchemeGroupVersion.WithKind(v1alpha2.ClusterStoreKind):     &v1alpha2.ClusterStore{},
	v1alpha2.SchemeGroupVersion.WithKind(v1alpha2.ClusterStackKind):     &v1alpha2.ClusterStack{},
	v1alpha2.SchemeGroupVersion.WithKind(v1alpha2.ClusterLifecycleKind): &v1alpha2.ClusterLifecycle{},
}

func init() {
	injection.Default.RegisterInformer(withStorageClassInformer)
}

func main() {
	webhookPort := defaultWebhookPort
	webhookPortEnv := os.Getenv("WEBHOOK_PORT")
	if parsedWebhookPort, err := strconv.Atoi(webhookPortEnv); err == nil {
		webhookPort = parsedWebhookPort
	}
	ctx := webhook.WithOptions(signals.NewContext(), webhook.Options{
		ServiceName: "kpack-webhook",
		Port:        webhookPort,
		SecretName:  "webhook-certs",
	})

	sharedmain.WebhookMainWithConfig(ctx, "webhook",
		injection.ParseAndGetRESTConfigOrDie(),
		certificates.NewController,
		defaultingAdmissionController,
		validatingAdmissionController,
		conversionController,
	)
}

func defaultingAdmissionController(ctx context.Context, _ configmap.Watcher) *controller.Impl {
	_ = "STUB: not implemented"
	return nil
}

// Name of the resource webhook.

// The path on which to serve the webhook.

// The resources to default.

// A function that infuses the context passed to Validate/SetDefaults with custom metadata.

// Whether to disallow unknown fields.

func validatingAdmissionController(ctx context.Context, _ configmap.Watcher) *controller.Impl {
	_ = "STUB: not implemented"
	return nil
}

// Name of the resource webhook.

// The path on which to serve the webhook.

// The resources to validate.

// A function that infuses the context passed to Validate/SetDefaults with custom metadata.

// Whether to disallow unknown fields.

func conversionController(ctx context.Context, _ configmap.Watcher) *controller.Impl {
	_ = "STUB: not implemented"
	return nil
}

func withCheckDefaultStorageClass(storageClassLister listersv1.StorageClassLister) func(context.Context) context.Context {
	_ = "STUB: not implemented"
	return nil
}

// storageClassInformerKey is used for associating the Informer inside the context.Context.
type storageClassInformerKey struct{}

func withStorageClassInformer(ctx context.Context) (context.Context, controller.Informer) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(controller.Informer)
}

func getStorageClassInformer(ctx context.Context) informersv1.StorageClassInformer {
	_ = "STUB: not implemented"
	return *new(informersv1.StorageClassInformer)
}
