package validator

import "testing"

func BenchmarkBCP47Variants(b *testing.B) {
	v := New()
	for _, tag := range []string{"en-US", "sl-rozaj", "sl-rozaj-biske", "sl-rozaj-biske-rozaj"} {
		b.Run(tag, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				_ = v.Var(tag, "bcp47_strict_language_tag")
			}
		})
	}
}
