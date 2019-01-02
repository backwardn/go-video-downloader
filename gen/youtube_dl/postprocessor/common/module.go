// Code generated by transpiler. DO NOT EDIT.

/**
 * Go Video Downloader
 *
 *    Copyright 2018 Tenta, LLC
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
 * common/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/postprocessor/common.py
 */

package common

import (
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	PostProcessingError λ.Object
	PostProcessor       λ.Object
)

func init() {
	λ.InitModule(func() {
		PostProcessingError = Ωutils.PostProcessingError
		PostProcessor = λ.Cal(λ.TypeType, λ.NewStr("PostProcessor"), λ.NewTuple(λ.ObjectType), func() λ.Dict {
			λ.NewStr("Post Processor class.\n\n    PostProcessor objects can be added to downloaders with their\n    add_post_processor() method. When the downloader has finished a\n    successful download, it will take its internal chain of PostProcessors\n    and start calling the run() method on each one of them, first with\n    an initial argument and then with the returned value of the previous\n    PostProcessor.\n\n    The chain will be stopped if one of them ever returns None or the end\n    of the chain is reached.\n\n    PostProcessor objects follow a \"mutual registration\" process similar\n    to InfoExtractor objects.\n\n    Optionally PostProcessor can use a list of additional command-line arguments\n    with self._configuration_args.\n    ")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{})
		}())
	})
}