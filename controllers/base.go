package controllers

type Controller interface {
	Construct(arg ...interface{}) interface{}
}

// @todo
//
// type Base struct{}

// func (*Base) Construct(args ...interface{}) interface{} {
// 	return &Base{}
// }

// func (*Base) Router(ctrl interface{}) (int, []byte) {
// 	log.Println(reflect.TypeOf(ctrl))

// 	return 202, []byte{}
// }
