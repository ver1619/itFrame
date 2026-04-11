package core

// RangeIterator produces a sequence of integers defined by start, end, and step.
type RangeIterator struct {
	current int
	end     int
	step    int
	done    bool
}

// Range creates a RangeIterator over [start, end) with the given step.
// Panics if step is zero.
func Range(start, end, step int) *RangeIterator {
	if step == 0 {
		panic("step cannot be zero")
	}

	if (step > 0 && start >= end) || (step < 0 && start <= end) {
		return &RangeIterator{
			current: start,
			end:     end,
			step:    step,
			done:    true,
		}
	}

	return &RangeIterator{
		current: start,
		end:     end,
		step:    step,
	}
}

// Next returns the next integer in the range, or (0, false) when exhausted.
func (r *RangeIterator) Next() (int, bool) {
	if r.done {
		return 0, false
	}

	if (r.step > 0 && r.current >= r.end) || (r.step < 0 && r.current <= r.end) {
		r.done = true
		return 0, false
	}

	val := r.current
	r.current += r.step
	return val, true
}

// Len returns the number of remaining elements.
func (r *RangeIterator) Len() int {
	if r.done {
		return 0
	}
	if r.step > 0 {
		if r.current >= r.end {
			return 0
		}
		return (r.end - r.current + r.step - 1) / r.step
	}
	if r.current <= r.end {
		return 0
	}
	return (r.current - r.end + (-r.step) - 1) / (-r.step)
}
