/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nanny

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

var (
	fullEstimator = LinearEstimator{
		Resources: []Resource{
			{
				Base:         resource.MustParse("0.3"),
				ExtraPerNode: resource.MustParse("1"),
				Name:         "cpu",
			},
			{
				Base:         resource.MustParse("30Mi"),
				ExtraPerNode: resource.MustParse("1Mi"),
				Name:         "memory",
			},
			{
				Base:         resource.MustParse("30Gi"),
				ExtraPerNode: resource.MustParse("1Gi"),
				Name:         "storage",
			},
		},
	}
	noCPUEstimator = LinearEstimator{
		Resources: []Resource{
			{
				Base:         resource.MustParse("30Mi"),
				ExtraPerNode: resource.MustParse("1Mi"),
				Name:         "memory",
			},
			{
				Base:         resource.MustParse("30Gi"),
				ExtraPerNode: resource.MustParse("1Gi"),
				Name:         "storage",
			},
		},
	}
	noMemoryEstimator = LinearEstimator{
		Resources: []Resource{
			{
				Base:         resource.MustParse("0.3"),
				ExtraPerNode: resource.MustParse("1"),
				Name:         "cpu",
			},
			{
				Base:         resource.MustParse("30Gi"),
				ExtraPerNode: resource.MustParse("1Gi"),
				Name:         "storage",
			},
		},
	}
	noStorageEstimator = LinearEstimator{
		Resources: []Resource{
			{
				Base:         resource.MustParse("0.3"),
				ExtraPerNode: resource.MustParse("1"),
				Name:         "cpu",
			},
			{
				Base:         resource.MustParse("30Mi"),
				ExtraPerNode: resource.MustParse("1Mi"),
				Name:         "memory",
			},
		},
	}
	lessThanMilliEstimator = LinearEstimator{
		Resources: []Resource{
			{
				Base:         resource.MustParse("0.3"),
				ExtraPerNode: resource.MustParse("0.5m"),
				Name:         "cpu",
			},
		},
	}
	emptyEstimator = LinearEstimator{
		Resources: []Resource{},
	}

	exponentialEstimator = ExponentialEstimator{
		Resources: []Resource{
			{
				Base:         resource.MustParse("0.3"),
				ExtraPerNode: resource.MustParse("1"),
				Name:         "cpu",
			},
			{
				Base:         resource.MustParse("30Mi"),
				ExtraPerNode: resource.MustParse("1Mi"),
				Name:         "memory",
			},
			{
				Base:         resource.MustParse("30Gi"),
				ExtraPerNode: resource.MustParse("1Gi"),
				Name:         "storage",
			},
		},
		ScaleFactor: 1.5,
	}
	exponentialLessThanMilliEstimator = ExponentialEstimator{
		Resources: []Resource{
			{
				Base:         resource.MustParse("0.3"),
				ExtraPerNode: resource.MustParse("0.5m"),
				Name:         "cpu",
			},
		},
		ScaleFactor: 1.5,
	}

	baseResources = corev1.ResourceList{
		"cpu":     resource.MustParse("0.3"),
		"memory":  resource.MustParse("30Mi"),
		"storage": resource.MustParse("30Gi"),
	}

	noCPUBaseResources = corev1.ResourceList{
		"memory":  resource.MustParse("30Mi"),
		"storage": resource.MustParse("30Gi"),
	}
	noMemoryBaseResources = corev1.ResourceList{
		"cpu":     resource.MustParse("0.3"),
		"storage": resource.MustParse("30Gi"),
	}
	noStorageBaseResources = corev1.ResourceList{
		"cpu":    resource.MustParse("0.3"),
		"memory": resource.MustParse("30Mi"),
	}
	singleNodeResources = corev1.ResourceList{
		"cpu":     resource.MustParse("1.3"),
		"memory":  resource.MustParse("31Mi"),
		"storage": resource.MustParse("31Gi"),
	}
	noCPUSingleNodeResources = corev1.ResourceList{
		"memory":  resource.MustParse("31Mi"),
		"storage": resource.MustParse("31Gi"),
	}
	noMemorySingleNodeResources = corev1.ResourceList{
		"cpu":     resource.MustParse("1.3"),
		"storage": resource.MustParse("31Gi"),
	}
	noStorageSingleNodeResources = corev1.ResourceList{
		"cpu":    resource.MustParse("1.3"),
		"memory": resource.MustParse("31Mi"),
	}
	threeNodeResources = corev1.ResourceList{
		"cpu":     resource.MustParse("3.3"),
		"memory":  resource.MustParse("33Mi"),
		"storage": resource.MustParse("33Gi"),
	}
	threeNodeNoCPUResources = corev1.ResourceList{
		"memory":  resource.MustParse("33Mi"),
		"storage": resource.MustParse("33Gi"),
	}
	threeNodeNoMemoryResources = corev1.ResourceList{
		"cpu":     resource.MustParse("3.3"),
		"storage": resource.MustParse("33Gi"),
	}
	threeNodeNoStorageResources = corev1.ResourceList{
		"cpu":    resource.MustParse("3.3"),
		"memory": resource.MustParse("33Mi"),
	}
	threeNodeLessThanMilliResources = corev1.ResourceList{
		"cpu": resource.MustParse("0.3015"),
	}
	sixteenNodeLessThanMilliExpResources = corev1.ResourceList{
		"cpu": resource.MustParse("0.308"),
	}
	fourNodeResources = corev1.ResourceList{
		"cpu":     resource.MustParse("4.3"),
		"memory":  resource.MustParse("34Mi"),
		"storage": resource.MustParse("34Gi"),
	}
	fourNodeNoCPUResources = corev1.ResourceList{
		"memory":  resource.MustParse("34Mi"),
		"storage": resource.MustParse("34Gi"),
	}
	fourNodeNoMemoryResources = corev1.ResourceList{
		"cpu":     resource.MustParse("4.3"),
		"storage": resource.MustParse("34Gi"),
	}
	fourNodeNoStorageResources = corev1.ResourceList{
		"cpu":    resource.MustParse("4.3"),
		"memory": resource.MustParse("34Mi"),
	}
	fourNodeLessThanMilliResources = corev1.ResourceList{
		"cpu": resource.MustParse("0.302"),
	}
	twentyFourNodeLessThanMilliExpResources = corev1.ResourceList{
		"cpu": resource.MustParse("0.312"),
	}
	noResources = corev1.ResourceList{}

	sixteenNodeResources = corev1.ResourceList{
		"cpu":     resource.MustParse("16.3"),
		"memory":  resource.MustParse("46Mi"),
		"storage": resource.MustParse("46Gi"),
	}
	seventeenNodeResources = corev1.ResourceList{
		"cpu":     resource.MustParse("17.3"),
		"memory":  resource.MustParse("47Mi"),
		"storage": resource.MustParse("47Gi"),
	}
	twentyFourNodeResources = corev1.ResourceList{
		"cpu":     resource.MustParse("24.3"),
		"memory":  resource.MustParse("54Mi"),
		"storage": resource.MustParse("54Gi"),
	}
	twentyFiveNodeResources = corev1.ResourceList{
		"cpu":     resource.MustParse("25.3"),
		"memory":  resource.MustParse("55Mi"),
		"storage": resource.MustParse("55Gi"),
	}
	thirtySixNodeResources = corev1.ResourceList{
		"cpu":     resource.MustParse("36.3"),
		"memory":  resource.MustParse("66Mi"),
		"storage": resource.MustParse("66Gi"),
	}
)

