package coverage

import (
	"errors"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// DO NOT EDIT THIS FUNCTION

func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

// var people = People{
// 	{firstName: "Lucas", lastName: "Hansen", birthDay: time.Date(1996, time.May, 24, 0, 0, 0, 0, time.UTC)},
// 	{firstName: "Sofie", lastName: "Johansen", birthDay: time.Date(1997, time.November, 8, 0, 0, 0, 0, time.UTC)},
// 	{firstName: "Oliver", lastName: "Eriksen", birthDay: time.Date(1996, time.April, 13, 0, 0, 0, 0, time.UTC)},
// 	{firstName: "Josef", lastName: "Andersen", birthDay: time.Date(1997, time.February, 19, 0, 0, 0, 0, time.UTC)},
// 	{firstName: "Alice", lastName: "Lund", birthDay: time.Date(1997, time.January, 22, 0, 0, 0, 0, time.UTC)},
// 	{firstName: "Oskar", lastName: "Nilsen", birthDay: time.Date(1996, time.August, 6, 0, 0, 0, 0, time.UTC)},
// 	{firstName: "Emma", lastName: "Olsen", birthDay: time.Date(1997, time.October, 3, 0, 0, 0, 0, time.UTC)},
// 	{firstName: "Oliver", lastName: "Evensen", birthDay: time.Date(1996, time.April, 13, 0, 0, 0, 0, time.UTC)},
// 	{firstName: "Olivia", lastName: "Andersen", birthDay: time.Date(1997, time.February, 19, 0, 0, 0, 0, time.UTC)},
// }

func BenchmarkPeople_Len(t *testing.B) {
	if got, want := people.Len(), 9; got != want {
		t.Errorf("wrong number of people: got %d, want %d", got, want)
	}
}

func BenchmarkPeople_Less(t *testing.B) {
	type args struct {
		i, j int
	}
	testCases := map[string]struct {
		people People
		args   args
		want   bool
	}{
		"in descending order by birthday":   {people: people, args: args{i: 1, j: 3}, want: true},
		"in ascending order by birthday":    {people: people, args: args{i: 4, j: 6}, want: false},
		"in ascending order by first name":  {people: people, args: args{i: 3, j: 8}, want: true},
		"in descending order by first name": {people: people, args: args{i: 8, j: 3}, want: false},
		"in ascending order by last name":   {people: people, args: args{i: 2, j: 7}, want: true},
		"in descending order by last name":  {people: people, args: args{i: 7, j: 2}, want: false},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.B) {
			if got := tc.people.Less(tc.args.i, tc.args.j); got != tc.want {
				t.Errorf("wrong result: got %t, want %t", got, tc.want)
			}
		})
	}
}

func BenchmarkPeople_Swap(t *testing.B) {
	people := People{
		{firstName: "Oliver", lastName: "Eriksen", birthDay: time.Date(1996, time.April, 13, 0, 0, 0, 0, time.UTC)},
		{firstName: "Josef", lastName: "Andersen", birthDay: time.Date(1997, time.February, 19, 0, 0, 0, 0, time.UTC)},
	}
	expected := People{
		{firstName: "Josef", lastName: "Andersen", birthDay: time.Date(1997, time.February, 19, 0, 0, 0, 0, time.UTC)},
		{firstName: "Oliver", lastName: "Eriksen", birthDay: time.Date(1996, time.April, 13, 0, 0, 0, 0, time.UTC)},
	}
	people.Swap(0, 1)
	assert.Equal(t, expected, people, "slices should be the same")
}

func BenchmarkMatrix_New(t *testing.B) {
	testCases := map[string]struct {
		args    string
		want    *Matrix
		wantErr error
	}{
		"correct 3x4 matrix": {
			args: "1 2 3 4 \t \n 5 6 7 8\n\t9 0 2 4 ",
			want: &Matrix{rows: 3, cols: 4, data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 2, 4}},
		},
		"correct 3x2 matrix": {
			args: "\t 1 2    \t \t\n 5 6 \n9 4",
			want: &Matrix{rows: 3, cols: 2, data: []int{1, 2, 5, 6, 9, 4}},
		},
		"different length rows": {
			args: "\t1 2 3 5\t \n7 9 1 ", 
			wantErr: errors.New("Rows need to be the same length"),
		},
		"not numbers in arguments": {
			args: "2 3 t\n8 0 9",
			wantErr: strconv.ErrSyntax,
		},
		"excessive blanks": {
			args: "7 5  7\n3 5 1",
			wantErr: strconv.ErrSyntax,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.B) {
			got, err := New(tc.args)
			if err != nil {
				if e, ok := err.(*strconv.NumError); ok {
					assert.Equal(t, tc.wantErr, errors.Unwrap(e))
				} else {
					assert.Equal(t, tc.wantErr, err)
				}
			} else {
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

func BenchmarkMatrix_Rows(t *testing.B) {
	matrix := Matrix{rows: 2, cols: 5, data: []int{0,1,2,3,4,5,6,7,8,9}}
	want := [][]int{{0,1,2,3,4},{5,6,7,8,9}}
	assert.Equal(t, want, matrix.Rows())
}

func BenchmarkMatrix_Cols(t *testing.B) {
	matrix := Matrix{rows: 5, cols: 2, data: []int{0,1,2,3,4,5,6,7,8,9}}
	want := [][]int{{0,2,4,6,8},{1,3,5,7,9}}
	assert.Equal(t, want, matrix.Cols())
}

func BenchmarkMatrix_Set(t *testing.B) {
	type args struct{
		row, col, value int
	}
	matrix := &Matrix{rows: 5, cols: 2, data: []int{0,1,2,3,4,5,6,7,8,9}}
	testCases := map[string]struct {
		matrix *Matrix
		args args
		want bool
	}{
		"success": {matrix: matrix, args: args{row: 3, col: 1, value: 25}, want: true},
		"nonexistent row": {matrix: matrix, args: args{row:5, col: 0, value: 12}, want: false},
		"nonexistent column": {matrix: matrix, args: args{row:4, col: 2, value: 18}, want: false},
	}
	
	for name, tc := range testCases {
		t.Run(name, func(t *testing.B) {
			assert.Equal(t, tc.want, matrix.Set(tc.args.row, tc.args.col, tc.args.value))
		})
	}
}
