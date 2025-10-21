package or

import "context"

func Or(channels ...<-chan interface{}) <-chan interface{} {
	newC := make(chan interface{})

	switch len(channels) {
	case 0:
		close(newC)
	case 1:
		go func() {
			defer close(newC)
			<-channels[0]
		}()
	default:
		go func() {
			defer close(newC)

			ctx, finish := context.WithCancel(context.Background())
			defer finish()

			out := make(chan struct{})
			for _, c := range channels {
				go func(ctx context.Context, in <-chan interface{}, out chan struct{}) {
					for {
						select {
						case <-ctx.Done():
							return
						case <-in:
							out <- struct{}{}
							return
						}
					}
				}(ctx, c, out)
			}
			<-out
			finish()
		}()
	}

	return newC
}
