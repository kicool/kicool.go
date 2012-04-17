package dump

import (
	"fmt"
	"io"
	"os"
	r "reflect"
	"strconv"
)

var emptyString = ""

// Prints to the writer the value with indentation.
func Fdump(out io.Writer, v_ interface{}) {
	// forward decl
	var dump0 func(r.Value, int)
	var dump func(r.Value, int, *string, *string)

	done := make(map[string]bool)

	dump = func(v r.Value, d int, prefix *string, suffix *string) {
		pad := func() {
			res := ""
			for i := 0; i < d; i++ {
				res += "  "
			}
			fmt.Fprintf(out, res)
		}

		padprefix := func() {
			if prefix != nil {
				fmt.Fprintf(out, *prefix)
			} else {
				res := ""
				for i := 0; i < d; i++ {
					res += "  "
				}
				fmt.Fprintf(out, res)
			}
		}

		vk := v.Kind()
		vt := v.Type()
		vts := vt.String()
		vks := vk.String()

		vl := 0
		vc := 0
		switch vk {
		case r.Map, r.String:
			vl = v.Len()

		case r.Array, r.Chan, r.Slice:
			vc = v.Cap()
			vl = v.Len()
		
		case r.Struct:
			vl = v.NumField()

		default:
			//panic
		}

		// prevent circular for composite types
		switch vk {
		case r.Array, r.Slice, r.Map, r.Ptr, r.Struct, r.Interface:
			if v.CanAddr() {
				addr := v.Addr()
				key := fmt.Sprintf("%x %v", addr, v.Type())
				if _, exists := done[key]; exists {
					padprefix()
					fmt.Fprintf(out, "<%s>", key)
					return
				} else {
					done[key] = true
				}
			}
		default:
			// do nothing
		}

		switch vk {
		case r.Array:
			padprefix()
			fmt.Fprintf(out, "%s:%s (l=%d c=%d) {\n", vks, vts, vl, vc)
			for i := 0; i < vl; i++ {
				dump0(v.Index(i), d+1)
				if i != vl-1 {
					fmt.Fprintln(out, ",")
				}
			}
			fmt.Fprintln(out)
			pad()
			fmt.Fprint(out, "}")

		case r.Slice:
			padprefix()
			fmt.Fprintf(out, "%s:%s (l=%d c=%d) {\n", vks, vts, vl, vc)
			for i := 0; i < vl; i++ {
				dump0(v.Index(i), d+1)
				if i != vl-1 {
					fmt.Fprintln(out, ",")
				}
			}
			fmt.Fprintln(out)
			pad()
			fmt.Fprint(out, "}")

		case r.Map:
			padprefix()
			fmt.Fprintf(out, "%s:%s (l=%d){\n", vks, vts, vl)
			for i, k := range v.MapKeys() {
				dump0(k, d+1)
				fmt.Fprint(out, ": ")
				dump(v.MapIndex(k), d+1, &emptyString, nil)
				if i != vl-1 {
					fmt.Fprint(out, ",")
				}
			}
			fmt.Fprintln(out)
			pad()
			fmt.Fprint(out, "}")

		case r.Ptr:
			padprefix()
			if v.Elem() == r.ValueOf(nil) { //Zero Value
				fmt.Fprintf(out, "(*%s) nil", vts)
			} else {
				fmt.Fprintf(out, "ptr:*%s:&", r.Indirect(v).Type().String())
				dump(v.Elem(), d, &emptyString, nil)
			}

		case r.Struct:
			padprefix()
			fmt.Fprintf(out, "%s {\n", vts)
			d += 1
			for i := 0; i < v.NumField(); i++ {
				pad()
				fmt.Fprint(out, v.Field(i).Type())
				fmt.Fprint(out, ": ")
				dump(v.Field(i), d, &emptyString, nil)
				if i != v.NumField()-1 {
					fmt.Fprintln(out, ",")
				}
			}
			d -= 1
			fmt.Fprintln(out)
			pad()
			fmt.Fprint(out, "}")

		case r.Interface:
			padprefix()
			fmt.Fprintf(out, "(%s) ", vts)
			dump(v.Elem(), d, &emptyString, nil)

		case r.String:
			padprefix()
			fmt.Fprintln(out, strconv.Quote(v.String()))

		case r.Bool,
			r.Int, r.Int8, r.Int16, r.Int32, r.Int64,
			r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64,
			r.Float32, r.Float64:
			padprefix()
			//printv(o.Interface());
			i := v.Interface()
			if stringer, ok := i.(interface {
				String() string
			}); ok {
				fmt.Fprintf(out, "(%s) %s", vts, stringer.String())
			} else {
				fmt.Fprint(out, i)
			}

		case r.Invalid:
			padprefix()
			fmt.Fprint(out, "nil")

		default:
			padprefix()
			fmt.Fprintf(out, "(%v) %v", vt, v.Interface())
		}
	}

	dump0 = func(v r.Value, d int) { dump(v, d, nil, nil) }

	v := r.ValueOf(v_)
	dump0(v, 0)
	fmt.Fprintf(out, "\n")
}

// Print to standard out the value that is passed as the argument with indentation.
// Pointers are dereferenced.
func Dump(v_ interface{}) { Fdump(os.Stdout, v_) }
