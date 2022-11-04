/*
Copyright 2022.

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

package controllers

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	messagev1 "example.com/message-controller/api/v1"
)

// MessageReconciler reconciles a Message object
type MessageReconciler struct {
	client.Client // kube-apiserverとやりとりするためのクライアント(client-go的なやつ)
	Scheme        *runtime.Scheme
}

//+kubebuilder:rbac:groups=message.example.com,resources=messages,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=message.example.com,resources=messages/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=message.example.com,resources=messages/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Message object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
// MessageReconcilerのメソッドとして定義(controller-runtimeで定義されたインターフェイスに対する実装)
func (r *MessageReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	defer fmt.Println("=== Finish Reconcile ===")
	fmt.Println("=== Start Reconcile ===")

	log := log.FromContext(ctx)

	var message messagev1.Message
	var err error

	// cacheからmessageを取得する
	// message_types.goのMessage structで定義した形で取得する
	if err = r.Get(ctx, req.NamespacedName, &message); err != nil {
		log.Error(err, "Faild to fetch message.")
		return ctrl.Result{}, client.IgnoreNotFound(err) // 取得に失敗したらエラーをスキップする
	}

	// [Debug] 取得したmessageを表示する
	// fmt.Println(message)

	// [Debug] cacheからmessage一覧を取得する
	// var messageList messagev1.MessageList
	// if err = r.List(ctx, &messageList); err != nil {
	// 	log.Error(err, "Faild to fetch messageList.")
	// 	return ctrl.Result{}, client.IgnoreNotFound(err)
	// }

	// [Debug] MessageListに含まれるmessage Nameを表示
	// for i, messageItem := range messageList.Items {
	// 	fmt.Println(i, ": ", messageItem.Name, " ", messageItem.Spec.Word)
	// }

	// Flags
	wordFlag := false
	numberFlag := false

	// messageのspec.wordとstatus.wordを取得する
	specWord := message.Spec.Word
	statusWord := message.Status.Word
	templateWord := "Hello "

	// Debug
	// fmt.Println(message)

	// specWordとstatusWordに差分があればstatusWordを更新する
	if statusWord != templateWord+specWord {
		message.Status.Word = templateWord + specWord
		wordFlag = true
	}

	// messageのspec.numberとstatus.numberを取得する
	specNumber := int32(0) // 初期値
	if message.Spec.Number != nil {
		specNumber = *message.Spec.Number
	}

	statusNumber := message.Status.Number

	// specNumberとstatusNumberに差分があればstatusNumberを更新する
	if specNumber != statusNumber {
		message.Status.Number = specNumber
		numberFlag = true
	}

	// message更新
	if wordFlag || numberFlag {
		log.Info("Update message")
		if err = r.Status().Update(ctx, &message); err != nil {
			log.Error(err, "Faild to update message.")
			return ctrl.Result{}, err
		}

		// Debug
		log.Info("[Debug] Status.Word: " + message.Status.Word)
		fmt.Println("[Debug] Status.Word: " + message.Status.Word)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MessageReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr). // &Builderのメンバーにmgrを登録
							For(&messagev1.Message{}). // reconcile対象のオブジェクトを指定
							Complete(r)                // controllerがbuildされる
}
