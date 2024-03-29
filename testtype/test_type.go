package testtype

type T interface {
	UserId()
}

type Default struct{}

func (d *Default) UserId() int32 {
	return 100
}

type User struct {
	*Default
	Id int32
}

func (_ *User) UserId() int32 {
	return 10001
}

func Run() {
	user := &User{}
	println(user.UserId())
}
