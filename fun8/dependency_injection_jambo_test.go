package fun8

import (
	"bytes"
	"testing"
)

func TestMusalimie(t *testing.T) {
	buffer := bytes.Buffer{}
	Musalimie(&buffer, "Chris")

	got := buffer.String()
	want := "Jambo, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

/*
Using fmt.Printf() was hard to test because it wrote data to somewhere we couldn't control
Therefore, we refactored the code so we could control where the data was written by injecting
a dependency which allowed us:
- test our code:  If you can't test a function easily, it's usually because of dependencies hard-wired
  into a function or global state.
  If you have a global database connection pool for instance that is used by some
  kind of service layer, it is likely going to be difficult to test and they will be
  slow to run. DI will motivate you to inject in a database dependency (via an interface)
  which you can then mock out with something you can control in your tests.

- Separate our concerns: decoupling where the data goes from how to generate it.
  If you ever feel like a method/function has too many responsibilities (generating
  data and writing to a db? handling HTTP requests and doing domain level logic?)
  DI is probably going to be the tool you need.

- Allow our code to be re-used in different contexts:The first "new" context our code can be used in is inside tests. But further on if someone wants to try something new
  with your function they can inject their own dependencies.
*/
