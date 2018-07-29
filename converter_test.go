package main

import (
	"testing"
)

func Test_convHeaders(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"01", args{"***test"}, "###test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convHeaders(tt.args.text); got != tt.want {
				t.Errorf("convHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convLists(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"01", args{"-test"}, "- test"},
		{"02", args{"--test"}, "  - test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convLists(tt.args.text); got != tt.want {
				t.Errorf("convLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convOrderLists(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"01", args{`123
+ 456
+ 789
abc`}, `123

1. 456
2. 789

abc`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convOrderedLists(tt.args.text); got != tt.want {
				t.Errorf("convOrderLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convCode(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"01", args{"\n 123\n 456\n"}, "\n```clike\n\t\t\t123\n\t\t\t456\n```\n"},
		{"02", args{"\n- 1\n 123\n 456\n"}, "\n- 1\n\n```clike\n\t\t\t123\n\t\t\t456\n```\n"},
		{"03", args{"\n 1\n2\n 3\n"}, "\n```clike\n\t\t\t1\n```\n\n2\n\n```clike\n\t\t\t3\n```\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convCode(tt.args.text); got != tt.want {
				t.Errorf("convCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convCode2(t *testing.T) {
	buf1 := `
- a
-- b
-- c
--- d
- e
`
	buf2 := `
- a
  - b
  - c
     - d
- e
`
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"01", args{"\n 1\n2\n 3\n"}, "\n```clike\n1\n```\n\n2\n\n```clike\n3\n```\n"},
		{"02", args{buf1 + "\n 1\n2\n 3\n"}, buf2 + "\n```clike\n1\n```\n\n2\n\n```clike\n3\n```\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convAll(tt.args.text); got != tt.want {
				t.Errorf("convAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
