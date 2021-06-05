package db

import (
	"crypto/md5"
	"errors"
	"fmt"
	"time"

	"github.com/dgraph-io/badger/v3"
)

var (
	ErrExists    = errors.New("identifier already exists")
	ErrNotExists = errors.New("identifier not found")

	// Returned when the master password is incorrect
	ErrMasterPassword = errors.New("incorrect master password")
)

type DB struct {
	db *badger.DB
}

// Open a password database. If the master password is incorrect then returns ErrMasterPassword.
func Open(path string, masterPassword string) (*DB, error) {
	return open(path, masterPassword)
}

func open(path string, masterPassword string) (*DB, error) {
	// Run the masterPassword through a hash function to generate a more uniform
	// encryption key. This operation doesn't actually have any security benefits
	// it just makes the encryption key expected and consistent.
	key := md5.Sum([]byte(masterPassword))

	opts := badger.DefaultOptions(path)
	opts.EncryptionKey = key[:]
	opts.EncryptionKeyRotationDuration = time.Hour * 24 // rotate our encryption each day
	opts.IndexCacheSize = 1024
	opts.Logger = nil

	db, err := badger.Open(opts)
	if errors.Is(err, badger.ErrEncryptionKeyMismatch) { // If this is the case then the master password is incorrect
		return nil, ErrMasterPassword
	} else if err != nil {
		return nil, fmt.Errorf("opening database: %w", err)
	}
	return &DB{db}, nil
}

// AddPassword adds the password to the database for an ID. If the ID
// already exists then ErrExists will be returned
func (db *DB) AddPassword(id, password string) error {
	return db.db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte(id)); err == nil {
			return ErrExists
		} else if !errors.Is(err, badger.ErrKeyNotFound) {
			return fmt.Errorf("checking if password exists: %w", err)
		}

		if err := txn.Set([]byte(id), []byte(password)); err != nil {
			return fmt.Errorf("setting password: %w", err)
		}
		return nil
	})
}

// GetPassword gets the password from the database for an ID. If the ID
// does not exist then ErrNotExists will be returned
func (db *DB) GetPassword(id string) (string, error) {
	var password []byte
	err := db.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(id))
		if errors.Is(err, badger.ErrKeyNotFound) {
			return ErrNotExists
		} else if err != nil {
			return fmt.Errorf("checking if password exists: %w", err)
		}
		if password, err = item.ValueCopy(password); err != nil {
			return fmt.Errorf("retrieving password: %w", err)
		}
		return nil
	})
	return string(password), err
}
