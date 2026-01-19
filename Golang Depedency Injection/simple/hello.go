package simple

type SayHello interface{
	Hello(name string) string
}

type Helloservice struct{
	SayHello SayHello
}

type SayHelloImpl struct{

}

func (s SayHelloImpl) Hello(name string) string{
	return  "Hello" +name
}

func NewSayHelloImpl () *SayHelloImpl{
	return  &SayHelloImpl{}
}

func NewHelloService (sayHello SayHello) *Helloservice{
	return  &Helloservice{SayHello: sayHello}
}



