package maria

import "time"

type Option func(*Maria)

func ConnsMaxLifeTime(time time.Duration) Option {
	return func(m *Maria) {
		m.ConnMaxLifeTime = time
	}
}

func MaxOpenConns(size int) Option {
	return func(m *Maria) {
		m.MaxOPenConns = size
	}
}

func MaxIdleConns(size int) Option {
	return func(m *Maria) {
		m.MaxIdleConns = size
	}
}