package erc1155

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/block"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/util"
)

type SwapState string

type SwapDirection string

const (
	SwapStateRequestOngoing     SwapState = "request_ongoing"
	SwapStateRequestRejected    SwapState = "request_rejected"
	SwapStateRequestConfirmed   SwapState = "request_confirmed"
	SwapStateFillTxDryRunFailed SwapState = "fill_tx_dry_run_failed"
	SwapStateFillTxCreated      SwapState = "fill_tx_created"
	SwapStateFillTxSent         SwapState = "fill_tx_sent"
	SwapStateFillTxConfirmed    SwapState = "fill_tx_confirmed"
	SwapStateFillTxFailed       SwapState = "fill_tx_failed"
	SwapStateFillTxMissing      SwapState = "fill_tx_missing"

	SwapDirectionForward  SwapDirection = "forward"
	SwapDirectionBackward SwapDirection = "backward"
)

type Swap struct {
	ID string `gorm:"size:26;primary_key"`

	// Basic Token Information
	SrcChainID   string `gorm:"not null"`
	DstChainID   string `gorm:"not null"`
	SrcTokenAddr string `gorm:"not null"`
	DstTokenAddr string
	Sender       string         `gorm:"not null"`
	Recipient    string         `gorm:"not null"`
	IDs          datatypes.JSON `gorm:"not null"`
	Amounts      datatypes.JSON `gorm:"not null"`
	Signature    string         `gorm:"not null"`

	// Swap State
	State         SwapState     `gorm:"not null"`
	SwapDirection SwapDirection `gorm:"not null"`

	// Request Transaction Information
	RequestTxHash     string     `gorm:"not null"`
	RequestHeight     int64      `gorm:"not null"`
	RequestBlockHash  string     `gorm:"not null"`
	RequestBlockLogID *string    `gorm:"size:26;index:foreign_key_request_block_log_id"`
	RequestBlockLog   *block.Log `gorm:"foreignKey:RequestBlockLogID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	RequestTrackRetry int64

	// Fill Transaction Information
	FillConsumedFeeAmount string
	FillGasPrice          string
	FillGasUsed           int64
	FillHeight            int64
	FillTxHash            string
	FillTrackRetry        int64
	FillBlockHash         string     `gorm:"not null"`
	FillBlockLogID        *string    `gorm:"size:26;index:foreign_key_fill_block_log_id"`
	FillBlockLog          *block.Log `gorm:"foreignKey:FillBlockLogID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	MessageLog string

	// Timestamp
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Swap) TableName() string {
	return "erc1155_swaps"
}

func (s *Swap) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = util.ULID()
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
	return nil
}

func (s *Swap) BeforeUpdate(tx *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	return nil
}

func (s *Swap) IsRequiredInfoValid() bool {
	return s.DstTokenAddr != ""
}

func (s *Swap) SetRequiredInfo(dstTokenAddr string) {
	s.DstTokenAddr = dstTokenAddr
}

func (s *Swap) SignaturePayload() string {
	return fmt.Sprintf("%v#%v#%v#%v#%v#%v#%v#%v#%v#%v#%v#%v#%v",
		s.State,
		s.SrcChainID,
		s.DstChainID,
		s.SrcTokenAddr,
		s.DstTokenAddr,
		s.Sender,
		s.Recipient,
		s.IDs.String(),
		s.Amounts.String(),
		s.RequestTxHash,
		s.RequestHeight,
		s.FillTxHash,
		s.FillHeight,
	)
}

func (s *Swap) VerifySignature(hmacKey string) bool {
	oldSig := s.Signature
	s.UpdateSignature(hmacKey)
	newSig := s.Signature

	return oldSig == newSig
}

func (s *Swap) UpdateSignature(hmacKey string) {
	mac := hmac.New(sha256.New, []byte(hmacKey))
	mac.Write([]byte(s.SignaturePayload()))
	s.Signature = hex.EncodeToString(mac.Sum(nil))
}
