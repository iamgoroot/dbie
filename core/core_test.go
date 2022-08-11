package core

import (
	"errors"
	"github.com/iamgoroot/dbie"
	"reflect"
	"sync/atomic"
	"testing"
)

func TestGenericBackend_SelectOne(t *testing.T) {
	type entity struct {
		ID  int
		Val string
	}
	type testCase struct {
		name         string
		wantErr      error
		wantPage     dbie.Page
		wantVal      string
		wantField    string
		wantOperator dbie.Op
		wantOrders   []dbie.Sort
		returnPage   dbie.Paginated[entity]
		wantResult   entity
	}
	for _, tt := range []testCase{
		{
			name:         "success. no sort",
			wantErr:      nil,
			wantPage:     dbie.Page{Limit: 1},
			wantVal:      "1",
			wantOperator: dbie.Eq,
			wantField:    "id",
			returnPage: dbie.Paginated[entity]{
				Data: []entity{
					{ID: 1, Val: "val1"},
				},
			},
			wantResult: entity{ID: 1, Val: "val1"},
		},
		{
			name:         "success. multiple sorting",
			wantErr:      nil,
			wantPage:     dbie.Page{Limit: 1},
			wantVal:      "1",
			wantOperator: dbie.Eq,
			wantField:    "id",
			wantOrders:   []dbie.Sort{{Field: "id", Order: dbie.ASC}, {Field: "id2", Order: dbie.DESC}},
			returnPage: dbie.Paginated[entity]{
				Data: []entity{
					{ID: 1, Val: "val1"},
				},
			},
			wantResult: entity{ID: 1, Val: "val1"},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			core := CoreMock[entity]{}.
				SelectPageExpect(tt.wantPage, tt.wantField, tt.wantOperator, tt.wantVal, tt.wantOrders...)(tt.returnPage, nil)
			repo := GenericBackend[entity]{Core: core}
			one, err := repo.SelectOne("id", dbie.Eq, "1", tt.wantOrders...)
			switch {
			case err != nil:
				t.Fatalf("expected error '%v' got %v", tt.wantErr, err)
			case !reflect.DeepEqual(tt.wantResult, one):
				t.Fatalf("expected result to be '%v' got %v", tt.wantResult, one)
			}
		})
	}
}

func TestGenericBackend_Select(t *testing.T) {
	type entity struct {
		ID  int
		Val string
	}
	type testCase struct {
		name         string
		wantErr      error
		wantPage     dbie.Page
		wantOperator dbie.Op
		returnPage   dbie.Paginated[entity]
		wantResult   []entity
		wantVal      string
	}
	for _, tt := range []testCase{
		{
			name:     "success",
			wantErr:  nil,
			wantPage: dbie.Page{Limit: 20}, //want page to be 20 as default
			wantVal:  "1",
			returnPage: dbie.Paginated[entity]{
				Data: []entity{{ID: 1, Val: "val1"}},
			},
			wantResult: []entity{
				{ID: 1, Val: "val1"},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			repo := GenericBackend[entity]{
				Core: CoreMock[entity]{
					SelectPageMock: func(page dbie.Page, field string, operator dbie.Op, val any, orders ...dbie.Sort) (items dbie.Paginated[entity], err error) {
						switch {
						case page != tt.wantPage:
							t.Fatalf("expected page to be '%v' got %v", tt.wantPage, page)
						case operator != tt.wantOperator:
							t.Fatalf("expected operator to be '%v' got %s)", tt.wantOperator, operator)
						case !reflect.DeepEqual(val, tt.wantVal):
							t.Fatalf("expected val to be '%v' got %v", tt.wantVal, val)
						}
						return tt.returnPage, err
					},
				},
			}
			one, err := repo.Select("val", dbie.Eq, "1")

			switch {
			case err != nil:
				t.Fatalf("expected error '%v' got %v", tt.wantErr, err)
			case !reflect.DeepEqual(tt.wantResult, one):
				t.Fatalf("expected result to be '%v' got %v", tt.wantResult, one)
			}
		})
	}
}

func TestGenericBackend_Close(t *testing.T) {
	type entity struct {
		ID  int
		Val string
	}
	mockRepoClose := func(errFunc func() error) dbie.Repo[entity] {
		return GenericBackend[entity]{
			Core: CoreMock[entity]{
				CloseMock: errFunc,
			},
		}
	}
	t.Run("test close triggers core.Close", func(t *testing.T) {
		var count int32
		got := mockRepoClose(func() error {
			atomic.AddInt32(&count, 1)
			return nil
		}).Close()
		if got != nil {
			t.Fatalf("expected no error got `%s`)", got)
		}
		if atomic.LoadInt32(&count) != 1 {
			t.Fatalf("expected core.Close to be called once got %d)", atomic.LoadInt32(&count))
		}
	})
	t.Run("test close passes error from core", func(t *testing.T) {
		expectError := errors.New("close error")
		got := mockRepoClose(func() error {
			return expectError
		}).Close()
		if got == nil || got.Error() != "close error" {
			t.Fatalf("expected `close error` got `%s`)", got)
		}
	})
}
