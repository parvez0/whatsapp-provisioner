package test_scripts

import (
	"context"
	kube_cli "github.com/parvez0/kube-api/kube"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

var kubeClientSet *kube_cli.KubeClientSet

// TestCreateClient creates a client with current active context
func TestCreateClient(t *testing.T)  {
	kubeClientSet = kube_cli.CreateClient("")
}

// TestNamespaceClient fetches all the namespaces in the cluster
func TestNamespaceClient(t *testing.T) {
	nsClient := kubeClientSet.GetNamespaceClient()
	ns, err := nsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil{
		t.Errorf("failed to fetch namespace list in the cluster - %+v", err)
	}
	t.Logf("successfully fetched all namespace = %+v", len(ns.Items))
}

// TestDeploymentClient fetches all the deployments in the default namespace
func TestDeploymentClient(t *testing.T) {
	deploymentClient := kubeClientSet.GetDeploymentClient("default")
	dep, err := deploymentClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil{
		t.Errorf("failed to fetch deployment list - %+v", err)
	}
	t.Logf("successfully fetched deployments list total deployments = %+v", len(dep.Items))
}

// Since all the other kube clients are working further client test for other resources not required
// Other services test cases can be written for the specific task