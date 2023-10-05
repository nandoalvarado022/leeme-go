// El problema: Diseñar de software que sea capaz de enviar notificaciones sms y email utilizando polimorfismo e interfaces.
// El patron de diseño factory nos permite centralizar en una sola función la creación de estructuras especificas de un mismo tipo.
// Con esto el cliente no tiene que conocer toda la causistica que tenemos que generar para entregart un elemento de un tipo especifico.

package abstractfactory

import "fmt"

// SMS Email

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending notification via SMS")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

func main() {

}
