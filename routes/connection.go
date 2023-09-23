package routes

var ConnectionCount int = 0
var MaxConnections = 2
func ConnectionCheck() bool {
	return ConnectionCount >= MaxConnections
}
func ConnectionIncrement() {
	ConnectionCount++
}
func ConnectionDecrement() {
	ConnectionCount--
}
