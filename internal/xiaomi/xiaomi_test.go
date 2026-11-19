package xiaomi

import "testing"

func TestMidrWakeAction(t *testing.T) {
	if siid, aiid := midrWakeAction("midr.cateye.sd400"); siid != 10 || aiid != 1 {
		t.Fatalf("SD400 action = %d/%d, want 10/1", siid, aiid)
	}
	if siid, aiid := midrWakeAction("midr.cateye.ph300"); siid != 2 || aiid != 3 {
		t.Fatalf("PH300 action = %d/%d, want 2/3", siid, aiid)
	}
}
