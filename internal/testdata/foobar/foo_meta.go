// GENERATED BY metatag, DO NOT EDIT
// (or edit away - I'm a comment, not a cop)

package foobar

import (
	"reflect"
	"time"
	"fmt"
)

// NewFoo creates a new Foo with the given initial values.
func NewFoo(name string, desc string, labels []string) Foo {
	return Foo{
		name: name,
		Desc: desc,
		labels: labels,
	}
}

// Name returns the value of name.
func (f Foo) Name() string {
	return f.name
}

// GetDesc returns the value of Desc.
func (f Foo) GetDesc() string {
	return f.Desc
}

// String returns the "native" format of Foo. Implements the fmt.Stringer interface.
func (f Foo) String() string {
	return fmt.Sprintf("%v %v %v", f.name, f.Desc, f.size)
}

// Size returns the value of size.
func (f *Foo) Size() int {
	return f.size
}

// SetSize sets the given value as size.
func (f *Foo) SetSize(i int) {
	f.size = i
}

// SetLabels sets the given value as labels.
func (f *Foo) SetLabels(s []string) {
	f.labels = s
}

// Labels returns the value of labels.
func (f Foo) Labels() []string {
	return f.labels
}

// FilterLabels returns a copy of labels, omitting elements that are rejected by the given function.
// The n argument determines the maximum number of elements to return (n < 1: all elements).
func (f Foo) FilterLabels(fn func(string) bool, n int) []string {
	cap := n
	if n < 1 {
		cap = len(f.labels)
	}
	result := make([]string, 0, cap)
	for i := range f.labels {
		if fn(f.labels[i]) {
			if result = append(result, f.labels[i]); len(result) >= cap {
				return result
			}
		}
	}
	return result
}

// MapLabelsToTime returns a new slice with the results of calling the given function for each element of labels.
func (f Foo) MapLabelsToTime(fn func(string) time.Time) []time.Time {
	result := make([]time.Time, len(f.labels))
	for i := range f.labels {
		result[i] = fn(f.labels[i])
	}
	return result
}

// SetStringer sets the given value as stringer.
func (f *Foo) SetStringer(s fmt.Stringer) {
	f.stringer = s
}

// String returns the "native" format of Bar. Implements the fmt.Stringer interface.
func (b Bar) String() string {
	return fmt.Sprintf("%v", b.name)
}

// Equal answers whether v is equivalent to b.
// Always returns false if v is not a Bar.
func (b Bar) Equal(v interface{}) bool {
	b2, ok := v.(Bar)
	if !ok {
		return false
	}
	if b.name != b2.name {
		return false
	}
	if !reflect.DeepEqual(b.times, b2.times) {
		return false
	}
	return true
}

// Foos returns the value of foos.
func (b Bar) Foos() []Foo {
	return b.foos
}

// SetFoos sets the given value as foos.
func (b *Bar) SetFoos(f []Foo) {
	b.foos = f
}

// MapFoosToString returns a new slice with the results of calling the given function for each element of foos.
func (b Bar) MapFoosToString(fn func(Foo) string) []string {
	result := make([]string, len(b.foos))
	for i := range b.foos {
		result[i] = fn(b.foos[i])
	}
	return result
}

// Pairs returns the value of pairs.
func (b Bar) Pairs() map[string]float64 {
	return b.pairs
}

// SetPairs sets the given value as pairs.
func (b *Bar) SetPairs(m map[string]float64) {
	b.pairs = m
}

// Times returns the value of times.
func (b Bar) Times() []time.Time {
	return b.times
}

// SetTimes sets the given value as times.
func (b *Bar) SetTimes(t []time.Time) {
	b.times = t
}

// FilterTimes returns a copy of times, omitting elements that are rejected by the given function.
// The n argument determines the maximum number of elements to return (n < 1: all elements).
func (b Bar) FilterTimes(fn func(time.Time) bool, n int) []time.Time {
	cap := n
	if n < 1 {
		cap = len(b.times)
	}
	result := make([]time.Time, 0, cap)
	for i := range b.times {
		if fn(b.times[i]) {
			if result = append(result, b.times[i]); len(result) >= cap {
				return result
			}
		}
	}
	return result
}

// MapTimesToInt64 returns a new slice with the results of calling the given function for each element of times.
func (b Bar) MapTimesToInt64(fn func(time.Time) int64) []int64 {
	result := make([]int64, len(b.times))
	for i := range b.times {
		result[i] = fn(b.times[i])
	}
	return result
}

// SetBaz sets the given value as baz.
func (b *Bar) SetBaz(bb bool) {
	b.baz = bb
}
