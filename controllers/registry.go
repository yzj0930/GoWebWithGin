package controllers

import "reflect"

// ControllerRegistry 控制器注册表
type ControllerRegistry struct {
	controllers []Controller
}

var globalRegistry = &ControllerRegistry{
	controllers: make([]Controller, 0),
}

// Register 注册控制器
func Register(controller Controller) {
	globalRegistry.controllers = append(globalRegistry.controllers, controller)
}

// GetControllers 获取所有注册的控制器
func GetControllers() []Controller {
	return globalRegistry.controllers
}

// GetControllersByType 根据类型获取控制器
func GetControllersByType(targetType Controller) []Controller {
	result := make([]Controller, 0)
	targetTypeName := reflect.TypeOf(targetType).Elem().Name()

	for _, c := range globalRegistry.controllers {
		if reflect.TypeOf(c).Elem().Name() == targetTypeName {
			result = append(result, c)
		}
	}
	return result
}
