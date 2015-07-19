package cocaine12

import (
	"errors"
	"fmt"
	"reflect"
)

var ErrNotAPointer = errors.New("result must be passed as a pointer to struct or slice")
var ErrBadType = errors.New("result must point to struct or slice")
var ErrNotEnoughFields = errors.New("not enough fields to unpack")
var ErrNotEnoughValues = errors.New("not enough values to unpack")

type ErrNotAssignable struct {
	Index            int
	SrcType, DstType reflect.Type
}

func NewErrNotAssignable(index int, dst, src reflect.Type) *ErrNotAssignable {
	return &ErrNotAssignable{
		Index:   index,
		SrcType: src,
		DstType: dst,
	}
}

func (e *ErrNotAssignable) Error() string {
	return fmt.Sprintf("(field %d) %v is not assignable to %v",
		e.Index, e.SrcType, e.DstType)
}

func unpackPayload(data []interface{}, res interface{}) error {
	resVal := reflect.ValueOf(res)
	switch resVal.Kind() {
	case reflect.Ptr:
		return unpack(data, resVal.Elem())
	default:
		return ErrNotAPointer
	}
}

func unpack(data []interface{}, res reflect.Value) error {
	switch res.Kind() {
	case reflect.Struct:
		return unpackToStruct(data, res)
	case reflect.Slice:
		return unpackToSlice(data, res)
	default:
		return ErrBadType
	}
}

func unpackToSlice(data []interface{}, res reflect.Value) error {
	return fmt.Errorf("not implemented")
}

func unpackToStruct(data []interface{}, res reflect.Value) error {
	fieldsNum := res.NumField()
	if len(data) != fieldsNum {
		switch {
		case len(data) < fieldsNum:
			return ErrNotEnoughValues
		default:
			return ErrNotEnoughFields
		}
	}

	for i := 0; i < fieldsNum; i++ {
		fieldVal := res.Field(i)
		dataVal := reflect.ValueOf(data[i])
		if dataVal.Type().AssignableTo(fieldVal.Type()) {
			fieldVal.Set(dataVal)
			continue
		}

		var err error

		switch res.Kind() {
		case reflect.Int:
			err = decodeInt(data[i], fieldVal)
		case reflect.Float32:
			err = decodeFloat(data[i], fieldVal)
		case reflect.Uint:
			err = decodeUint(data[i], fieldVal)
		default:
			return NewErrNotAssignable(i, dataVal.Type(), fieldVal.Type())
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func decodeInt(data interface{}, res reflect.Value) error {
	dataVal := reflect.ValueOf(data)

	switch dataVal.Kind() {
	case reflect.Int:
		res.SetInt(dataVal.Int())
	case reflect.Float32:
		res.SetInt(int64(dataVal.Float()))
	default:
		return fmt.Errorf("unable to convert %v to %v", dataVal.Type(), res.Type())
	}

	return nil
}

func decodeFloat(data interface{}, res reflect.Value) error {
	dataVal := reflect.ValueOf(data)

	return fmt.Errorf("unable to convert %v to %v", dataVal.Type(), res.Type())
}

func decodeUint(data interface{}, res reflect.Value) error {
	dataVal := reflect.ValueOf(data)

	return fmt.Errorf("unable to convert %v to %v", dataVal.Type(), res.Type())
}
