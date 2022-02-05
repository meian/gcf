package gcf_test

import (
	"fmt"
	"testing"

	"github.com/meian/gcf"
	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	type args struct {
		start int
		end   int
		step  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    []int
	}{
		{
			name: "positive step",
			args: args{
				start: 1,
				end:   10,
				step:  3,
			},
			want: []int{1, 4, 7, 10},
		},
		{
			name: "positive step with end is not match steps",
			args: args{
				start: 1,
				end:   10,
				step:  2,
			},
			want: []int{1, 3, 5, 7, 9},
		},
		{
			name: "positive big step",
			args: args{
				start: 1,
				end:   10,
				step:  10,
			},
			want: []int{1},
		},
		{
			name: "positive step but start greater than end",
			args: args{
				start: 1,
				end:   0,
				step:  2,
			},
			want: []int{},
		},
		{
			name: "positive step with start equals to end",
			args: args{
				start: 1,
				end:   1,
				step:  2,
			},
			want: []int{1},
		},
		{
			name: "negative step",
			args: args{
				start: 8,
				end:   -1,
				step:  -3,
			},
			want: []int{8, 5, 2, -1},
		},
		{
			name: "negative step with end is not match steps",
			args: args{
				start: 8,
				end:   -1,
				step:  -2,
			},
			want: []int{8, 6, 4, 2, 0},
		},
		{
			name: "negative big step",
			args: args{
				start: 8,
				end:   -1,
				step:  -10,
			},
			want: []int{8},
		},
		{
			name: "negative step but start less than end",
			args: args{
				start: 8,
				end:   9,
				step:  -2,
			},
			want: []int{},
		},
		{
			name: "negative step with start equals to end",
			args: args{
				start: 8,
				end:   8,
				step:  -2,
			},
			want: []int{8},
		},
		{
			name: "step is zero",
			args: args{
				start: 1,
				end:   2,
				step:  0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			itb, err := gcf.Range(tt.args.start, tt.args.end, tt.args.step)
			if tt.wantErr {
				assert.Error(err)
				assert.Empty(itb)
				return
			}
			assert.NoError(err)
			s := gcf.ToSlice(itb)
			assert.Equal(tt.want, s)
		})
	}

	itb, _ := gcf.Range(1, 3, 1)
	testBeforeAndAfter(t, itb)
	itb, _ = gcf.Range(3, 1, -1)
	testBeforeAndAfter(t, itb)

	testEmpties(t, func(itb gcf.Iterable[int]) gcf.Iterable[int] {
		itb, _ = gcf.Range(1, 0, 1)
		return itb
	})
	testEmpties(t, func(itb gcf.Iterable[int]) gcf.Iterable[int] {
		itb, _ = gcf.Range(1, 2, -1)
		return itb
	})
}

func ExampleRange() {
	itb, _ := gcf.Range(2, 10, 2)
	fmt.Println(gcf.ToSlice(itb))
	itb, _ = gcf.Range(10, 1, -2)
	fmt.Println(gcf.ToSlice(itb))
	// Output:
	// [2 4 6 8 10]
	// [10 8 6 4 2]
}
