// Auto-generated code. DO NOT EDIT!!!
// Generated by sts v0.0.1-alpha-dev.

package examples

func subsrc2anysub(src subsrc) anysub {
	return anysub{
		Test: src.Field1,
	}
}
func anysub2subsrc(src anysub) subsrc {
	return subsrc{
		Field1: src.Test,
	}
}
