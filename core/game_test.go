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
		name  string
		args  args
		want  [Size]cell
		want1 bool
	}{
		// One element
		{
			"Merge [2, 0, 0, 0]",
			args{row: [Size]cell{2, 0, 0, 0}},
			[Size]cell{2, 0, 0, 0},
			false,
		},
		{
			"Merge [4, 0, 0, 0]",
			args{row: [Size]cell{4, 0, 0, 0}},
			[Size]cell{4, 0, 0, 0},
			false,
		},
		{
			"Merge [0, 2, 0, 0]",
			args{row: [Size]cell{0, 2, 0, 0}},
			[Size]cell{2, 0, 0, 0},
			true,
		},
		{
			"Merge [0, 0, 0, 2]",
			args{row: [Size]cell{0, 0, 0, 2}},
			[Size]cell{2, 0, 0, 0},
			true,
		},
		// Two elements
		{
			"Merge [2, 2, 0, 0]",
			args{row: [Size]cell{2, 2, 0, 0}},
			[Size]cell{4, 0, 0, 0},
			true,
		},
		{
			"Merge [4, 4, 0, 0]",
			args{row: [Size]cell{4, 4, 0, 0}},
			[Size]cell{8, 0, 0, 0},
			true,
		},
		{
			"Merge [2, 0, 2, 0]",
			args{row: [Size]cell{2, 0, 2, 0}},
			[Size]cell{4, 0, 0, 0},
			true,
		},
		{
			"Merge [0, 2, 2, 0]",
			args{row: [Size]cell{0, 2, 2, 0}},
			[Size]cell{4, 0, 0, 0},
			true,
		},
		{
			"Merge [0, 2, 0, 2]",
			args{row: [Size]cell{0, 2, 0, 2}},
			[Size]cell{4, 0, 0, 0},
			true,
		},
		{
			"Merge [4, 2, 0, 0]",
			args{row: [Size]cell{4, 2, 0, 0}},
			[Size]cell{4, 2, 0, 0},
			false,
		},
		// Three elements
		{
			"Merge [2, 2, 2, 0]",
			args{row: [Size]cell{2, 2, 2, 0}},
			[Size]cell{4, 2, 0, 0},
			true,
		},
		{
			"Merge [2, 2, 0, 2]",
			args{row: [Size]cell{2, 2, 0, 2}},
			[Size]cell{4, 2, 0, 0},
			true,
		},
		{
			"Merge [0, 2, 2, 2]",
			args{row: [Size]cell{0, 2, 2, 2}},
			[Size]cell{4, 2, 0, 0},
			true,
		},
		{
			"Merge [4, 2, 2, 0]",
			args{row: [Size]cell{4, 2, 2, 0}},
			[Size]cell{4, 4, 0, 0},
			true,
		},
		{
			"Merge [2, 2, 4, 0]",
			args{row: [Size]cell{2, 2, 4, 0}},
			[Size]cell{4, 4, 0, 0},
			true,
		},
		{
			"Merge [4, 0, 2, 2]",
			args{row: [Size]cell{4, 0, 2, 2}},
			[Size]cell{4, 4, 0, 0},
			true,
		},
		{
			"Merge [8, 4, 2, 0]",
			args{row: [Size]cell{8, 4, 2, 0}},
			[Size]cell{8, 4, 2, 0},
			false,
		},
		// Four element
		{
			"Merge [2, 2, 2, 2]",
			args{row: [Size]cell{2, 2, 2, 2}},
			[Size]cell{4, 4, 0, 0},
			true,
		},
		{
			"Merge [4, 2, 2, 4]",
			args{row: [Size]cell{4, 2, 2, 4}},
			[Size]cell{4, 4, 4, 0},
			true,
		},
		{
			"Merge [16, 8, 4, 2]",
			args{row: [Size]cell{16, 8, 4, 2}},
			[Size]cell{16, 8, 4, 2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := merge(tt.args.row)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("merge() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGame_mergeLeft(t *testing.T) {
	type fields struct {
		board board
	}
	tests := []struct {
		name         string
		fields       fields
		gotFields    fields
		wantHasMoved bool
	}{
		{
			"MergeLeft",
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
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				board: tt.fields.board,
			}
			gotHasMoved := g.mergeLeft()
			if !reflect.DeepEqual(g.board, tt.gotFields.board) {
				t.Errorf("mergeLeft got %v, want %v", g.board, tt.gotFields.board)
			}
			if gotHasMoved != tt.wantHasMoved {
				t.Errorf("Game.mergeLeft() = %v, want %v", gotHasMoved, tt.wantHasMoved)
			}
		})
	}
}

func TestGame_mergeRight(t *testing.T) {
	type fields struct {
		board board
	}
	tests := []struct {
		name         string
		fields       fields
		gotFields    fields
		wantHasMoved bool
	}{
		{
			"MergeRight",
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
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				board: tt.fields.board,
			}
			gotHasMoved := g.mergeRight()
			if !reflect.DeepEqual(g.board, tt.gotFields.board) {
				t.Errorf("mergeRight got %v, want %v", g.board, tt.gotFields.board)
			}
			if gotHasMoved != tt.wantHasMoved {
				t.Errorf("Game.mergeRight() = %v, want %v", gotHasMoved, tt.wantHasMoved)
			}
		})
	}
}

