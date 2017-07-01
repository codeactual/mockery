package test

//go:generate mockery -name=Interface -case=underscore -inpkg -testonly

type Interface interface {
  Do()
}
