package work12_1

import (
	"fmt"
	"reflect"
	"strconv"
)

// Any formats any value as a string.
func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

// formatAtom formats a value without inspecting its internal structure.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			// XXX: 咱也不道对不对
			if key.Kind() == reflect.Struct {
				for i := 0; i < v.NumField(); i++ {
					fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
					display(fieldPath, v.Field(i))
				}
			} else {
				display(fmt.Sprintf("%s[%s]", path, formatAtom(key)), v.MapIndex(key))
			}
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

func ExampleDisplay1() {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},

		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	Display("strangelove", strangelove)
	// format.Display("os.Stderr", os.Stderr)

	// Output:
	// Display strangelove (display.Movie):
	// strangelove.Title = "Dr. Strangelove"
	// strangelove.Subtitle = "How I Learned to Stop Worrying and Love the Bomb"
	// strangelove.Year = 1964
	// strangelove.Color = false
	// strangelove.Actor["Gen. Buck Turgidson"] = "George C. Scott"
	// strangelove.Actor["Brig. Gen. Jack D. Ripper"] = "Sterling Hayden"
	// strangelove.Actor["Maj. T.J. \"King\" Kong"] = "Slim Pickens"
	// strangelove.Actor["Dr. Strangelove"] = "Peter Sellers"
	// strangelove.Actor["Grp. Capt. Lionel Mandrake"] = "Peter Sellers"
	// strangelove.Actor["Pres. Merkin Muffley"] = "Peter Sellers"
	// strangelove.Oscars[0] = "Best Actor (Nomin.)"
	// strangelove.Oscars[1] = "Best Adapted Screenplay (Nomin.)"
	// strangelove.Oscars[2] = "Best Director (Nomin.)"
	// strangelove.Oscars[3] = "Best Picture (Nomin.)"
	// strangelove.Sequel = nil

}
func ExampleDisplay2() {
	var i interface{} = 3

	Display("i", i)
	// Output:
	// Display i (int):
	// i = 3
}
func ExampleDisplay3() {
	var i interface{} = 3
	Display("&i", &i)
	// Output:
	// Display &i (*interface {}):
	// (*&i).type = int
	// (*&i).value = 3
}
