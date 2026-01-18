package util

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ValidationError 验证错误
type ValidationError struct {
    Field   string
    Message string
    Value   interface{}
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("字段 '%s' 验证失败: %s (值: %v)", e.Field, e.Message, e.Value)
}

// Validator 验证器接口
type Validator interface {
    Validate() []ValidationError
}

// Validate 通用验证函数
func Validate(data interface{}) []ValidationError {
    var errors []ValidationError
    
    v := reflect.ValueOf(data)
    t := reflect.TypeOf(data)
    
    // 如果是指针，获取其指向的值
    if v.Kind() == reflect.Ptr {
        v = v.Elem()
        t = t.Elem()
    }
    
    // 处理不同类型的验证
    switch v.Kind() {
    case reflect.Struct:
        errors = validateStruct(v, t)
    case reflect.String:
        if v.String() == "" {
            errors = append(errors, ValidationError{
                Field:   "value",
                Message: "字符串不能为空",
                Value:   v.String(),
            })
        }
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        if v.Int() == 0 {
            errors = append(errors, ValidationError{
                Field:   "value",
                Message: "数字不能为0",
                Value:   v.Int(),
            })
        }
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
        if v.Uint() == 0 {
            errors = append(errors, ValidationError{
                Field:   "value",
                Message: "数字不能为0",
                Value:   v.Uint(),
            })
        }
    case reflect.Float32, reflect.Float64:
        if v.Float() == 0 {
            errors = append(errors, ValidationError{
                Field:   "value",
                Message: "数字不能为0",
                Value:   v.Float(),
            })
        }
    case reflect.Slice, reflect.Array, reflect.Map:
        if v.Len() == 0 {
            errors = append(errors, ValidationError{
                Field:   "value",
                Message: "集合不能为空",
                Value:   v.Interface(),
            })
        }
    }
    
    return errors
}

// validateStruct 验证结构体
func validateStruct(v reflect.Value, t reflect.Type) []ValidationError {
    var errors []ValidationError
    
    for i := 0; i < v.NumField(); i++ {
        fieldValue := v.Field(i)
        fieldType := t.Field(i)
        
        // 获取tag中的验证规则
        tag := fieldType.Tag.Get("validate")
        if tag == "" {
            continue
        }
        
        // 解析验证规则
        rules := parseValidationRules(tag)
        
        // 检查是否必填
        if rules["required"] {
            fieldErrors := validateField(fieldType.Name, fieldValue, rules)
            errors = append(errors, fieldErrors...)
        }
        
        // 如果是嵌套结构体，递归验证
        if fieldValue.Kind() == reflect.Struct {
            nestedErrors := validateStruct(fieldValue, fieldValue.Type())
            // 为嵌套错误添加前缀
            for j := range nestedErrors {
                nestedErrors[j].Field = fmt.Sprintf("%s.%s", fieldType.Name, nestedErrors[j].Field)
            }
            errors = append(errors, nestedErrors...)
        }
        
        // 如果是结构体指针且不为nil，递归验证
        if fieldValue.Kind() == reflect.Ptr && !fieldValue.IsNil() {
            if fieldValue.Elem().Kind() == reflect.Struct {
                nestedErrors := validateStruct(fieldValue.Elem(), fieldValue.Elem().Type())
                for j := range nestedErrors {
                    nestedErrors[j].Field = fmt.Sprintf("%s.%s", fieldType.Name, nestedErrors[j].Field)
                }
                errors = append(errors, nestedErrors...)
            }
        }
    }
    
    return errors
}

