package mongo

import (
	"context"
	"fmt"
	"github.com/yorkie-team/codepair/backend-go/internal/database"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"log"
	"time"
)

// Client is a concrete implementation of the Database interface.
// It uses the official Client driver.
type Client struct {
	config *Config
	client *mongo.Client
}

// Dial creates an instance of Client and dials the given MongoDB.
func Dial(conf *Config) (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), conf.ParseConnectionTimeout())
	defer cancel()

	client, err := mongo.Connect(
		options.Client().
			ApplyURI(conf.ConnectionURI).
			SetConnectTimeout(conf.ParseConnectionTimeout()),
	)
	if err != nil {
		return nil, fmt.Errorf("connect to mongo: %w", err)
	}

	pingTimeout := conf.ParsePingTimeout()
	ctxPing, cancel := context.WithTimeout(ctx, pingTimeout)
	defer cancel()

	if err := client.Ping(ctxPing, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("ping mongo: %w", err)
	}

	if err := ensureIndexes(ctx, client.Database(conf.DatabaseName)); err != nil {
		return nil, err
	}

	log.Printf("MongoDB connected, URI: %s, DB: %s", conf.ConnectionURI, conf.DatabaseName)

	return &Client{
		config: conf,
		client: client,
	}, nil
}

//-----------------------------------------------
// UsersService-like queries
//-----------------------------------------------

// FindUserByID returns a user by its ID, or nil if not found.
func (m *Client) FindUserByID(ctx context.Context, userID string) (*database.User, error) {
	result := m.collection(ColUsers).FindOne(ctx, bson.M{"_id": userID})

	user := database.User{}
	if err := result.Decode(&user); err != nil {
		return nil, fmt.Errorf("find user by ID: %w", err)
	}

	return &user, nil
}

// FindUserBySocialUID checks if user exists with given provider & uid.
func (m *Client) FindUserBySocialUID(
	ctx context.Context,
	provider,
	uid string,
	createIfNotExist bool,
) (*database.User, error) {
	filter := bson.M{"social_provider": provider, "social_uid": uid}

	result := m.collection(ColUsers).FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments && createIfNotExist {
			return m.createUser(ctx, provider, uid)
		}
		return nil, fmt.Errorf("find user by social UID: %w", err)
	}

	user := database.User{}
	if err := result.Decode(&user); err != nil {
		return nil, fmt.Errorf("decode user: %w", err)
	}

	return &user, nil
}

// Ping is a fallback method to check Client connectivity.
func (m *Client) Ping(ctx context.Context) error {
	return m.client.Ping(ctx, nil)
}

// collection is a helper method to return a reference to a specific collection.
func (m *Client) collection(name string) *mongo.Collection {
	return m.client.Database(m.config.DatabaseName).Collection(name)
}

// createUser creates a new user with the given provider and uid.
func (m *Client) createUser(ctx context.Context, provider, uid string) (*database.User, error) {
	user := database.User{
		ID:             bson.NewObjectID(),
		SocialProvider: provider,
		SocialUID:      uid,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if _, err := m.collection(ColUsers).InsertOne(ctx, user); err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return &user, nil
}
