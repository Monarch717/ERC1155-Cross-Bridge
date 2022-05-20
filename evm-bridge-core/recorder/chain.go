package recorder

func (r *Recorder) ChainID() string {
	return r.conf.ChainID.String()
}

func (r *Recorder) ChainName() string {
	return r.conf.ChainName
}
