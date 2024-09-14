package fail

type IError interface {
	GetData() map[string]interface{}
	GetError() string
	GetResponseCode() int
	ToError() error
}