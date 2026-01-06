//go:build wireinject
// +build wireinject

package simple

import (
	"io"
	"os"

	"github.com/google/wire"
)

// membuat function untuk membuat service
// jikalau provider memiliki parameter, maka cukup tambahkan saja, maka nanti google wire akan memetakkan-
// parameter tersebut ke tempat yang sesuai
func InitializedService(isError bool) (*SimpleService, error) {
	// pengaturan untuk inject dari object mana ke object yang lain bisa diatur di sini
	// dengan menggunakan function build pada package wire
	// memberi tahu function provider mana saja yang akan digunakan untuk membuat dependency injection nya
	wire.Build(
		// untuk membuat service kita perlu
		// google wire akan otomatis menghandle function mana yang membutuhkan dependency di function lain
		NewSimpleRepository, NewSimpleService, 
	)

	// karena nanti kode nya akan digenerate otomatis oleh google wire
	return nil, nil 
}

// membuat function injector baru untuk database
func InitializedDatabaseRepository() *DatabaseRepository {

	wire.Build(
		NewDatabasePostgreSQL, NewDatabaseMongoDB, NewDatabaseRepository,
	)

	return nil
}

// melakukan grouping dengan provider set : untuk mempermudah pemanggilan
var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

// membuat function injector baru untuk foo bar service
func InitializedFooBarService() *FooBarService {

	wire.Build(
		fooSet, barSet, NewFooBarService,
	)

	return nil
}

// injector untuk hello 

// contoh yang salah
// func InitializedHelloService() *HelloService {
// 	// akan error pada saat di build, karena provider tidak menemukan interface sayHello,-
// 	// juga karena provider NewSayHelloImpl mereturn sebuah struct
// 	wire.Build(NewHelloService, NewSayHelloImpl)
	
// 	return nil
// }

// contoh yang benar
// agar injector bisa menerima provider dengan interface, maka perlu dibuatkan provider set-
// dengan menambahkan wire.binding
var helloSet = wire.NewSet(
	NewSayHelloImpl, 
	wire.Bind(new(SayHello), new(*SayHelloImpl)), // kalau ada yang membutuhkan interface SayHello, maka nanti akan menerapkan *SayHelloImpl
)

func InitializedHelloService() *HelloService {
	// dimana NewHelloService butuh interface SayHeloo-
	// maka, nanti akan memanggil wire binding di provider set sebelumnya
	wire.Build(helloSet, NewHelloService) 

	return nil
}

// implementasi untuk struct provider
func InitializedFooBar() *FooBar {
	wire.Build(
		NewFoo, 
		NewBar, 
		// disebutkan untuk attribute mana saja yang akan di inject
		// kalau mau disebutkan semua bisa pakai bintang
		wire.Struct(new(FooBar), "*"),  
	)

	return nil
}

// implementasi untuk binding values
// membuat nilai variabel yang sudah ada, dan nantinya akan digunkan
var fooValue = &Foo{}
var barValue = &Bar{}

func InitializedFooBarUsingValues() *FooBar {
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(FooBar), "*"),)
	return nil
}

// implementasi untuk binding values: dengan interface values
func InitializedReader() io.Reader {
	// jika ada yang mebutuhkan interface io reader, maka akan menginject os stdin
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

// implementasi untuk struct field provider
func InitializedConfiguration() *Configuration {
	wire.Build(
		NewApplication,
		// karena return pointer, maka juga pointer untuk application
		// kemudian karena data yang di return adalah configuration maka, ambil untuk-
		// field Configuration pada struct Application
		wire.FieldsOf(new(*Application), "Configuration"),
	)
	return nil
}

// implementasi untuk cleanup function / closure
// karena injector menggunakan closure, maka juga tambahkan return value anonymous function (closure)
func InitializedConnection(name string) (*Connection, func()) {
	// memanggil kedua provider
	wire.Build(NewConnection, NewFile)
	return nil, nil
}