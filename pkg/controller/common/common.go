// Package common contains common definitions and shared utilities for the different controllers
package common

import (
	policiesv1 "github.com/open-cluster-management/governance-policy-propagator/pkg/apis/policies/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/cluster-registry/pkg/apis/clusterregistry/v1alpha1"
)

// IsInClusterNamespace check if policy is in cluster namespace
func IsInClusterNamespace(ns string, allClusters []v1alpha1.Cluster) bool {
	for _, cluster := range allClusters {
		if ns == cluster.GetNamespace() {
			return true
		}
	}
	return false
}

// LabelsForRootPolicy returns the labels for given policy
func LabelsForRootPolicy(plc *policiesv1.Policy) map[string]string {
	return map[string]string{"root-policy": FullNameForPolicy(plc)}
}

// FullNameForPolicy returns the fully qualified name for given policy
// full qualified name: ${namespace}.${name}
func FullNameForPolicy(plc *policiesv1.Policy) string {
	return plc.GetNamespace() + "." + plc.GetName()
}

// CompareSpecAndAnnotation compares annotation and spec for given policies
// true if matches, false if doesn't match
func CompareSpecAndAnnotation(plc1 *policiesv1.Policy, plc2 *policiesv1.Policy) bool {
	return equality.Semantic.DeepEqual(plc1.GetAnnotations(), plc2.GetAnnotations()) &&
		equality.Semantic.DeepEqual(plc1.Spec, plc1.Spec)
}

// IsPbForPoicy compares group and kind with policy group and kind for given pb
func IsPbForPoicy(pb *policiesv1.PlacementBinding) bool {
	if pb.Spec.Subject.Kind == policiesv1.Kind && pb.Spec.Subject.APIGroup == policiesv1.SchemeGroupVersion.Group {
		return true
	}
	return false
}

// // GenerateLabelsForReplicatedPolicy generates labels needed for replicated policy
// func GenerateLabelsForReplicatedPolicy(plc *policiesv1.Policy) {
// 	labels := plc.GetLabels()
// 	if labels == nil {
// 		labels = map[string]string{}
// 	}
// 	labels["cluster-name"] = decision.ClusterName
// 	labels["cluster-namespace"] = decision.ClusterNamespace
// 	labels["root-policy"] = common.FullNameForPolicy(instance)
// }
