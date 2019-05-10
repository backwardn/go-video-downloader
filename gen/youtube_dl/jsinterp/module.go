// Code generated by transpiler. DO NOT EDIT.

/**
 * Go Video Downloader
 *
 *    Copyright 2019 Tenta, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * For any questions, please contact developer@tenta.io
 *
 * jsinterp/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/jsinterp.py
 */

package jsinterp

import (
	Ωjson "github.com/tenta-browser/go-video-downloader/gen/json"
	Ωoperator "github.com/tenta-browser/go-video-downloader/gen/operator"
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError     λ.Object
	JSInterpreter      λ.Object
	ϒ_ASSIGN_OPERATORS λ.Object
	ϒ_NAME_RE          λ.Object
	ϒ_OPERATORS        λ.Object
	ϒremove_quotes     λ.Object
)

func init() {
	λ.InitModule(func() {
		ExtractorError = Ωutils.ExtractorError
		ϒremove_quotes = Ωutils.ϒremove_quotes
		ϒ_OPERATORS = λ.NewList(
			λ.NewTuple(
				λ.NewStr("|"),
				Ωoperator.ϒor_,
			),
			λ.NewTuple(
				λ.NewStr("^"),
				Ωoperator.ϒxor,
			),
			λ.NewTuple(
				λ.NewStr("&"),
				Ωoperator.ϒand_,
			),
			λ.NewTuple(
				λ.NewStr(">>"),
				Ωoperator.ϒrshift,
			),
			λ.NewTuple(
				λ.NewStr("<<"),
				Ωoperator.ϒlshift,
			),
			λ.NewTuple(
				λ.NewStr("-"),
				Ωoperator.ϒsub,
			),
			λ.NewTuple(
				λ.NewStr("+"),
				Ωoperator.ϒadd,
			),
			λ.NewTuple(
				λ.NewStr("%"),
				Ωoperator.ϒmod,
			),
			λ.NewTuple(
				λ.NewStr("/"),
				Ωoperator.ϒtruediv,
			),
			λ.NewTuple(
				λ.NewStr("*"),
				Ωoperator.ϒmul,
			),
		)
		ϒ_ASSIGN_OPERATORS = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
			nil,
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
					var (
						ϒop     λ.Object
						ϒopfunc λ.Object
						τmp0    λ.Object
						τmp1    λ.Object
						τmp2    λ.Object
					)
					τmp0 = λ.Cal(λ.BuiltinIter, ϒ_OPERATORS)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒop = λ.GetItem(τmp2, λ.NewInt(0))
						ϒopfunc = λ.GetItem(τmp2, λ.NewInt(1))
						λgy.Yield(λ.NewTuple(
							λ.Add(ϒop, λ.NewStr("=")),
							ϒopfunc,
						))
					}
					return λ.None
				})
			})))
		λ.Cal(λ.GetAttr(ϒ_ASSIGN_OPERATORS, "append", nil), λ.NewTuple(
			λ.NewStr("="),
			λ.NewFunction("<lambda>",
				[]λ.Param{
					{Name: "cur"},
					{Name: "right"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcur   = λargs[0]
						ϒright = λargs[1]
					)
					_ = ϒcur
					return ϒright
				}),
		))
		ϒ_NAME_RE = λ.NewStr("[a-zA-Z_$][a-zA-Z_$0-9]*")
		JSInterpreter = λ.Cal(λ.TypeType, λ.NewStr("JSInterpreter"), λ.NewTuple(λ.ObjectType), func() λ.Dict {
			var (
				JSInterpreter___init__             λ.Object
				JSInterpreter_build_function       λ.Object
				JSInterpreter_extract_function     λ.Object
				JSInterpreter_extract_object       λ.Object
				JSInterpreter_interpret_expression λ.Object
				JSInterpreter_interpret_statement  λ.Object
			)
			JSInterpreter___init__ = λ.NewFunction("__init__",
				[]λ.Param{
					{Name: "self"},
					{Name: "code"},
					{Name: "objects", Def: λ.None},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcode    = λargs[1]
						ϒobjects = λargs[2]
						ϒself    = λargs[0]
					)
					if λ.IsTrue(λ.NewBool(ϒobjects == λ.None)) {
						ϒobjects = λ.NewDictWithTable(map[λ.Object]λ.Object{})
					}
					λ.SetAttr(ϒself, "code", ϒcode)
					λ.SetAttr(ϒself, "_functions", λ.NewDictWithTable(map[λ.Object]λ.Object{}))
					λ.SetAttr(ϒself, "_objects", ϒobjects)
					return λ.None
				})
			JSInterpreter_interpret_statement = λ.NewFunction("interpret_statement",
				[]λ.Param{
					{Name: "self"},
					{Name: "stmt"},
					{Name: "local_vars"},
					{Name: "allow_recursion", Def: λ.NewInt(100)},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒallow_recursion = λargs[3]
						ϒexpr            λ.Object
						ϒlocal_vars      = λargs[2]
						ϒreturn_m        λ.Object
						ϒself            = λargs[0]
						ϒshould_abort    λ.Object
						ϒstmt            = λargs[1]
						ϒstmt_m          λ.Object
						ϒv               λ.Object
					)
					if λ.IsTrue(λ.Lt(ϒallow_recursion, λ.NewInt(0))) {
						panic(λ.Raise(λ.Cal(ExtractorError, λ.NewStr("Recursion limit reached"))))
					}
					ϒshould_abort = λ.False
					ϒstmt = λ.Cal(λ.GetAttr(ϒstmt, "lstrip", nil))
					ϒstmt_m = λ.Cal(Ωre.ϒmatch, λ.NewStr("var\\s"), ϒstmt)
					if λ.IsTrue(ϒstmt_m) {
						ϒexpr = λ.GetItem(ϒstmt, λ.NewSlice(λ.Cal(λ.BuiltinLen, λ.Cal(λ.GetAttr(ϒstmt_m, "group", nil), λ.NewInt(0))), λ.None, λ.None))
					} else {
						ϒreturn_m = λ.Cal(Ωre.ϒmatch, λ.NewStr("return(?:\\s+|$)"), ϒstmt)
						if λ.IsTrue(ϒreturn_m) {
							ϒexpr = λ.GetItem(ϒstmt, λ.NewSlice(λ.Cal(λ.BuiltinLen, λ.Cal(λ.GetAttr(ϒreturn_m, "group", nil), λ.NewInt(0))), λ.None, λ.None))
							ϒshould_abort = λ.True
						} else {
							ϒexpr = ϒstmt
						}
					}
					ϒv = λ.Cal(λ.GetAttr(ϒself, "interpret_expression", nil), ϒexpr, ϒlocal_vars, ϒallow_recursion)
					return λ.NewTuple(
						ϒv,
						ϒshould_abort,
					)
				})
			JSInterpreter_interpret_expression = λ.NewFunction("interpret_expression",
				[]λ.Param{
					{Name: "self"},
					{Name: "expr"},
					{Name: "local_vars"},
					{Name: "allow_recursion"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒabort           λ.Object
						ϒallow_recursion = λargs[3]
						ϒarg_str         λ.Object
						ϒargvals         λ.Object
						ϒcur             λ.Object
						ϒexpr            = λargs[1]
						ϒfname           λ.Object
						ϒhowMany         λ.Object
						ϒi               λ.Object
						ϒidx             λ.Object
						ϒindex           λ.Object
						ϒlocal_vars      = λargs[2]
						ϒlvar            λ.Object
						ϒm               λ.Object
						ϒmember          λ.Object
						ϒobj             λ.Object
						ϒop              λ.Object
						ϒopfunc          λ.Object
						ϒparens_count    λ.Object
						ϒremaining_expr  λ.Object
						ϒres             λ.Object
						ϒright_val       λ.Object
						ϒself            = λargs[0]
						ϒsub_expr        λ.Object
						ϒsub_result      λ.Object
						ϒval             λ.Object
						ϒvar_m           λ.Object
						ϒvariable        λ.Object
						ϒx               λ.Object
						ϒy               λ.Object
						τmp0             λ.Object
						τmp1             λ.Object
						τmp2             λ.Object
					)
					_ = ϒi
					ϒexpr = λ.Cal(λ.GetAttr(ϒexpr, "strip", nil))
					if λ.IsTrue(λ.Eq(ϒexpr, λ.NewStr(""))) {
						return λ.None
					}
					if λ.IsTrue(λ.Cal(λ.GetAttr(ϒexpr, "startswith", nil), λ.NewStr("("))) {
						ϒparens_count = λ.NewInt(0)
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfinditer, λ.NewStr("[()]"), ϒexpr))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒm = τmp1
							if λ.IsTrue(λ.Eq(λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewInt(0)), λ.NewStr("("))) {
								τmp2 = λ.IAdd(ϒparens_count, λ.NewInt(1))
								ϒparens_count = τmp2
							} else {
								τmp2 = λ.ISub(ϒparens_count, λ.NewInt(1))
								ϒparens_count = τmp2
								if λ.IsTrue(λ.Eq(ϒparens_count, λ.NewInt(0))) {
									ϒsub_expr = λ.GetItem(ϒexpr, λ.NewSlice(λ.NewInt(1), λ.Cal(λ.GetAttr(ϒm, "start", nil)), λ.None))
									ϒsub_result = λ.Cal(λ.GetAttr(ϒself, "interpret_expression", nil), ϒsub_expr, ϒlocal_vars, ϒallow_recursion)
									ϒremaining_expr = λ.Cal(λ.GetAttr(λ.GetItem(ϒexpr, λ.NewSlice(λ.Cal(λ.GetAttr(ϒm, "end", nil)), λ.None, λ.None)), "strip", nil))
									if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒremaining_expr))) {
										return ϒsub_result
									} else {
										ϒexpr = λ.Add(λ.Cal(Ωjson.ϒdumps, ϒsub_result), ϒremaining_expr)
									}
									break
								}
							}
						}
						if τmp1 == λ.AfterLast {
							panic(λ.Raise(λ.Cal(ExtractorError, λ.Mod(λ.NewStr("Premature end of parens in %r"), ϒexpr))))
						}
					}
					τmp0 = λ.Cal(λ.BuiltinIter, ϒ_ASSIGN_OPERATORS)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒop = λ.GetItem(τmp2, λ.NewInt(0))
						ϒopfunc = λ.GetItem(τmp2, λ.NewInt(1))
						ϒm = λ.Cal(Ωre.ϒmatch, λ.Mod(λ.NewStr("(?x)\n                (?P<out>%s)(?:\\[(?P<index>[^\\]]+?)\\])?\n                \\s*%s\n                (?P<expr>.*)$"), λ.NewTuple(
							ϒ_NAME_RE,
							λ.Cal(Ωre.ϒescape, ϒop),
						)), ϒexpr)
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒm))) {
							continue
						}
						ϒright_val = λ.Cal(λ.GetAttr(ϒself, "interpret_expression", nil), λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("expr")), ϒlocal_vars, λ.Sub(ϒallow_recursion, λ.NewInt(1)))
						if λ.IsTrue(λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒm, "groupdict", nil)), "get", nil), λ.NewStr("index"))) {
							ϒlvar = λ.GetItem(ϒlocal_vars, λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("out")))
							ϒidx = λ.Cal(λ.GetAttr(ϒself, "interpret_expression", nil), λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("index")), ϒlocal_vars, ϒallow_recursion)
							if !λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒidx, λ.IntType)) {
								panic(λ.Raise(λ.Cal(λ.AssertionErrorType)))
							}
							ϒcur = λ.GetItem(ϒlvar, ϒidx)
							ϒval = λ.Cal(ϒopfunc, ϒcur, ϒright_val)
							λ.SetItem(ϒlvar, ϒidx, ϒval)
							return ϒval
						} else {
							ϒcur = λ.Cal(λ.GetAttr(ϒlocal_vars, "get", nil), λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("out")))
							ϒval = λ.Cal(ϒopfunc, ϒcur, ϒright_val)
							λ.SetItem(ϒlocal_vars, λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("out")), ϒval)
							return ϒval
						}
					}
					if λ.IsTrue(λ.Cal(λ.GetAttr(ϒexpr, "isdigit", nil))) {
						return λ.Cal(λ.IntType, ϒexpr)
					}
					ϒvar_m = λ.Cal(Ωre.ϒmatch, λ.Mod(λ.NewStr("(?!if|return|true|false)(?P<name>%s)$"), ϒ_NAME_RE), ϒexpr)
					if λ.IsTrue(ϒvar_m) {
						return λ.GetItem(ϒlocal_vars, λ.Cal(λ.GetAttr(ϒvar_m, "group", nil), λ.NewStr("name")))
					}
					τmp0, τmp1 = func() (λexit λ.Object, λret λ.Object) {
						defer λ.CatchMulti(
							nil,
							&λ.Catcher{λ.ValueErrorType, func(λex λ.BaseException) {
								// pass
							}},
						)
						λexit, λret = λ.BlockExitReturn, λ.Cal(Ωjson.ϒloads, ϒexpr)
						return
						return λ.BlockExitNormally, nil
					}()
					if τmp0 == λ.BlockExitReturn {
						return τmp1
					}
					ϒm = λ.Cal(Ωre.ϒmatch, λ.Mod(λ.NewStr("(?P<in>%s)\\[(?P<idx>.+)\\]$"), ϒ_NAME_RE), ϒexpr)
					if λ.IsTrue(ϒm) {
						ϒval = λ.GetItem(ϒlocal_vars, λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("in")))
						ϒidx = λ.Cal(λ.GetAttr(ϒself, "interpret_expression", nil), λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("idx")), ϒlocal_vars, λ.Sub(ϒallow_recursion, λ.NewInt(1)))
						return λ.GetItem(ϒval, ϒidx)
					}
					ϒm = λ.Cal(Ωre.ϒmatch, λ.Mod(λ.NewStr("(?P<var>%s)(?:\\.(?P<member>[^(]+)|\\[(?P<member2>[^]]+)\\])\\s*(?:\\(+(?P<args>[^()]*)\\))?$"), ϒ_NAME_RE), ϒexpr)
					if λ.IsTrue(ϒm) {
						ϒvariable = λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("var"))
						ϒmember = λ.Cal(ϒremove_quotes, func() λ.Object {
							if λv := λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("member")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("member2"))
							}
						}())
						ϒarg_str = λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("args"))
						if λ.IsTrue(λ.NewBool(λ.Contains(ϒlocal_vars, ϒvariable))) {
							ϒobj = λ.GetItem(ϒlocal_vars, ϒvariable)
						} else {
							if λ.IsTrue(λ.NewBool(!λ.Contains(λ.GetAttr(ϒself, "_objects", nil), ϒvariable))) {
								λ.SetItem(λ.GetAttr(ϒself, "_objects", nil), ϒvariable, λ.Cal(λ.GetAttr(ϒself, "extract_object", nil), ϒvariable))
							}
							ϒobj = λ.GetItem(λ.GetAttr(ϒself, "_objects", nil), ϒvariable)
						}
						if λ.IsTrue(λ.NewBool(ϒarg_str == λ.None)) {
							if λ.IsTrue(λ.Eq(ϒmember, λ.NewStr("length"))) {
								return λ.Cal(λ.BuiltinLen, ϒobj)
							}
							return λ.GetItem(ϒobj, ϒmember)
						}
						if !λ.IsTrue(λ.Cal(λ.GetAttr(ϒexpr, "endswith", nil), λ.NewStr(")"))) {
							panic(λ.Raise(λ.Cal(λ.AssertionErrorType)))
						}
						if λ.IsTrue(λ.Eq(ϒarg_str, λ.NewStr(""))) {
							ϒargvals = λ.Cal(λ.TupleType)
						} else {
							ϒargvals = λ.Cal(λ.TupleType, λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
								nil,
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
										var (
											ϒv   λ.Object
											τmp0 λ.Object
											τmp1 λ.Object
										)
										τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒarg_str, "split", nil), λ.NewStr(",")))
										for {
											if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
												break
											}
											ϒv = τmp1
											λgy.Yield(λ.Cal(λ.GetAttr(ϒself, "interpret_expression", nil), ϒv, ϒlocal_vars, ϒallow_recursion))
										}
										return λ.None
									})
								}))))
						}
						if λ.IsTrue(λ.Eq(ϒmember, λ.NewStr("split"))) {
							if !λ.IsTrue(λ.Eq(ϒargvals, λ.NewTuple(λ.NewStr("")))) {
								panic(λ.Raise(λ.Cal(λ.AssertionErrorType)))
							}
							return λ.Cal(λ.ListType, ϒobj)
						}
						if λ.IsTrue(λ.Eq(ϒmember, λ.NewStr("join"))) {
							if !λ.IsTrue(λ.Eq(λ.Cal(λ.BuiltinLen, ϒargvals), λ.NewInt(1))) {
								panic(λ.Raise(λ.Cal(λ.AssertionErrorType)))
							}
							return λ.Cal(λ.GetAttr(λ.GetItem(ϒargvals, λ.NewInt(0)), "join", nil), ϒobj)
						}
						if λ.IsTrue(λ.Eq(ϒmember, λ.NewStr("reverse"))) {
							if !λ.IsTrue(λ.Eq(λ.Cal(λ.BuiltinLen, ϒargvals), λ.NewInt(0))) {
								panic(λ.Raise(λ.Cal(λ.AssertionErrorType)))
							}
							λ.Cal(λ.GetAttr(ϒobj, "reverse", nil))
							return ϒobj
						}
						if λ.IsTrue(λ.Eq(ϒmember, λ.NewStr("slice"))) {
							if !λ.IsTrue(λ.Eq(λ.Cal(λ.BuiltinLen, ϒargvals), λ.NewInt(1))) {
								panic(λ.Raise(λ.Cal(λ.AssertionErrorType)))
							}
							return λ.GetItem(ϒobj, λ.NewSlice(λ.GetItem(ϒargvals, λ.NewInt(0)), λ.None, λ.None))
						}
						if λ.IsTrue(λ.Eq(ϒmember, λ.NewStr("splice"))) {
							if !λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒobj, λ.ListType)) {
								panic(λ.Raise(λ.Cal(λ.AssertionErrorType)))
							}
							τmp1 = ϒargvals
							ϒindex = λ.GetItem(τmp1, λ.NewInt(0))
							ϒhowMany = λ.GetItem(τmp1, λ.NewInt(1))
							ϒres = λ.NewList()
							τmp1 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.RangeType, ϒindex, λ.Cal(λ.BuiltinMin, λ.Add(ϒindex, ϒhowMany), λ.Cal(λ.BuiltinLen, ϒobj))))
							for {
								if τmp0 = λ.NextDefault(τmp1, λ.AfterLast); τmp0 == λ.AfterLast {
									break
								}
								ϒi = τmp0
								λ.Cal(λ.GetAttr(ϒres, "append", nil), λ.Cal(λ.GetAttr(ϒobj, "pop", nil), ϒindex))
							}
							return ϒres
						}
						return λ.Cal(λ.GetItem(ϒobj, ϒmember), ϒargvals)
					}
					τmp1 = λ.Cal(λ.BuiltinIter, ϒ_OPERATORS)
					for {
						if τmp0 = λ.NextDefault(τmp1, λ.AfterLast); τmp0 == λ.AfterLast {
							break
						}
						τmp2 = τmp0
						ϒop = λ.GetItem(τmp2, λ.NewInt(0))
						ϒopfunc = λ.GetItem(τmp2, λ.NewInt(1))
						ϒm = λ.Cal(Ωre.ϒmatch, λ.Mod(λ.NewStr("(?P<x>.+?)%s(?P<y>.+)"), λ.Cal(Ωre.ϒescape, ϒop)), ϒexpr)
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒm))) {
							continue
						}
						τmp2 = λ.Cal(λ.GetAttr(ϒself, "interpret_statement", nil), λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("x")), ϒlocal_vars, λ.Sub(ϒallow_recursion, λ.NewInt(1)))
						ϒx = λ.GetItem(τmp2, λ.NewInt(0))
						ϒabort = λ.GetItem(τmp2, λ.NewInt(1))
						if λ.IsTrue(ϒabort) {
							panic(λ.Raise(λ.Cal(ExtractorError, λ.Mod(λ.NewStr("Premature left-side return of %s in %r"), λ.NewTuple(
								ϒop,
								ϒexpr,
							)))))
						}
						τmp2 = λ.Cal(λ.GetAttr(ϒself, "interpret_statement", nil), λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("y")), ϒlocal_vars, λ.Sub(ϒallow_recursion, λ.NewInt(1)))
						ϒy = λ.GetItem(τmp2, λ.NewInt(0))
						ϒabort = λ.GetItem(τmp2, λ.NewInt(1))
						if λ.IsTrue(ϒabort) {
							panic(λ.Raise(λ.Cal(ExtractorError, λ.Mod(λ.NewStr("Premature right-side return of %s in %r"), λ.NewTuple(
								ϒop,
								ϒexpr,
							)))))
						}
						return λ.Cal(ϒopfunc, ϒx, ϒy)
					}
					ϒm = λ.Cal(Ωre.ϒmatch, λ.Mod(λ.NewStr("^(?P<func>%s)\\((?P<args>[a-zA-Z0-9_$,]*)\\)$"), ϒ_NAME_RE), ϒexpr)
					if λ.IsTrue(ϒm) {
						ϒfname = λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("func"))
						ϒargvals = func() λ.Object {
							if λ.IsTrue(λ.Gt(λ.Cal(λ.BuiltinLen, λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("args"))), λ.NewInt(0))) {
								return λ.Cal(λ.TupleType, λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
									nil,
									0, false, false,
									func(λargs []λ.Object) λ.Object {
										return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
											var (
												ϒv   λ.Object
												τmp0 λ.Object
												τmp1 λ.Object
											)
											τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("args")), "split", nil), λ.NewStr(",")))
											for {
												if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
													break
												}
												ϒv = τmp1
												λgy.Yield(func() λ.Object {
													if λ.IsTrue(λ.Cal(λ.GetAttr(ϒv, "isdigit", nil))) {
														return λ.Cal(λ.IntType, ϒv)
													} else {
														return λ.GetItem(ϒlocal_vars, ϒv)
													}
												}())
											}
											return λ.None
										})
									}))))
							} else {
								return λ.Cal(λ.TupleType)
							}
						}()
						if λ.IsTrue(λ.NewBool(!λ.Contains(λ.GetAttr(ϒself, "_functions", nil), ϒfname))) {
							λ.SetItem(λ.GetAttr(ϒself, "_functions", nil), ϒfname, λ.Cal(λ.GetAttr(ϒself, "extract_function", nil), ϒfname))
						}
						return λ.Cal(λ.GetItem(λ.GetAttr(ϒself, "_functions", nil), ϒfname), ϒargvals)
					}
					panic(λ.Raise(λ.Cal(ExtractorError, λ.Mod(λ.NewStr("Unsupported JS expression %r"), ϒexpr))))
					return λ.None
				})
			JSInterpreter_extract_object = λ.NewFunction("extract_object",
				[]λ.Param{
					{Name: "self"},
					{Name: "objname"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒ_FUNC_NAME_RE λ.Object
						ϒargnames      λ.Object
						ϒf             λ.Object
						ϒfields        λ.Object
						ϒfields_m      λ.Object
						ϒobj           λ.Object
						ϒobj_m         λ.Object
						ϒobjname       = λargs[1]
						ϒself          = λargs[0]
						τmp0           λ.Object
						τmp1           λ.Object
					)
					ϒ_FUNC_NAME_RE = λ.NewStr("(?:[a-zA-Z$0-9]+|\"[a-zA-Z$0-9]+\"|'[a-zA-Z$0-9]+')")
					ϒobj = λ.NewDictWithTable(map[λ.Object]λ.Object{})
					ϒobj_m = λ.Cal(Ωre.ϒsearch, λ.Mod(λ.NewStr("(?x)\n                (?<!this\\.)%s\\s*=\\s*{\\s*\n                    (?P<fields>(%s\\s*:\\s*function\\s*\\(.*?\\)\\s*{.*?}(?:,\\s*)?)*)\n                }\\s*;\n            "), λ.NewTuple(
						λ.Cal(Ωre.ϒescape, ϒobjname),
						ϒ_FUNC_NAME_RE,
					)), λ.GetAttr(ϒself, "code", nil))
					ϒfields = λ.Cal(λ.GetAttr(ϒobj_m, "group", nil), λ.NewStr("fields"))
					ϒfields_m = λ.Cal(Ωre.ϒfinditer, λ.Mod(λ.NewStr("(?x)\n                (?P<key>%s)\\s*:\\s*function\\s*\\((?P<args>[a-z,]+)\\){(?P<code>[^}]+)}\n            "), ϒ_FUNC_NAME_RE), ϒfields)
					τmp0 = λ.Cal(λ.BuiltinIter, ϒfields_m)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒf = τmp1
						ϒargnames = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒf, "group", nil), λ.NewStr("args")), "split", nil), λ.NewStr(","))
						λ.SetItem(ϒobj, λ.Cal(ϒremove_quotes, λ.Cal(λ.GetAttr(ϒf, "group", nil), λ.NewStr("key"))), λ.Cal(λ.GetAttr(ϒself, "build_function", nil), ϒargnames, λ.Cal(λ.GetAttr(ϒf, "group", nil), λ.NewStr("code"))))
					}
					return ϒobj
				})
			JSInterpreter_extract_function = λ.NewFunction("extract_function",
				[]λ.Param{
					{Name: "self"},
					{Name: "funcname"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒargnames λ.Object
						ϒfunc_m   λ.Object
						ϒfuncname = λargs[1]
						ϒself     = λargs[0]
					)
					ϒfunc_m = λ.Cal(Ωre.ϒsearch, λ.Mod(λ.NewStr("(?x)\n                (?:function\\s+%s|[{;,]\\s*%s\\s*=\\s*function|var\\s+%s\\s*=\\s*function)\\s*\n                \\((?P<args>[^)]*)\\)\\s*\n                \\{(?P<code>[^}]+)\\}"), λ.NewTuple(
						λ.Cal(Ωre.ϒescape, ϒfuncname),
						λ.Cal(Ωre.ϒescape, ϒfuncname),
						λ.Cal(Ωre.ϒescape, ϒfuncname),
					)), λ.GetAttr(ϒself, "code", nil))
					if λ.IsTrue(λ.NewBool(ϒfunc_m == λ.None)) {
						panic(λ.Raise(λ.Cal(ExtractorError, λ.Mod(λ.NewStr("Could not find JS function %r"), ϒfuncname))))
					}
					ϒargnames = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒfunc_m, "group", nil), λ.NewStr("args")), "split", nil), λ.NewStr(","))
					return λ.Cal(λ.GetAttr(ϒself, "build_function", nil), ϒargnames, λ.Cal(λ.GetAttr(ϒfunc_m, "group", nil), λ.NewStr("code")))
				})
			JSInterpreter_build_function = λ.NewFunction("build_function",
				[]λ.Param{
					{Name: "self"},
					{Name: "argnames"},
					{Name: "code"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒargnames = λargs[1]
						ϒcode     = λargs[2]
						ϒresf     λ.Object
						ϒself     = λargs[0]
					)
					ϒresf = λ.NewFunction("resf",
						[]λ.Param{
							{Name: "args"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒabort      λ.Object
								ϒargs       = λargs[0]
								ϒlocal_vars λ.Object
								ϒres        λ.Object
								ϒstmt       λ.Object
								τmp0        λ.Object
								τmp1        λ.Object
								τmp2        λ.Object
							)
							ϒlocal_vars = λ.Cal(λ.DictType, λ.Cal(λ.ZipIteratorType, ϒargnames, ϒargs))
							τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒcode, "split", nil), λ.NewStr(";")))
							for {
								if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
									break
								}
								ϒstmt = τmp1
								τmp2 = λ.Cal(λ.GetAttr(ϒself, "interpret_statement", nil), ϒstmt, ϒlocal_vars)
								ϒres = λ.GetItem(τmp2, λ.NewInt(0))
								ϒabort = λ.GetItem(τmp2, λ.NewInt(1))
								if λ.IsTrue(ϒabort) {
									break
								}
							}
							return ϒres
						})
					return ϒresf
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("__init__"):             JSInterpreter___init__,
				λ.NewStr("build_function"):       JSInterpreter_build_function,
				λ.NewStr("extract_function"):     JSInterpreter_extract_function,
				λ.NewStr("extract_object"):       JSInterpreter_extract_object,
				λ.NewStr("interpret_expression"): JSInterpreter_interpret_expression,
				λ.NewStr("interpret_statement"):  JSInterpreter_interpret_statement,
			})
		}())
	})
}
