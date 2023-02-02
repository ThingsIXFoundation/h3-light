package h3light_test

import (
	"testing"

	. "github.com/ThingsIXFoundation/h3-light"
)

func TestDatabaseCellFromCell(t *testing.T) {
	type args struct {
		cell uint64
	}
	tests := []struct {
		name string
		args args
		want DatabaseCell
	}{
		{
			name: "test DatabaseCellFromCell(85283473fffffff)=1406434",
			args: args{cell: uint64(MustCellFromString("85283473fffffff"))},
			want: "1406434",
		},
		{
			name: "test DatabaseCellFromCell(8019fffffffffff)=0c",
			args: args{cell: uint64(MustCellFromString("8019fffffffffff"))},
			want: "0c",
		},
		{
			name: "test DatabaseCellFromCell(83184dfffffffff)=0c115",
			args: args{cell: uint64(MustCellFromString("83184dfffffffff"))},
			want: "0c115",
		},
		{
			name: "test DatabaseCellFromCell(83186bfffffffff)=0c153",
			args: args{cell: uint64(MustCellFromString("83186bfffffffff"))},
			want: "0c153",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DatabaseCellFromCell(tt.args.cell); got != tt.want {
				t.Errorf("DatabaseCellFromCell() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseCell_ToCell(t *testing.T) {
	tests := []struct {
		name string
		dc   DatabaseCell
		want uint64
	}{
		{
			name: "test res5",
			dc:   DatabaseCellFromCell(uint64(MustCellFromString("85283473fffffff"))),
			want: uint64(MustCellFromString("85283473fffffff")),
		},
		{
			name: "test res0",
			dc:   DatabaseCellFromCell(uint64(MustCellFromString("8019fffffffffff"))),
			want: uint64(MustCellFromString("8019fffffffffff")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dc.Cell(); uint64(got) != tt.want {
				t.Errorf("ToCell() = %x, want %x", got, tt.want)
			}
		})
	}
}

func TestDatabaseCell_CellPtr(t *testing.T) {
	tests := []struct {
		name string
		dc   *DatabaseCell
		want *Cell
	}{
		{
			name: "test nil",
			dc:   nil,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dc.CellPtr(); got != tt.want {
				t.Errorf("DatabaseCell.CellPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}
