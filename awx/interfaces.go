package awx

import "net/url"

// Connection interface
type AwxConnection interface {
	Jobs() IdGetter
	WorkflowJobs() IdGetter
	Close()
	authenticatedGet(string, url.Values, interface{}) error
	authenticatedPost(string, url.Values, interface{}, interface{}) error
}

// Request interfaces

type Sender interface {
	Send() (interface{}, error)
	Filter(name string, value interface{}) Sender
}

// Resource interfaces

type Getter interface {
	Get() Sender
}

type Poster interface {
	Post() Sender
}

type Launcher interface {
	Launch() GetterPoster
}

type Ider interface {
	Id(id int) Getter
}

type GetterPoster interface {
	Getter
	Poster
}

type GetterLauncher interface {
	Getter
	Launcher
}

type IdGetter interface {
	Getter
	Ider
}

type IdPoster interface {
	Poster
	Ider
}
