package logger

type fieldType int

const (
	StringType fieldType = iota
	IntType
	FloatType
	BoolType
	ErrorType
)

type Field struct {
	Key   string
	Type  fieldType
	Str   string
	Int   int64
	Float float64
	Bool  bool
	Err   error
}

func String(key, val string) Field        { return Field{Key: key, Type: StringType, Str: val} }
func Int(key string, val int) Field       { return Field{Key: key, Type: IntType, Int: int64(val)} }
func Float(key string, val float64) Field { return Field{Key: key, Type: FloatType, Float: val} }
func Bool(key string, val bool) Field     { return Field{Key: key, Type: BoolType, Bool: val} }
func Error(key string, err error) Field   { return Field{Key: key, Type: ErrorType, Err: err} }

func (f Field) AppendValueTo(buf *fastBuffer) {
	switch f.Type {
	case StringType:
		buf.AppendString(f.Str)
	case IntType:
		buf.AppendInt(int(f.Int))
	case FloatType:
		buf.AppendFloat(f.Float)
	case BoolType:
		buf.AppendBool(f.Bool)
	case ErrorType:
		if f.Err != nil {
			buf.AppendString(f.Err.Error())
		} else {
			buf.AppendString("<nil>")
		}
	default:
		buf.AppendString("<unknown>")
	}
}