// validateField 验证单个字段
func validateField(fieldName string, fieldValue reflect.Value, rules map[string]bool) []ValidationError {
    var errors []ValidationError
    
    // 根据类型进行验证
    switch fieldValue.Kind() {
    case reflect.String:
        if fieldValue.String() == "" {
            errors = append(errors, ValidationError{
                Field:   fieldName,
                Message: "字符串不能为空",
                Value:   fieldValue.String(),
            })
        }
        
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        if fieldValue.Int() == 0 {
            errors = append(errors, ValidationError{
                Field:   fieldName,
                Message: "数字不能为0",
                Value:   fieldValue.Int(),
            })
        }
        
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
        if fieldValue.Uint() == 0 {
            errors = append(errors, ValidationError{
                Field:   fieldName,
                Message: "数字不能为0",
                Value:   fieldValue.Uint(),
            })
        }
        
    case reflect.Float32, reflect.Float64:
        if fieldValue.Float() == 0 {
            errors = append(errors, ValidationError{
                Field:   fieldName,
                Message: "数字不能为0",
                Value:   fieldValue.Float(),
            })
        }
        
    case reflect.Bool:
        // 布尔值不需要验证是否为空
        
    case reflect.Slice, reflect.Array:
        if fieldValue.Len() == 0 {
            errors = append(errors, ValidationError{
                Field:   fieldName,
                Message: "数组/切片不能为空",
                Value:   fieldValue.Interface(),
            })
        }
        
    case reflect.Map:
        if fieldValue.Len() == 0 {
            errors = append(errors, ValidationError{
                Field:   fieldName,
                Message: "映射不能为空",
                Value:   fieldValue.Interface(),
            })
        }
        
    case reflect.Ptr:
        if fieldValue.IsNil() {
            errors = append(errors, ValidationError{
                Field:   fieldName,
                Message: "指针不能为nil",
                Value:   nil,
            })
        }
        
    case reflect.Struct:
        // 结构体已经在validateStruct中处理了
        
    default:
        // 其他类型跳过验证
    }
    
    return errors
}

// parseValidationRules 解析验证规则
func parseValidationRules(tag string) map[string]bool {
    rules := make(map[string]bool)
    
    if tag == "" {
        return rules
    }
    
    parts := strings.Split(tag, ",")
    for _, part := range parts {
        part = strings.TrimSpace(part)
        if part != "" {
            rules[part] = true
        }
    }
    
    return rules
}

// ValidationRule 验证规则
type ValidationRule struct {
    Name  string
    Value string
}

// AdvancedValidator 增强验证器
type AdvancedValidator struct {
    errors []ValidationError
}

// NewValidator 创建验证器
func NewValidator() *AdvancedValidator {
    return &AdvancedValidator{}
}

// ValidateStruct 验证结构体
func (v *AdvancedValidator) ValidateStruct(data interface{}) []ValidationError {
    return Validate(data)
}

// ValidateWithRules 带自定义规则的验证
func ValidateWithRules(data interface{}, fieldRules map[string]string) []ValidationError {
    var errors []ValidationError
    
    val := reflect.ValueOf(data)
    typ := reflect.TypeOf(data)
    
    if val.Kind() == reflect.Ptr {
        val = val.Elem()
        typ = typ.Elem()
    }
    
    if val.Kind() != reflect.Struct {
        return append(errors, ValidationError{
            Field:   "root",
            Message: "需要结构体类型",
            Value:   data,
        })
    }
    
    for i := 0; i < val.NumField(); i++ {
        fieldVal := val.Field(i)
        fieldTyp := typ.Field(i)
        
        // 获取字段名
        fieldName := fieldTyp.Name
        jsonTag := fieldTyp.Tag.Get("json")
        if jsonTag != "" {
            parts := strings.Split(jsonTag, ",")
            if parts[0] != "" {
                fieldName = parts[0]
            }
        }
        
        // 如果有自定义规则，使用自定义规则
        if rules, exists := fieldRules[fieldTyp.Name]; exists {
            fieldErrors := validateFieldWithRules(fieldName, fieldVal, rules)
            errors = append(errors, fieldErrors...)
        } else {
            // 否则使用tag中的规则
            tag := fieldTyp.Tag.Get("validate")
            if tag != "" {
                fieldErrors := validateFieldWithRules(fieldName, fieldVal, tag)
                errors = append(errors, fieldErrors...)
            }
        }
    }
    
    return errors
}

