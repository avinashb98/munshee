package config

import (
	"fmt"
	"os"
)

type Server struct {
	Port string
}

type Mongo struct {
	HostURI                string
	MunsheeDB              string
	UsersCollection        string
	AccountsCollection     string
	TransactionsCollection string
	TagsCollection         string
}

type Config struct {
	Server Server
	Mongo  Mongo
}

var config *Config = nil

func Get() Config {
	if config != nil {
		return *config
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	mongoConfig, err := getMongoConfig()
	if err != nil {
		panic(err)
	}

	config = &Config{
		Server: Server{
			Port: port,
		},
		Mongo: *mongoConfig,
	}
	return *config
}

func getMongoConfig() (*Mongo, error) {
	host := os.Getenv("MONGODB_HOST_URI")
	if host == "" {
		return nil, fmt.Errorf("missing mongodb host uri in env")
	}
	munsheeDb := os.Getenv("MONGODB_MUNSHEE_DB")
	if munsheeDb == "" {
		return nil, fmt.Errorf("missing munsheeDb name in env")
	}
	usersCollection := os.Getenv("MONGODB_USERS_COLLECTION")
	if usersCollection == "" {
		return nil, fmt.Errorf("missing users collection name in env")
	}

	accountsCollection := os.Getenv("MONGODB_ACCOUNTS_COLLECTION")
	if accountsCollection == "" {
		return nil, fmt.Errorf("missing accounts collection name in env")
	}
	transactionsCollection := os.Getenv("MONGODB_TRANSACTIONS_COLLECTION")
	if transactionsCollection == "" {
		return nil, fmt.Errorf("missing transactions collection name in env")
	}
	tagsCollection := os.Getenv("MONGODB_TAGS_COLLECTION")
	if tagsCollection == "" {
		return nil, fmt.Errorf("missing tags collection name in env")
	}

	return &Mongo{
		HostURI:                host,
		MunsheeDB:              munsheeDb,
		UsersCollection:        usersCollection,
		AccountsCollection:     accountsCollection,
		TransactionsCollection: transactionsCollection,
		TagsCollection:         tagsCollection,
	}, nil
}
