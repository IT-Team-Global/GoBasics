package main

import "fmt"

//Tako definiramo podatkovno strukturo
type Controller struct {
	Number      string
	Gateway     string
	GatewayPort string
	AccessToken string
}

//Tako definiramo podatkovno strukturo, ki ima za vsak field določen kako se bo serializiral v JSON in deserializiral iz JSON-a
type Message struct {
	ToNumber string `json:"to_number"`
	Message  string `json:"message"`
	FromName string `json:"from_name"`
}

//Tako se doda funkcija na neko podatkovno strukturo/objekt
func (c *Controller) Init() (err error) {

	//Inicializacija nekega SMS gateway-a

	return nil

}

//Tako se doda funkcija na neko podatkovno strukturo/objekt
func (c *Controller) SendSMS(message Message) (err error) {

	//Koda za pošiljanje SMS-a

	return nil

}

//Primer treh načinov inicializacije objektov in prikaz uporabe funkcij obešenih/dodanih na objekt
func main() {

	//Naredimo Controller objekt
	smsController := Controller{
		Number:      "12345",
		Gateway:     "http://gateway.com",
		GatewayPort: "1234",
		AccessToken: "fsdklfjsvmpmgpiooniuvsn",
	}

	err := smsController.Init()
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	myMessage := Message{"051123456", "Hello World from SMS", "ITTIM"}

	err = smsController.SendSMS(myMessage)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	var myMessage2 Message
	myMessage2.ToNumber = "051123456"
	myMessage2.Message = "Hello World from SMS"
	myMessage2.FromName = "ITTIM"

	err = smsController.SendSMS(myMessage2)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

}
