package solid

import "io"

/*
单一职责原则 SRP

SRP (The Single Responsibility Principle)
There should never be more than one reason for a class to change.
每个类应该只有一个改变的原因。换句话说，一个类应该只负责一项职责。
解析：任何一个软件模块都应该只对某一类行为者负责。一个功能模块应该只做一件事或一类事。
举例：
对于字符串的处理函数应该都放在strings包（模块）下；
对于数学概念定义的函数、常数、公式等都应该放在math包（模块）下；

如果一个类有多个职责，则应该将这些职责分离到不同类中，以避免紧耦合和高复杂度。
*/

/*
开放封闭原则

OCP (Open/Closed Principle)
Software entities should be open for extension，but closed for modification.
软件实体应当对扩展开放，对修改关闭。
解析：软件实体（类、模块、函数等）应该是可扩展的，但不应被修改。这意味着当需求变化时，我们应该通过添加新的代码来扩展系统的功能，而不是修改已有的代码，以保持系统的稳定性。
OCP是系统框架设计的主导原则，其主要目的是让系统易于扩展，同时限制其每次被修改所影响的范围。实现的方式是通过将系统划分为一系列的组件，并且将这些组件的依赖关系按层次结构进行组织，使得高阶的组件不会因低阶组件被修改而受到影响。
*/

type Teacher interface {
	Teach()   // 教书
	Educate() // 育人
	//Review()  // 考评
}

type MathTeacher struct {
	Teacher
}

/*
里氏替换原则

LSP (Liskov Substitution Principle)
Functions that use pointers or references to base classes must be able to use objects of derived classes without knowing it.
使用基类对象指针或引用的函数必须能够在不了解衍生类的条件下使用衍生类的对象。

里氏代换原则中说，任何基类可以出现的地方，子类一定可以出现。只有当衍生类可以替换掉基类，软件单位的功能不受到影响时，基类才能真正被复用，而衍生类也能够在基类的基础上增加新的行为。
*/

type Base struct{}

func (b *Base) Move(x, y int)  {}
func (b *Base) Speed() float64 { return 0 }

type Bird struct {
	*Base
}

func (b *Bird) Jump() {}

func Do(b *Base /* *Bird */) {
	b.Speed()
}

/*
接口隔离原则

ISP (Interface Segregation Principle)
Clients should not be forced to depend upon interfaces that they do not use.
不应强制客户端依赖于它们不使用的接口。该原则还有另外一个定义：一个类对另一个类的依赖应该建立在最小的接口上（The dependency of one class to another one should depend on the smallest possible interface）

接口隔离原则和单一职责都是为了提高类的内聚性、降低它们之间的耦合性，体现了封装的思想，但两者是不同的：
-单一职责原则注重的是职责，而接口隔离原则注重的是对接口依赖的隔离。单一职责原则主要是约束类，它针对的是程序中的实现和细节；
-接口隔离原则主要约束接口，主要针对抽象和程序整体框架的构建。
*/

type MyReader interface {
	io.Reader
}

type MyWriter interface {
	io.Writer
	io.Closer
}

type MyReadWriteCloser interface {
	MyReader
	MyWriter
}

/*
依赖倒置原则

DIP (Dependency Inversion Principle)
High level modules should not depend upon low level modules. Both should depend upon abstractions.
高层次的模块不应该依赖低层次的模块，他们都应该依赖于抽象。

Abstractions should not depend upon details. Details should depend upon abstractions.
抽象不应该依赖于具体实现，具体实现应该依赖于抽象。

高层次的模块应该依赖低层次接口而不是具体的结构对象。

依赖倒转原则就是指：代码要依赖于抽象的类，而不要依赖于具体的类；要针对接口或抽象类编程，而不是针对具体类编程。
*/

// https://blog.csdn.net/q_17600689511/article/details/102893103
