package mutexraceconditions

import "testing"

func Test_mutex(t *testing.T) {
	msg = "Hallå"

	wg.Add(2)
	go updateMessage("Hej", &mutex)
	go updateMessage("Hej då", &mutex)
	wg.Wait()

	if msg != "Hej" && msg != "Hej då" {
		t.Error("Incorrec value in msg")
	}
}
