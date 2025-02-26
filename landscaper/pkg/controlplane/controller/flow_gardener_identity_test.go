// Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

import (
	"context"

	"github.com/gardener/gardener/landscaper/pkg/controlplane/apis/imports"
	"github.com/gardener/gardener/pkg/client/kubernetes"
	"github.com/gardener/gardener/pkg/client/kubernetes/fake"
	"github.com/gardener/gardener/pkg/logger"
	mockclient "github.com/gardener/gardener/pkg/mock/controller-runtime/client"
	kutil "github.com/gardener/gardener/pkg/utils/kubernetes"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("#GetOrDefaultDiffieHellmannKey", func() {
	var testOperation operation

	// mocking
	var (
		ctx              = context.TODO()
		ctrl             *gomock.Controller
		mockGardenClient *mockclient.MockClient
		gardenClient     kubernetes.Interface
		errNotFound      = &apierrors.StatusError{ErrStatus: metav1.Status{Reason: metav1.StatusReasonNotFound}}
	)

	AfterEach(func() {
		ctrl.Finish()
	})

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockGardenClient = mockclient.NewMockClient(ctrl)

		gardenClient = fake.NewClientSetBuilder().WithClient(mockGardenClient).Build()

		testOperation = operation{
			log:           logrus.NewEntry(logger.NewNopLogger()),
			runtimeClient: gardenClient,
			imports:       &imports.Imports{},
		}
	})

	It("should generate a new identity - config map not found", func() {
		mockGardenClient.EXPECT().Get(ctx, kutil.Key("kube-system", "cluster-identity"), gomock.AssignableToTypeOf(&corev1.ConfigMap{})).Return(errNotFound)

		Expect(testOperation.GetOrGenerateIdentity(ctx)).ToNot(HaveOccurred())
		Expect(testOperation.imports.Identity).ToNot(BeNil())
		Expect(*testOperation.imports.Identity).To(ContainSubstring("landscape-"))
	})

	It("should generate a new identity - config map does not contain the expected key", func() {
		mockGardenClient.EXPECT().Get(ctx, kutil.Key("kube-system", "cluster-identity"), gomock.AssignableToTypeOf(&corev1.ConfigMap{})).DoAndReturn(
			func(ctx context.Context, _ client.ObjectKey, obj client.Object) error {
				(&corev1.ConfigMap{
					ObjectMeta: metav1.ObjectMeta{
						Generation: 1,
					},
					Data: map[string]string{
						"invalid": "not used",
					},
				}).DeepCopyInto(obj.(*corev1.ConfigMap))
				return nil
			},
		)

		Expect(testOperation.GetOrGenerateIdentity(ctx)).ToNot(HaveOccurred())
		Expect(testOperation.imports.Identity).ToNot(BeNil())
		Expect(*testOperation.imports.Identity).To(ContainSubstring("landscape-"))
	})

	It("should generate a new landscape identity", func() {
		expectedIdentity := "landscape-demo"
		mockGardenClient.EXPECT().Get(ctx, kutil.Key("kube-system", "cluster-identity"), gomock.AssignableToTypeOf(&corev1.ConfigMap{})).DoAndReturn(
			func(ctx context.Context, _ client.ObjectKey, obj client.Object) error {
				(&corev1.ConfigMap{
					ObjectMeta: metav1.ObjectMeta{
						Generation: 1,
					},
					Data: map[string]string{
						"cluster-identity": expectedIdentity,
					},
				}).DeepCopyInto(obj.(*corev1.ConfigMap))
				return nil
			},
		)

		Expect(testOperation.GetOrGenerateIdentity(ctx)).ToNot(HaveOccurred())
		Expect(testOperation.imports.Identity).To(Equal(&expectedIdentity))
	})
})
