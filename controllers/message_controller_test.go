package controllers

import (
	"context"
	"time"

	messagev1 "example.com/message-controller/api/v1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// +kubebuilder:docs-gen:collapse=Imports

var _ = Describe("message controller", func() {

	// Define utility constants for object names and testing timeouts/durations and intervals.
	const (
		TestMessageName = "test-message"
		TestNamespace   = "test"
		TestWord        = "Test"
		TestNumber      = int32(1)

		timeout  = time.Second * 10
		duration = time.Second * 10
		interval = time.Millisecond * 250
	)

	// 各テストケース実行前に呼び出される
	BeforeEach(func() {
		err := k8sClient.DeleteAllOf(ctx, &messagev1.Message{}, client.InNamespace(TestNamespace)) // Namespaceに存在するすべてのmessageリソースを削除
		Expect(err).NotTo(HaveOccurred())
		time.Sleep(interval)
	})

	// 各テストケース実行後に呼び出される
	AfterEach(func() {
		By("tearing down the test environment")
		time.Sleep(100 * time.Millisecond)
	})

	// テストケースの大枠的なもの
	// この中で複数のテストケース(It)を定義する
	Context("When updating Message Status", func() {

		// テストケース(It)を記述
		It("Should be set Message Status.Word when new Message is created", func() {
			By("By creating a new Message") // テストケースの説明(メモみたいなもの)

			num := int32(1)

			ctx := context.Background()

			// Messageリソースの作成
			message := messagev1.Message{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "message.example.com/v1",
					Kind:       "Message",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      TestMessageName,
					Namespace: TestNamespace,
				},
				Spec: messagev1.MessageSpec{
					Word:   TestWord,
					Number: &num,
				},
			}

			// Messageリソースを作成
			err := k8sClient.Create(ctx, &message)
			Expect(err).NotTo(HaveOccurred())

			messageLookupKey := types.NamespacedName{Name: TestMessageName, Namespace: TestNamespace}
			createdMessage := messagev1.Message{}

			// Messageリソースが作成されてStatusが更新されるのを待つ
			time.Sleep(5000 * time.Millisecond)

			// Messageリソースを取得
			// 引数に与えられた関数を所定の間隔で繰り返す
			Eventually(func() bool {
				err := k8sClient.Get(ctx, messageLookupKey, &createdMessage)
				if err != nil {
					return false
				}
				return true
			}, timeout, interval).Should(BeTrue())

			// 作成したMessageリソースのspec.wordが正確に登録されていることを確認
			By("By checking the Message spec.word has set the correct message")
			Expect(createdMessage.Spec.Word).Should(Equal(TestWord))

			// Controllerの挙動を確認
			// 作成したMessageリソースのstatus.wordが正確に更新されていることを確認
			By("By checking the Message status.word has set the correct message")
			Expect(createdMessage.Status.Word).Should(Equal("Hello " + TestWord))
		})

	})

})
