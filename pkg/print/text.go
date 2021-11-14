package print

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-yaml/yaml"
	"io"
	"reflect"
	"strings"
	"text/tabwriter"
)

const (
	JSON PayloadType = iota
	YAML
	TEXT
)

type PayloadType uint

func PayloadTypeFromString(s string) PayloadType {
	switch s {
	case "json":
		return JSON
	case "yaml":
		return YAML
	case "text":
		return TEXT
	default:
		return TEXT
	}
}

type printer interface {
	Print(ctx context.Context, w io.Writer, pt PayloadType, v interface{}) error
}

func printYAML(w io.Writer, v interface{}) error {
	enc := yaml.NewEncoder(w)
	return enc.Encode(v)
}

func printJSON(w io.Writer, v interface{}) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

func printText(writer io.Writer, u interface{}) error {
	w := new(tabwriter.Writer)
	w.Init(writer, 10, 0, 2, '.', tabwriter.TabIndent)
	ru := reflect.ValueOf(u)
	ut := ru.Type()
	if ut.Kind() != reflect.Struct {
		return fmt.Errorf("not struct")
	}

	for i := 0; i < ut.NumField(); i++ {
		field := ut.Field(i)
		fieldTagName := field.Tag.Get("print")
		if fieldTagName != "" {
			fv := ru.Field(i)
			fieldTagValue := fv.Interface()
			ft := field.Type
			if ft.Kind() == reflect.Struct {
				_, _ = fmt.Fprintf(w, "%s:\t\n", fieldTagName)
				for j := 0; j < ft.NumField(); j++ {
					fld := ft.Field(j)
					fldTagName := fld.Tag.Get("print")
					if fldTagName != "" {
						fldTagValue := fv.Field(j).Interface()
						_, _ = fmt.Fprint(w, "    ", fldTagName, "\t", fldTagValue, "\n")
					}
				}

				continue
			}
			_, err := fmt.Fprint(w, fieldTagName, "\t", fieldTagValue, "\n")
			if err != nil {
				return fmt.Errorf("could not print: %w", err)
			}
		}
	}

	_, _ = fmt.Fprintln(w)
	err := w.Flush()
	if err != nil {
		return fmt.Errorf("could not flush tabwriter :%w", err)
	}

	return nil
}

// Out takes the payload and print it to the output in the
// specified format using the writer
// if the format chosen is print do pass a struct not a pointer
// pass a pointer if the format chosen is either json or yaml
func Out(ctx context.Context,prefix string, writer io.Writer, pt PayloadType, v interface{}) error {
	_, _ = writer.Write([]byte(fmt.Sprintf("\n%s\n-------------------------------\n",
		strings.ToUpper(prefix))))
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	switch pt {
	case JSON:
		return printJSON(writer, v)

	case YAML:
		return printYAML(writer, v)

	default:
		return printText(writer, v)
	}
}
