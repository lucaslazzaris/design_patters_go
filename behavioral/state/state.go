package main

import "fmt"

type TCPOctetStream interface {
}

type TCPState interface {
	transmit(TCPConnector, TCPOctetStream)
	activeOpen(TCPConnector)
	passiveOpen(TCPConnector)
	close(TCPConnector)
	synchronize(TCPConnector)
	acknowledge(TCPConnector)
	send(TCPConnector)
	changeState(TCPConnector, TCPState)
	getStateName() string
}

type TCPClosed struct {
}

func (state *TCPClosed) transmit(connector TCPConnector, stream TCPOctetStream){
	connector.processOctet(stream)
}

func (state *TCPClosed) activeOpen(connector TCPConnector){
	// state.changeState(connector, TCPEstablished)	
}
func (state *TCPClosed) passiveOpen(connector TCPConnector){
	// state.changeState(connector, TCPListen)	
}
func (state *TCPClosed) close(connector TCPConnector){
	// state.changeState(connector, TCPListen)	
}
func (state *TCPClosed) synchronize(connector TCPConnector){
}
func (state *TCPClosed) acknowledge(connector TCPConnector){
}
func (state *TCPClosed) send(connector TCPConnector){
	// state.changeState(connector, TCPEstablished)	
}
func (state *TCPClosed) changeState(connector TCPConnector, newState TCPState){
	connector.changeState(newState)
}

func (state *TCPClosed) getStateName() string {
	return "TCPClosed"
}


// This is not correct according to the TCP protocol, it's just a dummy to test
type TCPEstablished struct {
}

func (state *TCPEstablished) transmit(connector TCPConnector, stream TCPOctetStream){
	connector.processOctet(stream)
}

func (state *TCPEstablished) activeOpen(connector TCPConnector){
	// state.changeState(connector, TCPEstablished)	
}
func (state *TCPEstablished) passiveOpen(connector TCPConnector){
	// state.changeState(connector, TCPListen)	
}
func (state *TCPEstablished) close(connector TCPConnector){
	// state.changeState(connector, TCPListen)	
}
func (state *TCPEstablished) synchronize(connector TCPConnector){
}
func (state *TCPEstablished) acknowledge(connector TCPConnector){
}
func (state *TCPEstablished) send(connector TCPConnector){
	// state.changeState(connector, TCPEstablished)	
}
func (state *TCPEstablished) changeState(connector TCPConnector, newState TCPState){
	connector.changeState(newState)
}

func (state *TCPEstablished) getStateName() string {
	return "TCPEstablished"
}

type TCPConnector interface {
	activeOpen()
	passiveOpen()
	close()
	send()
	acknowledge()
	synchronize()
	processOctet(TCPOctetStream)
	changeState(TCPState)
}

type TCPConnection struct {
	state TCPState
}

func newTCPConnection() TCPConnection{ 
	return TCPConnection{state: new(TCPClosed)}
}

func (connection *TCPConnection) changeState(state TCPState) {
	connection.state = state
}

func (connection *TCPConnection) activeOpen() {
	connection.state.activeOpen(connection)
}

func (connection *TCPConnection) passiveOpen() {
	connection.state.passiveOpen(connection)
}

func (connection *TCPConnection) close() {
	connection.state.close(connection)
}

func (connection *TCPConnection) acknowledge() {
	connection.state.acknowledge(connection)
}

func (connection *TCPConnection) synchronize() {
	connection.state.synchronize(connection)
}

func (connection *TCPConnection) send() {
	connection.state.send(connection)
}

func (connection *TCPConnection) processOctet(stream TCPOctetStream) {
	fmt.Printf("Processing Octet")
}
