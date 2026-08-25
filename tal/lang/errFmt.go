package lang

import (
	goerr "errors"
	"fmt"
	"strings"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing/stringParsing/parser3"
	"github.com/pt-main/lc/public/errors"
	"github.com/pt-main/run/tal/shared"
	"github.com/pt-main/tap/color"
)

var (
	// whereSpace is a colored prefix for indentation (blue line).
	whereSpace = color.Set("  [?BE]|[?RT]")
	// whereRedSpace is a colored prefix for error context (red line).
	whereRedSpace = color.Set("  [?RD]|[?RT]")
	// errSpace is a colored prefix for error messages (red arrow).
	errSpace = color.Set(" [?RD]>[?RT] ")
)

func GetRealErrorReverse(err error) string {
	if err == nil {
		return ""
	}

	if ce, ok := err.(core.ErrorInterface); ok {
		return ce.Format()
	}

	var parts []string
	cur := err
	for cur != nil {

		if ce, ok := cur.(core.ErrorInterface); ok {

			innerFormatted := ce.Format()
			if len(parts) > 0 {

				outerMsg := strings.Join(reverse(parts), ": ")
				return outerMsg + ": " + innerFormatted
			}
			return innerFormatted
		}

		parts = append(parts, cur.Error())

		next := goerr.Unwrap(cur)
		if next == nil {
			break
		}
		cur = next
	}

	if len(parts) > 0 {
		return strings.Join(reverse(parts), ": ")
	}
	return err.Error()
}

func reverse(s []string) []string {
	res := make([]string, len(s))
	for i, v := range s {
		res[len(s)-1-i] = v
	}
	return res
}

// GetErr extracts the inner error if it is a core.ErrorInterface.
// If the inner error is not a core.ErrorInterface, it wraps it into a core.Error
// with SystemError code.
func GetErr(ei core.ErrorInterface) core.ErrorInterface {
	inner := ei.Unwrap()
	if inner == nil {
		return nil
	}
	res, ok := inner.(core.ErrorInterface)
	if !ok {
		res = &core.Error{
			Code:  shared.SystemError,
			Msg:   inner.Error(),
			Meta:  make(map[errors.ErrorMetaType]interface{}),
			Cause: nil,
		}
	}
	return res
}

// addSpace prefixes each line of the given string with a repeated space pattern.
func addSpace(code, space string, n int) string {
	prefix := strings.Repeat(space, n)
	return prefix + strings.ReplaceAll(code, "\n", "\n"+prefix)
}

func ErrFmt(ei core.ErrorInterface) string {
	return FormatError(ei, nil)
}

// FormatError formats a core.ErrorInterface into a human-readable colored string.
// It handles known error codes (SystemError, GenerationError, ParsingError, parser3 errors)
// and falls back to a generic output for unknown codes.
func FormatError(ei core.ErrorInterface, prev core.ErrorInterface) string {
	if ei == nil {
		return ""
	}

	inner := GetErr(ei)
	meta := ei.GetMeta()
	code := errors.ErrorCodeType(ei.GetCode())
	addFallback := false
	fallbackAdded := false

	var res strings.Builder

	switch code {
	case shared.SystemError:
		res.WriteString(color.Set("[?YW]System error:[?RT]\n"))
		res.WriteString(addSpace(ei.GetMsg(), errSpace, 1))

	case shared.GenerationError:
		res.WriteString(color.Set("[?YW]Generation error:[?RT]\n"))
		res.WriteString(addSpace(ei.GetMsg(), errSpace, 1))

	case errors.ParsingError, parser3.AdapterErrCode:
		fallback := func() {
			res.WriteString(color.Set("[?YW]Parsing error:[?RT]\n"))
			res.WriteString(addSpace(ei.GetMsg(), errSpace, 1))
		}
		if inner != nil {
			if inner.GetCode() == parser3.AdapterErrCode {
				fallbackAdded = true
				FormatError(inner.Unwrap().(core.ErrorInterface), ei)
			} else {
				fallback()
			}
		} else {
			fallback()
		}

	case parser3.ParseErrCode, parser3.GrammarErrCode:
		res.WriteString(color.Set("[?YW]Parser error (2):[?RT]\n"))
		text := ""
		switch v, _ := meta["Code"].(string); v {
		case "UnexpectedToken":
			expected, _ := meta["Expected"].(string)
			got, _ := meta["Got"].(string)
			raw, _ := meta["Raw"].(string)
			text += fmt.Sprintf("Expected '%s', got '%s'", expected, got)
			if raw != "" {
				text += "\n" + addSpace(raw, whereSpace, 1)
			}
		default:
			msg := ei.GetMsg()
			if v != "" {
				text += v + ":"
				if msg != "" {
					text += "\n"
				}
			}
			if msg != "" {
				text += msg
			}
		}
		if text != "" {
			res.WriteString(addSpace(text, errSpace, 1))
		}
	default:
		addFallback = true
	}

	fallback := func() {
		// Generic fallback
		if !addFallback {
			return
		}
		result := res.String()
		if len(result) > 0 && result[len(result)-1] != '\n' {
			res.WriteRune('\n')
		}
		res.WriteString(color.Set("[?BYW]Error:[?YW] " + ei.GetCode() + "[?RT]"))
		msg := ei.GetMsg()
		if msg != "" {
			res.WriteRune('\n')
			res.WriteString(addSpace(msg, errSpace, 1))
		}
	}

	if ei.GetCode() == "RepeatExpr" && prev != nil {
		if prev.GetCode() == "NodeExpr" && prev.GetMsg() == "building node 'file'" &&
			ei.GetMsg() == "expected at least 1 repetition(s), got 0 at idx=0 start=0-1" {
			// while parser in node file and repeats of blocks is not found
			res.WriteString("Do you forget to add block annotation?")
		} else {
			fallback()
		}
	} else {
		fallback()
	}

	if inner != nil {
		if !fallbackAdded {
			res.WriteString("\n")
		}
		result := FormatError(inner, ei)
		res.WriteString(addSpace(result, whereRedSpace, 1))
	}

	return res.String()
}
