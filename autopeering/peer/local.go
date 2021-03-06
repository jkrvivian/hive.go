package peer

import (
	"crypto/ed25519"
	"fmt"
	"sync"

	"github.com/iotaledger/hive.go/autopeering/peer/service"
	"github.com/iotaledger/hive.go/autopeering/salt"
)

// Local defines the struct of a local peer
type Local struct {
	Peer
	key PrivateKey
	db  *DB

	// everything below is protected by a lock
	mu            sync.RWMutex
	serviceRecord *service.Record
	publicSalt    *salt.Salt
	privateSalt   *salt.Salt
}

// PrivateKey is the type of Ed25519 private keys used for the local peer.
type PrivateKey ed25519.PrivateKey

// Public returns the PublicKey corresponding to priv.
func (priv PrivateKey) Public() PublicKey {
	publicKey := ed25519.PrivateKey(priv).Public()
	return PublicKey(publicKey.(ed25519.PublicKey))
}

// newLocal creates a new local peer.
func newLocal(key PrivateKey, serviceRecord *service.Record, db *DB) *Local {
	return &Local{
		Peer:          *NewPeer(key.Public(), serviceRecord),
		key:           key,
		db:            db,
		serviceRecord: serviceRecord,
	}
}

// NewLocal creates a new local peer linked to the provided db.
// If an optional seed is provided, the seed is used to generate the private key. Without a seed,
// the provided key is loaded from the provided database and generated if not stored there.
func NewLocal(serviceRecord *service.Record, db *DB, seed ...[]byte) (*Local, error) {
	var key PrivateKey
	if len(seed) > 0 {
		key = PrivateKey(ed25519.NewKeyFromSeed(seed[0]))
		if db != nil {
			if err := db.UpdateLocalPrivateKey(key); err != nil {
				return nil, err
			}
		}
	} else {
		var err error
		key, err = db.LocalPrivateKey()
		if err != nil {
			return nil, err
		}
	}

	if l := len(key); l != ed25519.PrivateKeySize {
		return nil, fmt.Errorf("invalid key length: %d, need %d", l, ed25519.PrivateKeySize)
	}
	return newLocal(key, serviceRecord, db), nil
}

// Database returns the node database associated with the local peer.
func (l *Local) Database() *DB {
	return l.db
}

// Sign signs the message with the local peer's private key and returns a signature.
func (l *Local) Sign(message []byte) []byte {
	return ed25519.Sign(ed25519.PrivateKey(l.key), message)
}

// UpdateService updates the endpoint address of the given local service.
func (l *Local) UpdateService(key service.Key, network string, address string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// update the service in the read protected map
	l.serviceRecord.Update(key, network, address)

	// create a new peer with the corresponding services
	l.Peer = *NewPeer(l.key.Public(), l.serviceRecord)

	return nil
}

// GetPublicSalt returns the public salt
func (l *Local) GetPublicSalt() *salt.Salt {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.publicSalt
}

// SetPublicSalt sets the public salt
func (l *Local) SetPublicSalt(salt *salt.Salt) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.publicSalt = salt
}

// GetPrivateSalt returns the private salt
func (l *Local) GetPrivateSalt() *salt.Salt {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.privateSalt
}

// SetPrivateSalt sets the private salt
func (l *Local) SetPrivateSalt(salt *salt.Salt) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.privateSalt = salt
}

// generatePrivateKey generates a private key that can be used for Local.
func generatePrivateKey() (PrivateKey, error) {
	_, priv, err := ed25519.GenerateKey(nil)
	return PrivateKey(priv), err
}
