package scope

// @author:      zhangzhewei
// @create:      2021-11-02 21:22
// @description:


type Runnable func()
type ThrowableRunnable func() error

type Supplier func() interface{}
type ThrowableSupplier func() (interface{}, error)

type Consumer func(t interface{})
type ThrowableConsumer func(t interface{}) error

type Function func(t interface{}) interface{}
type ThrowableFunction func(t interface{}) (interface{}, error)

type Predicate func(t interface{}) bool
type ThrowablePredicate func(t interface{}) (bool, error)