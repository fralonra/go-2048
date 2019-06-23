package core

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		row [Size]cell
	}
	tests := []struct {
		name string
		args args
		want [Size]cell
	}{
		// One element
		{
			"Merge [2, 0, 0, 0]",
			args{row: [Size]cell{2, 0, 0, 0}},
			[Size]cell{2, 0, 0, 0},
		},
		{
			"Merge [4, 0, 0, 0]",
			args{row: [Size]cell{4, 0, 0, 0}},
			[Size]cell{4, 0, 0, 0},
		},
		{
			"Merge [0, 2, 0, 0]",
			args{row: [Size]cell{0, 2, 0, 0}},
			[Size]cell{2, 0, 0, 0},
		},
		{
			"Merge [0, 0, 0, 2]",
			args{row: [Size]cell{0, 0, 0, 2}},
			[Size]cell{2, 0, 0, 0},
		},
		// Two elements
		{
			"Merge [2, 2, 0, 0]",
			args{row: [Size]cell{2, 2, 0, 0}},
			[Size]cell{4, 0, 0, 0},
		},
		{
			"Merge [4, 4, 0, 0]",
			args{row: [Size]cell{4, 4, 0, 0}},
			[Size]cell{8, 0, 0, 0},
		},
		{
			"Merge [2, 0, 2, 0]",
			args{row: [Size]cell{2, 0, 2, 0}},
			[Size]cell{4, 0, 0, 0},
		},
		{
			"Merge [0, 2, 2, 0]",
			args{row: [Size]cell{0, 2, 2, 0}},
			[Size]cell{4, 0, 0, 0},
		},
		{
			"Merge [0, 2, 0, 2]",
			args{row: [Size]cell{0, 2, 0, 2}},
			[Size]cell{4, 0, 0, 0},
		},
		{
			"Merge [4, 2, 0, 0]",
			args{row: [Size]cell{4, 2, 0, 0}},
			[Size]cell{4, 2, 0, 0},
		},
		// Three elements
		{
			"Merge [2, 2, 2, 0]",
			args{row: [Size]cell{2, 2, 2, 0}},
			[Size]cell{4, 2, 0, 0},
		},
		{
			"Merge [0, 2, 2, 2]",
			args{row: [Size]cell{0, 2, 2, 2}},
			[Size]cell{4, 2, 0, 0},
		},
		{
			"Merge [4, 2, 2, 0]",
			args{row: [Size]cell{4, 2, 2, 0}},
			[Size]cell{4, 4, 0, 0},
		},
		{
			"Merge [2, 2, 4, 0]",
			args{row: [Size]cell{2, 2, 4, 0}},
			[Size]cell{4, 4, 0, 0},
		},
		{
			"Merge [4, 0, 2, 2]",
			args{row: [Size]cell{4, 0, 2, 2}},
			[Size]cell{4, 4, 0, 0},
		},
		{
			"Merge [8, 4, 2, 0]",
			args{row: [Size]cell{8, 4, 2, 0}},
			[Size]cell{8, 4, 2, 0},
		},
		// Four element
		{
			"Merge [2, 2, 2, 2]",
			args{row: [Size]cell{2, 2, 2, 2}},
			[Size]cell{4, 4, 0, 0},
		},
		{
			"Merge [4, 2, 2, 4]",
			args{row: [Size]cell{4, 2, 2, 4}},
			[Size]cell{4, 4, 4, 0},
		},
		{
			"Merge [16, 8, 4, 2]",
			args{row: [Size]cell{16, 8, 4, 2}},
			[Size]cell{16, 8, 4, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.row); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGame_ToLeft(t *testing.T) {
	type fields struct {
		board board
	}
	tests := []struct {
		name   string
		fields fields
		want   fields
	}{
		{
			"ToLeft",
			fields{board{
				[Size]int{2, 0, 0, 0},
				[Size]int{2, 2, 0, 0},
				[Size]int{2, 2, 2, 0},
				[Size]int{4, 2, 0, 0},
			}},
			fields{board{
				[Size]int{2, 0, 0, 0},
				[Size]int{4, 0, 0, 0},
				[Size]int{4, 2, 0, 0},
				[Size]int{4, 2, 0, 0},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				board: tt.fields.board,
			}
			g.ToLeft()
			if !reflect.DeepEqual(g.board, tt.want.board) {
				t.Errorf("ToLeft() = %v, want %v", g.board, tt.want.board)
			}
		})
	}
}

func TestGame_ToRight(t *testing.T) {
	type fields struct {
		board board
	}
	tests := []struct {
		name   string
		fields fields
		want   fields
	}{
		{
			"ToRight",
			fields{board{
				[Size]int{2, 0, 0, 0},
				[Size]int{2, 2, 0, 0},
				[Size]int{2, 2, 2, 0},
				[Size]int{4, 2, 0, 0},
			}},
			fields{board{
				[Size]int{0, 0, 0, 2},
				[Size]int{0, 0, 0, 4},
				[Size]int{0, 0, 2, 4},
				[Size]int{0, 0, 4, 2},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				board: tt.fields.board,
			}
			g.ToRight()
			if !reflect.DeepEqual(g.board, tt.want.board) {
				t.Errorf("ToRight() = %v, want %v", g.board, tt.want.board)
			}
		})
	}
}

func TestGame_ToTop(t *testing.T) {
	type fields struct {
		board board
	}
	tests := []struct {
		name   string
		fields fields
		want   fields
	}{
		{
			"ToTop",
			fields{board{
				[Size]int{2, 0, 0, 0},
				[Size]int{2, 2, 0, 4},
				[Size]int{2, 2, 2, 0},
				[Size]int{4, 2, 0, 2},
			}},
			fields{board{
				[Size]int{4, 4, 2, 4},
				[Size]int{2, 2, 0, 2},
				[Size]int{4, 0, 0, 0},
				[Size]int{0, 0, 0, 0},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				board: tt.fields.board,
			}
			g.ToTop()
			if !reflect.DeepEqual(g.board, tt.want.board) {
				t.Errorf("ToTop() = %v, want %v", g.board, tt.want.board)
			}
		})
	}
}

func TestGame_ToBottom(t *testing.T) {
	type fields struct {
		board board
	}
	tests := []struct {
		name   string
		fields fields
		want   fields
	}{
		{
			"ToBottom",
			fields{board{
				[Size]int{2, 0, 0, 0},
				[Size]int{2, 2, 0, 4},
				[Size]int{2, 2, 2, 0},
				[Size]int{4, 2, 0, 2},
			}},
			fields{board{
				[Size]int{0, 0, 0, 0},
				[Size]int{2, 0, 0, 0},
				[Size]int{4, 2, 0, 4},
				[Size]int{4, 4, 2, 2},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				board: tt.fields.board,
			}
			g.ToBottom()
			if !reflect.DeepEqual(g.board, tt.want.board) {
				t.Errorf("ToBottom() = %v, want %v", g.board, tt.want.board)
			}
		})
	}
}

func TestGame_Get(t *testing.T) {
	type fields struct {
		board board
	}
	type args struct {
		row    int
		colomn int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   cell
	}{
		{
			"Get [0][0]",
			fields{board{
				[Size]int{2, 0, 0, 0},
				[Size]int{0, 2, 0, 0},
				[Size]int{4, 4, 2, 0},
				[Size]int{8, 4, 0, 0},
			}},
			args{0, 0},
			2,
		},
		{
			"Get [3][3]",
			fields{board{
				[Size]int{2, 0, 0, 0},
				[Size]int{0, 2, 0, 0},
				[Size]int{4, 4, 2, 0},
				[Size]int{8, 4, 0, 0},
			}},
			args{3, 3},
			0,
		},
		{
			"Get [2][3]",
			fields{board{
				[Size]int{2, 0, 0, 0},
				[Size]int{0, 2, 0, 0},
				[Size]int{4, 4, 2, 0},
				[Size]int{8, 4, 0, 0},
			}},
			args{2, 3},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				board: tt.fields.board,
			}
			if got := g.Get(tt.args.row, tt.args.colomn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Game.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGame_GetRow(t *testing.T) {
	type fields struct {
		board board
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   [Size]cell
	}{
		{
			"Get [0]",
			fields{board{
				[Size]int{2, 0, 0, 0},
				[Size]int{0, 2, 0, 0},
				[Size]int{4, 4, 2, 0},
				[Size]int{8, 4, 0, 0},
			}},
			args{0},
			[Size]int{2, 0, 0, 0},
		},
		{
			"Get [3]",
			fields{board{
				[Size]int{2, 0, 0, 0},
				[Size]int{0, 2, 0, 0},
				[Size]int{4, 4, 2, 0},
				[Size]int{8, 4, 0, 0},
			}},
			args{3},
			[Size]int{8, 4, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				board: tt.fields.board,
			}
			if got := g.GetRow(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Game.GetRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
