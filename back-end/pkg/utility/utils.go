package utility

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateTypeMeta(kind string, version string) *metav1.TypeMeta {
	return &metav1.TypeMeta{
		Kind:       kind,
		APIVersion: version,
	}
}

func CreateMetaData(name string, App string, Role string, Tier string) *metav1.ObjectMeta {
	return &metav1.ObjectMeta{
		Name: name,
		Labels: map[string]string{
			"App":  App,
			"Role": Role,
			"Tier": Tier,
		},
	}
}
