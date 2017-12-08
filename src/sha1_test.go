package main

import "testing"

func BenchmarkLuhnSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Digest([]byte("123"))
	}
}

func BenchmarkLuhnLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Digest([]byte("00123014764700968325001230147647009683250012301476470096832500123014764700968325"))
	}
}

func BenchmarkLuhnSmallParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Digest([]byte("123"))
		}
	})
}

func BenchmarkLuhnLargeParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Digest([]byte("00123014764700968325001230147647009683250012301476470096832500123014764700968325"))
		}
	})
}

// FIXME result sha1 of "a"
//86F7E437FAA5A7FCE15D1DDCB9EAEAEA377667B8
