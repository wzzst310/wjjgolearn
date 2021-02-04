package iaa_pipe_filter

import "testing"

/**
 * @author  wjj
 * @date  2020/9/8 2:31 上午
 * @description
 */
func TestSumElems(t *testing.T) {
	sf := NewSumFilter()
	ret, err := sf.Process([]int{1, 2, 3})
	if err != nil {
		t.Fatal(err)
	}
	expected := 6
	if ret != expected {
		t.Fatalf("The expected is %d, but actual is %d", expected, ret)
	}
}

func TestWrongInputForSumFilter(t *testing.T) {
	sf := NewSumFilter()
	_, err := sf.Process([]float32{1.1, 2.2, 3.1})

	if err == nil {
		t.Fatal("An error is expected.")
	}

}