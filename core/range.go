package core

type RangeIterator struct {
	current int
	end     int
	step    int
	done    bool
}

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

/*
- **RangeIterator** generates a sequence of integers like a loop.
- **Range(start, end, step)** defines the sequence.
- step must not be 0 (will panic).

- Supports:
  - forward iteration (step > 0)
  - backward iteration (step < 0)

- If the range is invalid (e.g., start ≥ end with positive step), iteration ends immediately.
*/
