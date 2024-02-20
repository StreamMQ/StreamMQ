package main

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/eclipse/paho.mqtt.golang"
)


func TestOnMessageReceived(t *testing.T) {
	// Mock the SQL database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	sqlDB, _ := sql.Open("clickhouse", "mocked_dsn")


	// Set up expectations for the mock SQL database
	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO measurements \\(name, value\\) VALUES \\(\\?, \\?\\)")
	mock.ExpectExec("INSERT INTO measurements \\(name, value\\) VALUES \\(\\?, \\?\\)").
		WithArgs("test_name", 123).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()


	// Override the real SQL database with the mock database
	sqlOpen = func(driverName, dataSourceName string) (*sql.DB, error) {
		return sqlDB, nil
	}
	defer func() { sqlOpen = sql.Open }()

	// Create a mock MQTT message
    msg := NewMockMessage("test/topic", []byte(`[{"name": "measurement1", "value": 10000}]`))


	// Call the function under test
	onMessageReceived(nil, msg)

	// Verify the expectations
	assert.NoError(t, mock.ExpectationsWereMet())

}

// Mock mqttMessage struct to avoid dependencies
// mockMessage implements the mqtt.Message interface for testing purposes
type mockMessage struct {
    topic   string
    payload []byte
}

// NewMockMessage creates a new mockMessage with the specified topic and payload
func NewMockMessage(topic string, payload []byte) mqtt.Message {
    return &mockMessage{topic: topic, payload: payload}
}

// Qos returns the message QoS level
func (m *mockMessage) Qos() byte {
    return 0
}

// Retained returns true if the message is retained
func (m *mockMessage) Retained() bool {
    return false
}

// Duplicate returns true if the message is a duplicate
func (m *mockMessage) Duplicate() bool {
    return false
}

// Topic returns the message topic
func (m *mockMessage) Topic() string {
    return m.topic
}

// MessageID returns the message ID
func (m *mockMessage) MessageID() uint16 {
    return 0
}

// Payload returns the message payload
func (m *mockMessage) Payload() []byte {
    return m.payload
}

// Ack is not implemented in this mock and will return an error
func (m *mockMessage) Ack() {
    // No-op
}

// Mock sql.Open function to provide a mock database
var sqlOpen = sql.Open

func init() {
	sqlOpen = func(driverName, dataSourceName string) (*sql.DB, error) {
		return nil, nil
	}
}
