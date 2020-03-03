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
 * cache/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/cache.py
 */

package cache

import (
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	Cache        λ.Object
	ϒexpand_path λ.Object
)

func init() {
	λ.InitModule(func() {
		ϒexpand_path = Ωutils.ϒexpand_path
		Cache = λ.Cal(λ.TypeType, λ.StrLiteral("Cache"), λ.NewTuple(λ.ObjectType), func() λ.Dict {
			var (
				Cache___init__ λ.Object
				Cache_enabled  λ.Object
				Cache_load     λ.Object
				Cache_store    λ.Object
			)
			Cache___init__ = λ.NewFunction("__init__",
				[]λ.Param{
					{Name: "self"},
					{Name: "ydl"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
						ϒydl  = λargs[1]
					)
					λ.SetAttr(ϒself, "_ydl", ϒydl)
					return λ.None
				})
			Cache_enabled = λ.NewFunction("enabled",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
					)
					return λ.NewBool(λ.Calm(λ.GetAttr(λ.GetAttr(ϒself, "_ydl", nil), "params", nil), "get", λ.StrLiteral("cachedir")) != λ.False)
				})
			Cache_enabled = λ.Cal(λ.PropertyType, Cache_enabled)
			Cache_store = λ.NewFunction("store",
				[]λ.Param{
					{Name: "self"},
					{Name: "section"},
					{Name: "key"},
					{Name: "data"},
					{Name: "dtype", Def: λ.StrLiteral("json")},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdata    = λargs[3]
						ϒdtype   = λargs[4]
						ϒfn      λ.Object
						ϒkey     = λargs[2]
						ϒsection = λargs[1]
						ϒself    = λargs[0]
						ϒtb      λ.Object
						τmp0     λ.Object
						τmp1     λ.Object
						τmp2     λ.Object
						τmp3     λ.Object
					)
					_ = τmp0
					_ = τmp1
					_ = τmp2
					_ = τmp3
					if !λ.Contains(λ.NewTuple(λ.StrLiteral("json")), ϒdtype) {
						panic(λ.Raise(λ.Cal(λ.AssertionErrorType)))
					}
					if !λ.IsTrue(λ.GetAttr(ϒself, "enabled", nil)) {
						return λ.None
					}
					ϒfn = λ.Calm(ϒself, "_get_cache_fn", ϒsection, ϒkey, ϒdtype)
					τmp0, τmp1 = func() (λexit λ.Object, λret λ.Object) {
						defer λ.CatchMulti(
							nil,
							&λ.Catcher{λ.ExceptionType, func(λex λ.BaseException) {
								ϒtb = λ.Calm(λ.None, "format_exc")
								λ.Calm(λ.GetAttr(ϒself, "_ydl", nil), "report_warning", λ.Mod(λ.StrLiteral("Writing cache to %r failed: %s"), λ.NewTuple(
									ϒfn,
									ϒtb,
								)))
							}},
						)
						τmp2, τmp3 = func() (λexit λ.Object, λret λ.Object) {
							defer λ.CatchMulti(
								nil,
								&λ.Catcher{λ.OSErrorType, func(λex λ.BaseException) {
									var ϒose λ.Object = λex
									if λ.IsTrue(λ.Ne(λ.GetAttr(ϒose, "errno", nil), λ.GetAttr(λ.None, "EEXIST", nil))) {
										panic(λ.Raise(λex))
									}
								}},
							)
							λ.Cal(λ.None, λ.Cal(λ.None, ϒfn))
							return λ.BlockExitNormally, nil
						}()
						λ.Cal(λ.None, ϒdata, ϒfn)
						return λ.BlockExitNormally, nil
					}()
					return λ.None
				})
			Cache_load = λ.NewFunction("load",
				[]λ.Param{
					{Name: "self"},
					{Name: "section"},
					{Name: "key"},
					{Name: "dtype", Def: λ.StrLiteral("json")},
					{Name: "default", Def: λ.None},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcache_fn  λ.Object
						ϒcachef    λ.Object
						ϒdefault   = λargs[4]
						ϒdtype     = λargs[3]
						ϒfile_size λ.Object
						ϒkey       = λargs[2]
						ϒsection   = λargs[1]
						ϒself      = λargs[0]
						τmp0       λ.Object
						τmp1       λ.Object
						τmp2       λ.Object
						τmp3       λ.Object
						τmp4       λ.Object
						τmp5       λ.Object
						τmp6       λ.Object
						τmp7       λ.Object
					)
					_ = τmp4
					_ = τmp5
					if !λ.Contains(λ.NewTuple(λ.StrLiteral("json")), ϒdtype) {
						panic(λ.Raise(λ.Cal(λ.AssertionErrorType)))
					}
					if !λ.IsTrue(λ.GetAttr(ϒself, "enabled", nil)) {
						return ϒdefault
					}
					ϒcache_fn = λ.Calm(ϒself, "_get_cache_fn", ϒsection, ϒkey, ϒdtype)
					τmp0, τmp1 = func() (λexit λ.Object, λret λ.Object) {
						defer λ.CatchMulti(
							nil,
							&λ.Catcher{λ.OSErrorType, func(λex λ.BaseException) {
								// pass
							}},
						)
						τmp2, τmp3 = func() (λexit λ.Object, λret λ.Object) {
							defer λ.CatchMulti(
								nil,
								&λ.Catcher{λ.ValueErrorType, func(λex λ.BaseException) {
									τmp5, τmp4 = func() (λexit λ.Object, λret λ.Object) {
										defer λ.CatchMulti(
											nil,
											&λ.Catcher{λ.NewTuple(
												λ.OSErrorType,
												λ.OSErrorType,
											), func(λex λ.BaseException) {
												var ϒoe λ.Object = λex
												ϒfile_size = λ.Cal(λ.StrType, ϒoe)
											}},
										)
										ϒfile_size = λ.Cal(λ.None, ϒcache_fn)
										return λ.BlockExitNormally, nil
									}()
									λ.Calm(λ.GetAttr(ϒself, "_ydl", nil), "report_warning", λ.Mod(λ.StrLiteral("Cache retrieval from %s failed (%s)"), λ.NewTuple(
										ϒcache_fn,
										ϒfile_size,
									)))
								}},
							)
							τmp4 = λ.Call(λ.None, λ.NewArgs(
								ϒcache_fn,
								λ.StrLiteral("r"),
							), λ.KWArgs{
								{Name: "encoding", Value: λ.StrLiteral("utf-8")},
							})
							τmp5 = λ.GetAttr(τmp4, "__exit__", nil)
							ϒcachef = λ.Calm(τmp4, "__enter__")
							τmp6, τmp7 = func() (λexit λ.Object, λret λ.Object) {
								defer λ.Catch(λ.BaseExceptionType, func(λex λ.BaseException) {
									if ret := λ.Cal(τmp5, λex.Type(), λex, λ.StrLiteral("..todo..traceback..")); λ.IsTrue(ret) {
										panic(λ.Raise(λex))
									}
								}, func() {
									λ.Cal(τmp5, λ.None, λ.None, λ.None)
								})
								λexit, λret = λ.BlockExitReturn, λ.Cal(λ.None, ϒcachef)
								return
								return λ.BlockExitNormally, nil
							}()
							if τmp6 == λ.BlockExitReturn {
								λexit, λret = λ.BlockExitReturn, τmp7
								return
							}
							return λ.BlockExitNormally, nil
						}()
						if τmp2 == λ.BlockExitReturn {
							λexit, λret = λ.BlockExitReturn, τmp3
							return
						}
						return λ.BlockExitNormally, nil
					}()
					if τmp0 == λ.BlockExitReturn {
						return τmp1
					}
					return ϒdefault
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"__init__": Cache___init__,
				"enabled":  Cache_enabled,
				"load":     Cache_load,
				"store":    Cache_store,
			})
		}())
	})
}
