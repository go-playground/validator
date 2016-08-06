package validator

import (
	"fmt"
	"reflect"
	"strconv"
)

const (
	arrayIndexFieldName = "%s" + leftBracket + "%d" + rightBracket
	mapIndexFieldName   = "%s" + leftBracket + "%v" + rightBracket
)

// per validate contruct
type validate struct {
	v              *Validate
	top            reflect.Value
	ns             []byte
	actualNs       []byte
	errs           ValidationErrors
	isPartial      bool
	hasExcludes    bool
	includeExclude map[string]struct{} // reset only if StructPartial or StructExcept are called, no need otherwise
	misc           []byte

	// StructLevel & FieldLevel fields
	slflParent reflect.Value
	slCurrent  reflect.Value
	slNs       []byte
	slStructNs []byte
	flField    reflect.Value
	flParam    string
}

// parent and current will be the same the first run of validateStruct
func (v *validate) validateStruct(parent reflect.Value, current reflect.Value, typ reflect.Type, ns []byte, structNs []byte, ct *cTag) {

	cs, ok := v.v.structCache.Get(typ)
	if !ok {
		cs = v.v.extractStructCache(current, typ.Name())
	}

	if len(ns) == 0 {

		ns = append(ns, cs.Name...)
		ns = append(ns, '.')

		structNs = append(structNs, cs.Name...)
		structNs = append(structNs, '.')
	}

	// ct is nil on top level struct, and structs as fields that have no tag info
	// so if nil or if not nil and the structonly tag isn't present
	if ct == nil || ct.typeof != typeStructOnly {

		for _, f := range cs.fields {

			if v.isPartial {

				_, ok = v.includeExclude[string(append(structNs, f.Name...))]

				if (ok && v.hasExcludes) || (!ok && !v.hasExcludes) {
					continue
				}
			}

			v.traverseField(parent, current.Field(f.Idx), ns, structNs, f, f.cTags)
		}
	}

	// check if any struct level validations, after all field validations already checked.
	// first iteration will have no info about nostructlevel tag, and is checked prior to
	// calling the next iteration of validateStruct called from traverseField.
	if cs.fn != nil {

		v.slflParent = parent
		v.slCurrent = current
		v.slNs = ns
		v.slStructNs = structNs

		cs.fn(v)
	}
}

