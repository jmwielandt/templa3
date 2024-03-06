package templa3

import (
	"fmt"

	"github.com/Masterminds/sprig/v3"
)

func (c *Template) registerFunctions() {
	c.customFuncs["printpass"] = c.printPass
	c.customFuncs["isdef"] = c.isDef
	// c.customFuncs["isdef_pass"] = c.isDefPass

	// gonna register all functions from sprig.
	for name, fn := range sprig.FuncMap() {
		/*
			full documentation: https://masterminds.github.io/sprig/

			some functions available:
			- trim, wrap, randAplha, plural
			- splitList, sortAlpha
			- add, max, mul
			- until, untilStep
			- addf, maxf, mulf
			- now, date
			- default, empty, coalesce, fromJson, toJson, toPrettyJson,
			  toRawJson, ternary
			- b64enc, b64dec
			- list, first, uniq
			- get, set, dict, hasKey, pluck, dig, deepCopy
			- atoy, int64, toString
			- base, dir, ext, clean, isAbs, osBase, osDir, osExt,
			  osClean, osIsAbs
			- fail
			- uuidv4
			- env, expandenv
			- semver, semverCompare
			- typeOf, kindIs, typeIsLike
			- derivePassword, sha256sum, genPrivateKey
			- getHostByName

		*/
		c.customFuncs[name] = fn
	}
}

func (c *Template) printPass(got ...any) []any {
	if c.verbose {
		fmt.Printf("printpass got: %+v\n", got)
	}
	return got
}

func (c *Template) isDef(item any, keys ...string) bool {
	for _, key := range keys {
		var element any
		switch item := item.(type) {
		case map[string]any:
			var ok bool
			element, ok = item[key]
			if !ok {
				return false
			}
		default:
			return false
		}
		item = element
	}
	return true
}

// func (c *Template) isDefPass(item any, keys ...string) (any, bool) {
// 	for _, key := range keys {
// 		var element any
// 		switch item := item.(type) {
// 		case map[string]any:
// 			var ok bool
// 			element, ok = item[key]
// 			if !ok {
// 				return nil, false
// 			}
// 		default:
// 			return nil, false
// 		}
// 		item = element
// 	}
// 	return item, true
// }
