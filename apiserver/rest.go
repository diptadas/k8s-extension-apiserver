package apiserver

import (
	"errors"
	"k8s-admission-webhook/apis/foocontroller/v1alpha1"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
)

type REST struct{}

var _ rest.Getter = &REST{}
var _ rest.GroupVersionKindProvider = &REST{}

func NewREST() *REST {
	return &REST{}
}

func (r *REST) New() runtime.Object {
	return &v1alpha1.Foo{}
}

func (r *REST) GroupVersionKind(containingGV schema.GroupVersion) schema.GroupVersionKind {
	return v1alpha1.SchemeGroupVersion.WithKind("Foo")
}

func (r *REST) Get(ctx apirequest.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	log.Println("Get...")

	ns, ok := apirequest.NamespaceFrom(ctx)
	if !ok {
		return nil, errors.New("missing namespace")
	}
	if len(name) == 0 {
		return nil, errors.New("missing search query")
	}

	resp := &v1alpha1.Foo{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "foocontroller.k8s.io/v1alpha1",
			Kind:       "Foo",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: ns,
		},
		Spec: "do-not-care",
	}

	return resp, nil
}
