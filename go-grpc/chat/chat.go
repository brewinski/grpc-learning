package chat

import (
	context "context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DBRow struct {
	ProductType string `json:"product_type"`
}

func connect() *sql.DB {
	var host = "localhost"
	var port = "5432"
	var user = "postgres"
	var password = "postgres"
	var dbname = "velocity_lead_service"

	var psqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var result, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalf("DB Connection Error : %s", err)
	}

	return result
}

var DB = connect()

var store = []*Message{
	{Body: "Hello From the Server! Found"},
	{Body: "Hello From the Server! Found"},
	{Body: "Hello From the Server!"},
}

type Server struct {
	UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	// log.Printf("Received message body from client: %s", message.Body)
	store = append(store, message)
	return &Message{Body: "Hello From the Server!"}, nil
}

func (s *Server) ReadChatByID(ctx context.Context, message *GetMessageRequest) (*Message, error) {

	rows, err := DB.Query("SELECT product_type FROM \"velocity\".\"leads\" LIMIT 10;")
	if err != nil {
		log.Fatal("Error querying records", err)
	}
	defer rows.Close()

	messages := []*Message{}

	for rows.Next() {
		row := &DBRow{}
		err = rows.Scan(&row.ProductType)
		if err != nil {
			log.Fatal("Error scanning records", err)
		}

		messages = append(messages, &Message{Body: row.ProductType})
	}
	// log.Printf("Received message body from client: %s", message.Body)
	return messages[0], nil
}
