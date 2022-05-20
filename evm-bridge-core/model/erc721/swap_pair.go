package erc721

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/block"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/util"
)

type SwapPairState string

const (
	SwapPairStateRegistrationOngoing    SwapPairState = "registration_ongoing"
	SwapPairStateRegistrationConfirmed  SwapPairState = "registration_confirmed"
	SwapPairStateCreationTxDryRunFailed SwapPairState = "creation_tx_dry_run_failed"
	SwapPairStateCreationTxCreated      SwapPairState = "creation_tx_created"
	SwapPairStateCreationTxSent         SwapPairState = "creation_tx_sent"
	SwapPairStateCreationTxConfirmed    SwapPairState = "creation_tx_confirmed"
	SwapPairStateCreationTxFailed       SwapPairState = "creation_tx_failed"
	SwapPairStateCreationTxMissing      SwapPairState = "creation_tx_missing"
)

type SwapPair struct {
	ID string `gorm:"size:26;primary_key"`

	// Basic Token Information
	SrcChainID   string `gorm:"not null;index:unique_registration,unique,priority:1"`
	DstChainID   string `gorm:"not null;index:unique_registration,unique,priority:2"`
	SrcTokenAddr string `gorm:"not null;index:unique_registration,unique,priority:3"`
	DstTokenAddr string
	SrcTokenName string `gorm:"not null"`
	DstTokenName string
	Sponsor      string `gorm:"not null"`
	Available    bool   `gorm:"not null"`
	Signature    string `gorm:"not null"`
	Symbol       string `gorm:"not null"`
	BaseURI      string

	// Pair State
	State SwapPairState `gorm:"not null"`

	// Registration Transaction Information
	RegisterTxHash     string     `gorm:"not null;index:unique_registration,unique,priority:4"`
	RegisterHeight     int64      `gorm:"not null;index:unique_registration,unique,priority:5"`
	RegisterBlockHash  string     `gorm:"not null"`
	RegisterBlockLogID *string    `gorm:"size:26;index:foreign_key_register_block_log_id"`
	RegisterBlockLog   *block.Log `gorm:"foreignKey:RegisterBlockLogID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	// Creation Transaction Information
	CreateConsumedFeeAmount string
	CreateGasPrice          string
	CreateGasUsed           int64
	CreateHeight            int64
	CreateTxHash            string
	CreateTrackRetry        int64
	CreateBlockHash         string     `gorm:"not null"`
	CreateBlockLogID        *string    `gorm:"size:26;index:foreign_key_create_block_log_id"`
	CreateBlockLog          *block.Log `gorm:"foreignKey:CreateBlockLogID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	MessageLog string

	// Timestamp
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (SwapPair) TableName() string {
	return "erc721_swap_pairs"
}

func (s *SwapPair) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = util.ULID()
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
	return nil
}

func (s *SwapPair) BeforeUpdate(tx *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	return nil
}

func (s *SwapPair) SignaturePayload() string {
	return fmt.Sprintf("%v#%v#%v#%v#%v#%v#%v#%v#%v#%v#%v#%v",
		s.State,
		s.Symbol,
		s.SrcChainID,
		s.DstChainID,
		s.SrcTokenAddr,
		s.DstTokenAddr,
		s.SrcTokenName,
		s.DstTokenName,
		s.RegisterTxHash,
		s.RegisterHeight,
		s.CreateTxHash,
		s.CreateHeight,
	)
}

func (s *SwapPair) VerifySignature(hmacKey string) bool {
	oldSig := s.Signature
	s.UpdateSignature(hmacKey)
	newSig := s.Signature

	return oldSig == newSig
}

func (s *SwapPair) UpdateSignature(hmacKey string) {
	mac := hmac.New(sha256.New, []byte(hmacKey))
	mac.Write([]byte(s.SignaturePayload()))
	s.Signature = hex.EncodeToString(mac.Sum(nil))
}
