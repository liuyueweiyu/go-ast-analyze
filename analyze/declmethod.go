package main



// GetDeclType GetDeclType
func (t *SimpleType) GetDeclType() string {
	return "SimpleType"
}

// SetTypeAndName SetTypeAndName
func (t *SimpleType) SetTypeAndName(types string, name string) {
	t.Type = types
}

// GetTypeName GetTypeName
func (t *SimpleType) GetTypeName() string {
	return t.Type
}

// GetRetuArgs GetRetuArgs
func (t *SimpleType) GetRetuArgs() []*BaseValue {
	return nil
}

// GetDeclType GetDeclType
func (t *ComplexType) GetDeclType() string {
	return "ComplexType"
}

// SetTypeAndName SetTypeAndName
func (t *ComplexType) SetTypeAndName(types string, name string) {
	t.Type = types
}

// GetTypeName GetTypeName
func (t *ComplexType) GetTypeName() string {
	return t.Type
}

// GetRetuArgs GetRetuArgs
func (t *ComplexType) GetRetuArgs() []*BaseValue {
	return nil
}

// GetDeclType GetDeclType
func (t *CustomizeType) GetDeclType() string {
	return "CustomizeType"
}

// GetRetuArgs GetRetuArgs
func (t *CustomizeType) GetRetuArgs() []*BaseValue {
	return nil
}

// GetTypeName GetTypeName
func (t *CustomizeType) GetTypeName() string {
	return t.TypeName
}

// SetTypeAndName SetTypeAndName
func (t *CustomizeType) SetTypeAndName(types string, name string) {
	t.Type = types
	t.TypeName = name
}

// GetDeclType GetDeclType
func (t *FuncType) GetDeclType() string {
	return "CustomizeType"
}

// SetTypeAndName SetTypeAndName
func (t *FuncType) SetTypeAndName(types string, name string) {
	t.Type = types
}

// GetTypeName GetTypeName
func (t *FuncType) GetTypeName() string {
	return t.Type
}

// GetRetuArgs GetRetuArgs
func (t *FuncType) GetRetuArgs() []*BaseValue {
	return t.RetuArgs
}

// GetDeclType GetDeclType
func (t *ReferType) GetDeclType() string {
	return "ReferType"
}

// SetTypeAndName SetTypeAndName
func (t *ReferType) SetTypeAndName(types string, name string) {
	t.Name = types
}

// GetTypeName GetTypeName
func (t *ReferType) GetTypeName() string {
	return t.Package + "." + t.Name
}

// GetRetuArgs GetRetuArgs
func (t *ReferType) GetRetuArgs() []*BaseValue {
	return nil
}
