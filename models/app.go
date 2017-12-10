package models

type Channel struct {
	Id         int
	Name       string
	FriendId   string
	Descript   string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
	Views      int
}

type Application struct {
	Id         int
	Name       string
	Descript   string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
	Views      int
}

type Pkh struct {
	Id         int
	Name       string
	Descript   string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
	Views      int
}

type Version struct {
	Id         int
	Name       string
	Descript   string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
	Views      int
}

type Code struct {
	Id         int
	Code       int
	Descript   string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
	Views      int
}

type Env struct {
	Id         int
	Name       string
	Descript   string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
	Views      int
}

type Build struct {
	Id         int
	Name       string
	Descript   string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
	Views      int
}

type Type struct {
	Id         int
	Name       string
	Descript   string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
	Views      int
}

type App struct {
	Id          int
	Type        string
	Application string
	Pkg         string
	Version     string
	Code        int
	Env         string
	Build       string
	Channel     string
	FriendId    string
	Descript    string
	Times       int
	Url         string
	Downs       int
	CreateId    int
	UpdateId    int
	CreateTime  int64
	UpdateTime  int64
	Views       int
}
