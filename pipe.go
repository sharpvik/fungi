package fungi

type pipe[R any] struct {
	Result func() R
}

// Pipe is a style-enhancing utility that allows one to apply functions in a
// left-to-right fashion like so:
//
//	err := Then(Then(
//	    Pipe(source, filter), mapper), looper).
//	    //   (arg) >>> (1) >>>> (2) >>>> (3)
//	    Result()
func Pipe[I, O any](x I, f func(I) O) *pipe[O] {
	return &pipe[O]{
		Result: func() O { return f(x) },
	}
}

// Then chains another function call onto a Pipe.
func Then[I, O any](p *pipe[I], f func(I) O) *pipe[O] {
	return &pipe[O]{
		Result: func() O { return f(p.Result()) },
	}
}