func verifyResources(t *testing.T, kind string, got, want corev1.ResourceList) {
	if len(got) != len(want) {
		t.Errorf("%s not equal got: %+v want: %+v", kind, got, want)
	}
	for res, val := range want {
		actVal, ok := got[res]
		if !ok {
			t.Errorf("missing resource %s in %s", res, kind)
		}
		if val.Cmp(actVal) != 0 {
			t.Errorf("not equal resource %s in %s, got: %+v, want: %+v", res, kind, actVal.String(), val.String())
		}
	}
}

func TestEstimateResources(t *testing.T) {
	testCases := []struct {
		e                  ResourceEstimator
		numNodes           uint64
		expectedLimits     corev1.ResourceList
		expectedRequests   corev1.ResourceList
		acceptableLimits   corev1.ResourceList
		acceptableRequests corev1.ResourceList
	}{
		{fullEstimator, 0, baseResources, baseResources, singleNodeResources, singleNodeResources},
		{fullEstimator, 3, threeNodeResources, threeNodeResources, fourNodeResources, fourNodeResources},
		{fullEstimator, 16, sixteenNodeResources, sixteenNodeResources, seventeenNodeResources, seventeenNodeResources},
		{fullEstimator, 24, twentyFourNodeResources, twentyFourNodeResources, twentyFiveNodeResources, twentyFiveNodeResources},
		{noCPUEstimator, 0, noCPUBaseResources, noCPUBaseResources, noCPUSingleNodeResources, noCPUSingleNodeResources},
		{noCPUEstimator, 3, threeNodeNoCPUResources, threeNodeNoCPUResources, fourNodeNoCPUResources, fourNodeNoCPUResources},
		{noMemoryEstimator, 0, noMemoryBaseResources, noMemoryBaseResources, noMemorySingleNodeResources, noMemorySingleNodeResources},
		{noMemoryEstimator, 3, threeNodeNoMemoryResources, threeNodeNoMemoryResources, fourNodeNoMemoryResources, fourNodeNoMemoryResources},
		{noStorageEstimator, 0, noStorageBaseResources, noStorageBaseResources, noStorageSingleNodeResources, noStorageSingleNodeResources},
		{noStorageEstimator, 3, threeNodeNoStorageResources, threeNodeNoStorageResources, fourNodeNoStorageResources, fourNodeNoStorageResources},
		{lessThanMilliEstimator, 3, threeNodeLessThanMilliResources, threeNodeLessThanMilliResources, fourNodeLessThanMilliResources, fourNodeLessThanMilliResources},
		{emptyEstimator, 0, noResources, noResources, noResources, noResources},
		{emptyEstimator, 3, noResources, noResources, noResources, noResources},
		{exponentialEstimator, 0, sixteenNodeResources, sixteenNodeResources, twentyFourNodeResources, twentyFourNodeResources},
		{exponentialEstimator, 3, sixteenNodeResources, sixteenNodeResources, twentyFourNodeResources, twentyFourNodeResources},
		{exponentialEstimator, 10, sixteenNodeResources, sixteenNodeResources, twentyFourNodeResources, twentyFourNodeResources},
		{exponentialEstimator, 16, sixteenNodeResources, sixteenNodeResources, twentyFourNodeResources, twentyFourNodeResources},
		{exponentialEstimator, 17, twentyFourNodeResources, twentyFourNodeResources, thirtySixNodeResources, thirtySixNodeResources},
		{exponentialEstimator, 20, twentyFourNodeResources, twentyFourNodeResources, thirtySixNodeResources, thirtySixNodeResources},
		{exponentialEstimator, 24, twentyFourNodeResources, twentyFourNodeResources, thirtySixNodeResources, thirtySixNodeResources},
		{exponentialLessThanMilliEstimator, 3, sixteenNodeLessThanMilliExpResources, sixteenNodeLessThanMilliExpResources, twentyFourNodeLessThanMilliExpResources, twentyFourNodeLessThanMilliExpResources},
	}

	for _, tc := range testCases {
		expected, acceptable := tc.e.scaleWithNodes(tc.numNodes)
		verifyResources(t, "expected limits", expected.Limits, tc.expectedLimits)
		verifyResources(t, "expected requests", expected.Requests, tc.expectedRequests)
		verifyResources(t, "acceptable limits", acceptable.Limits, tc.acceptableLimits)
		verifyResources(t, "acceptable requests", acceptable.Requests, tc.acceptableRequests)
	}
}
