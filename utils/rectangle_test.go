package utils

import "testing"

func TestRect_Intersects(t *testing.T) {
	type args struct {
		other *Rect
	}
	tests := []struct {
		name string
		rect *Rect
		args args
		want bool
	}{
		{
			name: "Null rects",
			rect: &Rect{0, 0, 0, 0},
			args: args{
				other: &Rect{0, 0, 0, 0},
			},
			want: false,
		},
		{
			name: "Non intersecting rects",
			rect: &Rect{0, 0, 1, 1},
			args: args{
				other: &Rect{2, 2, 1, 1},
			},
			want: false,
		},
		{
			name: "Intersecting rects",
			rect: &Rect{0, 0, 1, 1},
			args: args{
				other: &Rect{-1, -1, 2, 2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rect.Intersects(tt.args.other); got != tt.want {
				t.Errorf("Rect.Intersects() = %v, want %v", got, tt.want)
			}
		})
	}
}