// traverseField validates any field, be it a struct or single field, ensures it's validity and passes it along to be validated via it's tag options
func (v *validate) traverseField(parent reflect.Value, current reflect.Value, ns []byte, structNs []byte, cf *cField, ct *cTag) {

	var typ reflect.Type
	var kind reflect.Kind
	var nullable bool

	current, kind, nullable = v.extractTypeInternal(current, nullable)

	switch kind {
	case reflect.Ptr, reflect.Interface, reflect.Invalid:

		if ct == nil {
			return
		}

		if ct.typeof == typeOmitEmpty {
			return
		}

		if ct.hasTag {

			if kind == reflect.Invalid {

				v.errs = append(v.errs,
					&fieldError{
						tag:         ct.aliasTag,
						actualTag:   ct.tag,
						ns:          string(append(ns, cf.AltName...)),
						structNs:    string(append(structNs, cf.Name...)),
						field:       cf.AltName,
						structField: cf.Name,
						param:       ct.param,
						kind:        kind,
					},
				)

				return
			}

			v.errs = append(v.errs,
				&fieldError{
					tag:         ct.aliasTag,
					actualTag:   ct.tag,
					ns:          string(append(ns, cf.AltName...)),
					structNs:    string(append(structNs, cf.Name...)),
					field:       cf.AltName,
					structField: cf.Name,
					value:       current.Interface(),
					param:       ct.param,
					kind:        kind,
					typ:         current.Type(),
				},
			)

			return
		}

	case reflect.Struct:

		typ = current.Type()

		if typ != timeType {

			if ct != nil {
				ct = ct.next
			}

			if ct != nil && ct.typeof == typeNoStructLevel {
				return
			}

			// if len == 0 then validating using 'Var' or 'VarWithValue'
			// Var - doesn't make much sense to do it that way, should call 'Struct', but no harm...
			// VarWithField - this allows for validating against each field withing the struct against a specific value
			//                pretty handly in certain situations
			if len(ns) > 0 {
				ns = append(append(ns, cf.AltName...), '.')
				structNs = append(append(structNs, cf.Name...), '.')
			}

			v.validateStruct(current, current, typ, ns, structNs, ct)
			return
		}
	}

	if !ct.hasTag {
		return
	}

	typ = current.Type()

OUTER:
	for {
		if ct == nil {
			return
		}

		switch ct.typeof {

		case typeOmitEmpty:

			// set Field Level fields
			v.slflParent = parent
			v.flField = current
			v.flParam = ""

			if !nullable && !hasValue(v) {
				return
			}

			ct = ct.next
			continue

		case typeDive:

			ct = ct.next

			// traverse slice or map here
			// or panic ;)
			switch kind {
			case reflect.Slice, reflect.Array:

				var i64 int64
				reusableCF := &cField{}

				for i := 0; i < current.Len(); i++ {

					i64 = int64(i)

					v.misc = append(v.misc[0:0], cf.Name...)
					v.misc = append(v.misc, '[')
					v.misc = strconv.AppendInt(v.misc, i64, 10)
					v.misc = append(v.misc, ']')

					reusableCF.Name = string(v.misc)

					v.misc = append(v.misc[0:0], cf.AltName...)
					v.misc = append(v.misc, '[')
					v.misc = strconv.AppendInt(v.misc, i64, 10)
					v.misc = append(v.misc, ']')

					reusableCF.AltName = string(v.misc)

					v.traverseField(parent, current.Index(i), ns, structNs, reusableCF, ct)
				}

			case reflect.Map:

				var pv string
				reusableCF := &cField{}

				for _, key := range current.MapKeys() {

					pv = fmt.Sprintf("%v", key.Interface())

					v.misc = append(v.misc[0:0], cf.Name...)
					v.misc = append(v.misc, '[')
					v.misc = append(v.misc, pv...)
					v.misc = append(v.misc, ']')

					reusableCF.Name = string(v.misc)

					v.misc = append(v.misc[0:0], cf.AltName...)
					v.misc = append(v.misc, '[')
					v.misc = append(v.misc, pv...)
					v.misc = append(v.misc, ']')

					reusableCF.AltName = string(v.misc)

					v.traverseField(parent, current.MapIndex(key), ns, structNs, reusableCF, ct)
				}

			default:
				// throw error, if not a slice or map then should not have gotten here
				// bad dive tag
				panic("dive error! can't dive on a non slice or map")
			}

			return

		case typeOr:

			v.misc = v.misc[0:0]

			for {

				// set Field Level fields
				v.slflParent = parent
				v.flField = current
				v.flParam = ct.param

				if ct.fn(v) {

					// drain rest of the 'or' values, then continue or leave
					for {

						ct = ct.next

						if ct == nil {
							return
						}

						if ct.typeof != typeOr {
							continue OUTER
						}
					}
				}

				v.misc = append(v.misc, '|')
				v.misc = append(v.misc, ct.tag...)

				if ct.next == nil {
					// if we get here, no valid 'or' value and no more tags

					if ct.hasAlias {

						v.errs = append(v.errs,
							&fieldError{
								tag:         ct.aliasTag,
								actualTag:   ct.actualAliasTag,
								ns:          string(append(ns, cf.AltName...)),
								structNs:    string(append(structNs, cf.Name...)),
								field:       cf.AltName,
								structField: cf.Name,
								value:       current.Interface(),
								param:       ct.param,
								kind:        kind,
								typ:         typ,
							},
						)

					} else {

						v.errs = append(v.errs,
							&fieldError{
								tag:         string(v.misc)[1:],
								actualTag:   string(v.misc)[1:],
								ns:          string(append(ns, cf.AltName...)),
								structNs:    string(append(structNs, cf.Name...)),
								field:       cf.AltName,
								structField: cf.Name,
								value:       current.Interface(),
								param:       ct.param,
								kind:        kind,
								typ:         typ,
							},
						)
					}

					return
				}

				ct = ct.next
			}

		default:

			// set Field Level fields
			v.slflParent = parent
			v.flField = current
			v.flParam = ct.param

			if !ct.fn(v) {

				v.errs = append(v.errs,
					&fieldError{
						tag:         ct.aliasTag,
						actualTag:   ct.tag,
						ns:          string(append(ns, cf.AltName...)),
						structNs:    string(append(structNs, cf.Name...)),
						field:       cf.AltName,
						structField: cf.Name,
						value:       current.Interface(),
						param:       ct.param,
						kind:        kind,
						typ:         typ,
					},
				)

				return

			}

			ct = ct.next
		}
	}

}
