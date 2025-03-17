package query

import (
	"reflect"
	"strings"
	"testing"
)

func TestQuery(t *testing.T) {
	page := DefaultPage(-1)
	t.Log(page.Page(), page.Limit(), page.Sort(), page.Offset())

	SetMaxSize(1) // defaultMaxSize = 10

	t.Log(defaultMaxSize)

	page = NewPage(-1, 100, "id,-name")
	t.Log(page.Page(), page.Limit(), page.Sort(), page.Offset())
}

func TestParams_ConvertToPage(t *testing.T) {
	p := &Params{
		Page:  1,
		Limit: 50,
		Sort:  "age,-name",
	}
	order, limit, offset := p.ConvertToPage()
	t.Logf("order=%s, limit=%d, offset=%d", order, limit, offset)
}

func TestParams_ConvertToGormConditions(t *testing.T) {
	type args struct {
		columns []Column
	}
	tt := struct {
		name    string
		args    args
		want    string
		want1   []interface{}
		wantErr bool
	}{
		name: "2 column or",
		args: args{
			columns: []Column{
				{
					Name:  "name",
					Value: "ZhangSan",
					Logic: "||",
				},
				{
					Name:  "mail",
					Value: "123123",
					Exp:   "!=",
				},
			},
		},
		want:    "name = ? OR mail <> ?",
		want1:   []interface{}{"ZhangSan", "123123"},
		wantErr: false,
	}
	params := &Params{
		Columns: tt.args.columns,
	}
	got, got1, err := params.ConvertToGormConditions()
	if (err != nil) != tt.wantErr {
		t.Errorf("ConvertToGormConditions() error = %v, wantErr = %v", err, tt.wantErr)
		return
	}
	if got != tt.want {
		t.Errorf("ConvertToGormConditions() got = [%v], want = [%v]", got, tt.want)
	}
	if !reflect.DeepEqual(got1, tt.want1) {
		t.Errorf("ConvertToGormConditions() got1 = [%v], want = [%v]", got1, tt.want1)
	}

	t.Log(got, "|", got1)
	got = strings.Replace(got, "?", "%v", -1)
	t.Logf(got, got1...)
}

func TestParams_ConvertToStructColumn(t *testing.T) {
	d := struct {
		Name string `json:"name"`
		Mail string `json:"mail"`
	}{
		Name: "D---NAME123",
		Mail: "D---MAIL123",
	}
	p := &Params{
		Columns: []Column{
			{
				Name:  "name",
				Value: "ZhangSan",
				Logic: "||",
			},
			{
				Name:  "mail",
				Value: "123123",
				Exp:   "!=",
			},
		},
		RowColumns: d,
	}

	sql, params, e := p.ConvertToGormConditions()
	if e != nil {
		t.Error(e)
	}
	order, limit, offset := p.ConvertToPage()
	t.Log(order, limit, offset)
	t.Log(sql, params, "success")
}