// validateFieldWithRules 使用规则验证字段
func validateFieldWithRules(fieldName string, fieldValue reflect.Value, rules string) []ValidationError {
    var errors []ValidationError
    
    ruleList := parseRules(rules)
    
    for _, rule := range ruleList {
        switch rule.Name {
        case "required":
            if isEmptyValue(fieldValue) {
                errors = append(errors, ValidationError{
                    Field:   fieldName,
                    Message: "该字段是必填项",
                    Value:   getValue(fieldValue),
                })
            }
            
        case "min":
            min, err := strconv.ParseFloat(rule.Value, 64)
            if err == nil {
                num, ok := getNumericValue(fieldValue)
                if ok && num < min {
                    errors = append(errors, ValidationError{
                        Field:   fieldName,
                        Message: fmt.Sprintf("最小值是 %v", min),
                        Value:   getValue(fieldValue),
                    })
                }
            }
            
        case "max":
            max, err := strconv.ParseFloat(rule.Value, 64)
            if err == nil {
                num, ok := getNumericValue(fieldValue)
                if ok && num > max {
                    errors = append(errors, ValidationError{
                        Field:   fieldName,
                        Message: fmt.Sprintf("最大值是 %v", max),
                        Value:   getValue(fieldValue),
                    })
                }
            }
            
        case "email":
            if fieldValue.Kind() == reflect.String {
                email := fieldValue.String()
                if email != "" && !isValidEmail(email) {
                    errors = append(errors, ValidationError{
                        Field:   fieldName,
                        Message: "邮箱格式不正确",
                        Value:   email,
                    })
                }
            }
            
        case "len":
            expectedLen, err := strconv.Atoi(rule.Value)
            if err == nil {
                actualLen := 0
                switch fieldValue.Kind() {
                case reflect.String:
                    actualLen = len(fieldValue.String())
                case reflect.Slice, reflect.Array, reflect.Map:
                    actualLen = fieldValue.Len()
                }
                
                if actualLen != expectedLen {
                    errors = append(errors, ValidationError{
                        Field:   fieldName,
                        Message: fmt.Sprintf("长度必须是 %d", expectedLen),
                        Value:   getValue(fieldValue),
                    })
                }
            }
        }
    }
    
    return errors
}

// 辅助函数
func isEmptyValue(v reflect.Value) bool {
    switch v.Kind() {
    case reflect.String:
        return v.String() == ""
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return v.Int() == 0
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
        return v.Uint() == 0
    case reflect.Float32, reflect.Float64:
        return v.Float() == 0
    case reflect.Bool:
        return false
    case reflect.Slice, reflect.Array, reflect.Map:
        return v.Len() == 0
    case reflect.Ptr, reflect.Interface:
        return v.IsNil()
    case reflect.Struct:
        // 对于time.Time特殊处理
        if v.Type() == reflect.TypeOf(time.Time{}) {
            return v.Interface().(time.Time).IsZero()
        }
        // 其他结构体默认不为空
        return false
    default:
        return false
    }
}

func getValue(v reflect.Value) interface{} {
    if v.IsValid() && v.CanInterface() {
        return v.Interface()
    }
    return nil
}

func getNumericValue(v reflect.Value) (float64, bool) {
    switch v.Kind() {
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return float64(v.Int()), true
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
        return float64(v.Uint()), true
    case reflect.Float32, reflect.Float64:
        return v.Float(), true
    default:
        return 0, false
    }
}

func isValidEmail(email string) bool {
    pattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
    matched, _ := regexp.MatchString(pattern, email)
    return matched
}

func parseRules(rulesStr string) []ValidationRule {
    var rules []ValidationRule
    
    parts := strings.Split(rulesStr, ",")
    for _, part := range parts {
        part = strings.TrimSpace(part)
        if part == "" {
            continue
        }
        
        // 检查是否有参数（如 min=5）
        if strings.Contains(part, "=") {
            kv := strings.SplitN(part, "=", 2)
            rules = append(rules, ValidationRule{
                Name:  strings.TrimSpace(kv[0]),
                Value: strings.TrimSpace(kv[1]),
            })
        } else {
            rules = append(rules, ValidationRule{
                Name:  part,
                Value: "",
            })
        }
    }
    
    return rules
}