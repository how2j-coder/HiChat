package common

import (
	"com/chat/service/internal/dao"
	"com/chat/service/internal/database"
	customModel "com/chat/service/internal/model"
	"context"
	"github.com/casbin/casbin/v2"
	casbinModel "github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"sync"
)

var (
	enforcer  *casbin.SyncedEnforcer
	once      sync.Once
	enErr     error
	casbinDao dao.CasbinRuleDao
)

type adapterCasbin struct {
	Enforcer    *casbin.SyncedEnforcer
	EnforcerErr error
	ctx         context.Context
}

// CasbinEnforcer 获取adapter单例对象
func CasbinEnforcer(ctx context.Context) (enforcer *casbin.SyncedEnforcer, err error) {
	ac := newAdapter(ctx)
	enforcer = ac.Enforcer
	err = ac.EnforcerErr
	return
}

// 初始化adapter
func newAdapter(ctx context.Context) *adapterCasbin {
	adapter := new(adapterCasbin)
	adapter.ctx = ctx
	once.Do(func() {
		casbinDao = dao.NewCasbinRuleDao(database.GetDB())
		enforcer, enErr = initPolicy(adapter)
	})
	if enErr == nil && enforcer != nil {
		enforcer.SetAdapter(adapter)
	}
	adapter.Enforcer, adapter.EnforcerErr = enforcer, enErr
	return adapter
}

func initPolicy(adapter *adapterCasbin) (*casbin.SyncedEnforcer, error) {
	m, _ := casbinModel.NewModelFromString(`
	[request_definition]
	r = sub, obj, act
	
	[policy_definition]
	p = sub, obj, act
	
	[role_definition]
	g = _, _
	
	[policy_effect]
	e = some(where (p.eft == allow))
	
	[matchers]
	m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
	`)
	newEnforcer, err := casbin.NewSyncedEnforcer(m, adapter)
	return newEnforcer, err
}

func (a *adapterCasbin) dropTable() (err error) {
	return
}

func (a *adapterCasbin) createTable() (err error) {
	return
}

// SavePolicy saves policy to database.
func (a *adapterCasbin) SavePolicy(model casbinModel.Model) (err error) {
	err = a.dropTable()
	if err != nil {
		return
	}
	err = a.createTable()
	if err != nil {
	}
	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			err := casbinDao.Create(a.ctx, line)
			if err != nil {
				return err
			}
		}
	}

	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			err := casbinDao.Create(a.ctx, line)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// LoadPolicy loads policy from a database.
func (a *adapterCasbin) LoadPolicy(model casbinModel.Model) error {
	var lines []customModel.CasbinRule
	if err := casbinDao.GetByColumns(a.ctx, &lines); err != nil {
		return err
	}
	for _, line := range lines {
		loadPolicyLine(&line, model)
	}
	return nil
}

// AddPolicy adds a policy rule to the storage.
func (a *adapterCasbin) AddPolicy(sec string, ptype string, rule []string) error {
	line := savePolicyLine(ptype, rule)
	err := casbinDao.Create(a.ctx, line)
	return err
}

// RemovePolicy removes a policy rule from the storage.
func (a *adapterCasbin) RemovePolicy(sec string, ptype string, rule []string) error {
	line := savePolicyLine(ptype, rule)
	err := rawDelete(a.ctx, line)
	return err
}

func (a *adapterCasbin) AddPolicies(sec string, ptype string, rules [][]string) error {
	lines := make([]customModel.CasbinRule, len(rules))
	for k, rule := range rules {
		lines[k] = *savePolicyLine(ptype, rule)
	}
	err := casbinDao.CreateBatch(a.ctx, &lines)
	return err
}

// RemovePolicies removes policy rules from the storage. This is part of the Auto-Save feature.
func (a *adapterCasbin) RemovePolicies(sec string, ptype string, rules [][]string) error {
	for _, rule := range rules {
		err := a.RemovePolicy(sec, ptype, rule)
		if err != nil {
			return err
		}
	}
	return nil
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *adapterCasbin) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	line := &customModel.CasbinRule{}
	line.Ptype = ptype
	if fieldIndex <= 0 && 0 < fieldIndex+len(fieldValues) {
		line.V0 = fieldValues[0-fieldIndex]
	}
	if fieldIndex <= 1 && 1 < fieldIndex+len(fieldValues) {
		line.V1 = fieldValues[1-fieldIndex]
	}
	if fieldIndex <= 2 && 2 < fieldIndex+len(fieldValues) {
		line.V2 = fieldValues[2-fieldIndex]
	}
	if fieldIndex <= 3 && 3 < fieldIndex+len(fieldValues) {
		line.V3 = fieldValues[3-fieldIndex]
	}
	if fieldIndex <= 4 && 4 < fieldIndex+len(fieldValues) {
		line.V4 = fieldValues[4-fieldIndex]
	}
	if fieldIndex <= 5 && 5 < fieldIndex+len(fieldValues) {
		line.V5 = fieldValues[5-fieldIndex]
	}
	err := rawDelete(a.ctx, line)
	return err
}

func rawDelete(ctx context.Context, line *customModel.CasbinRule) error {
	queryArgs := []interface{}{line.Ptype}
	queryStr := "ptype = ?"
	if line.V0 != "" {
		queryStr += " and v0 = ?"
		queryArgs = append(queryArgs, line.V0)
	}
	if line.V1 != "" {
		queryStr += " and v1 = ?"
		queryArgs = append(queryArgs, line.V1)
	}
	if line.V2 != "" {
		queryStr += " and v2 = ?"
		queryArgs = append(queryArgs, line.V2)
	}
	if line.V3 != "" {
		queryStr += " and v3 = ?"
		queryArgs = append(queryArgs, line.V3)
	}
	if line.V4 != "" {
		queryStr += " and v4 = ?"
		queryArgs = append(queryArgs, line.V4)
	}
	if line.V5 != "" {
		queryStr += " and v5 = ?"
		queryArgs = append(queryArgs, line.V5)
	}
	args := append([]interface{}{queryStr}, queryArgs...)
	return casbinDao.Delete(ctx, line, args...)
}

func savePolicyLine(ptype string, rule []string) *customModel.CasbinRule {
	line := &customModel.CasbinRule{}
	line.Ptype = ptype
	if len(rule) > 0 {
		line.V0 = rule[0]
	}
	if len(rule) > 1 {
		line.V1 = rule[1]
	}
	if len(rule) > 2 {
		line.V2 = rule[2]
	}
	if len(rule) > 3 {
		line.V3 = rule[3]
	}
	if len(rule) > 4 {
		line.V4 = rule[4]
	}
	if len(rule) > 5 {
		line.V5 = rule[5]
	}
	return line
}

func loadPolicyLine(line *customModel.CasbinRule, model casbinModel.Model) {
	lineText := line.Ptype
	if line.V0 != "" {
		lineText += ", " + line.V0
	}
	if line.V1 != "" {
		lineText += ", " + line.V1
	}
	if line.V2 != "" {
		lineText += ", " + line.V2
	}
	if line.V3 != "" {
		lineText += ", " + line.V3
	}
	if line.V4 != "" {
		lineText += ", " + line.V4
	}
	if line.V5 != "" {
		lineText += ", " + line.V5
	}
	_ = persist.LoadPolicyLine(lineText, model)
}
