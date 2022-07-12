package constant

// MethodBelongs2HasOneFind 一对一关联查找
// MethodBelongs2HasOneCreate 一对一关联创建
// MethodBelongs2HasOneUpdate 一对一关联更新
// MethodHasManyFind 一对多关联查找
// MethodHasManyCreate 一对多关联创建
// MethodHasManyUpdate 一对多关联更新
const (
	MethodBelongs2HasOneFind = iota + 1
	MethodBelongs2HasOneCreate
	MethodBelongs2HasOneUpdate
	MethodHasManyFind
	MethodHasManyCreate
	MethodHasManyUpdate
)
