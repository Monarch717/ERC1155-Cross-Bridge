package engine

func (e *Engine) chainID() string {
	return e.conf.ChainID.String()
}
