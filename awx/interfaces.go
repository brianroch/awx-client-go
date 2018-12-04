package awx

import (
	"net/http"
	"net/url"
)

// Connection interface
type AwxConnection interface {
	authenticatedGet(path string, query url.Values, output interface{}) error
	authenticatedPost(path string, query url.Values, input interface{}, output interface{}) error
	Jobs() IdGetter
	WorkflowJobs() IdGetter
	Client() *http.Client
	Close()
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
