package create

import (
	"testing"

	"github.com/ray-project/kuberay/kubectl-plugin/pkg/util"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func TestRayCreateClusterComplete(t *testing.T) {
	testStreams, _, _, _ := genericclioptions.NewTestIOStreams()
	fakeCreateClusterOptions := NewCreateClusterOptions(testStreams)
	fakeArgs := []string{"testRayClusterName"}
	cmd := &cobra.Command{Use: "cluster"}

	err := fakeCreateClusterOptions.Complete(cmd, fakeArgs)
	assert.Nil(t, err)
	assert.Equal(t, "default", *fakeCreateClusterOptions.configFlags.Namespace)
	assert.Equal(t, "testRayClusterName", fakeCreateClusterOptions.clusterName)
}

func TestRayCreateClusterValidate(t *testing.T) {
	testStreams, _, _, _ := genericclioptions.NewTestIOStreams()

	testNS, testContext, testBT, testImpersonate := "test-namespace", "test-context", "test-bearer-token", "test-person"

	kubeConfigWithCurrentContext, err := util.CreateTempKubeConfigFile(t, testContext)
	assert.Nil(t, err)

	kubeConfigWithoutCurrentContext, err := util.CreateTempKubeConfigFile(t, "")
	assert.Nil(t, err)

	fakeConfigFlags := &genericclioptions.ConfigFlags{
		Namespace:        &testNS,
		Context:          &testContext,
		KubeConfig:       &kubeConfigWithCurrentContext,
		BearerToken:      &testBT,
		Impersonate:      &testImpersonate,
		ImpersonateGroup: &[]string{"fake-group"},
	}

	tests := []struct {
		name        string
		opts        *CreateClusterOptions
		expectError string
	}{
		{
			name: "Test validation when no context is set",
			opts: &CreateClusterOptions{
				configFlags: &genericclioptions.ConfigFlags{
					KubeConfig: &kubeConfigWithoutCurrentContext,
				},
				ioStreams: &testStreams,
			},
			expectError: "no context is currently set, use \"kubectl config use-context <context>\" to select a new one",
		},
		{
			name: "Successful submit job validation with RayJob",
			opts: &CreateClusterOptions{
				configFlags:    fakeConfigFlags,
				ioStreams:      &testStreams,
				clusterName:    "fakeclustername",
				rayVersion:     "ray-version",
				image:          "ray-image",
				headCPU:        "5",
				headMemory:     "5Gi",
				workerReplicas: 3,
				workerCPU:      "4",
				workerMemory:   "5Gi",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.opts.Validate()
			if tc.expectError != "" {
				assert.Error(t, err, tc.expectError)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
