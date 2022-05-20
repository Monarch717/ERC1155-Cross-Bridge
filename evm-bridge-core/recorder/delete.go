package recorder

import "gorm.io/gorm"

func (r *Recorder) Delete(tx *gorm.DB, height int64) error {
	return nil
}
