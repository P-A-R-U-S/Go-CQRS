# CQRS Pattern for Go language [![Travis-CI](https://travis-ci.org/P-A-R-U-S/Golang-CQRS.svg?branch=master)](https://travis-ci.org/P-A-R-U-S/Golang-CQRS)  [![License](https://img.shields.io/dub/l/vibe-d.svg)](https://opensource.org/licenses/MIT)

## Getting Started

Project distributed as open-source example and you can can copy and modify code snippet based on you need or requirements.
Project contains following parts:


### Create handler(s)

```GO
const ExampleEvent = "_EventExample"

type ExampleHandler1 struct {
	_name, _event string
	_isOnSubscribeFired, _isOnUnsubscribeFired, _isExecuteFired bool
	_isPanicOnEvent, _isPanicOnOnSubscribe, _isPanicOnOnUnsubscribe, _isPanicOnExecute bool
	_isDisableMessage bool
}

func (h *ExampleHandler1) Event() string {

	return ExampleEvent
}
func (h *ExampleHandler1) Execute(... interface{}) error {

	fmt.Println("Run Execute...")

	return nil
}
func (h *ExampleHandler1) OnSubscribe() {
	fmt.Println("Run OnSubscribe...")
}
func (h *ExampleHandler1) OnUnsubscribe() {
	fmt.Println("Run OnUnsubscribe...")
}

```

### Add handler into the Bus and send message

```Go
func main()  {

	eventBus := bus.New()

	h := &ExampleHandler1{}

	eventBus.Subscribe(h)

	eventBus.Publish(ExampleEvent, 1, 2, "Test Message", 4.5)

	eventBus.Unsubscribe(ExampleEvent)
}
```



## Contributing

* [Valentyn Ponomarenko](http://valentynponomarenko.com)

The project intended to use as a part of part of othe project, but if you want to contribute, feel free to send pull requests!

Have problems, bugs, feature ideas?
We are using the github [issue tracker](https://github.com/P-A-R-U-S/Golang-CQRS/issues) to manage them.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/P-A-R-U-S/Golang-CQRS/tags). 

## Authors

* **Valentyn Ponomarenko** - *Initial work* - [P-A-R-U-S](https://github.com/P-A-R-U-S/)

See also the list of [contributors](https://github.com/P-A-R-U-S/Golang-CQRS/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](https://opensource.org/licenses/MIT) file for details

## Acknowledgments

* Hat tip to anyone who's code was used
* Inspiration
* etc



