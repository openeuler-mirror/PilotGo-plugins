package service

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/global"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum/rule"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum/script"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/dangerous_rule/model"
)

func getetRulesFromRedis() ([]model.DangerousRule, error) {
	var rules []model.DangerousRule
	err := global.App.Redis.Get(DangerousRuleKey, &rules)
	if err != nil {
		return nil, err
	}
	return rules, nil
}

type Finding struct {
	RuleID      int             `json:"rule_id"`
	Description string          `json:"description"`
	Action      rule.ActionType `json:"action"`
	Line        int             `json:"line"`    // 1-based
	Snippet     string          `json:"snippet"` // 匹配文本或行片段
	Match       string          `json:"match"`   // 匹配到的文本
}

type DetectRule struct {
	ID          int             `json:"id"`
	Description string          `json:"description"`
	Action      rule.ActionType `json:"action"`
	Regex       *regexp.Regexp  `json:"regex"`
	Keywords    []string        `json:"keywords"`
}

func DetectRealtimely(script string, scriptType script.ScriptType) ([]DetectRule, error) {
	rules, err := detectRules(scriptType)
	if err != nil {
		return nil, err
	}
	return rules, nil
}

// Detect 脚本检测主方法
func Detect(script string, scriptType script.ScriptType) ([]Finding, error) {
	return detectInternal(script, scriptType)
}

// DetectWithVars 支持变量替换
func DetectWithVars(script string, scriptType script.ScriptType, params map[string]string) ([]Finding, error) {
	expanded := expandSimpleVars(script, params)
	return detectInternal(expanded, scriptType)
}

func detectInternal(script string, scriptType script.ScriptType) ([]Finding, error) {
	rules, err := detectRules(scriptType)
	if err != nil {
		return nil, err
	}
	if len(rules) == 0 {
		return nil, nil
	}

	lines := splitLines(script)
	var findings []Finding

	for i, line := range lines {
		trim := strings.TrimSpace(line)
		if trim == "" || isCommentLine(trim) {
			continue
		}
		lower := strings.ToLower(trim)

		for _, r := range rules {
			matched := false
			var matchText string

			// 正则匹配
			if r.Regex != nil {
				if loc := r.Regex.FindStringIndex(trim); loc != nil {
					matched = true
					matchText = trim[loc[0]:loc[1]]
				}
			}

			// 关键字匹配
			if !matched && len(r.Keywords) > 0 {
				for _, k := range r.Keywords {
					if strings.Contains(lower, strings.ToLower(k)) {
						matched = true
						matchText = k
						break
					}
				}
			}

			if matched {
				findings = append(findings, Finding{
					RuleID:      r.ID,
					Description: r.Description,
					Action:      r.Action,
					Line:        i + 1,
					Snippet:     trim,
					Match:       matchText,
				})
			}
		}
	}

	// 排序：Block 优先，Warning 次之；同类型按行号升序
	sort.Slice(findings, func(i, j int) bool {
		if findings[i].Action == findings[j].Action {
			return findings[i].Line < findings[j].Line
		}
		return findings[i].Action == rule.Block
	})

	return findings, nil
}
func detectRules(scriptType script.ScriptType) ([]DetectRule, error) {
	// 1. 从 Redis 获取高危规则
	dangerousRules, err := getetRulesFromRedis()
	if err != nil {
		return nil, fmt.Errorf("获取高危命令失败: %w", err)
	}

	// 2. 转换为可检测规则（Regex 或 Keywords）
	var rules []DetectRule
	for _, r := range dangerousRules {
		// 如果脚本类型不在规则的 ScriptTypes 中，跳过
		if !containsScriptType(r.ScriptTypes, scriptType) {
			continue
		}
		rules = append(rules, toDetectRule(r))
	}
	return rules, nil
}

// 判断 ScriptTypes 是否包含某个类型
func containsScriptType(arr script.ScriptTypeArr, t script.ScriptType) bool {
	for _, v := range arr {
		if v == t {
			return true
		}
	}
	return false
}

// 将 DangerousRule 转换为 DetectRule
func toDetectRule(r model.DangerousRule) DetectRule {
	dr := DetectRule{
		ID:          r.ID,
		Description: r.Description,
		Action:      r.Action,
	}

	if r.Expression != "" {
		if re, err := regexp.Compile(r.Expression); err == nil {
			dr.Regex = re
		} else {
			dr.Keywords = []string{r.Expression}
		}
	}

	return dr
}

// 简单变量替换 ${VAR} 或 $VAR
func expandSimpleVars(script string, params map[string]string) string {
	out := script
	for k, v := range params {
		out = strings.ReplaceAll(out, fmt.Sprintf("${%s}", k), v)
		out = strings.ReplaceAll(out, fmt.Sprintf("$%s", k), v)
	}
	return out
}

// 判断是否为注释行
func isCommentLine(line string) bool {
	trim := strings.TrimSpace(line)
	return strings.HasPrefix(trim, "#") || strings.HasPrefix(trim, "//")
}

// 按行切分脚本
func splitLines(s string) []string {
	return strings.Split(strings.ReplaceAll(s, "\r\n", "\n"), "\n")
}
