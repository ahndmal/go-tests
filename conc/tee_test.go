package conc

//func TestTeeChann(t *testing.T) {
//	tee := func(
//		done <-chan interface{},
//		in <-chan interface{},
//	) (_, _ <-chan interface{}) { <-chan interface{}) {
//out1 := make(chan interface{})
//out2 := make(chan interface{})
//go func() {defer close(out1)
//defer close(out2)
//for val := range orDone(done, in) {
//var out1, out2 = out1, out2
//for i := 0; i < 2; i++ {
//select {
//case <-done:
//case out1<-val:
//out1 = nil
//case out2<-val:
//out2 = nil
//}
//}
//}
//}()
//return out1, out2
//}
//}
