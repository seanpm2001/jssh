package jsshcmd

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/leizongmin/go/typeutil"
	"github.com/leizongmin/jssh/internal/jsexecutor"
	"log"
	"os"
)

var logPrefix = fmt.Sprintf("[%s] ", cmdName)
var stdLog = log.New(os.Stdout, logPrefix, log.LstdFlags)
var errLog = log.New(os.Stderr, logPrefix, log.LstdFlags)

func JsFnPrint(global typeutil.H) jsexecutor.JSFunction {
	return func(ctx *jsexecutor.JSContext, this jsexecutor.JSValue, args []jsexecutor.JSValue) jsexecutor.JSValue {
		if len(args) > 0 {
			s, err := jsexecutor.JSValueToAny(args[0])
			if err != nil {
				return ctx.ThrowError(err)
			}
			format, ok := s.(string)
			a := make([]interface{}, 0)
			for _, v := range args[1:] {
				v2, err := jsexecutor.JSValueToAny(v)
				if err != nil {
					return ctx.ThrowError(err)
				}
				a = append(a, v2)
			}
			if ok {
				fmt.Printf(format, a...)
			} else {
				a = append([]interface{}{s}, a...)
				fmt.Println(a)
			}
		}
		return ctx.Bool(true)
	}
}

func JsFnPrintln(global typeutil.H) jsexecutor.JSFunction {
	return func(ctx *jsexecutor.JSContext, this jsexecutor.JSValue, args []jsexecutor.JSValue) jsexecutor.JSValue {
		ret := global["print"].(jsexecutor.JSFunction)(ctx, this, args)
		fmt.Println()
		return ret
	}
}

func JsFnLogInfo(global typeutil.H) jsexecutor.JSFunction {
	green := color.FgGreen.Render
	return func(ctx *jsexecutor.JSContext, this jsexecutor.JSValue, args []jsexecutor.JSValue) jsexecutor.JSValue {
		if len(args) > 0 {
			s, err := jsexecutor.JSValueToAny(args[0])
			if err != nil {
				return ctx.ThrowError(err)
			}
			format, ok := s.(string)
			a := make([]interface{}, 0)
			for _, v := range args[1:] {
				v2, err := jsexecutor.JSValueToAny(v)
				if err != nil {
					return ctx.ThrowError(err)
				}
				a = append(a, v2)
			}
			if ok {
				stdLog.Printf(green(format), a...)
			} else {
				a = append([]interface{}{s}, a...)
				stdLog.Println(green(a))
			}
		}
		return ctx.Bool(true)
	}
}

func JsFnLogError(global typeutil.H) jsexecutor.JSFunction {
	red := color.FgRed.Render
	return func(ctx *jsexecutor.JSContext, this jsexecutor.JSValue, args []jsexecutor.JSValue) jsexecutor.JSValue {
		if len(args) > 0 {
			s, err := jsexecutor.JSValueToAny(args[0])
			if err != nil {
				return ctx.ThrowError(err)
			}
			format, ok := s.(string)
			a := make([]interface{}, 0)
			for _, v := range args[1:] {
				v2, err := jsexecutor.JSValueToAny(v)
				if err != nil {
					return ctx.ThrowError(err)
				}
				a = append(a, v2)
			}
			if ok {
				errLog.Printf(red(format), a...)
			} else {
				a = append([]interface{}{s}, a...)
				errLog.Println(red(a))
			}
		}
		return ctx.Bool(true)
	}
}
