package errx

import (
	"bytes"
	"fmt"
	"io"
	"path"
	"runtime"
	"strconv"
)

// stack represents a stack of program counters.
type stack []uintptr

// Format stack trace
func (s *stack) Format(fs fmt.State, verb rune) {
	switch verb {
	// case 'v', 's':
	case 'v':
		_, _ = s.WriteTo(fs)
	}
}

// StackLen for error
func (s *stack) StackLen() int {
	return len(*s)
}

// WriteTo for error
func (s *stack) WriteTo(w io.Writer) (int64, error) {
	if len(*s) == 0 {
		return 0, nil
	}

	nn, _ := w.Write([]byte("\nSTACK:\n"))
	for _, pc := range *s {
		// For historical reasons if pc is interpreted as a uintptr
		// its value represents the program counter + 1.
		fc := runtime.FuncForPC(pc - 1)
		if fc == nil {
			continue
		}
		file, line := fc.FileLine(pc - 1)
		location := fc.Name() + "()\n  " + file + ":" + strconv.Itoa(line) + "\n"
		n, _ := w.Write([]byte(location))
		nn += n
	}

	return int64(nn), nil
}

// String format to string
func (s *stack) String() string {
	var buf bytes.Buffer
	_, _ = s.WriteTo(&buf)
	return buf.String()
}

// StackFrames stack frame list
func (s *stack) StackFrames() *runtime.Frames {
	return runtime.CallersFrames(*s)
}

// CallerPC the caller PC value in the stack. it is first frame.
func (s *stack) CallerPC() uintptr {
	if len(*s) == 0 {
		return 0
	}

	// For historical reasons if pc is interpreted as a uintptr
	// its value represents the program counter + 1.
	return (*s)[0] - 1
}

/*************************************************************
 * For error caller func
 *************************************************************/

// Func struct
type Func struct {
	*runtime.Func
	pc uintptr
}

// FuncForPC create.
func FuncForPC(pc uintptr) *Func {
	fc := runtime.FuncForPC(pc)
	if fc == nil {
		return nil
	}

	return &Func{
		pc:   pc,
		Func: fc,
	}
}

// FileLine of the func
func (f *Func) FileLine() (file string, line int) {
	return f.Func.FileLine(f.pc)
}

// Location simple location info for the func
func (f *Func) Location() string {
	file, line := f.FileLine()

	return f.Name() + "(), " + path.Base(file) + ":" + strconv.Itoa(line)
}

// String of the func
func (f *Func) String() string {
	file, line := f.FileLine()

	return f.Name() + "()\n  At " + file + ":" + strconv.Itoa(line)
}

// MarshalText handle
func (f *Func) MarshalText() ([]byte, error) {
	return []byte(f.String()), nil
}

/*************************************************************
 * helper func for callers stacks
 *************************************************************/

// ErrStackOpt struct
type ErrStackOpt struct {
	SkipDepth  int
	TraceDepth int
}

// ????????????
var stdOpt = newErrOpt()

// ResetStdOpt ????????????
func ResetStdOpt() {
	stdOpt = newErrOpt()
}

func newErrOpt() *ErrStackOpt {
	return &ErrStackOpt{
		SkipDepth:  3,
		TraceDepth: 8,
	}
}

// Config ?????? stdOpt ??????
func Config(fns ...func(opt *ErrStackOpt)) {
	for _, fn := range fns {
		fn(stdOpt)
	}
}

// SkipDepth ??????????????????
func SkipDepth(skipDepth int) func(opt *ErrStackOpt) {
	return func(opt *ErrStackOpt) {
		opt.SkipDepth = skipDepth
	}
}

// TraceDepth ??????????????????
func TraceDepth(traceDepth int) func(opt *ErrStackOpt) {
	return func(opt *ErrStackOpt) {
		opt.TraceDepth = traceDepth
	}
}

func callersStack(skip, depth int) *stack {
	pcs := make([]uintptr, depth)
	num := runtime.Callers(skip, pcs[:])

	var st stack = pcs[0:num]
	return &st
}
