//go:build wireinject
// +build wireinject
package simple

import (
	"io"
	"os"
	"github.com/google/wire"
)

func InitializedService(isError bool) (*SimpleService,error) {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil,nil
}

func InitializedDatabase() *DatabaseRepository{
	wire.Build(
		NewDatabaseMongoDB,
		NewDatabasePostgreSQL,
		NewDatabaseRepository,
	)
	return nil
}

var fooSet = wire.NewSet(NewFooRepository,NewFooService)

var barSet = wire.NewSet(NewBarRepository,NewBarService)

func InitializedFoobarService() *FooBarService{
	wire.Build(fooSet,barSet,NewFooBarService)
	return  nil
}


// func InitializedHelloService() *Helloservice{
// 	wire.Build(NewHelloService,NewSayHelloImpl)
// 	return nil  
// }

var helloSet= wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello),new(*SayHelloImpl)),
)

func InitializedHelloService() *Helloservice{
	wire.Build(helloSet,NewHelloService)
	return nil  
}

func InitializedFoobar() *FooBar{
	wire.Build(NewFoo,NewBar,wire.Struct(new(FooBar), "Foo","Bar"))
	return nil
}

var fooValue= &Foo{}
var barValue= &Bar{}

func InitializedFoobarValue() *FooBar{
	wire.Build(wire.Value(fooValue), wire.Value(barValue),wire.Struct(new(FooBar),"*"))
	return nil
}

func InitializedReader() io.Reader{
	wire.Build(wire.InterfaceValue(new(io.Reader),os.Stdin))
	return nil
}

func InitializedConfiguration() *Configuration{
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application),"Configuration"),
	)
	return nil
}

func InitializedConnection (name string) (*Connection, func()){
	wire.Build(NewConnection,NewFile)
	return nil,nil
}