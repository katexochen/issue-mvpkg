package bar

import "github.com/katexochen/issue-mvpkg/foo"

type Bar struct {
	foo foo.Foo
}

func (b *Bar) Foo() {
	b.foo.Hello()
}