func TestGame_mergeTop(t *testing.T) {
	type fields struct {
		board board
	}
	tests := []struct {
		name         string
		fields       fields
		gotFields    fields
		wantHasMoved bool
	}{
		{
			"MergeTop",
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
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				board: tt.fields.board,
			}
			gotHasMoved := g.mergeTop()
			if !reflect.DeepEqual(g.board, tt.gotFields.board) {
				t.Errorf("mergeTop got %v, want %v", g.board, tt.gotFields.board)
			}
			if gotHasMoved != tt.wantHasMoved {
				t.Errorf("Game.mergeTop() = %v, want %v", gotHasMoved, tt.wantHasMoved)
			}
		})
	}
}

func TestGame_mergeBottom(t *testing.T) {
	type fields struct {
		board board
	}
	tests := []struct {
		name         string
		fields       fields
		gotFields    fields
		wantHasMoved bool
	}{
		{
			"MergeBottom",
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
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				board: tt.fields.board,
			}
			gotHasMoved := g.mergeBottom()
			if !reflect.DeepEqual(g.board, tt.gotFields.board) {
				t.Errorf("mergeBottom got %v, want %v", g.board, tt.gotFields.board)
			}
			if gotHasMoved != tt.wantHasMoved {
				t.Errorf("Game.mergeBottom() = %v, want %v", gotHasMoved, tt.wantHasMoved)
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

func TestGame_canMove(t *testing.T) {
	type fields struct {
		board board
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"canMove A",
			fields{board{
				[Size]int{2, 0, 0, 0},
				[Size]int{0, 0, 0, 0},
				[Size]int{0, 0, 0, 0},
				[Size]int{0, 0, 0, 0},
			}},
			true,
		},
		{
			"canMove B",
			fields{board{
				[Size]int{2, 2, 2, 2},
				[Size]int{0, 0, 0, 0},
				[Size]int{0, 0, 0, 0},
				[Size]int{0, 0, 0, 0},
			}},
			true,
		},
		{
			"canMove C",
			fields{board{
				[Size]int{2, 0, 0, 0},
				[Size]int{2, 0, 0, 0},
				[Size]int{2, 0, 0, 0},
				[Size]int{2, 0, 0, 0},
			}},
			true,
		},
		{
			"canMove D",
			fields{board{
				[Size]int{2, 0, 0, 0},
				[Size]int{0, 2, 0, 0},
				[Size]int{0, 0, 2, 0},
				[Size]int{0, 0, 0, 2},
			}},
			true,
		},
		{
			"canMove E",
			fields{board{
				[Size]int{2, 2, 2, 2},
				[Size]int{2, 2, 2, 2},
				[Size]int{2, 2, 2, 2},
				[Size]int{2, 2, 2, 2},
			}},
			true,
		},
		{
			"canMove F",
			fields{board{
				[Size]int{2, 16, 2, 16},
				[Size]int{4, 8, 4, 8},
				[Size]int{8, 4, 8, 4},
				[Size]int{16, 2, 16, 2},
			}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				board: tt.fields.board,
			}
			if got := g.canMove(); got != tt.want {
				t.Errorf("Game.canMove() = %v, want %v", got, tt.want)
			}
		})
	}
}
